// Generated from AdlP.g4 by ANTLR 4.7.

package adllp // AdlP
// See base listener file for example implementations

//import "github.com/wxio/goantlr"

// import "generate package"

//type AdlPVisitor struct {
//    *antlr.BaseParseTreeVisitor
//    indent string
//}

//var _ antlr.ParseTreeVisitor = &AdlPVisitor{}
//var _ antlr.AggregateResultVisitor = &AdlPVisitor{}
//var _ antlr.VisitNextCheck = &AdlPVisitor{}
//var _ antlr.VisitRestCheck = &AdlPVisitor{}
//var _ antlr.EnterEveryRuleVisitor = &AdlPVisitor{}
//var _ antlr.ExitEveryRuleVisitor = &AdlPVisitor{}

//var _ adllp.AdlContextVisitor = &AdlPVisitor{}
//var _ adllp.ModuleStatementContextVisitor = &AdlPVisitor{}
//var _ adllp.ImportScopedNameContextVisitor = &AdlPVisitor{}
//var _ adllp.ImportModuleNameContextVisitor = &AdlPVisitor{}
//var _ adllp.LocalAnnoContextVisitor = &AdlPVisitor{}
//var _ adllp.DocAnnoContextVisitor = &AdlPVisitor{}
//var _ adllp.StructOrUnionContextVisitor = &AdlPVisitor{}
//var _ adllp.TypeOrNewtypeContextVisitor = &AdlPVisitor{}
//var _ adllp.ModuleAnnotationContextVisitor = &AdlPVisitor{}
//var _ adllp.DeclAnnotationContextVisitor = &AdlPVisitor{}
//var _ adllp.FieldAnnotationContextVisitor = &AdlPVisitor{}
//var _ adllp.TypeParameterContextVisitor = &AdlPVisitor{}
//var _ adllp.TypeExprPrimOrParamContextVisitor = &AdlPVisitor{}
//var _ adllp.TypeExprTypeExprContextVisitor = &AdlPVisitor{}
//var _ adllp.TypeExpressionElemContextVisitor = &AdlPVisitor{}
//var _ adllp.FieldStatementContextVisitor = &AdlPVisitor{}
//var _ adllp.StringStatementContextVisitor = &AdlPVisitor{}
//var _ adllp.TrueFalseNullContextVisitor = &AdlPVisitor{}
//var _ adllp.NumberStatementContextVisitor = &AdlPVisitor{}
//var _ adllp.FloatStatementContextVisitor = &AdlPVisitor{}
//var _ adllp.ArrayStatementContextVisitor = &AdlPVisitor{}
//var _ adllp.ObjStatementContextVisitor = &AdlPVisitor{}

//func (v *AdlPVisitor) VisitNext(node antlr.Tree, current interface{}) bool {
//    return true
//}
//func (v *AdlPVisitor) VisitRest(node antlr.RuleNode, current interface{}) bool {
//    return true
//}
//func (v *AdlPVisitor) AggregateResult(aggregate, nextResult interface{}) (result interface{}) {
//    return nextResult
//}
//func (v *AdlPVisitor) VisitTerminal(node antlr.TerminalNode) {
//}
//func (v *AdlPVisitor) VisitErrorNode(node antlr.ErrorNode) {
//}
//func (v *AdlPVisitor) EnterEveryRule(ctx antlr.RuleNode) {
//    v.indent += " "
//    fmt.Printf("%s %T\n", v.indent, ctx)
//}
//func (v *AdlPVisitor) ExitEveryRule(ctx antlr.RuleNode) {
//    v.indent = v.indent[:len(v.indent)-1]
//}

//indent := ""
//parsetree.VisitFunc(&adllp.AdlPHandlers{
//EnterEveryRule: func(ctx antlr.RuleNode){indent += "\t"},
//ExitEveryRule:  func(ctx antlr.RuleNode){indent = indent[1:]},
//Adl: func(ctx *adllp.AdlContext, this *adllp.AdlPHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//ModuleStatement: func(ctx *adllp.ModuleStatementContext, this *adllp.AdlPHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//ImportScopedName: func(ctx *adllp.ImportScopedNameContext, this *adllp.AdlPHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//ImportModuleName: func(ctx *adllp.ImportModuleNameContext, this *adllp.AdlPHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//LocalAnno: func(ctx *adllp.LocalAnnoContext, this *adllp.AdlPHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//DocAnno: func(ctx *adllp.DocAnnoContext, this *adllp.AdlPHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//StructOrUnion: func(ctx *adllp.StructOrUnionContext, this *adllp.AdlPHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//TypeOrNewtype: func(ctx *adllp.TypeOrNewtypeContext, this *adllp.AdlPHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//ModuleAnnotation: func(ctx *adllp.ModuleAnnotationContext, this *adllp.AdlPHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//DeclAnnotation: func(ctx *adllp.DeclAnnotationContext, this *adllp.AdlPHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//FieldAnnotation: func(ctx *adllp.FieldAnnotationContext, this *adllp.AdlPHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//TypeParameter: func(ctx *adllp.TypeParameterContext, this *adllp.AdlPHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//TypeExprPrimOrParam: func(ctx *adllp.TypeExprPrimOrParamContext, this *adllp.AdlPHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//TypeExprTypeExpr: func(ctx *adllp.TypeExprTypeExprContext, this *adllp.AdlPHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//TypeExpressionElem: func(ctx *adllp.TypeExpressionElemContext, this *adllp.AdlPHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//FieldStatement: func(ctx *adllp.FieldStatementContext, this *adllp.AdlPHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//StringStatement: func(ctx *adllp.StringStatementContext, this *adllp.AdlPHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//TrueFalseNull: func(ctx *adllp.TrueFalseNullContext, this *adllp.AdlPHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//NumberStatement: func(ctx *adllp.NumberStatementContext, this *adllp.AdlPHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//FloatStatement: func(ctx *adllp.FloatStatementContext, this *adllp.AdlPHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//ArrayStatement: func(ctx *adllp.ArrayStatementContext, this *adllp.AdlPHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//ObjStatement: func(ctx *adllp.ObjStatementContext, this *adllp.AdlPHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},

//}

//func (v *AdlPVisitor) VisitAdl(ctx adllp.IAdlContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *AdlPVisitor) VisitModuleStatement(ctx adllp.IModuleStatementContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *AdlPVisitor) VisitImportScopedName(ctx adllp.IImportScopedNameContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *AdlPVisitor) VisitImportModuleName(ctx adllp.IImportModuleNameContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *AdlPVisitor) VisitLocalAnno(ctx adllp.ILocalAnnoContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *AdlPVisitor) VisitDocAnno(ctx adllp.IDocAnnoContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *AdlPVisitor) VisitStructOrUnion(ctx adllp.IStructOrUnionContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *AdlPVisitor) VisitTypeOrNewtype(ctx adllp.ITypeOrNewtypeContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *AdlPVisitor) VisitModuleAnnotation(ctx adllp.IModuleAnnotationContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *AdlPVisitor) VisitDeclAnnotation(ctx adllp.IDeclAnnotationContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *AdlPVisitor) VisitFieldAnnotation(ctx adllp.IFieldAnnotationContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *AdlPVisitor) VisitTypeParameter(ctx adllp.ITypeParameterContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *AdlPVisitor) VisitTypeExprPrimOrParam(ctx adllp.ITypeExprPrimOrParamContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *AdlPVisitor) VisitTypeExprTypeExpr(ctx adllp.ITypeExprTypeExprContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *AdlPVisitor) VisitTypeExpressionElem(ctx adllp.ITypeExpressionElemContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *AdlPVisitor) VisitFieldStatement(ctx adllp.IFieldStatementContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *AdlPVisitor) VisitStringStatement(ctx adllp.IStringStatementContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *AdlPVisitor) VisitTrueFalseNull(ctx adllp.ITrueFalseNullContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *AdlPVisitor) VisitNumberStatement(ctx adllp.INumberStatementContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *AdlPVisitor) VisitFloatStatement(ctx adllp.IFloatStatementContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *AdlPVisitor) VisitArrayStatement(ctx adllp.IArrayStatementContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *AdlPVisitor) VisitObjStatement(ctx adllp.IObjStatementContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}

//  TODO list rules here
//  Visit rules manually
//  eg a : b c* | d;
//  if ctx.GetB() != nil {
//    result1 = ctx.GetB(ctx, delegate, args)
//    for _, c := range ctx.GetC() {
//      resultS = c.GetC(ctx, delegate, args)
//    }
//  } else { ... }
//  OR visit all children rules
//  // before children
//  v.VisitChildren(ctx, delegate)
//  // afer children
//
//  return result
