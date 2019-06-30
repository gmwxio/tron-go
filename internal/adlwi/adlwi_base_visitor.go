// Generated from AdlWi.g4 by ANTLR 4.7.

package adlw1 // AdlWi
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

//var _ adlw1.AdlContextVisitor = &AdlWiVisitor{}
//var _ adlw1.JsonContextVisitor = &AdlWiVisitor{}
//var _ adlw1.ModuleContextVisitor = &AdlWiVisitor{}
//var _ adlw1.ImportModuleContextVisitor = &AdlWiVisitor{}
//var _ adlw1.ImportScopedModuleContextVisitor = &AdlWiVisitor{}
//var _ adlw1.ImportErrorContextVisitor = &AdlWiVisitor{}
//var _ adlw1.StructContextVisitor = &AdlWiVisitor{}
//var _ adlw1.UnionContextVisitor = &AdlWiVisitor{}
//var _ adlw1.TypeContextVisitor = &AdlWiVisitor{}
//var _ adlw1.NewtypeContextVisitor = &AdlWiVisitor{}
//var _ adlw1.ModAnnoContextVisitor = &AdlWiVisitor{}
//var _ adlw1.DeclAnnoContextVisitor = &AdlWiVisitor{}
//var _ adlw1.FieldAnnoContextVisitor = &AdlWiVisitor{}
//var _ adlw1.TypeParamErrorContextVisitor = &AdlWiVisitor{}
//var _ adlw1.TLDErrorContextVisitor = &AdlWiVisitor{}
//var _ adlw1.FieldContextVisitor = &AdlWiVisitor{}
//var _ adlw1.AnnotationContextVisitor = &AdlWiVisitor{}
//var _ adlw1.TypeExpr_ContextVisitor = &AdlWiVisitor{}
//var _ adlw1.TypeParamsContextVisitor = &AdlWiVisitor{}
//var _ adlw1.JsonStrContextVisitor = &AdlWiVisitor{}
//var _ adlw1.JsonBoolContextVisitor = &AdlWiVisitor{}
//var _ adlw1.JsonNullContextVisitor = &AdlWiVisitor{}
//var _ adlw1.JsonIntContextVisitor = &AdlWiVisitor{}
//var _ adlw1.JsonFloatContextVisitor = &AdlWiVisitor{}
//var _ adlw1.JsonArrayContextVisitor = &AdlWiVisitor{}
//var _ adlw1.JsonObjContextVisitor = &AdlWiVisitor{}
//var _ adlw1.JsonErrorContextVisitor = &AdlWiVisitor{}

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
//parsetree.VisitFunc(&adlw1.AdlWiHandlers{
//EnterEveryRule: func(ctx antlr.RuleNode){indent += "\t"},
//ExitEveryRule:  func(ctx antlr.RuleNode){indent = indent[1:]},
//Adl: func(ctx *adlw1.AdlContext, this *adlw1.AdlWiHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//Json: func(ctx *adlw1.JsonContext, this *adlw1.AdlWiHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//Module: func(ctx *adlw1.ModuleContext, this *adlw1.AdlWiHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//ImportModule: func(ctx *adlw1.ImportModuleContext, this *adlw1.AdlWiHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//ImportScopedModule: func(ctx *adlw1.ImportScopedModuleContext, this *adlw1.AdlWiHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//ImportError: func(ctx *adlw1.ImportErrorContext, this *adlw1.AdlWiHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//Struct: func(ctx *adlw1.StructContext, this *adlw1.AdlWiHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//Union: func(ctx *adlw1.UnionContext, this *adlw1.AdlWiHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//Type: func(ctx *adlw1.TypeContext, this *adlw1.AdlWiHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//Newtype: func(ctx *adlw1.NewtypeContext, this *adlw1.AdlWiHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//ModAnno: func(ctx *adlw1.ModAnnoContext, this *adlw1.AdlWiHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//DeclAnno: func(ctx *adlw1.DeclAnnoContext, this *adlw1.AdlWiHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//FieldAnno: func(ctx *adlw1.FieldAnnoContext, this *adlw1.AdlWiHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//TypeParamError: func(ctx *adlw1.TypeParamErrorContext, this *adlw1.AdlWiHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//TLDError: func(ctx *adlw1.TLDErrorContext, this *adlw1.AdlWiHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//Field: func(ctx *adlw1.FieldContext, this *adlw1.AdlWiHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//Annotation: func(ctx *adlw1.AnnotationContext, this *adlw1.AdlWiHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//TypeExpr_: func(ctx *adlw1.TypeExpr_Context, this *adlw1.AdlWiHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//TypeParams: func(ctx *adlw1.TypeParamsContext, this *adlw1.AdlWiHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//JsonStr: func(ctx *adlw1.JsonStrContext, this *adlw1.AdlWiHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//JsonBool: func(ctx *adlw1.JsonBoolContext, this *adlw1.AdlWiHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//JsonNull: func(ctx *adlw1.JsonNullContext, this *adlw1.AdlWiHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//JsonInt: func(ctx *adlw1.JsonIntContext, this *adlw1.AdlWiHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//JsonFloat: func(ctx *adlw1.JsonFloatContext, this *adlw1.AdlWiHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//JsonArray: func(ctx *adlw1.JsonArrayContext, this *adlw1.AdlWiHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//JsonObj: func(ctx *adlw1.JsonObjContext, this *adlw1.AdlWiHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//JsonError: func(ctx *adlw1.JsonErrorContext, this *adlw1.AdlWiHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},

//}

//func (v *AdlWiVisitor) VisitAdl(ctx adlw1.IAdlContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *AdlWiVisitor) VisitJson(ctx adlw1.IJsonContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *AdlWiVisitor) VisitModule(ctx adlw1.IModuleContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *AdlWiVisitor) VisitImportModule(ctx adlw1.IImportModuleContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *AdlWiVisitor) VisitImportScopedModule(ctx adlw1.IImportScopedModuleContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *AdlWiVisitor) VisitImportError(ctx adlw1.IImportErrorContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *AdlWiVisitor) VisitStruct(ctx adlw1.IStructContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *AdlWiVisitor) VisitUnion(ctx adlw1.IUnionContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *AdlWiVisitor) VisitType(ctx adlw1.ITypeContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *AdlWiVisitor) VisitNewtype(ctx adlw1.INewtypeContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *AdlWiVisitor) VisitModAnno(ctx adlw1.IModAnnoContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *AdlWiVisitor) VisitDeclAnno(ctx adlw1.IDeclAnnoContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *AdlWiVisitor) VisitFieldAnno(ctx adlw1.IFieldAnnoContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *AdlWiVisitor) VisitTypeParamError(ctx adlw1.ITypeParamErrorContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *AdlWiVisitor) VisitTLDError(ctx adlw1.ITLDErrorContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *AdlWiVisitor) VisitField(ctx adlw1.IFieldContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *AdlWiVisitor) VisitAnnotation(ctx adlw1.IAnnotationContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *AdlWiVisitor) VisitTypeExpr_(ctx adlw1.ITypeExpr_Context, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *AdlWiVisitor) VisitTypeParams(ctx adlw1.ITypeParamsContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *AdlWiVisitor) VisitJsonStr(ctx adlw1.IJsonStrContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *AdlWiVisitor) VisitJsonBool(ctx adlw1.IJsonBoolContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *AdlWiVisitor) VisitJsonNull(ctx adlw1.IJsonNullContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *AdlWiVisitor) VisitJsonInt(ctx adlw1.IJsonIntContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *AdlWiVisitor) VisitJsonFloat(ctx adlw1.IJsonFloatContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *AdlWiVisitor) VisitJsonArray(ctx adlw1.IJsonArrayContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *AdlWiVisitor) VisitJsonObj(ctx adlw1.IJsonObjContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *AdlWiVisitor) VisitJsonError(ctx adlw1.IJsonErrorContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}

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
