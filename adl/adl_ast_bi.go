package adl

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/golang/glog"
	antlr "github.com/wxio/goantlr"
	parser "github.com/wxio/tron-go/internal/adllp"
	"github.com/wxio/tron-go/internal/ctree"
)

func BuildAdlAST(str string) (ctree.Tree, *antlr.BaseLexer, antlr.TokenStream, errs) {
	el := &lexErr{}
	tbl := &ADLBuildListener{
		// debug: true,
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
	tbl.errs.lexErr = el.err
	tbl.errs.lexWarning = el.warning
	if el.err != nil || len(tbl.errs.parserErr) != 0 {
		return nil, nil, nil, tbl.errs
	}
	// fmt.Printf("--------%v %+v\n", el.err, tbl.errToks)
	antlr.ParseTreeWalkerDefault.Walk(tbl, ctx)
	return tbl.bldr.Build(), lexer.BaseLexer, stream, tbl.errs
}

type ADLBuildListener struct {
	*antlr.BaseParseTreeVisitor
	bldr  ctree.WalkableBuilder
	adl   *ADL
	errs  errs
	debug bool
}

type errs struct {
	lexErr     []interface{}
	lexWarning []interface{}
	parserErr  []interface{}
	errToks    []interface{}
	warnings   []interface{}
}

func (er errs) Error() error {
	if len(er.lexErr) == 0 && len(er.parserErr) == 0 && len(er.errToks) == 0 {
		return nil
	}
	return fmt.Errorf("lex:%v parser:%v syntax:%v", er.lexErr, er.parserErr, er.errToks)
}

// EnterEveryRule is called when any rule is entered.
func (v *ADLBuildListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	switch ctx := ctx.(type) {
	case *parser.AdlContext:
		v.bldr = ctree.NewBuild("ADL", ctx.GetStart(), parser.AdlPADL, nil)
		v.adl = &ADL{}
	case *parser.ModuleStatementContext:
		if ctx.GetKw().GetText() != "module" {
			et := Error{Expected: []string{"module"}, Received: ctx.GetKw().GetText()}
			v.bldr.AddNode(ctx.GetKw(), parser.AdlPERROR, et)
			v.errs.errToks = append(v.errs.errToks, et)
		} else {
			v.bldr.AddNode(ctx.GetKw(), parser.AdlPModule, Module{
				Name: strings.Join(tokens2strings(ctx.GetName()), ".")})
		}
		v.bldr.Down()
	case *parser.ImportModuleNameContext:
		if ctx.GetKw().GetText() != "import" {
			et := Error{Expected: []string{"import"}, Received: ctx.GetKw().GetText()}
			v.bldr.AddNode(ctx.GetKw(), parser.AdlPERROR, et)
			v.errs.errToks = append(v.errs.errToks, et)
		} else {
			path := strings.Join(tokens2strings(ctx.GetA()), ".")
			v.bldr.AddNode(ctx.GetKw(), parser.AdlPImportModule, Import{ModuleName: &path})
		}
	case *parser.ImportScopedNameContext:
		if ctx.GetKw().GetText() != "import" {
			et := Error{Expected: []string{"import"}, Received: ctx.GetKw().GetText()}
			v.bldr.AddNode(ctx.GetKw(), parser.AdlPERROR, et)
			v.errs.errToks = append(v.errs.errToks, et)
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
			et := Error{Expected: []string{"struct", "union"}, Received: ctx.GetKw().GetText()}
			v.bldr.AddNode(ctx.GetKw(), parser.AdlPERROR, et)
			v.errs.errToks = append(v.errs.errToks, et)
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
			et := Error{Expected: []string{"type", "newtype"}, Received: ctx.GetKw().GetText()}
			v.bldr.AddNode(ctx.GetKw(), parser.AdlPERROR, et)
			v.errs.errToks = append(v.errs.errToks, et)
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
			et := Error{Expected: []string{"annotation"}, Received: ctx.GetKw().GetText()}
			v.bldr.AddNode(ctx.GetStart(), parser.AdlPERROR, et)
			v.errs.errToks = append(v.errs.errToks, et)
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
			et := Error{Expected: []string{"annotation"}, Received: ctx.GetKw().GetText()}
			v.bldr.AddNode(ctx.GetStart(), parser.AdlPERROR, et)
			v.errs.errToks = append(v.errs.errToks, et)
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
			et := Error{Expected: []string{"annotation"}, Received: ctx.GetKw().GetText()}
			v.bldr.AddNode(ctx.GetStart(), parser.AdlPERROR, et)
			v.errs.errToks = append(v.errs.errToks, et)
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
			et := Error{Expected: []string{"true", "false", "null"}, Received: ctx.GetKw().GetText()}
			v.bldr.AddNode(ctx.GetStart(), parser.AdlPERROR, et)
			v.errs.errToks = append(v.errs.errToks, et)
		}
	case *parser.NumberStatementContext:
		i, err := strconv.ParseInt(ctx.GetN().GetText(), 10, 64)
		if err == nil {
			v.bldr.AddNode(ctx.GetStart(), parser.AdlPJsonInt, i)
		} else {
			et := Error{Expected: []string{"<number>"}, Received: ctx.GetN().GetText()}
			v.bldr.AddNode(ctx.GetStart(), parser.AdlPERROR, et)
			v.errs.errToks = append(v.errs.errToks, et)

		}
	case *parser.FloatStatementContext:
		i, err := strconv.ParseFloat(ctx.GetF().GetText(), 64)
		if err == nil {
			v.bldr.AddNode(ctx.GetStart(), parser.AdlPJsonFloat, i)
		} else {
			et := Error{Expected: []string{"<float>"}, Received: ctx.GetF().GetText()}
			v.bldr.AddNode(ctx.GetStart(), parser.AdlPERROR, et)
			v.errs.errToks = append(v.errs.errToks, et)

		}
	case *parser.ArrayStatementContext:
		v.bldr.AddNode(ctx.GetStart(), parser.AdlPJsonArray, JsonArray{}).Down()
	case *parser.ObjStatementContext:
		v.bldr.AddNode(ctx.GetStart(), parser.AdlPJsonObj, JsonObj{}).Down()
		// rules name occur on errors
	case *parser.Top_level_statementContext:
		et := Error{Expected: []string{"@|struct|union|annotation"}, Received: ctx.GetStart().GetText()}
		v.bldr.AddNode(ctx.GetStart(), parser.AdlPERROR, et)
		v.errs.errToks = append(v.errs.errToks, et)

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
	if strings.HasPrefix(msg, "reportAttemptingFullContext") { // TODO remove NewDiagnosticErrorListener and move warning to ReportAmbiguity etc. when getDecisionDescription is make public
		v.errs.warnings = append(v.errs.warnings, fmt.Sprintf("At %d:%d <%s>", line, column, msg))
		return
	}
	t, ok := offendingSymbol.(antlr.Token)
	if !ok && e != nil {
		t = e.GetOffendingToken()
	}
	v.errs.parserErr = append(v.errs.parserErr, t)
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
