// Generated from AdlWi.g4 by ANTLR 4.7.

package adlwi // AdlWi
// See base listener file for example implementations

//import "github.com/wxio/goantlr"

// import "generate package"

//type AdlWiVisitor struct {
//    *antlr.BaseParseTreeVisitor
//    indent string
//}

//var _ antlr.ParseTreeVisitor = &AdlWiVisitor{}
//var _ antlr.AggregateResultVisitor = &AdlWiVisitor{}
//var _ antlr.VisitNextCheck = &AdlWiVisitor{}
//var _ antlr.VisitRestCheck = &AdlWiVisitor{}
//var _ antlr.EnterEveryRuleVisitor = &AdlWiVisitor{}
//var _ antlr.ExitEveryRuleVisitor = &AdlWiVisitor{}

//var _ adlwi.AdlContextVisitor = &AdlWiVisitor{}
//var _ adlwi.JsonContextVisitor = &AdlWiVisitor{}
//var _ adlwi.ModuleContextVisitor = &AdlWiVisitor{}
//var _ adlwi.ImportModuleContextVisitor = &AdlWiVisitor{}
//var _ adlwi.ImportScopedModuleContextVisitor = &AdlWiVisitor{}
//var _ adlwi.ImportErrorContextVisitor = &AdlWiVisitor{}
//var _ adlwi.StructContextVisitor = &AdlWiVisitor{}
//var _ adlwi.UnionContextVisitor = &AdlWiVisitor{}
//var _ adlwi.TypeContextVisitor = &AdlWiVisitor{}
//var _ adlwi.NewtypeContextVisitor = &AdlWiVisitor{}
//var _ adlwi.ModAnnoContextVisitor = &AdlWiVisitor{}
//var _ adlwi.DeclAnnoContextVisitor = &AdlWiVisitor{}
//var _ adlwi.FieldAnnoContextVisitor = &AdlWiVisitor{}
//var _ adlwi.TypeParamErrorContextVisitor = &AdlWiVisitor{}
//var _ adlwi.TLDErrorContextVisitor = &AdlWiVisitor{}
//var _ adlwi.FieldContextVisitor = &AdlWiVisitor{}
//var _ adlwi.AnnotationContextVisitor = &AdlWiVisitor{}
//var _ adlwi.TypeExpr_ContextVisitor = &AdlWiVisitor{}
//var _ adlwi.TypeParamsContextVisitor = &AdlWiVisitor{}
//var _ adlwi.JsonStrContextVisitor = &AdlWiVisitor{}
//var _ adlwi.JsonBoolContextVisitor = &AdlWiVisitor{}
//var _ adlwi.JsonNullContextVisitor = &AdlWiVisitor{}
//var _ adlwi.JsonIntContextVisitor = &AdlWiVisitor{}
//var _ adlwi.JsonFloatContextVisitor = &AdlWiVisitor{}
//var _ adlwi.JsonArrayContextVisitor = &AdlWiVisitor{}
//var _ adlwi.JsonObjContextVisitor = &AdlWiVisitor{}
//var _ adlwi.JsonErrorContextVisitor = &AdlWiVisitor{}

//func (v *AdlWiVisitor) VisitNext(node antlr.Tree, current interface{}) bool {
//    return true
//}
//func (v *AdlWiVisitor) VisitRest(node antlr.RuleNode, current interface{}) bool {
//    return true
//}
//func (v *AdlWiVisitor) AggregateResult(aggregate, nextResult interface{}) (result interface{}) {
//    return nextResult
//}
//func (v *AdlWiVisitor) VisitTerminal(node antlr.TerminalNode) {
//}
//func (v *AdlWiVisitor) VisitErrorNode(node antlr.ErrorNode) {
//}
//func (v *AdlWiVisitor) EnterEveryRule(ctx antlr.RuleNode) {
//    v.indent += " "
//    fmt.Printf("%s %T\n", v.indent, ctx)
//}
//func (v *AdlWiVisitor) ExitEveryRule(ctx antlr.RuleNode) {
//    v.indent = v.indent[:len(v.indent)-1]
//}

//indent := ""
//parsetree.VisitFunc(&adlwi.AdlWiHandlers{
//EnterEveryRule: func(ctx antlr.RuleNode){indent += "\t"},
//ExitEveryRule:  func(ctx antlr.RuleNode){indent = indent[1:]},
//Adl: func(ctx *adlwi.AdlContext, this *adlwi.AdlWiHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//Json: func(ctx *adlwi.JsonContext, this *adlwi.AdlWiHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//Module: func(ctx *adlwi.ModuleContext, this *adlwi.AdlWiHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//ImportModule: func(ctx *adlwi.ImportModuleContext, this *adlwi.AdlWiHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//ImportScopedModule: func(ctx *adlwi.ImportScopedModuleContext, this *adlwi.AdlWiHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//ImportError: func(ctx *adlwi.ImportErrorContext, this *adlwi.AdlWiHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//Struct: func(ctx *adlwi.StructContext, this *adlwi.AdlWiHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//Union: func(ctx *adlwi.UnionContext, this *adlwi.AdlWiHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//Type: func(ctx *adlwi.TypeContext, this *adlwi.AdlWiHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//Newtype: func(ctx *adlwi.NewtypeContext, this *adlwi.AdlWiHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//ModAnno: func(ctx *adlwi.ModAnnoContext, this *adlwi.AdlWiHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//DeclAnno: func(ctx *adlwi.DeclAnnoContext, this *adlwi.AdlWiHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//FieldAnno: func(ctx *adlwi.FieldAnnoContext, this *adlwi.AdlWiHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//TypeParamError: func(ctx *adlwi.TypeParamErrorContext, this *adlwi.AdlWiHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//TLDError: func(ctx *adlwi.TLDErrorContext, this *adlwi.AdlWiHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//Field: func(ctx *adlwi.FieldContext, this *adlwi.AdlWiHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//Annotation: func(ctx *adlwi.AnnotationContext, this *adlwi.AdlWiHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//TypeExpr_: func(ctx *adlwi.TypeExpr_Context, this *adlwi.AdlWiHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//TypeParams: func(ctx *adlwi.TypeParamsContext, this *adlwi.AdlWiHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//JsonStr: func(ctx *adlwi.JsonStrContext, this *adlwi.AdlWiHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//JsonBool: func(ctx *adlwi.JsonBoolContext, this *adlwi.AdlWiHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//JsonNull: func(ctx *adlwi.JsonNullContext, this *adlwi.AdlWiHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//JsonInt: func(ctx *adlwi.JsonIntContext, this *adlwi.AdlWiHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//JsonFloat: func(ctx *adlwi.JsonFloatContext, this *adlwi.AdlWiHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//JsonArray: func(ctx *adlwi.JsonArrayContext, this *adlwi.AdlWiHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//JsonObj: func(ctx *adlwi.JsonObjContext, this *adlwi.AdlWiHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//JsonError: func(ctx *adlwi.JsonErrorContext, this *adlwi.AdlWiHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},

//}

//func (v *AdlWiVisitor) VisitAdl(ctx adlwi.IAdlContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *AdlWiVisitor) VisitJson(ctx adlwi.IJsonContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *AdlWiVisitor) VisitModule(ctx adlwi.IModuleContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *AdlWiVisitor) VisitImportModule(ctx adlwi.IImportModuleContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *AdlWiVisitor) VisitImportScopedModule(ctx adlwi.IImportScopedModuleContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *AdlWiVisitor) VisitImportError(ctx adlwi.IImportErrorContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *AdlWiVisitor) VisitStruct(ctx adlwi.IStructContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *AdlWiVisitor) VisitUnion(ctx adlwi.IUnionContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *AdlWiVisitor) VisitType(ctx adlwi.ITypeContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *AdlWiVisitor) VisitNewtype(ctx adlwi.INewtypeContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *AdlWiVisitor) VisitModAnno(ctx adlwi.IModAnnoContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *AdlWiVisitor) VisitDeclAnno(ctx adlwi.IDeclAnnoContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *AdlWiVisitor) VisitFieldAnno(ctx adlwi.IFieldAnnoContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *AdlWiVisitor) VisitTypeParamError(ctx adlwi.ITypeParamErrorContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *AdlWiVisitor) VisitTLDError(ctx adlwi.ITLDErrorContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *AdlWiVisitor) VisitField(ctx adlwi.IFieldContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *AdlWiVisitor) VisitAnnotation(ctx adlwi.IAnnotationContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *AdlWiVisitor) VisitTypeExpr_(ctx adlwi.ITypeExpr_Context, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *AdlWiVisitor) VisitTypeParams(ctx adlwi.ITypeParamsContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *AdlWiVisitor) VisitJsonStr(ctx adlwi.IJsonStrContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *AdlWiVisitor) VisitJsonBool(ctx adlwi.IJsonBoolContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *AdlWiVisitor) VisitJsonNull(ctx adlwi.IJsonNullContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *AdlWiVisitor) VisitJsonInt(ctx adlwi.IJsonIntContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *AdlWiVisitor) VisitJsonFloat(ctx adlwi.IJsonFloatContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *AdlWiVisitor) VisitJsonArray(ctx adlwi.IJsonArrayContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *AdlWiVisitor) VisitJsonObj(ctx adlwi.IJsonObjContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *AdlWiVisitor) VisitJsonError(ctx adlwi.IJsonErrorContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}

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
