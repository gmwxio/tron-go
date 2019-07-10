package lsp

import (
	"context"
	"fmt"

	"github.com/wxio/tron-go/internal/ctree"

	"github.com/golangq/q"
	antlr "github.com/wxio/goantlr"
	"github.com/wxio/tron-go/adl"
	"github.com/wxio/tron-go/internal/adlwi"
	"golang.org/x/tools/lsp/protocol"
)

func (svr *server) collect(ctx context.Context, req *protocol.CompletionParams) (*protocol.CompletionList, error) {
	defer func() {
		if r := recover(); r != nil {
			q.Q(r)
			qstack()
		}
	}()
	if req.Context.TriggerKind != protocol.Invoked {
		return nil, nil
	}
	text, err := svr.fileCache.get(req.TextDocument.URI)
	if err != nil {
		q.Q(err)
		return nil, nil
	}
	col := req.TextDocumentPositionParams.Position.Character
	line := req.TextDocumentPositionParams.Position.Line
	tr, atr, bl, ts, err1 := adl.BuildAdlAST(text)
	_, _, _, _ = tr, atr, bl, ts
	// adl.QTreeToken(ts, bl)
	if err1.Error() != nil {
		q.Q(err1)
	}
	q.Q(col)
	q.Q(line + 1)
	q.Q(req.TextDocumentPositionParams.Position)
	ccl := &complCollList{
		col:  int(col),
		line: int(line) + 1,
	}
	pa, errs := adl.WalkADLWo(tr, ccl)
	if errs.Error() != nil {
		q.Q(errs.Error())
	}
	// antlr.ParseTreeWalkerDefault.Walk(ccl, atr)
	debug := fmt.Sprintf("%d %v", ccl.size, ccl.treeNode)
	q.Q(debug)

	for _, x := range ccl.treeNode {
		tn := pa.GetSymbolicNames()[x.GetTokenType()]
		q.Q(fmt.Sprintf("%s %T", tn, x.Val()))
	}
	//
	if len(ccl.treeNode) < 2 {
		return nil, nil
	}
	tns := ccl.treeNode
	// allmod, err := svr.adl2ast(ctx, text, svr.client_log)
	// if err != nil {
	// 	q.Q(err)
	// 	return nil, nil
	// }
	if svr.allmod == nil {
		q.Q("no adl cache")
		return nil, nil
	}
	mod, ex := svr.allmod[tns[1].Val().(adl.Module).Name]
	_ = ex // adl is valid
	cl := &protocol.CompletionList{}
	var curAnn *adl.ScopedName
	var curAnnMod adl.Module
	var curDeclType *adl.DeclType
	for i, tn := range tns[2:] {
		switch tn.GetTokenType() {
		case adlwi.AdlWiModuleAnno:
			// TODO all wild card imports
			for _, a := range mod.Imports {
				if a.ScopedName != nil && a.ScopedName.Name == tn.Val() {
					curAnn = a.ScopedName
					break
				}
			}
			if curAnn == nil {
				q.Q("didn't find", tn.Val())
				q.Q("in", mod.Annotations)
				q.Q("of", tns[1].Val().(adl.Module).Name, mod)
				return nil, nil
			}
			curAnnMod, ex = svr.allmod[curAnn.ModuleName]
			_ = ex // adl is valid
			curDecl := curAnnMod.Decls[curAnn.Name]
			curDeclType = &curDecl.Type
		case adlwi.AdlWiJsonObj:
			// curDeclType.
		case adlwi.AdlWiJsonObjKey:
		case adlwi.AdlWiJsonArray:
		case adlwi.AdlWiJsonBool:
		case adlwi.AdlWiJsonFloat:
		case adlwi.AdlWiJsonInt:
		case adlwi.AdlWiJsonNull:
		case adlwi.AdlWiJsonStr:
		default:
			q.Q("Unhandled completion", i, tn)
			svr.client_log(ctx, protocol.Warning, fmt.Sprintf("Unhandled completion %v", tn))
		}
	}
	if curDeclType == nil {
		q.Q("keh", tns[2:])
		return nil, nil
	}
	q.Q(curDeclType)
	ci := declTypeComplete(*curDeclType).Complete(svr.allmod, req.TextDocumentPositionParams.Position)
	if ci != nil {
		cl.Items = append(cl.Items, *ci)
	}
	// if edit != nil {
	// 	changes := map[string][]protocol.TextEdit{
	// 		req.TextDocument.URI: []protocol.TextEdit{
	// 			{
	// 				Range: protocol.Range{
	// 					Start: req.TextDocumentPositionParams.Position,
	// 					End:   req.TextDocumentPositionParams.Position,
	// 				},
	// 				NewText: *edit,
	// 			},
	// 		},
	// 	}
	// 	wep := &protocol.ApplyWorkspaceEditParams{
	// 		Label: "adl-completion",
	// 		Edit: protocol.WorkspaceEdit{
	// 			Changes: &changes,
	// 		},
	// 	}
	// 	svr.client.ApplyEdit(ctx, wep)
	// }
	return cl, nil
}

type declTypeComplete adl.DeclType

func (dt declTypeComplete) Complete(allmod map[string]adl.Module, pos protocol.Position) *protocol.CompletionItem {
	if dt.Union != nil {
		ci := &protocol.CompletionItem{
			Label:  "union",
			Kind:   protocol.TextCompletion,
			Detail: "this is an adl completions",
			// Documentation
			// Deprecated
			// Preselect
			// SortText
			// FilterText
			// InsertText
			// InsertTextFormat
			// TextEdit
			// AdditionalTextEdits
			// CommitCharacters
			// Command
			// Data
		}
		return ci
	}
	if dt.Struct != nil {
		te := protocol.TextEdit{
			Range: protocol.Range{
				Start: pos,
				End:   pos,
			},
			NewText: "struct goes here",
		}
		ci := &protocol.CompletionItem{
			Label:  "struct",
			Kind:   protocol.TextCompletion,
			Detail: "this is an adl completions",
			// Documentation
			// Deprecated
			// Preselect
			// SortText
			// FilterText
			// InsertText
			// InsertTextFormat
			TextEdit: &te,
			// AdditionalTextEdits
			// CommitCharacters
			// Command
			// Data
		}
		return ci
	}
	return nil
}

type complCollList struct {
	line, col int
	treeNode  []ctree.TreeNode
	size      []int
}

func debug_(x string) {
	q.Q(x)
}

func (v *complCollList) collection(tn ctree.TreeNode) {
	b_line, b_col := tn.StartToken().GetLine(), tn.StartToken().GetColumn()
	a_line, a_col := tn.StopToken().GetLine(), tn.StopToken().GetColumn()+len(tn.StopToken().GetText())
	tcount := tn.StopToken().GetTokenIndex() - tn.StartToken().GetTokenIndex()
	debug := fmt.Sprintf("#%d b=%d:%d a:%d:%d pos=%d:%d start=%v stop=%v",
		tcount, b_line, b_col, a_line, a_col, v.line, v.col, tn.StartToken(), tn.StopToken())
	// debug_(debug)
	if ((b_line == v.line && b_col < v.col) || b_line < v.line) &&
		((a_line == v.line && a_col >= v.col) || a_line > v.line) {
		q.Q(debug)
		v.size = append(v.size, tcount)
		v.treeNode = append(v.treeNode, tn)
	}
}

func (v *complCollList) VisitTerminal(node antlr.TerminalNode) {
	if tn, ok := node.GetSymbol().(ctree.TreeNode); ok {
		v.collection(tn)
	}
}

func (v *complCollList) VisitErrorNode(node antlr.ErrorNode) {}
func (v *complCollList) EnterEveryRule(ctx antlr.ParserRuleContext) {
	// if tn, ok := ctx.GetStart().(ctree.TreeNode); ok {
	// 	v.collection(tn)
	// }
}
func (v *complCollList) ExitEveryRule(ctx antlr.ParserRuleContext) {}
