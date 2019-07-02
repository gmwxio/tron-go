// Generated from AdlWi.g4 by ANTLR 4.7.

package adlwi // AdlWi
//import "github.com/wxio/goantlr"
//import "generated code if in another package"

//// Commented out basic implementation for your convenience.

//func Example(s string) {
//  // Setup
//  input := antlr.NewInputStream(s)
//  lexer := adlwi.NewAdlWiLexer(input)
//  stream := antlr.NewCommonTokenStream(lexer, 0)
//  p := adlwi.NewAdlWiParser(stream)

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
//var _ adlwi.AdlWiListener = &AdlWiListener{}
//// implemented specific
//var _ adlwi.AdlEntryListener = &AdlWiListener{}
//var _ adlwi.AdlExitListener = &AdlWiListener{}

//var _ adlwi.JsonEntryListener = &AdlWiListener{}
//var _ adlwi.JsonExitListener = &AdlWiListener{}

//var _ adlwi.ModuleEntryListener = &AdlWiListener{}
//var _ adlwi.ModuleExitListener = &AdlWiListener{}

//var _ adlwi.ImportModuleEntryListener = &AdlWiListener{}
//var _ adlwi.ImportModuleExitListener = &AdlWiListener{}

//var _ adlwi.ImportScopedModuleEntryListener = &AdlWiListener{}
//var _ adlwi.ImportScopedModuleExitListener = &AdlWiListener{}

//var _ adlwi.ImportErrorEntryListener = &AdlWiListener{}
//var _ adlwi.ImportErrorExitListener = &AdlWiListener{}

//var _ adlwi.StructEntryListener = &AdlWiListener{}
//var _ adlwi.StructExitListener = &AdlWiListener{}

//var _ adlwi.UnionEntryListener = &AdlWiListener{}
//var _ adlwi.UnionExitListener = &AdlWiListener{}

//var _ adlwi.TypeEntryListener = &AdlWiListener{}
//var _ adlwi.TypeExitListener = &AdlWiListener{}

//var _ adlwi.NewtypeEntryListener = &AdlWiListener{}
//var _ adlwi.NewtypeExitListener = &AdlWiListener{}

//var _ adlwi.ModAnnoEntryListener = &AdlWiListener{}
//var _ adlwi.ModAnnoExitListener = &AdlWiListener{}

//var _ adlwi.DeclAnnoEntryListener = &AdlWiListener{}
//var _ adlwi.DeclAnnoExitListener = &AdlWiListener{}

//var _ adlwi.FieldAnnoEntryListener = &AdlWiListener{}
//var _ adlwi.FieldAnnoExitListener = &AdlWiListener{}

//var _ adlwi.TypeParamErrorEntryListener = &AdlWiListener{}
//var _ adlwi.TypeParamErrorExitListener = &AdlWiListener{}

//var _ adlwi.TLDErrorEntryListener = &AdlWiListener{}
//var _ adlwi.TLDErrorExitListener = &AdlWiListener{}

//var _ adlwi.FieldEntryListener = &AdlWiListener{}
//var _ adlwi.FieldExitListener = &AdlWiListener{}

//var _ adlwi.AnnotationEntryListener = &AdlWiListener{}
//var _ adlwi.AnnotationExitListener = &AdlWiListener{}

//var _ adlwi.TypeExpr_EntryListener = &AdlWiListener{}
//var _ adlwi.TypeExpr_ExitListener = &AdlWiListener{}

//var _ adlwi.TypeParamsEntryListener = &AdlWiListener{}
//var _ adlwi.TypeParamsExitListener = &AdlWiListener{}

//var _ adlwi.JsonStrEntryListener = &AdlWiListener{}
//var _ adlwi.JsonStrExitListener = &AdlWiListener{}

//var _ adlwi.JsonBoolEntryListener = &AdlWiListener{}
//var _ adlwi.JsonBoolExitListener = &AdlWiListener{}

//var _ adlwi.JsonNullEntryListener = &AdlWiListener{}
//var _ adlwi.JsonNullExitListener = &AdlWiListener{}

//var _ adlwi.JsonIntEntryListener = &AdlWiListener{}
//var _ adlwi.JsonIntExitListener = &AdlWiListener{}

//var _ adlwi.JsonFloatEntryListener = &AdlWiListener{}
//var _ adlwi.JsonFloatExitListener = &AdlWiListener{}

//var _ adlwi.JsonArrayEntryListener = &AdlWiListener{}
//var _ adlwi.JsonArrayExitListener = &AdlWiListener{}

//var _ adlwi.JsonObjEntryListener = &AdlWiListener{}
//var _ adlwi.JsonObjExitListener = &AdlWiListener{}

//var _ adlwi.JsonErrorEntryListener = &AdlWiListener{}
//var _ adlwi.JsonErrorExitListener = &AdlWiListener{}

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

//func (s *AdlWiListener) EnterAdl(ctx adlwi.*adlwi.AdlContext) {}
//func (s *AdlWiListener) ExitAdl(ctx adlwi.*adlwi.AdlContext) {}

//func (s *AdlWiListener) EnterJson(ctx adlwi.*adlwi.JsonContext) {}
//func (s *AdlWiListener) ExitJson(ctx adlwi.*adlwi.JsonContext) {}

//func (s *AdlWiListener) EnterModule(ctx adlwi.*adlwi.ModuleContext) {}
//func (s *AdlWiListener) ExitModule(ctx adlwi.*adlwi.ModuleContext) {}

//func (s *AdlWiListener) EnterImportModule(ctx adlwi.*adlwi.ImportModuleContext) {}
//func (s *AdlWiListener) ExitImportModule(ctx adlwi.*adlwi.ImportModuleContext) {}

//func (s *AdlWiListener) EnterImportScopedModule(ctx adlwi.*adlwi.ImportScopedModuleContext) {}
//func (s *AdlWiListener) ExitImportScopedModule(ctx adlwi.*adlwi.ImportScopedModuleContext) {}

//func (s *AdlWiListener) EnterImportError(ctx adlwi.*adlwi.ImportErrorContext) {}
//func (s *AdlWiListener) ExitImportError(ctx adlwi.*adlwi.ImportErrorContext) {}

//func (s *AdlWiListener) EnterStruct(ctx adlwi.*adlwi.StructContext) {}
//func (s *AdlWiListener) ExitStruct(ctx adlwi.*adlwi.StructContext) {}

//func (s *AdlWiListener) EnterUnion(ctx adlwi.*adlwi.UnionContext) {}
//func (s *AdlWiListener) ExitUnion(ctx adlwi.*adlwi.UnionContext) {}

//func (s *AdlWiListener) EnterType(ctx adlwi.*adlwi.TypeContext) {}
//func (s *AdlWiListener) ExitType(ctx adlwi.*adlwi.TypeContext) {}

//func (s *AdlWiListener) EnterNewtype(ctx adlwi.*adlwi.NewtypeContext) {}
//func (s *AdlWiListener) ExitNewtype(ctx adlwi.*adlwi.NewtypeContext) {}

//func (s *AdlWiListener) EnterModAnno(ctx adlwi.*adlwi.ModAnnoContext) {}
//func (s *AdlWiListener) ExitModAnno(ctx adlwi.*adlwi.ModAnnoContext) {}

//func (s *AdlWiListener) EnterDeclAnno(ctx adlwi.*adlwi.DeclAnnoContext) {}
//func (s *AdlWiListener) ExitDeclAnno(ctx adlwi.*adlwi.DeclAnnoContext) {}

//func (s *AdlWiListener) EnterFieldAnno(ctx adlwi.*adlwi.FieldAnnoContext) {}
//func (s *AdlWiListener) ExitFieldAnno(ctx adlwi.*adlwi.FieldAnnoContext) {}

//func (s *AdlWiListener) EnterTypeParamError(ctx adlwi.*adlwi.TypeParamErrorContext) {}
//func (s *AdlWiListener) ExitTypeParamError(ctx adlwi.*adlwi.TypeParamErrorContext) {}

//func (s *AdlWiListener) EnterTLDError(ctx adlwi.*adlwi.TLDErrorContext) {}
//func (s *AdlWiListener) ExitTLDError(ctx adlwi.*adlwi.TLDErrorContext) {}

//func (s *AdlWiListener) EnterField(ctx adlwi.*adlwi.FieldContext) {}
//func (s *AdlWiListener) ExitField(ctx adlwi.*adlwi.FieldContext) {}

//func (s *AdlWiListener) EnterAnnotation(ctx adlwi.*adlwi.AnnotationContext) {}
//func (s *AdlWiListener) ExitAnnotation(ctx adlwi.*adlwi.AnnotationContext) {}

//func (s *AdlWiListener) EnterTypeExpr_(ctx adlwi.*adlwi.TypeExpr_Context) {}
//func (s *AdlWiListener) ExitTypeExpr_(ctx adlwi.*adlwi.TypeExpr_Context) {}

//func (s *AdlWiListener) EnterTypeParams(ctx adlwi.*adlwi.TypeParamsContext) {}
//func (s *AdlWiListener) ExitTypeParams(ctx adlwi.*adlwi.TypeParamsContext) {}

//func (s *AdlWiListener) EnterJsonStr(ctx adlwi.*adlwi.JsonStrContext) {}
//func (s *AdlWiListener) ExitJsonStr(ctx adlwi.*adlwi.JsonStrContext) {}

//func (s *AdlWiListener) EnterJsonBool(ctx adlwi.*adlwi.JsonBoolContext) {}
//func (s *AdlWiListener) ExitJsonBool(ctx adlwi.*adlwi.JsonBoolContext) {}

//func (s *AdlWiListener) EnterJsonNull(ctx adlwi.*adlwi.JsonNullContext) {}
//func (s *AdlWiListener) ExitJsonNull(ctx adlwi.*adlwi.JsonNullContext) {}

//func (s *AdlWiListener) EnterJsonInt(ctx adlwi.*adlwi.JsonIntContext) {}
//func (s *AdlWiListener) ExitJsonInt(ctx adlwi.*adlwi.JsonIntContext) {}

//func (s *AdlWiListener) EnterJsonFloat(ctx adlwi.*adlwi.JsonFloatContext) {}
//func (s *AdlWiListener) ExitJsonFloat(ctx adlwi.*adlwi.JsonFloatContext) {}

//func (s *AdlWiListener) EnterJsonArray(ctx adlwi.*adlwi.JsonArrayContext) {}
//func (s *AdlWiListener) ExitJsonArray(ctx adlwi.*adlwi.JsonArrayContext) {}

//func (s *AdlWiListener) EnterJsonObj(ctx adlwi.*adlwi.JsonObjContext) {}
//func (s *AdlWiListener) ExitJsonObj(ctx adlwi.*adlwi.JsonObjContext) {}

//func (s *AdlWiListener) EnterJsonError(ctx adlwi.*adlwi.JsonErrorContext) {}
//func (s *AdlWiListener) ExitJsonError(ctx adlwi.*adlwi.JsonErrorContext) {}
