// Generated from AdlP.g4 by ANTLR 4.7.

package adllp // AdlP
import (
	"fmt"
	"reflect"
	"strconv"
)
import "github.com/antlr/antlr4/runtime/Go/antlr"

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 52, 269,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 3, 2,
	3, 2, 3, 2, 3, 3, 7, 3, 29, 10, 3, 12, 3, 14, 3, 32, 11, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 7, 3, 38, 10, 3, 12, 3, 14, 3, 41, 11, 3, 3, 3, 3, 3, 7, 3,
	45, 10, 3, 12, 3, 14, 3, 48, 11, 3, 3, 3, 7, 3, 51, 10, 3, 12, 3, 14, 3,
	54, 11, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 7, 4, 63, 10, 4, 12,
	4, 14, 4, 66, 11, 4, 3, 4, 3, 4, 5, 4, 70, 10, 4, 3, 4, 3, 4, 3, 5, 3,
	5, 3, 5, 3, 5, 5, 5, 78, 10, 5, 3, 6, 7, 6, 81, 10, 6, 12, 6, 14, 6, 84,
	11, 6, 3, 6, 3, 6, 3, 6, 5, 6, 89, 10, 6, 3, 6, 3, 6, 7, 6, 93, 10, 6,
	12, 6, 14, 6, 96, 11, 6, 3, 6, 3, 6, 3, 6, 7, 6, 101, 10, 6, 12, 6, 14,
	6, 104, 11, 6, 3, 6, 3, 6, 3, 6, 5, 6, 109, 10, 6, 3, 6, 3, 6, 3, 6, 5,
	6, 114, 10, 6, 3, 6, 3, 6, 5, 6, 118, 10, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3,
	6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3,
	6, 3, 6, 3, 6, 3, 6, 5, 6, 140, 10, 6, 3, 7, 3, 7, 3, 7, 3, 7, 7, 7, 146,
	10, 7, 12, 7, 14, 7, 149, 11, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 7, 7, 156,
	10, 7, 12, 7, 14, 7, 159, 11, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 6,
	7, 167, 10, 7, 13, 7, 14, 7, 168, 3, 7, 3, 7, 5, 7, 173, 10, 7, 3, 8, 3,
	8, 3, 8, 3, 8, 5, 8, 179, 10, 8, 3, 8, 3, 8, 3, 8, 5, 8, 184, 10, 8, 7,
	8, 186, 10, 8, 12, 8, 14, 8, 189, 11, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9,
	3, 9, 7, 9, 197, 10, 9, 12, 9, 14, 9, 200, 11, 9, 3, 9, 3, 9, 3, 10, 3,
	10, 3, 10, 3, 10, 3, 10, 7, 10, 209, 10, 10, 12, 10, 14, 10, 212, 11, 10,
	3, 10, 3, 10, 5, 10, 216, 10, 10, 3, 11, 7, 11, 219, 10, 11, 12, 11, 14,
	11, 222, 11, 11, 3, 11, 3, 11, 5, 11, 226, 10, 11, 3, 11, 3, 11, 3, 11,
	5, 11, 231, 10, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3,
	12, 3, 12, 3, 12, 7, 12, 243, 10, 12, 12, 12, 14, 12, 246, 11, 12, 5, 12,
	248, 10, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3,
	12, 7, 12, 259, 10, 12, 12, 12, 14, 12, 262, 11, 12, 5, 12, 264, 10, 12,
	3, 12, 5, 12, 267, 10, 12, 3, 12, 2, 2, 13, 2, 4, 6, 8, 10, 12, 14, 16,
	18, 20, 22, 2, 2, 2, 298, 2, 24, 3, 2, 2, 2, 4, 30, 3, 2, 2, 2, 6, 58,
	3, 2, 2, 2, 8, 77, 3, 2, 2, 2, 10, 139, 3, 2, 2, 2, 12, 172, 3, 2, 2, 2,
	14, 174, 3, 2, 2, 2, 16, 192, 3, 2, 2, 2, 18, 203, 3, 2, 2, 2, 20, 220,
	3, 2, 2, 2, 22, 266, 3, 2, 2, 2, 24, 25, 5, 4, 3, 2, 25, 26, 7, 2, 2, 3,
	26, 3, 3, 2, 2, 2, 27, 29, 5, 8, 5, 2, 28, 27, 3, 2, 2, 2, 29, 32, 3, 2,
	2, 2, 30, 28, 3, 2, 2, 2, 30, 31, 3, 2, 2, 2, 31, 33, 3, 2, 2, 2, 32, 30,
	3, 2, 2, 2, 33, 34, 7, 20, 2, 2, 34, 39, 7, 20, 2, 2, 35, 36, 7, 13, 2,
	2, 36, 38, 7, 20, 2, 2, 37, 35, 3, 2, 2, 2, 38, 41, 3, 2, 2, 2, 39, 37,
	3, 2, 2, 2, 39, 40, 3, 2, 2, 2, 40, 42, 3, 2, 2, 2, 41, 39, 3, 2, 2, 2,
	42, 46, 7, 3, 2, 2, 43, 45, 5, 6, 4, 2, 44, 43, 3, 2, 2, 2, 45, 48, 3,
	2, 2, 2, 46, 44, 3, 2, 2, 2, 46, 47, 3, 2, 2, 2, 47, 52, 3, 2, 2, 2, 48,
	46, 3, 2, 2, 2, 49, 51, 5, 10, 6, 2, 50, 49, 3, 2, 2, 2, 51, 54, 3, 2,
	2, 2, 52, 50, 3, 2, 2, 2, 52, 53, 3, 2, 2, 2, 53, 55, 3, 2, 2, 2, 54, 52,
	3, 2, 2, 2, 55, 56, 7, 4, 2, 2, 56, 57, 7, 10, 2, 2, 57, 5, 3, 2, 2, 2,
	58, 59, 7, 20, 2, 2, 59, 64, 7, 20, 2, 2, 60, 61, 7, 13, 2, 2, 61, 63,
	7, 20, 2, 2, 62, 60, 3, 2, 2, 2, 63, 66, 3, 2, 2, 2, 64, 62, 3, 2, 2, 2,
	64, 65, 3, 2, 2, 2, 65, 69, 3, 2, 2, 2, 66, 64, 3, 2, 2, 2, 67, 68, 7,
	13, 2, 2, 68, 70, 7, 17, 2, 2, 69, 67, 3, 2, 2, 2, 69, 70, 3, 2, 2, 2,
	70, 71, 3, 2, 2, 2, 71, 72, 7, 10, 2, 2, 72, 7, 3, 2, 2, 2, 73, 74, 7,
	18, 2, 2, 74, 75, 7, 20, 2, 2, 75, 78, 5, 22, 12, 2, 76, 78, 7, 24, 2,
	2, 77, 73, 3, 2, 2, 2, 77, 76, 3, 2, 2, 2, 78, 9, 3, 2, 2, 2, 79, 81, 5,
	8, 5, 2, 80, 79, 3, 2, 2, 2, 81, 84, 3, 2, 2, 2, 82, 80, 3, 2, 2, 2, 82,
	83, 3, 2, 2, 2, 83, 85, 3, 2, 2, 2, 84, 82, 3, 2, 2, 2, 85, 86, 7, 20,
	2, 2, 86, 88, 7, 20, 2, 2, 87, 89, 5, 12, 7, 2, 88, 87, 3, 2, 2, 2, 88,
	89, 3, 2, 2, 2, 89, 90, 3, 2, 2, 2, 90, 94, 7, 3, 2, 2, 91, 93, 5, 20,
	11, 2, 92, 91, 3, 2, 2, 2, 93, 96, 3, 2, 2, 2, 94, 92, 3, 2, 2, 2, 94,
	95, 3, 2, 2, 2, 95, 97, 3, 2, 2, 2, 96, 94, 3, 2, 2, 2, 97, 98, 7, 4, 2,
	2, 98, 140, 7, 10, 2, 2, 99, 101, 5, 8, 5, 2, 100, 99, 3, 2, 2, 2, 101,
	104, 3, 2, 2, 2, 102, 100, 3, 2, 2, 2, 102, 103, 3, 2, 2, 2, 103, 105,
	3, 2, 2, 2, 104, 102, 3, 2, 2, 2, 105, 106, 7, 20, 2, 2, 106, 108, 7, 20,
	2, 2, 107, 109, 5, 12, 7, 2, 108, 107, 3, 2, 2, 2, 108, 109, 3, 2, 2, 2,
	109, 110, 3, 2, 2, 2, 110, 111, 7, 7, 2, 2, 111, 113, 7, 20, 2, 2, 112,
	114, 5, 16, 9, 2, 113, 112, 3, 2, 2, 2, 113, 114, 3, 2, 2, 2, 114, 117,
	3, 2, 2, 2, 115, 116, 7, 7, 2, 2, 116, 118, 5, 22, 12, 2, 117, 115, 3,
	2, 2, 2, 117, 118, 3, 2, 2, 2, 118, 119, 3, 2, 2, 2, 119, 140, 7, 10, 2,
	2, 120, 121, 7, 20, 2, 2, 121, 122, 7, 20, 2, 2, 122, 123, 5, 22, 12, 2,
	123, 124, 7, 10, 2, 2, 124, 140, 3, 2, 2, 2, 125, 126, 7, 20, 2, 2, 126,
	127, 7, 20, 2, 2, 127, 128, 7, 20, 2, 2, 128, 129, 5, 22, 12, 2, 129, 130,
	7, 10, 2, 2, 130, 140, 3, 2, 2, 2, 131, 132, 7, 20, 2, 2, 132, 133, 7,
	20, 2, 2, 133, 134, 7, 11, 2, 2, 134, 135, 7, 20, 2, 2, 135, 136, 7, 20,
	2, 2, 136, 137, 5, 22, 12, 2, 137, 138, 7, 10, 2, 2, 138, 140, 3, 2, 2,
	2, 139, 82, 3, 2, 2, 2, 139, 102, 3, 2, 2, 2, 139, 120, 3, 2, 2, 2, 139,
	125, 3, 2, 2, 2, 139, 131, 3, 2, 2, 2, 140, 11, 3, 2, 2, 2, 141, 142, 7,
	15, 2, 2, 142, 147, 7, 20, 2, 2, 143, 144, 7, 14, 2, 2, 144, 146, 7, 20,
	2, 2, 145, 143, 3, 2, 2, 2, 146, 149, 3, 2, 2, 2, 147, 145, 3, 2, 2, 2,
	147, 148, 3, 2, 2, 2, 148, 150, 3, 2, 2, 2, 149, 147, 3, 2, 2, 2, 150,
	173, 7, 16, 2, 2, 151, 152, 7, 15, 2, 2, 152, 157, 5, 14, 8, 2, 153, 154,
	7, 14, 2, 2, 154, 156, 7, 20, 2, 2, 155, 153, 3, 2, 2, 2, 156, 159, 3,
	2, 2, 2, 157, 155, 3, 2, 2, 2, 157, 158, 3, 2, 2, 2, 158, 160, 3, 2, 2,
	2, 159, 157, 3, 2, 2, 2, 160, 161, 7, 16, 2, 2, 161, 173, 3, 2, 2, 2, 162,
	163, 7, 15, 2, 2, 163, 166, 7, 20, 2, 2, 164, 165, 7, 14, 2, 2, 165, 167,
	5, 14, 8, 2, 166, 164, 3, 2, 2, 2, 167, 168, 3, 2, 2, 2, 168, 166, 3, 2,
	2, 2, 168, 169, 3, 2, 2, 2, 169, 170, 3, 2, 2, 2, 170, 171, 7, 16, 2, 2,
	171, 173, 3, 2, 2, 2, 172, 141, 3, 2, 2, 2, 172, 151, 3, 2, 2, 2, 172,
	162, 3, 2, 2, 2, 173, 13, 3, 2, 2, 2, 174, 175, 7, 20, 2, 2, 175, 178,
	7, 15, 2, 2, 176, 179, 7, 20, 2, 2, 177, 179, 5, 14, 8, 2, 178, 176, 3,
	2, 2, 2, 178, 177, 3, 2, 2, 2, 179, 187, 3, 2, 2, 2, 180, 183, 7, 14, 2,
	2, 181, 184, 7, 20, 2, 2, 182, 184, 5, 14, 8, 2, 183, 181, 3, 2, 2, 2,
	183, 182, 3, 2, 2, 2, 184, 186, 3, 2, 2, 2, 185, 180, 3, 2, 2, 2, 186,
	189, 3, 2, 2, 2, 187, 185, 3, 2, 2, 2, 187, 188, 3, 2, 2, 2, 188, 190,
	3, 2, 2, 2, 189, 187, 3, 2, 2, 2, 190, 191, 7, 16, 2, 2, 191, 15, 3, 2,
	2, 2, 192, 193, 7, 15, 2, 2, 193, 198, 5, 18, 10, 2, 194, 195, 7, 14, 2,
	2, 195, 197, 5, 18, 10, 2, 196, 194, 3, 2, 2, 2, 197, 200, 3, 2, 2, 2,
	198, 196, 3, 2, 2, 2, 198, 199, 3, 2, 2, 2, 199, 201, 3, 2, 2, 2, 200,
	198, 3, 2, 2, 2, 201, 202, 7, 16, 2, 2, 202, 17, 3, 2, 2, 2, 203, 215,
	7, 20, 2, 2, 204, 205, 7, 15, 2, 2, 205, 210, 5, 18, 10, 2, 206, 207, 7,
	14, 2, 2, 207, 209, 5, 18, 10, 2, 208, 206, 3, 2, 2, 2, 209, 212, 3, 2,
	2, 2, 210, 208, 3, 2, 2, 2, 210, 211, 3, 2, 2, 2, 211, 213, 3, 2, 2, 2,
	212, 210, 3, 2, 2, 2, 213, 214, 7, 16, 2, 2, 214, 216, 3, 2, 2, 2, 215,
	204, 3, 2, 2, 2, 215, 216, 3, 2, 2, 2, 216, 19, 3, 2, 2, 2, 217, 219, 5,
	8, 5, 2, 218, 217, 3, 2, 2, 2, 219, 222, 3, 2, 2, 2, 220, 218, 3, 2, 2,
	2, 220, 221, 3, 2, 2, 2, 221, 223, 3, 2, 2, 2, 222, 220, 3, 2, 2, 2, 223,
	225, 7, 20, 2, 2, 224, 226, 5, 16, 9, 2, 225, 224, 3, 2, 2, 2, 225, 226,
	3, 2, 2, 2, 226, 227, 3, 2, 2, 2, 227, 230, 7, 20, 2, 2, 228, 229, 7, 7,
	2, 2, 229, 231, 5, 22, 12, 2, 230, 228, 3, 2, 2, 2, 230, 231, 3, 2, 2,
	2, 231, 232, 3, 2, 2, 2, 232, 233, 7, 10, 2, 2, 233, 21, 3, 2, 2, 2, 234,
	267, 7, 19, 2, 2, 235, 267, 7, 20, 2, 2, 236, 267, 7, 21, 2, 2, 237, 267,
	7, 22, 2, 2, 238, 247, 7, 5, 2, 2, 239, 244, 5, 22, 12, 2, 240, 241, 7,
	14, 2, 2, 241, 243, 5, 22, 12, 2, 242, 240, 3, 2, 2, 2, 243, 246, 3, 2,
	2, 2, 244, 242, 3, 2, 2, 2, 244, 245, 3, 2, 2, 2, 245, 248, 3, 2, 2, 2,
	246, 244, 3, 2, 2, 2, 247, 239, 3, 2, 2, 2, 247, 248, 3, 2, 2, 2, 248,
	249, 3, 2, 2, 2, 249, 267, 7, 6, 2, 2, 250, 263, 7, 3, 2, 2, 251, 252,
	7, 19, 2, 2, 252, 253, 7, 12, 2, 2, 253, 260, 5, 22, 12, 2, 254, 255, 7,
	14, 2, 2, 255, 256, 7, 19, 2, 2, 256, 257, 7, 12, 2, 2, 257, 259, 5, 22,
	12, 2, 258, 254, 3, 2, 2, 2, 259, 262, 3, 2, 2, 2, 260, 258, 3, 2, 2, 2,
	260, 261, 3, 2, 2, 2, 261, 264, 3, 2, 2, 2, 262, 260, 3, 2, 2, 2, 263,
	251, 3, 2, 2, 2, 263, 264, 3, 2, 2, 2, 264, 265, 3, 2, 2, 2, 265, 267,
	7, 4, 2, 2, 266, 234, 3, 2, 2, 2, 266, 235, 3, 2, 2, 2, 266, 236, 3, 2,
	2, 2, 266, 237, 3, 2, 2, 2, 266, 238, 3, 2, 2, 2, 266, 250, 3, 2, 2, 2,
	267, 23, 3, 2, 2, 2, 35, 30, 39, 46, 52, 64, 69, 77, 82, 88, 94, 102, 108,
	113, 117, 139, 147, 157, 168, 172, 178, 183, 187, 198, 210, 215, 220, 225,
	230, 244, 247, 260, 263, 266,
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
	"Module", "Import", "Annotation", "Struct", "Union", "Newtype", "Type",
	"TypeParam", "TypeExpr", "TypeExprElem", "Field", "Json", "JsonStr", "JsonBool",
	"JsonNull", "JsonInt", "JsonFloat", "JsonArray", "JsonObj", "ModuleAnno",
	"DeclAnno", "FieldAnno",
}

var ruleNames = []string{
	"adl", "module", "imports", "annon", "top_level_statement", "typeParam",
	"typeParamError", "typeExpr", "typeExprElem", "soruBody", "jsonValue",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type AdlP struct {
	*antlr.BaseParser
}

func NewAdlP(input antlr.TokenStream) *AdlP {
	this := new(AdlP)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "AdlP.g4"

	return this
}

// AdlP tokens.
const (
	AdlPEOF          = antlr.TokenEOF
	AdlPLCUR         = 1
	AdlPRCUR         = 2
	AdlPLSQ          = 3
	AdlPRSQ          = 4
	AdlPEQ           = 5
	AdlPDQ           = 6
	AdlPSQ           = 7
	AdlPSEMI         = 8
	AdlPDCOLON       = 9
	AdlPCOLON        = 10
	AdlPDOT          = 11
	AdlPCOMMA        = 12
	AdlPLCHEVR       = 13
	AdlPRCHEVR       = 14
	AdlPSTAR         = 15
	AdlPAT           = 16
	AdlPSTR          = 17
	AdlPID           = 18
	AdlPINT          = 19
	AdlPFLT          = 20
	AdlPWS           = 21
	AdlPLINE_DOC     = 22
	AdlPLINE_COMMENT = 23
	AdlPDOWN         = 24
	AdlPUP           = 25
	AdlPROOT         = 26
	AdlPERROR        = 27
	AdlPADL          = 28
	AdlPModule       = 29
	AdlPImport       = 30
	AdlPAnnotation   = 31
	AdlPStruct       = 32
	AdlPUnion        = 33
	AdlPNewtype      = 34
	AdlPType         = 35
	AdlPTypeParam    = 36
	AdlPTypeExpr     = 37
	AdlPTypeExprElem = 38
	AdlPField        = 39
	AdlPJson         = 40
	AdlPJsonStr      = 41
	AdlPJsonBool     = 42
	AdlPJsonNull     = 43
	AdlPJsonInt      = 44
	AdlPJsonFloat    = 45
	AdlPJsonArray    = 46
	AdlPJsonObj      = 47
	AdlPModuleAnno   = 48
	AdlPDeclAnno     = 49
	AdlPFieldAnno    = 50
)

// AdlP rules.
const (
	AdlPRULE_adl                 = 0
	AdlPRULE_module              = 1
	AdlPRULE_imports             = 2
	AdlPRULE_annon               = 3
	AdlPRULE_top_level_statement = 4
	AdlPRULE_typeParam           = 5
	AdlPRULE_typeParamError      = 6
	AdlPRULE_typeExpr            = 7
	AdlPRULE_typeExprElem        = 8
	AdlPRULE_soruBody            = 9
	AdlPRULE_jsonValue           = 10
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
	p.RuleIndex = AdlPRULE_adl
	return p
}

func (*AdlContext) IsAdlContext() {}

func NewAdlContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AdlContext {
	var p = new(AdlContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AdlPRULE_adl

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
	return s.GetToken(AdlPEOF, 0)
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
	h := hdls.(*AdlPHandlers)
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

func (p *AdlP) Adl() (localctx IAdlContext) {
	localctx = NewAdlContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, AdlPRULE_adl)

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
		p.SetState(22)
		p.Module()
	}
	{
		p.SetState(23)
		p.Match(AdlPEOF)
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
}

func NewEmptyModuleContext() *ModuleContext {
	var p = new(ModuleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AdlPRULE_module
	return p
}

func (*ModuleContext) IsModuleContext() {}

func NewModuleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ModuleContext {
	var p = new(ModuleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AdlPRULE_module

	return p
}

func (s *ModuleContext) GetParser() antlr.Parser { return s.parser }

//StructDecl tokenDecls

//StructDecl tokenTypeDecls

//StructDecl tokenListDecls

//StructDecl ruleContextDecls

//StructDecl ruleContextListDecls

//StructDecl attributeDecls

// Getters

//provideCopyFrom
func (s *ModuleContext) CopyFrom(ctx *ModuleContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ModuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ModuleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

//dispatchMethods
//extensionMembers

//Begin AltLabelStructDecl

type ModuleStatementContext struct {
	*ModuleContext
	//TokenDecl
	kw antlr.Token
	//TokenDecl
	_ID antlr.Token
	//TokenListDecl
	name []antlr.Token
}

func NewModuleStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ModuleStatementContext {
	var p = new(ModuleStatementContext)

	p.ModuleContext = NewEmptyModuleContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ModuleContext))

	return p
}

type IModuleStatementContext interface {
	//Current rule
	IModuleContext
	// start internal
	//Gets for raw elements
	//  ruleGetterDecl
	//  ruleListGetterDecl
	AllAnnon() []IAnnonContext
	AllImports() []IImportsContext
	AllTop_level_statement() []ITop_level_statementContext
	//  ruleListIndexedGetterDecl
	Annon(i int) IAnnonContext
	Imports(i int) IImportsContext
	Top_level_statement(i int) ITop_level_statementContext

	//  tokenGetterDecl
	LCUR() antlr.TerminalNode
	RCUR() antlr.TerminalNode
	SEMI() antlr.TerminalNode
	//  tokenListGetterDecl
	AllID() []antlr.TerminalNode
	AllDOT() []antlr.TerminalNode
	//  tokenListIndexedGetterDecl
	ID(i int) antlr.TerminalNode
	DOT(i int) antlr.TerminalNode
	// end internal
	//
	//Gets for labeled elements
	//  tokenDecls
	GetKw() antlr.Token

	//  tokenTypeDecls
	//  tokenListDecls
	GetName() []antlr.Token
	//  ruleContextDecls
	//  ruleContextListDecls
	//  attributeDecls

	// TODO dispatchMethods (needed?)
}

func (*ModuleStatementContext) IsModuleStatementContext() {}

//AltLabelStructDecl tokenDecls
func (s *ModuleStatementContext) GetKw() antlr.Token  { return s.kw }
func (s *ModuleStatementContext) SetKw(v antlr.Token) { s.kw = v }

func (s *ModuleStatementContext) Get_ID() antlr.Token  { return s._ID }
func (s *ModuleStatementContext) Set_ID(v antlr.Token) { s._ID = v }

//AltLabelStructDecl tokenTypeDecls

//AltLabelStructDecl tokenListDecls
func (s *ModuleStatementContext) GetName() []antlr.Token  { return s.name }
func (s *ModuleStatementContext) SetName(v []antlr.Token) { s.name = v }

//AltLabelStructDecl ruleContextDecls

//AltLabelStructDecl ruleContextListDecls

//AltLabelStructDecl attributeDecls

//getRuleContext
func (s *ModuleStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

//getters
func (s *ModuleStatementContext) LCUR() antlr.TerminalNode {
	return s.GetToken(AdlPLCUR, 0)
}

func (s *ModuleStatementContext) RCUR() antlr.TerminalNode {
	return s.GetToken(AdlPRCUR, 0)
}

func (s *ModuleStatementContext) SEMI() antlr.TerminalNode {
	return s.GetToken(AdlPSEMI, 0)
}

func (s *ModuleStatementContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(AdlPID)
}

func (s *ModuleStatementContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(AdlPID, i)
}

func (s *ModuleStatementContext) AllAnnon() []IAnnonContext {
	//ContextRuleListGetterDecl
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*AnnonContext)(nil)).Elem())
	var tst = make([]IAnnonContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAnnonContext)
		}
	}

	return tst
}

func (s *ModuleStatementContext) Annon(i int) IAnnonContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*AnnonContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAnnonContext)
}

func (s *ModuleStatementContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(AdlPDOT)
}

func (s *ModuleStatementContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(AdlPDOT, i)
}

func (s *ModuleStatementContext) AllImports() []IImportsContext {
	//ContextRuleListGetterDecl
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ImportsContext)(nil)).Elem())
	var tst = make([]IImportsContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IImportsContext)
		}
	}

	return tst
}

func (s *ModuleStatementContext) Imports(i int) IImportsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ImportsContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IImportsContext)
}

func (s *ModuleStatementContext) AllTop_level_statement() []ITop_level_statementContext {
	//ContextRuleListGetterDecl
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*Top_level_statementContext)(nil)).Elem())
	var tst = make([]ITop_level_statementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITop_level_statementContext)
		}
	}

	return tst
}

func (s *ModuleStatementContext) Top_level_statement(i int) ITop_level_statementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*Top_level_statementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITop_level_statementContext)
}

func (s *ModuleStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ModuleStatementEntryListener); ok {
		listenerT.EnterModuleStatement(s)
	}
}

func (s *ModuleStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ModuleStatementExitListener); ok {
		listenerT.ExitModuleStatement(s)
	}
}

func (s *ModuleStatementContext) VisitFunc(hdls antlr.ParserTreeVisitorHandlers, args ...interface{}) (result interface{}) {
	h := hdls.(*AdlPHandlers)
	if h.EnterEveryRule != nil {
		h.EnterEveryRule(s)
	}
	if h.ModuleStatement != nil {
		h.ModuleStatement(s, h, args...)
	} else {
		s.VisitChildrenFunc(h, args...)
	}
	if h.ExitEveryRule != nil {
		h.ExitEveryRule(s)
	}
	return
}

func (s *ModuleStatementContext) Visit(delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	switch t := delegate.(type) {
	case ModuleStatementContextVisitor:
		return t.VisitModuleStatement(s, delegate, args...)
	default:
		return delegate.VisitChildren(s, delegate, args...)
	}
}

//END AltLabelStructDecl

func (p *AdlP) Module() (localctx IModuleContext) {
	localctx = NewModuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, AdlPRULE_module)
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

	var _alt int

	localctx = NewModuleStatementContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(28)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == AdlPAT || _la == AdlPLINE_DOC {
		{
			p.SetState(25)
			p.Annon()
		}

		p.SetState(30)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(31)
		var _m = p.Match(AdlPID)
		localctx.(*ModuleStatementContext).kw = _m

	}
	{
		p.SetState(32)
		var _m = p.Match(AdlPID)
		localctx.(*ModuleStatementContext)._ID = _m

	}
	localctx.(*ModuleStatementContext).name = append(localctx.(*ModuleStatementContext).name, localctx.(*ModuleStatementContext)._ID)
	p.SetState(37)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == AdlPDOT {
		{
			p.SetState(33)
			p.Match(AdlPDOT)
		}
		{
			p.SetState(34)
			var _m = p.Match(AdlPID)
			localctx.(*ModuleStatementContext)._ID = _m

		}
		localctx.(*ModuleStatementContext).name = append(localctx.(*ModuleStatementContext).name, localctx.(*ModuleStatementContext)._ID)

		p.SetState(39)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(40)
		p.Match(AdlPLCUR)
	}
	p.SetState(44)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(41)
				p.Imports()
			}

		}
		p.SetState(46)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())
	}
	p.SetState(50)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AdlPAT)|(1<<AdlPID)|(1<<AdlPLINE_DOC))) != 0 {
		{
			p.SetState(47)
			p.Top_level_statement()
		}

		p.SetState(52)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(53)
		p.Match(AdlPRCUR)
	}
	{
		p.SetState(54)
		p.Match(AdlPSEMI)
	}

	return localctx
}

type IImportsContext interface {
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

	// IsImportsContext differentiates from other interfaces.
	//copyStruct,GetRuleContext and ToStringTree  from embedded

	//<if(dispatchMethods)>
	//<dispatchMethods; separator="\n\n">
	//<endif>

	//<if(extensionMembers)>
	//<extensionMembers; separator="\n\n">
	//<endif>
}

type ImportsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyImportsContext() *ImportsContext {
	var p = new(ImportsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AdlPRULE_imports
	return p
}

func (*ImportsContext) IsImportsContext() {}

func NewImportsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImportsContext {
	var p = new(ImportsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AdlPRULE_imports

	return p
}

func (s *ImportsContext) GetParser() antlr.Parser { return s.parser }

//StructDecl tokenDecls

//StructDecl tokenTypeDecls

//StructDecl tokenListDecls

//StructDecl ruleContextDecls

//StructDecl ruleContextListDecls

//StructDecl attributeDecls

// Getters

//provideCopyFrom
func (s *ImportsContext) CopyFrom(ctx *ImportsContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ImportsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImportsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

//dispatchMethods
//extensionMembers

//Begin AltLabelStructDecl

type ImportStatementContext struct {
	*ImportsContext
	//TokenDecl
	kw antlr.Token
	//TokenDecl
	_ID antlr.Token
	//TokenListDecl
	a []antlr.Token
	//TokenDecl
	s antlr.Token
}

func NewImportStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ImportStatementContext {
	var p = new(ImportStatementContext)

	p.ImportsContext = NewEmptyImportsContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ImportsContext))

	return p
}

type IImportStatementContext interface {
	//Current rule
	IImportsContext
	// start internal
	//Gets for raw elements
	//  ruleGetterDecl
	//  ruleListGetterDecl
	//  ruleListIndexedGetterDecl

	//  tokenGetterDecl
	SEMI() antlr.TerminalNode
	STAR() antlr.TerminalNode
	//  tokenListGetterDecl
	AllID() []antlr.TerminalNode
	AllDOT() []antlr.TerminalNode
	//  tokenListIndexedGetterDecl
	ID(i int) antlr.TerminalNode
	DOT(i int) antlr.TerminalNode
	// end internal
	//
	//Gets for labeled elements
	//  tokenDecls
	GetKw() antlr.Token
	GetS() antlr.Token

	//  tokenTypeDecls
	//  tokenListDecls
	GetA() []antlr.Token
	//  ruleContextDecls
	//  ruleContextListDecls
	//  attributeDecls

	// TODO dispatchMethods (needed?)
}

func (*ImportStatementContext) IsImportStatementContext() {}

//AltLabelStructDecl tokenDecls
func (s *ImportStatementContext) GetKw() antlr.Token  { return s.kw }
func (s *ImportStatementContext) SetKw(v antlr.Token) { s.kw = v }

func (s *ImportStatementContext) Get_ID() antlr.Token  { return s._ID }
func (s *ImportStatementContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *ImportStatementContext) GetS() antlr.Token  { return s.s }
func (s *ImportStatementContext) SetS(v antlr.Token) { s.s = v }

//AltLabelStructDecl tokenTypeDecls

//AltLabelStructDecl tokenListDecls
func (s *ImportStatementContext) GetA() []antlr.Token  { return s.a }
func (s *ImportStatementContext) SetA(v []antlr.Token) { s.a = v }

//AltLabelStructDecl ruleContextDecls

//AltLabelStructDecl ruleContextListDecls

//AltLabelStructDecl attributeDecls

//getRuleContext
func (s *ImportStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

//getters
func (s *ImportStatementContext) SEMI() antlr.TerminalNode {
	return s.GetToken(AdlPSEMI, 0)
}

func (s *ImportStatementContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(AdlPID)
}

func (s *ImportStatementContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(AdlPID, i)
}

func (s *ImportStatementContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(AdlPDOT)
}

func (s *ImportStatementContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(AdlPDOT, i)
}

func (s *ImportStatementContext) STAR() antlr.TerminalNode {
	return s.GetToken(AdlPSTAR, 0)
}

func (s *ImportStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ImportStatementEntryListener); ok {
		listenerT.EnterImportStatement(s)
	}
}

func (s *ImportStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ImportStatementExitListener); ok {
		listenerT.ExitImportStatement(s)
	}
}

func (s *ImportStatementContext) VisitFunc(hdls antlr.ParserTreeVisitorHandlers, args ...interface{}) (result interface{}) {
	h := hdls.(*AdlPHandlers)
	if h.EnterEveryRule != nil {
		h.EnterEveryRule(s)
	}
	if h.ImportStatement != nil {
		h.ImportStatement(s, h, args...)
	} else {
		s.VisitChildrenFunc(h, args...)
	}
	if h.ExitEveryRule != nil {
		h.ExitEveryRule(s)
	}
	return
}

func (s *ImportStatementContext) Visit(delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	switch t := delegate.(type) {
	case ImportStatementContextVisitor:
		return t.VisitImportStatement(s, delegate, args...)
	default:
		return delegate.VisitChildren(s, delegate, args...)
	}
}

//END AltLabelStructDecl

func (p *AdlP) Imports() (localctx IImportsContext) {
	localctx = NewImportsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, AdlPRULE_imports)
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

	var _alt int

	localctx = NewImportStatementContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(56)
		var _m = p.Match(AdlPID)
		localctx.(*ImportStatementContext).kw = _m

	}
	{
		p.SetState(57)
		var _m = p.Match(AdlPID)
		localctx.(*ImportStatementContext)._ID = _m

	}
	localctx.(*ImportStatementContext).a = append(localctx.(*ImportStatementContext).a, localctx.(*ImportStatementContext)._ID)
	p.SetState(62)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(58)
				p.Match(AdlPDOT)
			}
			{
				p.SetState(59)
				var _m = p.Match(AdlPID)
				localctx.(*ImportStatementContext)._ID = _m

			}
			localctx.(*ImportStatementContext).a = append(localctx.(*ImportStatementContext).a, localctx.(*ImportStatementContext)._ID)

		}
		p.SetState(64)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())
	}
	p.SetState(67)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == AdlPDOT {
		{
			p.SetState(65)
			p.Match(AdlPDOT)
		}
		{
			p.SetState(66)
			var _m = p.Match(AdlPSTAR)
			localctx.(*ImportStatementContext).s = _m

		}

	}
	{
		p.SetState(69)
		p.Match(AdlPSEMI)
	}

	return localctx
}

type IAnnonContext interface {
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

	// IsAnnonContext differentiates from other interfaces.
	//copyStruct,GetRuleContext and ToStringTree  from embedded

	//<if(dispatchMethods)>
	//<dispatchMethods; separator="\n\n">
	//<endif>

	//<if(extensionMembers)>
	//<extensionMembers; separator="\n\n">
	//<endif>
}

type AnnonContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAnnonContext() *AnnonContext {
	var p = new(AnnonContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AdlPRULE_annon
	return p
}

func (*AnnonContext) IsAnnonContext() {}

func NewAnnonContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AnnonContext {
	var p = new(AnnonContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AdlPRULE_annon

	return p
}

func (s *AnnonContext) GetParser() antlr.Parser { return s.parser }

//StructDecl tokenDecls

//StructDecl tokenTypeDecls

//StructDecl tokenListDecls

//StructDecl ruleContextDecls

//StructDecl ruleContextListDecls

//StructDecl attributeDecls

// Getters

//provideCopyFrom
func (s *AnnonContext) CopyFrom(ctx *AnnonContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *AnnonContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AnnonContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

//dispatchMethods
//extensionMembers

//Begin AltLabelStructDecl

type DocAnnoContext struct {
	*AnnonContext
	//TokenDecl
	a antlr.Token
}

func NewDocAnnoContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DocAnnoContext {
	var p = new(DocAnnoContext)

	p.AnnonContext = NewEmptyAnnonContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AnnonContext))

	return p
}

type IDocAnnoContext interface {
	//Current rule
	IAnnonContext
	// start internal
	//Gets for raw elements
	//  ruleGetterDecl
	//  ruleListGetterDecl
	//  ruleListIndexedGetterDecl

	//  tokenGetterDecl
	LINE_DOC() antlr.TerminalNode
	//  tokenListGetterDecl
	//  tokenListIndexedGetterDecl
	// end internal
	//
	//Gets for labeled elements
	//  tokenDecls
	GetA() antlr.Token

	//  tokenTypeDecls
	//  tokenListDecls
	//  ruleContextDecls
	//  ruleContextListDecls
	//  attributeDecls

	// TODO dispatchMethods (needed?)
}

func (*DocAnnoContext) IsDocAnnoContext() {}

//AltLabelStructDecl tokenDecls
func (s *DocAnnoContext) GetA() antlr.Token  { return s.a }
func (s *DocAnnoContext) SetA(v antlr.Token) { s.a = v }

//AltLabelStructDecl tokenTypeDecls

//AltLabelStructDecl tokenListDecls

//AltLabelStructDecl ruleContextDecls

//AltLabelStructDecl ruleContextListDecls

//AltLabelStructDecl attributeDecls

//getRuleContext
func (s *DocAnnoContext) GetRuleContext() antlr.RuleContext {
	return s
}

//getters
func (s *DocAnnoContext) LINE_DOC() antlr.TerminalNode {
	return s.GetToken(AdlPLINE_DOC, 0)
}

func (s *DocAnnoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DocAnnoEntryListener); ok {
		listenerT.EnterDocAnno(s)
	}
}

func (s *DocAnnoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DocAnnoExitListener); ok {
		listenerT.ExitDocAnno(s)
	}
}

func (s *DocAnnoContext) VisitFunc(hdls antlr.ParserTreeVisitorHandlers, args ...interface{}) (result interface{}) {
	h := hdls.(*AdlPHandlers)
	if h.EnterEveryRule != nil {
		h.EnterEveryRule(s)
	}
	if h.DocAnno != nil {
		h.DocAnno(s, h, args...)
	} else {
		s.VisitChildrenFunc(h, args...)
	}
	if h.ExitEveryRule != nil {
		h.ExitEveryRule(s)
	}
	return
}

func (s *DocAnnoContext) Visit(delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	switch t := delegate.(type) {
	case DocAnnoContextVisitor:
		return t.VisitDocAnno(s, delegate, args...)
	default:
		return delegate.VisitChildren(s, delegate, args...)
	}
}

//END AltLabelStructDecl

//Begin AltLabelStructDecl

type LocalAnnoContext struct {
	*AnnonContext
	//TokenDecl
	a antlr.Token
}

func NewLocalAnnoContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LocalAnnoContext {
	var p = new(LocalAnnoContext)

	p.AnnonContext = NewEmptyAnnonContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AnnonContext))

	return p
}

type ILocalAnnoContext interface {
	//Current rule
	IAnnonContext
	// start internal
	//Gets for raw elements
	//  ruleGetterDecl
	JsonValue() IJsonValueContext
	//  ruleListGetterDecl
	//  ruleListIndexedGetterDecl

	//  tokenGetterDecl
	AT() antlr.TerminalNode
	ID() antlr.TerminalNode
	//  tokenListGetterDecl
	//  tokenListIndexedGetterDecl
	// end internal
	//
	//Gets for labeled elements
	//  tokenDecls
	GetA() antlr.Token

	//  tokenTypeDecls
	//  tokenListDecls
	//  ruleContextDecls
	//  ruleContextListDecls
	//  attributeDecls

	// TODO dispatchMethods (needed?)
}

func (*LocalAnnoContext) IsLocalAnnoContext() {}

//AltLabelStructDecl tokenDecls
func (s *LocalAnnoContext) GetA() antlr.Token  { return s.a }
func (s *LocalAnnoContext) SetA(v antlr.Token) { s.a = v }

//AltLabelStructDecl tokenTypeDecls

//AltLabelStructDecl tokenListDecls

//AltLabelStructDecl ruleContextDecls

//AltLabelStructDecl ruleContextListDecls

//AltLabelStructDecl attributeDecls

//getRuleContext
func (s *LocalAnnoContext) GetRuleContext() antlr.RuleContext {
	return s
}

//getters
func (s *LocalAnnoContext) AT() antlr.TerminalNode {
	return s.GetToken(AdlPAT, 0)
}

func (s *LocalAnnoContext) JsonValue() IJsonValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*JsonValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IJsonValueContext)
}

func (s *LocalAnnoContext) ID() antlr.TerminalNode {
	return s.GetToken(AdlPID, 0)
}

func (s *LocalAnnoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LocalAnnoEntryListener); ok {
		listenerT.EnterLocalAnno(s)
	}
}

func (s *LocalAnnoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LocalAnnoExitListener); ok {
		listenerT.ExitLocalAnno(s)
	}
}

func (s *LocalAnnoContext) VisitFunc(hdls antlr.ParserTreeVisitorHandlers, args ...interface{}) (result interface{}) {
	h := hdls.(*AdlPHandlers)
	if h.EnterEveryRule != nil {
		h.EnterEveryRule(s)
	}
	if h.LocalAnno != nil {
		h.LocalAnno(s, h, args...)
	} else {
		s.VisitChildrenFunc(h, args...)
	}
	if h.ExitEveryRule != nil {
		h.ExitEveryRule(s)
	}
	return
}

func (s *LocalAnnoContext) Visit(delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	switch t := delegate.(type) {
	case LocalAnnoContextVisitor:
		return t.VisitLocalAnno(s, delegate, args...)
	default:
		return delegate.VisitChildren(s, delegate, args...)
	}
}

//END AltLabelStructDecl

func (p *AdlP) Annon() (localctx IAnnonContext) {
	localctx = NewAnnonContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, AdlPRULE_annon)

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

	p.SetState(75)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AdlPAT:
		localctx = NewLocalAnnoContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(71)
			p.Match(AdlPAT)
		}
		{
			p.SetState(72)
			var _m = p.Match(AdlPID)
			localctx.(*LocalAnnoContext).a = _m

		}
		{
			p.SetState(73)
			p.JsonValue()
		}

	case AdlPLINE_DOC:
		localctx = NewDocAnnoContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(74)
			var _m = p.Match(AdlPLINE_DOC)
			localctx.(*DocAnnoContext).a = _m

		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

type ITop_level_statementContext interface {
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

	// IsTop_level_statementContext differentiates from other interfaces.
	//copyStruct,GetRuleContext and ToStringTree  from embedded

	//<if(dispatchMethods)>
	//<dispatchMethods; separator="\n\n">
	//<endif>

	//<if(extensionMembers)>
	//<extensionMembers; separator="\n\n">
	//<endif>
}

type Top_level_statementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTop_level_statementContext() *Top_level_statementContext {
	var p = new(Top_level_statementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AdlPRULE_top_level_statement
	return p
}

func (*Top_level_statementContext) IsTop_level_statementContext() {}

func NewTop_level_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Top_level_statementContext {
	var p = new(Top_level_statementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AdlPRULE_top_level_statement

	return p
}

func (s *Top_level_statementContext) GetParser() antlr.Parser { return s.parser }

//StructDecl tokenDecls

//StructDecl tokenTypeDecls

//StructDecl tokenListDecls

//StructDecl ruleContextDecls

//StructDecl ruleContextListDecls

//StructDecl attributeDecls

// Getters

//provideCopyFrom
func (s *Top_level_statementContext) CopyFrom(ctx *Top_level_statementContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *Top_level_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Top_level_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

//dispatchMethods
//extensionMembers

//Begin AltLabelStructDecl

type StructOrUnionContext struct {
	*Top_level_statementContext
	//TokenDecl
	kw antlr.Token
	//TokenDecl
	a antlr.Token
}

func NewStructOrUnionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StructOrUnionContext {
	var p = new(StructOrUnionContext)

	p.Top_level_statementContext = NewEmptyTop_level_statementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Top_level_statementContext))

	return p
}

type IStructOrUnionContext interface {
	//Current rule
	ITop_level_statementContext
	// start internal
	//Gets for raw elements
	//  ruleGetterDecl
	TypeParam() ITypeParamContext
	//  ruleListGetterDecl
	AllAnnon() []IAnnonContext
	AllSoruBody() []ISoruBodyContext
	//  ruleListIndexedGetterDecl
	Annon(i int) IAnnonContext
	SoruBody(i int) ISoruBodyContext

	//  tokenGetterDecl
	LCUR() antlr.TerminalNode
	RCUR() antlr.TerminalNode
	SEMI() antlr.TerminalNode
	//  tokenListGetterDecl
	AllID() []antlr.TerminalNode
	//  tokenListIndexedGetterDecl
	ID(i int) antlr.TerminalNode
	// end internal
	//
	//Gets for labeled elements
	//  tokenDecls
	GetKw() antlr.Token
	GetA() antlr.Token

	//  tokenTypeDecls
	//  tokenListDecls
	//  ruleContextDecls
	//  ruleContextListDecls
	//  attributeDecls

	// TODO dispatchMethods (needed?)
}

func (*StructOrUnionContext) IsStructOrUnionContext() {}

//AltLabelStructDecl tokenDecls
func (s *StructOrUnionContext) GetKw() antlr.Token  { return s.kw }
func (s *StructOrUnionContext) SetKw(v antlr.Token) { s.kw = v }

func (s *StructOrUnionContext) GetA() antlr.Token  { return s.a }
func (s *StructOrUnionContext) SetA(v antlr.Token) { s.a = v }

//AltLabelStructDecl tokenTypeDecls

//AltLabelStructDecl tokenListDecls

//AltLabelStructDecl ruleContextDecls

//AltLabelStructDecl ruleContextListDecls

//AltLabelStructDecl attributeDecls

//getRuleContext
func (s *StructOrUnionContext) GetRuleContext() antlr.RuleContext {
	return s
}

//getters
func (s *StructOrUnionContext) LCUR() antlr.TerminalNode {
	return s.GetToken(AdlPLCUR, 0)
}

func (s *StructOrUnionContext) RCUR() antlr.TerminalNode {
	return s.GetToken(AdlPRCUR, 0)
}

func (s *StructOrUnionContext) SEMI() antlr.TerminalNode {
	return s.GetToken(AdlPSEMI, 0)
}

func (s *StructOrUnionContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(AdlPID)
}

func (s *StructOrUnionContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(AdlPID, i)
}

func (s *StructOrUnionContext) AllAnnon() []IAnnonContext {
	//ContextRuleListGetterDecl
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*AnnonContext)(nil)).Elem())
	var tst = make([]IAnnonContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAnnonContext)
		}
	}

	return tst
}

func (s *StructOrUnionContext) Annon(i int) IAnnonContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*AnnonContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAnnonContext)
}

func (s *StructOrUnionContext) TypeParam() ITypeParamContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*TypeParamContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeParamContext)
}

func (s *StructOrUnionContext) AllSoruBody() []ISoruBodyContext {
	//ContextRuleListGetterDecl
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*SoruBodyContext)(nil)).Elem())
	var tst = make([]ISoruBodyContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISoruBodyContext)
		}
	}

	return tst
}

func (s *StructOrUnionContext) SoruBody(i int) ISoruBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*SoruBodyContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISoruBodyContext)
}

func (s *StructOrUnionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StructOrUnionEntryListener); ok {
		listenerT.EnterStructOrUnion(s)
	}
}

func (s *StructOrUnionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StructOrUnionExitListener); ok {
		listenerT.ExitStructOrUnion(s)
	}
}

func (s *StructOrUnionContext) VisitFunc(hdls antlr.ParserTreeVisitorHandlers, args ...interface{}) (result interface{}) {
	h := hdls.(*AdlPHandlers)
	if h.EnterEveryRule != nil {
		h.EnterEveryRule(s)
	}
	if h.StructOrUnion != nil {
		h.StructOrUnion(s, h, args...)
	} else {
		s.VisitChildrenFunc(h, args...)
	}
	if h.ExitEveryRule != nil {
		h.ExitEveryRule(s)
	}
	return
}

func (s *StructOrUnionContext) Visit(delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	switch t := delegate.(type) {
	case StructOrUnionContextVisitor:
		return t.VisitStructOrUnion(s, delegate, args...)
	default:
		return delegate.VisitChildren(s, delegate, args...)
	}
}

//END AltLabelStructDecl

//Begin AltLabelStructDecl

type DeclAnnotationContext struct {
	*Top_level_statementContext
	//TokenDecl
	kw antlr.Token
	//TokenDecl
	a antlr.Token
	//TokenDecl
	b antlr.Token
}

func NewDeclAnnotationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DeclAnnotationContext {
	var p = new(DeclAnnotationContext)

	p.Top_level_statementContext = NewEmptyTop_level_statementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Top_level_statementContext))

	return p
}

type IDeclAnnotationContext interface {
	//Current rule
	ITop_level_statementContext
	// start internal
	//Gets for raw elements
	//  ruleGetterDecl
	JsonValue() IJsonValueContext
	//  ruleListGetterDecl
	//  ruleListIndexedGetterDecl

	//  tokenGetterDecl
	SEMI() antlr.TerminalNode
	//  tokenListGetterDecl
	AllID() []antlr.TerminalNode
	//  tokenListIndexedGetterDecl
	ID(i int) antlr.TerminalNode
	// end internal
	//
	//Gets for labeled elements
	//  tokenDecls
	GetKw() antlr.Token
	GetA() antlr.Token
	GetB() antlr.Token

	//  tokenTypeDecls
	//  tokenListDecls
	//  ruleContextDecls
	//  ruleContextListDecls
	//  attributeDecls

	// TODO dispatchMethods (needed?)
}

func (*DeclAnnotationContext) IsDeclAnnotationContext() {}

//AltLabelStructDecl tokenDecls
func (s *DeclAnnotationContext) GetKw() antlr.Token  { return s.kw }
func (s *DeclAnnotationContext) SetKw(v antlr.Token) { s.kw = v }

func (s *DeclAnnotationContext) GetA() antlr.Token  { return s.a }
func (s *DeclAnnotationContext) SetA(v antlr.Token) { s.a = v }

func (s *DeclAnnotationContext) GetB() antlr.Token  { return s.b }
func (s *DeclAnnotationContext) SetB(v antlr.Token) { s.b = v }

//AltLabelStructDecl tokenTypeDecls

//AltLabelStructDecl tokenListDecls

//AltLabelStructDecl ruleContextDecls

//AltLabelStructDecl ruleContextListDecls

//AltLabelStructDecl attributeDecls

//getRuleContext
func (s *DeclAnnotationContext) GetRuleContext() antlr.RuleContext {
	return s
}

//getters
func (s *DeclAnnotationContext) JsonValue() IJsonValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*JsonValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IJsonValueContext)
}

func (s *DeclAnnotationContext) SEMI() antlr.TerminalNode {
	return s.GetToken(AdlPSEMI, 0)
}

func (s *DeclAnnotationContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(AdlPID)
}

func (s *DeclAnnotationContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(AdlPID, i)
}

func (s *DeclAnnotationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DeclAnnotationEntryListener); ok {
		listenerT.EnterDeclAnnotation(s)
	}
}

func (s *DeclAnnotationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DeclAnnotationExitListener); ok {
		listenerT.ExitDeclAnnotation(s)
	}
}

func (s *DeclAnnotationContext) VisitFunc(hdls antlr.ParserTreeVisitorHandlers, args ...interface{}) (result interface{}) {
	h := hdls.(*AdlPHandlers)
	if h.EnterEveryRule != nil {
		h.EnterEveryRule(s)
	}
	if h.DeclAnnotation != nil {
		h.DeclAnnotation(s, h, args...)
	} else {
		s.VisitChildrenFunc(h, args...)
	}
	if h.ExitEveryRule != nil {
		h.ExitEveryRule(s)
	}
	return
}

func (s *DeclAnnotationContext) Visit(delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	switch t := delegate.(type) {
	case DeclAnnotationContextVisitor:
		return t.VisitDeclAnnotation(s, delegate, args...)
	default:
		return delegate.VisitChildren(s, delegate, args...)
	}
}

//END AltLabelStructDecl

//Begin AltLabelStructDecl

type FieldAnnotationContext struct {
	*Top_level_statementContext
	//TokenDecl
	kw antlr.Token
	//TokenDecl
	a antlr.Token
	//TokenDecl
	b antlr.Token
	//TokenDecl
	c antlr.Token
}

func NewFieldAnnotationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FieldAnnotationContext {
	var p = new(FieldAnnotationContext)

	p.Top_level_statementContext = NewEmptyTop_level_statementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Top_level_statementContext))

	return p
}

type IFieldAnnotationContext interface {
	//Current rule
	ITop_level_statementContext
	// start internal
	//Gets for raw elements
	//  ruleGetterDecl
	JsonValue() IJsonValueContext
	//  ruleListGetterDecl
	//  ruleListIndexedGetterDecl

	//  tokenGetterDecl
	DCOLON() antlr.TerminalNode
	SEMI() antlr.TerminalNode
	//  tokenListGetterDecl
	AllID() []antlr.TerminalNode
	//  tokenListIndexedGetterDecl
	ID(i int) antlr.TerminalNode
	// end internal
	//
	//Gets for labeled elements
	//  tokenDecls
	GetKw() antlr.Token
	GetA() antlr.Token
	GetB() antlr.Token
	GetC() antlr.Token

	//  tokenTypeDecls
	//  tokenListDecls
	//  ruleContextDecls
	//  ruleContextListDecls
	//  attributeDecls

	// TODO dispatchMethods (needed?)
}

func (*FieldAnnotationContext) IsFieldAnnotationContext() {}

//AltLabelStructDecl tokenDecls
func (s *FieldAnnotationContext) GetKw() antlr.Token  { return s.kw }
func (s *FieldAnnotationContext) SetKw(v antlr.Token) { s.kw = v }

func (s *FieldAnnotationContext) GetA() antlr.Token  { return s.a }
func (s *FieldAnnotationContext) SetA(v antlr.Token) { s.a = v }

func (s *FieldAnnotationContext) GetB() antlr.Token  { return s.b }
func (s *FieldAnnotationContext) SetB(v antlr.Token) { s.b = v }

func (s *FieldAnnotationContext) GetC() antlr.Token  { return s.c }
func (s *FieldAnnotationContext) SetC(v antlr.Token) { s.c = v }

//AltLabelStructDecl tokenTypeDecls

//AltLabelStructDecl tokenListDecls

//AltLabelStructDecl ruleContextDecls

//AltLabelStructDecl ruleContextListDecls

//AltLabelStructDecl attributeDecls

//getRuleContext
func (s *FieldAnnotationContext) GetRuleContext() antlr.RuleContext {
	return s
}

//getters
func (s *FieldAnnotationContext) DCOLON() antlr.TerminalNode {
	return s.GetToken(AdlPDCOLON, 0)
}

func (s *FieldAnnotationContext) JsonValue() IJsonValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*JsonValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IJsonValueContext)
}

func (s *FieldAnnotationContext) SEMI() antlr.TerminalNode {
	return s.GetToken(AdlPSEMI, 0)
}

func (s *FieldAnnotationContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(AdlPID)
}

func (s *FieldAnnotationContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(AdlPID, i)
}

func (s *FieldAnnotationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FieldAnnotationEntryListener); ok {
		listenerT.EnterFieldAnnotation(s)
	}
}

func (s *FieldAnnotationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FieldAnnotationExitListener); ok {
		listenerT.ExitFieldAnnotation(s)
	}
}

func (s *FieldAnnotationContext) VisitFunc(hdls antlr.ParserTreeVisitorHandlers, args ...interface{}) (result interface{}) {
	h := hdls.(*AdlPHandlers)
	if h.EnterEveryRule != nil {
		h.EnterEveryRule(s)
	}
	if h.FieldAnnotation != nil {
		h.FieldAnnotation(s, h, args...)
	} else {
		s.VisitChildrenFunc(h, args...)
	}
	if h.ExitEveryRule != nil {
		h.ExitEveryRule(s)
	}
	return
}

func (s *FieldAnnotationContext) Visit(delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	switch t := delegate.(type) {
	case FieldAnnotationContextVisitor:
		return t.VisitFieldAnnotation(s, delegate, args...)
	default:
		return delegate.VisitChildren(s, delegate, args...)
	}
}

//END AltLabelStructDecl

//Begin AltLabelStructDecl

type TypeOrNewtypeContext struct {
	*Top_level_statementContext
	//TokenDecl
	kw antlr.Token
	//TokenDecl
	a antlr.Token
	//TokenDecl
	b antlr.Token
}

func NewTypeOrNewtypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypeOrNewtypeContext {
	var p = new(TypeOrNewtypeContext)

	p.Top_level_statementContext = NewEmptyTop_level_statementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Top_level_statementContext))

	return p
}

type ITypeOrNewtypeContext interface {
	//Current rule
	ITop_level_statementContext
	// start internal
	//Gets for raw elements
	//  ruleGetterDecl
	TypeParam() ITypeParamContext
	TypeExpr() ITypeExprContext
	JsonValue() IJsonValueContext
	//  ruleListGetterDecl
	AllAnnon() []IAnnonContext
	//  ruleListIndexedGetterDecl
	Annon(i int) IAnnonContext

	//  tokenGetterDecl
	SEMI() antlr.TerminalNode
	//  tokenListGetterDecl
	AllEQ() []antlr.TerminalNode
	AllID() []antlr.TerminalNode
	//  tokenListIndexedGetterDecl
	EQ(i int) antlr.TerminalNode
	ID(i int) antlr.TerminalNode
	// end internal
	//
	//Gets for labeled elements
	//  tokenDecls
	GetKw() antlr.Token
	GetA() antlr.Token
	GetB() antlr.Token

	//  tokenTypeDecls
	//  tokenListDecls
	//  ruleContextDecls
	//  ruleContextListDecls
	//  attributeDecls

	// TODO dispatchMethods (needed?)
}

func (*TypeOrNewtypeContext) IsTypeOrNewtypeContext() {}

//AltLabelStructDecl tokenDecls
func (s *TypeOrNewtypeContext) GetKw() antlr.Token  { return s.kw }
func (s *TypeOrNewtypeContext) SetKw(v antlr.Token) { s.kw = v }

func (s *TypeOrNewtypeContext) GetA() antlr.Token  { return s.a }
func (s *TypeOrNewtypeContext) SetA(v antlr.Token) { s.a = v }

func (s *TypeOrNewtypeContext) GetB() antlr.Token  { return s.b }
func (s *TypeOrNewtypeContext) SetB(v antlr.Token) { s.b = v }

//AltLabelStructDecl tokenTypeDecls

//AltLabelStructDecl tokenListDecls

//AltLabelStructDecl ruleContextDecls

//AltLabelStructDecl ruleContextListDecls

//AltLabelStructDecl attributeDecls

//getRuleContext
func (s *TypeOrNewtypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

//getters
func (s *TypeOrNewtypeContext) AllEQ() []antlr.TerminalNode {
	return s.GetTokens(AdlPEQ)
}

func (s *TypeOrNewtypeContext) EQ(i int) antlr.TerminalNode {
	return s.GetToken(AdlPEQ, i)
}

func (s *TypeOrNewtypeContext) SEMI() antlr.TerminalNode {
	return s.GetToken(AdlPSEMI, 0)
}

func (s *TypeOrNewtypeContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(AdlPID)
}

func (s *TypeOrNewtypeContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(AdlPID, i)
}

func (s *TypeOrNewtypeContext) AllAnnon() []IAnnonContext {
	//ContextRuleListGetterDecl
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*AnnonContext)(nil)).Elem())
	var tst = make([]IAnnonContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAnnonContext)
		}
	}

	return tst
}

func (s *TypeOrNewtypeContext) Annon(i int) IAnnonContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*AnnonContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAnnonContext)
}

func (s *TypeOrNewtypeContext) TypeParam() ITypeParamContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*TypeParamContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeParamContext)
}

func (s *TypeOrNewtypeContext) TypeExpr() ITypeExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*TypeExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeExprContext)
}

func (s *TypeOrNewtypeContext) JsonValue() IJsonValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*JsonValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IJsonValueContext)
}

func (s *TypeOrNewtypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TypeOrNewtypeEntryListener); ok {
		listenerT.EnterTypeOrNewtype(s)
	}
}

func (s *TypeOrNewtypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TypeOrNewtypeExitListener); ok {
		listenerT.ExitTypeOrNewtype(s)
	}
}

func (s *TypeOrNewtypeContext) VisitFunc(hdls antlr.ParserTreeVisitorHandlers, args ...interface{}) (result interface{}) {
	h := hdls.(*AdlPHandlers)
	if h.EnterEveryRule != nil {
		h.EnterEveryRule(s)
	}
	if h.TypeOrNewtype != nil {
		h.TypeOrNewtype(s, h, args...)
	} else {
		s.VisitChildrenFunc(h, args...)
	}
	if h.ExitEveryRule != nil {
		h.ExitEveryRule(s)
	}
	return
}

func (s *TypeOrNewtypeContext) Visit(delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	switch t := delegate.(type) {
	case TypeOrNewtypeContextVisitor:
		return t.VisitTypeOrNewtype(s, delegate, args...)
	default:
		return delegate.VisitChildren(s, delegate, args...)
	}
}

//END AltLabelStructDecl

//Begin AltLabelStructDecl

type ModuleAnnotationContext struct {
	*Top_level_statementContext
	//TokenDecl
	kw antlr.Token
	//TokenDecl
	a antlr.Token
}

func NewModuleAnnotationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ModuleAnnotationContext {
	var p = new(ModuleAnnotationContext)

	p.Top_level_statementContext = NewEmptyTop_level_statementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Top_level_statementContext))

	return p
}

type IModuleAnnotationContext interface {
	//Current rule
	ITop_level_statementContext
	// start internal
	//Gets for raw elements
	//  ruleGetterDecl
	JsonValue() IJsonValueContext
	//  ruleListGetterDecl
	//  ruleListIndexedGetterDecl

	//  tokenGetterDecl
	SEMI() antlr.TerminalNode
	//  tokenListGetterDecl
	AllID() []antlr.TerminalNode
	//  tokenListIndexedGetterDecl
	ID(i int) antlr.TerminalNode
	// end internal
	//
	//Gets for labeled elements
	//  tokenDecls
	GetKw() antlr.Token
	GetA() antlr.Token

	//  tokenTypeDecls
	//  tokenListDecls
	//  ruleContextDecls
	//  ruleContextListDecls
	//  attributeDecls

	// TODO dispatchMethods (needed?)
}

func (*ModuleAnnotationContext) IsModuleAnnotationContext() {}

//AltLabelStructDecl tokenDecls
func (s *ModuleAnnotationContext) GetKw() antlr.Token  { return s.kw }
func (s *ModuleAnnotationContext) SetKw(v antlr.Token) { s.kw = v }

func (s *ModuleAnnotationContext) GetA() antlr.Token  { return s.a }
func (s *ModuleAnnotationContext) SetA(v antlr.Token) { s.a = v }

//AltLabelStructDecl tokenTypeDecls

//AltLabelStructDecl tokenListDecls

//AltLabelStructDecl ruleContextDecls

//AltLabelStructDecl ruleContextListDecls

//AltLabelStructDecl attributeDecls

//getRuleContext
func (s *ModuleAnnotationContext) GetRuleContext() antlr.RuleContext {
	return s
}

//getters
func (s *ModuleAnnotationContext) JsonValue() IJsonValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*JsonValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IJsonValueContext)
}

func (s *ModuleAnnotationContext) SEMI() antlr.TerminalNode {
	return s.GetToken(AdlPSEMI, 0)
}

func (s *ModuleAnnotationContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(AdlPID)
}

func (s *ModuleAnnotationContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(AdlPID, i)
}

func (s *ModuleAnnotationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ModuleAnnotationEntryListener); ok {
		listenerT.EnterModuleAnnotation(s)
	}
}

func (s *ModuleAnnotationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ModuleAnnotationExitListener); ok {
		listenerT.ExitModuleAnnotation(s)
	}
}

func (s *ModuleAnnotationContext) VisitFunc(hdls antlr.ParserTreeVisitorHandlers, args ...interface{}) (result interface{}) {
	h := hdls.(*AdlPHandlers)
	if h.EnterEveryRule != nil {
		h.EnterEveryRule(s)
	}
	if h.ModuleAnnotation != nil {
		h.ModuleAnnotation(s, h, args...)
	} else {
		s.VisitChildrenFunc(h, args...)
	}
	if h.ExitEveryRule != nil {
		h.ExitEveryRule(s)
	}
	return
}

func (s *ModuleAnnotationContext) Visit(delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	switch t := delegate.(type) {
	case ModuleAnnotationContextVisitor:
		return t.VisitModuleAnnotation(s, delegate, args...)
	default:
		return delegate.VisitChildren(s, delegate, args...)
	}
}

//END AltLabelStructDecl

func (p *AdlP) Top_level_statement() (localctx ITop_level_statementContext) {
	localctx = NewTop_level_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, AdlPRULE_top_level_statement)
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

	p.SetState(137)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext()) {
	case 1:
		localctx = NewStructOrUnionContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		p.SetState(80)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == AdlPAT || _la == AdlPLINE_DOC {
			{
				p.SetState(77)
				p.Annon()
			}

			p.SetState(82)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(83)
			var _m = p.Match(AdlPID)
			localctx.(*StructOrUnionContext).kw = _m

		}
		{
			p.SetState(84)
			var _m = p.Match(AdlPID)
			localctx.(*StructOrUnionContext).a = _m

		}
		p.SetState(86)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == AdlPLCHEVR {
			{
				p.SetState(85)
				p.TypeParam()
			}

		}
		{
			p.SetState(88)
			p.Match(AdlPLCUR)
		}
		p.SetState(92)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AdlPAT)|(1<<AdlPID)|(1<<AdlPLINE_DOC))) != 0 {
			{
				p.SetState(89)
				p.SoruBody()
			}

			p.SetState(94)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(95)
			p.Match(AdlPRCUR)
		}
		{
			p.SetState(96)
			p.Match(AdlPSEMI)
		}

	case 2:
		localctx = NewTypeOrNewtypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		p.SetState(100)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == AdlPAT || _la == AdlPLINE_DOC {
			{
				p.SetState(97)
				p.Annon()
			}

			p.SetState(102)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(103)
			var _m = p.Match(AdlPID)
			localctx.(*TypeOrNewtypeContext).kw = _m

		}
		{
			p.SetState(104)
			var _m = p.Match(AdlPID)
			localctx.(*TypeOrNewtypeContext).a = _m

		}
		p.SetState(106)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == AdlPLCHEVR {
			{
				p.SetState(105)
				p.TypeParam()
			}

		}
		{
			p.SetState(108)
			p.Match(AdlPEQ)
		}
		{
			p.SetState(109)
			var _m = p.Match(AdlPID)
			localctx.(*TypeOrNewtypeContext).b = _m

		}
		p.SetState(111)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == AdlPLCHEVR {
			{
				p.SetState(110)
				p.TypeExpr()
			}

		}
		p.SetState(115)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == AdlPEQ {
			{
				p.SetState(113)
				p.Match(AdlPEQ)
			}
			{
				p.SetState(114)
				p.JsonValue()
			}

		}
		{
			p.SetState(117)
			p.Match(AdlPSEMI)
		}

	case 3:
		localctx = NewModuleAnnotationContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(118)
			var _m = p.Match(AdlPID)
			localctx.(*ModuleAnnotationContext).kw = _m

		}
		{
			p.SetState(119)
			var _m = p.Match(AdlPID)
			localctx.(*ModuleAnnotationContext).a = _m

		}
		{
			p.SetState(120)
			p.JsonValue()
		}
		{
			p.SetState(121)
			p.Match(AdlPSEMI)
		}

	case 4:
		localctx = NewDeclAnnotationContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(123)
			var _m = p.Match(AdlPID)
			localctx.(*DeclAnnotationContext).kw = _m

		}
		{
			p.SetState(124)
			var _m = p.Match(AdlPID)
			localctx.(*DeclAnnotationContext).a = _m

		}
		{
			p.SetState(125)
			var _m = p.Match(AdlPID)
			localctx.(*DeclAnnotationContext).b = _m

		}
		{
			p.SetState(126)
			p.JsonValue()
		}
		{
			p.SetState(127)
			p.Match(AdlPSEMI)
		}

	case 5:
		localctx = NewFieldAnnotationContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(129)
			var _m = p.Match(AdlPID)
			localctx.(*FieldAnnotationContext).kw = _m

		}
		{
			p.SetState(130)
			var _m = p.Match(AdlPID)
			localctx.(*FieldAnnotationContext).a = _m

		}
		{
			p.SetState(131)
			p.Match(AdlPDCOLON)
		}
		{
			p.SetState(132)
			var _m = p.Match(AdlPID)
			localctx.(*FieldAnnotationContext).b = _m

		}
		{
			p.SetState(133)
			var _m = p.Match(AdlPID)
			localctx.(*FieldAnnotationContext).c = _m

		}
		{
			p.SetState(134)
			p.JsonValue()
		}
		{
			p.SetState(135)
			p.Match(AdlPSEMI)
		}

	}

	return localctx
}

type ITypeParamContext interface {
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

	// IsTypeParamContext differentiates from other interfaces.
	//copyStruct,GetRuleContext and ToStringTree  from embedded

	//<if(dispatchMethods)>
	//<dispatchMethods; separator="\n\n">
	//<endif>

	//<if(extensionMembers)>
	//<extensionMembers; separator="\n\n">
	//<endif>
}

type TypeParamContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeParamContext() *TypeParamContext {
	var p = new(TypeParamContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AdlPRULE_typeParam
	return p
}

func (*TypeParamContext) IsTypeParamContext() {}

func NewTypeParamContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeParamContext {
	var p = new(TypeParamContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AdlPRULE_typeParam

	return p
}

func (s *TypeParamContext) GetParser() antlr.Parser { return s.parser }

//StructDecl tokenDecls

//StructDecl tokenTypeDecls

//StructDecl tokenListDecls

//StructDecl ruleContextDecls

//StructDecl ruleContextListDecls

//StructDecl attributeDecls

// Getters

//provideCopyFrom
func (s *TypeParamContext) CopyFrom(ctx *TypeParamContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *TypeParamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeParamContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

//dispatchMethods
//extensionMembers

//Begin AltLabelStructDecl

type TypeParameterContext struct {
	*TypeParamContext
	//TokenDecl
	_ID antlr.Token
	//TokenListDecl
	typep []antlr.Token
}

func NewTypeParameterContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypeParameterContext {
	var p = new(TypeParameterContext)

	p.TypeParamContext = NewEmptyTypeParamContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TypeParamContext))

	return p
}

type ITypeParameterContext interface {
	//Current rule
	ITypeParamContext
	// start internal
	//Gets for raw elements
	//  ruleGetterDecl
	//  ruleListGetterDecl
	//  ruleListIndexedGetterDecl

	//  tokenGetterDecl
	LCHEVR() antlr.TerminalNode
	RCHEVR() antlr.TerminalNode
	//  tokenListGetterDecl
	AllID() []antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	//  tokenListIndexedGetterDecl
	ID(i int) antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode
	// end internal
	//
	//Gets for labeled elements
	//  tokenDecls

	//  tokenTypeDecls
	//  tokenListDecls
	GetTypep() []antlr.Token
	//  ruleContextDecls
	//  ruleContextListDecls
	//  attributeDecls

	// TODO dispatchMethods (needed?)
}

func (*TypeParameterContext) IsTypeParameterContext() {}

//AltLabelStructDecl tokenDecls
func (s *TypeParameterContext) Get_ID() antlr.Token  { return s._ID }
func (s *TypeParameterContext) Set_ID(v antlr.Token) { s._ID = v }

//AltLabelStructDecl tokenTypeDecls

//AltLabelStructDecl tokenListDecls
func (s *TypeParameterContext) GetTypep() []antlr.Token  { return s.typep }
func (s *TypeParameterContext) SetTypep(v []antlr.Token) { s.typep = v }

//AltLabelStructDecl ruleContextDecls

//AltLabelStructDecl ruleContextListDecls

//AltLabelStructDecl attributeDecls

//getRuleContext
func (s *TypeParameterContext) GetRuleContext() antlr.RuleContext {
	return s
}

//getters
func (s *TypeParameterContext) LCHEVR() antlr.TerminalNode {
	return s.GetToken(AdlPLCHEVR, 0)
}

func (s *TypeParameterContext) RCHEVR() antlr.TerminalNode {
	return s.GetToken(AdlPRCHEVR, 0)
}

func (s *TypeParameterContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(AdlPID)
}

func (s *TypeParameterContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(AdlPID, i)
}

func (s *TypeParameterContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(AdlPCOMMA)
}

func (s *TypeParameterContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(AdlPCOMMA, i)
}

func (s *TypeParameterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TypeParameterEntryListener); ok {
		listenerT.EnterTypeParameter(s)
	}
}

func (s *TypeParameterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TypeParameterExitListener); ok {
		listenerT.ExitTypeParameter(s)
	}
}

func (s *TypeParameterContext) VisitFunc(hdls antlr.ParserTreeVisitorHandlers, args ...interface{}) (result interface{}) {
	h := hdls.(*AdlPHandlers)
	if h.EnterEveryRule != nil {
		h.EnterEveryRule(s)
	}
	if h.TypeParameter != nil {
		h.TypeParameter(s, h, args...)
	} else {
		s.VisitChildrenFunc(h, args...)
	}
	if h.ExitEveryRule != nil {
		h.ExitEveryRule(s)
	}
	return
}

func (s *TypeParameterContext) Visit(delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	switch t := delegate.(type) {
	case TypeParameterContextVisitor:
		return t.VisitTypeParameter(s, delegate, args...)
	default:
		return delegate.VisitChildren(s, delegate, args...)
	}
}

//END AltLabelStructDecl

//Begin AltLabelStructDecl

type ErrorTypeParamContext struct {
	*TypeParamContext
}

func NewErrorTypeParamContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ErrorTypeParamContext {
	var p = new(ErrorTypeParamContext)

	p.TypeParamContext = NewEmptyTypeParamContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TypeParamContext))

	return p
}

type IErrorTypeParamContext interface {
	//Current rule
	ITypeParamContext
	// start internal
	//Gets for raw elements
	//  ruleGetterDecl
	//  ruleListGetterDecl
	AllTypeParamError() []ITypeParamErrorContext
	//  ruleListIndexedGetterDecl
	TypeParamError(i int) ITypeParamErrorContext

	//  tokenGetterDecl
	LCHEVR() antlr.TerminalNode
	RCHEVR() antlr.TerminalNode
	//  tokenListGetterDecl
	AllCOMMA() []antlr.TerminalNode
	AllID() []antlr.TerminalNode
	//  tokenListIndexedGetterDecl
	COMMA(i int) antlr.TerminalNode
	ID(i int) antlr.TerminalNode
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

func (*ErrorTypeParamContext) IsErrorTypeParamContext() {}

//AltLabelStructDecl tokenDecls

//AltLabelStructDecl tokenTypeDecls

//AltLabelStructDecl tokenListDecls

//AltLabelStructDecl ruleContextDecls

//AltLabelStructDecl ruleContextListDecls

//AltLabelStructDecl attributeDecls

//getRuleContext
func (s *ErrorTypeParamContext) GetRuleContext() antlr.RuleContext {
	return s
}

//getters
func (s *ErrorTypeParamContext) LCHEVR() antlr.TerminalNode {
	return s.GetToken(AdlPLCHEVR, 0)
}

func (s *ErrorTypeParamContext) AllTypeParamError() []ITypeParamErrorContext {
	//ContextRuleListGetterDecl
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*TypeParamErrorContext)(nil)).Elem())
	var tst = make([]ITypeParamErrorContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITypeParamErrorContext)
		}
	}

	return tst
}

func (s *ErrorTypeParamContext) TypeParamError(i int) ITypeParamErrorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*TypeParamErrorContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITypeParamErrorContext)
}

func (s *ErrorTypeParamContext) RCHEVR() antlr.TerminalNode {
	return s.GetToken(AdlPRCHEVR, 0)
}

func (s *ErrorTypeParamContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(AdlPCOMMA)
}

func (s *ErrorTypeParamContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(AdlPCOMMA, i)
}

func (s *ErrorTypeParamContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(AdlPID)
}

func (s *ErrorTypeParamContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(AdlPID, i)
}

func (s *ErrorTypeParamContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ErrorTypeParamEntryListener); ok {
		listenerT.EnterErrorTypeParam(s)
	}
}

func (s *ErrorTypeParamContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ErrorTypeParamExitListener); ok {
		listenerT.ExitErrorTypeParam(s)
	}
}

func (s *ErrorTypeParamContext) VisitFunc(hdls antlr.ParserTreeVisitorHandlers, args ...interface{}) (result interface{}) {
	h := hdls.(*AdlPHandlers)
	if h.EnterEveryRule != nil {
		h.EnterEveryRule(s)
	}
	if h.ErrorTypeParam != nil {
		h.ErrorTypeParam(s, h, args...)
	} else {
		s.VisitChildrenFunc(h, args...)
	}
	if h.ExitEveryRule != nil {
		h.ExitEveryRule(s)
	}
	return
}

func (s *ErrorTypeParamContext) Visit(delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	switch t := delegate.(type) {
	case ErrorTypeParamContextVisitor:
		return t.VisitErrorTypeParam(s, delegate, args...)
	default:
		return delegate.VisitChildren(s, delegate, args...)
	}
}

//END AltLabelStructDecl

func (p *AdlP) TypeParam() (localctx ITypeParamContext) {
	localctx = NewTypeParamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, AdlPRULE_typeParam)
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

	p.SetState(170)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext()) {
	case 1:
		localctx = NewTypeParameterContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(139)
			p.Match(AdlPLCHEVR)
		}
		{
			p.SetState(140)
			var _m = p.Match(AdlPID)
			localctx.(*TypeParameterContext)._ID = _m

		}
		localctx.(*TypeParameterContext).typep = append(localctx.(*TypeParameterContext).typep, localctx.(*TypeParameterContext)._ID)
		p.SetState(145)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == AdlPCOMMA {
			{
				p.SetState(141)
				p.Match(AdlPCOMMA)
			}
			{
				p.SetState(142)
				var _m = p.Match(AdlPID)
				localctx.(*TypeParameterContext)._ID = _m

			}
			localctx.(*TypeParameterContext).typep = append(localctx.(*TypeParameterContext).typep, localctx.(*TypeParameterContext)._ID)

			p.SetState(147)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(148)
			p.Match(AdlPRCHEVR)
		}

	case 2:
		localctx = NewErrorTypeParamContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(149)
			p.Match(AdlPLCHEVR)
		}
		{
			p.SetState(150)
			p.TypeParamError()
		}
		p.SetState(155)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == AdlPCOMMA {
			{
				p.SetState(151)
				p.Match(AdlPCOMMA)
			}
			{
				p.SetState(152)
				p.Match(AdlPID)
			}

			p.SetState(157)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(158)
			p.Match(AdlPRCHEVR)
		}

	case 3:
		localctx = NewErrorTypeParamContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(160)
			p.Match(AdlPLCHEVR)
		}
		{
			p.SetState(161)
			p.Match(AdlPID)
		}
		p.SetState(164)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == AdlPCOMMA {
			{
				p.SetState(162)
				p.Match(AdlPCOMMA)
			}
			{
				p.SetState(163)
				p.TypeParamError()
			}

			p.SetState(166)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(168)
			p.Match(AdlPRCHEVR)
		}

	}

	return localctx
}

type ITypeParamErrorContext interface {
	antlr.ParserRuleContext
	// GetParser returns the parser.
	GetParser() antlr.Parser

	// start internal
	//  ruleGetterDecl
	//  ruleListGetterDecl
	AllTypeParamError() []ITypeParamErrorContext
	//  ruleListIndexedGetterDecl
	TypeParamError(i int) ITypeParamErrorContext
	//  ruleContextDecls
	//  ruleContextListDecls
	// end internal
	//Gets for labeled elements
	//tokenDecls
	//tokenTypeDecls
	//tokenListDecls
	//attributeDecls
	//tokenGetterDecl
	LCHEVR() antlr.TerminalNode
	RCHEVR() antlr.TerminalNode
	//tokenListGetterDecl
	AllID() []antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	//tokenListIndexedGetterDecl
	ID(i int) antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsTypeParamErrorContext differentiates from other interfaces.
	//copyStruct,GetRuleContext and ToStringTree  from embedded

	//<if(dispatchMethods)>
	//<dispatchMethods; separator="\n\n">
	//<endif>

	//<if(extensionMembers)>
	//<extensionMembers; separator="\n\n">
	//<endif>
}

type TypeParamErrorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeParamErrorContext() *TypeParamErrorContext {
	var p = new(TypeParamErrorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AdlPRULE_typeParamError
	return p
}

func (*TypeParamErrorContext) IsTypeParamErrorContext() {}

func NewTypeParamErrorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeParamErrorContext {
	var p = new(TypeParamErrorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AdlPRULE_typeParamError

	return p
}

func (s *TypeParamErrorContext) GetParser() antlr.Parser { return s.parser }

//StructDecl tokenDecls

//StructDecl tokenTypeDecls

//StructDecl tokenListDecls

//StructDecl ruleContextDecls

//StructDecl ruleContextListDecls

//StructDecl attributeDecls

// Getters
func (s *TypeParamErrorContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(AdlPID)
}

func (s *TypeParamErrorContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(AdlPID, i)
}

func (s *TypeParamErrorContext) LCHEVR() antlr.TerminalNode {
	return s.GetToken(AdlPLCHEVR, 0)
}

func (s *TypeParamErrorContext) RCHEVR() antlr.TerminalNode {
	return s.GetToken(AdlPRCHEVR, 0)
}

func (s *TypeParamErrorContext) AllTypeParamError() []ITypeParamErrorContext {
	//ContextRuleListGetterDecl
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*TypeParamErrorContext)(nil)).Elem())
	var tst = make([]ITypeParamErrorContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITypeParamErrorContext)
		}
	}

	return tst
}

func (s *TypeParamErrorContext) TypeParamError(i int) ITypeParamErrorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*TypeParamErrorContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITypeParamErrorContext)
}

func (s *TypeParamErrorContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(AdlPCOMMA)
}

func (s *TypeParamErrorContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(AdlPCOMMA, i)
}

//provideCopyFrom
func (s *TypeParamErrorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeParamErrorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

//dispatchMethods
func (s *TypeParamErrorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TypeParamErrorEntryListener); ok {
		listenerT.EnterTypeParamError(s)
	}
}

func (s *TypeParamErrorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TypeParamErrorExitListener); ok {
		listenerT.ExitTypeParamError(s)
	}
}

func (s *TypeParamErrorContext) VisitFunc(hdls antlr.ParserTreeVisitorHandlers, args ...interface{}) (result interface{}) {
	h := hdls.(*AdlPHandlers)
	if h.EnterEveryRule != nil {
		h.EnterEveryRule(s)
	}
	if h.TypeParamError != nil {
		h.TypeParamError(s, h, args...)
	} else {
		s.VisitChildrenFunc(h, args...)
	}
	if h.ExitEveryRule != nil {
		h.ExitEveryRule(s)
	}
	return
}

func (s *TypeParamErrorContext) Visit(delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	switch t := delegate.(type) {
	case TypeParamErrorContextVisitor:
		return t.VisitTypeParamError(s, delegate, args...)
	default:
		return delegate.VisitChildren(s, delegate, args...)
	}
}

//extensionMembers

func (p *AdlP) TypeParamError() (localctx ITypeParamErrorContext) {
	localctx = NewTypeParamErrorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, AdlPRULE_typeParamError)
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
		p.SetState(172)
		p.Match(AdlPID)
	}
	{
		p.SetState(173)
		p.Match(AdlPLCHEVR)
	}
	p.SetState(176)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(174)
			p.Match(AdlPID)
		}

	case 2:
		{
			p.SetState(175)
			p.TypeParamError()
		}

	}
	p.SetState(185)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == AdlPCOMMA {
		{
			p.SetState(178)
			p.Match(AdlPCOMMA)
		}
		p.SetState(181)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(179)
				p.Match(AdlPID)
			}

		case 2:
			{
				p.SetState(180)
				p.TypeParamError()
			}

		}

		p.SetState(187)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(188)
		p.Match(AdlPRCHEVR)
	}

	return localctx
}

type ITypeExprContext interface {
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
}

func NewEmptyTypeExprContext() *TypeExprContext {
	var p = new(TypeExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AdlPRULE_typeExpr
	return p
}

func (*TypeExprContext) IsTypeExprContext() {}

func NewTypeExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeExprContext {
	var p = new(TypeExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AdlPRULE_typeExpr

	return p
}

func (s *TypeExprContext) GetParser() antlr.Parser { return s.parser }

//StructDecl tokenDecls

//StructDecl tokenTypeDecls

//StructDecl tokenListDecls

//StructDecl ruleContextDecls

//StructDecl ruleContextListDecls

//StructDecl attributeDecls

// Getters

//provideCopyFrom
func (s *TypeExprContext) CopyFrom(ctx *TypeExprContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *TypeExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

//dispatchMethods
//extensionMembers

//Begin AltLabelStructDecl

type TypeExpressionContext struct {
	*TypeExprContext
	//RuleContextDecl
	_typeExprElem ITypeExprElemContext
	//RuleContextListDecl
	typep []ITypeExprElemContext
}

func NewTypeExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypeExpressionContext {
	var p = new(TypeExpressionContext)

	p.TypeExprContext = NewEmptyTypeExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TypeExprContext))

	return p
}

type ITypeExpressionContext interface {
	//Current rule
	ITypeExprContext
	// start internal
	//Gets for raw elements
	//  ruleGetterDecl
	//  ruleListGetterDecl
	AllTypeExprElem() []ITypeExprElemContext
	//  ruleListIndexedGetterDecl
	TypeExprElem(i int) ITypeExprElemContext

	//  tokenGetterDecl
	LCHEVR() antlr.TerminalNode
	RCHEVR() antlr.TerminalNode
	//  tokenListGetterDecl
	AllCOMMA() []antlr.TerminalNode
	//  tokenListIndexedGetterDecl
	COMMA(i int) antlr.TerminalNode
	// end internal
	//
	//Gets for labeled elements
	//  tokenDecls
	//  tokenTypeDecls
	//  tokenListDecls
	//  ruleContextDecls

	//  ruleContextListDecls
	GetTypep() []ITypeExprElemContext

	//  attributeDecls

	// TODO dispatchMethods (needed?)
}

func (*TypeExpressionContext) IsTypeExpressionContext() {}

//AltLabelStructDecl tokenDecls

//AltLabelStructDecl tokenTypeDecls

//AltLabelStructDecl tokenListDecls

//AltLabelStructDecl ruleContextDecls
func (s *TypeExpressionContext) Get_typeExprElem() ITypeExprElemContext  { return s._typeExprElem }
func (s *TypeExpressionContext) Set_typeExprElem(v ITypeExprElemContext) { s._typeExprElem = v }

//AltLabelStructDecl ruleContextListDecls
func (s *TypeExpressionContext) GetTypep() []ITypeExprElemContext  { return s.typep }
func (s *TypeExpressionContext) SetTypep(v []ITypeExprElemContext) { s.typep = v }

//AltLabelStructDecl attributeDecls

//getRuleContext
func (s *TypeExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

//getters
func (s *TypeExpressionContext) LCHEVR() antlr.TerminalNode {
	return s.GetToken(AdlPLCHEVR, 0)
}

func (s *TypeExpressionContext) RCHEVR() antlr.TerminalNode {
	return s.GetToken(AdlPRCHEVR, 0)
}

func (s *TypeExpressionContext) AllTypeExprElem() []ITypeExprElemContext {
	//ContextRuleListGetterDecl
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*TypeExprElemContext)(nil)).Elem())
	var tst = make([]ITypeExprElemContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITypeExprElemContext)
		}
	}

	return tst
}

func (s *TypeExpressionContext) TypeExprElem(i int) ITypeExprElemContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*TypeExprElemContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITypeExprElemContext)
}

func (s *TypeExpressionContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(AdlPCOMMA)
}

func (s *TypeExpressionContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(AdlPCOMMA, i)
}

func (s *TypeExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TypeExpressionEntryListener); ok {
		listenerT.EnterTypeExpression(s)
	}
}

func (s *TypeExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TypeExpressionExitListener); ok {
		listenerT.ExitTypeExpression(s)
	}
}

func (s *TypeExpressionContext) VisitFunc(hdls antlr.ParserTreeVisitorHandlers, args ...interface{}) (result interface{}) {
	h := hdls.(*AdlPHandlers)
	if h.EnterEveryRule != nil {
		h.EnterEveryRule(s)
	}
	if h.TypeExpression != nil {
		h.TypeExpression(s, h, args...)
	} else {
		s.VisitChildrenFunc(h, args...)
	}
	if h.ExitEveryRule != nil {
		h.ExitEveryRule(s)
	}
	return
}

func (s *TypeExpressionContext) Visit(delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	switch t := delegate.(type) {
	case TypeExpressionContextVisitor:
		return t.VisitTypeExpression(s, delegate, args...)
	default:
		return delegate.VisitChildren(s, delegate, args...)
	}
}

//END AltLabelStructDecl

func (p *AdlP) TypeExpr() (localctx ITypeExprContext) {
	localctx = NewTypeExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, AdlPRULE_typeExpr)
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

	localctx = NewTypeExpressionContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(190)
		p.Match(AdlPLCHEVR)
	}
	{
		p.SetState(191)

		var _x = p.TypeExprElem()

		localctx.(*TypeExpressionContext)._typeExprElem = _x

	}
	localctx.(*TypeExpressionContext).typep = append(localctx.(*TypeExpressionContext).typep, localctx.(*TypeExpressionContext)._typeExprElem)
	p.SetState(196)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == AdlPCOMMA {
		{
			p.SetState(192)
			p.Match(AdlPCOMMA)
		}
		{
			p.SetState(193)

			var _x = p.TypeExprElem()

			localctx.(*TypeExpressionContext)._typeExprElem = _x

		}
		localctx.(*TypeExpressionContext).typep = append(localctx.(*TypeExpressionContext).typep, localctx.(*TypeExpressionContext)._typeExprElem)

		p.SetState(198)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(199)
		p.Match(AdlPRCHEVR)
	}

	return localctx
}

type ITypeExprElemContext interface {
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

	// IsTypeExprElemContext differentiates from other interfaces.
	//copyStruct,GetRuleContext and ToStringTree  from embedded

	//<if(dispatchMethods)>
	//<dispatchMethods; separator="\n\n">
	//<endif>

	//<if(extensionMembers)>
	//<extensionMembers; separator="\n\n">
	//<endif>
}

type TypeExprElemContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeExprElemContext() *TypeExprElemContext {
	var p = new(TypeExprElemContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AdlPRULE_typeExprElem
	return p
}

func (*TypeExprElemContext) IsTypeExprElemContext() {}

func NewTypeExprElemContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeExprElemContext {
	var p = new(TypeExprElemContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AdlPRULE_typeExprElem

	return p
}

func (s *TypeExprElemContext) GetParser() antlr.Parser { return s.parser }

//StructDecl tokenDecls

//StructDecl tokenTypeDecls

//StructDecl tokenListDecls

//StructDecl ruleContextDecls

//StructDecl ruleContextListDecls

//StructDecl attributeDecls

// Getters

//provideCopyFrom
func (s *TypeExprElemContext) CopyFrom(ctx *TypeExprElemContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *TypeExprElemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeExprElemContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

//dispatchMethods
//extensionMembers

//Begin AltLabelStructDecl

type TypeExpressionElemContext struct {
	*TypeExprElemContext
	//TokenDecl
	a antlr.Token
	//RuleContextDecl
	_typeExprElem ITypeExprElemContext
	//RuleContextListDecl
	typep []ITypeExprElemContext
}

func NewTypeExpressionElemContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypeExpressionElemContext {
	var p = new(TypeExpressionElemContext)

	p.TypeExprElemContext = NewEmptyTypeExprElemContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TypeExprElemContext))

	return p
}

type ITypeExpressionElemContext interface {
	//Current rule
	ITypeExprElemContext
	// start internal
	//Gets for raw elements
	//  ruleGetterDecl
	//  ruleListGetterDecl
	AllTypeExprElem() []ITypeExprElemContext
	//  ruleListIndexedGetterDecl
	TypeExprElem(i int) ITypeExprElemContext

	//  tokenGetterDecl
	ID() antlr.TerminalNode
	LCHEVR() antlr.TerminalNode
	RCHEVR() antlr.TerminalNode
	//  tokenListGetterDecl
	AllCOMMA() []antlr.TerminalNode
	//  tokenListIndexedGetterDecl
	COMMA(i int) antlr.TerminalNode
	// end internal
	//
	//Gets for labeled elements
	//  tokenDecls
	GetA() antlr.Token

	//  tokenTypeDecls
	//  tokenListDecls
	//  ruleContextDecls

	//  ruleContextListDecls
	GetTypep() []ITypeExprElemContext

	//  attributeDecls

	// TODO dispatchMethods (needed?)
}

func (*TypeExpressionElemContext) IsTypeExpressionElemContext() {}

//AltLabelStructDecl tokenDecls
func (s *TypeExpressionElemContext) GetA() antlr.Token  { return s.a }
func (s *TypeExpressionElemContext) SetA(v antlr.Token) { s.a = v }

//AltLabelStructDecl tokenTypeDecls

//AltLabelStructDecl tokenListDecls

//AltLabelStructDecl ruleContextDecls
func (s *TypeExpressionElemContext) Get_typeExprElem() ITypeExprElemContext  { return s._typeExprElem }
func (s *TypeExpressionElemContext) Set_typeExprElem(v ITypeExprElemContext) { s._typeExprElem = v }

//AltLabelStructDecl ruleContextListDecls
func (s *TypeExpressionElemContext) GetTypep() []ITypeExprElemContext  { return s.typep }
func (s *TypeExpressionElemContext) SetTypep(v []ITypeExprElemContext) { s.typep = v }

//AltLabelStructDecl attributeDecls

//getRuleContext
func (s *TypeExpressionElemContext) GetRuleContext() antlr.RuleContext {
	return s
}

//getters
func (s *TypeExpressionElemContext) ID() antlr.TerminalNode {
	return s.GetToken(AdlPID, 0)
}

func (s *TypeExpressionElemContext) LCHEVR() antlr.TerminalNode {
	return s.GetToken(AdlPLCHEVR, 0)
}

func (s *TypeExpressionElemContext) RCHEVR() antlr.TerminalNode {
	return s.GetToken(AdlPRCHEVR, 0)
}

func (s *TypeExpressionElemContext) AllTypeExprElem() []ITypeExprElemContext {
	//ContextRuleListGetterDecl
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*TypeExprElemContext)(nil)).Elem())
	var tst = make([]ITypeExprElemContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITypeExprElemContext)
		}
	}

	return tst
}

func (s *TypeExpressionElemContext) TypeExprElem(i int) ITypeExprElemContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*TypeExprElemContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITypeExprElemContext)
}

func (s *TypeExpressionElemContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(AdlPCOMMA)
}

func (s *TypeExpressionElemContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(AdlPCOMMA, i)
}

func (s *TypeExpressionElemContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TypeExpressionElemEntryListener); ok {
		listenerT.EnterTypeExpressionElem(s)
	}
}

func (s *TypeExpressionElemContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TypeExpressionElemExitListener); ok {
		listenerT.ExitTypeExpressionElem(s)
	}
}

func (s *TypeExpressionElemContext) VisitFunc(hdls antlr.ParserTreeVisitorHandlers, args ...interface{}) (result interface{}) {
	h := hdls.(*AdlPHandlers)
	if h.EnterEveryRule != nil {
		h.EnterEveryRule(s)
	}
	if h.TypeExpressionElem != nil {
		h.TypeExpressionElem(s, h, args...)
	} else {
		s.VisitChildrenFunc(h, args...)
	}
	if h.ExitEveryRule != nil {
		h.ExitEveryRule(s)
	}
	return
}

func (s *TypeExpressionElemContext) Visit(delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	switch t := delegate.(type) {
	case TypeExpressionElemContextVisitor:
		return t.VisitTypeExpressionElem(s, delegate, args...)
	default:
		return delegate.VisitChildren(s, delegate, args...)
	}
}

//END AltLabelStructDecl

func (p *AdlP) TypeExprElem() (localctx ITypeExprElemContext) {
	localctx = NewTypeExprElemContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, AdlPRULE_typeExprElem)
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

	localctx = NewTypeExpressionElemContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(201)
		var _m = p.Match(AdlPID)
		localctx.(*TypeExpressionElemContext).a = _m

	}
	p.SetState(213)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == AdlPLCHEVR {
		{
			p.SetState(202)
			p.Match(AdlPLCHEVR)
		}
		{
			p.SetState(203)

			var _x = p.TypeExprElem()

			localctx.(*TypeExpressionElemContext)._typeExprElem = _x

		}
		localctx.(*TypeExpressionElemContext).typep = append(localctx.(*TypeExpressionElemContext).typep, localctx.(*TypeExpressionElemContext)._typeExprElem)
		p.SetState(208)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == AdlPCOMMA {
			{
				p.SetState(204)
				p.Match(AdlPCOMMA)
			}
			{
				p.SetState(205)

				var _x = p.TypeExprElem()

				localctx.(*TypeExpressionElemContext)._typeExprElem = _x

			}
			localctx.(*TypeExpressionElemContext).typep = append(localctx.(*TypeExpressionElemContext).typep, localctx.(*TypeExpressionElemContext)._typeExprElem)

			p.SetState(210)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(211)
			p.Match(AdlPRCHEVR)
		}

	}

	return localctx
}

type ISoruBodyContext interface {
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

	// IsSoruBodyContext differentiates from other interfaces.
	//copyStruct,GetRuleContext and ToStringTree  from embedded

	//<if(dispatchMethods)>
	//<dispatchMethods; separator="\n\n">
	//<endif>

	//<if(extensionMembers)>
	//<extensionMembers; separator="\n\n">
	//<endif>
}

type SoruBodyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySoruBodyContext() *SoruBodyContext {
	var p = new(SoruBodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AdlPRULE_soruBody
	return p
}

func (*SoruBodyContext) IsSoruBodyContext() {}

func NewSoruBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SoruBodyContext {
	var p = new(SoruBodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AdlPRULE_soruBody

	return p
}

func (s *SoruBodyContext) GetParser() antlr.Parser { return s.parser }

//StructDecl tokenDecls

//StructDecl tokenTypeDecls

//StructDecl tokenListDecls

//StructDecl ruleContextDecls

//StructDecl ruleContextListDecls

//StructDecl attributeDecls

// Getters

//provideCopyFrom
func (s *SoruBodyContext) CopyFrom(ctx *SoruBodyContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *SoruBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SoruBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

//dispatchMethods
//extensionMembers

//Begin AltLabelStructDecl

type FieldStatementContext struct {
	*SoruBodyContext
	//TokenDecl
	a antlr.Token
	//TokenDecl
	b antlr.Token
}

func NewFieldStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FieldStatementContext {
	var p = new(FieldStatementContext)

	p.SoruBodyContext = NewEmptySoruBodyContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SoruBodyContext))

	return p
}

type IFieldStatementContext interface {
	//Current rule
	ISoruBodyContext
	// start internal
	//Gets for raw elements
	//  ruleGetterDecl
	TypeExpr() ITypeExprContext
	JsonValue() IJsonValueContext
	//  ruleListGetterDecl
	AllAnnon() []IAnnonContext
	//  ruleListIndexedGetterDecl
	Annon(i int) IAnnonContext

	//  tokenGetterDecl
	SEMI() antlr.TerminalNode
	EQ() antlr.TerminalNode
	//  tokenListGetterDecl
	AllID() []antlr.TerminalNode
	//  tokenListIndexedGetterDecl
	ID(i int) antlr.TerminalNode
	// end internal
	//
	//Gets for labeled elements
	//  tokenDecls
	GetA() antlr.Token
	GetB() antlr.Token

	//  tokenTypeDecls
	//  tokenListDecls
	//  ruleContextDecls
	//  ruleContextListDecls
	//  attributeDecls

	// TODO dispatchMethods (needed?)
}

func (*FieldStatementContext) IsFieldStatementContext() {}

//AltLabelStructDecl tokenDecls
func (s *FieldStatementContext) GetA() antlr.Token  { return s.a }
func (s *FieldStatementContext) SetA(v antlr.Token) { s.a = v }

func (s *FieldStatementContext) GetB() antlr.Token  { return s.b }
func (s *FieldStatementContext) SetB(v antlr.Token) { s.b = v }

//AltLabelStructDecl tokenTypeDecls

//AltLabelStructDecl tokenListDecls

//AltLabelStructDecl ruleContextDecls

//AltLabelStructDecl ruleContextListDecls

//AltLabelStructDecl attributeDecls

//getRuleContext
func (s *FieldStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

//getters
func (s *FieldStatementContext) SEMI() antlr.TerminalNode {
	return s.GetToken(AdlPSEMI, 0)
}

func (s *FieldStatementContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(AdlPID)
}

func (s *FieldStatementContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(AdlPID, i)
}

func (s *FieldStatementContext) AllAnnon() []IAnnonContext {
	//ContextRuleListGetterDecl
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*AnnonContext)(nil)).Elem())
	var tst = make([]IAnnonContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAnnonContext)
		}
	}

	return tst
}

func (s *FieldStatementContext) Annon(i int) IAnnonContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*AnnonContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAnnonContext)
}

func (s *FieldStatementContext) TypeExpr() ITypeExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*TypeExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeExprContext)
}

func (s *FieldStatementContext) EQ() antlr.TerminalNode {
	return s.GetToken(AdlPEQ, 0)
}

func (s *FieldStatementContext) JsonValue() IJsonValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*JsonValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IJsonValueContext)
}

func (s *FieldStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FieldStatementEntryListener); ok {
		listenerT.EnterFieldStatement(s)
	}
}

func (s *FieldStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FieldStatementExitListener); ok {
		listenerT.ExitFieldStatement(s)
	}
}

func (s *FieldStatementContext) VisitFunc(hdls antlr.ParserTreeVisitorHandlers, args ...interface{}) (result interface{}) {
	h := hdls.(*AdlPHandlers)
	if h.EnterEveryRule != nil {
		h.EnterEveryRule(s)
	}
	if h.FieldStatement != nil {
		h.FieldStatement(s, h, args...)
	} else {
		s.VisitChildrenFunc(h, args...)
	}
	if h.ExitEveryRule != nil {
		h.ExitEveryRule(s)
	}
	return
}

func (s *FieldStatementContext) Visit(delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	switch t := delegate.(type) {
	case FieldStatementContextVisitor:
		return t.VisitFieldStatement(s, delegate, args...)
	default:
		return delegate.VisitChildren(s, delegate, args...)
	}
}

//END AltLabelStructDecl

func (p *AdlP) SoruBody() (localctx ISoruBodyContext) {
	localctx = NewSoruBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, AdlPRULE_soruBody)
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

	localctx = NewFieldStatementContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(218)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == AdlPAT || _la == AdlPLINE_DOC {
		{
			p.SetState(215)
			p.Annon()
		}

		p.SetState(220)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(221)
		var _m = p.Match(AdlPID)
		localctx.(*FieldStatementContext).a = _m

	}
	p.SetState(223)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == AdlPLCHEVR {
		{
			p.SetState(222)
			p.TypeExpr()
		}

	}
	{
		p.SetState(225)
		var _m = p.Match(AdlPID)
		localctx.(*FieldStatementContext).b = _m

	}
	p.SetState(228)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == AdlPEQ {
		{
			p.SetState(226)
			p.Match(AdlPEQ)
		}
		{
			p.SetState(227)
			p.JsonValue()
		}

	}
	{
		p.SetState(230)
		p.Match(AdlPSEMI)
	}

	return localctx
}

type IJsonValueContext interface {
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

	// IsJsonValueContext differentiates from other interfaces.
	//copyStruct,GetRuleContext and ToStringTree  from embedded

	//<if(dispatchMethods)>
	//<dispatchMethods; separator="\n\n">
	//<endif>

	//<if(extensionMembers)>
	//<extensionMembers; separator="\n\n">
	//<endif>
}

type JsonValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJsonValueContext() *JsonValueContext {
	var p = new(JsonValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AdlPRULE_jsonValue
	return p
}

func (*JsonValueContext) IsJsonValueContext() {}

func NewJsonValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *JsonValueContext {
	var p = new(JsonValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AdlPRULE_jsonValue

	return p
}

func (s *JsonValueContext) GetParser() antlr.Parser { return s.parser }

//StructDecl tokenDecls

//StructDecl tokenTypeDecls

//StructDecl tokenListDecls

//StructDecl ruleContextDecls

//StructDecl ruleContextListDecls

//StructDecl attributeDecls

// Getters

//provideCopyFrom
func (s *JsonValueContext) CopyFrom(ctx *JsonValueContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *JsonValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JsonValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

//dispatchMethods
//extensionMembers

//Begin AltLabelStructDecl

type TrueFalseNullContext struct {
	*JsonValueContext
	//TokenDecl
	kw antlr.Token
}

func NewTrueFalseNullContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TrueFalseNullContext {
	var p = new(TrueFalseNullContext)

	p.JsonValueContext = NewEmptyJsonValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*JsonValueContext))

	return p
}

type ITrueFalseNullContext interface {
	//Current rule
	IJsonValueContext
	// start internal
	//Gets for raw elements
	//  ruleGetterDecl
	//  ruleListGetterDecl
	//  ruleListIndexedGetterDecl

	//  tokenGetterDecl
	ID() antlr.TerminalNode
	//  tokenListGetterDecl
	//  tokenListIndexedGetterDecl
	// end internal
	//
	//Gets for labeled elements
	//  tokenDecls
	GetKw() antlr.Token

	//  tokenTypeDecls
	//  tokenListDecls
	//  ruleContextDecls
	//  ruleContextListDecls
	//  attributeDecls

	// TODO dispatchMethods (needed?)
}

func (*TrueFalseNullContext) IsTrueFalseNullContext() {}

//AltLabelStructDecl tokenDecls
func (s *TrueFalseNullContext) GetKw() antlr.Token  { return s.kw }
func (s *TrueFalseNullContext) SetKw(v antlr.Token) { s.kw = v }

//AltLabelStructDecl tokenTypeDecls

//AltLabelStructDecl tokenListDecls

//AltLabelStructDecl ruleContextDecls

//AltLabelStructDecl ruleContextListDecls

//AltLabelStructDecl attributeDecls

//getRuleContext
func (s *TrueFalseNullContext) GetRuleContext() antlr.RuleContext {
	return s
}

//getters
func (s *TrueFalseNullContext) ID() antlr.TerminalNode {
	return s.GetToken(AdlPID, 0)
}

func (s *TrueFalseNullContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TrueFalseNullEntryListener); ok {
		listenerT.EnterTrueFalseNull(s)
	}
}

func (s *TrueFalseNullContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TrueFalseNullExitListener); ok {
		listenerT.ExitTrueFalseNull(s)
	}
}

func (s *TrueFalseNullContext) VisitFunc(hdls antlr.ParserTreeVisitorHandlers, args ...interface{}) (result interface{}) {
	h := hdls.(*AdlPHandlers)
	if h.EnterEveryRule != nil {
		h.EnterEveryRule(s)
	}
	if h.TrueFalseNull != nil {
		h.TrueFalseNull(s, h, args...)
	} else {
		s.VisitChildrenFunc(h, args...)
	}
	if h.ExitEveryRule != nil {
		h.ExitEveryRule(s)
	}
	return
}

func (s *TrueFalseNullContext) Visit(delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	switch t := delegate.(type) {
	case TrueFalseNullContextVisitor:
		return t.VisitTrueFalseNull(s, delegate, args...)
	default:
		return delegate.VisitChildren(s, delegate, args...)
	}
}

//END AltLabelStructDecl

//Begin AltLabelStructDecl

type ObjStatementContext struct {
	*JsonValueContext
	//TokenDecl
	_STR antlr.Token
	//TokenListDecl
	k []antlr.Token
	//RuleContextDecl
	_jsonValue IJsonValueContext
	//RuleContextListDecl
	v []IJsonValueContext
}

func NewObjStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ObjStatementContext {
	var p = new(ObjStatementContext)

	p.JsonValueContext = NewEmptyJsonValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*JsonValueContext))

	return p
}

type IObjStatementContext interface {
	//Current rule
	IJsonValueContext
	// start internal
	//Gets for raw elements
	//  ruleGetterDecl
	//  ruleListGetterDecl
	AllJsonValue() []IJsonValueContext
	//  ruleListIndexedGetterDecl
	JsonValue(i int) IJsonValueContext

	//  tokenGetterDecl
	LCUR() antlr.TerminalNode
	RCUR() antlr.TerminalNode
	//  tokenListGetterDecl
	AllCOLON() []antlr.TerminalNode
	AllSTR() []antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	//  tokenListIndexedGetterDecl
	COLON(i int) antlr.TerminalNode
	STR(i int) antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode
	// end internal
	//
	//Gets for labeled elements
	//  tokenDecls

	//  tokenTypeDecls
	//  tokenListDecls
	GetK() []antlr.Token
	//  ruleContextDecls

	//  ruleContextListDecls
	GetV() []IJsonValueContext

	//  attributeDecls

	// TODO dispatchMethods (needed?)
}

func (*ObjStatementContext) IsObjStatementContext() {}

//AltLabelStructDecl tokenDecls
func (s *ObjStatementContext) Get_STR() antlr.Token  { return s._STR }
func (s *ObjStatementContext) Set_STR(v antlr.Token) { s._STR = v }

//AltLabelStructDecl tokenTypeDecls

//AltLabelStructDecl tokenListDecls
func (s *ObjStatementContext) GetK() []antlr.Token  { return s.k }
func (s *ObjStatementContext) SetK(v []antlr.Token) { s.k = v }

//AltLabelStructDecl ruleContextDecls
func (s *ObjStatementContext) Get_jsonValue() IJsonValueContext  { return s._jsonValue }
func (s *ObjStatementContext) Set_jsonValue(v IJsonValueContext) { s._jsonValue = v }

//AltLabelStructDecl ruleContextListDecls
func (s *ObjStatementContext) GetV() []IJsonValueContext  { return s.v }
func (s *ObjStatementContext) SetV(v []IJsonValueContext) { s.v = v }

//AltLabelStructDecl attributeDecls

//getRuleContext
func (s *ObjStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

//getters
func (s *ObjStatementContext) LCUR() antlr.TerminalNode {
	return s.GetToken(AdlPLCUR, 0)
}

func (s *ObjStatementContext) RCUR() antlr.TerminalNode {
	return s.GetToken(AdlPRCUR, 0)
}

func (s *ObjStatementContext) AllCOLON() []antlr.TerminalNode {
	return s.GetTokens(AdlPCOLON)
}

func (s *ObjStatementContext) COLON(i int) antlr.TerminalNode {
	return s.GetToken(AdlPCOLON, i)
}

func (s *ObjStatementContext) AllSTR() []antlr.TerminalNode {
	return s.GetTokens(AdlPSTR)
}

func (s *ObjStatementContext) STR(i int) antlr.TerminalNode {
	return s.GetToken(AdlPSTR, i)
}

func (s *ObjStatementContext) AllJsonValue() []IJsonValueContext {
	//ContextRuleListGetterDecl
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*JsonValueContext)(nil)).Elem())
	var tst = make([]IJsonValueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IJsonValueContext)
		}
	}

	return tst
}

func (s *ObjStatementContext) JsonValue(i int) IJsonValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*JsonValueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IJsonValueContext)
}

func (s *ObjStatementContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(AdlPCOMMA)
}

func (s *ObjStatementContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(AdlPCOMMA, i)
}

func (s *ObjStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ObjStatementEntryListener); ok {
		listenerT.EnterObjStatement(s)
	}
}

func (s *ObjStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ObjStatementExitListener); ok {
		listenerT.ExitObjStatement(s)
	}
}

func (s *ObjStatementContext) VisitFunc(hdls antlr.ParserTreeVisitorHandlers, args ...interface{}) (result interface{}) {
	h := hdls.(*AdlPHandlers)
	if h.EnterEveryRule != nil {
		h.EnterEveryRule(s)
	}
	if h.ObjStatement != nil {
		h.ObjStatement(s, h, args...)
	} else {
		s.VisitChildrenFunc(h, args...)
	}
	if h.ExitEveryRule != nil {
		h.ExitEveryRule(s)
	}
	return
}

func (s *ObjStatementContext) Visit(delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	switch t := delegate.(type) {
	case ObjStatementContextVisitor:
		return t.VisitObjStatement(s, delegate, args...)
	default:
		return delegate.VisitChildren(s, delegate, args...)
	}
}

//END AltLabelStructDecl

//Begin AltLabelStructDecl

type FloatStatementContext struct {
	*JsonValueContext
	//TokenDecl
	f antlr.Token
}

func NewFloatStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FloatStatementContext {
	var p = new(FloatStatementContext)

	p.JsonValueContext = NewEmptyJsonValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*JsonValueContext))

	return p
}

type IFloatStatementContext interface {
	//Current rule
	IJsonValueContext
	// start internal
	//Gets for raw elements
	//  ruleGetterDecl
	//  ruleListGetterDecl
	//  ruleListIndexedGetterDecl

	//  tokenGetterDecl
	FLT() antlr.TerminalNode
	//  tokenListGetterDecl
	//  tokenListIndexedGetterDecl
	// end internal
	//
	//Gets for labeled elements
	//  tokenDecls
	GetF() antlr.Token

	//  tokenTypeDecls
	//  tokenListDecls
	//  ruleContextDecls
	//  ruleContextListDecls
	//  attributeDecls

	// TODO dispatchMethods (needed?)
}

func (*FloatStatementContext) IsFloatStatementContext() {}

//AltLabelStructDecl tokenDecls
func (s *FloatStatementContext) GetF() antlr.Token  { return s.f }
func (s *FloatStatementContext) SetF(v antlr.Token) { s.f = v }

//AltLabelStructDecl tokenTypeDecls

//AltLabelStructDecl tokenListDecls

//AltLabelStructDecl ruleContextDecls

//AltLabelStructDecl ruleContextListDecls

//AltLabelStructDecl attributeDecls

//getRuleContext
func (s *FloatStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

//getters
func (s *FloatStatementContext) FLT() antlr.TerminalNode {
	return s.GetToken(AdlPFLT, 0)
}

func (s *FloatStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FloatStatementEntryListener); ok {
		listenerT.EnterFloatStatement(s)
	}
}

func (s *FloatStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FloatStatementExitListener); ok {
		listenerT.ExitFloatStatement(s)
	}
}

func (s *FloatStatementContext) VisitFunc(hdls antlr.ParserTreeVisitorHandlers, args ...interface{}) (result interface{}) {
	h := hdls.(*AdlPHandlers)
	if h.EnterEveryRule != nil {
		h.EnterEveryRule(s)
	}
	if h.FloatStatement != nil {
		h.FloatStatement(s, h, args...)
	} else {
		s.VisitChildrenFunc(h, args...)
	}
	if h.ExitEveryRule != nil {
		h.ExitEveryRule(s)
	}
	return
}

func (s *FloatStatementContext) Visit(delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	switch t := delegate.(type) {
	case FloatStatementContextVisitor:
		return t.VisitFloatStatement(s, delegate, args...)
	default:
		return delegate.VisitChildren(s, delegate, args...)
	}
}

//END AltLabelStructDecl

//Begin AltLabelStructDecl

type ArrayStatementContext struct {
	*JsonValueContext
	//RuleContextDecl
	_jsonValue IJsonValueContext
	//RuleContextListDecl
	jv []IJsonValueContext
}

func NewArrayStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArrayStatementContext {
	var p = new(ArrayStatementContext)

	p.JsonValueContext = NewEmptyJsonValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*JsonValueContext))

	return p
}

type IArrayStatementContext interface {
	//Current rule
	IJsonValueContext
	// start internal
	//Gets for raw elements
	//  ruleGetterDecl
	//  ruleListGetterDecl
	AllJsonValue() []IJsonValueContext
	//  ruleListIndexedGetterDecl
	JsonValue(i int) IJsonValueContext

	//  tokenGetterDecl
	LSQ() antlr.TerminalNode
	RSQ() antlr.TerminalNode
	//  tokenListGetterDecl
	AllCOMMA() []antlr.TerminalNode
	//  tokenListIndexedGetterDecl
	COMMA(i int) antlr.TerminalNode
	// end internal
	//
	//Gets for labeled elements
	//  tokenDecls
	//  tokenTypeDecls
	//  tokenListDecls
	//  ruleContextDecls

	//  ruleContextListDecls
	GetJv() []IJsonValueContext

	//  attributeDecls

	// TODO dispatchMethods (needed?)
}

func (*ArrayStatementContext) IsArrayStatementContext() {}

//AltLabelStructDecl tokenDecls

//AltLabelStructDecl tokenTypeDecls

//AltLabelStructDecl tokenListDecls

//AltLabelStructDecl ruleContextDecls
func (s *ArrayStatementContext) Get_jsonValue() IJsonValueContext  { return s._jsonValue }
func (s *ArrayStatementContext) Set_jsonValue(v IJsonValueContext) { s._jsonValue = v }

//AltLabelStructDecl ruleContextListDecls
func (s *ArrayStatementContext) GetJv() []IJsonValueContext  { return s.jv }
func (s *ArrayStatementContext) SetJv(v []IJsonValueContext) { s.jv = v }

//AltLabelStructDecl attributeDecls

//getRuleContext
func (s *ArrayStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

//getters
func (s *ArrayStatementContext) LSQ() antlr.TerminalNode {
	return s.GetToken(AdlPLSQ, 0)
}

func (s *ArrayStatementContext) RSQ() antlr.TerminalNode {
	return s.GetToken(AdlPRSQ, 0)
}

func (s *ArrayStatementContext) AllJsonValue() []IJsonValueContext {
	//ContextRuleListGetterDecl
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*JsonValueContext)(nil)).Elem())
	var tst = make([]IJsonValueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IJsonValueContext)
		}
	}

	return tst
}

func (s *ArrayStatementContext) JsonValue(i int) IJsonValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*JsonValueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IJsonValueContext)
}

func (s *ArrayStatementContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(AdlPCOMMA)
}

func (s *ArrayStatementContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(AdlPCOMMA, i)
}

func (s *ArrayStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArrayStatementEntryListener); ok {
		listenerT.EnterArrayStatement(s)
	}
}

func (s *ArrayStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArrayStatementExitListener); ok {
		listenerT.ExitArrayStatement(s)
	}
}

func (s *ArrayStatementContext) VisitFunc(hdls antlr.ParserTreeVisitorHandlers, args ...interface{}) (result interface{}) {
	h := hdls.(*AdlPHandlers)
	if h.EnterEveryRule != nil {
		h.EnterEveryRule(s)
	}
	if h.ArrayStatement != nil {
		h.ArrayStatement(s, h, args...)
	} else {
		s.VisitChildrenFunc(h, args...)
	}
	if h.ExitEveryRule != nil {
		h.ExitEveryRule(s)
	}
	return
}

func (s *ArrayStatementContext) Visit(delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	switch t := delegate.(type) {
	case ArrayStatementContextVisitor:
		return t.VisitArrayStatement(s, delegate, args...)
	default:
		return delegate.VisitChildren(s, delegate, args...)
	}
}

//END AltLabelStructDecl

//Begin AltLabelStructDecl

type NumberStatementContext struct {
	*JsonValueContext
	//TokenDecl
	n antlr.Token
}

func NewNumberStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NumberStatementContext {
	var p = new(NumberStatementContext)

	p.JsonValueContext = NewEmptyJsonValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*JsonValueContext))

	return p
}

type INumberStatementContext interface {
	//Current rule
	IJsonValueContext
	// start internal
	//Gets for raw elements
	//  ruleGetterDecl
	//  ruleListGetterDecl
	//  ruleListIndexedGetterDecl

	//  tokenGetterDecl
	INT() antlr.TerminalNode
	//  tokenListGetterDecl
	//  tokenListIndexedGetterDecl
	// end internal
	//
	//Gets for labeled elements
	//  tokenDecls
	GetN() antlr.Token

	//  tokenTypeDecls
	//  tokenListDecls
	//  ruleContextDecls
	//  ruleContextListDecls
	//  attributeDecls

	// TODO dispatchMethods (needed?)
}

func (*NumberStatementContext) IsNumberStatementContext() {}

//AltLabelStructDecl tokenDecls
func (s *NumberStatementContext) GetN() antlr.Token  { return s.n }
func (s *NumberStatementContext) SetN(v antlr.Token) { s.n = v }

//AltLabelStructDecl tokenTypeDecls

//AltLabelStructDecl tokenListDecls

//AltLabelStructDecl ruleContextDecls

//AltLabelStructDecl ruleContextListDecls

//AltLabelStructDecl attributeDecls

//getRuleContext
func (s *NumberStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

//getters
func (s *NumberStatementContext) INT() antlr.TerminalNode {
	return s.GetToken(AdlPINT, 0)
}

func (s *NumberStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumberStatementEntryListener); ok {
		listenerT.EnterNumberStatement(s)
	}
}

func (s *NumberStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NumberStatementExitListener); ok {
		listenerT.ExitNumberStatement(s)
	}
}

func (s *NumberStatementContext) VisitFunc(hdls antlr.ParserTreeVisitorHandlers, args ...interface{}) (result interface{}) {
	h := hdls.(*AdlPHandlers)
	if h.EnterEveryRule != nil {
		h.EnterEveryRule(s)
	}
	if h.NumberStatement != nil {
		h.NumberStatement(s, h, args...)
	} else {
		s.VisitChildrenFunc(h, args...)
	}
	if h.ExitEveryRule != nil {
		h.ExitEveryRule(s)
	}
	return
}

func (s *NumberStatementContext) Visit(delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	switch t := delegate.(type) {
	case NumberStatementContextVisitor:
		return t.VisitNumberStatement(s, delegate, args...)
	default:
		return delegate.VisitChildren(s, delegate, args...)
	}
}

//END AltLabelStructDecl

//Begin AltLabelStructDecl

type StringStatementContext struct {
	*JsonValueContext
	//TokenDecl
	s antlr.Token
}

func NewStringStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringStatementContext {
	var p = new(StringStatementContext)

	p.JsonValueContext = NewEmptyJsonValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*JsonValueContext))

	return p
}

type IStringStatementContext interface {
	//Current rule
	IJsonValueContext
	// start internal
	//Gets for raw elements
	//  ruleGetterDecl
	//  ruleListGetterDecl
	//  ruleListIndexedGetterDecl

	//  tokenGetterDecl
	STR() antlr.TerminalNode
	//  tokenListGetterDecl
	//  tokenListIndexedGetterDecl
	// end internal
	//
	//Gets for labeled elements
	//  tokenDecls
	GetS() antlr.Token

	//  tokenTypeDecls
	//  tokenListDecls
	//  ruleContextDecls
	//  ruleContextListDecls
	//  attributeDecls

	// TODO dispatchMethods (needed?)
}

func (*StringStatementContext) IsStringStatementContext() {}

//AltLabelStructDecl tokenDecls
func (s *StringStatementContext) GetS() antlr.Token  { return s.s }
func (s *StringStatementContext) SetS(v antlr.Token) { s.s = v }

//AltLabelStructDecl tokenTypeDecls

//AltLabelStructDecl tokenListDecls

//AltLabelStructDecl ruleContextDecls

//AltLabelStructDecl ruleContextListDecls

//AltLabelStructDecl attributeDecls

//getRuleContext
func (s *StringStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

//getters
func (s *StringStatementContext) STR() antlr.TerminalNode {
	return s.GetToken(AdlPSTR, 0)
}

func (s *StringStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StringStatementEntryListener); ok {
		listenerT.EnterStringStatement(s)
	}
}

func (s *StringStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(StringStatementExitListener); ok {
		listenerT.ExitStringStatement(s)
	}
}

func (s *StringStatementContext) VisitFunc(hdls antlr.ParserTreeVisitorHandlers, args ...interface{}) (result interface{}) {
	h := hdls.(*AdlPHandlers)
	if h.EnterEveryRule != nil {
		h.EnterEveryRule(s)
	}
	if h.StringStatement != nil {
		h.StringStatement(s, h, args...)
	} else {
		s.VisitChildrenFunc(h, args...)
	}
	if h.ExitEveryRule != nil {
		h.ExitEveryRule(s)
	}
	return
}

func (s *StringStatementContext) Visit(delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	switch t := delegate.(type) {
	case StringStatementContextVisitor:
		return t.VisitStringStatement(s, delegate, args...)
	default:
		return delegate.VisitChildren(s, delegate, args...)
	}
}

//END AltLabelStructDecl

func (p *AdlP) JsonValue() (localctx IJsonValueContext) {
	localctx = NewJsonValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, AdlPRULE_jsonValue)
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

	p.SetState(264)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AdlPSTR:
		localctx = NewStringStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(232)
			var _m = p.Match(AdlPSTR)
			localctx.(*StringStatementContext).s = _m

		}

	case AdlPID:
		localctx = NewTrueFalseNullContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(233)
			var _m = p.Match(AdlPID)
			localctx.(*TrueFalseNullContext).kw = _m

		}

	case AdlPINT:
		localctx = NewNumberStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(234)
			var _m = p.Match(AdlPINT)
			localctx.(*NumberStatementContext).n = _m

		}

	case AdlPFLT:
		localctx = NewFloatStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(235)
			var _m = p.Match(AdlPFLT)
			localctx.(*FloatStatementContext).f = _m

		}

	case AdlPLSQ:
		localctx = NewArrayStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(236)
			p.Match(AdlPLSQ)
		}
		p.SetState(245)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AdlPLCUR)|(1<<AdlPLSQ)|(1<<AdlPSTR)|(1<<AdlPID)|(1<<AdlPINT)|(1<<AdlPFLT))) != 0 {
			{
				p.SetState(237)

				var _x = p.JsonValue()

				localctx.(*ArrayStatementContext)._jsonValue = _x

			}
			localctx.(*ArrayStatementContext).jv = append(localctx.(*ArrayStatementContext).jv, localctx.(*ArrayStatementContext)._jsonValue)
			p.SetState(242)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == AdlPCOMMA {
				{
					p.SetState(238)
					p.Match(AdlPCOMMA)
				}
				{
					p.SetState(239)

					var _x = p.JsonValue()

					localctx.(*ArrayStatementContext)._jsonValue = _x

				}
				localctx.(*ArrayStatementContext).jv = append(localctx.(*ArrayStatementContext).jv, localctx.(*ArrayStatementContext)._jsonValue)

				p.SetState(244)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(247)
			p.Match(AdlPRSQ)
		}

	case AdlPLCUR:
		localctx = NewObjStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(248)
			p.Match(AdlPLCUR)
		}
		p.SetState(261)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == AdlPSTR {
			{
				p.SetState(249)
				var _m = p.Match(AdlPSTR)
				localctx.(*ObjStatementContext)._STR = _m

			}
			localctx.(*ObjStatementContext).k = append(localctx.(*ObjStatementContext).k, localctx.(*ObjStatementContext)._STR)
			{
				p.SetState(250)
				p.Match(AdlPCOLON)
			}
			{
				p.SetState(251)

				var _x = p.JsonValue()

				localctx.(*ObjStatementContext)._jsonValue = _x

			}
			localctx.(*ObjStatementContext).v = append(localctx.(*ObjStatementContext).v, localctx.(*ObjStatementContext)._jsonValue)
			p.SetState(258)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == AdlPCOMMA {
				{
					p.SetState(252)
					p.Match(AdlPCOMMA)
				}
				{
					p.SetState(253)
					var _m = p.Match(AdlPSTR)
					localctx.(*ObjStatementContext)._STR = _m

				}
				localctx.(*ObjStatementContext).k = append(localctx.(*ObjStatementContext).k, localctx.(*ObjStatementContext)._STR)
				{
					p.SetState(254)
					p.Match(AdlPCOLON)
				}
				{
					p.SetState(255)

					var _x = p.JsonValue()

					localctx.(*ObjStatementContext)._jsonValue = _x

				}
				localctx.(*ObjStatementContext).v = append(localctx.(*ObjStatementContext).v, localctx.(*ObjStatementContext)._jsonValue)

				p.SetState(260)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(263)
			p.Match(AdlPRCUR)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}
