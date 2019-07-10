// Generated from AdlAst.g4 by ANTLR 4.7.

package adlast // AdlAst
import (
	"fmt"
	"reflect"
	"strconv"
)
import "github.com/wxio/goantlr"

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 64, 226,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3,
	3, 3, 7, 3, 38, 10, 3, 12, 3, 14, 3, 41, 11, 3, 3, 3, 3, 3, 7, 3, 45, 10,
	3, 12, 3, 14, 3, 48, 11, 3, 3, 3, 3, 3, 3, 3, 7, 3, 53, 10, 3, 12, 3, 14,
	3, 56, 11, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 64, 10, 4, 3, 4,
	3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 5, 5, 72, 10, 5, 3, 5, 3, 5, 3, 5, 3, 5,
	7, 5, 78, 10, 5, 12, 5, 14, 5, 81, 11, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6,
	3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 5, 7, 97, 10, 7,
	3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 7, 8, 104, 10, 8, 12, 8, 14, 8, 107, 11,
	8, 3, 8, 7, 8, 110, 10, 8, 12, 8, 14, 8, 113, 11, 8, 3, 8, 5, 8, 116, 10,
	8, 3, 9, 3, 9, 3, 9, 7, 9, 121, 10, 9, 12, 9, 14, 9, 124, 11, 9, 3, 9,
	7, 9, 127, 10, 9, 12, 9, 14, 9, 130, 11, 9, 3, 9, 5, 9, 133, 10, 9, 3,
	10, 3, 10, 3, 10, 7, 10, 138, 10, 10, 12, 10, 14, 10, 141, 11, 10, 3, 10,
	3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 7, 11, 149, 10, 11, 12, 11, 14, 11,
	152, 11, 11, 3, 11, 3, 11, 5, 11, 156, 10, 11, 3, 11, 3, 11, 3, 12, 3,
	12, 3, 12, 3, 12, 3, 12, 3, 12, 5, 12, 166, 10, 12, 3, 12, 3, 12, 3, 12,
	7, 12, 171, 10, 12, 12, 12, 14, 12, 174, 11, 12, 3, 12, 3, 12, 3, 13, 3,
	13, 3, 13, 3, 13, 7, 13, 182, 10, 13, 12, 13, 14, 13, 185, 11, 13, 3, 13,
	3, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 5, 14, 194, 10, 14, 3, 15, 3,
	15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 6, 15, 204, 10, 15, 13, 15,
	14, 15, 205, 3, 15, 3, 15, 5, 15, 210, 10, 15, 3, 15, 3, 15, 3, 15, 6,
	15, 215, 10, 15, 13, 15, 14, 15, 216, 3, 15, 3, 15, 5, 15, 221, 10, 15,
	3, 15, 5, 15, 224, 10, 15, 3, 15, 2, 2, 16, 2, 4, 6, 8, 10, 12, 14, 16,
	18, 20, 22, 24, 26, 28, 2, 2, 2, 245, 2, 30, 3, 2, 2, 2, 4, 33, 3, 2, 2,
	2, 6, 59, 3, 2, 2, 2, 8, 67, 3, 2, 2, 2, 10, 84, 3, 2, 2, 2, 12, 90, 3,
	2, 2, 2, 14, 100, 3, 2, 2, 2, 16, 117, 3, 2, 2, 2, 18, 134, 3, 2, 2, 2,
	20, 145, 3, 2, 2, 2, 22, 159, 3, 2, 2, 2, 24, 177, 3, 2, 2, 2, 26, 193,
	3, 2, 2, 2, 28, 223, 3, 2, 2, 2, 30, 31, 5, 4, 3, 2, 31, 32, 7, 2, 2, 3,
	32, 3, 3, 2, 2, 2, 33, 34, 7, 31, 2, 2, 34, 35, 7, 26, 2, 2, 35, 39, 7,
	20, 2, 2, 36, 38, 5, 6, 4, 2, 37, 36, 3, 2, 2, 2, 38, 41, 3, 2, 2, 2, 39,
	37, 3, 2, 2, 2, 39, 40, 3, 2, 2, 2, 40, 46, 3, 2, 2, 2, 41, 39, 3, 2, 2,
	2, 42, 43, 7, 20, 2, 2, 43, 45, 5, 8, 5, 2, 44, 42, 3, 2, 2, 2, 45, 48,
	3, 2, 2, 2, 46, 44, 3, 2, 2, 2, 46, 47, 3, 2, 2, 2, 47, 54, 3, 2, 2, 2,
	48, 46, 3, 2, 2, 2, 49, 50, 5, 10, 6, 2, 50, 51, 5, 28, 15, 2, 51, 53,
	3, 2, 2, 2, 52, 49, 3, 2, 2, 2, 53, 56, 3, 2, 2, 2, 54, 52, 3, 2, 2, 2,
	54, 55, 3, 2, 2, 2, 55, 57, 3, 2, 2, 2, 56, 54, 3, 2, 2, 2, 57, 58, 7,
	27, 2, 2, 58, 5, 3, 2, 2, 2, 59, 60, 7, 32, 2, 2, 60, 63, 7, 26, 2, 2,
	61, 64, 7, 20, 2, 2, 62, 64, 5, 10, 6, 2, 63, 61, 3, 2, 2, 2, 63, 62, 3,
	2, 2, 2, 64, 65, 3, 2, 2, 2, 65, 66, 7, 27, 2, 2, 66, 7, 3, 2, 2, 2, 67,
	68, 7, 35, 2, 2, 68, 69, 7, 26, 2, 2, 69, 71, 7, 20, 2, 2, 70, 72, 7, 21,
	2, 2, 71, 70, 3, 2, 2, 2, 71, 72, 3, 2, 2, 2, 72, 73, 3, 2, 2, 2, 73, 79,
	5, 12, 7, 2, 74, 75, 5, 10, 6, 2, 75, 76, 5, 28, 15, 2, 76, 78, 3, 2, 2,
	2, 77, 74, 3, 2, 2, 2, 78, 81, 3, 2, 2, 2, 79, 77, 3, 2, 2, 2, 79, 80,
	3, 2, 2, 2, 80, 82, 3, 2, 2, 2, 81, 79, 3, 2, 2, 2, 82, 83, 7, 27, 2, 2,
	83, 9, 3, 2, 2, 2, 84, 85, 7, 36, 2, 2, 85, 86, 7, 26, 2, 2, 86, 87, 7,
	20, 2, 2, 87, 88, 7, 20, 2, 2, 88, 89, 7, 27, 2, 2, 89, 11, 3, 2, 2, 2,
	90, 91, 7, 37, 2, 2, 91, 96, 7, 26, 2, 2, 92, 97, 5, 14, 8, 2, 93, 97,
	5, 16, 9, 2, 94, 97, 5, 18, 10, 2, 95, 97, 5, 20, 11, 2, 96, 92, 3, 2,
	2, 2, 96, 93, 3, 2, 2, 2, 96, 94, 3, 2, 2, 2, 96, 95, 3, 2, 2, 2, 97, 98,
	3, 2, 2, 2, 98, 99, 7, 27, 2, 2, 99, 13, 3, 2, 2, 2, 100, 115, 7, 45, 2,
	2, 101, 105, 7, 26, 2, 2, 102, 104, 7, 20, 2, 2, 103, 102, 3, 2, 2, 2,
	104, 107, 3, 2, 2, 2, 105, 103, 3, 2, 2, 2, 105, 106, 3, 2, 2, 2, 106,
	111, 3, 2, 2, 2, 107, 105, 3, 2, 2, 2, 108, 110, 5, 22, 12, 2, 109, 108,
	3, 2, 2, 2, 110, 113, 3, 2, 2, 2, 111, 109, 3, 2, 2, 2, 111, 112, 3, 2,
	2, 2, 112, 114, 3, 2, 2, 2, 113, 111, 3, 2, 2, 2, 114, 116, 7, 27, 2, 2,
	115, 101, 3, 2, 2, 2, 115, 116, 3, 2, 2, 2, 116, 15, 3, 2, 2, 2, 117, 132,
	7, 46, 2, 2, 118, 122, 7, 26, 2, 2, 119, 121, 7, 20, 2, 2, 120, 119, 3,
	2, 2, 2, 121, 124, 3, 2, 2, 2, 122, 120, 3, 2, 2, 2, 122, 123, 3, 2, 2,
	2, 123, 128, 3, 2, 2, 2, 124, 122, 3, 2, 2, 2, 125, 127, 5, 22, 12, 2,
	126, 125, 3, 2, 2, 2, 127, 130, 3, 2, 2, 2, 128, 126, 3, 2, 2, 2, 128,
	129, 3, 2, 2, 2, 129, 131, 3, 2, 2, 2, 130, 128, 3, 2, 2, 2, 131, 133,
	7, 27, 2, 2, 132, 118, 3, 2, 2, 2, 132, 133, 3, 2, 2, 2, 133, 17, 3, 2,
	2, 2, 134, 135, 7, 38, 2, 2, 135, 139, 7, 26, 2, 2, 136, 138, 7, 20, 2,
	2, 137, 136, 3, 2, 2, 2, 138, 141, 3, 2, 2, 2, 139, 137, 3, 2, 2, 2, 139,
	140, 3, 2, 2, 2, 140, 142, 3, 2, 2, 2, 141, 139, 3, 2, 2, 2, 142, 143,
	5, 24, 13, 2, 143, 144, 7, 27, 2, 2, 144, 19, 3, 2, 2, 2, 145, 146, 7,
	39, 2, 2, 146, 150, 7, 26, 2, 2, 147, 149, 7, 20, 2, 2, 148, 147, 3, 2,
	2, 2, 149, 152, 3, 2, 2, 2, 150, 148, 3, 2, 2, 2, 150, 151, 3, 2, 2, 2,
	151, 153, 3, 2, 2, 2, 152, 150, 3, 2, 2, 2, 153, 155, 5, 24, 13, 2, 154,
	156, 5, 28, 15, 2, 155, 154, 3, 2, 2, 2, 155, 156, 3, 2, 2, 2, 156, 157,
	3, 2, 2, 2, 157, 158, 7, 27, 2, 2, 158, 21, 3, 2, 2, 2, 159, 160, 7, 52,
	2, 2, 160, 161, 7, 26, 2, 2, 161, 162, 7, 20, 2, 2, 162, 163, 7, 20, 2,
	2, 163, 165, 5, 24, 13, 2, 164, 166, 5, 28, 15, 2, 165, 164, 3, 2, 2, 2,
	165, 166, 3, 2, 2, 2, 166, 172, 3, 2, 2, 2, 167, 168, 5, 10, 6, 2, 168,
	169, 5, 28, 15, 2, 169, 171, 3, 2, 2, 2, 170, 167, 3, 2, 2, 2, 171, 174,
	3, 2, 2, 2, 172, 170, 3, 2, 2, 2, 172, 173, 3, 2, 2, 2, 173, 175, 3, 2,
	2, 2, 174, 172, 3, 2, 2, 2, 175, 176, 7, 27, 2, 2, 176, 23, 3, 2, 2, 2,
	177, 178, 7, 40, 2, 2, 178, 179, 7, 26, 2, 2, 179, 183, 5, 26, 14, 2, 180,
	182, 5, 24, 13, 2, 181, 180, 3, 2, 2, 2, 182, 185, 3, 2, 2, 2, 183, 181,
	3, 2, 2, 2, 183, 184, 3, 2, 2, 2, 184, 186, 3, 2, 2, 2, 185, 183, 3, 2,
	2, 2, 186, 187, 7, 27, 2, 2, 187, 25, 3, 2, 2, 2, 188, 189, 7, 44, 2, 2,
	189, 194, 7, 20, 2, 2, 190, 191, 7, 49, 2, 2, 191, 194, 7, 20, 2, 2, 192,
	194, 5, 10, 6, 2, 193, 188, 3, 2, 2, 2, 193, 190, 3, 2, 2, 2, 193, 192,
	3, 2, 2, 2, 194, 27, 3, 2, 2, 2, 195, 224, 7, 54, 2, 2, 196, 224, 7, 55,
	2, 2, 197, 224, 7, 56, 2, 2, 198, 224, 7, 57, 2, 2, 199, 224, 7, 58, 2,
	2, 200, 209, 7, 59, 2, 2, 201, 203, 7, 26, 2, 2, 202, 204, 5, 28, 15, 2,
	203, 202, 3, 2, 2, 2, 204, 205, 3, 2, 2, 2, 205, 203, 3, 2, 2, 2, 205,
	206, 3, 2, 2, 2, 206, 207, 3, 2, 2, 2, 207, 208, 7, 27, 2, 2, 208, 210,
	3, 2, 2, 2, 209, 201, 3, 2, 2, 2, 209, 210, 3, 2, 2, 2, 210, 224, 3, 2,
	2, 2, 211, 220, 7, 60, 2, 2, 212, 214, 7, 26, 2, 2, 213, 215, 5, 28, 15,
	2, 214, 213, 3, 2, 2, 2, 215, 216, 3, 2, 2, 2, 216, 214, 3, 2, 2, 2, 216,
	217, 3, 2, 2, 2, 217, 218, 3, 2, 2, 2, 218, 219, 7, 27, 2, 2, 219, 221,
	3, 2, 2, 2, 220, 212, 3, 2, 2, 2, 220, 221, 3, 2, 2, 2, 221, 224, 3, 2,
	2, 2, 222, 224, 7, 29, 2, 2, 223, 195, 3, 2, 2, 2, 223, 196, 3, 2, 2, 2,
	223, 197, 3, 2, 2, 2, 223, 198, 3, 2, 2, 2, 223, 199, 3, 2, 2, 2, 223,
	200, 3, 2, 2, 2, 223, 211, 3, 2, 2, 2, 223, 222, 3, 2, 2, 2, 224, 29, 3,
	2, 2, 2, 27, 39, 46, 54, 63, 71, 79, 96, 105, 111, 115, 122, 128, 132,
	139, 150, 155, 165, 172, 183, 193, 205, 209, 216, 220, 223,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'{'", "'}'", "'['", "']'", "'='", "'\"'", "'''", "';'", "'::'", "':'",
	"'.'", "','", "'<'", "'>'", "'*'", "'@'",
}
var symbolicNames = []string{
	"", "LCUR", "RCUR", "LSQ", "RSQ", "EQ", "DQ", "SQ", "SEMI", "DCOLON", "COLON",
	"DOT", "COMMA", "LCHEVR", "RCHEVR", "STAR", "AT", "STR", "ID", "INT", "FLT",
	"WS", "LINE_DOC", "LINE_COMMENT", "DOWN", "UP", "ROOT", "ERROR", "ADL",
	"Module", "Import", "ImportModule", "ImportScopedName", "Decl", "ScopedName",
	"DeclType", "TypeDef", "NewType", "TypeExpr", "Annotation", "AnnotationNotScoped",
	"AnnotationScoped", "Primitive", "Struct", "Union", "Newtype", "Type",
	"TypeParam", "TypeExprSimple", "TypeExprGeneric", "Field", "Json", "JsonStr",
	"JsonBool", "JsonNull", "JsonInt", "JsonFloat", "JsonArray", "JsonObj",
	"JsonObjKey", "ModuleAnno", "DeclAnno", "FieldAnno",
}

var ruleNames = []string{
	"adl", "module", "import_", "decl", "scopedName", "declType", "struct_",
	"union_", "typeDef", "newType", "field", "typeExpr", "typeRef", "json",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type AdlAst struct {
	*antlr.BaseParser
}

func NewAdlAst(input antlr.TokenStream) *AdlAst {
	this := new(AdlAst)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "AdlAst.g4"

	return this
}

// AdlAst tokens.
const (
	AdlAstEOF                 = antlr.TokenEOF
	AdlAstLCUR                = 1
	AdlAstRCUR                = 2
	AdlAstLSQ                 = 3
	AdlAstRSQ                 = 4
	AdlAstEQ                  = 5
	AdlAstDQ                  = 6
	AdlAstSQ                  = 7
	AdlAstSEMI                = 8
	AdlAstDCOLON              = 9
	AdlAstCOLON               = 10
	AdlAstDOT                 = 11
	AdlAstCOMMA               = 12
	AdlAstLCHEVR              = 13
	AdlAstRCHEVR              = 14
	AdlAstSTAR                = 15
	AdlAstAT                  = 16
	AdlAstSTR                 = 17
	AdlAstID                  = 18
	AdlAstINT                 = 19
	AdlAstFLT                 = 20
	AdlAstWS                  = 21
	AdlAstLINE_DOC            = 22
	AdlAstLINE_COMMENT        = 23
	AdlAstDOWN                = 24
	AdlAstUP                  = 25
	AdlAstROOT                = 26
	AdlAstERROR               = 27
	AdlAstADL                 = 28
	AdlAstModule              = 29
	AdlAstImport              = 30
	AdlAstImportModule        = 31
	AdlAstImportScopedName    = 32
	AdlAstDecl                = 33
	AdlAstScopedName          = 34
	AdlAstDeclType            = 35
	AdlAstTypeDef             = 36
	AdlAstNewType             = 37
	AdlAstTypeExpr            = 38
	AdlAstAnnotation          = 39
	AdlAstAnnotationNotScoped = 40
	AdlAstAnnotationScoped    = 41
	AdlAstPrimitive           = 42
	AdlAstStruct              = 43
	AdlAstUnion               = 44
	AdlAstNewtype             = 45
	AdlAstType                = 46
	AdlAstTypeParam           = 47
	AdlAstTypeExprSimple      = 48
	AdlAstTypeExprGeneric     = 49
	AdlAstField               = 50
	AdlAstJson                = 51
	AdlAstJsonStr             = 52
	AdlAstJsonBool            = 53
	AdlAstJsonNull            = 54
	AdlAstJsonInt             = 55
	AdlAstJsonFloat           = 56
	AdlAstJsonArray           = 57
	AdlAstJsonObj             = 58
	AdlAstJsonObjKey          = 59
	AdlAstModuleAnno          = 60
	AdlAstDeclAnno            = 61
	AdlAstFieldAnno           = 62
)

// AdlAst rules.
const (
	AdlAstRULE_adl        = 0
	AdlAstRULE_module     = 1
	AdlAstRULE_import_    = 2
	AdlAstRULE_decl       = 3
	AdlAstRULE_scopedName = 4
	AdlAstRULE_declType   = 5
	AdlAstRULE_struct_    = 6
	AdlAstRULE_union_     = 7
	AdlAstRULE_typeDef    = 8
	AdlAstRULE_newType    = 9
	AdlAstRULE_field      = 10
	AdlAstRULE_typeExpr   = 11
	AdlAstRULE_typeRef    = 12
	AdlAstRULE_json       = 13
)

type IAdlContext interface {
	antlr.ParserRuleContext
	// GetParser returns the parser.
	GetParser() antlr.Parser

	// start internal
	//  ruleGetterDecl
	Module() IModuleContext
	//  ruleListGetterDecl
	//  ruleListIndexedGetterDecl
	//  ruleContextDecls
	//  ruleContextListDecls
	// end internal
	//Gets for labeled elements
	//tokenDecls
	//tokenTypeDecls
	//tokenListDecls
	//attributeDecls
	//tokenGetterDecl
	EOF() antlr.TerminalNode
	//tokenListGetterDecl
	//tokenListIndexedGetterDecl

	// IsAdlContext differentiates from other interfaces.
	//copyStruct,GetRuleContext and ToStringTree  from embedded

	//<if(dispatchMethods)>
	//<dispatchMethods; separator="\n\n">
	//<endif>

	//<if(extensionMembers)>
	//<extensionMembers; separator="\n\n">
	//<endif>
}

type AdlContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAdlContext() *AdlContext {
	var p = new(AdlContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AdlAstRULE_adl
	return p
}

func (*AdlContext) IsAdlContext() {}

func NewAdlContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AdlContext {
	var p = new(AdlContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AdlAstRULE_adl

	return p
}

func (s *AdlContext) GetParser() antlr.Parser { return s.parser }

//StructDecl tokenDecls

//StructDecl tokenTypeDecls

//StructDecl tokenListDecls

//StructDecl ruleContextDecls

//StructDecl ruleContextListDecls

//StructDecl attributeDecls

// Getters
func (s *AdlContext) Module() IModuleContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ModuleContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IModuleContext)
}

func (s *AdlContext) EOF() antlr.TerminalNode {
	return s.GetToken(AdlAstEOF, 0)
}

//provideCopyFrom
func (s *AdlContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AdlContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

//dispatchMethods
func (s *AdlContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AdlEntryListener); ok {
		listenerT.EnterAdl(s)
	}
}

func (s *AdlContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AdlExitListener); ok {
		listenerT.ExitAdl(s)
	}
}

func (s *AdlContext) VisitFunc(hdls antlr.ParserTreeVisitorHandlers, args ...interface{}) (result interface{}) {
	h := hdls.(*AdlAstHandlers)
	if h.EnterEveryRule != nil {
		h.EnterEveryRule(s)
	}
	if h.Adl != nil {
		h.Adl(s, h, args...)
	} else {
		s.VisitChildrenFunc(h, args...)
	}
	if h.ExitEveryRule != nil {
		h.ExitEveryRule(s)
	}
	return
}

func (s *AdlContext) Visit(delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	switch t := delegate.(type) {
	case AdlContextVisitor:
		return t.VisitAdl(s, delegate, args...)
	default:
		return delegate.VisitChildren(s, delegate, args...)
	}
}

//extensionMembers

func (p *AdlAst) Adl() (localctx IAdlContext) {
	localctx = NewAdlContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, AdlAstRULE_adl)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(28)
		p.Module()
	}
	{
		p.SetState(29)
		p.Match(AdlAstEOF)
	}

	return localctx
}

type IModuleContext interface {
	antlr.ParserRuleContext
	// GetParser returns the parser.
	GetParser() antlr.Parser

	// start internal
	//  ruleGetterDecl
	//  ruleListGetterDecl
	AllImport_() []IImport_Context
	AllDecl() []IDeclContext
	AllScopedName() []IScopedNameContext
	AllJson() []IJsonContext
	//  ruleListIndexedGetterDecl
	Import_(i int) IImport_Context
	Decl(i int) IDeclContext
	ScopedName(i int) IScopedNameContext
	Json(i int) IJsonContext
	//  ruleContextDecls
	//  ruleContextListDecls
	// end internal
	//Gets for labeled elements
	//tokenDecls
	GetModuleName() antlr.Token

	//tokenTypeDecls
	//tokenListDecls
	//attributeDecls
	//tokenGetterDecl
	Module() antlr.TerminalNode
	DOWN() antlr.TerminalNode
	UP() antlr.TerminalNode
	//tokenListGetterDecl
	AllID() []antlr.TerminalNode
	//tokenListIndexedGetterDecl
	ID(i int) antlr.TerminalNode

	// IsModuleContext differentiates from other interfaces.
	//copyStruct,GetRuleContext and ToStringTree  from embedded

	//<if(dispatchMethods)>
	//<dispatchMethods; separator="\n\n">
	//<endif>

	//<if(extensionMembers)>
	//<extensionMembers; separator="\n\n">
	//<endif>
}

type ModuleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	//TokenDecl
	moduleName antlr.Token
}

func NewEmptyModuleContext() *ModuleContext {
	var p = new(ModuleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AdlAstRULE_module
	return p
}

func (*ModuleContext) IsModuleContext() {}

func NewModuleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ModuleContext {
	var p = new(ModuleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AdlAstRULE_module

	return p
}

func (s *ModuleContext) GetParser() antlr.Parser { return s.parser }

//StructDecl tokenDecls
func (s *ModuleContext) GetModuleName() antlr.Token  { return s.moduleName }
func (s *ModuleContext) SetModuleName(v antlr.Token) { s.moduleName = v }

//StructDecl tokenTypeDecls

//StructDecl tokenListDecls

//StructDecl ruleContextDecls

//StructDecl ruleContextListDecls

//StructDecl attributeDecls

// Getters
func (s *ModuleContext) Module() antlr.TerminalNode {
	return s.GetToken(AdlAstModule, 0)
}

func (s *ModuleContext) DOWN() antlr.TerminalNode {
	return s.GetToken(AdlAstDOWN, 0)
}

func (s *ModuleContext) UP() antlr.TerminalNode {
	return s.GetToken(AdlAstUP, 0)
}

func (s *ModuleContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(AdlAstID)
}

func (s *ModuleContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(AdlAstID, i)
}

func (s *ModuleContext) AllImport_() []IImport_Context {
	//ContextRuleListGetterDecl
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*Import_Context)(nil)).Elem())
	var tst = make([]IImport_Context, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IImport_Context)
		}
	}

	return tst
}

func (s *ModuleContext) Import_(i int) IImport_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*Import_Context)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IImport_Context)
}

func (s *ModuleContext) AllDecl() []IDeclContext {
	//ContextRuleListGetterDecl
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*DeclContext)(nil)).Elem())
	var tst = make([]IDeclContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDeclContext)
		}
	}

	return tst
}

func (s *ModuleContext) Decl(i int) IDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*DeclContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDeclContext)
}

func (s *ModuleContext) AllScopedName() []IScopedNameContext {
	//ContextRuleListGetterDecl
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ScopedNameContext)(nil)).Elem())
	var tst = make([]IScopedNameContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IScopedNameContext)
		}
	}

	return tst
}

func (s *ModuleContext) ScopedName(i int) IScopedNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ScopedNameContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IScopedNameContext)
}

func (s *ModuleContext) AllJson() []IJsonContext {
	//ContextRuleListGetterDecl
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*JsonContext)(nil)).Elem())
	var tst = make([]IJsonContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IJsonContext)
		}
	}

	return tst
}

func (s *ModuleContext) Json(i int) IJsonContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*JsonContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IJsonContext)
}

//provideCopyFrom
func (s *ModuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ModuleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

//dispatchMethods
func (s *ModuleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ModuleEntryListener); ok {
		listenerT.EnterModule(s)
	}
}

func (s *ModuleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ModuleExitListener); ok {
		listenerT.ExitModule(s)
	}
}

func (s *ModuleContext) VisitFunc(hdls antlr.ParserTreeVisitorHandlers, args ...interface{}) (result interface{}) {
	h := hdls.(*AdlAstHandlers)
	if h.EnterEveryRule != nil {
		h.EnterEveryRule(s)
	}
	if h.Module != nil {
		h.Module(s, h, args...)
	} else {
		s.VisitChildrenFunc(h, args...)
	}
	if h.ExitEveryRule != nil {
		h.ExitEveryRule(s)
	}
	return
}

func (s *ModuleContext) Visit(delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	switch t := delegate.(type) {
	case ModuleContextVisitor:
		return t.VisitModule(s, delegate, args...)
	default:
		return delegate.VisitChildren(s, delegate, args...)
	}
}

//extensionMembers

func (p *AdlAst) Module() (localctx IModuleContext) {
	localctx = NewModuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, AdlAstRULE_module)
	var //TokenTypeDecl
	_la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(31)
		p.Match(AdlAstModule)
	}
	{
		p.SetState(32)
		p.Match(AdlAstDOWN)
	}
	{
		p.SetState(33)
		var _m = p.Match(AdlAstID)
		localctx.(*ModuleContext).moduleName = _m

	}
	p.SetState(37)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == AdlAstImport {
		{
			p.SetState(34)
			p.Import_()
		}

		p.SetState(39)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(44)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == AdlAstID {
		{
			p.SetState(40)
			p.Match(AdlAstID)
		}
		{
			p.SetState(41)
			p.Decl()
		}

		p.SetState(46)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(52)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == AdlAstScopedName {
		{
			p.SetState(47)
			p.ScopedName()
		}
		{
			p.SetState(48)
			p.Json()
		}

		p.SetState(54)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(55)
		p.Match(AdlAstUP)
	}

	return localctx
}

type IImport_Context interface {
	antlr.ParserRuleContext
	// GetParser returns the parser.
	GetParser() antlr.Parser

	// start internal
	//  ruleGetterDecl
	ScopedName() IScopedNameContext
	//  ruleListGetterDecl
	//  ruleListIndexedGetterDecl
	//  ruleContextDecls
	//  ruleContextListDecls
	// end internal
	//Gets for labeled elements
	//tokenDecls
	GetModuleName() antlr.Token

	//tokenTypeDecls
	//tokenListDecls
	//attributeDecls
	//tokenGetterDecl
	Import() antlr.TerminalNode
	DOWN() antlr.TerminalNode
	UP() antlr.TerminalNode
	ID() antlr.TerminalNode
	//tokenListGetterDecl
	//tokenListIndexedGetterDecl

	// IsImport_Context differentiates from other interfaces.
	//copyStruct,GetRuleContext and ToStringTree  from embedded

	//<if(dispatchMethods)>
	//<dispatchMethods; separator="\n\n">
	//<endif>

	//<if(extensionMembers)>
	//<extensionMembers; separator="\n\n">
	//<endif>
}

type Import_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	//TokenDecl
	moduleName antlr.Token
}

func NewEmptyImport_Context() *Import_Context {
	var p = new(Import_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AdlAstRULE_import_
	return p
}

func (*Import_Context) IsImport_Context() {}

func NewImport_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Import_Context {
	var p = new(Import_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AdlAstRULE_import_

	return p
}

func (s *Import_Context) GetParser() antlr.Parser { return s.parser }

//StructDecl tokenDecls
func (s *Import_Context) GetModuleName() antlr.Token  { return s.moduleName }
func (s *Import_Context) SetModuleName(v antlr.Token) { s.moduleName = v }

//StructDecl tokenTypeDecls

//StructDecl tokenListDecls

//StructDecl ruleContextDecls

//StructDecl ruleContextListDecls

//StructDecl attributeDecls

// Getters
func (s *Import_Context) Import() antlr.TerminalNode {
	return s.GetToken(AdlAstImport, 0)
}

func (s *Import_Context) DOWN() antlr.TerminalNode {
	return s.GetToken(AdlAstDOWN, 0)
}

func (s *Import_Context) UP() antlr.TerminalNode {
	return s.GetToken(AdlAstUP, 0)
}

func (s *Import_Context) ScopedName() IScopedNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ScopedNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IScopedNameContext)
}

func (s *Import_Context) ID() antlr.TerminalNode {
	return s.GetToken(AdlAstID, 0)
}

//provideCopyFrom
func (s *Import_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Import_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

//dispatchMethods
func (s *Import_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Import_EntryListener); ok {
		listenerT.EnterImport_(s)
	}
}

func (s *Import_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Import_ExitListener); ok {
		listenerT.ExitImport_(s)
	}
}

func (s *Import_Context) VisitFunc(hdls antlr.ParserTreeVisitorHandlers, args ...interface{}) (result interface{}) {
	h := hdls.(*AdlAstHandlers)
	if h.EnterEveryRule != nil {
		h.EnterEveryRule(s)
	}
	if h.Import_ != nil {
		h.Import_(s, h, args...)
	} else {
		s.VisitChildrenFunc(h, args...)
	}
	if h.ExitEveryRule != nil {
		h.ExitEveryRule(s)
	}
	return
}

func (s *Import_Context) Visit(delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	switch t := delegate.(type) {
	case Import_ContextVisitor:
		return t.VisitImport_(s, delegate, args...)
	default:
		return delegate.VisitChildren(s, delegate, args...)
	}
}

//extensionMembers

func (p *AdlAst) Import_() (localctx IImport_Context) {
	localctx = NewImport_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, AdlAstRULE_import_)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(57)
		p.Match(AdlAstImport)
	}
	{
		p.SetState(58)
		p.Match(AdlAstDOWN)
	}
	p.SetState(61)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AdlAstID:
		{
			p.SetState(59)
			var _m = p.Match(AdlAstID)
			localctx.(*Import_Context).moduleName = _m

		}

	case AdlAstScopedName:
		{
			p.SetState(60)
			p.ScopedName()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(63)
		p.Match(AdlAstUP)
	}

	return localctx
}

type IDeclContext interface {
	antlr.ParserRuleContext
	// GetParser returns the parser.
	GetParser() antlr.Parser

	// start internal
	//  ruleGetterDecl
	DeclType() IDeclTypeContext
	//  ruleListGetterDecl
	AllScopedName() []IScopedNameContext
	AllJson() []IJsonContext
	//  ruleListIndexedGetterDecl
	ScopedName(i int) IScopedNameContext
	Json(i int) IJsonContext
	//  ruleContextDecls
	GetType_() IDeclTypeContext

	//  ruleContextListDecls
	// end internal
	//Gets for labeled elements
	//tokenDecls
	GetName() antlr.Token
	GetVersion() antlr.Token

	//tokenTypeDecls
	//tokenListDecls
	//attributeDecls
	//tokenGetterDecl
	Decl() antlr.TerminalNode
	DOWN() antlr.TerminalNode
	UP() antlr.TerminalNode
	ID() antlr.TerminalNode
	INT() antlr.TerminalNode
	//tokenListGetterDecl
	//tokenListIndexedGetterDecl

	// IsDeclContext differentiates from other interfaces.
	//copyStruct,GetRuleContext and ToStringTree  from embedded

	//<if(dispatchMethods)>
	//<dispatchMethods; separator="\n\n">
	//<endif>

	//<if(extensionMembers)>
	//<extensionMembers; separator="\n\n">
	//<endif>
}

type DeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	//TokenDecl
	name antlr.Token
	//TokenDecl
	version antlr.Token
	//RuleContextDecl
	type_ IDeclTypeContext
}

func NewEmptyDeclContext() *DeclContext {
	var p = new(DeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AdlAstRULE_decl
	return p
}

func (*DeclContext) IsDeclContext() {}

func NewDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclContext {
	var p = new(DeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AdlAstRULE_decl

	return p
}

func (s *DeclContext) GetParser() antlr.Parser { return s.parser }

//StructDecl tokenDecls
func (s *DeclContext) GetName() antlr.Token  { return s.name }
func (s *DeclContext) SetName(v antlr.Token) { s.name = v }

func (s *DeclContext) GetVersion() antlr.Token  { return s.version }
func (s *DeclContext) SetVersion(v antlr.Token) { s.version = v }

//StructDecl tokenTypeDecls

//StructDecl tokenListDecls

//StructDecl ruleContextDecls
func (s *DeclContext) GetType_() IDeclTypeContext  { return s.type_ }
func (s *DeclContext) SetType_(v IDeclTypeContext) { s.type_ = v }

//StructDecl ruleContextListDecls

//StructDecl attributeDecls

// Getters
func (s *DeclContext) Decl() antlr.TerminalNode {
	return s.GetToken(AdlAstDecl, 0)
}

func (s *DeclContext) DOWN() antlr.TerminalNode {
	return s.GetToken(AdlAstDOWN, 0)
}

func (s *DeclContext) UP() antlr.TerminalNode {
	return s.GetToken(AdlAstUP, 0)
}

func (s *DeclContext) ID() antlr.TerminalNode {
	return s.GetToken(AdlAstID, 0)
}

func (s *DeclContext) DeclType() IDeclTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*DeclTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDeclTypeContext)
}

func (s *DeclContext) AllScopedName() []IScopedNameContext {
	//ContextRuleListGetterDecl
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ScopedNameContext)(nil)).Elem())
	var tst = make([]IScopedNameContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IScopedNameContext)
		}
	}

	return tst
}

func (s *DeclContext) ScopedName(i int) IScopedNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ScopedNameContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IScopedNameContext)
}

func (s *DeclContext) AllJson() []IJsonContext {
	//ContextRuleListGetterDecl
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*JsonContext)(nil)).Elem())
	var tst = make([]IJsonContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IJsonContext)
		}
	}

	return tst
}

func (s *DeclContext) Json(i int) IJsonContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*JsonContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IJsonContext)
}

func (s *DeclContext) INT() antlr.TerminalNode {
	return s.GetToken(AdlAstINT, 0)
}

//provideCopyFrom
func (s *DeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

//dispatchMethods
func (s *DeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DeclEntryListener); ok {
		listenerT.EnterDecl(s)
	}
}

func (s *DeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DeclExitListener); ok {
		listenerT.ExitDecl(s)
	}
}

func (s *DeclContext) VisitFunc(hdls antlr.ParserTreeVisitorHandlers, args ...interface{}) (result interface{}) {
	h := hdls.(*AdlAstHandlers)
	if h.EnterEveryRule != nil {
		h.EnterEveryRule(s)
	}
	if h.Decl != nil {
		h.Decl(s, h, args...)
	} else {
		s.VisitChildrenFunc(h, args...)
	}
	if h.ExitEveryRule != nil {
		h.ExitEveryRule(s)
	}
	return
}

func (s *DeclContext) Visit(delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	switch t := delegate.(type) {
	case DeclContextVisitor:
		return t.VisitDecl(s, delegate, args...)
	default:
		return delegate.VisitChildren(s, delegate, args...)
	}
}

//extensionMembers

func (p *AdlAst) Decl() (localctx IDeclContext) {
	localctx = NewDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, AdlAstRULE_decl)
	var //TokenTypeDecl
	_la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(65)
		p.Match(AdlAstDecl)
	}
	{
		p.SetState(66)
		p.Match(AdlAstDOWN)
	}
	{
		p.SetState(67)
		var _m = p.Match(AdlAstID)
		localctx.(*DeclContext).name = _m

	}
	p.SetState(69)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == AdlAstINT {
		{
			p.SetState(68)
			var _m = p.Match(AdlAstINT)
			localctx.(*DeclContext).version = _m

		}

	}
	{
		p.SetState(71)

		var _x = p.DeclType()

		localctx.(*DeclContext).type_ = _x

	}
	p.SetState(77)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == AdlAstScopedName {
		{
			p.SetState(72)
			p.ScopedName()
		}
		{
			p.SetState(73)
			p.Json()
		}

		p.SetState(79)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(80)
		p.Match(AdlAstUP)
	}

	return localctx
}

type IScopedNameContext interface {
	antlr.ParserRuleContext
	// GetParser returns the parser.
	GetParser() antlr.Parser

	// start internal
	//  ruleGetterDecl
	//  ruleListGetterDecl
	//  ruleListIndexedGetterDecl
	//  ruleContextDecls
	//  ruleContextListDecls
	// end internal
	//Gets for labeled elements
	//tokenDecls
	GetModuleName() antlr.Token
	GetName() antlr.Token

	//tokenTypeDecls
	//tokenListDecls
	//attributeDecls
	//tokenGetterDecl
	ScopedName() antlr.TerminalNode
	DOWN() antlr.TerminalNode
	UP() antlr.TerminalNode
	//tokenListGetterDecl
	AllID() []antlr.TerminalNode
	//tokenListIndexedGetterDecl
	ID(i int) antlr.TerminalNode

	// IsScopedNameContext differentiates from other interfaces.
	//copyStruct,GetRuleContext and ToStringTree  from embedded

	//<if(dispatchMethods)>
	//<dispatchMethods; separator="\n\n">
	//<endif>

	//<if(extensionMembers)>
	//<extensionMembers; separator="\n\n">
	//<endif>
}

type ScopedNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	//TokenDecl
	moduleName antlr.Token
	//TokenDecl
	name antlr.Token
}

func NewEmptyScopedNameContext() *ScopedNameContext {
	var p = new(ScopedNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AdlAstRULE_scopedName
	return p
}

func (*ScopedNameContext) IsScopedNameContext() {}

func NewScopedNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ScopedNameContext {
	var p = new(ScopedNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AdlAstRULE_scopedName

	return p
}

func (s *ScopedNameContext) GetParser() antlr.Parser { return s.parser }

//StructDecl tokenDecls
func (s *ScopedNameContext) GetModuleName() antlr.Token  { return s.moduleName }
func (s *ScopedNameContext) SetModuleName(v antlr.Token) { s.moduleName = v }

func (s *ScopedNameContext) GetName() antlr.Token  { return s.name }
func (s *ScopedNameContext) SetName(v antlr.Token) { s.name = v }

//StructDecl tokenTypeDecls

//StructDecl tokenListDecls

//StructDecl ruleContextDecls

//StructDecl ruleContextListDecls

//StructDecl attributeDecls

// Getters
func (s *ScopedNameContext) ScopedName() antlr.TerminalNode {
	return s.GetToken(AdlAstScopedName, 0)
}

func (s *ScopedNameContext) DOWN() antlr.TerminalNode {
	return s.GetToken(AdlAstDOWN, 0)
}

func (s *ScopedNameContext) UP() antlr.TerminalNode {
	return s.GetToken(AdlAstUP, 0)
}

func (s *ScopedNameContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(AdlAstID)
}

func (s *ScopedNameContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(AdlAstID, i)
}

//provideCopyFrom
func (s *ScopedNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ScopedNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

//dispatchMethods
func (s *ScopedNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScopedNameEntryListener); ok {
		listenerT.EnterScopedName(s)
	}
}

func (s *ScopedNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ScopedNameExitListener); ok {
		listenerT.ExitScopedName(s)
	}
}

func (s *ScopedNameContext) VisitFunc(hdls antlr.ParserTreeVisitorHandlers, args ...interface{}) (result interface{}) {
	h := hdls.(*AdlAstHandlers)
	if h.EnterEveryRule != nil {
		h.EnterEveryRule(s)
	}
	if h.ScopedName != nil {
		h.ScopedName(s, h, args...)
	} else {
		s.VisitChildrenFunc(h, args...)
	}
	if h.ExitEveryRule != nil {
		h.ExitEveryRule(s)
	}
	return
}

func (s *ScopedNameContext) Visit(delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	switch t := delegate.(type) {
	case ScopedNameContextVisitor:
		return t.VisitScopedName(s, delegate, args...)
	default:
		return delegate.VisitChildren(s, delegate, args...)
	}
}

//extensionMembers

func (p *AdlAst) ScopedName() (localctx IScopedNameContext) {
	localctx = NewScopedNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, AdlAstRULE_scopedName)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(82)
		p.Match(AdlAstScopedName)
	}
	{
		p.SetState(83)
		p.Match(AdlAstDOWN)
	}
	{
		p.SetState(84)
		var _m = p.Match(AdlAstID)
		localctx.(*ScopedNameContext).moduleName = _m

	}
	{
		p.SetState(85)
		var _m = p.Match(AdlAstID)
		localctx.(*ScopedNameContext).name = _m

	}
	{
		p.SetState(86)
		p.Match(AdlAstUP)
	}

	return localctx
}

type IDeclTypeContext interface {
	antlr.ParserRuleContext
	// GetParser returns the parser.
	GetParser() antlr.Parser

	// start internal
	//  ruleGetterDecl
	Struct_() IStruct_Context
	Union_() IUnion_Context
	TypeDef() ITypeDefContext
	NewType() INewTypeContext
	//  ruleListGetterDecl
	//  ruleListIndexedGetterDecl
	//  ruleContextDecls
	//  ruleContextListDecls
	// end internal
	//Gets for labeled elements
	//tokenDecls
	//tokenTypeDecls
	//tokenListDecls
	//attributeDecls
	//tokenGetterDecl
	DeclType() antlr.TerminalNode
	DOWN() antlr.TerminalNode
	UP() antlr.TerminalNode
	//tokenListGetterDecl
	//tokenListIndexedGetterDecl

	// IsDeclTypeContext differentiates from other interfaces.
	//copyStruct,GetRuleContext and ToStringTree  from embedded

	//<if(dispatchMethods)>
	//<dispatchMethods; separator="\n\n">
	//<endif>

	//<if(extensionMembers)>
	//<extensionMembers; separator="\n\n">
	//<endif>
}

type DeclTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclTypeContext() *DeclTypeContext {
	var p = new(DeclTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AdlAstRULE_declType
	return p
}

func (*DeclTypeContext) IsDeclTypeContext() {}

func NewDeclTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclTypeContext {
	var p = new(DeclTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AdlAstRULE_declType

	return p
}

func (s *DeclTypeContext) GetParser() antlr.Parser { return s.parser }

//StructDecl tokenDecls

//StructDecl tokenTypeDecls

//StructDecl tokenListDecls

//StructDecl ruleContextDecls

//StructDecl ruleContextListDecls

//StructDecl attributeDecls

// Getters
func (s *DeclTypeContext) DeclType() antlr.TerminalNode {
	return s.GetToken(AdlAstDeclType, 0)
}

func (s *DeclTypeContext) DOWN() antlr.TerminalNode {
	return s.GetToken(AdlAstDOWN, 0)
}

func (s *DeclTypeContext) UP() antlr.TerminalNode {
	return s.GetToken(AdlAstUP, 0)
}

func (s *DeclTypeContext) Struct_() IStruct_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*Struct_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStruct_Context)
}

func (s *DeclTypeContext) Union_() IUnion_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*Union_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUnion_Context)
}

func (s *DeclTypeContext) TypeDef() ITypeDefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*TypeDefContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeDefContext)
}

func (s *DeclTypeContext) NewType() INewTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*NewTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INewTypeContext)
}

//provideCopyFrom
func (s *DeclTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

//dispatchMethods
func (s *DeclTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DeclTypeEntryListener); ok {
		listenerT.EnterDeclType(s)
	}
}

func (s *DeclTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DeclTypeExitListener); ok {
		listenerT.ExitDeclType(s)
	}
}

func (s *DeclTypeContext) VisitFunc(hdls antlr.ParserTreeVisitorHandlers, args ...interface{}) (result interface{}) {
	h := hdls.(*AdlAstHandlers)
	if h.EnterEveryRule != nil {
		h.EnterEveryRule(s)
	}
	if h.DeclType != nil {
		h.DeclType(s, h, args...)
	} else {
		s.VisitChildrenFunc(h, args...)
	}
	if h.ExitEveryRule != nil {
		h.ExitEveryRule(s)
	}
	return
}

func (s *DeclTypeContext) Visit(delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	switch t := delegate.(type) {
	case DeclTypeContextVisitor:
		return t.VisitDeclType(s, delegate, args...)
	default:
		return delegate.VisitChildren(s, delegate, args...)
	}
}

//extensionMembers

func (p *AdlAst) DeclType() (localctx IDeclTypeContext) {
	localctx = NewDeclTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, AdlAstRULE_declType)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(88)
		p.Match(AdlAstDeclType)
	}
	{
		p.SetState(89)
		p.Match(AdlAstDOWN)
	}
	p.SetState(94)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AdlAstStruct:
		{
			p.SetState(90)
			p.Struct_()
		}

	case AdlAstUnion:
		{
			p.SetState(91)
			p.Union_()
		}

	case AdlAstTypeDef:
		{
			p.SetState(92)
			p.TypeDef()
		}

	case AdlAstNewType:
		{
			p.SetState(93)
			p.NewType()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(96)
		p.Match(AdlAstUP)
	}

	return localctx
}

type IStruct_Context interface {
	antlr.ParserRuleContext
	// GetParser returns the parser.
	GetParser() antlr.Parser

	// start internal
	//  ruleGetterDecl
	//  ruleListGetterDecl
	AllField() []IFieldContext
	//  ruleListIndexedGetterDecl
	Field(i int) IFieldContext
	//  ruleContextDecls

	//  ruleContextListDecls
	GetFields() []IFieldContext

	// end internal
	//Gets for labeled elements
	//tokenDecls

	//tokenTypeDecls
	//tokenListDecls
	GetTypeParams() []antlr.Token
	//attributeDecls
	//tokenGetterDecl
	Struct() antlr.TerminalNode
	DOWN() antlr.TerminalNode
	UP() antlr.TerminalNode
	//tokenListGetterDecl
	AllID() []antlr.TerminalNode
	//tokenListIndexedGetterDecl
	ID(i int) antlr.TerminalNode

	// IsStruct_Context differentiates from other interfaces.
	//copyStruct,GetRuleContext and ToStringTree  from embedded

	//<if(dispatchMethods)>
	//<dispatchMethods; separator="\n\n">
	//<endif>

	//<if(extensionMembers)>
	//<extensionMembers; separator="\n\n">
	//<endif>
}

type Struct_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	//TokenDecl
	_ID antlr.Token
	//TokenListDecl
	typeParams []antlr.Token
	//RuleContextDecl
	_field IFieldContext
	//RuleContextListDecl
	fields []IFieldContext
}

func NewEmptyStruct_Context() *Struct_Context {
	var p = new(Struct_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AdlAstRULE_struct_
	return p
}

func (*Struct_Context) IsStruct_Context() {}

func NewStruct_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Struct_Context {
	var p = new(Struct_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AdlAstRULE_struct_

	return p
}

func (s *Struct_Context) GetParser() antlr.Parser { return s.parser }

//StructDecl tokenDecls
func (s *Struct_Context) Get_ID() antlr.Token  { return s._ID }
func (s *Struct_Context) Set_ID(v antlr.Token) { s._ID = v }

//StructDecl tokenTypeDecls

//StructDecl tokenListDecls
func (s *Struct_Context) GetTypeParams() []antlr.Token  { return s.typeParams }
func (s *Struct_Context) SetTypeParams(v []antlr.Token) { s.typeParams = v }

//StructDecl ruleContextDecls
func (s *Struct_Context) Get_field() IFieldContext  { return s._field }
func (s *Struct_Context) Set_field(v IFieldContext) { s._field = v }

//StructDecl ruleContextListDecls
func (s *Struct_Context) GetFields() []IFieldContext  { return s.fields }
func (s *Struct_Context) SetFields(v []IFieldContext) { s.fields = v }

//StructDecl attributeDecls

// Getters
func (s *Struct_Context) Struct() antlr.TerminalNode {
	return s.GetToken(AdlAstStruct, 0)
}

func (s *Struct_Context) DOWN() antlr.TerminalNode {
	return s.GetToken(AdlAstDOWN, 0)
}

func (s *Struct_Context) UP() antlr.TerminalNode {
	return s.GetToken(AdlAstUP, 0)
}

func (s *Struct_Context) AllID() []antlr.TerminalNode {
	return s.GetTokens(AdlAstID)
}

func (s *Struct_Context) ID(i int) antlr.TerminalNode {
	return s.GetToken(AdlAstID, i)
}

func (s *Struct_Context) AllField() []IFieldContext {
	//ContextRuleListGetterDecl
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*FieldContext)(nil)).Elem())
	var tst = make([]IFieldContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFieldContext)
		}
	}

	return tst
}

func (s *Struct_Context) Field(i int) IFieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*FieldContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFieldContext)
}

//provideCopyFrom
func (s *Struct_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Struct_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

//dispatchMethods
func (s *Struct_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Struct_EntryListener); ok {
		listenerT.EnterStruct_(s)
	}
}

func (s *Struct_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Struct_ExitListener); ok {
		listenerT.ExitStruct_(s)
	}
}

func (s *Struct_Context) VisitFunc(hdls antlr.ParserTreeVisitorHandlers, args ...interface{}) (result interface{}) {
	h := hdls.(*AdlAstHandlers)
	if h.EnterEveryRule != nil {
		h.EnterEveryRule(s)
	}
	if h.Struct_ != nil {
		h.Struct_(s, h, args...)
	} else {
		s.VisitChildrenFunc(h, args...)
	}
	if h.ExitEveryRule != nil {
		h.ExitEveryRule(s)
	}
	return
}

func (s *Struct_Context) Visit(delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	switch t := delegate.(type) {
	case Struct_ContextVisitor:
		return t.VisitStruct_(s, delegate, args...)
	default:
		return delegate.VisitChildren(s, delegate, args...)
	}
}

//extensionMembers

func (p *AdlAst) Struct_() (localctx IStruct_Context) {
	localctx = NewStruct_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, AdlAstRULE_struct_)
	var //TokenTypeDecl
	_la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(98)
		p.Match(AdlAstStruct)
	}
	p.SetState(113)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == AdlAstDOWN {
		{
			p.SetState(99)
			p.Match(AdlAstDOWN)
		}
		p.SetState(103)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == AdlAstID {
			{
				p.SetState(100)
				var _m = p.Match(AdlAstID)
				localctx.(*Struct_Context)._ID = _m

			}
			localctx.(*Struct_Context).typeParams = append(localctx.(*Struct_Context).typeParams, localctx.(*Struct_Context)._ID)

			p.SetState(105)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		p.SetState(109)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == AdlAstField {
			{
				p.SetState(106)

				var _x = p.Field()

				localctx.(*Struct_Context)._field = _x

			}
			localctx.(*Struct_Context).fields = append(localctx.(*Struct_Context).fields, localctx.(*Struct_Context)._field)

			p.SetState(111)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(112)
			p.Match(AdlAstUP)
		}

	}

	return localctx
}

type IUnion_Context interface {
	antlr.ParserRuleContext
	// GetParser returns the parser.
	GetParser() antlr.Parser

	// start internal
	//  ruleGetterDecl
	//  ruleListGetterDecl
	AllField() []IFieldContext
	//  ruleListIndexedGetterDecl
	Field(i int) IFieldContext
	//  ruleContextDecls

	//  ruleContextListDecls
	GetFields() []IFieldContext

	// end internal
	//Gets for labeled elements
	//tokenDecls

	//tokenTypeDecls
	//tokenListDecls
	GetTypeParams() []antlr.Token
	//attributeDecls
	//tokenGetterDecl
	Union() antlr.TerminalNode
	DOWN() antlr.TerminalNode
	UP() antlr.TerminalNode
	//tokenListGetterDecl
	AllID() []antlr.TerminalNode
	//tokenListIndexedGetterDecl
	ID(i int) antlr.TerminalNode

	// IsUnion_Context differentiates from other interfaces.
	//copyStruct,GetRuleContext and ToStringTree  from embedded

	//<if(dispatchMethods)>
	//<dispatchMethods; separator="\n\n">
	//<endif>

	//<if(extensionMembers)>
	//<extensionMembers; separator="\n\n">
	//<endif>
}

type Union_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	//TokenDecl
	_ID antlr.Token
	//TokenListDecl
	typeParams []antlr.Token
	//RuleContextDecl
	_field IFieldContext
	//RuleContextListDecl
	fields []IFieldContext
}

func NewEmptyUnion_Context() *Union_Context {
	var p = new(Union_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AdlAstRULE_union_
	return p
}

func (*Union_Context) IsUnion_Context() {}

func NewUnion_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Union_Context {
	var p = new(Union_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AdlAstRULE_union_

	return p
}

func (s *Union_Context) GetParser() antlr.Parser { return s.parser }

//StructDecl tokenDecls
func (s *Union_Context) Get_ID() antlr.Token  { return s._ID }
func (s *Union_Context) Set_ID(v antlr.Token) { s._ID = v }

//StructDecl tokenTypeDecls

//StructDecl tokenListDecls
func (s *Union_Context) GetTypeParams() []antlr.Token  { return s.typeParams }
func (s *Union_Context) SetTypeParams(v []antlr.Token) { s.typeParams = v }

//StructDecl ruleContextDecls
func (s *Union_Context) Get_field() IFieldContext  { return s._field }
func (s *Union_Context) Set_field(v IFieldContext) { s._field = v }

//StructDecl ruleContextListDecls
func (s *Union_Context) GetFields() []IFieldContext  { return s.fields }
func (s *Union_Context) SetFields(v []IFieldContext) { s.fields = v }

//StructDecl attributeDecls

// Getters
func (s *Union_Context) Union() antlr.TerminalNode {
	return s.GetToken(AdlAstUnion, 0)
}

func (s *Union_Context) DOWN() antlr.TerminalNode {
	return s.GetToken(AdlAstDOWN, 0)
}

func (s *Union_Context) UP() antlr.TerminalNode {
	return s.GetToken(AdlAstUP, 0)
}

func (s *Union_Context) AllID() []antlr.TerminalNode {
	return s.GetTokens(AdlAstID)
}

func (s *Union_Context) ID(i int) antlr.TerminalNode {
	return s.GetToken(AdlAstID, i)
}

func (s *Union_Context) AllField() []IFieldContext {
	//ContextRuleListGetterDecl
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*FieldContext)(nil)).Elem())
	var tst = make([]IFieldContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFieldContext)
		}
	}

	return tst
}

func (s *Union_Context) Field(i int) IFieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*FieldContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFieldContext)
}

//provideCopyFrom
func (s *Union_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Union_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

//dispatchMethods
func (s *Union_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Union_EntryListener); ok {
		listenerT.EnterUnion_(s)
	}
}

func (s *Union_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Union_ExitListener); ok {
		listenerT.ExitUnion_(s)
	}
}

func (s *Union_Context) VisitFunc(hdls antlr.ParserTreeVisitorHandlers, args ...interface{}) (result interface{}) {
	h := hdls.(*AdlAstHandlers)
	if h.EnterEveryRule != nil {
		h.EnterEveryRule(s)
	}
	if h.Union_ != nil {
		h.Union_(s, h, args...)
	} else {
		s.VisitChildrenFunc(h, args...)
	}
	if h.ExitEveryRule != nil {
		h.ExitEveryRule(s)
	}
	return
}

func (s *Union_Context) Visit(delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	switch t := delegate.(type) {
	case Union_ContextVisitor:
		return t.VisitUnion_(s, delegate, args...)
	default:
		return delegate.VisitChildren(s, delegate, args...)
	}
}

//extensionMembers

func (p *AdlAst) Union_() (localctx IUnion_Context) {
	localctx = NewUnion_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, AdlAstRULE_union_)
	var //TokenTypeDecl
	_la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(115)
		p.Match(AdlAstUnion)
	}
	p.SetState(130)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == AdlAstDOWN {
		{
			p.SetState(116)
			p.Match(AdlAstDOWN)
		}
		p.SetState(120)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == AdlAstID {
			{
				p.SetState(117)
				var _m = p.Match(AdlAstID)
				localctx.(*Union_Context)._ID = _m

			}
			localctx.(*Union_Context).typeParams = append(localctx.(*Union_Context).typeParams, localctx.(*Union_Context)._ID)

			p.SetState(122)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		p.SetState(126)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == AdlAstField {
			{
				p.SetState(123)

				var _x = p.Field()

				localctx.(*Union_Context)._field = _x

			}
			localctx.(*Union_Context).fields = append(localctx.(*Union_Context).fields, localctx.(*Union_Context)._field)

			p.SetState(128)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(129)
			p.Match(AdlAstUP)
		}

	}

	return localctx
}

type ITypeDefContext interface {
	antlr.ParserRuleContext
	// GetParser returns the parser.
	GetParser() antlr.Parser

	// start internal
	//  ruleGetterDecl
	TypeExpr() ITypeExprContext
	//  ruleListGetterDecl
	//  ruleListIndexedGetterDecl
	//  ruleContextDecls
	//  ruleContextListDecls
	// end internal
	//Gets for labeled elements
	//tokenDecls

	//tokenTypeDecls
	//tokenListDecls
	GetTypeParams() []antlr.Token
	//attributeDecls
	//tokenGetterDecl
	TypeDef() antlr.TerminalNode
	DOWN() antlr.TerminalNode
	UP() antlr.TerminalNode
	//tokenListGetterDecl
	AllID() []antlr.TerminalNode
	//tokenListIndexedGetterDecl
	ID(i int) antlr.TerminalNode

	// IsTypeDefContext differentiates from other interfaces.
	//copyStruct,GetRuleContext and ToStringTree  from embedded

	//<if(dispatchMethods)>
	//<dispatchMethods; separator="\n\n">
	//<endif>

	//<if(extensionMembers)>
	//<extensionMembers; separator="\n\n">
	//<endif>
}

type TypeDefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	//TokenDecl
	_ID antlr.Token
	//TokenListDecl
	typeParams []antlr.Token
}

func NewEmptyTypeDefContext() *TypeDefContext {
	var p = new(TypeDefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AdlAstRULE_typeDef
	return p
}

func (*TypeDefContext) IsTypeDefContext() {}

func NewTypeDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeDefContext {
	var p = new(TypeDefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AdlAstRULE_typeDef

	return p
}

func (s *TypeDefContext) GetParser() antlr.Parser { return s.parser }

//StructDecl tokenDecls
func (s *TypeDefContext) Get_ID() antlr.Token  { return s._ID }
func (s *TypeDefContext) Set_ID(v antlr.Token) { s._ID = v }

//StructDecl tokenTypeDecls

//StructDecl tokenListDecls
func (s *TypeDefContext) GetTypeParams() []antlr.Token  { return s.typeParams }
func (s *TypeDefContext) SetTypeParams(v []antlr.Token) { s.typeParams = v }

//StructDecl ruleContextDecls

//StructDecl ruleContextListDecls

//StructDecl attributeDecls

// Getters
func (s *TypeDefContext) TypeDef() antlr.TerminalNode {
	return s.GetToken(AdlAstTypeDef, 0)
}

func (s *TypeDefContext) DOWN() antlr.TerminalNode {
	return s.GetToken(AdlAstDOWN, 0)
}

func (s *TypeDefContext) TypeExpr() ITypeExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*TypeExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeExprContext)
}

func (s *TypeDefContext) UP() antlr.TerminalNode {
	return s.GetToken(AdlAstUP, 0)
}

func (s *TypeDefContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(AdlAstID)
}

func (s *TypeDefContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(AdlAstID, i)
}

//provideCopyFrom
func (s *TypeDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

//dispatchMethods
func (s *TypeDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TypeDefEntryListener); ok {
		listenerT.EnterTypeDef(s)
	}
}

func (s *TypeDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TypeDefExitListener); ok {
		listenerT.ExitTypeDef(s)
	}
}

func (s *TypeDefContext) VisitFunc(hdls antlr.ParserTreeVisitorHandlers, args ...interface{}) (result interface{}) {
	h := hdls.(*AdlAstHandlers)
	if h.EnterEveryRule != nil {
		h.EnterEveryRule(s)
	}
	if h.TypeDef != nil {
		h.TypeDef(s, h, args...)
	} else {
		s.VisitChildrenFunc(h, args...)
	}
	if h.ExitEveryRule != nil {
		h.ExitEveryRule(s)
	}
	return
}

func (s *TypeDefContext) Visit(delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	switch t := delegate.(type) {
	case TypeDefContextVisitor:
		return t.VisitTypeDef(s, delegate, args...)
	default:
		return delegate.VisitChildren(s, delegate, args...)
	}
}

//extensionMembers

func (p *AdlAst) TypeDef() (localctx ITypeDefContext) {
	localctx = NewTypeDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, AdlAstRULE_typeDef)
	var //TokenTypeDecl
	_la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(132)
		p.Match(AdlAstTypeDef)
	}
	{
		p.SetState(133)
		p.Match(AdlAstDOWN)
	}
	p.SetState(137)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == AdlAstID {
		{
			p.SetState(134)
			var _m = p.Match(AdlAstID)
			localctx.(*TypeDefContext)._ID = _m

		}
		localctx.(*TypeDefContext).typeParams = append(localctx.(*TypeDefContext).typeParams, localctx.(*TypeDefContext)._ID)

		p.SetState(139)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(140)
		p.TypeExpr()
	}
	{
		p.SetState(141)
		p.Match(AdlAstUP)
	}

	return localctx
}

type INewTypeContext interface {
	antlr.ParserRuleContext
	// GetParser returns the parser.
	GetParser() antlr.Parser

	// start internal
	//  ruleGetterDecl
	TypeExpr() ITypeExprContext
	Json() IJsonContext
	//  ruleListGetterDecl
	//  ruleListIndexedGetterDecl
	//  ruleContextDecls
	GetDefault_() IJsonContext

	//  ruleContextListDecls
	// end internal
	//Gets for labeled elements
	//tokenDecls

	//tokenTypeDecls
	//tokenListDecls
	GetTypeParams() []antlr.Token
	//attributeDecls
	//tokenGetterDecl
	NewType() antlr.TerminalNode
	DOWN() antlr.TerminalNode
	UP() antlr.TerminalNode
	//tokenListGetterDecl
	AllID() []antlr.TerminalNode
	//tokenListIndexedGetterDecl
	ID(i int) antlr.TerminalNode

	// IsNewTypeContext differentiates from other interfaces.
	//copyStruct,GetRuleContext and ToStringTree  from embedded

	//<if(dispatchMethods)>
	//<dispatchMethods; separator="\n\n">
	//<endif>

	//<if(extensionMembers)>
	//<extensionMembers; separator="\n\n">
	//<endif>
}

type NewTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	//TokenDecl
	_ID antlr.Token
	//TokenListDecl
	typeParams []antlr.Token
	//RuleContextDecl
	default_ IJsonContext
}

func NewEmptyNewTypeContext() *NewTypeContext {
	var p = new(NewTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AdlAstRULE_newType
	return p
}

func (*NewTypeContext) IsNewTypeContext() {}

func NewNewTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NewTypeContext {
	var p = new(NewTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AdlAstRULE_newType

	return p
}

func (s *NewTypeContext) GetParser() antlr.Parser { return s.parser }

//StructDecl tokenDecls
func (s *NewTypeContext) Get_ID() antlr.Token  { return s._ID }
func (s *NewTypeContext) Set_ID(v antlr.Token) { s._ID = v }

//StructDecl tokenTypeDecls

//StructDecl tokenListDecls
func (s *NewTypeContext) GetTypeParams() []antlr.Token  { return s.typeParams }
func (s *NewTypeContext) SetTypeParams(v []antlr.Token) { s.typeParams = v }

//StructDecl ruleContextDecls
func (s *NewTypeContext) GetDefault_() IJsonContext  { return s.default_ }
func (s *NewTypeContext) SetDefault_(v IJsonContext) { s.default_ = v }

//StructDecl ruleContextListDecls

//StructDecl attributeDecls

// Getters
func (s *NewTypeContext) NewType() antlr.TerminalNode {
	return s.GetToken(AdlAstNewType, 0)
}

func (s *NewTypeContext) DOWN() antlr.TerminalNode {
	return s.GetToken(AdlAstDOWN, 0)
}

func (s *NewTypeContext) TypeExpr() ITypeExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*TypeExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeExprContext)
}

func (s *NewTypeContext) UP() antlr.TerminalNode {
	return s.GetToken(AdlAstUP, 0)
}

func (s *NewTypeContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(AdlAstID)
}

func (s *NewTypeContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(AdlAstID, i)
}

func (s *NewTypeContext) Json() IJsonContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*JsonContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IJsonContext)
}

//provideCopyFrom
func (s *NewTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NewTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

//dispatchMethods
func (s *NewTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NewTypeEntryListener); ok {
		listenerT.EnterNewType(s)
	}
}

func (s *NewTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NewTypeExitListener); ok {
		listenerT.ExitNewType(s)
	}
}

func (s *NewTypeContext) VisitFunc(hdls antlr.ParserTreeVisitorHandlers, args ...interface{}) (result interface{}) {
	h := hdls.(*AdlAstHandlers)
	if h.EnterEveryRule != nil {
		h.EnterEveryRule(s)
	}
	if h.NewType != nil {
		h.NewType(s, h, args...)
	} else {
		s.VisitChildrenFunc(h, args...)
	}
	if h.ExitEveryRule != nil {
		h.ExitEveryRule(s)
	}
	return
}

func (s *NewTypeContext) Visit(delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	switch t := delegate.(type) {
	case NewTypeContextVisitor:
		return t.VisitNewType(s, delegate, args...)
	default:
		return delegate.VisitChildren(s, delegate, args...)
	}
}

//extensionMembers

func (p *AdlAst) NewType() (localctx INewTypeContext) {
	localctx = NewNewTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, AdlAstRULE_newType)
	var //TokenTypeDecl
	_la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(143)
		p.Match(AdlAstNewType)
	}
	{
		p.SetState(144)
		p.Match(AdlAstDOWN)
	}
	p.SetState(148)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == AdlAstID {
		{
			p.SetState(145)
			var _m = p.Match(AdlAstID)
			localctx.(*NewTypeContext)._ID = _m

		}
		localctx.(*NewTypeContext).typeParams = append(localctx.(*NewTypeContext).typeParams, localctx.(*NewTypeContext)._ID)

		p.SetState(150)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(151)
		p.TypeExpr()
	}
	p.SetState(153)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la-27)&-(0x1f+1)) == 0 && ((1<<uint((_la-27)))&((1<<(AdlAstERROR-27))|(1<<(AdlAstJsonStr-27))|(1<<(AdlAstJsonBool-27))|(1<<(AdlAstJsonNull-27))|(1<<(AdlAstJsonInt-27))|(1<<(AdlAstJsonFloat-27))|(1<<(AdlAstJsonArray-27))|(1<<(AdlAstJsonObj-27)))) != 0 {
		{
			p.SetState(152)

			var _x = p.Json()

			localctx.(*NewTypeContext).default_ = _x

		}

	}
	{
		p.SetState(155)
		p.Match(AdlAstUP)
	}

	return localctx
}

type IFieldContext interface {
	antlr.ParserRuleContext
	// GetParser returns the parser.
	GetParser() antlr.Parser

	// start internal
	//  ruleGetterDecl
	TypeExpr() ITypeExprContext
	//  ruleListGetterDecl
	AllScopedName() []IScopedNameContext
	AllJson() []IJsonContext
	//  ruleListIndexedGetterDecl
	ScopedName(i int) IScopedNameContext
	Json(i int) IJsonContext
	//  ruleContextDecls
	GetDefault_() IJsonContext

	//  ruleContextListDecls
	// end internal
	//Gets for labeled elements
	//tokenDecls
	GetName() antlr.Token
	GetSerializedName() antlr.Token

	//tokenTypeDecls
	//tokenListDecls
	//attributeDecls
	//tokenGetterDecl
	Field() antlr.TerminalNode
	DOWN() antlr.TerminalNode
	UP() antlr.TerminalNode
	//tokenListGetterDecl
	AllID() []antlr.TerminalNode
	//tokenListIndexedGetterDecl
	ID(i int) antlr.TerminalNode

	// IsFieldContext differentiates from other interfaces.
	//copyStruct,GetRuleContext and ToStringTree  from embedded

	//<if(dispatchMethods)>
	//<dispatchMethods; separator="\n\n">
	//<endif>

	//<if(extensionMembers)>
	//<extensionMembers; separator="\n\n">
	//<endif>
}

type FieldContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	//TokenDecl
	name antlr.Token
	//TokenDecl
	serializedName antlr.Token
	//RuleContextDecl
	default_ IJsonContext
}

func NewEmptyFieldContext() *FieldContext {
	var p = new(FieldContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AdlAstRULE_field
	return p
}

func (*FieldContext) IsFieldContext() {}

func NewFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldContext {
	var p = new(FieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AdlAstRULE_field

	return p
}

func (s *FieldContext) GetParser() antlr.Parser { return s.parser }

//StructDecl tokenDecls
func (s *FieldContext) GetName() antlr.Token  { return s.name }
func (s *FieldContext) SetName(v antlr.Token) { s.name = v }

func (s *FieldContext) GetSerializedName() antlr.Token  { return s.serializedName }
func (s *FieldContext) SetSerializedName(v antlr.Token) { s.serializedName = v }

//StructDecl tokenTypeDecls

//StructDecl tokenListDecls

//StructDecl ruleContextDecls
func (s *FieldContext) GetDefault_() IJsonContext  { return s.default_ }
func (s *FieldContext) SetDefault_(v IJsonContext) { s.default_ = v }

//StructDecl ruleContextListDecls

//StructDecl attributeDecls

// Getters
func (s *FieldContext) Field() antlr.TerminalNode {
	return s.GetToken(AdlAstField, 0)
}

func (s *FieldContext) DOWN() antlr.TerminalNode {
	return s.GetToken(AdlAstDOWN, 0)
}

func (s *FieldContext) TypeExpr() ITypeExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*TypeExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeExprContext)
}

func (s *FieldContext) UP() antlr.TerminalNode {
	return s.GetToken(AdlAstUP, 0)
}

func (s *FieldContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(AdlAstID)
}

func (s *FieldContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(AdlAstID, i)
}

func (s *FieldContext) AllScopedName() []IScopedNameContext {
	//ContextRuleListGetterDecl
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ScopedNameContext)(nil)).Elem())
	var tst = make([]IScopedNameContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IScopedNameContext)
		}
	}

	return tst
}

func (s *FieldContext) ScopedName(i int) IScopedNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ScopedNameContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IScopedNameContext)
}

func (s *FieldContext) AllJson() []IJsonContext {
	//ContextRuleListGetterDecl
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*JsonContext)(nil)).Elem())
	var tst = make([]IJsonContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IJsonContext)
		}
	}

	return tst
}

func (s *FieldContext) Json(i int) IJsonContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*JsonContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IJsonContext)
}

//provideCopyFrom
func (s *FieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

//dispatchMethods
func (s *FieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FieldEntryListener); ok {
		listenerT.EnterField(s)
	}
}

func (s *FieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FieldExitListener); ok {
		listenerT.ExitField(s)
	}
}

func (s *FieldContext) VisitFunc(hdls antlr.ParserTreeVisitorHandlers, args ...interface{}) (result interface{}) {
	h := hdls.(*AdlAstHandlers)
	if h.EnterEveryRule != nil {
		h.EnterEveryRule(s)
	}
	if h.Field != nil {
		h.Field(s, h, args...)
	} else {
		s.VisitChildrenFunc(h, args...)
	}
	if h.ExitEveryRule != nil {
		h.ExitEveryRule(s)
	}
	return
}

func (s *FieldContext) Visit(delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	switch t := delegate.(type) {
	case FieldContextVisitor:
		return t.VisitField(s, delegate, args...)
	default:
		return delegate.VisitChildren(s, delegate, args...)
	}
}

//extensionMembers

func (p *AdlAst) Field() (localctx IFieldContext) {
	localctx = NewFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, AdlAstRULE_field)
	var //TokenTypeDecl
	_la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(157)
		p.Match(AdlAstField)
	}
	{
		p.SetState(158)
		p.Match(AdlAstDOWN)
	}
	{
		p.SetState(159)
		var _m = p.Match(AdlAstID)
		localctx.(*FieldContext).name = _m

	}
	{
		p.SetState(160)
		var _m = p.Match(AdlAstID)
		localctx.(*FieldContext).serializedName = _m

	}
	{
		p.SetState(161)
		p.TypeExpr()
	}
	p.SetState(163)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la-27)&-(0x1f+1)) == 0 && ((1<<uint((_la-27)))&((1<<(AdlAstERROR-27))|(1<<(AdlAstJsonStr-27))|(1<<(AdlAstJsonBool-27))|(1<<(AdlAstJsonNull-27))|(1<<(AdlAstJsonInt-27))|(1<<(AdlAstJsonFloat-27))|(1<<(AdlAstJsonArray-27))|(1<<(AdlAstJsonObj-27)))) != 0 {
		{
			p.SetState(162)

			var _x = p.Json()

			localctx.(*FieldContext).default_ = _x

		}

	}
	p.SetState(170)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == AdlAstScopedName {
		{
			p.SetState(165)
			p.ScopedName()
		}
		{
			p.SetState(166)
			p.Json()
		}

		p.SetState(172)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(173)
		p.Match(AdlAstUP)
	}

	return localctx
}

type ITypeExprContext interface {
	antlr.ParserRuleContext
	// GetParser returns the parser.
	GetParser() antlr.Parser

	// start internal
	//  ruleGetterDecl
	TypeRef() ITypeRefContext
	//  ruleListGetterDecl
	AllTypeExpr() []ITypeExprContext
	//  ruleListIndexedGetterDecl
	TypeExpr(i int) ITypeExprContext
	//  ruleContextDecls

	//  ruleContextListDecls
	GetParameters() []ITypeExprContext

	// end internal
	//Gets for labeled elements
	//tokenDecls
	//tokenTypeDecls
	//tokenListDecls
	//attributeDecls
	//tokenGetterDecl
	TypeExpr() antlr.TerminalNode
	DOWN() antlr.TerminalNode
	UP() antlr.TerminalNode
	//tokenListGetterDecl
	//tokenListIndexedGetterDecl

	// IsTypeExprContext differentiates from other interfaces.
	//copyStruct,GetRuleContext and ToStringTree  from embedded

	//<if(dispatchMethods)>
	//<dispatchMethods; separator="\n\n">
	//<endif>

	//<if(extensionMembers)>
	//<extensionMembers; separator="\n\n">
	//<endif>
}

type TypeExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	//RuleContextDecl
	_typeExpr ITypeExprContext
	//RuleContextListDecl
	parameters []ITypeExprContext
}

func NewEmptyTypeExprContext() *TypeExprContext {
	var p = new(TypeExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AdlAstRULE_typeExpr
	return p
}

func (*TypeExprContext) IsTypeExprContext() {}

func NewTypeExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeExprContext {
	var p = new(TypeExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AdlAstRULE_typeExpr

	return p
}

func (s *TypeExprContext) GetParser() antlr.Parser { return s.parser }

//StructDecl tokenDecls

//StructDecl tokenTypeDecls

//StructDecl tokenListDecls

//StructDecl ruleContextDecls
func (s *TypeExprContext) Get_typeExpr() ITypeExprContext  { return s._typeExpr }
func (s *TypeExprContext) Set_typeExpr(v ITypeExprContext) { s._typeExpr = v }

//StructDecl ruleContextListDecls
func (s *TypeExprContext) GetParameters() []ITypeExprContext  { return s.parameters }
func (s *TypeExprContext) SetParameters(v []ITypeExprContext) { s.parameters = v }

//StructDecl attributeDecls

// Getters
func (s *TypeExprContext) TypeExpr() antlr.TerminalNode {
	return s.GetToken(AdlAstTypeExpr, 0)
}

func (s *TypeExprContext) DOWN() antlr.TerminalNode {
	return s.GetToken(AdlAstDOWN, 0)
}

func (s *TypeExprContext) TypeRef() ITypeRefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*TypeRefContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeRefContext)
}

func (s *TypeExprContext) UP() antlr.TerminalNode {
	return s.GetToken(AdlAstUP, 0)
}

func (s *TypeExprContext) AllTypeExpr() []ITypeExprContext {
	//ContextRuleListGetterDecl
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*TypeExprContext)(nil)).Elem())
	var tst = make([]ITypeExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITypeExprContext)
		}
	}

	return tst
}

func (s *TypeExprContext) TypeExpr(i int) ITypeExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*TypeExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITypeExprContext)
}

//provideCopyFrom
func (s *TypeExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

//dispatchMethods
func (s *TypeExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TypeExprEntryListener); ok {
		listenerT.EnterTypeExpr(s)
	}
}

func (s *TypeExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TypeExprExitListener); ok {
		listenerT.ExitTypeExpr(s)
	}
}

func (s *TypeExprContext) VisitFunc(hdls antlr.ParserTreeVisitorHandlers, args ...interface{}) (result interface{}) {
	h := hdls.(*AdlAstHandlers)
	if h.EnterEveryRule != nil {
		h.EnterEveryRule(s)
	}
	if h.TypeExpr != nil {
		h.TypeExpr(s, h, args...)
	} else {
		s.VisitChildrenFunc(h, args...)
	}
	if h.ExitEveryRule != nil {
		h.ExitEveryRule(s)
	}
	return
}

func (s *TypeExprContext) Visit(delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	switch t := delegate.(type) {
	case TypeExprContextVisitor:
		return t.VisitTypeExpr(s, delegate, args...)
	default:
		return delegate.VisitChildren(s, delegate, args...)
	}
}

//extensionMembers

func (p *AdlAst) TypeExpr() (localctx ITypeExprContext) {
	localctx = NewTypeExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, AdlAstRULE_typeExpr)
	var //TokenTypeDecl
	_la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(175)
		p.Match(AdlAstTypeExpr)
	}
	{
		p.SetState(176)
		p.Match(AdlAstDOWN)
	}
	{
		p.SetState(177)
		p.TypeRef()
	}
	p.SetState(181)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == AdlAstTypeExpr {
		{
			p.SetState(178)

			var _x = p.TypeExpr()

			localctx.(*TypeExprContext)._typeExpr = _x

		}
		localctx.(*TypeExprContext).parameters = append(localctx.(*TypeExprContext).parameters, localctx.(*TypeExprContext)._typeExpr)

		p.SetState(183)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(184)
		p.Match(AdlAstUP)
	}

	return localctx
}

type ITypeRefContext interface {
	antlr.ParserRuleContext
	// GetParser returns the parser.
	GetParser() antlr.Parser

	// start internal
	//  ruleGetterDecl
	ScopedName() IScopedNameContext
	//  ruleListGetterDecl
	//  ruleListIndexedGetterDecl
	//  ruleContextDecls
	//  ruleContextListDecls
	// end internal
	//Gets for labeled elements
	//tokenDecls
	//tokenTypeDecls
	//tokenListDecls
	//attributeDecls
	//tokenGetterDecl
	Primitive() antlr.TerminalNode
	ID() antlr.TerminalNode
	TypeParam() antlr.TerminalNode
	//tokenListGetterDecl
	//tokenListIndexedGetterDecl

	// IsTypeRefContext differentiates from other interfaces.
	//copyStruct,GetRuleContext and ToStringTree  from embedded

	//<if(dispatchMethods)>
	//<dispatchMethods; separator="\n\n">
	//<endif>

	//<if(extensionMembers)>
	//<extensionMembers; separator="\n\n">
	//<endif>
}

type TypeRefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeRefContext() *TypeRefContext {
	var p = new(TypeRefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AdlAstRULE_typeRef
	return p
}

func (*TypeRefContext) IsTypeRefContext() {}

func NewTypeRefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeRefContext {
	var p = new(TypeRefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AdlAstRULE_typeRef

	return p
}

func (s *TypeRefContext) GetParser() antlr.Parser { return s.parser }

//StructDecl tokenDecls

//StructDecl tokenTypeDecls

//StructDecl tokenListDecls

//StructDecl ruleContextDecls

//StructDecl ruleContextListDecls

//StructDecl attributeDecls

// Getters
func (s *TypeRefContext) Primitive() antlr.TerminalNode {
	return s.GetToken(AdlAstPrimitive, 0)
}

func (s *TypeRefContext) ID() antlr.TerminalNode {
	return s.GetToken(AdlAstID, 0)
}

func (s *TypeRefContext) TypeParam() antlr.TerminalNode {
	return s.GetToken(AdlAstTypeParam, 0)
}

func (s *TypeRefContext) ScopedName() IScopedNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ScopedNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IScopedNameContext)
}

//provideCopyFrom
func (s *TypeRefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeRefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

//dispatchMethods
func (s *TypeRefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TypeRefEntryListener); ok {
		listenerT.EnterTypeRef(s)
	}
}

func (s *TypeRefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TypeRefExitListener); ok {
		listenerT.ExitTypeRef(s)
	}
}

func (s *TypeRefContext) VisitFunc(hdls antlr.ParserTreeVisitorHandlers, args ...interface{}) (result interface{}) {
	h := hdls.(*AdlAstHandlers)
	if h.EnterEveryRule != nil {
		h.EnterEveryRule(s)
	}
	if h.TypeRef != nil {
		h.TypeRef(s, h, args...)
	} else {
		s.VisitChildrenFunc(h, args...)
	}
	if h.ExitEveryRule != nil {
		h.ExitEveryRule(s)
	}
	return
}

func (s *TypeRefContext) Visit(delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	switch t := delegate.(type) {
	case TypeRefContextVisitor:
		return t.VisitTypeRef(s, delegate, args...)
	default:
		return delegate.VisitChildren(s, delegate, args...)
	}
}

//extensionMembers

func (p *AdlAst) TypeRef() (localctx ITypeRefContext) {
	localctx = NewTypeRefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, AdlAstRULE_typeRef)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(191)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AdlAstPrimitive:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(186)
			p.Match(AdlAstPrimitive)
		}
		{
			p.SetState(187)
			p.Match(AdlAstID)
		}

	case AdlAstTypeParam:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(188)
			p.Match(AdlAstTypeParam)
		}
		{
			p.SetState(189)
			p.Match(AdlAstID)
		}

	case AdlAstScopedName:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(190)
			p.ScopedName()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

type IJsonContext interface {
	antlr.ParserRuleContext
	// GetParser returns the parser.
	GetParser() antlr.Parser

	// start internal
	//  ruleGetterDecl
	//  ruleListGetterDecl
	//  ruleListIndexedGetterDecl
	//  ruleContextDecls
	//  ruleContextListDecls
	// end internal
	//Gets for labeled elements
	//tokenDecls
	//tokenTypeDecls
	//tokenListDecls
	//attributeDecls
	//tokenGetterDecl
	//tokenListGetterDecl
	//tokenListIndexedGetterDecl

	// IsJsonContext differentiates from other interfaces.
	//copyStruct,GetRuleContext and ToStringTree  from embedded

	//<if(dispatchMethods)>
	//<dispatchMethods; separator="\n\n">
	//<endif>

	//<if(extensionMembers)>
	//<extensionMembers; separator="\n\n">
	//<endif>
}

type JsonContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJsonContext() *JsonContext {
	var p = new(JsonContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AdlAstRULE_json
	return p
}

func (*JsonContext) IsJsonContext() {}

func NewJsonContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *JsonContext {
	var p = new(JsonContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AdlAstRULE_json

	return p
}

func (s *JsonContext) GetParser() antlr.Parser { return s.parser }

//StructDecl tokenDecls

//StructDecl tokenTypeDecls

//StructDecl tokenListDecls

//StructDecl ruleContextDecls

//StructDecl ruleContextListDecls

//StructDecl attributeDecls

// Getters

//provideCopyFrom
func (s *JsonContext) CopyFrom(ctx *JsonContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *JsonContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JsonContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

//dispatchMethods
//extensionMembers

//Begin AltLabelStructDecl

type JsonStrContext struct {
	*JsonContext
	//TokenDecl
	tok antlr.Token
}

func NewJsonStrContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *JsonStrContext {
	var p = new(JsonStrContext)

	p.JsonContext = NewEmptyJsonContext()
	p.parser = parser
	p.CopyFrom(ctx.(*JsonContext))

	return p
}

type IJsonStrContext interface {
	//Current rule
	IJsonContext
	// start internal
	//Gets for raw elements
	//  ruleGetterDecl
	//  ruleListGetterDecl
	//  ruleListIndexedGetterDecl

	//  tokenGetterDecl
	JsonStr() antlr.TerminalNode
	//  tokenListGetterDecl
	//  tokenListIndexedGetterDecl
	// end internal
	//
	//Gets for labeled elements
	//  tokenDecls
	GetTok() antlr.Token

	//  tokenTypeDecls
	//  tokenListDecls
	//  ruleContextDecls
	//  ruleContextListDecls
	//  attributeDecls

	// TODO dispatchMethods (needed?)
}

func (*JsonStrContext) IsJsonStrContext() {}

//AltLabelStructDecl tokenDecls
func (s *JsonStrContext) GetTok() antlr.Token  { return s.tok }
func (s *JsonStrContext) SetTok(v antlr.Token) { s.tok = v }

//AltLabelStructDecl tokenTypeDecls

//AltLabelStructDecl tokenListDecls

//AltLabelStructDecl ruleContextDecls

//AltLabelStructDecl ruleContextListDecls

//AltLabelStructDecl attributeDecls

//getRuleContext
func (s *JsonStrContext) GetRuleContext() antlr.RuleContext {
	return s
}

//getters
func (s *JsonStrContext) JsonStr() antlr.TerminalNode {
	return s.GetToken(AdlAstJsonStr, 0)
}

func (s *JsonStrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonStrEntryListener); ok {
		listenerT.EnterJsonStr(s)
	}
}

func (s *JsonStrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonStrExitListener); ok {
		listenerT.ExitJsonStr(s)
	}
}

func (s *JsonStrContext) VisitFunc(hdls antlr.ParserTreeVisitorHandlers, args ...interface{}) (result interface{}) {
	h := hdls.(*AdlAstHandlers)
	if h.EnterEveryRule != nil {
		h.EnterEveryRule(s)
	}
	if h.JsonStr != nil {
		h.JsonStr(s, h, args...)
	} else {
		s.VisitChildrenFunc(h, args...)
	}
	if h.ExitEveryRule != nil {
		h.ExitEveryRule(s)
	}
	return
}

func (s *JsonStrContext) Visit(delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	switch t := delegate.(type) {
	case JsonStrContextVisitor:
		return t.VisitJsonStr(s, delegate, args...)
	default:
		return delegate.VisitChildren(s, delegate, args...)
	}
}

//END AltLabelStructDecl

//Begin AltLabelStructDecl

type JsonArrayContext struct {
	*JsonContext
	//TokenDecl
	tok antlr.Token
}

func NewJsonArrayContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *JsonArrayContext {
	var p = new(JsonArrayContext)

	p.JsonContext = NewEmptyJsonContext()
	p.parser = parser
	p.CopyFrom(ctx.(*JsonContext))

	return p
}

type IJsonArrayContext interface {
	//Current rule
	IJsonContext
	// start internal
	//Gets for raw elements
	//  ruleGetterDecl
	//  ruleListGetterDecl
	AllJson() []IJsonContext
	//  ruleListIndexedGetterDecl
	Json(i int) IJsonContext

	//  tokenGetterDecl
	JsonArray() antlr.TerminalNode
	DOWN() antlr.TerminalNode
	UP() antlr.TerminalNode
	//  tokenListGetterDecl
	//  tokenListIndexedGetterDecl
	// end internal
	//
	//Gets for labeled elements
	//  tokenDecls
	GetTok() antlr.Token

	//  tokenTypeDecls
	//  tokenListDecls
	//  ruleContextDecls
	//  ruleContextListDecls
	//  attributeDecls

	// TODO dispatchMethods (needed?)
}

func (*JsonArrayContext) IsJsonArrayContext() {}

//AltLabelStructDecl tokenDecls
func (s *JsonArrayContext) GetTok() antlr.Token  { return s.tok }
func (s *JsonArrayContext) SetTok(v antlr.Token) { s.tok = v }

//AltLabelStructDecl tokenTypeDecls

//AltLabelStructDecl tokenListDecls

//AltLabelStructDecl ruleContextDecls

//AltLabelStructDecl ruleContextListDecls

//AltLabelStructDecl attributeDecls

//getRuleContext
func (s *JsonArrayContext) GetRuleContext() antlr.RuleContext {
	return s
}

//getters
func (s *JsonArrayContext) JsonArray() antlr.TerminalNode {
	return s.GetToken(AdlAstJsonArray, 0)
}

func (s *JsonArrayContext) DOWN() antlr.TerminalNode {
	return s.GetToken(AdlAstDOWN, 0)
}

func (s *JsonArrayContext) UP() antlr.TerminalNode {
	return s.GetToken(AdlAstUP, 0)
}

func (s *JsonArrayContext) AllJson() []IJsonContext {
	//ContextRuleListGetterDecl
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*JsonContext)(nil)).Elem())
	var tst = make([]IJsonContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IJsonContext)
		}
	}

	return tst
}

func (s *JsonArrayContext) Json(i int) IJsonContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*JsonContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IJsonContext)
}

func (s *JsonArrayContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonArrayEntryListener); ok {
		listenerT.EnterJsonArray(s)
	}
}

func (s *JsonArrayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonArrayExitListener); ok {
		listenerT.ExitJsonArray(s)
	}
}

func (s *JsonArrayContext) VisitFunc(hdls antlr.ParserTreeVisitorHandlers, args ...interface{}) (result interface{}) {
	h := hdls.(*AdlAstHandlers)
	if h.EnterEveryRule != nil {
		h.EnterEveryRule(s)
	}
	if h.JsonArray != nil {
		h.JsonArray(s, h, args...)
	} else {
		s.VisitChildrenFunc(h, args...)
	}
	if h.ExitEveryRule != nil {
		h.ExitEveryRule(s)
	}
	return
}

func (s *JsonArrayContext) Visit(delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	switch t := delegate.(type) {
	case JsonArrayContextVisitor:
		return t.VisitJsonArray(s, delegate, args...)
	default:
		return delegate.VisitChildren(s, delegate, args...)
	}
}

//END AltLabelStructDecl

//Begin AltLabelStructDecl

type JsonFloatContext struct {
	*JsonContext
	//TokenDecl
	tok antlr.Token
}

func NewJsonFloatContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *JsonFloatContext {
	var p = new(JsonFloatContext)

	p.JsonContext = NewEmptyJsonContext()
	p.parser = parser
	p.CopyFrom(ctx.(*JsonContext))

	return p
}

type IJsonFloatContext interface {
	//Current rule
	IJsonContext
	// start internal
	//Gets for raw elements
	//  ruleGetterDecl
	//  ruleListGetterDecl
	//  ruleListIndexedGetterDecl

	//  tokenGetterDecl
	JsonFloat() antlr.TerminalNode
	//  tokenListGetterDecl
	//  tokenListIndexedGetterDecl
	// end internal
	//
	//Gets for labeled elements
	//  tokenDecls
	GetTok() antlr.Token

	//  tokenTypeDecls
	//  tokenListDecls
	//  ruleContextDecls
	//  ruleContextListDecls
	//  attributeDecls

	// TODO dispatchMethods (needed?)
}

func (*JsonFloatContext) IsJsonFloatContext() {}

//AltLabelStructDecl tokenDecls
func (s *JsonFloatContext) GetTok() antlr.Token  { return s.tok }
func (s *JsonFloatContext) SetTok(v antlr.Token) { s.tok = v }

//AltLabelStructDecl tokenTypeDecls

//AltLabelStructDecl tokenListDecls

//AltLabelStructDecl ruleContextDecls

//AltLabelStructDecl ruleContextListDecls

//AltLabelStructDecl attributeDecls

//getRuleContext
func (s *JsonFloatContext) GetRuleContext() antlr.RuleContext {
	return s
}

//getters
func (s *JsonFloatContext) JsonFloat() antlr.TerminalNode {
	return s.GetToken(AdlAstJsonFloat, 0)
}

func (s *JsonFloatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonFloatEntryListener); ok {
		listenerT.EnterJsonFloat(s)
	}
}

func (s *JsonFloatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonFloatExitListener); ok {
		listenerT.ExitJsonFloat(s)
	}
}

func (s *JsonFloatContext) VisitFunc(hdls antlr.ParserTreeVisitorHandlers, args ...interface{}) (result interface{}) {
	h := hdls.(*AdlAstHandlers)
	if h.EnterEveryRule != nil {
		h.EnterEveryRule(s)
	}
	if h.JsonFloat != nil {
		h.JsonFloat(s, h, args...)
	} else {
		s.VisitChildrenFunc(h, args...)
	}
	if h.ExitEveryRule != nil {
		h.ExitEveryRule(s)
	}
	return
}

func (s *JsonFloatContext) Visit(delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	switch t := delegate.(type) {
	case JsonFloatContextVisitor:
		return t.VisitJsonFloat(s, delegate, args...)
	default:
		return delegate.VisitChildren(s, delegate, args...)
	}
}

//END AltLabelStructDecl

//Begin AltLabelStructDecl

type JsonObjContext struct {
	*JsonContext
	//TokenDecl
	tok antlr.Token
}

func NewJsonObjContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *JsonObjContext {
	var p = new(JsonObjContext)

	p.JsonContext = NewEmptyJsonContext()
	p.parser = parser
	p.CopyFrom(ctx.(*JsonContext))

	return p
}

type IJsonObjContext interface {
	//Current rule
	IJsonContext
	// start internal
	//Gets for raw elements
	//  ruleGetterDecl
	//  ruleListGetterDecl
	AllJson() []IJsonContext
	//  ruleListIndexedGetterDecl
	Json(i int) IJsonContext

	//  tokenGetterDecl
	JsonObj() antlr.TerminalNode
	DOWN() antlr.TerminalNode
	UP() antlr.TerminalNode
	//  tokenListGetterDecl
	//  tokenListIndexedGetterDecl
	// end internal
	//
	//Gets for labeled elements
	//  tokenDecls
	GetTok() antlr.Token

	//  tokenTypeDecls
	//  tokenListDecls
	//  ruleContextDecls
	//  ruleContextListDecls
	//  attributeDecls

	// TODO dispatchMethods (needed?)
}

func (*JsonObjContext) IsJsonObjContext() {}

//AltLabelStructDecl tokenDecls
func (s *JsonObjContext) GetTok() antlr.Token  { return s.tok }
func (s *JsonObjContext) SetTok(v antlr.Token) { s.tok = v }

//AltLabelStructDecl tokenTypeDecls

//AltLabelStructDecl tokenListDecls

//AltLabelStructDecl ruleContextDecls

//AltLabelStructDecl ruleContextListDecls

//AltLabelStructDecl attributeDecls

//getRuleContext
func (s *JsonObjContext) GetRuleContext() antlr.RuleContext {
	return s
}

//getters
func (s *JsonObjContext) JsonObj() antlr.TerminalNode {
	return s.GetToken(AdlAstJsonObj, 0)
}

func (s *JsonObjContext) DOWN() antlr.TerminalNode {
	return s.GetToken(AdlAstDOWN, 0)
}

func (s *JsonObjContext) UP() antlr.TerminalNode {
	return s.GetToken(AdlAstUP, 0)
}

func (s *JsonObjContext) AllJson() []IJsonContext {
	//ContextRuleListGetterDecl
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*JsonContext)(nil)).Elem())
	var tst = make([]IJsonContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IJsonContext)
		}
	}

	return tst
}

func (s *JsonObjContext) Json(i int) IJsonContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*JsonContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IJsonContext)
}

func (s *JsonObjContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonObjEntryListener); ok {
		listenerT.EnterJsonObj(s)
	}
}

func (s *JsonObjContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonObjExitListener); ok {
		listenerT.ExitJsonObj(s)
	}
}

func (s *JsonObjContext) VisitFunc(hdls antlr.ParserTreeVisitorHandlers, args ...interface{}) (result interface{}) {
	h := hdls.(*AdlAstHandlers)
	if h.EnterEveryRule != nil {
		h.EnterEveryRule(s)
	}
	if h.JsonObj != nil {
		h.JsonObj(s, h, args...)
	} else {
		s.VisitChildrenFunc(h, args...)
	}
	if h.ExitEveryRule != nil {
		h.ExitEveryRule(s)
	}
	return
}

func (s *JsonObjContext) Visit(delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	switch t := delegate.(type) {
	case JsonObjContextVisitor:
		return t.VisitJsonObj(s, delegate, args...)
	default:
		return delegate.VisitChildren(s, delegate, args...)
	}
}

//END AltLabelStructDecl

//Begin AltLabelStructDecl

type JsonBoolContext struct {
	*JsonContext
	//TokenDecl
	tok antlr.Token
}

func NewJsonBoolContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *JsonBoolContext {
	var p = new(JsonBoolContext)

	p.JsonContext = NewEmptyJsonContext()
	p.parser = parser
	p.CopyFrom(ctx.(*JsonContext))

	return p
}

type IJsonBoolContext interface {
	//Current rule
	IJsonContext
	// start internal
	//Gets for raw elements
	//  ruleGetterDecl
	//  ruleListGetterDecl
	//  ruleListIndexedGetterDecl

	//  tokenGetterDecl
	JsonBool() antlr.TerminalNode
	//  tokenListGetterDecl
	//  tokenListIndexedGetterDecl
	// end internal
	//
	//Gets for labeled elements
	//  tokenDecls
	GetTok() antlr.Token

	//  tokenTypeDecls
	//  tokenListDecls
	//  ruleContextDecls
	//  ruleContextListDecls
	//  attributeDecls

	// TODO dispatchMethods (needed?)
}

func (*JsonBoolContext) IsJsonBoolContext() {}

//AltLabelStructDecl tokenDecls
func (s *JsonBoolContext) GetTok() antlr.Token  { return s.tok }
func (s *JsonBoolContext) SetTok(v antlr.Token) { s.tok = v }

//AltLabelStructDecl tokenTypeDecls

//AltLabelStructDecl tokenListDecls

//AltLabelStructDecl ruleContextDecls

//AltLabelStructDecl ruleContextListDecls

//AltLabelStructDecl attributeDecls

//getRuleContext
func (s *JsonBoolContext) GetRuleContext() antlr.RuleContext {
	return s
}

//getters
func (s *JsonBoolContext) JsonBool() antlr.TerminalNode {
	return s.GetToken(AdlAstJsonBool, 0)
}

func (s *JsonBoolContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonBoolEntryListener); ok {
		listenerT.EnterJsonBool(s)
	}
}

func (s *JsonBoolContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonBoolExitListener); ok {
		listenerT.ExitJsonBool(s)
	}
}

func (s *JsonBoolContext) VisitFunc(hdls antlr.ParserTreeVisitorHandlers, args ...interface{}) (result interface{}) {
	h := hdls.(*AdlAstHandlers)
	if h.EnterEveryRule != nil {
		h.EnterEveryRule(s)
	}
	if h.JsonBool != nil {
		h.JsonBool(s, h, args...)
	} else {
		s.VisitChildrenFunc(h, args...)
	}
	if h.ExitEveryRule != nil {
		h.ExitEveryRule(s)
	}
	return
}

func (s *JsonBoolContext) Visit(delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	switch t := delegate.(type) {
	case JsonBoolContextVisitor:
		return t.VisitJsonBool(s, delegate, args...)
	default:
		return delegate.VisitChildren(s, delegate, args...)
	}
}

//END AltLabelStructDecl

//Begin AltLabelStructDecl

type JsonErrorContext struct {
	*JsonContext
}

func NewJsonErrorContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *JsonErrorContext {
	var p = new(JsonErrorContext)

	p.JsonContext = NewEmptyJsonContext()
	p.parser = parser
	p.CopyFrom(ctx.(*JsonContext))

	return p
}

type IJsonErrorContext interface {
	//Current rule
	IJsonContext
	// start internal
	//Gets for raw elements
	//  ruleGetterDecl
	//  ruleListGetterDecl
	//  ruleListIndexedGetterDecl

	//  tokenGetterDecl
	ERROR() antlr.TerminalNode
	//  tokenListGetterDecl
	//  tokenListIndexedGetterDecl
	// end internal
	//
	//Gets for labeled elements
	//  tokenDecls
	//  tokenTypeDecls
	//  tokenListDecls
	//  ruleContextDecls
	//  ruleContextListDecls
	//  attributeDecls

	// TODO dispatchMethods (needed?)
}

func (*JsonErrorContext) IsJsonErrorContext() {}

//AltLabelStructDecl tokenDecls

//AltLabelStructDecl tokenTypeDecls

//AltLabelStructDecl tokenListDecls

//AltLabelStructDecl ruleContextDecls

//AltLabelStructDecl ruleContextListDecls

//AltLabelStructDecl attributeDecls

//getRuleContext
func (s *JsonErrorContext) GetRuleContext() antlr.RuleContext {
	return s
}

//getters
func (s *JsonErrorContext) ERROR() antlr.TerminalNode {
	return s.GetToken(AdlAstERROR, 0)
}

func (s *JsonErrorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonErrorEntryListener); ok {
		listenerT.EnterJsonError(s)
	}
}

func (s *JsonErrorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonErrorExitListener); ok {
		listenerT.ExitJsonError(s)
	}
}

func (s *JsonErrorContext) VisitFunc(hdls antlr.ParserTreeVisitorHandlers, args ...interface{}) (result interface{}) {
	h := hdls.(*AdlAstHandlers)
	if h.EnterEveryRule != nil {
		h.EnterEveryRule(s)
	}
	if h.JsonError != nil {
		h.JsonError(s, h, args...)
	} else {
		s.VisitChildrenFunc(h, args...)
	}
	if h.ExitEveryRule != nil {
		h.ExitEveryRule(s)
	}
	return
}

func (s *JsonErrorContext) Visit(delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	switch t := delegate.(type) {
	case JsonErrorContextVisitor:
		return t.VisitJsonError(s, delegate, args...)
	default:
		return delegate.VisitChildren(s, delegate, args...)
	}
}

//END AltLabelStructDecl

//Begin AltLabelStructDecl

type JsonIntContext struct {
	*JsonContext
	//TokenDecl
	tok antlr.Token
}

func NewJsonIntContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *JsonIntContext {
	var p = new(JsonIntContext)

	p.JsonContext = NewEmptyJsonContext()
	p.parser = parser
	p.CopyFrom(ctx.(*JsonContext))

	return p
}

type IJsonIntContext interface {
	//Current rule
	IJsonContext
	// start internal
	//Gets for raw elements
	//  ruleGetterDecl
	//  ruleListGetterDecl
	//  ruleListIndexedGetterDecl

	//  tokenGetterDecl
	JsonInt() antlr.TerminalNode
	//  tokenListGetterDecl
	//  tokenListIndexedGetterDecl
	// end internal
	//
	//Gets for labeled elements
	//  tokenDecls
	GetTok() antlr.Token

	//  tokenTypeDecls
	//  tokenListDecls
	//  ruleContextDecls
	//  ruleContextListDecls
	//  attributeDecls

	// TODO dispatchMethods (needed?)
}

func (*JsonIntContext) IsJsonIntContext() {}

//AltLabelStructDecl tokenDecls
func (s *JsonIntContext) GetTok() antlr.Token  { return s.tok }
func (s *JsonIntContext) SetTok(v antlr.Token) { s.tok = v }

//AltLabelStructDecl tokenTypeDecls

//AltLabelStructDecl tokenListDecls

//AltLabelStructDecl ruleContextDecls

//AltLabelStructDecl ruleContextListDecls

//AltLabelStructDecl attributeDecls

//getRuleContext
func (s *JsonIntContext) GetRuleContext() antlr.RuleContext {
	return s
}

//getters
func (s *JsonIntContext) JsonInt() antlr.TerminalNode {
	return s.GetToken(AdlAstJsonInt, 0)
}

func (s *JsonIntContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonIntEntryListener); ok {
		listenerT.EnterJsonInt(s)
	}
}

func (s *JsonIntContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonIntExitListener); ok {
		listenerT.ExitJsonInt(s)
	}
}

func (s *JsonIntContext) VisitFunc(hdls antlr.ParserTreeVisitorHandlers, args ...interface{}) (result interface{}) {
	h := hdls.(*AdlAstHandlers)
	if h.EnterEveryRule != nil {
		h.EnterEveryRule(s)
	}
	if h.JsonInt != nil {
		h.JsonInt(s, h, args...)
	} else {
		s.VisitChildrenFunc(h, args...)
	}
	if h.ExitEveryRule != nil {
		h.ExitEveryRule(s)
	}
	return
}

func (s *JsonIntContext) Visit(delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	switch t := delegate.(type) {
	case JsonIntContextVisitor:
		return t.VisitJsonInt(s, delegate, args...)
	default:
		return delegate.VisitChildren(s, delegate, args...)
	}
}

//END AltLabelStructDecl

//Begin AltLabelStructDecl

type JsonNullContext struct {
	*JsonContext
	//TokenDecl
	tok antlr.Token
}

func NewJsonNullContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *JsonNullContext {
	var p = new(JsonNullContext)

	p.JsonContext = NewEmptyJsonContext()
	p.parser = parser
	p.CopyFrom(ctx.(*JsonContext))

	return p
}

type IJsonNullContext interface {
	//Current rule
	IJsonContext
	// start internal
	//Gets for raw elements
	//  ruleGetterDecl
	//  ruleListGetterDecl
	//  ruleListIndexedGetterDecl

	//  tokenGetterDecl
	JsonNull() antlr.TerminalNode
	//  tokenListGetterDecl
	//  tokenListIndexedGetterDecl
	// end internal
	//
	//Gets for labeled elements
	//  tokenDecls
	GetTok() antlr.Token

	//  tokenTypeDecls
	//  tokenListDecls
	//  ruleContextDecls
	//  ruleContextListDecls
	//  attributeDecls

	// TODO dispatchMethods (needed?)
}

func (*JsonNullContext) IsJsonNullContext() {}

//AltLabelStructDecl tokenDecls
func (s *JsonNullContext) GetTok() antlr.Token  { return s.tok }
func (s *JsonNullContext) SetTok(v antlr.Token) { s.tok = v }

//AltLabelStructDecl tokenTypeDecls

//AltLabelStructDecl tokenListDecls

//AltLabelStructDecl ruleContextDecls

//AltLabelStructDecl ruleContextListDecls

//AltLabelStructDecl attributeDecls

//getRuleContext
func (s *JsonNullContext) GetRuleContext() antlr.RuleContext {
	return s
}

//getters
func (s *JsonNullContext) JsonNull() antlr.TerminalNode {
	return s.GetToken(AdlAstJsonNull, 0)
}

func (s *JsonNullContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonNullEntryListener); ok {
		listenerT.EnterJsonNull(s)
	}
}

func (s *JsonNullContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonNullExitListener); ok {
		listenerT.ExitJsonNull(s)
	}
}

func (s *JsonNullContext) VisitFunc(hdls antlr.ParserTreeVisitorHandlers, args ...interface{}) (result interface{}) {
	h := hdls.(*AdlAstHandlers)
	if h.EnterEveryRule != nil {
		h.EnterEveryRule(s)
	}
	if h.JsonNull != nil {
		h.JsonNull(s, h, args...)
	} else {
		s.VisitChildrenFunc(h, args...)
	}
	if h.ExitEveryRule != nil {
		h.ExitEveryRule(s)
	}
	return
}

func (s *JsonNullContext) Visit(delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	switch t := delegate.(type) {
	case JsonNullContextVisitor:
		return t.VisitJsonNull(s, delegate, args...)
	default:
		return delegate.VisitChildren(s, delegate, args...)
	}
}

//END AltLabelStructDecl

func (p *AdlAst) Json() (localctx IJsonContext) {
	localctx = NewJsonContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, AdlAstRULE_json)
	var //TokenTypeDecl
	_la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(221)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AdlAstJsonStr:
		localctx = NewJsonStrContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(193)
			var _m = p.Match(AdlAstJsonStr)
			localctx.(*JsonStrContext).tok = _m

		}

	case AdlAstJsonBool:
		localctx = NewJsonBoolContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(194)
			var _m = p.Match(AdlAstJsonBool)
			localctx.(*JsonBoolContext).tok = _m

		}

	case AdlAstJsonNull:
		localctx = NewJsonNullContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(195)
			var _m = p.Match(AdlAstJsonNull)
			localctx.(*JsonNullContext).tok = _m

		}

	case AdlAstJsonInt:
		localctx = NewJsonIntContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(196)
			var _m = p.Match(AdlAstJsonInt)
			localctx.(*JsonIntContext).tok = _m

		}

	case AdlAstJsonFloat:
		localctx = NewJsonFloatContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(197)
			var _m = p.Match(AdlAstJsonFloat)
			localctx.(*JsonFloatContext).tok = _m

		}

	case AdlAstJsonArray:
		localctx = NewJsonArrayContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(198)
			var _m = p.Match(AdlAstJsonArray)
			localctx.(*JsonArrayContext).tok = _m

		}
		p.SetState(207)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == AdlAstDOWN {
			{
				p.SetState(199)
				p.Match(AdlAstDOWN)
			}
			p.SetState(201)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for ok := true; ok; ok = (((_la-27)&-(0x1f+1)) == 0 && ((1<<uint((_la-27)))&((1<<(AdlAstERROR-27))|(1<<(AdlAstJsonStr-27))|(1<<(AdlAstJsonBool-27))|(1<<(AdlAstJsonNull-27))|(1<<(AdlAstJsonInt-27))|(1<<(AdlAstJsonFloat-27))|(1<<(AdlAstJsonArray-27))|(1<<(AdlAstJsonObj-27)))) != 0) {
				{
					p.SetState(200)
					p.Json()
				}

				p.SetState(203)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(205)
				p.Match(AdlAstUP)
			}

		}

	case AdlAstJsonObj:
		localctx = NewJsonObjContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(209)
			var _m = p.Match(AdlAstJsonObj)
			localctx.(*JsonObjContext).tok = _m

		}
		p.SetState(218)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == AdlAstDOWN {
			{
				p.SetState(210)
				p.Match(AdlAstDOWN)
			}
			p.SetState(212)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for ok := true; ok; ok = (((_la-27)&-(0x1f+1)) == 0 && ((1<<uint((_la-27)))&((1<<(AdlAstERROR-27))|(1<<(AdlAstJsonStr-27))|(1<<(AdlAstJsonBool-27))|(1<<(AdlAstJsonNull-27))|(1<<(AdlAstJsonInt-27))|(1<<(AdlAstJsonFloat-27))|(1<<(AdlAstJsonArray-27))|(1<<(AdlAstJsonObj-27)))) != 0) {
				{
					p.SetState(211)
					p.Json()
				}

				p.SetState(214)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(216)
				p.Match(AdlAstUP)
			}

		}

	case AdlAstERROR:
		localctx = NewJsonErrorContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(220)
			p.Match(AdlAstERROR)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}
