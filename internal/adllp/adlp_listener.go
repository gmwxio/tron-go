// Generated from AdlP.g4 by ANTLR 4.7.

package adllp // AdlP
import "github.com/wxio/goantlr"

// AdlPListener is a complete listener for a parse tree produced by AdlP.
type AdlPListener interface {
	antlr.ParseTreeListener

	AdlEntryListener
	AdlExitListener

	ModuleStatementEntryListener
	ModuleStatementExitListener

	ImportScopedNameEntryListener
	ImportScopedNameExitListener

	ImportModuleNameEntryListener
	ImportModuleNameExitListener

	LocalAnnoEntryListener
	LocalAnnoExitListener

	DocAnnoEntryListener
	DocAnnoExitListener

	StructOrUnionEntryListener
	StructOrUnionExitListener

	TypeOrNewtypeEntryListener
	TypeOrNewtypeExitListener

	ModuleAnnotationEntryListener
	ModuleAnnotationExitListener

	DeclAnnotationEntryListener
	DeclAnnotationExitListener

	FieldAnnotationEntryListener
	FieldAnnotationExitListener

	TypeParameterEntryListener
	TypeParameterExitListener

	TypeExprPrimOrParamEntryListener
	TypeExprPrimOrParamExitListener

	TypeExprTypeExprEntryListener
	TypeExprTypeExprExitListener

	TypeExpressionElemEntryListener
	TypeExpressionElemExitListener

	FieldStatementEntryListener
	FieldStatementExitListener

	StringStatementEntryListener
	StringStatementExitListener

	TrueFalseNullEntryListener
	TrueFalseNullExitListener

	NumberStatementEntryListener
	NumberStatementExitListener

	FloatStatementEntryListener
	FloatStatementExitListener

	ArrayStatementEntryListener
	ArrayStatementExitListener

	ObjStatementEntryListener
	ObjStatementExitListener
}

//
// Rules with unnamed alternatives
//
type AdlEntryListener interface {
	EnterAdl(c *AdlContext)
}
type AdlExitListener interface {
	ExitAdl(c *AdlContext)
}

//
// Named alternatives
//
//
// From Rule 'module'
type ModuleStatementEntryListener interface {
	EnterModuleStatement(c *ModuleStatementContext)
}
type ModuleStatementExitListener interface {
	ExitModuleStatement(c *ModuleStatementContext)
}

// From Rule 'imports'
type ImportScopedNameEntryListener interface {
	EnterImportScopedName(c *ImportScopedNameContext)
}
type ImportScopedNameExitListener interface {
	ExitImportScopedName(c *ImportScopedNameContext)
}

// From Rule 'imports'
type ImportModuleNameEntryListener interface {
	EnterImportModuleName(c *ImportModuleNameContext)
}
type ImportModuleNameExitListener interface {
	ExitImportModuleName(c *ImportModuleNameContext)
}

// From Rule 'annon'
type LocalAnnoEntryListener interface {
	EnterLocalAnno(c *LocalAnnoContext)
}
type LocalAnnoExitListener interface {
	ExitLocalAnno(c *LocalAnnoContext)
}

// From Rule 'annon'
type DocAnnoEntryListener interface {
	EnterDocAnno(c *DocAnnoContext)
}
type DocAnnoExitListener interface {
	ExitDocAnno(c *DocAnnoContext)
}

// From Rule 'top_level_statement'
type StructOrUnionEntryListener interface {
	EnterStructOrUnion(c *StructOrUnionContext)
}
type StructOrUnionExitListener interface {
	ExitStructOrUnion(c *StructOrUnionContext)
}

// From Rule 'top_level_statement'
type TypeOrNewtypeEntryListener interface {
	EnterTypeOrNewtype(c *TypeOrNewtypeContext)
}
type TypeOrNewtypeExitListener interface {
	ExitTypeOrNewtype(c *TypeOrNewtypeContext)
}

// From Rule 'top_level_statement'
type ModuleAnnotationEntryListener interface {
	EnterModuleAnnotation(c *ModuleAnnotationContext)
}
type ModuleAnnotationExitListener interface {
	ExitModuleAnnotation(c *ModuleAnnotationContext)
}

// From Rule 'top_level_statement'
type DeclAnnotationEntryListener interface {
	EnterDeclAnnotation(c *DeclAnnotationContext)
}
type DeclAnnotationExitListener interface {
	ExitDeclAnnotation(c *DeclAnnotationContext)
}

// From Rule 'top_level_statement'
type FieldAnnotationEntryListener interface {
	EnterFieldAnnotation(c *FieldAnnotationContext)
}
type FieldAnnotationExitListener interface {
	ExitFieldAnnotation(c *FieldAnnotationContext)
}

// From Rule 'typeParam'
type TypeParameterEntryListener interface {
	EnterTypeParameter(c *TypeParameterContext)
}
type TypeParameterExitListener interface {
	ExitTypeParameter(c *TypeParameterContext)
}

// From Rule 'typeExpr'
type TypeExprPrimOrParamEntryListener interface {
	EnterTypeExprPrimOrParam(c *TypeExprPrimOrParamContext)
}
type TypeExprPrimOrParamExitListener interface {
	ExitTypeExprPrimOrParam(c *TypeExprPrimOrParamContext)
}

// From Rule 'typeExpr'
type TypeExprTypeExprEntryListener interface {
	EnterTypeExprTypeExpr(c *TypeExprTypeExprContext)
}
type TypeExprTypeExprExitListener interface {
	ExitTypeExprTypeExpr(c *TypeExprTypeExprContext)
}

// From Rule 'typeExprElem'
type TypeExpressionElemEntryListener interface {
	EnterTypeExpressionElem(c *TypeExpressionElemContext)
}
type TypeExpressionElemExitListener interface {
	ExitTypeExpressionElem(c *TypeExpressionElemContext)
}

// From Rule 'soruBody'
type FieldStatementEntryListener interface {
	EnterFieldStatement(c *FieldStatementContext)
}
type FieldStatementExitListener interface {
	ExitFieldStatement(c *FieldStatementContext)
}

// From Rule 'jsonValue'
type StringStatementEntryListener interface {
	EnterStringStatement(c *StringStatementContext)
}
type StringStatementExitListener interface {
	ExitStringStatement(c *StringStatementContext)
}

// From Rule 'jsonValue'
type TrueFalseNullEntryListener interface {
	EnterTrueFalseNull(c *TrueFalseNullContext)
}
type TrueFalseNullExitListener interface {
	ExitTrueFalseNull(c *TrueFalseNullContext)
}

// From Rule 'jsonValue'
type NumberStatementEntryListener interface {
	EnterNumberStatement(c *NumberStatementContext)
}
type NumberStatementExitListener interface {
	ExitNumberStatement(c *NumberStatementContext)
}

// From Rule 'jsonValue'
type FloatStatementEntryListener interface {
	EnterFloatStatement(c *FloatStatementContext)
}
type FloatStatementExitListener interface {
	ExitFloatStatement(c *FloatStatementContext)
}

// From Rule 'jsonValue'
type ArrayStatementEntryListener interface {
	EnterArrayStatement(c *ArrayStatementContext)
}
type ArrayStatementExitListener interface {
	ExitArrayStatement(c *ArrayStatementContext)
}

// From Rule 'jsonValue'
type ObjStatementEntryListener interface {
	EnterObjStatement(c *ObjStatementContext)
}
type ObjStatementExitListener interface {
	ExitObjStatement(c *ObjStatementContext)
}
