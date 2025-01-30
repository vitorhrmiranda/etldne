package dne

import (
	"fmt"
	"path/filepath"
)

type LogCPC string

func (cpc LogCPC) Migrate(baseDir string) string {
	return fmt.Sprintf(`
    CREATE TABLE IF NOT EXISTS LOG_CPC AS SELECT *
    FROM read_csv("%s", auto_detect=FALSE, sep='@', columns={'CPC_NU': INTEGER, 'UFE_SG': TEXT, 'LOC_NU': TEXT, 'CPC_NO': TEXT, 'CPC_ENDERECO': TEXT, 'CEP': TEXT});
	`, filepath.Join(baseDir, string(cpc)))
}
