package dne

import (
	"fmt"
	"path/filepath"
)

type LogLogradouro string

func (logr LogLogradouro) Migrate(baseDir string) string {
	return fmt.Sprintf(`
    CREATE TABLE IF NOT EXISTS LOG_LOGRADOURO AS SELECT *
    FROM read_csv("%s", auto_detect=FALSE, sep='@', columns={'LOG_NU': INTEGER, 'UFE_SG': TEXT, 'LOC_NU': TEXT, 'BAI_NU_INI': INTEGER, 'BAI_NU_FIM': INTEGER, 'LOG_NO': TEXT, 'LOG_COMPLEMENTO': TEXT, 'CEP': TEXT, 'TLO_TX': TEXT, 'LOG_STA_TLO': TEXT, 'LOG_NO_ABREV': TEXT});
	`, filepath.Join(baseDir, string(logr)))
}
