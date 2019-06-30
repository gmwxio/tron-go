package adl

import (
	"fmt"
	"strconv"

	"strings"

	"github.com/golang/glog"
	antlr "github.com/wxio/goantlr"
	parser "github.com/wxio/tron-go/adllp"
	"github.com/wxio/tron-go/internal/ctree"
)

func BuildAdlAST(str string) (ctree.Tree, *antlr.BaseLexer, antlr.TokenStream, error) {
	el := &antlrEL{}
	tbl := &ADLBuildListener{
		debug: true,
	}
	is := antlr.NewInputStream(str)
	lexer := parser.NewAdlL(is)
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(el)

	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewAdlP(stream)
	p.RemoveErrorListeners()
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.AddErrorListener(tbl)
	p.BuildParseTrees = true
	ctx := p.Adl()
	if el.err != nil {
		return nil, nil, nil, el.err
	}
	fmt.Printf("--------%v %v\n", el.err, tbl.errToks)
	antlr.ParseTreeWalkerDefault.Walk(tbl, ctx)
	// if tbl.err != "" {
	// 	return nil, lexer.BaseLexer, stream, fmt.Errorf("ERROR:%v", tbl.err)
	// }
	// return nil, nil, nil, nil
	var errToks error
	if len(tbl.errToks) != 0 {
		errToks = fmt.Errorf("%v", tbl.errToks)
	}
	return tbl.bldr.Build(), lexer.BaseLexer, stream, errToks
}

type ADLBuildListener struct {
	*antlr.BaseParseTreeVisitor
	bldr    ctree.WalkableBuilder
	adl     *ADL
	errToks []interface{}
	debug   bool
}

// EnterEveryRule is called when any rule is entered.
func (v *ADLBuildListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	switch ctx := ctx.(type) {
	case *parser.AdlContext:
		v.bldr = ctree.NewBuild("ADL", ctx.GetStart(), parser.AdlPADL, nil)
		v.adl = &ADL{}
	case *parser.ModuleStatementContext:
		if ctx.GetKw().GetText() != "module" {
			v.bldr.AddNode(
				ctx.GetKw(), parser.AdlPERROR, Error{
					Expected: []string{"module"}, Received: ctx.GetKw().GetText()},
			)
		} else {
			v.bldr.AddNode(
				ctx.GetKw(), parser.AdlPModule, Module{
					Name: strings.Join(tokens2strings(ctx.GetName()), ".")},
			)
		}
		v.bldr.Down()
	case *parser.ImportModuleNameContext:
		if ctx.GetKw().GetText() != "import" {
			v.bldr.AddNode(ctx.GetKw(), parser.AdlPERROR, Error{
				Expected: []string{"import"}, Received: ctx.GetKw().GetText()})
		} else {
			path := strings.Join(tokens2strings(ctx.GetA()), ".")
			v.bldr.AddNode(ctx.GetKw(), parser.AdlPImportModule, Import{ModuleName: &path})
		}
	case *parser.ImportScopedNameContext:
		if ctx.GetKw().GetText() != "import" {
			v.bldr.AddNode(ctx.GetKw(), parser.AdlPERROR, Error{
				Expected: []string{"import"}, Received: ctx.GetKw().GetText()})
		} else {
			toks := ctx.GetA()
			v.bldr.AddNode(ctx.GetKw(), parser.AdlPImportScopedName, Import{ScopedName: &ScopedName{
				ModuleName: strings.Join(tokens2strings(toks[:len(toks)-1]), "."),
				Name:       toks[len(toks)-1].GetText(),
			}})
		}
	case *parser.LocalAnnoContext:
		an := Annotation{Key: ScopedName{Name: ctx.GetA().GetText()}}
		v.bldr.AddNode(ctx.GetA(), parser.AdlPAnnotationNotScoped, an)
		v.bldr.Down()
	case *parser.DocAnnoContext:
		an := Annotation{Key: ScopedName{ModuleName: "sys.annotations", Name: "Doc"}, Val: ctx.GetA().GetText()}
		v.bldr.AddNode(ctx.GetA(), parser.AdlPAnnotation, an)
	case *parser.StructOrUnionContext:
		switch ctx.GetKw().GetText() {
		case "struct":
			v.bldr.AddNode(ctx.GetKw(), parser.AdlPStruct, Decl{Name: ctx.GetA().GetText(), Type: DeclType{
				Struct: &Name{},
			}})
		case "union":
			v.bldr.AddNode(ctx.GetKw(), parser.AdlPUnion, Decl{Name: ctx.GetA().GetText(), Type: DeclType{
				Union: &Name{},
			}})
		default:
			v.bldr.AddNode(ctx.GetKw(), parser.AdlPERROR, Error{
				Expected: []string{"struct", "union"}, Received: ctx.GetKw().GetText()})
		}
		v.bldr.Down()
	case *parser.TypeOrNewtypeContext:
		switch ctx.GetKw().GetText() {
		case "type":
			v.bldr.AddNode(ctx.GetKw(), parser.AdlPType, Decl{Name: ctx.GetA().GetText(), Type: DeclType{
				Type: &TypeDef{}}})
		case "newtype":
			v.bldr.AddNode(ctx.GetKw(), parser.AdlPNewtype, Decl{Name: ctx.GetA().GetText(), Type: DeclType{
				Newtype: &NewType{}}})
		default:
			v.bldr.AddNode(ctx.GetKw(), parser.AdlPERROR, Error{
				Expected: []string{"type", "newtype"}, Received: ctx.GetKw().GetText()})
		}
		v.bldr.Down()
	case *parser.TypeParameterContext:
		v.bldr.AddNode(ctx.GetStart(), parser.AdlPTypeParam, tokens2strings(ctx.GetTypep()))
	case *parser.TypeExpressionElemContext:
		v.bldr.AddNode(ctx.GetStart(), parser.AdlPTypeExprElem, ctx.GetA().GetText())
		v.bldr.Down()
	case *parser.TypeExprPrimOrParamContext:
		v.bldr.AddNode(ctx.GetStart(), parser.AdlPTypeExprPrimOrParam, ctx.GetB().GetText())
	case *parser.TypeExprTypeExprContext:
		v.bldr.AddNode(ctx.GetStart(), parser.AdlPTypeExprTypeExpr, ctx.GetB().GetText())
		v.bldr.Down()
	case *parser.ModuleAnnotationContext:
		switch ctx.GetKw().GetText() {
		case "annotation":
			v.bldr.AddNode(ctx.GetStart(), parser.AdlPModuleAnno, ctx.GetA().GetText())
		default:
			v.bldr.AddNode(ctx.GetStart(), parser.AdlPERROR, Error{
				Expected: []string{"annotation"}, Received: ctx.GetKw().GetText()})
		}
		v.bldr.Down()
	case *parser.DeclAnnotationContext:
		switch ctx.GetKw().GetText() {
		case "annotation":
			v.bldr.AddNode(ctx.GetStart(), parser.AdlPDeclAnno, []string{
				ctx.GetA().GetText(),
				ctx.GetB().GetText(),
			})
		default:
			v.bldr.AddNode(ctx.GetStart(), parser.AdlPERROR, Error{
				Expected: []string{"annotation"}, Received: ctx.GetKw().GetText()})
		}
		v.bldr.Down()
	case *parser.FieldAnnotationContext:
		switch ctx.GetKw().GetText() {
		case "annotation":
			v.bldr.AddNode(ctx.GetStart(), parser.AdlPFieldAnno, []string{
				ctx.GetA().GetText(),
				ctx.GetB().GetText(),
				ctx.GetC().GetText(),
			})
		default:
			v.bldr.AddNode(ctx.GetStart(), parser.AdlPERROR, Error{
				Expected: []string{"annotation"}, Received: ctx.GetKw().GetText()})
		}
		v.bldr.Down()
	case *parser.FieldStatementContext:
		v.bldr.AddNode(ctx.GetStart(), parser.AdlPField, Field{
			Name: ctx.GetB().GetText(),
		}).Down()
	case *parser.StringStatementContext:
		v.bldr.AddNode(ctx.GetStart(), parser.AdlPJsonStr, ctx.GetS().GetText())
	case *parser.TrueFalseNullContext:
		switch strings.ToLower(ctx.GetKw().GetText()) {
		case "true":
			v.bldr.AddNode(ctx.GetStart(), parser.AdlPJsonBool, true)
		case "false":
			v.bldr.AddNode(ctx.GetStart(), parser.AdlPJsonBool, false)
		case "null":
			v.bldr.AddNode(ctx.GetStart(), parser.AdlPJsonNull, nil)
		default:
			v.bldr.AddNode(ctx.GetStart(), parser.AdlPERROR, Error{
				Expected: []string{"true", "false", "null"}, Received: ctx.GetKw().GetText()})
		}
	case *parser.NumberStatementContext:
		i, err := strconv.ParseInt(ctx.GetN().GetText(), 10, 64)
		if err == nil {
			v.bldr.AddNode(ctx.GetStart(), parser.AdlPJsonInt, i)
		} else {
			v.bldr.AddNode(ctx.GetStart(), parser.AdlPERROR, Error{
				Expected: []string{"<number>"}, Received: ctx.GetN().GetText()})
		}
	case *parser.FloatStatementContext:
		i, err := strconv.ParseFloat(ctx.GetF().GetText(), 64)
		if err == nil {
			v.bldr.AddNode(ctx.GetStart(), parser.AdlPJsonFloat, i)
		} else {
			v.bldr.AddNode(ctx.GetStart(), parser.AdlPERROR, Error{
				Expected: []string{"<float>"}, Received: ctx.GetF().GetText()})
		}
	case *parser.ArrayStatementContext:
		v.bldr.AddNode(ctx.GetStart(), parser.AdlPJsonArray, JsonArray{}).Down()
	case *parser.ObjStatementContext:
		v.bldr.AddNode(ctx.GetStart(), parser.AdlPJsonObj, JsonObj{}).Down()
		// rules name occur on errors
	case *parser.Top_level_statementContext:
		v.bldr.AddNode(ctx.GetStart(), parser.AdlPERROR, Error{
			Expected: []string{"@|struct|union|annotation"}, Received: ctx.GetStart().GetText()})
	// case *parser.ErrorTypeParamContext:
	// 	n := &ErrorNode{
	// 		MyToken:  MyToken{Token: ctx.GetStart(), TType: parser.ADLParserERROR},
	// 		Expected: "Need Type Param looks like a Type Expression. ie type A<B> ok, type A<B<C>> not", Received: ctx.GetStart().GetText()}
	// 	tr.Builder.Add(n)
	default:
		glog.Warningf("Unhandled in >EnterEveryRule case %T:\n", ctx)
	}
	// fmt.Printf("\n%s>>%T ", tr.indent, ctx)
	// tr.indent += "  "
}

// ExitEveryRule
func (v *ADLBuildListener) ExitEveryRule(ctx antlr.ParserRuleContext) {
	switch ctx.(type) {
	case *parser.AdlContext:
	case *parser.ModuleStatementContext:
		v.bldr.Up()
	case *parser.ImportModuleNameContext:
	case *parser.ImportScopedNameContext:
	case *parser.LocalAnnoContext:
		v.bldr.Up()
	case *parser.DocAnnoContext:
	case *parser.StructOrUnionContext:
		v.bldr.Up()
	case *parser.TypeOrNewtypeContext:
		v.bldr.Up()
	case *parser.TypeExpressionElemContext:
		v.bldr.Up()
	case *parser.TypeParameterContext:
	case *parser.TypeExprPrimOrParamContext:
	case *parser.TypeExprTypeExprContext:
		v.bldr.Up()
	case *parser.ModuleAnnotationContext:
		v.bldr.Up()
	case *parser.DeclAnnotationContext:
		v.bldr.Up()
	case *parser.FieldAnnotationContext:
		v.bldr.Up()
	case *parser.FieldStatementContext:
		v.bldr.Up()
	case *parser.StringStatementContext:
	case *parser.TrueFalseNullContext:
	case *parser.NumberStatementContext:
	case *parser.FloatStatementContext:
	case *parser.ArrayStatementContext:
		v.bldr.Up()
	case *parser.ObjStatementContext:
		v.bldr.Up()
	default:
		glog.Warningf("Unhandled in <ExitEveryRule case %T:\n", ctx)
	}
}

func (v *ADLBuildListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	t, ok := offendingSymbol.(antlr.Token)
	if !ok && e != nil {
		t = e.GetOffendingToken()
	}
	v.errToks = append(v.errToks, t)
}

func (tbl *ADLBuildListener) ReportAmbiguity(
	recognizer antlr.Parser,
	dfa *antlr.DFA,
	startIndex, stopIndex int,
	exact bool,
	ambigAlts *antlr.BitSet,
	configs antlr.ATNConfigSet) {
	// panic("ReportAmbiguity")
	if tbl.debug {
		fmt.Printf("ReportAmbiguity rec:%v dfs:%v start:%d stop:%d, exact:%v, ambigAlts:%v config:%v\n", recognizer, dfa, startIndex, stopIndex, exact, ambigAlts, configs)
	}
}

func (tbl *ADLBuildListener) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
	// panic("ReportAttemptingFullContext")
	if tbl.debug {
		fmt.Printf("ReportAttemptingFullContext rec:%v dfs:%v start:%d stop:%d, conflictingAlts:%v config:%v\n", recognizer, dfa, startIndex, stopIndex, conflictingAlts, configs)
	}
}

func (tbl *ADLBuildListener) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex, prediction int, configs antlr.ATNConfigSet) {
	//	panic("ReportContextSensitivity")
}
