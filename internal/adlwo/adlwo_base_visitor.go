// Generated from AdlWo.g4 by ANTLR 4.7.

package adlwo // AdlWo
// See base listener file for example implementations

//import "github.com/wxio/goantlr"

// import "generate package"

//type AdlWoVisitor struct {
//    *antlr.BaseParseTreeVisitor
//    indent string
//}

//var _ antlr.ParseTreeVisitor = &AdlWoVisitor{}
//var _ antlr.AggregateResultVisitor = &AdlWoVisitor{}
//var _ antlr.VisitNextCheck = &AdlWoVisitor{}
//var _ antlr.VisitRestCheck = &AdlWoVisitor{}
//var _ antlr.EnterEveryRuleVisitor = &AdlWoVisitor{}
//var _ antlr.ExitEveryRuleVisitor = &AdlWoVisitor{}

//var _ adlwo.AdlContextVisitor = &AdlWoVisitor{}
//var _ adlwo.JsonContextVisitor = &AdlWoVisitor{}
//var _ adlwo.ModuleContextVisitor = &AdlWoVisitor{}
//var _ adlwo.ImportModuleContextVisitor = &AdlWoVisitor{}
//var _ adlwo.ImportScopedModuleContextVisitor = &AdlWoVisitor{}
//var _ adlwo.StructContextVisitor = &AdlWoVisitor{}
//var _ adlwo.UnionContextVisitor = &AdlWoVisitor{}
//var _ adlwo.TypeContextVisitor = &AdlWoVisitor{}
//var _ adlwo.NewtypeContextVisitor = &AdlWoVisitor{}
//var _ adlwo.ModAnnoContextVisitor = &AdlWoVisitor{}
//var _ adlwo.DeclAnnoContextVisitor = &AdlWoVisitor{}
//var _ adlwo.FieldAnnoContextVisitor = &AdlWoVisitor{}
//var _ adlwo.FieldContextVisitor = &AdlWoVisitor{}
//var _ adlwo.AnnotationContextVisitor = &AdlWoVisitor{}
//var _ adlwo.TypeExpr_ContextVisitor = &AdlWoVisitor{}
//var _ adlwo.TypeParamsContextVisitor = &AdlWoVisitor{}
//var _ adlwo.JsonStrContextVisitor = &AdlWoVisitor{}
//var _ adlwo.JsonBoolContextVisitor = &AdlWoVisitor{}
//var _ adlwo.JsonNullContextVisitor = &AdlWoVisitor{}
//var _ adlwo.JsonIntContextVisitor = &AdlWoVisitor{}
//var _ adlwo.JsonFloatContextVisitor = &AdlWoVisitor{}
//var _ adlwo.JsonArrayContextVisitor = &AdlWoVisitor{}
//var _ adlwo.JsonObjContextVisitor = &AdlWoVisitor{}

//func (v *AdlWoVisitor) VisitNext(node antlr.Tree, current interface{}) bool {
//    return true
//}
//func (v *AdlWoVisitor) VisitRest(node antlr.RuleNode, current interface{}) bool {
//    return true
//}
//func (v *AdlWoVisitor) AggregateResult(aggregate, nextResult interface{}) (result interface{}) {
//    return nextResult
//}
//func (v *AdlWoVisitor) VisitTerminal(node antlr.TerminalNode) {
//}
//func (v *AdlWoVisitor) VisitErrorNode(node antlr.ErrorNode) {
//}
//func (v *AdlWoVisitor) EnterEveryRule(ctx antlr.RuleNode) {
//    v.indent += " "
//    fmt.Printf("%s %T\n", v.indent, ctx)
//}
//func (v *AdlWoVisitor) ExitEveryRule(ctx antlr.RuleNode) {
//    v.indent = v.indent[:len(v.indent)-1]
//}

//indent := ""
//parsetree.VisitFunc(&adlwo.AdlWoHandlers{
//EnterEveryRule: func(ctx antlr.RuleNode){indent += "\t"},
//ExitEveryRule:  func(ctx antlr.RuleNode){indent = indent[1:]},
//Adl: func(ctx *adlwo.AdlContext, this *adlwo.AdlWoHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//Json: func(ctx *adlwo.JsonContext, this *adlwo.AdlWoHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//Module: func(ctx *adlwo.ModuleContext, this *adlwo.AdlWoHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//ImportModule: func(ctx *adlwo.ImportModuleContext, this *adlwo.AdlWoHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//ImportScopedModule: func(ctx *adlwo.ImportScopedModuleContext, this *adlwo.AdlWoHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//Struct: func(ctx *adlwo.StructContext, this *adlwo.AdlWoHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//Union: func(ctx *adlwo.UnionContext, this *adlwo.AdlWoHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//Type: func(ctx *adlwo.TypeContext, this *adlwo.AdlWoHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//Newtype: func(ctx *adlwo.NewtypeContext, this *adlwo.AdlWoHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//ModAnno: func(ctx *adlwo.ModAnnoContext, this *adlwo.AdlWoHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//DeclAnno: func(ctx *adlwo.DeclAnnoContext, this *adlwo.AdlWoHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//FieldAnno: func(ctx *adlwo.FieldAnnoContext, this *adlwo.AdlWoHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//Field: func(ctx *adlwo.FieldContext, this *adlwo.AdlWoHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//Annotation: func(ctx *adlwo.AnnotationContext, this *adlwo.AdlWoHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//TypeExpr_: func(ctx *adlwo.TypeExpr_Context, this *adlwo.AdlWoHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//TypeParams: func(ctx *adlwo.TypeParamsContext, this *adlwo.AdlWoHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//JsonStr: func(ctx *adlwo.JsonStrContext, this *adlwo.AdlWoHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//JsonBool: func(ctx *adlwo.JsonBoolContext, this *adlwo.AdlWoHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//JsonNull: func(ctx *adlwo.JsonNullContext, this *adlwo.AdlWoHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//JsonInt: func(ctx *adlwo.JsonIntContext, this *adlwo.AdlWoHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//JsonFloat: func(ctx *adlwo.JsonFloatContext, this *adlwo.AdlWoHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//JsonArray: func(ctx *adlwo.JsonArrayContext, this *adlwo.AdlWoHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//JsonObj: func(ctx *adlwo.JsonObjContext, this *adlwo.AdlWoHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},

//}

//func (v *AdlWoVisitor) VisitAdl(ctx adlwo.IAdlContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *AdlWoVisitor) VisitJson(ctx adlwo.IJsonContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *AdlWoVisitor) VisitModule(ctx adlwo.IModuleContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *AdlWoVisitor) VisitImportModule(ctx adlwo.IImportModuleContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *AdlWoVisitor) VisitImportScopedModule(ctx adlwo.IImportScopedModuleContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *AdlWoVisitor) VisitStruct(ctx adlwo.IStructContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *AdlWoVisitor) VisitUnion(ctx adlwo.IUnionContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *AdlWoVisitor) VisitType(ctx adlwo.ITypeContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *AdlWoVisitor) VisitNewtype(ctx adlwo.INewtypeContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *AdlWoVisitor) VisitModAnno(ctx adlwo.IModAnnoContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *AdlWoVisitor) VisitDeclAnno(ctx adlwo.IDeclAnnoContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *AdlWoVisitor) VisitFieldAnno(ctx adlwo.IFieldAnnoContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *AdlWoVisitor) VisitField(ctx adlwo.IFieldContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *AdlWoVisitor) VisitAnnotation(ctx adlwo.IAnnotationContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *AdlWoVisitor) VisitTypeExpr_(ctx adlwo.ITypeExpr_Context, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *AdlWoVisitor) VisitTypeParams(ctx adlwo.ITypeParamsContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *AdlWoVisitor) VisitJsonStr(ctx adlwo.IJsonStrContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *AdlWoVisitor) VisitJsonBool(ctx adlwo.IJsonBoolContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *AdlWoVisitor) VisitJsonNull(ctx adlwo.IJsonNullContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *AdlWoVisitor) VisitJsonInt(ctx adlwo.IJsonIntContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *AdlWoVisitor) VisitJsonFloat(ctx adlwo.IJsonFloatContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *AdlWoVisitor) VisitJsonArray(ctx adlwo.IJsonArrayContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *AdlWoVisitor) VisitJsonObj(ctx adlwo.IJsonObjContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}

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
