package dne

import (
	"fmt"
	"path/filepath"
)

type LogLocalidade string

func (loc LogLocalidade) Migrate(baseDir string) string {
	return fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS LOG_LOCALIDADE AS SELECT *
    FROM read_csv("%s", auto_detect=FALSE, sep='@', columns={'LOC_NU': INTEGER, 'UFE_SG': TEXT, 'LOC_NO': TEXT, 'CEP': TEXT, 'LOC_IN_SIT': TEXT, 'LOC_IN_TIPO_LOC': TEXT, 'LOC_NU_SUB': INTEGER, 'LOC_NO_ABREV': TEXT, 'MUN_NU': TEXT});
`, filepath.Join(baseDir, string(loc)))
}
