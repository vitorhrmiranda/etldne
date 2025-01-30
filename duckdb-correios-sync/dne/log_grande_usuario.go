package dne

import (
	"fmt"
	"path/filepath"
)

type LogGrandeUsuario string

func (gu LogGrandeUsuario) Migrate(baseDir string) string {
	return fmt.Sprintf(`
    CREATE TABLE IF NOT EXISTS LOG_GRANDE_USUARIO AS SELECT *
    FROM read_csv("%s", auto_detect=FALSE, sep='@', columns={'GRU_NU': INTEGER, 'UFE_SG': TEXT, 'LOC_NU': INTEGER, 'BAI_NU': INTEGER, 'LOG_NU': INTEGER, 'GRU_NO': TEXT, 'GRU_ENDERECO': TEXT, 'CEP': TEXT, 'GRU_NO_ABREV': TEXT});
	`, filepath.Join(baseDir, string(gu)))
}
