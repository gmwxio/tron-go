package lsp

import (
	"context"

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
	cc := &complColl{
		cl:   &protocol.CompletionList{},
		col:  int(col),
		line: int(line),
	}
	adl.VisitADL(tr, cc)
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
	return cc.cl, nil
}

type complColl struct {
	*antlr.BaseParseTreeVisitor
	cl        *protocol.CompletionList
	line, col int
}

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
	n := ctx.GetTok().(*ctree.TreeNode).Val.(adl.Module).Name
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
	n := ctx.GetTok().(*ctree.TreeNode).Val.(adl.Module).Name
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
