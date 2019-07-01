// Generated from AdlWo.g4 by ANTLR 4.7.

package adlwo // AdlWo
import "github.com/wxio/goantlr"

// AdlWoListener is a complete listener for a parse tree produced by AdlWo.
type AdlWoListener interface {
	antlr.ParseTreeListener

	AdlEntryListener
	AdlExitListener

	JsonEntryListener
	JsonExitListener

	ModuleEntryListener
	ModuleExitListener

	AnnotationEntryListener
	AnnotationExitListener

	TypeExpr_EntryListener
	TypeExpr_ExitListener

	ImportModuleEntryListener
	ImportModuleExitListener

	ImportScopedModuleEntryListener
	ImportScopedModuleExitListener

	StructEntryListener
	StructExitListener

	UnionEntryListener
	UnionExitListener

	TypeEntryListener
	TypeExitListener

	NewtypeEntryListener
	NewtypeExitListener

	ModAnnoEntryListener
	ModAnnoExitListener

	DeclAnnoEntryListener
	DeclAnnoExitListener

	FieldAnnoEntryListener
	FieldAnnoExitListener

	FieldEntryListener
	FieldExitListener

	TypeParamsEntryListener
	TypeParamsExitListener

	JsonStrEntryListener
	JsonStrExitListener

	JsonBoolEntryListener
	JsonBoolExitListener

	JsonNullEntryListener
	JsonNullExitListener

	JsonIntEntryListener
	JsonIntExitListener

	JsonFloatEntryListener
	JsonFloatExitListener

	JsonArrayEntryListener
	JsonArrayExitListener

	JsonObjEntryListener
	JsonObjExitListener
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

type JsonEntryListener interface {
	EnterJson(c *JsonContext)
}
type JsonExitListener interface {
	ExitJson(c *JsonContext)
}

type ModuleEntryListener interface {
	EnterModule(c *ModuleContext)
}
type ModuleExitListener interface {
	ExitModule(c *ModuleContext)
}

type AnnotationEntryListener interface {
	EnterAnnotation(c *AnnotationContext)
}
type AnnotationExitListener interface {
	ExitAnnotation(c *AnnotationContext)
}

type TypeExpr_EntryListener interface {
	EnterTypeExpr_(c *TypeExpr_Context)
}
type TypeExpr_ExitListener interface {
	ExitTypeExpr_(c *TypeExpr_Context)
}

//
// Named alternatives
//
//
// From Rule 'import_'
type ImportModuleEntryListener interface {
	EnterImportModule(c *ImportModuleContext)
}
type ImportModuleExitListener interface {
	ExitImportModule(c *ImportModuleContext)
}

// From Rule 'import_'
type ImportScopedModuleEntryListener interface {
	EnterImportScopedModule(c *ImportScopedModuleContext)
}
type ImportScopedModuleExitListener interface {
	ExitImportScopedModule(c *ImportScopedModuleContext)
}

// From Rule 'tld'
type StructEntryListener interface {
	EnterStruct(c *StructContext)
}
type StructExitListener interface {
	ExitStruct(c *StructContext)
}

// From Rule 'tld'
type UnionEntryListener interface {
	EnterUnion(c *UnionContext)
}
type UnionExitListener interface {
	ExitUnion(c *UnionContext)
}

// From Rule 'tld'
type TypeEntryListener interface {
	EnterType(c *TypeContext)
}
type TypeExitListener interface {
	ExitType(c *TypeContext)
}

// From Rule 'tld'
type NewtypeEntryListener interface {
	EnterNewtype(c *NewtypeContext)
}
type NewtypeExitListener interface {
	ExitNewtype(c *NewtypeContext)
}

// From Rule 'tld'
type ModAnnoEntryListener interface {
	EnterModAnno(c *ModAnnoContext)
}
type ModAnnoExitListener interface {
	ExitModAnno(c *ModAnnoContext)
}

// From Rule 'tld'
type DeclAnnoEntryListener interface {
	EnterDeclAnno(c *DeclAnnoContext)
}
type DeclAnnoExitListener interface {
	ExitDeclAnno(c *DeclAnnoContext)
}

// From Rule 'tld'
type FieldAnnoEntryListener interface {
	EnterFieldAnno(c *FieldAnnoContext)
}
type FieldAnnoExitListener interface {
	ExitFieldAnno(c *FieldAnnoContext)
}

// From Rule 'nameBody'
type FieldEntryListener interface {
	EnterField(c *FieldContext)
}
type FieldExitListener interface {
	ExitField(c *FieldContext)
}

// From Rule 'typeExprElem_'
type TypeParamsEntryListener interface {
	EnterTypeParams(c *TypeParamsContext)
}
type TypeParamsExitListener interface {
	ExitTypeParams(c *TypeParamsContext)
}

// From Rule 'jsonVal'
type JsonStrEntryListener interface {
	EnterJsonStr(c *JsonStrContext)
}
type JsonStrExitListener interface {
	ExitJsonStr(c *JsonStrContext)
}

// From Rule 'jsonVal'
type JsonBoolEntryListener interface {
	EnterJsonBool(c *JsonBoolContext)
}
type JsonBoolExitListener interface {
	ExitJsonBool(c *JsonBoolContext)
}

// From Rule 'jsonVal'
type JsonNullEntryListener interface {
	EnterJsonNull(c *JsonNullContext)
}
type JsonNullExitListener interface {
	ExitJsonNull(c *JsonNullContext)
}

// From Rule 'jsonVal'
type JsonIntEntryListener interface {
	EnterJsonInt(c *JsonIntContext)
}
type JsonIntExitListener interface {
	ExitJsonInt(c *JsonIntContext)
}

// From Rule 'jsonVal'
type JsonFloatEntryListener interface {
	EnterJsonFloat(c *JsonFloatContext)
}
type JsonFloatExitListener interface {
	ExitJsonFloat(c *JsonFloatContext)
}

// From Rule 'jsonVal'
type JsonArrayEntryListener interface {
	EnterJsonArray(c *JsonArrayContext)
}
type JsonArrayExitListener interface {
	ExitJsonArray(c *JsonArrayContext)
}

// From Rule 'jsonVal'
type JsonObjEntryListener interface {
	EnterJsonObj(c *JsonObjContext)
}
type JsonObjExitListener interface {
	ExitJsonObj(c *JsonObjContext)
}
