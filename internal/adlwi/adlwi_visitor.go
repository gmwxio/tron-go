// Generated from AdlWi.g4 by ANTLR 4.7.

package adlwi // AdlWi
import "github.com/wxio/goantlr"

// Struct of Handlers
type AdlWiHandlers struct {
	EnterEveryRule func(ctx antlr.RuleNode)
	ExitEveryRule  func(ctx antlr.RuleNode)

	Adl                func(ctx IAdlContext, this *AdlWiHandlers, args ...interface{}) (result interface{})
	Json               func(ctx IJsonContext, this *AdlWiHandlers, args ...interface{}) (result interface{})
	Module             func(ctx IModuleContext, this *AdlWiHandlers, args ...interface{}) (result interface{})
	ImportModule       func(ctx IImportModuleContext, this *AdlWiHandlers, args ...interface{}) (result interface{})
	ImportScopedModule func(ctx IImportScopedModuleContext, this *AdlWiHandlers, args ...interface{}) (result interface{})
	ImportError        func(ctx IImportErrorContext, this *AdlWiHandlers, args ...interface{}) (result interface{})
	Struct             func(ctx IStructContext, this *AdlWiHandlers, args ...interface{}) (result interface{})
	Union              func(ctx IUnionContext, this *AdlWiHandlers, args ...interface{}) (result interface{})
	Type               func(ctx ITypeContext, this *AdlWiHandlers, args ...interface{}) (result interface{})
	Newtype            func(ctx INewtypeContext, this *AdlWiHandlers, args ...interface{}) (result interface{})
	ModAnno            func(ctx IModAnnoContext, this *AdlWiHandlers, args ...interface{}) (result interface{})
	DeclAnno           func(ctx IDeclAnnoContext, this *AdlWiHandlers, args ...interface{}) (result interface{})
	FieldAnno          func(ctx IFieldAnnoContext, this *AdlWiHandlers, args ...interface{}) (result interface{})
	TypeParamError     func(ctx ITypeParamErrorContext, this *AdlWiHandlers, args ...interface{}) (result interface{})
	TLDError           func(ctx ITLDErrorContext, this *AdlWiHandlers, args ...interface{}) (result interface{})
	Field              func(ctx IFieldContext, this *AdlWiHandlers, args ...interface{}) (result interface{})
	Annotation         func(ctx IAnnotationContext, this *AdlWiHandlers, args ...interface{}) (result interface{})
	TypeExprSimple     func(ctx ITypeExprSimpleContext, this *AdlWiHandlers, args ...interface{}) (result interface{})
	TypeExprGeneric    func(ctx ITypeExprGenericContext, this *AdlWiHandlers, args ...interface{}) (result interface{})
	JsonStr            func(ctx IJsonStrContext, this *AdlWiHandlers, args ...interface{}) (result interface{})
	JsonBool           func(ctx IJsonBoolContext, this *AdlWiHandlers, args ...interface{}) (result interface{})
	JsonNull           func(ctx IJsonNullContext, this *AdlWiHandlers, args ...interface{}) (result interface{})
	JsonInt            func(ctx IJsonIntContext, this *AdlWiHandlers, args ...interface{}) (result interface{})
	JsonFloat          func(ctx IJsonFloatContext, this *AdlWiHandlers, args ...interface{}) (result interface{})
	JsonArray          func(ctx IJsonArrayContext, this *AdlWiHandlers, args ...interface{}) (result interface{})
	JsonObj            func(ctx IJsonObjContext, this *AdlWiHandlers, args ...interface{}) (result interface{})
	JsonError          func(ctx IJsonErrorContext, this *AdlWiHandlers, args ...interface{}) (result interface{})
}

// A complete Visitor for a parse tree produced by AdlWi.
type AdlWiVisitor interface {
	antlr.ParseTreeVisitor
	AdlContextVisitor
	JsonContextVisitor
	ModuleContextVisitor
	ImportModuleContextVisitor
	ImportScopedModuleContextVisitor
	ImportErrorContextVisitor
	StructContextVisitor
	UnionContextVisitor
	TypeContextVisitor
	NewtypeContextVisitor
	ModAnnoContextVisitor
	DeclAnnoContextVisitor
	FieldAnnoContextVisitor
	TypeParamErrorContextVisitor
	TLDErrorContextVisitor
	FieldContextVisitor
	AnnotationContextVisitor
	TypeExprSimpleContextVisitor
	TypeExprGenericContextVisitor
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
type JsonContextVisitor interface {
	VisitJson(ctx IJsonContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{})
}
type ModuleContextVisitor interface {
	VisitModule(ctx IModuleContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{})
}
type ImportModuleContextVisitor interface {
	VisitImportModule(ctx IImportModuleContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{})
}
type ImportScopedModuleContextVisitor interface {
	VisitImportScopedModule(ctx IImportScopedModuleContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{})
}
type ImportErrorContextVisitor interface {
	VisitImportError(ctx IImportErrorContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{})
}
type StructContextVisitor interface {
	VisitStruct(ctx IStructContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{})
}
type UnionContextVisitor interface {
	VisitUnion(ctx IUnionContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{})
}
type TypeContextVisitor interface {
	VisitType(ctx ITypeContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{})
}
type NewtypeContextVisitor interface {
	VisitNewtype(ctx INewtypeContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{})
}
type ModAnnoContextVisitor interface {
	VisitModAnno(ctx IModAnnoContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{})
}
type DeclAnnoContextVisitor interface {
	VisitDeclAnno(ctx IDeclAnnoContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{})
}
type FieldAnnoContextVisitor interface {
	VisitFieldAnno(ctx IFieldAnnoContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{})
}
type TypeParamErrorContextVisitor interface {
	VisitTypeParamError(ctx ITypeParamErrorContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{})
}
type TLDErrorContextVisitor interface {
	VisitTLDError(ctx ITLDErrorContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{})
}
type FieldContextVisitor interface {
	VisitField(ctx IFieldContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{})
}
type AnnotationContextVisitor interface {
	VisitAnnotation(ctx IAnnotationContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{})
}
type TypeExprSimpleContextVisitor interface {
	VisitTypeExprSimple(ctx ITypeExprSimpleContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{})
}
type TypeExprGenericContextVisitor interface {
	VisitTypeExprGeneric(ctx ITypeExprGenericContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{})
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
