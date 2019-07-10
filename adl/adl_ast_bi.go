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

func BuildAdlAST(str string) (ctree.Tree, antlr.Tree, *antlr.BaseLexer, antlr.TokenStream, errs) {
	el := &lexErr{}
	is := antlr.NewInputStream(str)
	lexer := parser.NewAdlL(is)
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(el)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewAdlP(stream)
	p.RemoveErrorListeners()
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	pl := &parseErr{}
	p.AddErrorListener(pl)
	p.BuildParseTrees = true
	ctx := p.Adl()
	errs := errs{}
	errs.LexErr = el.err
	errs.LexWarning = el.warning
	errs.ParseErr = pl.ParseErr
	errs.SyntaxWarning = pl.SyntaxWarning
	if len(el.err) != 0 || len(pl.ParseErr) != 0 {
		return nil, ctx, lexer.BaseLexer, stream, errs
	}
	// fmt.Printf("--------%v %+v\n", el.err, tbl.SyntaxErr)
	tbl := &ADLBuildListener{}
	antlr.ParseTreeWalkerDefault.Walk(tbl, ctx)
	errs.SyntaxErr = tbl.errs.SyntaxErr
	return tbl.bldr.Build(), ctx, lexer.BaseLexer, stream, errs
}

type ADLBuildListener struct {
	*antlr.BaseParseTreeVisitor
	bldr ctree.WalkableBuilder
	adl  *ADL
	errs struct {
		SyntaxErr []DiagMessage
	}
	debug bool
}

type errs struct {
	LexErr        []DiagMessage
	LexWarning    []interface{}
	ParseErr      []DiagMessage
	SyntaxErr     []DiagMessage
	SyntaxWarning []interface{}
}

func (er errs) Error() error {
	if len(er.LexErr) == 0 && len(er.ParseErr) == 0 && len(er.SyntaxErr) == 0 {
		return nil
	}
	return fmt.Errorf("lex:%v parser:%v syntax:%v", er.LexErr, er.ParseErr, er.SyntaxErr)
}

// EnterEveryRule is called when any rule is entered.
func (v *ADLBuildListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	switch ctx := ctx.(type) {
	case *parser.AdlContext:
		v.bldr = ctree.NewBuild("ADL", ctx.GetStart(), ctx.GetStop(), parser.AdlPADL, nil)
		v.adl = &ADL{}
	case *parser.ModuleStatementContext:
		if ctx.GetKw().GetText() != "module" {
			et := Error{Start: ctx.GetKw(), Stop: ctx.SEMI().GetSymbol(), Expected: []string{"module"}, Received: ctx.GetKw().GetText()}
			v.bldr.AddNode(ctx.GetKw(), ctx.GetStop(), parser.AdlPERROR, et)
			v.errs.SyntaxErr = append(v.errs.SyntaxErr, et)
		} else {
			v.bldr.AddNode(ctx.GetKw(), ctx.SEMI().GetSymbol(), parser.AdlPModule, Module{
				Name: strings.Join(tokens2strings(ctx.GetName()), ".")})
		}
		v.bldr.Down()
	case *parser.ImportModuleNameContext:
		if ctx.GetKw().GetText() != "import" {
			et := Error{Start: ctx.GetStart(), Stop: ctx.SEMI().GetSymbol(), Expected: []string{"import"}, Received: ctx.GetKw().GetText()}
			v.bldr.AddNode(ctx.GetKw(), ctx.SEMI().GetSymbol(), parser.AdlPERROR, et)
			v.errs.SyntaxErr = append(v.errs.SyntaxErr, et)
		} else {
			path := strings.Join(tokens2strings(ctx.GetA()), ".")
			v.bldr.AddNode(ctx.GetKw(), ctx.SEMI().GetSymbol(), parser.AdlPImportModule, Import{ModuleName: &path})
		}
	case *parser.ImportScopedNameContext:
		if ctx.GetKw().GetText() != "import" {
			et := Error{Start: ctx.GetStart(), Stop: ctx.SEMI().GetSymbol(), Expected: []string{"import"}, Received: ctx.GetKw().GetText()}
			v.bldr.AddNode(ctx.GetKw(), ctx.SEMI().GetSymbol(), parser.AdlPERROR, et)
			v.errs.SyntaxErr = append(v.errs.SyntaxErr, et)
		} else {
			toks := ctx.GetA()
			v.bldr.AddNode(ctx.GetKw(), ctx.SEMI().GetSymbol(), parser.AdlPImportScopedName, Import{ScopedName: &ScopedName{
				ModuleName: strings.Join(tokens2strings(toks[:len(toks)-1]), "."),
				Name:       toks[len(toks)-1].GetText(),
			}})
		}
	case *parser.LocalAnnoContext:
		an := Annotation{Key: ScopedName{Name: ctx.GetA().GetText()}}
		v.bldr.AddNode(ctx.GetA(), ctx.GetStop(), parser.AdlPAnnotationNotScoped, an)
		v.bldr.Down()
	case *parser.DocAnnoContext:
		an := Annotation{Key: ScopedName{ModuleName: "sys.annotations", Name: "Doc"}, Val: ctx.GetA().GetText()}
		v.bldr.AddNode(ctx.GetA(), ctx.GetStop(), parser.AdlPAnnotation, an)
	case *parser.StructOrUnionContext:
		switch ctx.GetKw().GetText() {
		case "struct":
			v.bldr.AddNode(ctx.GetKw(), ctx.SEMI().GetSymbol(), parser.AdlPStruct, Decl{Name: ctx.GetA().GetText(), Type: DeclType{
				Struct: &Name{},
			}})
		case "union":
			v.bldr.AddNode(ctx.GetKw(), ctx.SEMI().GetSymbol(), parser.AdlPUnion, Decl{Name: ctx.GetA().GetText(), Type: DeclType{
				Union: &Name{},
			}})
		default:
			et := Error{Start: ctx.GetStart(), Stop: ctx.SEMI().GetSymbol(), Expected: []string{"struct", "union"}, Received: ctx.GetKw().GetText()}
			v.bldr.AddNode(ctx.GetKw(), ctx.SEMI().GetSymbol(), parser.AdlPERROR, et)
			v.errs.SyntaxErr = append(v.errs.SyntaxErr, et)
		}
		v.bldr.Down()
	case *parser.TypeOrNewtypeContext:
		switch ctx.GetKw().GetText() {
		case "type":
			v.bldr.AddNode(ctx.GetKw(), ctx.SEMI().GetSymbol(), parser.AdlPType, Decl{Name: ctx.GetA().GetText(), Type: DeclType{
				Type: &TypeDef{}}})
		case "newtype":
			v.bldr.AddNode(ctx.GetKw(), ctx.SEMI().GetSymbol(), parser.AdlPNewtype, Decl{Name: ctx.GetA().GetText(), Type: DeclType{
				Newtype: &NewType{}}})
		default:
			et := Error{Start: ctx.GetStart(), Stop: ctx.SEMI().GetSymbol(), Expected: []string{"type", "newtype"}, Received: ctx.GetKw().GetText()}
			v.bldr.AddNode(ctx.GetKw(), ctx.SEMI().GetSymbol(), parser.AdlPERROR, et)
			v.errs.SyntaxErr = append(v.errs.SyntaxErr, et)
		}
		v.bldr.Down()
	case *parser.TypeParameterContext:
		v.bldr.AddNode(ctx.GetStart(), ctx.GetStop(), parser.AdlPTypeParam, tokens2strings(ctx.GetTypep()))
	case *parser.TypeExprSimpleContext:
		v.bldr.AddNode(ctx.GetStart(), ctx.GetStop(), parser.AdlPTypeExprSimple, ctx.GetB().GetText())
	case *parser.TypeExprGenericContext:
		v.bldr.AddNode(ctx.GetStart(), ctx.GetStop(), parser.AdlPTypeExprGeneric, ctx.GetB().GetText())
		v.bldr.Down()
	case *parser.ModuleAnnotationContext:
		switch ctx.GetKw().GetText() {
		case "annotation":
			v.bldr.AddNode(ctx.GetStart(), ctx.SEMI().GetSymbol(), parser.AdlPModuleAnno, ctx.GetA().GetText())
		default:
			et := Error{Start: ctx.GetStart(), Stop: ctx.SEMI().GetSymbol(), Expected: []string{"annotation"}, Received: ctx.GetKw().GetText()}
			v.bldr.AddNode(ctx.GetStart(), ctx.SEMI().GetSymbol(), parser.AdlPERROR, et)
			v.errs.SyntaxErr = append(v.errs.SyntaxErr, et)
		}
		v.bldr.Down()
	case *parser.DeclAnnotationContext:
		switch ctx.GetKw().GetText() {
		case "annotation":
			v.bldr.AddNode(ctx.GetStart(), ctx.SEMI().GetSymbol(), parser.AdlPDeclAnno, []string{
				ctx.GetA().GetText(),
				ctx.GetB().GetText(),
			})
		default:
			et := Error{Start: ctx.GetStart(), Stop: ctx.SEMI().GetSymbol(), Expected: []string{"annotation"}, Received: ctx.GetKw().GetText()}
			v.bldr.AddNode(ctx.GetStart(), ctx.SEMI().GetSymbol(), parser.AdlPERROR, et)
			v.errs.SyntaxErr = append(v.errs.SyntaxErr, et)
		}
		v.bldr.Down()
	case *parser.FieldAnnotationContext:
		switch ctx.GetKw().GetText() {
		case "annotation":
			v.bldr.AddNode(ctx.GetStart(), ctx.SEMI().GetSymbol(), parser.AdlPFieldAnno, []string{
				ctx.GetA().GetText(),
				ctx.GetB().GetText(),
				ctx.GetC().GetText(),
			})
		default:
			et := Error{Start: ctx.GetStart(), Stop: ctx.SEMI().GetSymbol(), Expected: []string{"annotation"}, Received: ctx.GetKw().GetText()}
			v.bldr.AddNode(ctx.GetStart(), ctx.SEMI().GetSymbol(), parser.AdlPERROR, et)
			v.errs.SyntaxErr = append(v.errs.SyntaxErr, et)
		}
		v.bldr.Down()
	case *parser.FieldStatementContext:
		v.bldr.AddNode(ctx.GetStart(), ctx.SEMI().GetSymbol(), parser.AdlPField, Field{
			Name: ctx.GetB().GetText(),
		}).Down()
	case *parser.StringStatementContext:
		v.bldr.AddNode(ctx.GetStart(), ctx.GetS(), parser.AdlPJsonStr, ctx.GetS().GetText())
	case *parser.TrueFalseNullContext:
		switch strings.ToLower(ctx.GetKw().GetText()) {
		case "true":
			v.bldr.AddNode(ctx.GetStart(), ctx.GetKw(), parser.AdlPJsonBool, true)
		case "false":
			v.bldr.AddNode(ctx.GetStart(), ctx.GetKw(), parser.AdlPJsonBool, false)
		case "null":
			v.bldr.AddNode(ctx.GetStart(), ctx.GetKw(), parser.AdlPJsonNull, nil)
		default:
			et := Error{Start: ctx.GetStart(), Stop: ctx.GetKw(), Expected: []string{"true", "false", "null"}, Received: ctx.GetKw().GetText()}
			v.bldr.AddNode(ctx.GetStart(), ctx.GetKw(), parser.AdlPERROR, et)
			v.errs.SyntaxErr = append(v.errs.SyntaxErr, et)
		}
	case *parser.NumberStatementContext:
		i, err := strconv.ParseInt(ctx.GetN().GetText(), 10, 64)
		if err == nil {
			v.bldr.AddNode(ctx.GetStart(), ctx.GetN(), parser.AdlPJsonInt, i)
		} else {
			et := Error{Start: ctx.GetStart(), Stop: ctx.GetN(), Expected: []string{"<number>"}, Received: ctx.GetN().GetText()}
			v.bldr.AddNode(ctx.GetStart(), ctx.GetN(), parser.AdlPERROR, et)
			v.errs.SyntaxErr = append(v.errs.SyntaxErr, et)

		}
	case *parser.FloatStatementContext:
		i, err := strconv.ParseFloat(ctx.GetF().GetText(), 64)
		if err == nil {
			v.bldr.AddNode(ctx.GetStart(), ctx.GetF(), parser.AdlPJsonFloat, i)
		} else {
			et := Error{Start: ctx.GetStart(), Stop: ctx.GetF(), Expected: []string{"<float>"}, Received: ctx.GetF().GetText()}
			v.bldr.AddNode(ctx.GetStart(), ctx.GetF(), parser.AdlPERROR, et)
			v.errs.SyntaxErr = append(v.errs.SyntaxErr, et)

		}
	case *parser.ArrayStatementContext:
		v.bldr.AddNode(ctx.GetStart(), ctx.RSQ().GetSymbol(), parser.AdlPJsonArray, JsonArray{}).Down()
	case *parser.ObjStatementContext:
		v.bldr.AddNode(ctx.GetStart(), ctx.RCUR().GetSymbol(), parser.AdlPJsonObj, JsonObj{}).Down()
		// rules name occur on errors
	case *parser.JsonObjStatementContext:
		v.bldr.AddNode(ctx.GetStart(), ctx.GetV().GetStop(), parser.AdlPJsonObjKey, ctx.GetText())
	case *parser.Top_level_statementContext:
		et := Error{Start: ctx.GetStart(), Stop: ctx.GetStop(), Expected: []string{"@|struct|union|annotation"}, Received: ctx.GetStart().GetText()}
		v.bldr.AddNode(ctx.GetStart(), ctx.GetStop(), parser.AdlPERROR, et)
		v.errs.SyntaxErr = append(v.errs.SyntaxErr, et)

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
	case *parser.TypeParameterContext:
	case *parser.TypeExprSimpleContext:
	case *parser.TypeExprGenericContext:
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
	case *parser.JsonObjStatementContext:
	case *parser.Top_level_statementContext:
	default:
		glog.Warningf("Unhandled in <ExitEveryRule case %T:\n", ctx)
	}
}

// func (v *ADLBuildListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
// 	if strings.HasPrefix(msg, "reportAttemptingFullContext") { // TODO remove NewDiagnosticErrorListener and move warning to ReportAmbiguity etc. when getDecisionDescription is make public
// 		v.errs.SyntaxWarning = append(v.errs.SyntaxWarning, fmt.Sprintf("At %d:%d <%s>", line, column, msg))
// 		return
// 	}
// 	t, ok := offendingSymbol.(antlr.Token)
// 	if !ok && e != nil {
// 		t = e.GetOffendingToken()
// 	}
// 	v.errs.ParseErr = append(v.errs.ParseErr, t)
// }

// func (tbl *ADLBuildListener) ReportAmbiguity(
// 	recognizer antlr.Parser,
// 	dfa *antlr.DFA,
// 	startIndex, stopIndex int,
// 	exact bool,
// 	ambigAlts *antlr.BitSet,
// 	configs antlr.ATNConfigSet) {
// 	// panic("ReportAmbiguity")
// 	if tbl.debug {
// 		fmt.Printf("ReportAmbiguity rec:%v dfs:%v start:%d stop:%d, exact:%v, ambigAlts:%v config:%v\n", recognizer, dfa, startIndex, stopIndex, exact, ambigAlts, configs)
// 	}
// }

// func (tbl *ADLBuildListener) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
// 	// panic("ReportAttemptingFullContext")
// 	if tbl.debug {
// 		fmt.Printf("ReportAttemptingFullContext rec:%v dfs:%v start:%d stop:%d, conflictingAlts:%v config:%v\n", recognizer, dfa, startIndex, stopIndex, conflictingAlts, configs)
// 	}
// }

// func (tbl *ADLBuildListener) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex, prediction int, configs antlr.ATNConfigSet) {
// 	//	panic("ReportContextSensitivity")
// }
