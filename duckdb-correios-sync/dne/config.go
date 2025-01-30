package dne

type Migratable interface {
	Migrate(baseDir string) string
}

func ListEntities() []Migratable {
	return []Migratable{
		LogBairro("LOG_BAIRRO.TXT"),
		LogCPC("LOG_CPC.TXT"),
		LogGrandeUsuario("LOG_GRANDE_USUARIO.TXT"),
		LogLocalidade("LOG_LOCALIDADE.TXT"),
		LogLogradouro("LOG_LOGRADOURO_*.TXT"),
		LogUnidOper("LOG_UNID_OPER.TXT"),
	}
}
