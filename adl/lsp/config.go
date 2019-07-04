package lsp

// Mirrors config in
// adl/vscode-ext/package.json

type TronCfg struct {
	*TronExtCfg
	*TronLangCfg
}

type TronExtCfg struct {
	Includes     []string `json:"includes"`
	MaxIssues    float64  `json:"maxIssues"`
	AdlcPath     string   `json:"adlc.path"`
	TraceServer  string   `json:"trace.server"`
	TronLspServe struct {
		Port       int    `json:"port"`
		AdlAstPath string `json:"adlast.path"`
		Roots      []struct {
			Dir_Template string `json:"root"`
			Prefix       string `json:"prefix"`
		} `json:"roots"`
	} `json:"lspserve"`

	DevFeatures struct {
		Format       bool `json:"format"`
		AutoComplete bool `json:"autoComplete"`
	} `json:"devFeatures"`
}

type TronLangCfg struct {
	InsertSpaces     bool `json:"editor.insertSpaces"`
	FormatOnSave     bool `json:"editor.formatOnSave"`
	CodeActonsOnSave struct {
		OrganisImports bool `json:"source.organizeImports"`
	} `json:"editor.codeActionsOnSave"`
}
