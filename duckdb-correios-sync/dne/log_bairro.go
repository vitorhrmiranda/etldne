package dne

import (
	"fmt"
	"path/filepath"
)

type LogBairro string

func (bairro LogBairro) Migrate(baseDir string) string {
	return fmt.Sprintf(`
    CREATE TABLE IF NOT EXISTS LOG_BAIRRO AS SELECT *
    FROM read_csv("%s", auto_detect=FALSE, sep='@', columns={'BAI_NU': INTEGER, 'UFE_SG': TEXT, 'LOC_NU': INTEGER, 'BAI_NO': TEXT, 'BAI_NO_ABREV': TEXT});
	`, filepath.Join(baseDir, string(bairro)))
}
