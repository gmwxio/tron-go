// Generated from AdlWo.g4 by ANTLR 4.7.

package adlwo // AdlWo
//import "github.com/wxio/goantlr"
//import "generated code if in another package"

//// Commented out basic implementation for your convenience.

//func Example(s string) {
//  // Setup
//  input := antlr.NewInputStream(s)
//  lexer := adlwo.NewAdlWoLexer(input)
//  stream := antlr.NewCommonTokenStream(lexer, 0)
//  p := adlwo.NewAdlWoParser(stream)

//  // Antlr error listener - turns reports (ambiguity etc) into syntax errors
//  p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))

//  // Custom error listener, register before the parse
//  el := &AdlWoErrorListener{}
//  p.AddErrorListener(el)

//  // Parse - start rule
//  tree := p.Start()

//  // Antlr provided parse tree representation
//  sexpr := antlr.TreesStringTree(tree, nil, p)
//  fmt.Printf("%s\n", sexpr)

//  // Custom listener
//  l := &AdlWoListener{}
//  antlr.ParseTreeWalkerDefault.Walk(l, tree)

//  // Custom visitor
//  v := &AdlWoVisitor{}
//  tree.Accept(v)
// }

//// implemented all listeners methods
//var _ adlwo.AdlWoListener = &AdlWoListener{}
//// implemented specific
//var _ adlwo.AdlEntryListener = &AdlWoListener{}
//var _ adlwo.AdlExitListener = &AdlWoListener{}

//var _ adlwo.JsonEntryListener = &AdlWoListener{}
//var _ adlwo.JsonExitListener = &AdlWoListener{}

//var _ adlwo.ModuleEntryListener = &AdlWoListener{}
//var _ adlwo.ModuleExitListener = &AdlWoListener{}

//var _ adlwo.ImportModuleEntryListener = &AdlWoListener{}
//var _ adlwo.ImportModuleExitListener = &AdlWoListener{}

//var _ adlwo.ImportScopedModuleEntryListener = &AdlWoListener{}
//var _ adlwo.ImportScopedModuleExitListener = &AdlWoListener{}

//var _ adlwo.StructEntryListener = &AdlWoListener{}
//var _ adlwo.StructExitListener = &AdlWoListener{}

//var _ adlwo.UnionEntryListener = &AdlWoListener{}
//var _ adlwo.UnionExitListener = &AdlWoListener{}

//var _ adlwo.TypeEntryListener = &AdlWoListener{}
//var _ adlwo.TypeExitListener = &AdlWoListener{}

//var _ adlwo.NewtypeEntryListener = &AdlWoListener{}
//var _ adlwo.NewtypeExitListener = &AdlWoListener{}

//var _ adlwo.ModAnnoEntryListener = &AdlWoListener{}
//var _ adlwo.ModAnnoExitListener = &AdlWoListener{}

//var _ adlwo.DeclAnnoEntryListener = &AdlWoListener{}
//var _ adlwo.DeclAnnoExitListener = &AdlWoListener{}

//var _ adlwo.FieldAnnoEntryListener = &AdlWoListener{}
//var _ adlwo.FieldAnnoExitListener = &AdlWoListener{}

//var _ adlwo.FieldEntryListener = &AdlWoListener{}
//var _ adlwo.FieldExitListener = &AdlWoListener{}

//var _ adlwo.AnnotationEntryListener = &AdlWoListener{}
//var _ adlwo.AnnotationExitListener = &AdlWoListener{}

//var _ adlwo.TypeExpr_EntryListener = &AdlWoListener{}
//var _ adlwo.TypeExpr_ExitListener = &AdlWoListener{}

//var _ adlwo.TypeParamsEntryListener = &AdlWoListener{}
//var _ adlwo.TypeParamsExitListener = &AdlWoListener{}

//var _ adlwo.JsonStrEntryListener = &AdlWoListener{}
//var _ adlwo.JsonStrExitListener = &AdlWoListener{}

//var _ adlwo.JsonBoolEntryListener = &AdlWoListener{}
//var _ adlwo.JsonBoolExitListener = &AdlWoListener{}

//var _ adlwo.JsonNullEntryListener = &AdlWoListener{}
//var _ adlwo.JsonNullExitListener = &AdlWoListener{}

//var _ adlwo.JsonIntEntryListener = &AdlWoListener{}
//var _ adlwo.JsonIntExitListener = &AdlWoListener{}

//var _ adlwo.JsonFloatEntryListener = &AdlWoListener{}
//var _ adlwo.JsonFloatExitListener = &AdlWoListener{}

//var _ adlwo.JsonArrayEntryListener = &AdlWoListener{}
//var _ adlwo.JsonArrayExitListener = &AdlWoListener{}

//var _ adlwo.JsonObjEntryListener = &AdlWoListener{}
//var _ adlwo.JsonObjExitListener = &AdlWoListener{}

//type AdlWoListener struct {
//}

//type AdlWoErrorListener struct {
//    Warning string
//    Err     error
//    Debug   bool
//}

// func (cb *AdlWoErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
//  if cb.Debug {
//      fmt.Printf("SyntaxError %d:%d <%s>\n", line, column, msg)
//  }
//  if strings.HasPrefix(msg, "report") { // TODO remove NewDiagnosticErrorListener and move warning to ReportAmbiguity etc. when getDecisionDescription is make public
//      cb.Warning = fmt.Sprintf("At %d:%d <%s>", line, column, msg)
//  } else {
//      cb.Err = fmt.Errorf("SyntaxError %d:%d <%s>", line, column, msg)
//  }
// }
// func (cb *AdlWoErrorListener) ReportAmbiguity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int,
//  exact bool, ambigAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
//  if cb.Debug {
//      fmt.Printf("ReportAmbiguity rec:%v dfs:%v start:%d stop:%d, exact:%v, ambigAlts:%v config:%v\n", recognizer, dfa, startIndex, stopIndex, exact, ambigAlts, configs)
//  }
// }
// func (cb *AdlWoErrorListener) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
//  if cb.Debug {
//      fmt.Printf("ReportAttemptingFullContext rec:%v dfs:%v start:%d stop:%d, conflictingAlts:%v config:%v\n", recognizer, dfa, startIndex, stopIndex, conflictingAlts, configs)
//  }
// }
// func (cb *AdlWoErrorListener) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex, prediction int, configs antlr.ATNConfigSet) {
//  if cb.Debug {
//      fmt.Printf("ReportContextSensitivity rec:%v dfs:%v start:%d stop:%d, prediction:%v configs:%v\n", recognizer, dfa, startIndex, stopIndex, prediction, configs)
//  }
// }

//// antlr.ParseTreeListener implementation.
//// All required.

//func (s *AdlWoListener ) VisitTerminal(node  antlr.TerminalNode) {   }
//func (s *AdlWoListener ) VisitErrorNode(node antlr.ErrorNode)    {   }
//func (s *AdlWoListener ) EnterEveryRule(ctx antlr.ParserRuleContext) {  }
//func (s *AdlWoListener ) ExitEveryRule(ctx antlr.ParserRuleContext) {  }

//// Only implemented as needed.

//func (s *AdlWoListener) EnterAdl(ctx adlwo.*adlwo.AdlContext) {}
//func (s *AdlWoListener) ExitAdl(ctx adlwo.*adlwo.AdlContext) {}

//func (s *AdlWoListener) EnterJson(ctx adlwo.*adlwo.JsonContext) {}
//func (s *AdlWoListener) ExitJson(ctx adlwo.*adlwo.JsonContext) {}

//func (s *AdlWoListener) EnterModule(ctx adlwo.*adlwo.ModuleContext) {}
//func (s *AdlWoListener) ExitModule(ctx adlwo.*adlwo.ModuleContext) {}

//func (s *AdlWoListener) EnterImportModule(ctx adlwo.*adlwo.ImportModuleContext) {}
//func (s *AdlWoListener) ExitImportModule(ctx adlwo.*adlwo.ImportModuleContext) {}

//func (s *AdlWoListener) EnterImportScopedModule(ctx adlwo.*adlwo.ImportScopedModuleContext) {}
//func (s *AdlWoListener) ExitImportScopedModule(ctx adlwo.*adlwo.ImportScopedModuleContext) {}

//func (s *AdlWoListener) EnterStruct(ctx adlwo.*adlwo.StructContext) {}
//func (s *AdlWoListener) ExitStruct(ctx adlwo.*adlwo.StructContext) {}

//func (s *AdlWoListener) EnterUnion(ctx adlwo.*adlwo.UnionContext) {}
//func (s *AdlWoListener) ExitUnion(ctx adlwo.*adlwo.UnionContext) {}

//func (s *AdlWoListener) EnterType(ctx adlwo.*adlwo.TypeContext) {}
//func (s *AdlWoListener) ExitType(ctx adlwo.*adlwo.TypeContext) {}

//func (s *AdlWoListener) EnterNewtype(ctx adlwo.*adlwo.NewtypeContext) {}
//func (s *AdlWoListener) ExitNewtype(ctx adlwo.*adlwo.NewtypeContext) {}

//func (s *AdlWoListener) EnterModAnno(ctx adlwo.*adlwo.ModAnnoContext) {}
//func (s *AdlWoListener) ExitModAnno(ctx adlwo.*adlwo.ModAnnoContext) {}

//func (s *AdlWoListener) EnterDeclAnno(ctx adlwo.*adlwo.DeclAnnoContext) {}
//func (s *AdlWoListener) ExitDeclAnno(ctx adlwo.*adlwo.DeclAnnoContext) {}

//func (s *AdlWoListener) EnterFieldAnno(ctx adlwo.*adlwo.FieldAnnoContext) {}
//func (s *AdlWoListener) ExitFieldAnno(ctx adlwo.*adlwo.FieldAnnoContext) {}

//func (s *AdlWoListener) EnterField(ctx adlwo.*adlwo.FieldContext) {}
//func (s *AdlWoListener) ExitField(ctx adlwo.*adlwo.FieldContext) {}

//func (s *AdlWoListener) EnterAnnotation(ctx adlwo.*adlwo.AnnotationContext) {}
//func (s *AdlWoListener) ExitAnnotation(ctx adlwo.*adlwo.AnnotationContext) {}

//func (s *AdlWoListener) EnterTypeExpr_(ctx adlwo.*adlwo.TypeExpr_Context) {}
//func (s *AdlWoListener) ExitTypeExpr_(ctx adlwo.*adlwo.TypeExpr_Context) {}

//func (s *AdlWoListener) EnterTypeParams(ctx adlwo.*adlwo.TypeParamsContext) {}
//func (s *AdlWoListener) ExitTypeParams(ctx adlwo.*adlwo.TypeParamsContext) {}

//func (s *AdlWoListener) EnterJsonStr(ctx adlwo.*adlwo.JsonStrContext) {}
//func (s *AdlWoListener) ExitJsonStr(ctx adlwo.*adlwo.JsonStrContext) {}

//func (s *AdlWoListener) EnterJsonBool(ctx adlwo.*adlwo.JsonBoolContext) {}
//func (s *AdlWoListener) ExitJsonBool(ctx adlwo.*adlwo.JsonBoolContext) {}

//func (s *AdlWoListener) EnterJsonNull(ctx adlwo.*adlwo.JsonNullContext) {}
//func (s *AdlWoListener) ExitJsonNull(ctx adlwo.*adlwo.JsonNullContext) {}

//func (s *AdlWoListener) EnterJsonInt(ctx adlwo.*adlwo.JsonIntContext) {}
//func (s *AdlWoListener) ExitJsonInt(ctx adlwo.*adlwo.JsonIntContext) {}

//func (s *AdlWoListener) EnterJsonFloat(ctx adlwo.*adlwo.JsonFloatContext) {}
//func (s *AdlWoListener) ExitJsonFloat(ctx adlwo.*adlwo.JsonFloatContext) {}

//func (s *AdlWoListener) EnterJsonArray(ctx adlwo.*adlwo.JsonArrayContext) {}
//func (s *AdlWoListener) ExitJsonArray(ctx adlwo.*adlwo.JsonArrayContext) {}

//func (s *AdlWoListener) EnterJsonObj(ctx adlwo.*adlwo.JsonObjContext) {}
//func (s *AdlWoListener) ExitJsonObj(ctx adlwo.*adlwo.JsonObjContext) {}
