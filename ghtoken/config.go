package ghtoken

type GhTokenProviderCfg struct {
	AppIdVarName     string
	AppInstIdVarName string
	AppPemKeyVarName string
	PatVarName       string
	GhApiUrlVarName  string
}

var DefaultGhTokenProviderCfg = GhTokenProviderCfg{
	AppIdVarName:     "GH_TKN_APP_ID",
	AppInstIdVarName: "GH_TKN_APP_INST_ID",
	AppPemKeyVarName: "GH_TKN_APP_PRIVATE_KEY",
	PatVarName:       "GH_TKN",
	GhApiUrlVarName:  "GH_TKN_API_URL",
}
