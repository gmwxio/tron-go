// Generated from AdlAst.g4 by ANTLR 4.7.

package adlast // AdlAst
import "github.com/wxio/goantlr"

// AdlAstListener is a complete listener for a parse tree produced by AdlAst.
type AdlAstListener interface {
	antlr.ParseTreeListener

	AdlEntryListener
	AdlExitListener

	ModuleEntryListener
	ModuleExitListener

	Import_EntryListener
	Import_ExitListener

	DeclEntryListener
	DeclExitListener

	ScopedNameEntryListener
	ScopedNameExitListener

	DeclTypeEntryListener
	DeclTypeExitListener

	Struct_EntryListener
	Struct_ExitListener

	Union_EntryListener
	Union_ExitListener

	TypeDefEntryListener
	TypeDefExitListener

	NewTypeEntryListener
	NewTypeExitListener

	FieldEntryListener
	FieldExitListener

	TypeExprEntryListener
	TypeExprExitListener

	TypeRefEntryListener
	TypeRefExitListener

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

	JsonErrorEntryListener
	JsonErrorExitListener
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

type ModuleEntryListener interface {
	EnterModule(c *ModuleContext)
}
type ModuleExitListener interface {
	ExitModule(c *ModuleContext)
}

type Import_EntryListener interface {
	EnterImport_(c *Import_Context)
}
type Import_ExitListener interface {
	ExitImport_(c *Import_Context)
}

type DeclEntryListener interface {
	EnterDecl(c *DeclContext)
}
type DeclExitListener interface {
	ExitDecl(c *DeclContext)
}

type ScopedNameEntryListener interface {
	EnterScopedName(c *ScopedNameContext)
}
type ScopedNameExitListener interface {
	ExitScopedName(c *ScopedNameContext)
}

type DeclTypeEntryListener interface {
	EnterDeclType(c *DeclTypeContext)
}
type DeclTypeExitListener interface {
	ExitDeclType(c *DeclTypeContext)
}

type Struct_EntryListener interface {
	EnterStruct_(c *Struct_Context)
}
type Struct_ExitListener interface {
	ExitStruct_(c *Struct_Context)
}

type Union_EntryListener interface {
	EnterUnion_(c *Union_Context)
}
type Union_ExitListener interface {
	ExitUnion_(c *Union_Context)
}

type TypeDefEntryListener interface {
	EnterTypeDef(c *TypeDefContext)
}
type TypeDefExitListener interface {
	ExitTypeDef(c *TypeDefContext)
}

type NewTypeEntryListener interface {
	EnterNewType(c *NewTypeContext)
}
type NewTypeExitListener interface {
	ExitNewType(c *NewTypeContext)
}

type FieldEntryListener interface {
	EnterField(c *FieldContext)
}
type FieldExitListener interface {
	ExitField(c *FieldContext)
}

type TypeExprEntryListener interface {
	EnterTypeExpr(c *TypeExprContext)
}
type TypeExprExitListener interface {
	ExitTypeExpr(c *TypeExprContext)
}

type TypeRefEntryListener interface {
	EnterTypeRef(c *TypeRefContext)
}
type TypeRefExitListener interface {
	ExitTypeRef(c *TypeRefContext)
}

//
// Named alternatives
//
//
// From Rule 'json'
type JsonStrEntryListener interface {
	EnterJsonStr(c *JsonStrContext)
}
type JsonStrExitListener interface {
	ExitJsonStr(c *JsonStrContext)
}

// From Rule 'json'
type JsonBoolEntryListener interface {
	EnterJsonBool(c *JsonBoolContext)
}
type JsonBoolExitListener interface {
	ExitJsonBool(c *JsonBoolContext)
}

// From Rule 'json'
type JsonNullEntryListener interface {
	EnterJsonNull(c *JsonNullContext)
}
type JsonNullExitListener interface {
	ExitJsonNull(c *JsonNullContext)
}

// From Rule 'json'
type JsonIntEntryListener interface {
	EnterJsonInt(c *JsonIntContext)
}
type JsonIntExitListener interface {
	ExitJsonInt(c *JsonIntContext)
}

// From Rule 'json'
type JsonFloatEntryListener interface {
	EnterJsonFloat(c *JsonFloatContext)
}
type JsonFloatExitListener interface {
	ExitJsonFloat(c *JsonFloatContext)
}

// From Rule 'json'
type JsonArrayEntryListener interface {
	EnterJsonArray(c *JsonArrayContext)
}
type JsonArrayExitListener interface {
	ExitJsonArray(c *JsonArrayContext)
}

// From Rule 'json'
type JsonObjEntryListener interface {
	EnterJsonObj(c *JsonObjContext)
}
type JsonObjExitListener interface {
	ExitJsonObj(c *JsonObjContext)
}

// From Rule 'json'
type JsonErrorEntryListener interface {
	EnterJsonError(c *JsonErrorContext)
}
type JsonErrorExitListener interface {
	ExitJsonError(c *JsonErrorContext)
}
