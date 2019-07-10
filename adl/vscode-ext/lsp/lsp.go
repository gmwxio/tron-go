package lsp

import (
	"context"
	"io/ioutil"
	"net"
	"os"
	"strings"

	"github.com/golangq/q"
	"golang.org/x/tools/jsonrpc2"
	"golang.org/x/tools/lsp/protocol"
)

type server struct {
	version          string
	adlcPath         string
	tempDir          string
	lastFileUri      string
	webAddr          net.Addr
	extConfig        TronExtCfg
	langCfg          TronLangCfg
	workspaceFolders []protocol.WorkspaceFolder
	workspaceFolder  protocol.WorkspaceFolder
	client           protocol.Client
	conn             *jsonrpc2.Conn
	initParams       *protocol.InitializeParams
	cancel           context.CancelFunc
	// tcpConn    net.Conn
	fileCache filecache
	astCache  astcache
}

func (svr *server) Initialize(ctx context.Context, req *protocol.InitializeParams) (*protocol.InitializeResult, error) {
	q.Q(svr.adlcPath)
	tempDir, err := ioutil.TempDir("", "adl-lsp")
	if err != nil {
		q.Q(err)
		return nil, err
	}
	svr.tempDir = tempDir
	q.Q(req)
	svr.fileCache = newCache()
	svr.astCache = newAstCache()
	svr.initParams = req
	// type ServerCapabilities struct {
	// 	InnerServerCapabilities
	// 	ImplementationServerCapabilities
	// 	TypeDefinitionServerCapabilities
	// 	WorkspaceFoldersServerCapabilities
	// 	ColorServerCapabilities
	// 	FoldingRangeServerCapabilities
	// 	DeclarationServerCapabilities
	// 	SelectionRangeServerCapabilities
	// }
	// q.Q(spew.Sdump(req))
	textDocumentSyncKind := protocol.Full
	ret := &protocol.InitializeResult{
		Capabilities: protocol.ServerCapabilities{
			InnerServerCapabilities: protocol.InnerServerCapabilities{
				CodeActionProvider: true,
				CompletionProvider: &protocol.CompletionOptions{
					TriggerCharacters: []string{"."},
				},
				DefinitionProvider:              true,
				DocumentFormattingProvider:      true,
				DocumentRangeFormattingProvider: true,
				DocumentSymbolProvider:          true,
				ExecuteCommandProvider: &protocol.ExecuteCommandOptions{
					Commands: []string{"tron.compile", "tron.browse", "tron.lsp.readconfig"},
				},
				HoverProvider:             true,
				DocumentHighlightProvider: true,
				SignatureHelpProvider: &protocol.SignatureHelpOptions{
					TriggerCharacters: []string{"(", ","},
				},
				TextDocumentSync: &protocol.TextDocumentSyncOptions{
					Change:    textDocumentSyncKind,
					OpenClose: true,
				},
			},
			TypeDefinitionServerCapabilities: protocol.TypeDefinitionServerCapabilities{
				TypeDefinitionProvider: true,
			},
			ImplementationServerCapabilities: protocol.ImplementationServerCapabilities{
				ImplementationProvider: true,
			},
			WorkspaceFoldersServerCapabilities: protocol.WorkspaceFoldersServerCapabilities{
				Workspace: &struct {
					WorkspaceFolders *struct {
						Supported           bool   `json:"supported,omitempty"`
						ChangeNotifications string `json:"changeNotifications,omitempty"`
					} `json:"workspaceFolders,omitempty"`
				}{
					WorkspaceFolders: &struct {
						Supported           bool   `json:"supported,omitempty"`
						ChangeNotifications string `json:"changeNotifications,omitempty"`
					}{
						Supported:           true,
						ChangeNotifications: "true",
					},
				},
			},
			DeclarationServerCapabilities: protocol.DeclarationServerCapabilities{
				DeclarationProvider: true,
			},
		},
	}

	// ret := &protocol.InitializeResult{
	// 	Capabilities: protocol.ServerCapabilities{},
	// 	Custom:       make(map[string]interface{}),
	// }
	return ret, nil
}
func (svr *server) Initialized(ctx context.Context, req *protocol.InitializedParams) error {
	q.Q(req)
	defer func() {
		svr.conn.Notify(ctx, "window/logMessage", protocol.LogMessageParams{
			Message: "TRON LSP (version " + svr.version + ")",
			Type:    protocol.Info,
		})
		err := svr.client.ShowMessage(ctx, &protocol.ShowMessageParams{
			Message: "Started TRON LSP (version " + svr.version + ")",
			Type:    protocol.Info,
		})
		if err != nil {
			q.Q(err)
			// return err // TODO what does returning an error do
		}
	}()

	// svr.initParams.Capabilities
	// Check if the client supports configuration messages.
	if x, ok := svr.initParams.Capabilities["workspace"].(map[string]interface{}); ok {
		if x, ok := x["configuration"].(bool); ok {
			q.Q("configurationSupported", x)
		}
		if x, ok := x["didChangeConfiguration"].(map[string]interface{}); ok {
			if x, ok := x["dynamicRegistration"].(bool); ok {
				q.Q("dynamicConfigurationSupported", x)
			}
		}
	}

	// err := svr.client.RegisterCapability(ctx, &protocol.RegistrationParams{
	// 	Registrations: []protocol.Registration{{
	// 		ID:     "1234567890",
	// 		Method: "workspace/didChangeConfiguration",
	// 	}},
	// })
	// if err != nil {
	// 	q.Q(err)
	// 	// return err // TODO what does returning an error do
	// }

	svr.config()
	svr.serverBrowser()

	return nil
}
func (svr *server) Shutdown(context.Context) error {
	q.Q("shutdown")
	svr.cancel()
	os.Exit(0)
	// svr.tcpConn.Close()
	return nil
}
func (svr *server) Exit(context.Context) error {
	q.Q("exit")
	return nil
}
func (svr *server) DidChangeWorkspaceFolders(ctx context.Context, req *protocol.DidChangeWorkspaceFoldersParams) error {
	q.Q(req)
	return nil
}
func (svr *server) DidChangeConfiguration(ctx context.Context, req *protocol.DidChangeConfigurationParams) error {
	q.Q(req)
	return nil
}
func (svr *server) DidChangeWatchedFiles(ctx context.Context, req *protocol.DidChangeWatchedFilesParams) error {
	q.Q(req)
	return nil
}
func (svr *server) Symbols(ctx context.Context, req *protocol.WorkspaceSymbolParams) ([]protocol.SymbolInformation, error) {
	q.Q(req)
	return nil, nil
}
func (svr *server) ExecuteCommand(ctx context.Context, req *protocol.ExecuteCommandParams) (interface{}, error) {
	q.Q(req)
	switch req.Command {
	case "tron.browse":
		if svr.lastFileUri != "" {
			err := svr.openbrowser()
			if err != nil {
				svr.client.ShowMessage(ctx, &protocol.ShowMessageParams{
					Message: "Error launching browser",
					Type:    protocol.Warning,
				})
			}
		} else {
			svr.client.ShowMessage(ctx, &protocol.ShowMessageParams{
				Message: "No file selected",
				Type:    protocol.Info,
			})
		}
	case "tron.lsp.readconfig":
		svr.config()
	case "tron.compile":
		if ok, _ := svr.lastFileStable(ctx, svr.client_msg_log); ok {
			cur, _ := svr.fileCache.get(svr.lastFileUri)
			allmod, err := svr.adl2ast(ctx, cur, svr.client_msg_log)
			if err != nil {
				q.Q(err)
				return nil, nil
			}
			svr.compile(ctx, cur, allmod, svr.client_msg_log)
		}
	default:
		svr.client.ShowMessage(ctx, &protocol.ShowMessageParams{
			Message: "unhandled command '" + req.Command + "'",
			Type:    protocol.Warning,
		})
	}

	return nil, nil
}
func (svr *server) DidOpen(ctx context.Context, req *protocol.DidOpenTextDocumentParams) error {
	q.Q(req)
	if strings.HasSuffix(req.TextDocument.URI, ".mod") {
		return nil
	}
	svr.fileCache.put(req.TextDocument.URI, req.TextDocument.Text)
	dss := svr.diag(ctx, req.TextDocument.Text)
	svr.client.PublishDiagnostics(ctx, &protocol.PublishDiagnosticsParams{
		Diagnostics: dss,
		URI:         req.TextDocument.URI,
	})
	//
	if len(dss) == 0 {
		cur, _ := svr.fileCache.get(svr.lastFileUri)
		allmod, err := svr.adl2ast(ctx, cur, svr.client_log)
		if err != nil {
			q.Q(err)
			return nil
		}
		svr.compile(ctx, cur, allmod, svr.client_log)
	}

	return nil
}
func (svr *server) DidChange(ctx context.Context, req *protocol.DidChangeTextDocumentParams) error {
	q.Q(req)
	if strings.HasSuffix(req.TextDocument.URI, ".mod") {
		return nil
	}
	if len(req.ContentChanges) < 1 {
		q.Q("no change")
		return nil
	}
	change := req.ContentChanges[0]
	svr.fileCache.put(req.TextDocument.URI, change.Text)
	dss := svr.diag(ctx, change.Text)
	svr.client.PublishDiagnostics(ctx, &protocol.PublishDiagnosticsParams{
		Diagnostics: dss,
		URI:         req.TextDocument.URI,
	})
	//
	if len(dss) == 0 {
		cur, _ := svr.fileCache.get(svr.lastFileUri)
		allmod, err := svr.adl2ast(ctx, cur, svr.client_log)
		if err != nil {
			q.Q(err)
			return nil
		}
		svr.compile(ctx, cur, allmod, svr.client_log)
	}

	return nil
}
func (svr *server) WillSave(ctx context.Context, req *protocol.WillSaveTextDocumentParams) error {
	q.Q(req)
	return nil
}
func (svr *server) WillSaveWaitUntil(ctx context.Context, req *protocol.WillSaveTextDocumentParams) ([]protocol.TextEdit, error) {
	q.Q(req)
	return nil, nil
}
func (svr *server) DidSave(ctx context.Context, req *protocol.DidSaveTextDocumentParams) error {
	q.Q(req)
	return nil
}
func (svr *server) DidClose(ctx context.Context, req *protocol.DidCloseTextDocumentParams) error {
	q.Q(req)
	svr.lastFileUri = req.TextDocument.URI
	if strings.HasSuffix(req.TextDocument.URI, ".mod") {
		return nil
	}
	svr.fileCache.delete(req.TextDocument.URI)
	svr.client.PublishDiagnostics(ctx, &protocol.PublishDiagnosticsParams{
		Diagnostics: []protocol.Diagnostic{},
		URI:         req.TextDocument.URI,
	})
	return nil
}
func (svr *server) Completion(ctx context.Context, req *protocol.CompletionParams) (*protocol.CompletionList, error) {
	q.Q(req)
	return svr.collect(ctx, req)
}

func (svr *server) CompletionResolve(ctx context.Context, req *protocol.CompletionItem) (*protocol.CompletionItem, error) {
	q.Q(req)
	return nil, nil
}
func (svr *server) Hover(ctx context.Context, req *protocol.TextDocumentPositionParams) (*protocol.Hover, error) {
	q.Q(req)
	return nil, nil
}
func (svr *server) SignatureHelp(ctx context.Context, req *protocol.TextDocumentPositionParams) (*protocol.SignatureHelp, error) {
	q.Q(req)
	return nil, nil
}
func (svr *server) Definition(ctx context.Context, req *protocol.TextDocumentPositionParams) ([]protocol.Location, error) {
	q.Q(req)
	return nil, nil
}
func (svr *server) TypeDefinition(ctx context.Context, req *protocol.TextDocumentPositionParams) ([]protocol.Location, error) {
	q.Q(req)
	return nil, nil
}
func (svr *server) Implementation(ctx context.Context, req *protocol.TextDocumentPositionParams) ([]protocol.Location, error) {
	q.Q(req)
	return nil, nil
}
func (svr *server) References(ctx context.Context, req *protocol.ReferenceParams) ([]protocol.Location, error) {
	q.Q(req)
	return nil, nil
}
func (svr *server) DocumentHighlight(ctx context.Context, req *protocol.TextDocumentPositionParams) ([]protocol.DocumentHighlight, error) {
	q.Q(req)
	return nil, nil
}
func (svr *server) DocumentSymbol(ctx context.Context, req *protocol.DocumentSymbolParams) ([]protocol.DocumentSymbol, error) {
	defer func() {
		if r := recover(); r != nil {
			q.Q(r)
			qstack()
		}
	}()
	q.Q(req)
	txt, err := svr.fileCache.get(req.TextDocument.URI)
	if err != nil {
		return []protocol.DocumentSymbol{}, nil
	}
	dsa, err := buildDocSym(txt, req.TextDocument.URI)
	if err != nil {
		q.Q(err)
		return []protocol.DocumentSymbol{}, nil
	}
	return dsa, nil
}

func (svr *server) CodeAction(ctx context.Context, req *protocol.CodeActionParams) ([]protocol.CodeAction, error) {
	q.Q(req)
	if strings.HasSuffix(req.TextDocument.URI, ".mod") {
		return nil, nil
	}
	svr.lastFileUri = req.TextDocument.URI
	return nil, nil
}
func (svr *server) CodeLens(ctx context.Context, req *protocol.CodeLensParams) ([]protocol.CodeLens, error) {
	q.Q(req)
	return nil, nil
}
func (svr *server) CodeLensResolve(ctx context.Context, req *protocol.CodeLens) (*protocol.CodeLens, error) {
	q.Q(req)
	return nil, nil
}
func (svr *server) DocumentLink(ctx context.Context, req *protocol.DocumentLinkParams) ([]protocol.DocumentLink, error) {
	q.Q(req)
	return nil, nil
}
func (svr *server) DocumentLinkResolve(ctx context.Context, req *protocol.DocumentLink) (*protocol.DocumentLink, error) {
	q.Q(req)
	return nil, nil
}
func (svr *server) DocumentColor(ctx context.Context, req *protocol.DocumentColorParams) ([]protocol.ColorInformation, error) {
	q.Q(req)
	return nil, nil
}
func (svr *server) ColorPresentation(ctx context.Context, req *protocol.ColorPresentationParams) ([]protocol.ColorPresentation, error) {
	q.Q(req)
	return nil, nil
}
func (svr *server) Formatting(ctx context.Context, req *protocol.DocumentFormattingParams) ([]protocol.TextEdit, error) {
	q.Q(req)
	return nil, nil
}
func (svr *server) RangeFormatting(ctx context.Context, req *protocol.DocumentRangeFormattingParams) ([]protocol.TextEdit, error) {
	q.Q(req)
	return nil, nil
}
func (svr *server) OnTypeFormatting(ctx context.Context, req *protocol.DocumentOnTypeFormattingParams) ([]protocol.TextEdit, error) {
	q.Q(req)
	return nil, nil
}
func (svr *server) Rename(ctx context.Context, req *protocol.RenameParams) ([]protocol.WorkspaceEdit, error) {
	q.Q(req)
	return nil, nil
}
func (svr *server) FoldingRanges(ctx context.Context, req *protocol.FoldingRangeParams) ([]protocol.FoldingRange, error) {
	q.Q(req)
	return nil, nil
}
