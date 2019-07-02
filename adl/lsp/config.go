package lsp

// Mirrors config in
// adl/vscode-ext/package.json

type TronCfg struct {
	*TronExtCfg
	*TronLangCfg
}

type TronExtCfg struct {
	Includes    []string `json:"includes"`
	MaxIssues   float64  `json:"maxIssues"`
	TraceServer string   `json:"trace.server"`
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
