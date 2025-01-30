package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"correiosync/dne"

	_ "embed"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/marcboeker/go-duckdb"
)

//go:embed listar_enderecos.sql
var QueryListarEnderecos string

func main() {
	var canInsertIntoPostgres bool
	flag.BoolVar(&canInsertIntoPostgres, "postgres", false, "Insert data into Postgres")
	flag.Parse()

	start := time.Now()

	log.Println("running ETL...")
	defer func() {
		log.Printf("running ETL... done")
		log.Printf("elapsed time: %s", time.Since(start))
	}()

	dir := "./converted_files"
	db, err := openDuckDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := loadFilesIntoDB(db, dir); err != nil {
		log.Fatal(err)
	}

	if _, err := db.Exec(fmt.Sprintf("CREATE TABLE IF NOT EXISTS ZIPCODES AS (%s);", QueryListarEnderecos)); err != nil {
		log.Fatal(fmt.Errorf("create table ZIPCODES err: %w", err))
	}

	log.Println("zipcodes table created")
	if !canInsertIntoPostgres {
		return
	}

	log.Println("inserting data into Postgres...")
	defer log.Println("inserting data into Postgres... done!")
	if err := insertIntoPostgres(db); err != nil {
		log.Fatal(fmt.Errorf("insert into postgres err: %w", err))
	}
}

func loadFilesIntoDB(db *sql.DB, dir string) error {
	for _, entity := range dne.ListEntities() {
		if _, err := db.Exec(entity.Migrate(dir)); err != nil {
			return fmt.Errorf("migration %s err: %w", entity, err)
		}
	}
	return nil
}

func openDuckDB() (*sql.DB, error) {
	fileName := "./database/duck.db"
	if _, err := os.Stat(fileName); err == nil {
		_ = os.Remove(fileName)
	}

	db, err := sql.Open("duckdb", fileName)
	if err != nil {
		return nil, fmt.Errorf("open DuckDB err: %w", err)
	}
	return db, nil
}

func openPostgres() (*pgxpool.Pool, error) {
	poolConfig, err := pgxpool.ParseConfig("postgresql://postgres:postgres@localhost:5432/development")
	if err != nil {
		return nil, fmt.Errorf("parse Postgres err: %w", err)
	}
	db, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		return nil, fmt.Errorf("unable to create pg connection pool: %w", err)
	}
	return db, nil
}

func insertIntoPostgres(db *sql.DB) error {
	rows, err := db.Query("SELECT * FROM ZIPCODES")
	if err != nil {
		log.Fatal(fmt.Errorf("query list zipcodes err: %w", err))
	}
	defer rows.Close()

	pg, err := openPostgres()
	if err != nil {
		log.Fatal(err)
	}
	if _, err := pg.Exec(context.Background(), "TRUNCATE ZIPCODES"); err != nil {
		log.Fatal(fmt.Errorf("truncate zipcodes err: %w", err))
	}

	var (
		buffer []any
		mask   string
		count  int
		wg     sync.WaitGroup
	)
	// batchSize is max value supported to insert in Postgres
	const batchSize = 8191

	for rowNumber := 1; rows.Next(); rowNumber++ {

		var uf, cidade, bairro, logradouro, cep, nome, kind string
		var complemento sql.NullString

		if err := rows.Scan(&uf, &cidade, &bairro, &logradouro, &cep, &complemento, &nome, &kind); err != nil {
			log.Fatal(fmt.Errorf("scan row err: %w", err))
		}

		buffer = append(buffer, uf, cidade, bairro, logradouro, cep, complemento, nome, kind)
		mask += fmt.Sprintf(
			"($%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d),",
			rowNumber*8-7, rowNumber*8-6, rowNumber*8-5, rowNumber*8-4, rowNumber*8-3, rowNumber*8-2, rowNumber*8-1, rowNumber*8,
		)

		count++
		if (rowNumber % batchSize) != 0 {
			continue
		}

		wg.Add(1)
		go func(mask string, buffer []any) {
			defer wg.Done()
			fmt.Printf("\rprocessing row %d", count)

			if err := (PG{mask: mask, Pool: pg}).Insert(buffer...); err != nil {
				log.Fatal(fmt.Errorf("insert zipcodes err: %w", err))
			}
		}(mask, buffer)

		buffer, mask, rowNumber = []any{}, "", 0
	}

	wg.Wait()
	if len(buffer) == 0 {
		return nil
	}
	return (PG{mask: mask, Pool: pg}).Insert(buffer...)
}

type PG struct {
	*pgxpool.Pool
	mask string
}

func (pg PG) Insert(args ...any) error {
	const insertInto = "INSERT INTO ZIPCODES (uf, cidade, bairro, logradouro, cep, complemento, nome, kind) VALUES %s"
	_, err := pg.Exec(
		context.Background(),
		fmt.Sprintf(insertInto, pg.mask[:len(pg.mask)-1]),
		args...,
	)
	return err
}
