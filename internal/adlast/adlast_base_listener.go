// Generated from AdlAst.g4 by ANTLR 4.7.

package adlast // AdlAst
//import "github.com/wxio/goantlr"
//import "generated code if in another package"

//// Commented out basic implementation for your convenience.

//func Example(s string) {
//  // Setup
//  input := antlr.NewInputStream(s)
//  lexer := adlast.NewAdlAstLexer(input)
//  stream := antlr.NewCommonTokenStream(lexer, 0)
//  p := adlast.NewAdlAstParser(stream)

//  // Antlr error listener - turns reports (ambiguity etc) into syntax errors
//  p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))

//  // Custom error listener, register before the parse
//  el := &AdlAstErrorListener{}
//  p.AddErrorListener(el)

//  // Parse - start rule
//  tree := p.Start()

//  // Antlr provided parse tree representation
//  sexpr := antlr.TreesStringTree(tree, nil, p)
//  fmt.Printf("%s\n", sexpr)

//  // Custom listener
//  l := &AdlAstListener{}
//  antlr.ParseTreeWalkerDefault.Walk(l, tree)

//  // Custom visitor
//  v := &AdlAstVisitor{}
//  tree.Accept(v)
// }

//// implemented all listeners methods
//var _ adlast.AdlAstListener = &AdlAstListener{}
//// implemented specific
//var _ adlast.AdlEntryListener = &AdlAstListener{}
//var _ adlast.AdlExitListener = &AdlAstListener{}

//var _ adlast.ModuleEntryListener = &AdlAstListener{}
//var _ adlast.ModuleExitListener = &AdlAstListener{}

//var _ adlast.Import_EntryListener = &AdlAstListener{}
//var _ adlast.Import_ExitListener = &AdlAstListener{}

//var _ adlast.DeclEntryListener = &AdlAstListener{}
//var _ adlast.DeclExitListener = &AdlAstListener{}

//var _ adlast.ScopedNameEntryListener = &AdlAstListener{}
//var _ adlast.ScopedNameExitListener = &AdlAstListener{}

//var _ adlast.DeclTypeEntryListener = &AdlAstListener{}
//var _ adlast.DeclTypeExitListener = &AdlAstListener{}

//var _ adlast.Struct_EntryListener = &AdlAstListener{}
//var _ adlast.Struct_ExitListener = &AdlAstListener{}

//var _ adlast.Union_EntryListener = &AdlAstListener{}
//var _ adlast.Union_ExitListener = &AdlAstListener{}

//var _ adlast.TypeDefEntryListener = &AdlAstListener{}
//var _ adlast.TypeDefExitListener = &AdlAstListener{}

//var _ adlast.NewTypeEntryListener = &AdlAstListener{}
//var _ adlast.NewTypeExitListener = &AdlAstListener{}

//var _ adlast.FieldEntryListener = &AdlAstListener{}
//var _ adlast.FieldExitListener = &AdlAstListener{}

//var _ adlast.TypeExprEntryListener = &AdlAstListener{}
//var _ adlast.TypeExprExitListener = &AdlAstListener{}

//var _ adlast.TypeRefEntryListener = &AdlAstListener{}
//var _ adlast.TypeRefExitListener = &AdlAstListener{}

//var _ adlast.JsonStrEntryListener = &AdlAstListener{}
//var _ adlast.JsonStrExitListener = &AdlAstListener{}

//var _ adlast.JsonBoolEntryListener = &AdlAstListener{}
//var _ adlast.JsonBoolExitListener = &AdlAstListener{}

//var _ adlast.JsonNullEntryListener = &AdlAstListener{}
//var _ adlast.JsonNullExitListener = &AdlAstListener{}

//var _ adlast.JsonIntEntryListener = &AdlAstListener{}
//var _ adlast.JsonIntExitListener = &AdlAstListener{}

//var _ adlast.JsonFloatEntryListener = &AdlAstListener{}
//var _ adlast.JsonFloatExitListener = &AdlAstListener{}

//var _ adlast.JsonArrayEntryListener = &AdlAstListener{}
//var _ adlast.JsonArrayExitListener = &AdlAstListener{}

//var _ adlast.JsonObjEntryListener = &AdlAstListener{}
//var _ adlast.JsonObjExitListener = &AdlAstListener{}

//var _ adlast.JsonErrorEntryListener = &AdlAstListener{}
//var _ adlast.JsonErrorExitListener = &AdlAstListener{}

//type AdlAstListener struct {
//}

//type AdlAstErrorListener struct {
//    Warning string
//    Err     error
//    Debug   bool
//}

// func (cb *AdlAstErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
//  if cb.Debug {
//      fmt.Printf("SyntaxError %d:%d <%s>\n", line, column, msg)
//  }
//  if strings.HasPrefix(msg, "report") { // TODO remove NewDiagnosticErrorListener and move warning to ReportAmbiguity etc. when getDecisionDescription is make public
//      cb.Warning = fmt.Sprintf("At %d:%d <%s>", line, column, msg)
//  } else {
//      cb.Err = fmt.Errorf("SyntaxError %d:%d <%s>", line, column, msg)
//  }
// }
// func (cb *AdlAstErrorListener) ReportAmbiguity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int,
//  exact bool, ambigAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
//  if cb.Debug {
//      fmt.Printf("ReportAmbiguity rec:%v dfs:%v start:%d stop:%d, exact:%v, ambigAlts:%v config:%v\n", recognizer, dfa, startIndex, stopIndex, exact, ambigAlts, configs)
//  }
// }
// func (cb *AdlAstErrorListener) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
//  if cb.Debug {
//      fmt.Printf("ReportAttemptingFullContext rec:%v dfs:%v start:%d stop:%d, conflictingAlts:%v config:%v\n", recognizer, dfa, startIndex, stopIndex, conflictingAlts, configs)
//  }
// }
// func (cb *AdlAstErrorListener) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex, prediction int, configs antlr.ATNConfigSet) {
//  if cb.Debug {
//      fmt.Printf("ReportContextSensitivity rec:%v dfs:%v start:%d stop:%d, prediction:%v configs:%v\n", recognizer, dfa, startIndex, stopIndex, prediction, configs)
//  }
// }

//// antlr.ParseTreeListener implementation.
//// All required.

//func (s *AdlAstListener ) VisitTerminal(node  antlr.TerminalNode) {   }
//func (s *AdlAstListener ) VisitErrorNode(node antlr.ErrorNode)    {   }
//func (s *AdlAstListener ) EnterEveryRule(ctx antlr.ParserRuleContext) {  }
//func (s *AdlAstListener ) ExitEveryRule(ctx antlr.ParserRuleContext) {  }

//// Only implemented as needed.

//func (s *AdlAstListener) EnterAdl(ctx adlast.*adlast.AdlContext) {}
//func (s *AdlAstListener) ExitAdl(ctx adlast.*adlast.AdlContext) {}

//func (s *AdlAstListener) EnterModule(ctx adlast.*adlast.ModuleContext) {}
//func (s *AdlAstListener) ExitModule(ctx adlast.*adlast.ModuleContext) {}

//func (s *AdlAstListener) EnterImport_(ctx adlast.*adlast.Import_Context) {}
//func (s *AdlAstListener) ExitImport_(ctx adlast.*adlast.Import_Context) {}

//func (s *AdlAstListener) EnterDecl(ctx adlast.*adlast.DeclContext) {}
//func (s *AdlAstListener) ExitDecl(ctx adlast.*adlast.DeclContext) {}

//func (s *AdlAstListener) EnterScopedName(ctx adlast.*adlast.ScopedNameContext) {}
//func (s *AdlAstListener) ExitScopedName(ctx adlast.*adlast.ScopedNameContext) {}

//func (s *AdlAstListener) EnterDeclType(ctx adlast.*adlast.DeclTypeContext) {}
//func (s *AdlAstListener) ExitDeclType(ctx adlast.*adlast.DeclTypeContext) {}

//func (s *AdlAstListener) EnterStruct_(ctx adlast.*adlast.Struct_Context) {}
//func (s *AdlAstListener) ExitStruct_(ctx adlast.*adlast.Struct_Context) {}

//func (s *AdlAstListener) EnterUnion_(ctx adlast.*adlast.Union_Context) {}
//func (s *AdlAstListener) ExitUnion_(ctx adlast.*adlast.Union_Context) {}

//func (s *AdlAstListener) EnterTypeDef(ctx adlast.*adlast.TypeDefContext) {}
//func (s *AdlAstListener) ExitTypeDef(ctx adlast.*adlast.TypeDefContext) {}

//func (s *AdlAstListener) EnterNewType(ctx adlast.*adlast.NewTypeContext) {}
//func (s *AdlAstListener) ExitNewType(ctx adlast.*adlast.NewTypeContext) {}

//func (s *AdlAstListener) EnterField(ctx adlast.*adlast.FieldContext) {}
//func (s *AdlAstListener) ExitField(ctx adlast.*adlast.FieldContext) {}

//func (s *AdlAstListener) EnterTypeExpr(ctx adlast.*adlast.TypeExprContext) {}
//func (s *AdlAstListener) ExitTypeExpr(ctx adlast.*adlast.TypeExprContext) {}

//func (s *AdlAstListener) EnterTypeRef(ctx adlast.*adlast.TypeRefContext) {}
//func (s *AdlAstListener) ExitTypeRef(ctx adlast.*adlast.TypeRefContext) {}

//func (s *AdlAstListener) EnterJsonStr(ctx adlast.*adlast.JsonStrContext) {}
//func (s *AdlAstListener) ExitJsonStr(ctx adlast.*adlast.JsonStrContext) {}

//func (s *AdlAstListener) EnterJsonBool(ctx adlast.*adlast.JsonBoolContext) {}
//func (s *AdlAstListener) ExitJsonBool(ctx adlast.*adlast.JsonBoolContext) {}

//func (s *AdlAstListener) EnterJsonNull(ctx adlast.*adlast.JsonNullContext) {}
//func (s *AdlAstListener) ExitJsonNull(ctx adlast.*adlast.JsonNullContext) {}

//func (s *AdlAstListener) EnterJsonInt(ctx adlast.*adlast.JsonIntContext) {}
//func (s *AdlAstListener) ExitJsonInt(ctx adlast.*adlast.JsonIntContext) {}

//func (s *AdlAstListener) EnterJsonFloat(ctx adlast.*adlast.JsonFloatContext) {}
//func (s *AdlAstListener) ExitJsonFloat(ctx adlast.*adlast.JsonFloatContext) {}

//func (s *AdlAstListener) EnterJsonArray(ctx adlast.*adlast.JsonArrayContext) {}
//func (s *AdlAstListener) ExitJsonArray(ctx adlast.*adlast.JsonArrayContext) {}

//func (s *AdlAstListener) EnterJsonObj(ctx adlast.*adlast.JsonObjContext) {}
//func (s *AdlAstListener) ExitJsonObj(ctx adlast.*adlast.JsonObjContext) {}

//func (s *AdlAstListener) EnterJsonError(ctx adlast.*adlast.JsonErrorContext) {}
//func (s *AdlAstListener) ExitJsonError(ctx adlast.*adlast.JsonErrorContext) {}
