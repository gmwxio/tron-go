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
	text, err := svr.fileCache.get(req.TextDocument.URI)
	if err != nil {
		q.Q(err)
		return nil, nil
	}
	svr.astCache.get(req.TextDocument.URI)
	if req.Context.TriggerKind != protocol.Invoked {
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
	// cc := &complColl{
	// 	cl:   &protocol.CompletionList{},
	// 	col:  int(col),
	// 	line: int(line) + 1,
	// }
	ccl := &complCollList{
		cl:   &protocol.CompletionList{},
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
		q.Q(fmt.Sprintf("%s %#+[1]v", tn, x))
	}
	// adl.VisitADLP(ts, cc)
	// q.Q(cc.before, cc.after)
	// cl := protocol.CompletionList{
	// 	IsIncomplete: false,
	// 	Items: []protocol.CompletionItem{
	// 		{
	// 			Label:         "test",
	// 			Kind:          protocol.TextCompletion,
	// 			Detail:        "This is a test",
	// 			Documentation: "this is the docs",
	// 		},
	// 		{
	// 			Label:         "test2",
	// 			Kind:          protocol.TextCompletion,
	// 			Detail:        "This is a test2",
	// 			Documentation: "this is the docs2",
	// 		},
	// 	},
	// }
	return ccl.cl, nil
}

type complCollList struct {
	cl        *protocol.CompletionList
	line, col int
	treeNode  []ctree.TreeNode
	size      []int
	// before, after     antlr.TerminalNode
	// r_before, r_after antlr.RuleNode
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

		// tcount := tn.StopToken().GetTokenIndex() - tn.StartToken().GetTokenIndex()
		// debug := fmt.Sprintf("rule     #%d b=%d:%d a:%d:%d pos=%d:%d start=%v stop=%v",
		// 	tcount, b_line, b_col, a_line, a_col, v.line, v.col, tn.StartToken(), tn.StopToken())
		q.Q(debug)
		// debug_("   ---")
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

type complColl struct {
	*antlr.BaseParseTreeVisitor
	cl                *protocol.CompletionList
	line, col         int
	before, after     antlr.TerminalNode
	r_before, r_after antlr.RuleNode
}

func (v *complColl) VisitTerminal(node antlr.TerminalNode) {
	tok := node.GetSymbol()
	if v.after == nil {
		q.Q("terminal ", tok.GetLine(), tok.GetColumn(), v.col, v.line)
		if tok.GetLine() == v.line {
			if tok.GetColumn() >= v.col {
				v.after = node
			} else if v.after == nil {
				v.before = node
			} else {
				v.before = node
			}
		}
	}
}

// func (v *complColl) VisitErrorNode(node antlr.ErrorNode) {
// }
func (v *complColl) EnterEveryRule(ctx antlr.RuleNode) {
	tok := ctx.GetPayload().(antlr.Token)
	if v.after == nil {
		q.Q("rule ", tok.GetLine(), tok.GetColumn(), v.col, v.line)
		if tok.GetLine() == v.line && tok.GetColumn() >= v.col {
			v.r_after = ctx
		} else if v.r_after == nil {
			v.r_before = ctx
		} else {
			v.r_before = ctx
		}
	}
}

// func (v *complColl) ExitEveryRule(ctx antlr.RuleNode) {
//    v.indent = v.indent[:len(v.indent)-1]
// }

func (v *complColl) VisitAnnotation(ctx adlwi.IAnnotationContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {

	result = v.VisitChildren(ctx, delegate, args...)
	return
}
func (v *complColl) VisitModAnno(ctx adlwi.IModAnnoContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	result = v.VisitChildren(ctx, delegate, args...)
	return
}
func (v *complColl) VisitDeclAnno(ctx adlwi.IDeclAnnoContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	result = v.VisitChildren(ctx, delegate, args...)
	return
}
func (v *complColl) VisitFieldAnno(ctx adlwi.IFieldAnnoContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	result = v.VisitChildren(ctx, delegate, args...)
	return
}

func (v *complColl) VisitAdl(ctx adlwi.IAdlContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	result = v.VisitChildren(ctx, delegate, args...)
	return
}
func (v *complColl) VisitJson(ctx adlwi.IJsonContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	result = v.VisitChildren(ctx, delegate, args...)
	return
}
func (v *complColl) VisitModule(ctx adlwi.IModuleContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	n := ctx.GetTok().(ctree.TreeNode).Val().(adl.Module).Name
	result = v.VisitChildren(ctx, delegate, []string{n})
	return
}
func (v *complColl) VisitImportModule(ctx adlwi.IImportModuleContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	result = v.VisitChildren(ctx, delegate, args...)
	return
}
func (v *complColl) VisitImportScopedModule(ctx adlwi.IImportScopedModuleContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	result = v.VisitChildren(ctx, delegate, args...)
	return
}
func (v *complColl) VisitImportError(ctx adlwi.IImportErrorContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	result = v.VisitChildren(ctx, delegate, args...)
	return
}
func (v *complColl) VisitStruct(ctx adlwi.IStructContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	n := ctx.GetTok().(ctree.TreeNode).Val().(adl.Module).Name
	q.Q(n)
	result = v.VisitChildren(ctx, delegate, args...)
	return
}
func (v *complColl) VisitUnion(ctx adlwi.IUnionContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	result = v.VisitChildren(ctx, delegate, args...)
	return
}
func (v *complColl) VisitType(ctx adlwi.ITypeContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	result = v.VisitChildren(ctx, delegate, args...)
	return
}
func (v *complColl) VisitNewtype(ctx adlwi.INewtypeContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	result = v.VisitChildren(ctx, delegate, args...)
	return
}
func (v *complColl) VisitTypeParamError(ctx adlwi.ITypeParamErrorContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	result = v.VisitChildren(ctx, delegate, args...)
	return
}
func (v *complColl) VisitTLDError(ctx adlwi.ITLDErrorContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	result = v.VisitChildren(ctx, delegate, args...)
	return
}
func (v *complColl) VisitField(ctx adlwi.IFieldContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	result = v.VisitChildren(ctx, delegate, args...)
	return
}
func (v *complColl) VisitTypeExpr_(ctx adlwi.ITypeExpr_Context, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	result = v.VisitChildren(ctx, delegate, args...)
	return
}
func (v *complColl) VisitJsonStr(ctx adlwi.IJsonStrContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	result = v.VisitChildren(ctx, delegate, args...)
	return
}
func (v *complColl) VisitJsonBool(ctx adlwi.IJsonBoolContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	result = v.VisitChildren(ctx, delegate, args...)
	return
}
func (v *complColl) VisitJsonNull(ctx adlwi.IJsonNullContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	result = v.VisitChildren(ctx, delegate, args...)
	return
}
func (v *complColl) VisitJsonInt(ctx adlwi.IJsonIntContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	result = v.VisitChildren(ctx, delegate, args...)
	return
}
func (v *complColl) VisitJsonFloat(ctx adlwi.IJsonFloatContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	result = v.VisitChildren(ctx, delegate, args...)
	return
}
func (v *complColl) VisitJsonArray(ctx adlwi.IJsonArrayContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	result = v.VisitChildren(ctx, delegate, args...)
	return
}
func (v *complColl) VisitJsonObj(ctx adlwi.IJsonObjContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	result = v.VisitChildren(ctx, delegate, args...)
	return
}
func (v *complColl) VisitJsonError(ctx adlwi.IJsonErrorContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	result = v.VisitChildren(ctx, delegate, args...)
	return
}
