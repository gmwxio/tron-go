package lsp

import (
	"fmt"

	"github.com/wxio/tron-go/internal/ctree"

	antlr "github.com/wxio/goantlr"
	"github.com/wxio/tron-go/internal/adlwo"
	"golang.org/x/tools/lsp/protocol"
)

type docSym struct {
	*antlr.BaseParseTreeVisitor
	dsa       *[]protocol.DocumentSymbol
	lexStream antlr.TokenStream
	furi      string
}

func (v *docSym) VisitAdl(ctx adlwo.IAdlContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	result = v.VisitChildren(ctx, delegate, args...)
	return
}
func (v *docSym) VisitModule(ctx adlwo.IModuleContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	// adl.QTreeToken(v.lexStream, ctx.GetParser())
	sym := ctx.Module().GetSymbol().(*ctree.TreeNode)
	sidx := sym.GetStop()
	// q.Q(sidx)
	stok := v.lexStream.Get(sidx)
	// q.Q(fmt.Sprintf("%#+v", sym))
	// // q.Q(v.lexStream)
	// // var stok antlr.Token
	// // stok := v.lexStream.Get(sidx)
	// // if sidx == -1 {
	// // 	stok = ctx.GetParser().GetTokenStream().Get
	// // } else {
	// // 	stok = ctx.GetParser().GetTokenStream().Get(sidx)
	// // }

	// ds := protocol.DocumentSymbol{
	// 	Name: ctx.Module().GetText(),
	// 	Kind: protocol.Package,
	// 	Range: protocol.Range{
	// 		Start: protocol.Position{
	// 			Line:      float64(ctx.GetStart().GetLine() - 1),
	// 			Character: float64(ctx.GetStart().GetColumn() + 1),
	// 		},
	// 		End: protocol.Position{
	// 			Line:      float64(stok.GetLine() - 1),
	// 			Character: float64(stok.GetColumn() + 1),
	// 		},
	// 	},
	// }

	start := protocol.Position{
		Line:      float64(ctx.GetStart().GetLine() - 1),
		Character: float64(ctx.GetStart().GetColumn() + 1),
	}
	end := protocol.Position{
		Line:      float64(stok.GetLine() - 1),
		Character: float64(stok.GetColumn()) + 1,
	}
	ds := &protocol.DocumentSymbol{
		Detail: fmt.Sprintf("%s", ctx.Module().GetText()),
		Name:   fmt.Sprintf("%v", ctx.Module().GetPayload()),
		Kind:   protocol.Module,
		Location: protocol.Location{
			Range: protocol.Range{
				Start: start,
				End:   end,
			},
			URI: v.furi,
		},
		Range: protocol.Range{
			Start: start,
			End:   end,
		},
	}

	// q.Q(ds)
	ds.Children = make([]protocol.DocumentSymbol, 0)
	// result = v.VisitChildren(ctx, delegate, &ds.Children)
	result = v.VisitChildren(ctx, delegate, v.dsa)
	*v.dsa = append(*v.dsa, *ds)
	return
}
func (v *docSym) VisitImportModule(ctx adlwo.IImportModuleContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	result = v.VisitChildren(ctx, delegate, args...)
	return
}
func (v *docSym) VisitImportScopedModule(ctx adlwo.IImportScopedModuleContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	result = v.VisitChildren(ctx, delegate, args...)
	return
}
func (v *docSym) VisitStruct(ctx adlwo.IStructContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	// q.Q("struct")
	arr := args[0].(*[]protocol.DocumentSymbol)
	sym := ctx.Struct().GetSymbol().(*ctree.TreeNode)
	sidx := sym.GetStop()
	stok := v.lexStream.Get(sidx)
	start := protocol.Position{
		Line:      float64(ctx.GetStart().GetLine() - 1),
		Character: float64(ctx.GetStart().GetColumn() + 1),
	}
	end := protocol.Position{
		Line:      float64(stok.GetLine() - 1),
		Character: float64(stok.GetColumn() + 1),
	}
	// p := fmt.Sprintf("%T", ctx.Struct().GetPayload())
	// q.Q(p)
	ds := protocol.DocumentSymbol{
		Detail: fmt.Sprintf("%s", ctx.Struct().GetText()),
		Name:   fmt.Sprintf("%v", ctx.Struct().GetPayload()),
		Kind:   protocol.Struct,
		Location: protocol.Location{
			Range: protocol.Range{
				Start: start,
				End:   end,
			},
			URI: v.furi,
		},
		Range: protocol.Range{
			Start: start,
			End:   end,
		},
	}
	// x := fmt.Sprintf("%#+v", ds)
	// q.Q(arr)
	*arr = append(*arr, ds)
	ds.Children = make([]protocol.DocumentSymbol, 0)
	result = v.VisitChildren(ctx, delegate, &ds.Children)
	return
}
func (v *docSym) VisitUnion(ctx adlwo.IUnionContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	result = v.VisitChildren(ctx, delegate, args...)
	return
}
func (v *docSym) VisitType(ctx adlwo.ITypeContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	result = v.VisitChildren(ctx, delegate, args...)
	return
}
func (v *docSym) VisitNewtype(ctx adlwo.INewtypeContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	result = v.VisitChildren(ctx, delegate, args...)
	return
}
func (v *docSym) VisitModAnno(ctx adlwo.IModAnnoContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	result = v.VisitChildren(ctx, delegate, args...)
	return
}
func (v *docSym) VisitDeclAnno(ctx adlwo.IDeclAnnoContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	result = v.VisitChildren(ctx, delegate, args...)
	return
}
func (v *docSym) VisitFieldAnno(ctx adlwo.IFieldAnnoContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	result = v.VisitChildren(ctx, delegate, args...)
	return
}
func (v *docSym) VisitField(ctx adlwo.IFieldContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	result = v.VisitChildren(ctx, delegate, args...)
	return
}
func (v *docSym) VisitAnnotation(ctx adlwo.IAnnotationContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	result = v.VisitChildren(ctx, delegate, args...)
	return
}
func (v *docSym) VisitTypeExpr_(ctx adlwo.ITypeExpr_Context, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	result = v.VisitChildren(ctx, delegate, args...)
	return
}
func (v *docSym) VisitTypeParams(ctx adlwo.ITypeParamsContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	result = v.VisitChildren(ctx, delegate, args...)
	return
}
func (v *docSym) VisitJson(ctx adlwo.IJsonContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	result = v.VisitChildren(ctx, delegate, args...)
	return
}
func (v *docSym) VisitJsonStr(ctx adlwo.IJsonStrContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	result = v.VisitChildren(ctx, delegate, args...)
	return
}
func (v *docSym) VisitJsonBool(ctx adlwo.IJsonBoolContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	result = v.VisitChildren(ctx, delegate, args...)
	return
}
func (v *docSym) VisitJsonNull(ctx adlwo.IJsonNullContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	result = v.VisitChildren(ctx, delegate, args...)
	return
}
func (v *docSym) VisitJsonInt(ctx adlwo.IJsonIntContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	result = v.VisitChildren(ctx, delegate, args...)
	return
}
func (v *docSym) VisitJsonFloat(ctx adlwo.IJsonFloatContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	result = v.VisitChildren(ctx, delegate, args...)
	return
}
func (v *docSym) VisitJsonArray(ctx adlwo.IJsonArrayContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	result = v.VisitChildren(ctx, delegate, args...)
	return
}
func (v *docSym) VisitJsonObj(ctx adlwo.IJsonObjContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	result = v.VisitChildren(ctx, delegate, args...)
	return
}
