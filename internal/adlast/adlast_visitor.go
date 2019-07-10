// Generated from AdlAst.g4 by ANTLR 4.7.

package adlast // AdlAst
import "github.com/wxio/goantlr"

// Struct of Handlers
type AdlAstHandlers struct {
	EnterEveryRule func(ctx antlr.RuleNode)
	ExitEveryRule  func(ctx antlr.RuleNode)

	Adl        func(ctx IAdlContext, this *AdlAstHandlers, args ...interface{}) (result interface{})
	Module     func(ctx IModuleContext, this *AdlAstHandlers, args ...interface{}) (result interface{})
	Import_    func(ctx IImport_Context, this *AdlAstHandlers, args ...interface{}) (result interface{})
	Decl       func(ctx IDeclContext, this *AdlAstHandlers, args ...interface{}) (result interface{})
	ScopedName func(ctx IScopedNameContext, this *AdlAstHandlers, args ...interface{}) (result interface{})
	DeclType   func(ctx IDeclTypeContext, this *AdlAstHandlers, args ...interface{}) (result interface{})
	Struct_    func(ctx IStruct_Context, this *AdlAstHandlers, args ...interface{}) (result interface{})
	Union_     func(ctx IUnion_Context, this *AdlAstHandlers, args ...interface{}) (result interface{})
	TypeDef    func(ctx ITypeDefContext, this *AdlAstHandlers, args ...interface{}) (result interface{})
	NewType    func(ctx INewTypeContext, this *AdlAstHandlers, args ...interface{}) (result interface{})
	Field      func(ctx IFieldContext, this *AdlAstHandlers, args ...interface{}) (result interface{})
	TypeExpr   func(ctx ITypeExprContext, this *AdlAstHandlers, args ...interface{}) (result interface{})
	TypeRef    func(ctx ITypeRefContext, this *AdlAstHandlers, args ...interface{}) (result interface{})
	JsonStr    func(ctx IJsonStrContext, this *AdlAstHandlers, args ...interface{}) (result interface{})
	JsonBool   func(ctx IJsonBoolContext, this *AdlAstHandlers, args ...interface{}) (result interface{})
	JsonNull   func(ctx IJsonNullContext, this *AdlAstHandlers, args ...interface{}) (result interface{})
	JsonInt    func(ctx IJsonIntContext, this *AdlAstHandlers, args ...interface{}) (result interface{})
	JsonFloat  func(ctx IJsonFloatContext, this *AdlAstHandlers, args ...interface{}) (result interface{})
	JsonArray  func(ctx IJsonArrayContext, this *AdlAstHandlers, args ...interface{}) (result interface{})
	JsonObj    func(ctx IJsonObjContext, this *AdlAstHandlers, args ...interface{}) (result interface{})
	JsonError  func(ctx IJsonErrorContext, this *AdlAstHandlers, args ...interface{}) (result interface{})
}

// A complete Visitor for a parse tree produced by AdlAst.
type AdlAstVisitor interface {
	antlr.ParseTreeVisitor
	AdlContextVisitor
	ModuleContextVisitor
	Import_ContextVisitor
	DeclContextVisitor
	ScopedNameContextVisitor
	DeclTypeContextVisitor
	Struct_ContextVisitor
	Union_ContextVisitor
	TypeDefContextVisitor
	NewTypeContextVisitor
	FieldContextVisitor
	TypeExprContextVisitor
	TypeRefContextVisitor
	JsonStrContextVisitor
	JsonBoolContextVisitor
	JsonNullContextVisitor
	JsonIntContextVisitor
	JsonFloatContextVisitor
	JsonArrayContextVisitor
	JsonObjContextVisitor
	JsonErrorContextVisitor
}

type AdlContextVisitor interface {
	VisitAdl(ctx IAdlContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{})
}
type ModuleContextVisitor interface {
	VisitModule(ctx IModuleContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{})
}
type Import_ContextVisitor interface {
	VisitImport_(ctx IImport_Context, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{})
}
type DeclContextVisitor interface {
	VisitDecl(ctx IDeclContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{})
}
type ScopedNameContextVisitor interface {
	VisitScopedName(ctx IScopedNameContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{})
}
type DeclTypeContextVisitor interface {
	VisitDeclType(ctx IDeclTypeContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{})
}
type Struct_ContextVisitor interface {
	VisitStruct_(ctx IStruct_Context, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{})
}
type Union_ContextVisitor interface {
	VisitUnion_(ctx IUnion_Context, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{})
}
type TypeDefContextVisitor interface {
	VisitTypeDef(ctx ITypeDefContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{})
}
type NewTypeContextVisitor interface {
	VisitNewType(ctx INewTypeContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{})
}
type FieldContextVisitor interface {
	VisitField(ctx IFieldContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{})
}
type TypeExprContextVisitor interface {
	VisitTypeExpr(ctx ITypeExprContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{})
}
type TypeRefContextVisitor interface {
	VisitTypeRef(ctx ITypeRefContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{})
}
type JsonStrContextVisitor interface {
	VisitJsonStr(ctx IJsonStrContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{})
}
type JsonBoolContextVisitor interface {
	VisitJsonBool(ctx IJsonBoolContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{})
}
type JsonNullContextVisitor interface {
	VisitJsonNull(ctx IJsonNullContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{})
}
type JsonIntContextVisitor interface {
	VisitJsonInt(ctx IJsonIntContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{})
}
type JsonFloatContextVisitor interface {
	VisitJsonFloat(ctx IJsonFloatContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{})
}
type JsonArrayContextVisitor interface {
	VisitJsonArray(ctx IJsonArrayContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{})
}
type JsonObjContextVisitor interface {
	VisitJsonObj(ctx IJsonObjContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{})
}
type JsonErrorContextVisitor interface {
	VisitJsonError(ctx IJsonErrorContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{})
}
