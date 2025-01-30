SELECT
	LOG_LOGRADOURO.UFE_SG AS uf,
	LOG_LOCALIDADE.LOC_NO AS cidade,
	LOG_BAIRRO.BAI_NO AS bairro,
	LOG_LOGRADOURO.TLO_TX || ' ' || LOG_LOGRADOURO.LOG_NO AS logradouro,
	LOG_LOGRADOURO.CEP AS cep,
	LOG_LOGRADOURO.LOG_COMPLEMENTO AS complemento,
	'' AS nome,
	TLO_TX as kind
FROM
	LOG_LOGRADOURO,
	LOG_LOCALIDADE,
	LOG_BAIRRO
WHERE
	LOG_LOGRADOURO.LOC_NU = LOG_LOCALIDADE.LOC_NU
	AND LOG_LOGRADOURO.BAI_NU_INI = LOG_BAIRRO.BAI_NU
	AND LOG_LOGRADOURO.LOG_STA_TLO = 'S'
UNION
SELECT
	LOG_LOGRADOURO.UFE_SG AS uf,
	LOG_LOCALIDADE.LOC_NO AS cidade,
	LOG_BAIRRO.BAI_NO AS bairro,
	LOG_LOGRADOURO.LOG_NO AS logradouro,
	LOG_LOGRADOURO.CEP AS cep,
	LOG_LOGRADOURO.LOG_COMPLEMENTO AS complemento,
	'' AS nome,
	TLO_TX as kind
FROM
	LOG_LOGRADOURO,
	LOG_LOCALIDADE,
	LOG_BAIRRO
WHERE
	LOG_LOGRADOURO.LOC_NU = LOG_LOCALIDADE.LOC_NU
	AND LOG_LOGRADOURO.BAI_NU_INI = LOG_BAIRRO.BAI_NU
	AND LOG_LOGRADOURO.LOG_STA_TLO = 'N'
UNION
SELECT
	LOC.UFE_SG AS uf,
	LOC.LOC_NO AS cidade,
	'' AS bairro,
	'' AS logradouro,
	LOC.CEP AS cep,
	'' AS complemento,
	'' AS nome,
	CASE LOC_IN_TIPO_LOC
		WHEN 'D' THEN 'Distrito'
		WHEN 'M' THEN 'Município'
		WHEN 'P' THEN 'Povoado'
		ELSE ''
	END as kind
FROM
	LOG_LOCALIDADE AS LOC
WHERE
	LOC.CEP IS NOT NULL
	AND LOC.LOC_NU_SUB IS NULL
UNION
SELECT
	LOC.UFE_SG AS uf,
	LOCSUB.LOC_NO AS cidade,
	LOC.LOC_NO AS bairro,
	'' AS logradouro,
	LOC.CEP AS cep,
	'' AS complemento,
	'' AS nome,
	CASE COALESCE(NULLIF(LOC.LOC_IN_TIPO_LOC, ''), LOCSUB.LOC_IN_TIPO_LOC)
		WHEN 'D' THEN 'Distrito'
		WHEN 'M' THEN 'Município'
		WHEN 'P' THEN 'Povoado'
		ELSE ''
	END as kind
FROM
	LOG_LOCALIDADE AS LOC,
	LOG_LOCALIDADE AS LOCSUB
WHERE
	LOC.CEP IS NOT NULL
	AND LOC.LOC_NU_SUB IS NOT NULL
	AND LOC.LOC_NU_SUB = LOCSUB.LOC_NU
UNION
SELECT
	LOG_CPC.UFE_SG AS uf,
	LOG_LOCALIDADE.LOC_NO AS cidade,
	'' AS bairro,
	LOG_CPC.CPC_ENDERECO AS logradouro,
	LOG_CPC.CEP AS cep,
	'' AS complemento,
	CPC_NO AS nome,
	CASE LOC_IN_TIPO_LOC
		WHEN 'D' THEN 'Distrito'
		WHEN 'M' THEN 'Município'
		WHEN 'P' THEN 'Povoado'
		ELSE ''
	END as kind
FROM
	LOG_CPC,
	LOG_LOCALIDADE
WHERE
	LOG_CPC.LOC_NU = LOG_LOCALIDADE.LOC_NU
UNION
SELECT
	LOG_GRANDE_USUARIO.UFE_SG AS uf,
	LOG_LOCALIDADE.LOC_NO AS cidade,
	LOG_BAIRRO.BAI_NO AS bairro,
	LOG_GRANDE_USUARIO.GRU_ENDERECO AS logradouro,
	LOG_GRANDE_USUARIO.CEP AS cep,
	'' AS complemento,
	GRU_NO AS nome,
	CASE LOC_IN_TIPO_LOC
		WHEN 'D' THEN 'Distrito'
		WHEN 'M' THEN 'Município'
		WHEN 'P' THEN 'Povoado'
		ELSE ''
	END as kind
FROM
	LOG_GRANDE_USUARIO,
	LOG_LOCALIDADE,
	LOG_BAIRRO
WHERE
	LOG_GRANDE_USUARIO.LOC_NU = LOG_LOCALIDADE.LOC_NU
	AND LOG_GRANDE_USUARIO.BAI_NU = LOG_BAIRRO.BAI_NU
UNION
SELECT
	LOG_UNID_OPER.UFE_SG AS uf,
	LOG_LOCALIDADE.LOC_NO AS cidade,
	LOG_BAIRRO.BAI_NO AS bairro,
	LOG_UNID_OPER.UOP_ENDERECO AS logradouro,
	LOG_UNID_OPER.CEP AS cep,
	'' AS complemento,
	UOP_NO AS nome,
	CASE LOC_IN_TIPO_LOC
		WHEN 'D' THEN 'Distrito'
		WHEN 'M' THEN 'Município'
		WHEN 'P' THEN 'Povoado'
		ELSE ''
	END as kind
FROM
	LOG_UNID_OPER,
	LOG_LOCALIDADE,
	LOG_BAIRRO
WHERE
	LOG_UNID_OPER.LOC_NU = LOG_LOCALIDADE.LOC_NU
	AND LOG_UNID_OPER.BAI_NU = LOG_BAIRRO.BAI_NU
