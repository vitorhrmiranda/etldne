package dne

import (
	"fmt"
	"path/filepath"
)

type LogUnidOper string

func (unid LogUnidOper) Migrate(baseDir string) string {
	return fmt.Sprintf(`
    CREATE TABLE IF NOT EXISTS LOG_UNID_OPER AS SELECT *
    FROM read_csv("%s", auto_detect=FALSE, sep='@', columns={'UOP_NU': INTEGER, 'UFE_SG': TEXT, 'LOC_NU': TEXT, 'BAI_NU': INTEGER, 'LOG_NU': INTEGER, 'UOP_NO': TEXT, 'UOP_ENDERECO': TEXT, 'CEP': TEXT, 'UOP_IN_CP': TEXT, 'UOP_NO_ABREV': TEXT});
	`, filepath.Join(baseDir, string(unid)))
}
