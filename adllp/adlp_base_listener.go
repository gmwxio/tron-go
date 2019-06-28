// Generated from AdlP.g4 by ANTLR 4.7.

package adllp // AdlP
//import "github.com/antlr/antlr4/runtime/Go/antlr"
//import "generated code if in another package"

//// Commented out basic implementation for your convenience.

//func Example(s string) {
//  // Setup
//  input := antlr.NewInputStream(s)
//  lexer := adllp.NewAdlPLexer(input)
//  stream := antlr.NewCommonTokenStream(lexer, 0)
//  p := adllp.NewAdlPParser(stream)

//  // Antlr error listener - turns reports (ambiguity etc) into syntax errors
//  p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))

//  // Custom error listener, register before the parse
//  el := &AdlPErrorListener{}
//  p.AddErrorListener(el)

//  // Parse - start rule
//  tree := p.Start()

//  // Antlr provided parse tree representation
//  sexpr := antlr.TreesStringTree(tree, nil, p)
//  fmt.Printf("%s\n", sexpr)

//  // Custom listener
//  l := &AdlPListener{}
//  antlr.ParseTreeWalkerDefault.Walk(l, tree)

//  // Custom visitor
//  v := &AdlPVisitor{}
//  tree.Accept(v)
// }

//// implemented all listeners methods
//var _ adllp.AdlPListener = &AdlPListener{}
//// implemented specific
//var _ adllp.AdlEntryListener = &AdlPListener{}
//var _ adllp.AdlExitListener = &AdlPListener{}

//var _ adllp.ModuleStatementEntryListener = &AdlPListener{}
//var _ adllp.ModuleStatementExitListener = &AdlPListener{}

//var _ adllp.ImportStatementEntryListener = &AdlPListener{}
//var _ adllp.ImportStatementExitListener = &AdlPListener{}

//var _ adllp.LocalAnnoEntryListener = &AdlPListener{}
//var _ adllp.LocalAnnoExitListener = &AdlPListener{}

//var _ adllp.DocAnnoEntryListener = &AdlPListener{}
//var _ adllp.DocAnnoExitListener = &AdlPListener{}

//var _ adllp.StructOrUnionEntryListener = &AdlPListener{}
//var _ adllp.StructOrUnionExitListener = &AdlPListener{}

//var _ adllp.TypeOrNewtypeEntryListener = &AdlPListener{}
//var _ adllp.TypeOrNewtypeExitListener = &AdlPListener{}

//var _ adllp.ModuleAnnotationEntryListener = &AdlPListener{}
//var _ adllp.ModuleAnnotationExitListener = &AdlPListener{}

//var _ adllp.DeclAnnotationEntryListener = &AdlPListener{}
//var _ adllp.DeclAnnotationExitListener = &AdlPListener{}

//var _ adllp.FieldAnnotationEntryListener = &AdlPListener{}
//var _ adllp.FieldAnnotationExitListener = &AdlPListener{}

//var _ adllp.TypeParameterEntryListener = &AdlPListener{}
//var _ adllp.TypeParameterExitListener = &AdlPListener{}

//var _ adllp.ErrorTypeParamEntryListener = &AdlPListener{}
//var _ adllp.ErrorTypeParamExitListener = &AdlPListener{}

//var _ adllp.TypeParamErrorEntryListener = &AdlPListener{}
//var _ adllp.TypeParamErrorExitListener = &AdlPListener{}

//var _ adllp.TypeExpressionEntryListener = &AdlPListener{}
//var _ adllp.TypeExpressionExitListener = &AdlPListener{}

//var _ adllp.TypeExpressionElemEntryListener = &AdlPListener{}
//var _ adllp.TypeExpressionElemExitListener = &AdlPListener{}

//var _ adllp.FieldStatementEntryListener = &AdlPListener{}
//var _ adllp.FieldStatementExitListener = &AdlPListener{}

//var _ adllp.StringStatementEntryListener = &AdlPListener{}
//var _ adllp.StringStatementExitListener = &AdlPListener{}

//var _ adllp.TrueFalseNullEntryListener = &AdlPListener{}
//var _ adllp.TrueFalseNullExitListener = &AdlPListener{}

//var _ adllp.NumberStatementEntryListener = &AdlPListener{}
//var _ adllp.NumberStatementExitListener = &AdlPListener{}

//var _ adllp.FloatStatementEntryListener = &AdlPListener{}
//var _ adllp.FloatStatementExitListener = &AdlPListener{}

//var _ adllp.ArrayStatementEntryListener = &AdlPListener{}
//var _ adllp.ArrayStatementExitListener = &AdlPListener{}

//var _ adllp.ObjStatementEntryListener = &AdlPListener{}
//var _ adllp.ObjStatementExitListener = &AdlPListener{}

//type AdlPListener struct {
//}

//type AdlPErrorListener struct {
//    Warning string
//    Err     error
//    Debug   bool
//}

// func (cb *AdlPErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
//  if cb.Debug {
//      fmt.Printf("SyntaxError %d:%d <%s>\n", line, column, msg)
//  }
//  if strings.HasPrefix(msg, "report") { // TODO remove NewDiagnosticErrorListener and move warning to ReportAmbiguity etc. when getDecisionDescription is make public
//      cb.Warning = fmt.Sprintf("At %d:%d <%s>", line, column, msg)
//  } else {
//      cb.Err = fmt.Errorf("SyntaxError %d:%d <%s>", line, column, msg)
//  }
// }
// func (cb *AdlPErrorListener) ReportAmbiguity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int,
//  exact bool, ambigAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
//  if cb.Debug {
//      fmt.Printf("ReportAmbiguity rec:%v dfs:%v start:%d stop:%d, exact:%v, ambigAlts:%v config:%v\n", recognizer, dfa, startIndex, stopIndex, exact, ambigAlts, configs)
//  }
// }
// func (cb *AdlPErrorListener) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
//  if cb.Debug {
//      fmt.Printf("ReportAttemptingFullContext rec:%v dfs:%v start:%d stop:%d, conflictingAlts:%v config:%v\n", recognizer, dfa, startIndex, stopIndex, conflictingAlts, configs)
//  }
// }
// func (cb *AdlPErrorListener) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex, prediction int, configs antlr.ATNConfigSet) {
//  if cb.Debug {
//      fmt.Printf("ReportContextSensitivity rec:%v dfs:%v start:%d stop:%d, prediction:%v configs:%v\n", recognizer, dfa, startIndex, stopIndex, prediction, configs)
//  }
// }

//// antlr.ParseTreeListener implementation.
//// All required.

//func (s *AdlPListener ) VisitTerminal(node  antlr.TerminalNode) {   }
//func (s *AdlPListener ) VisitErrorNode(node antlr.ErrorNode)    {   }
//func (s *AdlPListener ) EnterEveryRule(ctx antlr.ParserRuleContext) {  }
//func (s *AdlPListener ) ExitEveryRule(ctx antlr.ParserRuleContext) {  }

//// Only implemented as needed.

//func (s *AdlPListener) EnterAdl(ctx adllp.*adllp.AdlContext) {}
//func (s *AdlPListener) ExitAdl(ctx adllp.*adllp.AdlContext) {}

//func (s *AdlPListener) EnterModuleStatement(ctx adllp.*adllp.ModuleStatementContext) {}
//func (s *AdlPListener) ExitModuleStatement(ctx adllp.*adllp.ModuleStatementContext) {}

//func (s *AdlPListener) EnterImportStatement(ctx adllp.*adllp.ImportStatementContext) {}
//func (s *AdlPListener) ExitImportStatement(ctx adllp.*adllp.ImportStatementContext) {}

//func (s *AdlPListener) EnterLocalAnno(ctx adllp.*adllp.LocalAnnoContext) {}
//func (s *AdlPListener) ExitLocalAnno(ctx adllp.*adllp.LocalAnnoContext) {}

//func (s *AdlPListener) EnterDocAnno(ctx adllp.*adllp.DocAnnoContext) {}
//func (s *AdlPListener) ExitDocAnno(ctx adllp.*adllp.DocAnnoContext) {}

//func (s *AdlPListener) EnterStructOrUnion(ctx adllp.*adllp.StructOrUnionContext) {}
//func (s *AdlPListener) ExitStructOrUnion(ctx adllp.*adllp.StructOrUnionContext) {}

//func (s *AdlPListener) EnterTypeOrNewtype(ctx adllp.*adllp.TypeOrNewtypeContext) {}
//func (s *AdlPListener) ExitTypeOrNewtype(ctx adllp.*adllp.TypeOrNewtypeContext) {}

//func (s *AdlPListener) EnterModuleAnnotation(ctx adllp.*adllp.ModuleAnnotationContext) {}
//func (s *AdlPListener) ExitModuleAnnotation(ctx adllp.*adllp.ModuleAnnotationContext) {}

//func (s *AdlPListener) EnterDeclAnnotation(ctx adllp.*adllp.DeclAnnotationContext) {}
//func (s *AdlPListener) ExitDeclAnnotation(ctx adllp.*adllp.DeclAnnotationContext) {}

//func (s *AdlPListener) EnterFieldAnnotation(ctx adllp.*adllp.FieldAnnotationContext) {}
//func (s *AdlPListener) ExitFieldAnnotation(ctx adllp.*adllp.FieldAnnotationContext) {}

//func (s *AdlPListener) EnterTypeParameter(ctx adllp.*adllp.TypeParameterContext) {}
//func (s *AdlPListener) ExitTypeParameter(ctx adllp.*adllp.TypeParameterContext) {}

//func (s *AdlPListener) EnterErrorTypeParam(ctx adllp.*adllp.ErrorTypeParamContext) {}
//func (s *AdlPListener) ExitErrorTypeParam(ctx adllp.*adllp.ErrorTypeParamContext) {}

//func (s *AdlPListener) EnterTypeParamError(ctx adllp.*adllp.TypeParamErrorContext) {}
//func (s *AdlPListener) ExitTypeParamError(ctx adllp.*adllp.TypeParamErrorContext) {}

//func (s *AdlPListener) EnterTypeExpression(ctx adllp.*adllp.TypeExpressionContext) {}
//func (s *AdlPListener) ExitTypeExpression(ctx adllp.*adllp.TypeExpressionContext) {}

//func (s *AdlPListener) EnterTypeExpressionElem(ctx adllp.*adllp.TypeExpressionElemContext) {}
//func (s *AdlPListener) ExitTypeExpressionElem(ctx adllp.*adllp.TypeExpressionElemContext) {}

//func (s *AdlPListener) EnterFieldStatement(ctx adllp.*adllp.FieldStatementContext) {}
//func (s *AdlPListener) ExitFieldStatement(ctx adllp.*adllp.FieldStatementContext) {}

//func (s *AdlPListener) EnterStringStatement(ctx adllp.*adllp.StringStatementContext) {}
//func (s *AdlPListener) ExitStringStatement(ctx adllp.*adllp.StringStatementContext) {}

//func (s *AdlPListener) EnterTrueFalseNull(ctx adllp.*adllp.TrueFalseNullContext) {}
//func (s *AdlPListener) ExitTrueFalseNull(ctx adllp.*adllp.TrueFalseNullContext) {}

//func (s *AdlPListener) EnterNumberStatement(ctx adllp.*adllp.NumberStatementContext) {}
//func (s *AdlPListener) ExitNumberStatement(ctx adllp.*adllp.NumberStatementContext) {}

//func (s *AdlPListener) EnterFloatStatement(ctx adllp.*adllp.FloatStatementContext) {}
//func (s *AdlPListener) ExitFloatStatement(ctx adllp.*adllp.FloatStatementContext) {}

//func (s *AdlPListener) EnterArrayStatement(ctx adllp.*adllp.ArrayStatementContext) {}
//func (s *AdlPListener) ExitArrayStatement(ctx adllp.*adllp.ArrayStatementContext) {}

//func (s *AdlPListener) EnterObjStatement(ctx adllp.*adllp.ObjStatementContext) {}
//func (s *AdlPListener) ExitObjStatement(ctx adllp.*adllp.ObjStatementContext) {}
