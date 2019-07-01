// Generated from AdlWo.g4 by ANTLR 4.7.

package adlwo // AdlWo
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 59, 258,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 3, 2, 3, 2, 3, 2, 3,
	2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3,
	4, 7, 4, 39, 10, 4, 12, 4, 14, 4, 42, 11, 4, 3, 4, 7, 4, 45, 10, 4, 12,
	4, 14, 4, 48, 11, 4, 3, 4, 7, 4, 51, 10, 4, 12, 4, 14, 4, 54, 11, 4, 3,
	4, 5, 4, 57, 10, 4, 3, 5, 3, 5, 5, 5, 61, 10, 5, 3, 6, 3, 6, 3, 6, 7, 6,
	66, 10, 6, 12, 6, 14, 6, 69, 11, 6, 3, 6, 5, 6, 72, 10, 6, 3, 6, 7, 6,
	75, 10, 6, 12, 6, 14, 6, 78, 11, 6, 3, 6, 5, 6, 81, 10, 6, 3, 6, 3, 6,
	3, 6, 7, 6, 86, 10, 6, 12, 6, 14, 6, 89, 11, 6, 3, 6, 5, 6, 92, 10, 6,
	3, 6, 7, 6, 95, 10, 6, 12, 6, 14, 6, 98, 11, 6, 3, 6, 5, 6, 101, 10, 6,
	3, 6, 3, 6, 3, 6, 7, 6, 106, 10, 6, 12, 6, 14, 6, 109, 11, 6, 3, 6, 5,
	6, 112, 10, 6, 3, 6, 5, 6, 115, 10, 6, 3, 6, 7, 6, 118, 10, 6, 12, 6, 14,
	6, 121, 11, 6, 3, 6, 5, 6, 124, 10, 6, 3, 6, 3, 6, 3, 6, 7, 6, 129, 10,
	6, 12, 6, 14, 6, 132, 11, 6, 3, 6, 5, 6, 135, 10, 6, 3, 6, 5, 6, 138, 10,
	6, 3, 6, 7, 6, 141, 10, 6, 12, 6, 14, 6, 144, 11, 6, 3, 6, 5, 6, 147, 10,
	6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3,
	6, 3, 6, 3, 6, 3, 6, 5, 6, 164, 10, 6, 3, 7, 3, 7, 3, 7, 7, 7, 169, 10,
	7, 12, 7, 14, 7, 172, 11, 7, 3, 7, 5, 7, 175, 10, 7, 3, 7, 5, 7, 178, 10,
	7, 3, 7, 5, 7, 181, 10, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 5, 8, 188, 10,
	8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 5, 8, 195, 10, 8, 3, 8, 3, 8, 3, 8, 3,
	8, 3, 8, 5, 8, 202, 10, 8, 5, 8, 204, 10, 8, 3, 9, 3, 9, 3, 9, 3, 9, 6,
	9, 210, 10, 9, 13, 9, 14, 9, 211, 3, 9, 3, 9, 5, 9, 216, 10, 9, 3, 10,
	3, 10, 3, 10, 6, 10, 221, 10, 10, 13, 10, 14, 10, 222, 3, 10, 3, 10, 5,
	10, 227, 10, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11,
	6, 11, 237, 10, 11, 13, 11, 14, 11, 238, 3, 11, 3, 11, 5, 11, 243, 10,
	11, 3, 11, 3, 11, 3, 11, 6, 11, 248, 10, 11, 13, 11, 14, 11, 249, 3, 11,
	3, 11, 5, 11, 254, 10, 11, 5, 11, 256, 10, 11, 3, 11, 2, 2, 12, 2, 4, 6,
	8, 10, 12, 14, 16, 18, 20, 2, 2, 2, 299, 2, 22, 3, 2, 2, 2, 4, 30, 3, 2,
	2, 2, 6, 35, 3, 2, 2, 2, 8, 60, 3, 2, 2, 2, 10, 163, 3, 2, 2, 2, 12, 165,
	3, 2, 2, 2, 14, 203, 3, 2, 2, 2, 16, 215, 3, 2, 2, 2, 18, 217, 3, 2, 2,
	2, 20, 255, 3, 2, 2, 2, 22, 23, 7, 26, 2, 2, 23, 24, 7, 30, 2, 2, 24, 25,
	7, 26, 2, 2, 25, 26, 5, 6, 4, 2, 26, 27, 7, 27, 2, 2, 27, 28, 7, 27, 2,
	2, 28, 29, 7, 2, 2, 3, 29, 3, 3, 2, 2, 2, 30, 31, 7, 26, 2, 2, 31, 32,
	5, 20, 11, 2, 32, 33, 7, 27, 2, 2, 33, 34, 7, 2, 2, 3, 34, 5, 3, 2, 2,
	2, 35, 56, 7, 31, 2, 2, 36, 40, 7, 26, 2, 2, 37, 39, 5, 14, 8, 2, 38, 37,
	3, 2, 2, 2, 39, 42, 3, 2, 2, 2, 40, 38, 3, 2, 2, 2, 40, 41, 3, 2, 2, 2,
	41, 46, 3, 2, 2, 2, 42, 40, 3, 2, 2, 2, 43, 45, 5, 8, 5, 2, 44, 43, 3,
	2, 2, 2, 45, 48, 3, 2, 2, 2, 46, 44, 3, 2, 2, 2, 46, 47, 3, 2, 2, 2, 47,
	52, 3, 2, 2, 2, 48, 46, 3, 2, 2, 2, 49, 51, 5, 10, 6, 2, 50, 49, 3, 2,
	2, 2, 51, 54, 3, 2, 2, 2, 52, 50, 3, 2, 2, 2, 52, 53, 3, 2, 2, 2, 53, 55,
	3, 2, 2, 2, 54, 52, 3, 2, 2, 2, 55, 57, 7, 27, 2, 2, 56, 36, 3, 2, 2, 2,
	56, 57, 3, 2, 2, 2, 57, 7, 3, 2, 2, 2, 58, 61, 7, 32, 2, 2, 59, 61, 7,
	33, 2, 2, 60, 58, 3, 2, 2, 2, 60, 59, 3, 2, 2, 2, 61, 9, 3, 2, 2, 2, 62,
	80, 7, 37, 2, 2, 63, 67, 7, 26, 2, 2, 64, 66, 5, 14, 8, 2, 65, 64, 3, 2,
	2, 2, 66, 69, 3, 2, 2, 2, 67, 65, 3, 2, 2, 2, 67, 68, 3, 2, 2, 2, 68, 71,
	3, 2, 2, 2, 69, 67, 3, 2, 2, 2, 70, 72, 7, 41, 2, 2, 71, 70, 3, 2, 2, 2,
	71, 72, 3, 2, 2, 2, 72, 76, 3, 2, 2, 2, 73, 75, 5, 12, 7, 2, 74, 73, 3,
	2, 2, 2, 75, 78, 3, 2, 2, 2, 76, 74, 3, 2, 2, 2, 76, 77, 3, 2, 2, 2, 77,
	79, 3, 2, 2, 2, 78, 76, 3, 2, 2, 2, 79, 81, 7, 27, 2, 2, 80, 63, 3, 2,
	2, 2, 80, 81, 3, 2, 2, 2, 81, 164, 3, 2, 2, 2, 82, 100, 7, 38, 2, 2, 83,
	87, 7, 26, 2, 2, 84, 86, 5, 14, 8, 2, 85, 84, 3, 2, 2, 2, 86, 89, 3, 2,
	2, 2, 87, 85, 3, 2, 2, 2, 87, 88, 3, 2, 2, 2, 88, 91, 3, 2, 2, 2, 89, 87,
	3, 2, 2, 2, 90, 92, 7, 41, 2, 2, 91, 90, 3, 2, 2, 2, 91, 92, 3, 2, 2, 2,
	92, 96, 3, 2, 2, 2, 93, 95, 5, 12, 7, 2, 94, 93, 3, 2, 2, 2, 95, 98, 3,
	2, 2, 2, 96, 94, 3, 2, 2, 2, 96, 97, 3, 2, 2, 2, 97, 99, 3, 2, 2, 2, 98,
	96, 3, 2, 2, 2, 99, 101, 7, 27, 2, 2, 100, 83, 3, 2, 2, 2, 100, 101, 3,
	2, 2, 2, 101, 164, 3, 2, 2, 2, 102, 123, 7, 40, 2, 2, 103, 107, 7, 26,
	2, 2, 104, 106, 5, 14, 8, 2, 105, 104, 3, 2, 2, 2, 106, 109, 3, 2, 2, 2,
	107, 105, 3, 2, 2, 2, 107, 108, 3, 2, 2, 2, 108, 111, 3, 2, 2, 2, 109,
	107, 3, 2, 2, 2, 110, 112, 7, 41, 2, 2, 111, 110, 3, 2, 2, 2, 111, 112,
	3, 2, 2, 2, 112, 114, 3, 2, 2, 2, 113, 115, 5, 16, 9, 2, 114, 113, 3, 2,
	2, 2, 114, 115, 3, 2, 2, 2, 115, 119, 3, 2, 2, 2, 116, 118, 5, 20, 11,
	2, 117, 116, 3, 2, 2, 2, 118, 121, 3, 2, 2, 2, 119, 117, 3, 2, 2, 2, 119,
	120, 3, 2, 2, 2, 120, 122, 3, 2, 2, 2, 121, 119, 3, 2, 2, 2, 122, 124,
	7, 27, 2, 2, 123, 103, 3, 2, 2, 2, 123, 124, 3, 2, 2, 2, 124, 164, 3, 2,
	2, 2, 125, 146, 7, 39, 2, 2, 126, 130, 7, 26, 2, 2, 127, 129, 5, 14, 8,
	2, 128, 127, 3, 2, 2, 2, 129, 132, 3, 2, 2, 2, 130, 128, 3, 2, 2, 2, 130,
	131, 3, 2, 2, 2, 131, 134, 3, 2, 2, 2, 132, 130, 3, 2, 2, 2, 133, 135,
	7, 41, 2, 2, 134, 133, 3, 2, 2, 2, 134, 135, 3, 2, 2, 2, 135, 137, 3, 2,
	2, 2, 136, 138, 5, 16, 9, 2, 137, 136, 3, 2, 2, 2, 137, 138, 3, 2, 2, 2,
	138, 142, 3, 2, 2, 2, 139, 141, 5, 20, 11, 2, 140, 139, 3, 2, 2, 2, 141,
	144, 3, 2, 2, 2, 142, 140, 3, 2, 2, 2, 142, 143, 3, 2, 2, 2, 143, 145,
	3, 2, 2, 2, 144, 142, 3, 2, 2, 2, 145, 147, 7, 27, 2, 2, 146, 126, 3, 2,
	2, 2, 146, 147, 3, 2, 2, 2, 147, 164, 3, 2, 2, 2, 148, 149, 7, 54, 2, 2,
	149, 150, 7, 26, 2, 2, 150, 151, 5, 20, 11, 2, 151, 152, 7, 27, 2, 2, 152,
	164, 3, 2, 2, 2, 153, 154, 7, 55, 2, 2, 154, 155, 7, 26, 2, 2, 155, 156,
	5, 20, 11, 2, 156, 157, 7, 27, 2, 2, 157, 164, 3, 2, 2, 2, 158, 159, 7,
	56, 2, 2, 159, 160, 7, 26, 2, 2, 160, 161, 5, 20, 11, 2, 161, 162, 7, 27,
	2, 2, 162, 164, 3, 2, 2, 2, 163, 62, 3, 2, 2, 2, 163, 82, 3, 2, 2, 2, 163,
	102, 3, 2, 2, 2, 163, 125, 3, 2, 2, 2, 163, 148, 3, 2, 2, 2, 163, 153,
	3, 2, 2, 2, 163, 158, 3, 2, 2, 2, 164, 11, 3, 2, 2, 2, 165, 180, 7, 45,
	2, 2, 166, 170, 7, 26, 2, 2, 167, 169, 5, 14, 8, 2, 168, 167, 3, 2, 2,
	2, 169, 172, 3, 2, 2, 2, 170, 168, 3, 2, 2, 2, 170, 171, 3, 2, 2, 2, 171,
	174, 3, 2, 2, 2, 172, 170, 3, 2, 2, 2, 173, 175, 5, 16, 9, 2, 174, 173,
	3, 2, 2, 2, 174, 175, 3, 2, 2, 2, 175, 177, 3, 2, 2, 2, 176, 178, 5, 20,
	11, 2, 177, 176, 3, 2, 2, 2, 177, 178, 3, 2, 2, 2, 178, 179, 3, 2, 2, 2,
	179, 181, 7, 27, 2, 2, 180, 166, 3, 2, 2, 2, 180, 181, 3, 2, 2, 2, 181,
	13, 3, 2, 2, 2, 182, 187, 7, 34, 2, 2, 183, 184, 7, 26, 2, 2, 184, 185,
	5, 20, 11, 2, 185, 186, 7, 27, 2, 2, 186, 188, 3, 2, 2, 2, 187, 183, 3,
	2, 2, 2, 187, 188, 3, 2, 2, 2, 188, 204, 3, 2, 2, 2, 189, 194, 7, 35, 2,
	2, 190, 191, 7, 26, 2, 2, 191, 192, 5, 20, 11, 2, 192, 193, 7, 27, 2, 2,
	193, 195, 3, 2, 2, 2, 194, 190, 3, 2, 2, 2, 194, 195, 3, 2, 2, 2, 195,
	204, 3, 2, 2, 2, 196, 201, 7, 36, 2, 2, 197, 198, 7, 26, 2, 2, 198, 199,
	5, 20, 11, 2, 199, 200, 7, 27, 2, 2, 200, 202, 3, 2, 2, 2, 201, 197, 3,
	2, 2, 2, 201, 202, 3, 2, 2, 2, 202, 204, 3, 2, 2, 2, 203, 182, 3, 2, 2,
	2, 203, 189, 3, 2, 2, 2, 203, 196, 3, 2, 2, 2, 204, 15, 3, 2, 2, 2, 205,
	216, 7, 42, 2, 2, 206, 207, 7, 43, 2, 2, 207, 209, 7, 26, 2, 2, 208, 210,
	5, 18, 10, 2, 209, 208, 3, 2, 2, 2, 210, 211, 3, 2, 2, 2, 211, 209, 3,
	2, 2, 2, 211, 212, 3, 2, 2, 2, 212, 213, 3, 2, 2, 2, 213, 214, 7, 27, 2,
	2, 214, 216, 3, 2, 2, 2, 215, 205, 3, 2, 2, 2, 215, 206, 3, 2, 2, 2, 216,
	17, 3, 2, 2, 2, 217, 226, 7, 44, 2, 2, 218, 220, 7, 26, 2, 2, 219, 221,
	5, 18, 10, 2, 220, 219, 3, 2, 2, 2, 221, 222, 3, 2, 2, 2, 222, 220, 3,
	2, 2, 2, 222, 223, 3, 2, 2, 2, 223, 224, 3, 2, 2, 2, 224, 225, 7, 27, 2,
	2, 225, 227, 3, 2, 2, 2, 226, 218, 3, 2, 2, 2, 226, 227, 3, 2, 2, 2, 227,
	19, 3, 2, 2, 2, 228, 256, 7, 47, 2, 2, 229, 256, 7, 48, 2, 2, 230, 256,
	7, 49, 2, 2, 231, 256, 7, 50, 2, 2, 232, 256, 7, 51, 2, 2, 233, 242, 7,
	52, 2, 2, 234, 236, 7, 26, 2, 2, 235, 237, 5, 20, 11, 2, 236, 235, 3, 2,
	2, 2, 237, 238, 3, 2, 2, 2, 238, 236, 3, 2, 2, 2, 238, 239, 3, 2, 2, 2,
	239, 240, 3, 2, 2, 2, 240, 241, 7, 27, 2, 2, 241, 243, 3, 2, 2, 2, 242,
	234, 3, 2, 2, 2, 242, 243, 3, 2, 2, 2, 243, 256, 3, 2, 2, 2, 244, 253,
	7, 53, 2, 2, 245, 247, 7, 26, 2, 2, 246, 248, 5, 20, 11, 2, 247, 246, 3,
	2, 2, 2, 248, 249, 3, 2, 2, 2, 249, 247, 3, 2, 2, 2, 249, 250, 3, 2, 2,
	2, 250, 251, 3, 2, 2, 2, 251, 252, 7, 27, 2, 2, 252, 254, 3, 2, 2, 2, 253,
	245, 3, 2, 2, 2, 253, 254, 3, 2, 2, 2, 254, 256, 3, 2, 2, 2, 255, 228,
	3, 2, 2, 2, 255, 229, 3, 2, 2, 2, 255, 230, 3, 2, 2, 2, 255, 231, 3, 2,
	2, 2, 255, 232, 3, 2, 2, 2, 255, 233, 3, 2, 2, 2, 255, 244, 3, 2, 2, 2,
	256, 21, 3, 2, 2, 2, 43, 40, 46, 52, 56, 60, 67, 71, 76, 80, 87, 91, 96,
	100, 107, 111, 114, 119, 123, 130, 134, 137, 142, 146, 163, 170, 174, 177,
	180, 187, 194, 201, 203, 211, 215, 222, 226, 238, 242, 249, 253, 255,
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
	"Module", "ImportModule", "ImportScopedName", "Annotation", "AnnotationNotScoped",
	"AnnotationScoped", "Struct", "Union", "Newtype", "Type", "TypeParam",
	"TypeExprPrimOrParam", "TypeExprTypeExpr", "TypeExprElem", "Field", "Json",
	"JsonStr", "JsonBool", "JsonNull", "JsonInt", "JsonFloat", "JsonArray",
	"JsonObj", "ModuleAnno", "DeclAnno", "FieldAnno", "DNAC", "Name", "Exnotation",
}

var ruleNames = []string{
	"adl", "json", "module", "import_", "tld", "nameBody", "annotation", "typeExpr_",
	"typeExprElem_", "jsonVal",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type AdlWo struct {
	*antlr.BaseParser
}

func NewAdlWo(input antlr.TokenStream) *AdlWo {
	this := new(AdlWo)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "AdlWo.g4"

	return this
}

// AdlWo tokens.
const (
	AdlWoEOF                 = antlr.TokenEOF
	AdlWoLCUR                = 1
	AdlWoRCUR                = 2
	AdlWoLSQ                 = 3
	AdlWoRSQ                 = 4
	AdlWoEQ                  = 5
	AdlWoDQ                  = 6
	AdlWoSQ                  = 7
	AdlWoSEMI                = 8
	AdlWoDCOLON              = 9
	AdlWoCOLON               = 10
	AdlWoDOT                 = 11
	AdlWoCOMMA               = 12
	AdlWoLCHEVR              = 13
	AdlWoRCHEVR              = 14
	AdlWoSTAR                = 15
	AdlWoAT                  = 16
	AdlWoSTR                 = 17
	AdlWoID                  = 18
	AdlWoINT                 = 19
	AdlWoFLT                 = 20
	AdlWoWS                  = 21
	AdlWoLINE_DOC            = 22
	AdlWoLINE_COMMENT        = 23
	AdlWoDOWN                = 24
	AdlWoUP                  = 25
	AdlWoROOT                = 26
	AdlWoERROR               = 27
	AdlWoADL                 = 28
	AdlWoModule              = 29
	AdlWoImportModule        = 30
	AdlWoImportScopedName    = 31
	AdlWoAnnotation          = 32
	AdlWoAnnotationNotScoped = 33
	AdlWoAnnotationScoped    = 34
	AdlWoStruct              = 35
	AdlWoUnion               = 36
	AdlWoNewtype             = 37
	AdlWoType                = 38
	AdlWoTypeParam           = 39
	AdlWoTypeExprPrimOrParam = 40
	AdlWoTypeExprTypeExpr    = 41
	AdlWoTypeExprElem        = 42
	AdlWoField               = 43
	AdlWoJson                = 44
	AdlWoJsonStr             = 45
	AdlWoJsonBool            = 46
	AdlWoJsonNull            = 47
	AdlWoJsonInt             = 48
	AdlWoJsonFloat           = 49
	AdlWoJsonArray           = 50
	AdlWoJsonObj             = 51
	AdlWoModuleAnno          = 52
	AdlWoDeclAnno            = 53
	AdlWoFieldAnno           = 54
	AdlWoDNAC                = 55
	AdlWoName                = 56
	AdlWoExnotation          = 57
)

// AdlWo rules.
const (
	AdlWoRULE_adl           = 0
	AdlWoRULE_json          = 1
	AdlWoRULE_module        = 2
	AdlWoRULE_import_       = 3
	AdlWoRULE_tld           = 4
	AdlWoRULE_nameBody      = 5
	AdlWoRULE_annotation    = 6
	AdlWoRULE_typeExpr_     = 7
	AdlWoRULE_typeExprElem_ = 8
	AdlWoRULE_jsonVal       = 9
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
	GetTok() antlr.Token

	//tokenTypeDecls
	//tokenListDecls
	//attributeDecls
	//tokenGetterDecl
	EOF() antlr.TerminalNode
	ADL() antlr.TerminalNode
	//tokenListGetterDecl
	AllDOWN() []antlr.TerminalNode
	AllUP() []antlr.TerminalNode
	//tokenListIndexedGetterDecl
	DOWN(i int) antlr.TerminalNode
	UP(i int) antlr.TerminalNode

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
	//TokenDecl
	tok antlr.Token
}

func NewEmptyAdlContext() *AdlContext {
	var p = new(AdlContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AdlWoRULE_adl
	return p
}

func (*AdlContext) IsAdlContext() {}

func NewAdlContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AdlContext {
	var p = new(AdlContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AdlWoRULE_adl

	return p
}

func (s *AdlContext) GetParser() antlr.Parser { return s.parser }

//StructDecl tokenDecls
func (s *AdlContext) GetTok() antlr.Token  { return s.tok }
func (s *AdlContext) SetTok(v antlr.Token) { s.tok = v }

//StructDecl tokenTypeDecls

//StructDecl tokenListDecls

//StructDecl ruleContextDecls

//StructDecl ruleContextListDecls

//StructDecl attributeDecls

// Getters
func (s *AdlContext) AllDOWN() []antlr.TerminalNode {
	return s.GetTokens(AdlWoDOWN)
}

func (s *AdlContext) DOWN(i int) antlr.TerminalNode {
	return s.GetToken(AdlWoDOWN, i)
}

func (s *AdlContext) Module() IModuleContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ModuleContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IModuleContext)
}

func (s *AdlContext) AllUP() []antlr.TerminalNode {
	return s.GetTokens(AdlWoUP)
}

func (s *AdlContext) UP(i int) antlr.TerminalNode {
	return s.GetToken(AdlWoUP, i)
}

func (s *AdlContext) EOF() antlr.TerminalNode {
	return s.GetToken(AdlWoEOF, 0)
}

func (s *AdlContext) ADL() antlr.TerminalNode {
	return s.GetToken(AdlWoADL, 0)
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
	h := hdls.(*AdlWoHandlers)
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

func (p *AdlWo) Adl() (localctx IAdlContext) {
	localctx = NewAdlContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, AdlWoRULE_adl)

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
		p.SetState(20)
		p.Match(AdlWoDOWN)
	}
	{
		p.SetState(21)
		var _m = p.Match(AdlWoADL)
		localctx.(*AdlContext).tok = _m

	}
	{
		p.SetState(22)
		p.Match(AdlWoDOWN)
	}
	{
		p.SetState(23)
		p.Module()
	}
	{
		p.SetState(24)
		p.Match(AdlWoUP)
	}
	{
		p.SetState(25)
		p.Match(AdlWoUP)
	}
	{
		p.SetState(26)
		p.Match(AdlWoEOF)
	}

	return localctx
}

type IJsonContext interface {
	antlr.ParserRuleContext
	// GetParser returns the parser.
	GetParser() antlr.Parser

	// start internal
	//  ruleGetterDecl
	JsonVal() IJsonValContext
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
	DOWN() antlr.TerminalNode
	UP() antlr.TerminalNode
	EOF() antlr.TerminalNode
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
	p.RuleIndex = AdlWoRULE_json
	return p
}

func (*JsonContext) IsJsonContext() {}

func NewJsonContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *JsonContext {
	var p = new(JsonContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AdlWoRULE_json

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
func (s *JsonContext) DOWN() antlr.TerminalNode {
	return s.GetToken(AdlWoDOWN, 0)
}

func (s *JsonContext) JsonVal() IJsonValContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*JsonValContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IJsonValContext)
}

func (s *JsonContext) UP() antlr.TerminalNode {
	return s.GetToken(AdlWoUP, 0)
}

func (s *JsonContext) EOF() antlr.TerminalNode {
	return s.GetToken(AdlWoEOF, 0)
}

//provideCopyFrom
func (s *JsonContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JsonContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

//dispatchMethods
func (s *JsonContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonEntryListener); ok {
		listenerT.EnterJson(s)
	}
}

func (s *JsonContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(JsonExitListener); ok {
		listenerT.ExitJson(s)
	}
}

func (s *JsonContext) VisitFunc(hdls antlr.ParserTreeVisitorHandlers, args ...interface{}) (result interface{}) {
	h := hdls.(*AdlWoHandlers)
	if h.EnterEveryRule != nil {
		h.EnterEveryRule(s)
	}
	if h.Json != nil {
		h.Json(s, h, args...)
	} else {
		s.VisitChildrenFunc(h, args...)
	}
	if h.ExitEveryRule != nil {
		h.ExitEveryRule(s)
	}
	return
}

func (s *JsonContext) Visit(delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	switch t := delegate.(type) {
	case JsonContextVisitor:
		return t.VisitJson(s, delegate, args...)
	default:
		return delegate.VisitChildren(s, delegate, args...)
	}
}

//extensionMembers

func (p *AdlWo) Json() (localctx IJsonContext) {
	localctx = NewJsonContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, AdlWoRULE_json)

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
		p.Match(AdlWoDOWN)
	}
	{
		p.SetState(29)
		p.JsonVal()
	}
	{
		p.SetState(30)
		p.Match(AdlWoUP)
	}
	{
		p.SetState(31)
		p.Match(AdlWoEOF)
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
	AllAnnotation() []IAnnotationContext
	AllImport_() []IImport_Context
	AllTld() []ITldContext
	//  ruleListIndexedGetterDecl
	Annotation(i int) IAnnotationContext
	Import_(i int) IImport_Context
	Tld(i int) ITldContext
	//  ruleContextDecls
	//  ruleContextListDecls
	// end internal
	//Gets for labeled elements
	//tokenDecls
	GetTok() antlr.Token

	//tokenTypeDecls
	//tokenListDecls
	//attributeDecls
	//tokenGetterDecl
	Module() antlr.TerminalNode
	DOWN() antlr.TerminalNode
	UP() antlr.TerminalNode
	//tokenListGetterDecl
	//tokenListIndexedGetterDecl

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
	tok antlr.Token
}

func NewEmptyModuleContext() *ModuleContext {
	var p = new(ModuleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AdlWoRULE_module
	return p
}

func (*ModuleContext) IsModuleContext() {}

func NewModuleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ModuleContext {
	var p = new(ModuleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AdlWoRULE_module

	return p
}

func (s *ModuleContext) GetParser() antlr.Parser { return s.parser }

//StructDecl tokenDecls
func (s *ModuleContext) GetTok() antlr.Token  { return s.tok }
func (s *ModuleContext) SetTok(v antlr.Token) { s.tok = v }

//StructDecl tokenTypeDecls

//StructDecl tokenListDecls

//StructDecl ruleContextDecls

//StructDecl ruleContextListDecls

//StructDecl attributeDecls

// Getters
func (s *ModuleContext) Module() antlr.TerminalNode {
	return s.GetToken(AdlWoModule, 0)
}

func (s *ModuleContext) DOWN() antlr.TerminalNode {
	return s.GetToken(AdlWoDOWN, 0)
}

func (s *ModuleContext) UP() antlr.TerminalNode {
	return s.GetToken(AdlWoUP, 0)
}

func (s *ModuleContext) AllAnnotation() []IAnnotationContext {
	//ContextRuleListGetterDecl
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*AnnotationContext)(nil)).Elem())
	var tst = make([]IAnnotationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAnnotationContext)
		}
	}

	return tst
}

func (s *ModuleContext) Annotation(i int) IAnnotationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*AnnotationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAnnotationContext)
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

func (s *ModuleContext) AllTld() []ITldContext {
	//ContextRuleListGetterDecl
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*TldContext)(nil)).Elem())
	var tst = make([]ITldContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITldContext)
		}
	}

	return tst
}

func (s *ModuleContext) Tld(i int) ITldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*TldContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITldContext)
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
	h := hdls.(*AdlWoHandlers)
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

func (p *AdlWo) Module() (localctx IModuleContext) {
	localctx = NewModuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, AdlWoRULE_module)
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
		p.SetState(33)
		var _m = p.Match(AdlWoModule)
		localctx.(*ModuleContext).tok = _m

	}
	p.SetState(54)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == AdlWoDOWN {
		{
			p.SetState(34)
			p.Match(AdlWoDOWN)
		}
		p.SetState(38)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(AdlWoAnnotation-32))|(1<<(AdlWoAnnotationNotScoped-32))|(1<<(AdlWoAnnotationScoped-32)))) != 0 {
			{
				p.SetState(35)
				p.Annotation()
			}

			p.SetState(40)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		p.SetState(44)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == AdlWoImportModule || _la == AdlWoImportScopedName {
			{
				p.SetState(41)
				p.Import_()
			}

			p.SetState(46)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		p.SetState(50)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ((_la-35)&-(0x1f+1)) == 0 && ((1<<uint((_la-35)))&((1<<(AdlWoStruct-35))|(1<<(AdlWoUnion-35))|(1<<(AdlWoNewtype-35))|(1<<(AdlWoType-35))|(1<<(AdlWoModuleAnno-35))|(1<<(AdlWoDeclAnno-35))|(1<<(AdlWoFieldAnno-35)))) != 0 {
			{
				p.SetState(47)
				p.Tld()
			}

			p.SetState(52)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(53)
			p.Match(AdlWoUP)
		}

	}

	return localctx
}

type IImport_Context interface {
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
}

func NewEmptyImport_Context() *Import_Context {
	var p = new(Import_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AdlWoRULE_import_
	return p
}

func (*Import_Context) IsImport_Context() {}

func NewImport_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Import_Context {
	var p = new(Import_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AdlWoRULE_import_

	return p
}

func (s *Import_Context) GetParser() antlr.Parser { return s.parser }

//StructDecl tokenDecls

//StructDecl tokenTypeDecls

//StructDecl tokenListDecls

//StructDecl ruleContextDecls

//StructDecl ruleContextListDecls

//StructDecl attributeDecls

// Getters

//provideCopyFrom
func (s *Import_Context) CopyFrom(ctx *Import_Context) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *Import_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Import_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

//dispatchMethods
//extensionMembers

//Begin AltLabelStructDecl

type ImportScopedModuleContext struct {
	*Import_Context
}

func NewImportScopedModuleContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ImportScopedModuleContext {
	var p = new(ImportScopedModuleContext)

	p.Import_Context = NewEmptyImport_Context()
	p.parser = parser
	p.CopyFrom(ctx.(*Import_Context))

	return p
}

type IImportScopedModuleContext interface {
	//Current rule
	IImport_Context
	// start internal
	//Gets for raw elements
	//  ruleGetterDecl
	//  ruleListGetterDecl
	//  ruleListIndexedGetterDecl

	//  tokenGetterDecl
	ImportScopedName() antlr.TerminalNode
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

func (*ImportScopedModuleContext) IsImportScopedModuleContext() {}

//AltLabelStructDecl tokenDecls

//AltLabelStructDecl tokenTypeDecls

//AltLabelStructDecl tokenListDecls

//AltLabelStructDecl ruleContextDecls

//AltLabelStructDecl ruleContextListDecls

//AltLabelStructDecl attributeDecls

//getRuleContext
func (s *ImportScopedModuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

//getters
func (s *ImportScopedModuleContext) ImportScopedName() antlr.TerminalNode {
	return s.GetToken(AdlWoImportScopedName, 0)
}

func (s *ImportScopedModuleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ImportScopedModuleEntryListener); ok {
		listenerT.EnterImportScopedModule(s)
	}
}

func (s *ImportScopedModuleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ImportScopedModuleExitListener); ok {
		listenerT.ExitImportScopedModule(s)
	}
}

func (s *ImportScopedModuleContext) VisitFunc(hdls antlr.ParserTreeVisitorHandlers, args ...interface{}) (result interface{}) {
	h := hdls.(*AdlWoHandlers)
	if h.EnterEveryRule != nil {
		h.EnterEveryRule(s)
	}
	if h.ImportScopedModule != nil {
		h.ImportScopedModule(s, h, args...)
	} else {
		s.VisitChildrenFunc(h, args...)
	}
	if h.ExitEveryRule != nil {
		h.ExitEveryRule(s)
	}
	return
}

func (s *ImportScopedModuleContext) Visit(delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	switch t := delegate.(type) {
	case ImportScopedModuleContextVisitor:
		return t.VisitImportScopedModule(s, delegate, args...)
	default:
		return delegate.VisitChildren(s, delegate, args...)
	}
}

//END AltLabelStructDecl

//Begin AltLabelStructDecl

type ImportModuleContext struct {
	*Import_Context
}

func NewImportModuleContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ImportModuleContext {
	var p = new(ImportModuleContext)

	p.Import_Context = NewEmptyImport_Context()
	p.parser = parser
	p.CopyFrom(ctx.(*Import_Context))

	return p
}

type IImportModuleContext interface {
	//Current rule
	IImport_Context
	// start internal
	//Gets for raw elements
	//  ruleGetterDecl
	//  ruleListGetterDecl
	//  ruleListIndexedGetterDecl

	//  tokenGetterDecl
	ImportModule() antlr.TerminalNode
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

func (*ImportModuleContext) IsImportModuleContext() {}

//AltLabelStructDecl tokenDecls

//AltLabelStructDecl tokenTypeDecls

//AltLabelStructDecl tokenListDecls

//AltLabelStructDecl ruleContextDecls

//AltLabelStructDecl ruleContextListDecls

//AltLabelStructDecl attributeDecls

//getRuleContext
func (s *ImportModuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

//getters
func (s *ImportModuleContext) ImportModule() antlr.TerminalNode {
	return s.GetToken(AdlWoImportModule, 0)
}

func (s *ImportModuleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ImportModuleEntryListener); ok {
		listenerT.EnterImportModule(s)
	}
}

func (s *ImportModuleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ImportModuleExitListener); ok {
		listenerT.ExitImportModule(s)
	}
}

func (s *ImportModuleContext) VisitFunc(hdls antlr.ParserTreeVisitorHandlers, args ...interface{}) (result interface{}) {
	h := hdls.(*AdlWoHandlers)
	if h.EnterEveryRule != nil {
		h.EnterEveryRule(s)
	}
	if h.ImportModule != nil {
		h.ImportModule(s, h, args...)
	} else {
		s.VisitChildrenFunc(h, args...)
	}
	if h.ExitEveryRule != nil {
		h.ExitEveryRule(s)
	}
	return
}

func (s *ImportModuleContext) Visit(delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	switch t := delegate.(type) {
	case ImportModuleContextVisitor:
		return t.VisitImportModule(s, delegate, args...)
	default:
		return delegate.VisitChildren(s, delegate, args...)
	}
}

//END AltLabelStructDecl

func (p *AdlWo) Import_() (localctx IImport_Context) {
	localctx = NewImport_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, AdlWoRULE_import_)

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

	p.SetState(58)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AdlWoImportModule:
		localctx = NewImportModuleContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(56)
			p.Match(AdlWoImportModule)
		}

	case AdlWoImportScopedName:
		localctx = NewImportScopedModuleContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(57)
			p.Match(AdlWoImportScopedName)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

type ITldContext interface {
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

	// IsTldContext differentiates from other interfaces.
	//copyStruct,GetRuleContext and ToStringTree  from embedded

	//<if(dispatchMethods)>
	//<dispatchMethods; separator="\n\n">
	//<endif>

	//<if(extensionMembers)>
	//<extensionMembers; separator="\n\n">
	//<endif>
}

type TldContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTldContext() *TldContext {
	var p = new(TldContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AdlWoRULE_tld
	return p
}

func (*TldContext) IsTldContext() {}

func NewTldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TldContext {
	var p = new(TldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AdlWoRULE_tld

	return p
}

func (s *TldContext) GetParser() antlr.Parser { return s.parser }

//StructDecl tokenDecls

//StructDecl tokenTypeDecls

//StructDecl tokenListDecls

//StructDecl ruleContextDecls

//StructDecl ruleContextListDecls

//StructDecl attributeDecls

// Getters

//provideCopyFrom
func (s *TldContext) CopyFrom(ctx *TldContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *TldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

//dispatchMethods
//extensionMembers

//Begin AltLabelStructDecl

type TypeContext struct {
	*TldContext
	//TokenDecl
	tok antlr.Token
}

func NewTypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypeContext {
	var p = new(TypeContext)

	p.TldContext = NewEmptyTldContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TldContext))

	return p
}

type ITypeContext interface {
	//Current rule
	ITldContext
	// start internal
	//Gets for raw elements
	//  ruleGetterDecl
	TypeExpr_() ITypeExpr_Context
	//  ruleListGetterDecl
	AllAnnotation() []IAnnotationContext
	AllJsonVal() []IJsonValContext
	//  ruleListIndexedGetterDecl
	Annotation(i int) IAnnotationContext
	JsonVal(i int) IJsonValContext

	//  tokenGetterDecl
	Type() antlr.TerminalNode
	DOWN() antlr.TerminalNode
	UP() antlr.TerminalNode
	TypeParam() antlr.TerminalNode
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

func (*TypeContext) IsTypeContext() {}

//AltLabelStructDecl tokenDecls
func (s *TypeContext) GetTok() antlr.Token  { return s.tok }
func (s *TypeContext) SetTok(v antlr.Token) { s.tok = v }

//AltLabelStructDecl tokenTypeDecls

//AltLabelStructDecl tokenListDecls

//AltLabelStructDecl ruleContextDecls

//AltLabelStructDecl ruleContextListDecls

//AltLabelStructDecl attributeDecls

//getRuleContext
func (s *TypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

//getters
func (s *TypeContext) Type() antlr.TerminalNode {
	return s.GetToken(AdlWoType, 0)
}

func (s *TypeContext) DOWN() antlr.TerminalNode {
	return s.GetToken(AdlWoDOWN, 0)
}

func (s *TypeContext) UP() antlr.TerminalNode {
	return s.GetToken(AdlWoUP, 0)
}

func (s *TypeContext) AllAnnotation() []IAnnotationContext {
	//ContextRuleListGetterDecl
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*AnnotationContext)(nil)).Elem())
	var tst = make([]IAnnotationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAnnotationContext)
		}
	}

	return tst
}

func (s *TypeContext) Annotation(i int) IAnnotationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*AnnotationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAnnotationContext)
}

func (s *TypeContext) TypeParam() antlr.TerminalNode {
	return s.GetToken(AdlWoTypeParam, 0)
}

func (s *TypeContext) TypeExpr_() ITypeExpr_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*TypeExpr_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeExpr_Context)
}

func (s *TypeContext) AllJsonVal() []IJsonValContext {
	//ContextRuleListGetterDecl
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*JsonValContext)(nil)).Elem())
	var tst = make([]IJsonValContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IJsonValContext)
		}
	}

	return tst
}

func (s *TypeContext) JsonVal(i int) IJsonValContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*JsonValContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IJsonValContext)
}

func (s *TypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TypeEntryListener); ok {
		listenerT.EnterType(s)
	}
}

func (s *TypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TypeExitListener); ok {
		listenerT.ExitType(s)
	}
}

func (s *TypeContext) VisitFunc(hdls antlr.ParserTreeVisitorHandlers, args ...interface{}) (result interface{}) {
	h := hdls.(*AdlWoHandlers)
	if h.EnterEveryRule != nil {
		h.EnterEveryRule(s)
	}
	if h.Type != nil {
		h.Type(s, h, args...)
	} else {
		s.VisitChildrenFunc(h, args...)
	}
	if h.ExitEveryRule != nil {
		h.ExitEveryRule(s)
	}
	return
}

func (s *TypeContext) Visit(delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	switch t := delegate.(type) {
	case TypeContextVisitor:
		return t.VisitType(s, delegate, args...)
	default:
		return delegate.VisitChildren(s, delegate, args...)
	}
}

//END AltLabelStructDecl

//Begin AltLabelStructDecl

type NewtypeContext struct {
	*TldContext
	//TokenDecl
	tok antlr.Token
}

func NewNewtypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NewtypeContext {
	var p = new(NewtypeContext)

	p.TldContext = NewEmptyTldContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TldContext))

	return p
}

type INewtypeContext interface {
	//Current rule
	ITldContext
	// start internal
	//Gets for raw elements
	//  ruleGetterDecl
	TypeExpr_() ITypeExpr_Context
	//  ruleListGetterDecl
	AllAnnotation() []IAnnotationContext
	AllJsonVal() []IJsonValContext
	//  ruleListIndexedGetterDecl
	Annotation(i int) IAnnotationContext
	JsonVal(i int) IJsonValContext

	//  tokenGetterDecl
	Newtype() antlr.TerminalNode
	DOWN() antlr.TerminalNode
	UP() antlr.TerminalNode
	TypeParam() antlr.TerminalNode
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

func (*NewtypeContext) IsNewtypeContext() {}

//AltLabelStructDecl tokenDecls
func (s *NewtypeContext) GetTok() antlr.Token  { return s.tok }
func (s *NewtypeContext) SetTok(v antlr.Token) { s.tok = v }

//AltLabelStructDecl tokenTypeDecls

//AltLabelStructDecl tokenListDecls

//AltLabelStructDecl ruleContextDecls

//AltLabelStructDecl ruleContextListDecls

//AltLabelStructDecl attributeDecls

//getRuleContext
func (s *NewtypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

//getters
func (s *NewtypeContext) Newtype() antlr.TerminalNode {
	return s.GetToken(AdlWoNewtype, 0)
}

func (s *NewtypeContext) DOWN() antlr.TerminalNode {
	return s.GetToken(AdlWoDOWN, 0)
}

func (s *NewtypeContext) UP() antlr.TerminalNode {
	return s.GetToken(AdlWoUP, 0)
}

func (s *NewtypeContext) AllAnnotation() []IAnnotationContext {
	//ContextRuleListGetterDecl
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*AnnotationContext)(nil)).Elem())
	var tst = make([]IAnnotationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAnnotationContext)
		}
	}

	return tst
}

func (s *NewtypeContext) Annotation(i int) IAnnotationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*AnnotationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAnnotationContext)
}

func (s *NewtypeContext) TypeParam() antlr.TerminalNode {
	return s.GetToken(AdlWoTypeParam, 0)
}

func (s *NewtypeContext) TypeExpr_() ITypeExpr_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*TypeExpr_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeExpr_Context)
}

func (s *NewtypeContext) AllJsonVal() []IJsonValContext {
	//ContextRuleListGetterDecl
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*JsonValContext)(nil)).Elem())
	var tst = make([]IJsonValContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IJsonValContext)
		}
	}

	return tst
}

func (s *NewtypeContext) JsonVal(i int) IJsonValContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*JsonValContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IJsonValContext)
}

func (s *NewtypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NewtypeEntryListener); ok {
		listenerT.EnterNewtype(s)
	}
}

func (s *NewtypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NewtypeExitListener); ok {
		listenerT.ExitNewtype(s)
	}
}

func (s *NewtypeContext) VisitFunc(hdls antlr.ParserTreeVisitorHandlers, args ...interface{}) (result interface{}) {
	h := hdls.(*AdlWoHandlers)
	if h.EnterEveryRule != nil {
		h.EnterEveryRule(s)
	}
	if h.Newtype != nil {
		h.Newtype(s, h, args...)
	} else {
		s.VisitChildrenFunc(h, args...)
	}
	if h.ExitEveryRule != nil {
		h.ExitEveryRule(s)
	}
	return
}

func (s *NewtypeContext) Visit(delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	switch t := delegate.(type) {
	case NewtypeContextVisitor:
		return t.VisitNewtype(s, delegate, args...)
	default:
		return delegate.VisitChildren(s, delegate, args...)
	}
}

//END AltLabelStructDecl

//Begin AltLabelStructDecl

type ModAnnoContext struct {
	*TldContext
	//TokenDecl
	tok antlr.Token
}

func NewModAnnoContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ModAnnoContext {
	var p = new(ModAnnoContext)

	p.TldContext = NewEmptyTldContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TldContext))

	return p
}

type IModAnnoContext interface {
	//Current rule
	ITldContext
	// start internal
	//Gets for raw elements
	//  ruleGetterDecl
	JsonVal() IJsonValContext
	//  ruleListGetterDecl
	//  ruleListIndexedGetterDecl

	//  tokenGetterDecl
	DOWN() antlr.TerminalNode
	UP() antlr.TerminalNode
	ModuleAnno() antlr.TerminalNode
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

func (*ModAnnoContext) IsModAnnoContext() {}

//AltLabelStructDecl tokenDecls
func (s *ModAnnoContext) GetTok() antlr.Token  { return s.tok }
func (s *ModAnnoContext) SetTok(v antlr.Token) { s.tok = v }

//AltLabelStructDecl tokenTypeDecls

//AltLabelStructDecl tokenListDecls

//AltLabelStructDecl ruleContextDecls

//AltLabelStructDecl ruleContextListDecls

//AltLabelStructDecl attributeDecls

//getRuleContext
func (s *ModAnnoContext) GetRuleContext() antlr.RuleContext {
	return s
}

//getters
func (s *ModAnnoContext) DOWN() antlr.TerminalNode {
	return s.GetToken(AdlWoDOWN, 0)
}

func (s *ModAnnoContext) JsonVal() IJsonValContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*JsonValContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IJsonValContext)
}

func (s *ModAnnoContext) UP() antlr.TerminalNode {
	return s.GetToken(AdlWoUP, 0)
}

func (s *ModAnnoContext) ModuleAnno() antlr.TerminalNode {
	return s.GetToken(AdlWoModuleAnno, 0)
}

func (s *ModAnnoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ModAnnoEntryListener); ok {
		listenerT.EnterModAnno(s)
	}
}

func (s *ModAnnoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ModAnnoExitListener); ok {
		listenerT.ExitModAnno(s)
	}
}

func (s *ModAnnoContext) VisitFunc(hdls antlr.ParserTreeVisitorHandlers, args ...interface{}) (result interface{}) {
	h := hdls.(*AdlWoHandlers)
	if h.EnterEveryRule != nil {
		h.EnterEveryRule(s)
	}
	if h.ModAnno != nil {
		h.ModAnno(s, h, args...)
	} else {
		s.VisitChildrenFunc(h, args...)
	}
	if h.ExitEveryRule != nil {
		h.ExitEveryRule(s)
	}
	return
}

func (s *ModAnnoContext) Visit(delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	switch t := delegate.(type) {
	case ModAnnoContextVisitor:
		return t.VisitModAnno(s, delegate, args...)
	default:
		return delegate.VisitChildren(s, delegate, args...)
	}
}

//END AltLabelStructDecl

//Begin AltLabelStructDecl

type FieldAnnoContext struct {
	*TldContext
	//TokenDecl
	tok antlr.Token
}

func NewFieldAnnoContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FieldAnnoContext {
	var p = new(FieldAnnoContext)

	p.TldContext = NewEmptyTldContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TldContext))

	return p
}

type IFieldAnnoContext interface {
	//Current rule
	ITldContext
	// start internal
	//Gets for raw elements
	//  ruleGetterDecl
	JsonVal() IJsonValContext
	//  ruleListGetterDecl
	//  ruleListIndexedGetterDecl

	//  tokenGetterDecl
	DOWN() antlr.TerminalNode
	UP() antlr.TerminalNode
	FieldAnno() antlr.TerminalNode
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

func (*FieldAnnoContext) IsFieldAnnoContext() {}

//AltLabelStructDecl tokenDecls
func (s *FieldAnnoContext) GetTok() antlr.Token  { return s.tok }
func (s *FieldAnnoContext) SetTok(v antlr.Token) { s.tok = v }

//AltLabelStructDecl tokenTypeDecls

//AltLabelStructDecl tokenListDecls

//AltLabelStructDecl ruleContextDecls

//AltLabelStructDecl ruleContextListDecls

//AltLabelStructDecl attributeDecls

//getRuleContext
func (s *FieldAnnoContext) GetRuleContext() antlr.RuleContext {
	return s
}

//getters
func (s *FieldAnnoContext) DOWN() antlr.TerminalNode {
	return s.GetToken(AdlWoDOWN, 0)
}

func (s *FieldAnnoContext) JsonVal() IJsonValContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*JsonValContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IJsonValContext)
}

func (s *FieldAnnoContext) UP() antlr.TerminalNode {
	return s.GetToken(AdlWoUP, 0)
}

func (s *FieldAnnoContext) FieldAnno() antlr.TerminalNode {
	return s.GetToken(AdlWoFieldAnno, 0)
}

func (s *FieldAnnoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FieldAnnoEntryListener); ok {
		listenerT.EnterFieldAnno(s)
	}
}

func (s *FieldAnnoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FieldAnnoExitListener); ok {
		listenerT.ExitFieldAnno(s)
	}
}

func (s *FieldAnnoContext) VisitFunc(hdls antlr.ParserTreeVisitorHandlers, args ...interface{}) (result interface{}) {
	h := hdls.(*AdlWoHandlers)
	if h.EnterEveryRule != nil {
		h.EnterEveryRule(s)
	}
	if h.FieldAnno != nil {
		h.FieldAnno(s, h, args...)
	} else {
		s.VisitChildrenFunc(h, args...)
	}
	if h.ExitEveryRule != nil {
		h.ExitEveryRule(s)
	}
	return
}

func (s *FieldAnnoContext) Visit(delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	switch t := delegate.(type) {
	case FieldAnnoContextVisitor:
		return t.VisitFieldAnno(s, delegate, args...)
	default:
		return delegate.VisitChildren(s, delegate, args...)
	}
}

//END AltLabelStructDecl

//Begin AltLabelStructDecl

type UnionContext struct {
	*TldContext
	//TokenDecl
	tok antlr.Token
}

func NewUnionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UnionContext {
	var p = new(UnionContext)

	p.TldContext = NewEmptyTldContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TldContext))

	return p
}

type IUnionContext interface {
	//Current rule
	ITldContext
	// start internal
	//Gets for raw elements
	//  ruleGetterDecl
	//  ruleListGetterDecl
	AllAnnotation() []IAnnotationContext
	AllNameBody() []INameBodyContext
	//  ruleListIndexedGetterDecl
	Annotation(i int) IAnnotationContext
	NameBody(i int) INameBodyContext

	//  tokenGetterDecl
	Union() antlr.TerminalNode
	DOWN() antlr.TerminalNode
	UP() antlr.TerminalNode
	TypeParam() antlr.TerminalNode
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

func (*UnionContext) IsUnionContext() {}

//AltLabelStructDecl tokenDecls
func (s *UnionContext) GetTok() antlr.Token  { return s.tok }
func (s *UnionContext) SetTok(v antlr.Token) { s.tok = v }

//AltLabelStructDecl tokenTypeDecls

//AltLabelStructDecl tokenListDecls

//AltLabelStructDecl ruleContextDecls

//AltLabelStructDecl ruleContextListDecls

//AltLabelStructDecl attributeDecls

//getRuleContext
func (s *UnionContext) GetRuleContext() antlr.RuleContext {
	return s
}

//getters
func (s *UnionContext) Union() antlr.TerminalNode {
	return s.GetToken(AdlWoUnion, 0)
}

func (s *UnionContext) DOWN() antlr.TerminalNode {
	return s.GetToken(AdlWoDOWN, 0)
}

func (s *UnionContext) UP() antlr.TerminalNode {
	return s.GetToken(AdlWoUP, 0)
}

func (s *UnionContext) AllAnnotation() []IAnnotationContext {
	//ContextRuleListGetterDecl
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*AnnotationContext)(nil)).Elem())
	var tst = make([]IAnnotationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAnnotationContext)
		}
	}

	return tst
}

func (s *UnionContext) Annotation(i int) IAnnotationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*AnnotationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAnnotationContext)
}

func (s *UnionContext) TypeParam() antlr.TerminalNode {
	return s.GetToken(AdlWoTypeParam, 0)
}

func (s *UnionContext) AllNameBody() []INameBodyContext {
	//ContextRuleListGetterDecl
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*NameBodyContext)(nil)).Elem())
	var tst = make([]INameBodyContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INameBodyContext)
		}
	}

	return tst
}

func (s *UnionContext) NameBody(i int) INameBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*NameBodyContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INameBodyContext)
}

func (s *UnionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UnionEntryListener); ok {
		listenerT.EnterUnion(s)
	}
}

func (s *UnionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UnionExitListener); ok {
		listenerT.ExitUnion(s)
	}
}

func (s *UnionContext) VisitFunc(hdls antlr.ParserTreeVisitorHandlers, args ...interface{}) (result interface{}) {
	h := hdls.(*AdlWoHandlers)
	if h.EnterEveryRule != nil {
		h.EnterEveryRule(s)
	}
	if h.Union != nil {
		h.Union(s, h, args...)
	} else {
		s.VisitChildrenFunc(h, args...)
	}
	if h.ExitEveryRule != nil {
		h.ExitEveryRule(s)
	}
	return
}

func (s *UnionContext) Visit(delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	switch t := delegate.(type) {
	case UnionContextVisitor:
		return t.VisitUnion(s, delegate, args...)
	default:
		return delegate.VisitChildren(s, delegate, args...)
	}
}

//END AltLabelStructDecl

//Begin AltLabelStructDecl

type DeclAnnoContext struct {
	*TldContext
	//TokenDecl
	tok antlr.Token
}

func NewDeclAnnoContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DeclAnnoContext {
	var p = new(DeclAnnoContext)

	p.TldContext = NewEmptyTldContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TldContext))

	return p
}

type IDeclAnnoContext interface {
	//Current rule
	ITldContext
	// start internal
	//Gets for raw elements
	//  ruleGetterDecl
	JsonVal() IJsonValContext
	//  ruleListGetterDecl
	//  ruleListIndexedGetterDecl

	//  tokenGetterDecl
	DOWN() antlr.TerminalNode
	UP() antlr.TerminalNode
	DeclAnno() antlr.TerminalNode
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

func (*DeclAnnoContext) IsDeclAnnoContext() {}

//AltLabelStructDecl tokenDecls
func (s *DeclAnnoContext) GetTok() antlr.Token  { return s.tok }
func (s *DeclAnnoContext) SetTok(v antlr.Token) { s.tok = v }

//AltLabelStructDecl tokenTypeDecls

//AltLabelStructDecl tokenListDecls

//AltLabelStructDecl ruleContextDecls

//AltLabelStructDecl ruleContextListDecls

//AltLabelStructDecl attributeDecls

//getRuleContext
func (s *DeclAnnoContext) GetRuleContext() antlr.RuleContext {
	return s
}

//getters
func (s *DeclAnnoContext) DOWN() antlr.TerminalNode {
	return s.GetToken(AdlWoDOWN, 0)
}

func (s *DeclAnnoContext) JsonVal() IJsonValContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*JsonValContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IJsonValContext)
}

func (s *DeclAnnoContext) UP() antlr.TerminalNode {
	return s.GetToken(AdlWoUP, 0)
}

func (s *DeclAnnoContext) DeclAnno() antlr.TerminalNode {
	return s.GetToken(AdlWoDeclAnno, 0)
}

func (s *DeclAnnoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DeclAnnoEntryListener); ok {
		listenerT.EnterDeclAnno(s)
	}
}

func (s *DeclAnnoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DeclAnnoExitListener); ok {
		listenerT.ExitDeclAnno(s)
	}
}

func (s *DeclAnnoContext) VisitFunc(hdls antlr.ParserTreeVisitorHandlers, args ...interface{}) (result interface{}) {
	h := hdls.(*AdlWoHandlers)
	if h.EnterEveryRule != nil {
		h.EnterEveryRule(s)
	}
	if h.DeclAnno != nil {
		h.DeclAnno(s, h, args...)
	} else {
		s.VisitChildrenFunc(h, args...)
	}
	if h.ExitEveryRule != nil {
		h.ExitEveryRule(s)
	}
	return
}

func (s *DeclAnnoContext) Visit(delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	switch t := delegate.(type) {
	case DeclAnnoContextVisitor:
		return t.VisitDeclAnno(s, delegate, args...)
	default:
		return delegate.VisitChildren(s, delegate, args...)
	}
}

//END AltLabelStructDecl

//Begin AltLabelStructDecl

type StructContext struct {
	*TldContext
	//TokenDecl
	tok antlr.Token
}

func NewStructContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StructContext {
	var p = new(StructContext)

	p.TldContext = NewEmptyTldContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TldContext))

	return p
}

type IStructContext interface {
	//Current rule
	ITldContext
	// start internal
	//Gets for raw elements
	//  ruleGetterDecl
	//  ruleListGetterDecl
	AllAnnotation() []IAnnotationContext
	AllNameBody() []INameBodyContext
	//  ruleListIndexedGetterDecl
	Annotation(i int) IAnnotationContext
	NameBody(i int) INameBodyContext

	//  tokenGetterDecl
	Struct() antlr.TerminalNode
	DOWN() antlr.TerminalNode
	UP() antlr.TerminalNode
	TypeParam() antlr.TerminalNode
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

func (*StructContext) IsStructContext() {}

//AltLabelStructDecl tokenDecls
func (s *StructContext) GetTok() antlr.Token  { return s.tok }
func (s *StructContext) SetTok(v antlr.Token) { s.tok = v }

//AltLabelStructDecl tokenTypeDecls

//AltLabelStructDecl tokenListDecls

//AltLabelStructDecl ruleContextDecls

//AltLabelStructDecl ruleContextListDecls

//AltLabelStructDecl attributeDecls

//getRuleContext
func (s *StructContext) GetRuleContext() antlr.RuleContext {
	return s
}

//getters
func (s *StructContext) Struct() antlr.TerminalNode {
	return s.GetToken(AdlWoStruct, 0)
}

func (s *StructContext) DOWN() antlr.TerminalNode {
	return s.GetToken(AdlWoDOWN, 0)
}

func (s *StructContext) UP() antlr.TerminalNode {
	return s.GetToken(AdlWoUP, 0)
}

func (s *StructContext) AllAnnotation() []IAnnotationContext {
	//ContextRuleListGetterDecl
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*AnnotationContext)(nil)).Elem())
	var tst = make([]IAnnotationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAnnotationContext)
		}
	}

	return tst
}

func (s *StructContext) Annotation(i int) IAnnotationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*AnnotationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAnnotationContext)
}

func (s *StructContext) TypeParam() antlr.TerminalNode {
	return s.GetToken(AdlWoTypeParam, 0)
}

func (s *StructContext) AllNameBody() []INameBodyContext {
	//ContextRuleListGetterDecl
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*NameBodyContext)(nil)).Elem())
	var tst = make([]INameBodyContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INameBodyContext)
		}
	}

	return tst
}

func (s *StructContext) NameBody(i int) INameBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*NameBodyContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INameBodyContext)
}

func (s *StructContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StructEntryListener); ok {
		listenerT.EnterStruct(s)
	}
}

func (s *StructContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StructExitListener); ok {
		listenerT.ExitStruct(s)
	}
}

func (s *StructContext) VisitFunc(hdls antlr.ParserTreeVisitorHandlers, args ...interface{}) (result interface{}) {
	h := hdls.(*AdlWoHandlers)
	if h.EnterEveryRule != nil {
		h.EnterEveryRule(s)
	}
	if h.Struct != nil {
		h.Struct(s, h, args...)
	} else {
		s.VisitChildrenFunc(h, args...)
	}
	if h.ExitEveryRule != nil {
		h.ExitEveryRule(s)
	}
	return
}

func (s *StructContext) Visit(delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	switch t := delegate.(type) {
	case StructContextVisitor:
		return t.VisitStruct(s, delegate, args...)
	default:
		return delegate.VisitChildren(s, delegate, args...)
	}
}

//END AltLabelStructDecl

func (p *AdlWo) Tld() (localctx ITldContext) {
	localctx = NewTldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, AdlWoRULE_tld)
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

	p.SetState(161)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AdlWoStruct:
		localctx = NewStructContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(60)
			var _m = p.Match(AdlWoStruct)
			localctx.(*StructContext).tok = _m

		}
		p.SetState(78)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == AdlWoDOWN {
			{
				p.SetState(61)
				p.Match(AdlWoDOWN)
			}
			p.SetState(65)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for ((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(AdlWoAnnotation-32))|(1<<(AdlWoAnnotationNotScoped-32))|(1<<(AdlWoAnnotationScoped-32)))) != 0 {
				{
					p.SetState(62)
					p.Annotation()
				}

				p.SetState(67)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			p.SetState(69)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == AdlWoTypeParam {
				{
					p.SetState(68)
					p.Match(AdlWoTypeParam)
				}

			}
			p.SetState(74)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == AdlWoField {
				{
					p.SetState(71)
					p.NameBody()
				}

				p.SetState(76)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(77)
				p.Match(AdlWoUP)
			}

		}

	case AdlWoUnion:
		localctx = NewUnionContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(80)
			var _m = p.Match(AdlWoUnion)
			localctx.(*UnionContext).tok = _m

		}
		p.SetState(98)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == AdlWoDOWN {
			{
				p.SetState(81)
				p.Match(AdlWoDOWN)
			}
			p.SetState(85)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for ((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(AdlWoAnnotation-32))|(1<<(AdlWoAnnotationNotScoped-32))|(1<<(AdlWoAnnotationScoped-32)))) != 0 {
				{
					p.SetState(82)
					p.Annotation()
				}

				p.SetState(87)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			p.SetState(89)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == AdlWoTypeParam {
				{
					p.SetState(88)
					p.Match(AdlWoTypeParam)
				}

			}
			p.SetState(94)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == AdlWoField {
				{
					p.SetState(91)
					p.NameBody()
				}

				p.SetState(96)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(97)
				p.Match(AdlWoUP)
			}

		}

	case AdlWoType:
		localctx = NewTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(100)
			var _m = p.Match(AdlWoType)
			localctx.(*TypeContext).tok = _m

		}
		p.SetState(121)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == AdlWoDOWN {
			{
				p.SetState(101)
				p.Match(AdlWoDOWN)
			}
			p.SetState(105)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for ((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(AdlWoAnnotation-32))|(1<<(AdlWoAnnotationNotScoped-32))|(1<<(AdlWoAnnotationScoped-32)))) != 0 {
				{
					p.SetState(102)
					p.Annotation()
				}

				p.SetState(107)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			p.SetState(109)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == AdlWoTypeParam {
				{
					p.SetState(108)
					p.Match(AdlWoTypeParam)
				}

			}
			p.SetState(112)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == AdlWoTypeExprPrimOrParam || _la == AdlWoTypeExprTypeExpr {
				{
					p.SetState(111)
					p.TypeExpr_()
				}

			}
			p.SetState(117)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for ((_la-45)&-(0x1f+1)) == 0 && ((1<<uint((_la-45)))&((1<<(AdlWoJsonStr-45))|(1<<(AdlWoJsonBool-45))|(1<<(AdlWoJsonNull-45))|(1<<(AdlWoJsonInt-45))|(1<<(AdlWoJsonFloat-45))|(1<<(AdlWoJsonArray-45))|(1<<(AdlWoJsonObj-45)))) != 0 {
				{
					p.SetState(114)
					p.JsonVal()
				}

				p.SetState(119)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(120)
				p.Match(AdlWoUP)
			}

		}

	case AdlWoNewtype:
		localctx = NewNewtypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(123)
			var _m = p.Match(AdlWoNewtype)
			localctx.(*NewtypeContext).tok = _m

		}
		p.SetState(144)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == AdlWoDOWN {
			{
				p.SetState(124)
				p.Match(AdlWoDOWN)
			}
			p.SetState(128)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for ((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(AdlWoAnnotation-32))|(1<<(AdlWoAnnotationNotScoped-32))|(1<<(AdlWoAnnotationScoped-32)))) != 0 {
				{
					p.SetState(125)
					p.Annotation()
				}

				p.SetState(130)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			p.SetState(132)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == AdlWoTypeParam {
				{
					p.SetState(131)
					p.Match(AdlWoTypeParam)
				}

			}
			p.SetState(135)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == AdlWoTypeExprPrimOrParam || _la == AdlWoTypeExprTypeExpr {
				{
					p.SetState(134)
					p.TypeExpr_()
				}

			}
			p.SetState(140)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for ((_la-45)&-(0x1f+1)) == 0 && ((1<<uint((_la-45)))&((1<<(AdlWoJsonStr-45))|(1<<(AdlWoJsonBool-45))|(1<<(AdlWoJsonNull-45))|(1<<(AdlWoJsonInt-45))|(1<<(AdlWoJsonFloat-45))|(1<<(AdlWoJsonArray-45))|(1<<(AdlWoJsonObj-45)))) != 0 {
				{
					p.SetState(137)
					p.JsonVal()
				}

				p.SetState(142)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(143)
				p.Match(AdlWoUP)
			}

		}

	case AdlWoModuleAnno:
		localctx = NewModAnnoContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(146)
			var _m = p.Match(AdlWoModuleAnno)
			localctx.(*ModAnnoContext).tok = _m

		}
		{
			p.SetState(147)
			p.Match(AdlWoDOWN)
		}
		{
			p.SetState(148)
			p.JsonVal()
		}
		{
			p.SetState(149)
			p.Match(AdlWoUP)
		}

	case AdlWoDeclAnno:
		localctx = NewDeclAnnoContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(151)
			var _m = p.Match(AdlWoDeclAnno)
			localctx.(*DeclAnnoContext).tok = _m

		}
		{
			p.SetState(152)
			p.Match(AdlWoDOWN)
		}
		{
			p.SetState(153)
			p.JsonVal()
		}
		{
			p.SetState(154)
			p.Match(AdlWoUP)
		}

	case AdlWoFieldAnno:
		localctx = NewFieldAnnoContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(156)
			var _m = p.Match(AdlWoFieldAnno)
			localctx.(*FieldAnnoContext).tok = _m

		}
		{
			p.SetState(157)
			p.Match(AdlWoDOWN)
		}
		{
			p.SetState(158)
			p.JsonVal()
		}
		{
			p.SetState(159)
			p.Match(AdlWoUP)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

type INameBodyContext interface {
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

	// IsNameBodyContext differentiates from other interfaces.
	//copyStruct,GetRuleContext and ToStringTree  from embedded

	//<if(dispatchMethods)>
	//<dispatchMethods; separator="\n\n">
	//<endif>

	//<if(extensionMembers)>
	//<extensionMembers; separator="\n\n">
	//<endif>
}

type NameBodyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNameBodyContext() *NameBodyContext {
	var p = new(NameBodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AdlWoRULE_nameBody
	return p
}

func (*NameBodyContext) IsNameBodyContext() {}

func NewNameBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NameBodyContext {
	var p = new(NameBodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AdlWoRULE_nameBody

	return p
}

func (s *NameBodyContext) GetParser() antlr.Parser { return s.parser }

//StructDecl tokenDecls

//StructDecl tokenTypeDecls

//StructDecl tokenListDecls

//StructDecl ruleContextDecls

//StructDecl ruleContextListDecls

//StructDecl attributeDecls

// Getters

//provideCopyFrom
func (s *NameBodyContext) CopyFrom(ctx *NameBodyContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *NameBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NameBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

//dispatchMethods
//extensionMembers

//Begin AltLabelStructDecl

type FieldContext struct {
	*NameBodyContext
	//TokenDecl
	tok antlr.Token
}

func NewFieldContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FieldContext {
	var p = new(FieldContext)

	p.NameBodyContext = NewEmptyNameBodyContext()
	p.parser = parser
	p.CopyFrom(ctx.(*NameBodyContext))

	return p
}

type IFieldContext interface {
	//Current rule
	INameBodyContext
	// start internal
	//Gets for raw elements
	//  ruleGetterDecl
	TypeExpr_() ITypeExpr_Context
	JsonVal() IJsonValContext
	//  ruleListGetterDecl
	AllAnnotation() []IAnnotationContext
	//  ruleListIndexedGetterDecl
	Annotation(i int) IAnnotationContext

	//  tokenGetterDecl
	Field() antlr.TerminalNode
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

func (*FieldContext) IsFieldContext() {}

//AltLabelStructDecl tokenDecls
func (s *FieldContext) GetTok() antlr.Token  { return s.tok }
func (s *FieldContext) SetTok(v antlr.Token) { s.tok = v }

//AltLabelStructDecl tokenTypeDecls

//AltLabelStructDecl tokenListDecls

//AltLabelStructDecl ruleContextDecls

//AltLabelStructDecl ruleContextListDecls

//AltLabelStructDecl attributeDecls

//getRuleContext
func (s *FieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

//getters
func (s *FieldContext) Field() antlr.TerminalNode {
	return s.GetToken(AdlWoField, 0)
}

func (s *FieldContext) DOWN() antlr.TerminalNode {
	return s.GetToken(AdlWoDOWN, 0)
}

func (s *FieldContext) UP() antlr.TerminalNode {
	return s.GetToken(AdlWoUP, 0)
}

func (s *FieldContext) AllAnnotation() []IAnnotationContext {
	//ContextRuleListGetterDecl
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*AnnotationContext)(nil)).Elem())
	var tst = make([]IAnnotationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAnnotationContext)
		}
	}

	return tst
}

func (s *FieldContext) Annotation(i int) IAnnotationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*AnnotationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAnnotationContext)
}

func (s *FieldContext) TypeExpr_() ITypeExpr_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*TypeExpr_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeExpr_Context)
}

func (s *FieldContext) JsonVal() IJsonValContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*JsonValContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IJsonValContext)
}

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
	h := hdls.(*AdlWoHandlers)
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

//END AltLabelStructDecl

func (p *AdlWo) NameBody() (localctx INameBodyContext) {
	localctx = NewNameBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, AdlWoRULE_nameBody)
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

	localctx = NewFieldContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(163)
		var _m = p.Match(AdlWoField)
		localctx.(*FieldContext).tok = _m

	}
	p.SetState(178)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == AdlWoDOWN {
		{
			p.SetState(164)
			p.Match(AdlWoDOWN)
		}
		p.SetState(168)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(AdlWoAnnotation-32))|(1<<(AdlWoAnnotationNotScoped-32))|(1<<(AdlWoAnnotationScoped-32)))) != 0 {
			{
				p.SetState(165)
				p.Annotation()
			}

			p.SetState(170)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		p.SetState(172)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == AdlWoTypeExprPrimOrParam || _la == AdlWoTypeExprTypeExpr {
			{
				p.SetState(171)
				p.TypeExpr_()
			}

		}
		p.SetState(175)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la-45)&-(0x1f+1)) == 0 && ((1<<uint((_la-45)))&((1<<(AdlWoJsonStr-45))|(1<<(AdlWoJsonBool-45))|(1<<(AdlWoJsonNull-45))|(1<<(AdlWoJsonInt-45))|(1<<(AdlWoJsonFloat-45))|(1<<(AdlWoJsonArray-45))|(1<<(AdlWoJsonObj-45)))) != 0 {
			{
				p.SetState(174)
				p.JsonVal()
			}

		}
		{
			p.SetState(177)
			p.Match(AdlWoUP)
		}

	}

	return localctx
}

type IAnnotationContext interface {
	antlr.ParserRuleContext
	// GetParser returns the parser.
	GetParser() antlr.Parser

	// start internal
	//  ruleGetterDecl
	JsonVal() IJsonValContext
	//  ruleListGetterDecl
	//  ruleListIndexedGetterDecl
	//  ruleContextDecls
	//  ruleContextListDecls
	// end internal
	//Gets for labeled elements
	//tokenDecls
	GetTok() antlr.Token

	//tokenTypeDecls
	//tokenListDecls
	//attributeDecls
	//tokenGetterDecl
	Annotation() antlr.TerminalNode
	DOWN() antlr.TerminalNode
	UP() antlr.TerminalNode
	AnnotationNotScoped() antlr.TerminalNode
	AnnotationScoped() antlr.TerminalNode
	//tokenListGetterDecl
	//tokenListIndexedGetterDecl

	// IsAnnotationContext differentiates from other interfaces.
	//copyStruct,GetRuleContext and ToStringTree  from embedded

	//<if(dispatchMethods)>
	//<dispatchMethods; separator="\n\n">
	//<endif>

	//<if(extensionMembers)>
	//<extensionMembers; separator="\n\n">
	//<endif>
}

type AnnotationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	//TokenDecl
	tok antlr.Token
}

func NewEmptyAnnotationContext() *AnnotationContext {
	var p = new(AnnotationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AdlWoRULE_annotation
	return p
}

func (*AnnotationContext) IsAnnotationContext() {}

func NewAnnotationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AnnotationContext {
	var p = new(AnnotationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AdlWoRULE_annotation

	return p
}

func (s *AnnotationContext) GetParser() antlr.Parser { return s.parser }

//StructDecl tokenDecls
func (s *AnnotationContext) GetTok() antlr.Token  { return s.tok }
func (s *AnnotationContext) SetTok(v antlr.Token) { s.tok = v }

//StructDecl tokenTypeDecls

//StructDecl tokenListDecls

//StructDecl ruleContextDecls

//StructDecl ruleContextListDecls

//StructDecl attributeDecls

// Getters
func (s *AnnotationContext) Annotation() antlr.TerminalNode {
	return s.GetToken(AdlWoAnnotation, 0)
}

func (s *AnnotationContext) DOWN() antlr.TerminalNode {
	return s.GetToken(AdlWoDOWN, 0)
}

func (s *AnnotationContext) JsonVal() IJsonValContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*JsonValContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IJsonValContext)
}

func (s *AnnotationContext) UP() antlr.TerminalNode {
	return s.GetToken(AdlWoUP, 0)
}

func (s *AnnotationContext) AnnotationNotScoped() antlr.TerminalNode {
	return s.GetToken(AdlWoAnnotationNotScoped, 0)
}

func (s *AnnotationContext) AnnotationScoped() antlr.TerminalNode {
	return s.GetToken(AdlWoAnnotationScoped, 0)
}

//provideCopyFrom
func (s *AnnotationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AnnotationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

//dispatchMethods
func (s *AnnotationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnnotationEntryListener); ok {
		listenerT.EnterAnnotation(s)
	}
}

func (s *AnnotationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AnnotationExitListener); ok {
		listenerT.ExitAnnotation(s)
	}
}

func (s *AnnotationContext) VisitFunc(hdls antlr.ParserTreeVisitorHandlers, args ...interface{}) (result interface{}) {
	h := hdls.(*AdlWoHandlers)
	if h.EnterEveryRule != nil {
		h.EnterEveryRule(s)
	}
	if h.Annotation != nil {
		h.Annotation(s, h, args...)
	} else {
		s.VisitChildrenFunc(h, args...)
	}
	if h.ExitEveryRule != nil {
		h.ExitEveryRule(s)
	}
	return
}

func (s *AnnotationContext) Visit(delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	switch t := delegate.(type) {
	case AnnotationContextVisitor:
		return t.VisitAnnotation(s, delegate, args...)
	default:
		return delegate.VisitChildren(s, delegate, args...)
	}
}

//extensionMembers

func (p *AdlWo) Annotation() (localctx IAnnotationContext) {
	localctx = NewAnnotationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, AdlWoRULE_annotation)
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

	p.SetState(201)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AdlWoAnnotation:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(180)
			var _m = p.Match(AdlWoAnnotation)
			localctx.(*AnnotationContext).tok = _m

		}
		p.SetState(185)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == AdlWoDOWN {
			{
				p.SetState(181)
				p.Match(AdlWoDOWN)
			}
			{
				p.SetState(182)
				p.JsonVal()
			}
			{
				p.SetState(183)
				p.Match(AdlWoUP)
			}

		}

	case AdlWoAnnotationNotScoped:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(187)
			var _m = p.Match(AdlWoAnnotationNotScoped)
			localctx.(*AnnotationContext).tok = _m

		}
		p.SetState(192)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == AdlWoDOWN {
			{
				p.SetState(188)
				p.Match(AdlWoDOWN)
			}
			{
				p.SetState(189)
				p.JsonVal()
			}
			{
				p.SetState(190)
				p.Match(AdlWoUP)
			}

		}

	case AdlWoAnnotationScoped:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(194)
			var _m = p.Match(AdlWoAnnotationScoped)
			localctx.(*AnnotationContext).tok = _m

		}
		p.SetState(199)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == AdlWoDOWN {
			{
				p.SetState(195)
				p.Match(AdlWoDOWN)
			}
			{
				p.SetState(196)
				p.JsonVal()
			}
			{
				p.SetState(197)
				p.Match(AdlWoUP)
			}

		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

type ITypeExpr_Context interface {
	antlr.ParserRuleContext
	// GetParser returns the parser.
	GetParser() antlr.Parser

	// start internal
	//  ruleGetterDecl
	//  ruleListGetterDecl
	AllTypeExprElem_() []ITypeExprElem_Context
	//  ruleListIndexedGetterDecl
	TypeExprElem_(i int) ITypeExprElem_Context
	//  ruleContextDecls
	//  ruleContextListDecls
	// end internal
	//Gets for labeled elements
	//tokenDecls
	GetTok() antlr.Token

	//tokenTypeDecls
	//tokenListDecls
	//attributeDecls
	//tokenGetterDecl
	TypeExprPrimOrParam() antlr.TerminalNode
	TypeExprTypeExpr() antlr.TerminalNode
	DOWN() antlr.TerminalNode
	UP() antlr.TerminalNode
	//tokenListGetterDecl
	//tokenListIndexedGetterDecl

	// IsTypeExpr_Context differentiates from other interfaces.
	//copyStruct,GetRuleContext and ToStringTree  from embedded

	//<if(dispatchMethods)>
	//<dispatchMethods; separator="\n\n">
	//<endif>

	//<if(extensionMembers)>
	//<extensionMembers; separator="\n\n">
	//<endif>
}

type TypeExpr_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	//TokenDecl
	tok antlr.Token
}

func NewEmptyTypeExpr_Context() *TypeExpr_Context {
	var p = new(TypeExpr_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AdlWoRULE_typeExpr_
	return p
}

func (*TypeExpr_Context) IsTypeExpr_Context() {}

func NewTypeExpr_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeExpr_Context {
	var p = new(TypeExpr_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AdlWoRULE_typeExpr_

	return p
}

func (s *TypeExpr_Context) GetParser() antlr.Parser { return s.parser }

//StructDecl tokenDecls
func (s *TypeExpr_Context) GetTok() antlr.Token  { return s.tok }
func (s *TypeExpr_Context) SetTok(v antlr.Token) { s.tok = v }

//StructDecl tokenTypeDecls

//StructDecl tokenListDecls

//StructDecl ruleContextDecls

//StructDecl ruleContextListDecls

//StructDecl attributeDecls

// Getters
func (s *TypeExpr_Context) TypeExprPrimOrParam() antlr.TerminalNode {
	return s.GetToken(AdlWoTypeExprPrimOrParam, 0)
}

func (s *TypeExpr_Context) TypeExprTypeExpr() antlr.TerminalNode {
	return s.GetToken(AdlWoTypeExprTypeExpr, 0)
}

func (s *TypeExpr_Context) DOWN() antlr.TerminalNode {
	return s.GetToken(AdlWoDOWN, 0)
}

func (s *TypeExpr_Context) UP() antlr.TerminalNode {
	return s.GetToken(AdlWoUP, 0)
}

func (s *TypeExpr_Context) AllTypeExprElem_() []ITypeExprElem_Context {
	//ContextRuleListGetterDecl
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*TypeExprElem_Context)(nil)).Elem())
	var tst = make([]ITypeExprElem_Context, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITypeExprElem_Context)
		}
	}

	return tst
}

func (s *TypeExpr_Context) TypeExprElem_(i int) ITypeExprElem_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*TypeExprElem_Context)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITypeExprElem_Context)
}

//provideCopyFrom
func (s *TypeExpr_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeExpr_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

//dispatchMethods
func (s *TypeExpr_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TypeExpr_EntryListener); ok {
		listenerT.EnterTypeExpr_(s)
	}
}

func (s *TypeExpr_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TypeExpr_ExitListener); ok {
		listenerT.ExitTypeExpr_(s)
	}
}

func (s *TypeExpr_Context) VisitFunc(hdls antlr.ParserTreeVisitorHandlers, args ...interface{}) (result interface{}) {
	h := hdls.(*AdlWoHandlers)
	if h.EnterEveryRule != nil {
		h.EnterEveryRule(s)
	}
	if h.TypeExpr_ != nil {
		h.TypeExpr_(s, h, args...)
	} else {
		s.VisitChildrenFunc(h, args...)
	}
	if h.ExitEveryRule != nil {
		h.ExitEveryRule(s)
	}
	return
}

func (s *TypeExpr_Context) Visit(delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	switch t := delegate.(type) {
	case TypeExpr_ContextVisitor:
		return t.VisitTypeExpr_(s, delegate, args...)
	default:
		return delegate.VisitChildren(s, delegate, args...)
	}
}

//extensionMembers

func (p *AdlWo) TypeExpr_() (localctx ITypeExpr_Context) {
	localctx = NewTypeExpr_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, AdlWoRULE_typeExpr_)
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

	p.SetState(213)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AdlWoTypeExprPrimOrParam:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(203)
			var _m = p.Match(AdlWoTypeExprPrimOrParam)
			localctx.(*TypeExpr_Context).tok = _m

		}

	case AdlWoTypeExprTypeExpr:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(204)
			var _m = p.Match(AdlWoTypeExprTypeExpr)
			localctx.(*TypeExpr_Context).tok = _m

		}

		{
			p.SetState(205)
			p.Match(AdlWoDOWN)
		}
		p.SetState(207)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == AdlWoTypeExprElem {
			{
				p.SetState(206)
				p.TypeExprElem_()
			}

			p.SetState(209)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(211)
			p.Match(AdlWoUP)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

type ITypeExprElem_Context interface {
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

	// IsTypeExprElem_Context differentiates from other interfaces.
	//copyStruct,GetRuleContext and ToStringTree  from embedded

	//<if(dispatchMethods)>
	//<dispatchMethods; separator="\n\n">
	//<endif>

	//<if(extensionMembers)>
	//<extensionMembers; separator="\n\n">
	//<endif>
}

type TypeExprElem_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeExprElem_Context() *TypeExprElem_Context {
	var p = new(TypeExprElem_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AdlWoRULE_typeExprElem_
	return p
}

func (*TypeExprElem_Context) IsTypeExprElem_Context() {}

func NewTypeExprElem_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeExprElem_Context {
	var p = new(TypeExprElem_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AdlWoRULE_typeExprElem_

	return p
}

func (s *TypeExprElem_Context) GetParser() antlr.Parser { return s.parser }

//StructDecl tokenDecls

//StructDecl tokenTypeDecls

//StructDecl tokenListDecls

//StructDecl ruleContextDecls

//StructDecl ruleContextListDecls

//StructDecl attributeDecls

// Getters

//provideCopyFrom
func (s *TypeExprElem_Context) CopyFrom(ctx *TypeExprElem_Context) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *TypeExprElem_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeExprElem_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

//dispatchMethods
//extensionMembers

//Begin AltLabelStructDecl

type TypeParamsContext struct {
	*TypeExprElem_Context
	//TokenDecl
	tok antlr.Token
}

func NewTypeParamsContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypeParamsContext {
	var p = new(TypeParamsContext)

	p.TypeExprElem_Context = NewEmptyTypeExprElem_Context()
	p.parser = parser
	p.CopyFrom(ctx.(*TypeExprElem_Context))

	return p
}

type ITypeParamsContext interface {
	//Current rule
	ITypeExprElem_Context
	// start internal
	//Gets for raw elements
	//  ruleGetterDecl
	//  ruleListGetterDecl
	AllTypeExprElem_() []ITypeExprElem_Context
	//  ruleListIndexedGetterDecl
	TypeExprElem_(i int) ITypeExprElem_Context

	//  tokenGetterDecl
	TypeExprElem() antlr.TerminalNode
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

func (*TypeParamsContext) IsTypeParamsContext() {}

//AltLabelStructDecl tokenDecls
func (s *TypeParamsContext) GetTok() antlr.Token  { return s.tok }
func (s *TypeParamsContext) SetTok(v antlr.Token) { s.tok = v }

//AltLabelStructDecl tokenTypeDecls

//AltLabelStructDecl tokenListDecls

//AltLabelStructDecl ruleContextDecls

//AltLabelStructDecl ruleContextListDecls

//AltLabelStructDecl attributeDecls

//getRuleContext
func (s *TypeParamsContext) GetRuleContext() antlr.RuleContext {
	return s
}

//getters
func (s *TypeParamsContext) TypeExprElem() antlr.TerminalNode {
	return s.GetToken(AdlWoTypeExprElem, 0)
}

func (s *TypeParamsContext) DOWN() antlr.TerminalNode {
	return s.GetToken(AdlWoDOWN, 0)
}

func (s *TypeParamsContext) UP() antlr.TerminalNode {
	return s.GetToken(AdlWoUP, 0)
}

func (s *TypeParamsContext) AllTypeExprElem_() []ITypeExprElem_Context {
	//ContextRuleListGetterDecl
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*TypeExprElem_Context)(nil)).Elem())
	var tst = make([]ITypeExprElem_Context, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITypeExprElem_Context)
		}
	}

	return tst
}

func (s *TypeParamsContext) TypeExprElem_(i int) ITypeExprElem_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*TypeExprElem_Context)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITypeExprElem_Context)
}

func (s *TypeParamsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TypeParamsEntryListener); ok {
		listenerT.EnterTypeParams(s)
	}
}

func (s *TypeParamsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TypeParamsExitListener); ok {
		listenerT.ExitTypeParams(s)
	}
}

func (s *TypeParamsContext) VisitFunc(hdls antlr.ParserTreeVisitorHandlers, args ...interface{}) (result interface{}) {
	h := hdls.(*AdlWoHandlers)
	if h.EnterEveryRule != nil {
		h.EnterEveryRule(s)
	}
	if h.TypeParams != nil {
		h.TypeParams(s, h, args...)
	} else {
		s.VisitChildrenFunc(h, args...)
	}
	if h.ExitEveryRule != nil {
		h.ExitEveryRule(s)
	}
	return
}

func (s *TypeParamsContext) Visit(delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	switch t := delegate.(type) {
	case TypeParamsContextVisitor:
		return t.VisitTypeParams(s, delegate, args...)
	default:
		return delegate.VisitChildren(s, delegate, args...)
	}
}

//END AltLabelStructDecl

func (p *AdlWo) TypeExprElem_() (localctx ITypeExprElem_Context) {
	localctx = NewTypeExprElem_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, AdlWoRULE_typeExprElem_)
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

	localctx = NewTypeParamsContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(215)
		var _m = p.Match(AdlWoTypeExprElem)
		localctx.(*TypeParamsContext).tok = _m

	}
	p.SetState(224)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == AdlWoDOWN {
		{
			p.SetState(216)
			p.Match(AdlWoDOWN)
		}
		p.SetState(218)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == AdlWoTypeExprElem {
			{
				p.SetState(217)
				p.TypeExprElem_()
			}

			p.SetState(220)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(222)
			p.Match(AdlWoUP)
		}

	}

	return localctx
}

type IJsonValContext interface {
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

	// IsJsonValContext differentiates from other interfaces.
	//copyStruct,GetRuleContext and ToStringTree  from embedded

	//<if(dispatchMethods)>
	//<dispatchMethods; separator="\n\n">
	//<endif>

	//<if(extensionMembers)>
	//<extensionMembers; separator="\n\n">
	//<endif>
}

type JsonValContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJsonValContext() *JsonValContext {
	var p = new(JsonValContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AdlWoRULE_jsonVal
	return p
}

func (*JsonValContext) IsJsonValContext() {}

func NewJsonValContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *JsonValContext {
	var p = new(JsonValContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AdlWoRULE_jsonVal

	return p
}

func (s *JsonValContext) GetParser() antlr.Parser { return s.parser }

//StructDecl tokenDecls

//StructDecl tokenTypeDecls

//StructDecl tokenListDecls

//StructDecl ruleContextDecls

//StructDecl ruleContextListDecls

//StructDecl attributeDecls

// Getters

//provideCopyFrom
func (s *JsonValContext) CopyFrom(ctx *JsonValContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *JsonValContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JsonValContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

//dispatchMethods
//extensionMembers

//Begin AltLabelStructDecl

type JsonStrContext struct {
	*JsonValContext
	//TokenDecl
	tok antlr.Token
}

func NewJsonStrContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *JsonStrContext {
	var p = new(JsonStrContext)

	p.JsonValContext = NewEmptyJsonValContext()
	p.parser = parser
	p.CopyFrom(ctx.(*JsonValContext))

	return p
}

type IJsonStrContext interface {
	//Current rule
	IJsonValContext
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
	return s.GetToken(AdlWoJsonStr, 0)
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
	h := hdls.(*AdlWoHandlers)
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
	*JsonValContext
	//TokenDecl
	tok antlr.Token
}

func NewJsonArrayContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *JsonArrayContext {
	var p = new(JsonArrayContext)

	p.JsonValContext = NewEmptyJsonValContext()
	p.parser = parser
	p.CopyFrom(ctx.(*JsonValContext))

	return p
}

type IJsonArrayContext interface {
	//Current rule
	IJsonValContext
	// start internal
	//Gets for raw elements
	//  ruleGetterDecl
	//  ruleListGetterDecl
	AllJsonVal() []IJsonValContext
	//  ruleListIndexedGetterDecl
	JsonVal(i int) IJsonValContext

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
	return s.GetToken(AdlWoJsonArray, 0)
}

func (s *JsonArrayContext) DOWN() antlr.TerminalNode {
	return s.GetToken(AdlWoDOWN, 0)
}

func (s *JsonArrayContext) UP() antlr.TerminalNode {
	return s.GetToken(AdlWoUP, 0)
}

func (s *JsonArrayContext) AllJsonVal() []IJsonValContext {
	//ContextRuleListGetterDecl
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*JsonValContext)(nil)).Elem())
	var tst = make([]IJsonValContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IJsonValContext)
		}
	}

	return tst
}

func (s *JsonArrayContext) JsonVal(i int) IJsonValContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*JsonValContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IJsonValContext)
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
	h := hdls.(*AdlWoHandlers)
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
	*JsonValContext
	//TokenDecl
	tok antlr.Token
}

func NewJsonFloatContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *JsonFloatContext {
	var p = new(JsonFloatContext)

	p.JsonValContext = NewEmptyJsonValContext()
	p.parser = parser
	p.CopyFrom(ctx.(*JsonValContext))

	return p
}

type IJsonFloatContext interface {
	//Current rule
	IJsonValContext
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
	return s.GetToken(AdlWoJsonFloat, 0)
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
	h := hdls.(*AdlWoHandlers)
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
	*JsonValContext
	//TokenDecl
	tok antlr.Token
}

func NewJsonObjContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *JsonObjContext {
	var p = new(JsonObjContext)

	p.JsonValContext = NewEmptyJsonValContext()
	p.parser = parser
	p.CopyFrom(ctx.(*JsonValContext))

	return p
}

type IJsonObjContext interface {
	//Current rule
	IJsonValContext
	// start internal
	//Gets for raw elements
	//  ruleGetterDecl
	//  ruleListGetterDecl
	AllJsonVal() []IJsonValContext
	//  ruleListIndexedGetterDecl
	JsonVal(i int) IJsonValContext

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
	return s.GetToken(AdlWoJsonObj, 0)
}

func (s *JsonObjContext) DOWN() antlr.TerminalNode {
	return s.GetToken(AdlWoDOWN, 0)
}

func (s *JsonObjContext) UP() antlr.TerminalNode {
	return s.GetToken(AdlWoUP, 0)
}

func (s *JsonObjContext) AllJsonVal() []IJsonValContext {
	//ContextRuleListGetterDecl
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*JsonValContext)(nil)).Elem())
	var tst = make([]IJsonValContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IJsonValContext)
		}
	}

	return tst
}

func (s *JsonObjContext) JsonVal(i int) IJsonValContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*JsonValContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IJsonValContext)
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
	h := hdls.(*AdlWoHandlers)
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
	*JsonValContext
	//TokenDecl
	tok antlr.Token
}

func NewJsonBoolContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *JsonBoolContext {
	var p = new(JsonBoolContext)

	p.JsonValContext = NewEmptyJsonValContext()
	p.parser = parser
	p.CopyFrom(ctx.(*JsonValContext))

	return p
}

type IJsonBoolContext interface {
	//Current rule
	IJsonValContext
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
	return s.GetToken(AdlWoJsonBool, 0)
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
	h := hdls.(*AdlWoHandlers)
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

type JsonIntContext struct {
	*JsonValContext
	//TokenDecl
	tok antlr.Token
}

func NewJsonIntContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *JsonIntContext {
	var p = new(JsonIntContext)

	p.JsonValContext = NewEmptyJsonValContext()
	p.parser = parser
	p.CopyFrom(ctx.(*JsonValContext))

	return p
}

type IJsonIntContext interface {
	//Current rule
	IJsonValContext
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
	return s.GetToken(AdlWoJsonInt, 0)
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
	h := hdls.(*AdlWoHandlers)
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
	*JsonValContext
	//TokenDecl
	tok antlr.Token
}

func NewJsonNullContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *JsonNullContext {
	var p = new(JsonNullContext)

	p.JsonValContext = NewEmptyJsonValContext()
	p.parser = parser
	p.CopyFrom(ctx.(*JsonValContext))

	return p
}

type IJsonNullContext interface {
	//Current rule
	IJsonValContext
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
	return s.GetToken(AdlWoJsonNull, 0)
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
	h := hdls.(*AdlWoHandlers)
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

func (p *AdlWo) JsonVal() (localctx IJsonValContext) {
	localctx = NewJsonValContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, AdlWoRULE_jsonVal)
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

	p.SetState(253)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AdlWoJsonStr:
		localctx = NewJsonStrContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(226)
			var _m = p.Match(AdlWoJsonStr)
			localctx.(*JsonStrContext).tok = _m

		}

	case AdlWoJsonBool:
		localctx = NewJsonBoolContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(227)
			var _m = p.Match(AdlWoJsonBool)
			localctx.(*JsonBoolContext).tok = _m

		}

	case AdlWoJsonNull:
		localctx = NewJsonNullContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(228)
			var _m = p.Match(AdlWoJsonNull)
			localctx.(*JsonNullContext).tok = _m

		}

	case AdlWoJsonInt:
		localctx = NewJsonIntContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(229)
			var _m = p.Match(AdlWoJsonInt)
			localctx.(*JsonIntContext).tok = _m

		}

	case AdlWoJsonFloat:
		localctx = NewJsonFloatContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(230)
			var _m = p.Match(AdlWoJsonFloat)
			localctx.(*JsonFloatContext).tok = _m

		}

	case AdlWoJsonArray:
		localctx = NewJsonArrayContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(231)
			var _m = p.Match(AdlWoJsonArray)
			localctx.(*JsonArrayContext).tok = _m

		}
		p.SetState(240)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == AdlWoDOWN {
			{
				p.SetState(232)
				p.Match(AdlWoDOWN)
			}
			p.SetState(234)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for ok := true; ok; ok = (((_la-45)&-(0x1f+1)) == 0 && ((1<<uint((_la-45)))&((1<<(AdlWoJsonStr-45))|(1<<(AdlWoJsonBool-45))|(1<<(AdlWoJsonNull-45))|(1<<(AdlWoJsonInt-45))|(1<<(AdlWoJsonFloat-45))|(1<<(AdlWoJsonArray-45))|(1<<(AdlWoJsonObj-45)))) != 0) {
				{
					p.SetState(233)
					p.JsonVal()
				}

				p.SetState(236)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(238)
				p.Match(AdlWoUP)
			}

		}

	case AdlWoJsonObj:
		localctx = NewJsonObjContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(242)
			var _m = p.Match(AdlWoJsonObj)
			localctx.(*JsonObjContext).tok = _m

		}
		p.SetState(251)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == AdlWoDOWN {
			{
				p.SetState(243)
				p.Match(AdlWoDOWN)
			}
			p.SetState(245)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for ok := true; ok; ok = (((_la-45)&-(0x1f+1)) == 0 && ((1<<uint((_la-45)))&((1<<(AdlWoJsonStr-45))|(1<<(AdlWoJsonBool-45))|(1<<(AdlWoJsonNull-45))|(1<<(AdlWoJsonInt-45))|(1<<(AdlWoJsonFloat-45))|(1<<(AdlWoJsonArray-45))|(1<<(AdlWoJsonObj-45)))) != 0) {
				{
					p.SetState(244)
					p.JsonVal()
				}

				p.SetState(247)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(249)
				p.Match(AdlWoUP)
			}

		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}
