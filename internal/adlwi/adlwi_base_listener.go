// Generated from AdlWi.g4 by ANTLR 4.7.

package adlw1 // AdlWi
//import "github.com/wxio/goantlr"
//import "generated code if in another package"

//// Commented out basic implementation for your convenience.

//func Example(s string) {
//  // Setup
//  input := antlr.NewInputStream(s)
//  lexer := adlw1.NewAdlWiLexer(input)
//  stream := antlr.NewCommonTokenStream(lexer, 0)
//  p := adlw1.NewAdlWiParser(stream)

//  // Antlr error listener - turns reports (ambiguity etc) into syntax errors
//  p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))

//  // Custom error listener, register before the parse
//  el := &AdlWiErrorListener{}
//  p.AddErrorListener(el)

//  // Parse - start rule
//  tree := p.Start()

//  // Antlr provided parse tree representation
//  sexpr := antlr.TreesStringTree(tree, nil, p)
//  fmt.Printf("%s\n", sexpr)

//  // Custom listener
//  l := &AdlWiListener{}
//  antlr.ParseTreeWalkerDefault.Walk(l, tree)

//  // Custom visitor
//  v := &AdlWiVisitor{}
//  tree.Accept(v)
// }

//// implemented all listeners methods
//var _ adlw1.AdlWiListener = &AdlWiListener{}
//// implemented specific
//var _ adlw1.AdlEntryListener = &AdlWiListener{}
//var _ adlw1.AdlExitListener = &AdlWiListener{}

//var _ adlw1.JsonEntryListener = &AdlWiListener{}
//var _ adlw1.JsonExitListener = &AdlWiListener{}

//var _ adlw1.ModuleEntryListener = &AdlWiListener{}
//var _ adlw1.ModuleExitListener = &AdlWiListener{}

//var _ adlw1.ImportModuleEntryListener = &AdlWiListener{}
//var _ adlw1.ImportModuleExitListener = &AdlWiListener{}

//var _ adlw1.ImportScopedModuleEntryListener = &AdlWiListener{}
//var _ adlw1.ImportScopedModuleExitListener = &AdlWiListener{}

//var _ adlw1.ImportErrorEntryListener = &AdlWiListener{}
//var _ adlw1.ImportErrorExitListener = &AdlWiListener{}

//var _ adlw1.StructEntryListener = &AdlWiListener{}
//var _ adlw1.StructExitListener = &AdlWiListener{}

//var _ adlw1.UnionEntryListener = &AdlWiListener{}
//var _ adlw1.UnionExitListener = &AdlWiListener{}

//var _ adlw1.TypeEntryListener = &AdlWiListener{}
//var _ adlw1.TypeExitListener = &AdlWiListener{}

//var _ adlw1.NewtypeEntryListener = &AdlWiListener{}
//var _ adlw1.NewtypeExitListener = &AdlWiListener{}

//var _ adlw1.ModAnnoEntryListener = &AdlWiListener{}
//var _ adlw1.ModAnnoExitListener = &AdlWiListener{}

//var _ adlw1.DeclAnnoEntryListener = &AdlWiListener{}
//var _ adlw1.DeclAnnoExitListener = &AdlWiListener{}

//var _ adlw1.FieldAnnoEntryListener = &AdlWiListener{}
//var _ adlw1.FieldAnnoExitListener = &AdlWiListener{}

//var _ adlw1.TypeParamErrorEntryListener = &AdlWiListener{}
//var _ adlw1.TypeParamErrorExitListener = &AdlWiListener{}

//var _ adlw1.TLDErrorEntryListener = &AdlWiListener{}
//var _ adlw1.TLDErrorExitListener = &AdlWiListener{}

//var _ adlw1.FieldEntryListener = &AdlWiListener{}
//var _ adlw1.FieldExitListener = &AdlWiListener{}

//var _ adlw1.AnnotationEntryListener = &AdlWiListener{}
//var _ adlw1.AnnotationExitListener = &AdlWiListener{}

//var _ adlw1.TypeExpr_EntryListener = &AdlWiListener{}
//var _ adlw1.TypeExpr_ExitListener = &AdlWiListener{}

//var _ adlw1.TypeParamsEntryListener = &AdlWiListener{}
//var _ adlw1.TypeParamsExitListener = &AdlWiListener{}

//var _ adlw1.JsonStrEntryListener = &AdlWiListener{}
//var _ adlw1.JsonStrExitListener = &AdlWiListener{}

//var _ adlw1.JsonBoolEntryListener = &AdlWiListener{}
//var _ adlw1.JsonBoolExitListener = &AdlWiListener{}

//var _ adlw1.JsonNullEntryListener = &AdlWiListener{}
//var _ adlw1.JsonNullExitListener = &AdlWiListener{}

//var _ adlw1.JsonIntEntryListener = &AdlWiListener{}
//var _ adlw1.JsonIntExitListener = &AdlWiListener{}

//var _ adlw1.JsonFloatEntryListener = &AdlWiListener{}
//var _ adlw1.JsonFloatExitListener = &AdlWiListener{}

//var _ adlw1.JsonArrayEntryListener = &AdlWiListener{}
//var _ adlw1.JsonArrayExitListener = &AdlWiListener{}

//var _ adlw1.JsonObjEntryListener = &AdlWiListener{}
//var _ adlw1.JsonObjExitListener = &AdlWiListener{}

//var _ adlw1.JsonErrorEntryListener = &AdlWiListener{}
//var _ adlw1.JsonErrorExitListener = &AdlWiListener{}

//type AdlWiListener struct {
//}

//type AdlWiErrorListener struct {
//    Warning string
//    Err     error
//    Debug   bool
//}

// func (cb *AdlWiErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
//  if cb.Debug {
//      fmt.Printf("SyntaxError %d:%d <%s>\n", line, column, msg)
//  }
//  if strings.HasPrefix(msg, "report") { // TODO remove NewDiagnosticErrorListener and move warning to ReportAmbiguity etc. when getDecisionDescription is make public
//      cb.Warning = fmt.Sprintf("At %d:%d <%s>", line, column, msg)
//  } else {
//      cb.Err = fmt.Errorf("SyntaxError %d:%d <%s>", line, column, msg)
//  }
// }
// func (cb *AdlWiErrorListener) ReportAmbiguity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int,
//  exact bool, ambigAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
//  if cb.Debug {
//      fmt.Printf("ReportAmbiguity rec:%v dfs:%v start:%d stop:%d, exact:%v, ambigAlts:%v config:%v\n", recognizer, dfa, startIndex, stopIndex, exact, ambigAlts, configs)
//  }
// }
// func (cb *AdlWiErrorListener) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
//  if cb.Debug {
//      fmt.Printf("ReportAttemptingFullContext rec:%v dfs:%v start:%d stop:%d, conflictingAlts:%v config:%v\n", recognizer, dfa, startIndex, stopIndex, conflictingAlts, configs)
//  }
// }
// func (cb *AdlWiErrorListener) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex, prediction int, configs antlr.ATNConfigSet) {
//  if cb.Debug {
//      fmt.Printf("ReportContextSensitivity rec:%v dfs:%v start:%d stop:%d, prediction:%v configs:%v\n", recognizer, dfa, startIndex, stopIndex, prediction, configs)
//  }
// }

//// antlr.ParseTreeListener implementation.
//// All required.

//func (s *AdlWiListener ) VisitTerminal(node  antlr.TerminalNode) {   }
//func (s *AdlWiListener ) VisitErrorNode(node antlr.ErrorNode)    {   }
//func (s *AdlWiListener ) EnterEveryRule(ctx antlr.ParserRuleContext) {  }
//func (s *AdlWiListener ) ExitEveryRule(ctx antlr.ParserRuleContext) {  }

//// Only implemented as needed.

//func (s *AdlWiListener) EnterAdl(ctx adlw1.*adlw1.AdlContext) {}
//func (s *AdlWiListener) ExitAdl(ctx adlw1.*adlw1.AdlContext) {}

//func (s *AdlWiListener) EnterJson(ctx adlw1.*adlw1.JsonContext) {}
//func (s *AdlWiListener) ExitJson(ctx adlw1.*adlw1.JsonContext) {}

//func (s *AdlWiListener) EnterModule(ctx adlw1.*adlw1.ModuleContext) {}
//func (s *AdlWiListener) ExitModule(ctx adlw1.*adlw1.ModuleContext) {}

//func (s *AdlWiListener) EnterImportModule(ctx adlw1.*adlw1.ImportModuleContext) {}
//func (s *AdlWiListener) ExitImportModule(ctx adlw1.*adlw1.ImportModuleContext) {}

//func (s *AdlWiListener) EnterImportScopedModule(ctx adlw1.*adlw1.ImportScopedModuleContext) {}
//func (s *AdlWiListener) ExitImportScopedModule(ctx adlw1.*adlw1.ImportScopedModuleContext) {}

//func (s *AdlWiListener) EnterImportError(ctx adlw1.*adlw1.ImportErrorContext) {}
//func (s *AdlWiListener) ExitImportError(ctx adlw1.*adlw1.ImportErrorContext) {}

//func (s *AdlWiListener) EnterStruct(ctx adlw1.*adlw1.StructContext) {}
//func (s *AdlWiListener) ExitStruct(ctx adlw1.*adlw1.StructContext) {}

//func (s *AdlWiListener) EnterUnion(ctx adlw1.*adlw1.UnionContext) {}
//func (s *AdlWiListener) ExitUnion(ctx adlw1.*adlw1.UnionContext) {}

//func (s *AdlWiListener) EnterType(ctx adlw1.*adlw1.TypeContext) {}
//func (s *AdlWiListener) ExitType(ctx adlw1.*adlw1.TypeContext) {}

//func (s *AdlWiListener) EnterNewtype(ctx adlw1.*adlw1.NewtypeContext) {}
//func (s *AdlWiListener) ExitNewtype(ctx adlw1.*adlw1.NewtypeContext) {}

//func (s *AdlWiListener) EnterModAnno(ctx adlw1.*adlw1.ModAnnoContext) {}
//func (s *AdlWiListener) ExitModAnno(ctx adlw1.*adlw1.ModAnnoContext) {}

//func (s *AdlWiListener) EnterDeclAnno(ctx adlw1.*adlw1.DeclAnnoContext) {}
//func (s *AdlWiListener) ExitDeclAnno(ctx adlw1.*adlw1.DeclAnnoContext) {}

//func (s *AdlWiListener) EnterFieldAnno(ctx adlw1.*adlw1.FieldAnnoContext) {}
//func (s *AdlWiListener) ExitFieldAnno(ctx adlw1.*adlw1.FieldAnnoContext) {}

//func (s *AdlWiListener) EnterTypeParamError(ctx adlw1.*adlw1.TypeParamErrorContext) {}
//func (s *AdlWiListener) ExitTypeParamError(ctx adlw1.*adlw1.TypeParamErrorContext) {}

//func (s *AdlWiListener) EnterTLDError(ctx adlw1.*adlw1.TLDErrorContext) {}
//func (s *AdlWiListener) ExitTLDError(ctx adlw1.*adlw1.TLDErrorContext) {}

//func (s *AdlWiListener) EnterField(ctx adlw1.*adlw1.FieldContext) {}
//func (s *AdlWiListener) ExitField(ctx adlw1.*adlw1.FieldContext) {}

//func (s *AdlWiListener) EnterAnnotation(ctx adlw1.*adlw1.AnnotationContext) {}
//func (s *AdlWiListener) ExitAnnotation(ctx adlw1.*adlw1.AnnotationContext) {}

//func (s *AdlWiListener) EnterTypeExpr_(ctx adlw1.*adlw1.TypeExpr_Context) {}
//func (s *AdlWiListener) ExitTypeExpr_(ctx adlw1.*adlw1.TypeExpr_Context) {}

//func (s *AdlWiListener) EnterTypeParams(ctx adlw1.*adlw1.TypeParamsContext) {}
//func (s *AdlWiListener) ExitTypeParams(ctx adlw1.*adlw1.TypeParamsContext) {}

//func (s *AdlWiListener) EnterJsonStr(ctx adlw1.*adlw1.JsonStrContext) {}
//func (s *AdlWiListener) ExitJsonStr(ctx adlw1.*adlw1.JsonStrContext) {}

//func (s *AdlWiListener) EnterJsonBool(ctx adlw1.*adlw1.JsonBoolContext) {}
//func (s *AdlWiListener) ExitJsonBool(ctx adlw1.*adlw1.JsonBoolContext) {}

//func (s *AdlWiListener) EnterJsonNull(ctx adlw1.*adlw1.JsonNullContext) {}
//func (s *AdlWiListener) ExitJsonNull(ctx adlw1.*adlw1.JsonNullContext) {}

//func (s *AdlWiListener) EnterJsonInt(ctx adlw1.*adlw1.JsonIntContext) {}
//func (s *AdlWiListener) ExitJsonInt(ctx adlw1.*adlw1.JsonIntContext) {}

//func (s *AdlWiListener) EnterJsonFloat(ctx adlw1.*adlw1.JsonFloatContext) {}
//func (s *AdlWiListener) ExitJsonFloat(ctx adlw1.*adlw1.JsonFloatContext) {}

//func (s *AdlWiListener) EnterJsonArray(ctx adlw1.*adlw1.JsonArrayContext) {}
//func (s *AdlWiListener) ExitJsonArray(ctx adlw1.*adlw1.JsonArrayContext) {}

//func (s *AdlWiListener) EnterJsonObj(ctx adlw1.*adlw1.JsonObjContext) {}
//func (s *AdlWiListener) ExitJsonObj(ctx adlw1.*adlw1.JsonObjContext) {}

//func (s *AdlWiListener) EnterJsonError(ctx adlw1.*adlw1.JsonErrorContext) {}
//func (s *AdlWiListener) ExitJsonError(ctx adlw1.*adlw1.JsonErrorContext) {}
