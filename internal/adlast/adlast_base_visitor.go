// Generated from AdlAst.g4 by ANTLR 4.7.

package adlast // AdlAst
// See base listener file for example implementations

//import "github.com/wxio/goantlr"

// import "generate package"

//type AdlAstVisitor struct {
//    *antlr.BaseParseTreeVisitor
//    indent string
//}

//var _ antlr.ParseTreeVisitor = &AdlAstVisitor{}
//var _ antlr.AggregateResultVisitor = &AdlAstVisitor{}
//var _ antlr.VisitNextCheck = &AdlAstVisitor{}
//var _ antlr.VisitRestCheck = &AdlAstVisitor{}
//var _ antlr.EnterEveryRuleVisitor = &AdlAstVisitor{}
//var _ antlr.ExitEveryRuleVisitor = &AdlAstVisitor{}

//var _ adlast.AdlContextVisitor = &AdlAstVisitor{}
//var _ adlast.ModuleContextVisitor = &AdlAstVisitor{}
//var _ adlast.Import_ContextVisitor = &AdlAstVisitor{}
//var _ adlast.DeclContextVisitor = &AdlAstVisitor{}
//var _ adlast.ScopedNameContextVisitor = &AdlAstVisitor{}
//var _ adlast.DeclTypeContextVisitor = &AdlAstVisitor{}
//var _ adlast.Struct_ContextVisitor = &AdlAstVisitor{}
//var _ adlast.Union_ContextVisitor = &AdlAstVisitor{}
//var _ adlast.TypeDefContextVisitor = &AdlAstVisitor{}
//var _ adlast.NewTypeContextVisitor = &AdlAstVisitor{}
//var _ adlast.FieldContextVisitor = &AdlAstVisitor{}
//var _ adlast.TypeExprContextVisitor = &AdlAstVisitor{}
//var _ adlast.TypeRefContextVisitor = &AdlAstVisitor{}
//var _ adlast.JsonStrContextVisitor = &AdlAstVisitor{}
//var _ adlast.JsonBoolContextVisitor = &AdlAstVisitor{}
//var _ adlast.JsonNullContextVisitor = &AdlAstVisitor{}
//var _ adlast.JsonIntContextVisitor = &AdlAstVisitor{}
//var _ adlast.JsonFloatContextVisitor = &AdlAstVisitor{}
//var _ adlast.JsonArrayContextVisitor = &AdlAstVisitor{}
//var _ adlast.JsonObjContextVisitor = &AdlAstVisitor{}
//var _ adlast.JsonErrorContextVisitor = &AdlAstVisitor{}

//func (v *AdlAstVisitor) VisitNext(node antlr.Tree, current interface{}) bool {
//    return true
//}
//func (v *AdlAstVisitor) VisitRest(node antlr.RuleNode, current interface{}) bool {
//    return true
//}
//func (v *AdlAstVisitor) AggregateResult(aggregate, nextResult interface{}) (result interface{}) {
//    return nextResult
//}
//func (v *AdlAstVisitor) VisitTerminal(node antlr.TerminalNode) {
//}
//func (v *AdlAstVisitor) VisitErrorNode(node antlr.ErrorNode) {
//}
//func (v *AdlAstVisitor) EnterEveryRule(ctx antlr.RuleNode) {
//    v.indent += " "
//    fmt.Printf("%s %T\n", v.indent, ctx)
//}
//func (v *AdlAstVisitor) ExitEveryRule(ctx antlr.RuleNode) {
//    v.indent = v.indent[:len(v.indent)-1]
//}

//indent := ""
//parsetree.VisitFunc(&adlast.AdlAstHandlers{
//EnterEveryRule: func(ctx antlr.RuleNode){indent += "\t"},
//ExitEveryRule:  func(ctx antlr.RuleNode){indent = indent[1:]},
//Adl: func(ctx *adlast.AdlContext, this *adlast.AdlAstHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//Module: func(ctx *adlast.ModuleContext, this *adlast.AdlAstHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//Import_: func(ctx *adlast.Import_Context, this *adlast.AdlAstHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//Decl: func(ctx *adlast.DeclContext, this *adlast.AdlAstHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//ScopedName: func(ctx *adlast.ScopedNameContext, this *adlast.AdlAstHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//DeclType: func(ctx *adlast.DeclTypeContext, this *adlast.AdlAstHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//Struct_: func(ctx *adlast.Struct_Context, this *adlast.AdlAstHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//Union_: func(ctx *adlast.Union_Context, this *adlast.AdlAstHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//TypeDef: func(ctx *adlast.TypeDefContext, this *adlast.AdlAstHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//NewType: func(ctx *adlast.NewTypeContext, this *adlast.AdlAstHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//Field: func(ctx *adlast.FieldContext, this *adlast.AdlAstHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//TypeExpr: func(ctx *adlast.TypeExprContext, this *adlast.AdlAstHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//TypeRef: func(ctx *adlast.TypeRefContext, this *adlast.AdlAstHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//JsonStr: func(ctx *adlast.JsonStrContext, this *adlast.AdlAstHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//JsonBool: func(ctx *adlast.JsonBoolContext, this *adlast.AdlAstHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//JsonNull: func(ctx *adlast.JsonNullContext, this *adlast.AdlAstHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//JsonInt: func(ctx *adlast.JsonIntContext, this *adlast.AdlAstHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//JsonFloat: func(ctx *adlast.JsonFloatContext, this *adlast.AdlAstHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//JsonArray: func(ctx *adlast.JsonArrayContext, this *adlast.AdlAstHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//JsonObj: func(ctx *adlast.JsonObjContext, this *adlast.AdlAstHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},
//JsonError: func(ctx *adlast.JsonErrorContext, this *adlast.AdlAstHandlers, args ...interface{}) (result interface{}) {fmt.Printf("%s%T\n", indent, ctx);ctx.VisitChildrenFunc(this, args...);return},

//}

//func (v *AdlAstVisitor) VisitAdl(ctx adlast.IAdlContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *AdlAstVisitor) VisitModule(ctx adlast.IModuleContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *AdlAstVisitor) VisitImport_(ctx adlast.IImport_Context, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *AdlAstVisitor) VisitDecl(ctx adlast.IDeclContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *AdlAstVisitor) VisitScopedName(ctx adlast.IScopedNameContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *AdlAstVisitor) VisitDeclType(ctx adlast.IDeclTypeContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *AdlAstVisitor) VisitStruct_(ctx adlast.IStruct_Context, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *AdlAstVisitor) VisitUnion_(ctx adlast.IUnion_Context, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *AdlAstVisitor) VisitTypeDef(ctx adlast.ITypeDefContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *AdlAstVisitor) VisitNewType(ctx adlast.INewTypeContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *AdlAstVisitor) VisitField(ctx adlast.IFieldContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *AdlAstVisitor) VisitTypeExpr(ctx adlast.ITypeExprContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *AdlAstVisitor) VisitTypeRef(ctx adlast.ITypeRefContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *AdlAstVisitor) VisitJsonStr(ctx adlast.IJsonStrContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *AdlAstVisitor) VisitJsonBool(ctx adlast.IJsonBoolContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *AdlAstVisitor) VisitJsonNull(ctx adlast.IJsonNullContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *AdlAstVisitor) VisitJsonInt(ctx adlast.IJsonIntContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *AdlAstVisitor) VisitJsonFloat(ctx adlast.IJsonFloatContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *AdlAstVisitor) VisitJsonArray(ctx adlast.IJsonArrayContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *AdlAstVisitor) VisitJsonObj(ctx adlast.IJsonObjContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}
//func (v *AdlAstVisitor) VisitJsonError(ctx adlast.IJsonErrorContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    result = v.VisitChildren(ctx, delegate, args...);return}

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
