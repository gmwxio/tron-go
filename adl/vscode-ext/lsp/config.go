package lsp

import (
	"context"
	"fmt"
	"path"

	"github.com/golangq/q"
	"golang.org/x/tools/lsp/protocol"
)

// Mirrors config in
// adl/vscode-ext/package.json

type TronCfg struct {
	*TronExtCfg
	*TronLangCfg
}

type TronExtCfg struct {
	ApplyPTComp  bool     `json:"autoApplyStructCompletions"`
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

func (svr *server) config() {
	ctx2 := context.Background()
	wfs, err := svr.client.WorkspaceFolders(ctx2)
	if err != nil {
		q.Q(err)
		// return err // TODO what does returning an error do
	}
	q.Q(wfs)
	q.Q(svr.initParams.RootURI)
	if len(wfs) == 0 {
		q.Q("empty WorkspaceFolder root:", svr.initParams.RootURI)
		if svr.initParams.RootURI != "" {
			wfs = []protocol.WorkspaceFolder{{
				URI:  svr.initParams.RootURI,
				Name: path.Base(svr.initParams.RootURI),
			}}
			q.Q(wfs)
		} else {
			// no folders and no root, single file mode
			//TODO(iancottrell): not sure how to do single file mode yet
			//issue: golang.org/issue/31168
			q.Q(fmt.Errorf("single file mode not supported yet"))
		}
	}
	svr.workspaceFolders = wfs

	var items []protocol.ConfigurationItem
	for _, wf := range wfs {
		q.Q(wf)
		items = append(items, []protocol.ConfigurationItem{
			{ScopeURI: wf.URI, Section: "tron"},
			{ScopeURI: wf.URI, Section: "[tron]"},
		}...,
		)
	}
	var xx interface{}
	svr.conn.Call(ctx2, "workspace/configuration", &protocol.ConfigurationParams{Items: items}, &xx)
	q.Q(xx)
	var result []TronCfg
	if err := svr.conn.Call(ctx2, "workspace/configuration", &protocol.ConfigurationParams{Items: items}, &result); err != nil {
		q.Q(err)
		err := svr.client.ShowMessage(ctx2, &protocol.ShowMessageParams{
			Message: fmt.Sprintf("Error reading config %v", err),
			Type:    protocol.Info,
		})
		if err != nil {
			q.Q(err)
			// return err // TODO what does returning an error do
		}

		return
	}
	for _, x := range result {
		if x.TronExtCfg != nil {
			svr.extConfig = *x.TronExtCfg
			q.Q(svr.extConfig.TronLspServe)
			q.Q(svr.extConfig.TronLspServe.Roots)
		}
		if x.TronLangCfg != nil {
			svr.langCfg = *x.TronLangCfg
			q.Q(svr.langCfg)
		}
	}

}
