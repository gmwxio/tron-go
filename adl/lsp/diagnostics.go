package lsp

import (
	"context"
	"fmt"
	"runtime"

	"github.com/golangq/q"
	antlr "github.com/wxio/goantlr"
	"github.com/wxio/tron-go/adl"
	"golang.org/x/tools/lsp/protocol"
)

func qstack() {
	pc := make([]uintptr, 10)
	n := runtime.Callers(0, pc)
	if n == 0 {
		// No pcs available. Stop now.
		// This can happen if the first argument to runtime.Callers is large.
		return
	}
	pc = pc[:n] // pass only valid pcs to runtime.CallersFrames
	frames := runtime.CallersFrames(pc)
	// Loop to get frames.
	// A fixed number of pcs can expand to an indefinite number of Frames.
	for {
		frame, more := frames.Next()
		// To keep this example's output stable
		// even if there are changes in the testing package,
		// stop unwinding when we leave package runtime.
		// if !strings.Contains(frame.File, "runtime/") {
		// 	break
		// }
		q.Q(frame)
		if !more {
			break
		}
	}
}

type errColl struct {
	errs []antlr.ErrorNode
}

func (s *errColl) VisitTerminal(node antlr.TerminalNode)      {}
func (s *errColl) EnterEveryRule(ctx antlr.ParserRuleContext) {}
func (s *errColl) ExitEveryRule(ctx antlr.ParserRuleContext)  {}
func (s *errColl) VisitErrorNode(node antlr.ErrorNode) {
	s.errs = append(s.errs, node)
}

func (svr *server) diag(ctx context.Context, fname string, text string) {
	defer func() {
		if r := recover(); r != nil {
			q.Q(r)
			qstack()
		}
	}()
	// if strings.HasPrefix(fname, "file://") {
	// 	fname = fname[len("file://"):]
	// }
	// by, err := ioutil.ReadFile(fname)
	// if err != nil {
	// 	q.Q(err)
	// 	return
	// }
	dss := []protocol.Diagnostic{}
	tr, atr, bl, ts, err1 := adl.BuildAdlAST(text)
	_, _, _, _ = tr, atr, bl, ts
	// adl.QTreeToken(ts, bl)
	if err1.Error() != nil {
		// q.Q("%v", tr.TreeString())
		errC := &errColl{}
		antlr.ParseTreeWalkerDefault.Walk(errC, atr)
		// q.Q("Lex Errors")
		for i, er := range err1.LexErr {
			_ = i
			// q.Q(i, er)
			ds := protocol.Diagnostic{
				Range: protocol.Range{
					Start: protocol.Position{
						Line:      float64(er.Line() - 1),
						Character: float64(er.Column()),
					},
					End: protocol.Position{
						Line:      float64(er.Line() - 1),
						Character: float64(er.Column() + er.Len()),
					},
				},
				Severity:           protocol.SeverityError,
				Code:               er.Text(),
				Source:             "ADL-LEX",
				Message:            er.Message(),
				Tags:               []protocol.DiagnosticTag{},
				RelatedInformation: []protocol.DiagnosticRelatedInformation{},
			}
			dss = append(dss, ds)
		}
		// q.Q("Parse Errors")
		for i, er := range err1.ParseErr {
			_ = i
			// q.Q(i, er)
			ds := protocol.Diagnostic{
				Range: protocol.Range{
					Start: protocol.Position{
						Line:      float64(er.Line()),
						Character: float64(er.Column()),
					},
					End: protocol.Position{
						Line:      float64(er.Line()),
						Character: float64(er.Column() + er.Len()),
					},
				},
				Severity:           protocol.SeverityError,
				Code:               er.Text(),
				Source:             "ADL-PARSER",
				Message:            er.Msg,
				Tags:               []protocol.DiagnosticTag{},
				RelatedInformation: []protocol.DiagnosticRelatedInformation{},
			}
			dss = append(dss, ds)
		}
		// q.Q("Syntax Errors")
		for i, er := range err1.SyntaxErr {
			_ = i
			// q.Q(i, er)
			ds := protocol.Diagnostic{
				Range: protocol.Range{
					Start: protocol.Position{
						Line:      float64(er.Line()),
						Character: float64(er.Column()),
					},
					End: protocol.Position{
						Line:      float64(er.Line()),
						Character: float64(er.Column() + er.Len()),
					},
				},
				Severity:           protocol.SeverityError,
				Code:               er.Text(),
				Source:             "ADL-SYNTAX",
				Message:            er.Message(),
				Tags:               []protocol.DiagnosticTag{},
				RelatedInformation: []protocol.DiagnosticRelatedInformation{},
			}
			dss = append(dss, ds)
		}
		// q.Q("Error Nodes")
		for i, er := range errC.errs {
			// q.Q(i, er.GetSymbol())
			ds := protocol.Diagnostic{
				Range: protocol.Range{
					Start: protocol.Position{
						Line:      float64(er.GetSymbol().GetLine() - 1),
						Character: float64(er.GetSymbol().GetColumn()),
					},
					End: protocol.Position{
						Line:      float64(er.GetSymbol().GetLine() - 1),
						Character: float64(er.GetSymbol().GetColumn() + len(er.GetText())),
					},
				},
				Severity:           protocol.SeverityError,
				Code:               er.GetText(),
				Source:             "ADL-WALK",
				Message:            fmt.Sprintf("error at token %v", er),
				Tags:               []protocol.DiagnosticTag{},
				RelatedInformation: []protocol.DiagnosticRelatedInformation{},
			}
			dss = append(dss, ds)
			if i > 9 {
				// q.Q("  ... total errs ", len(errC.errs))
				break
			}
		}
		// if tr != nil {
		// 	errC = &errColl{}
		// 	err3 := adl.WalkADLWo(tr, errC)
		// 	q.Q("Parse Errors")
		// 	for i, er := range err3.ParseErr {
		// 		q.Q(i, er)
		// 		ds := protocol.Diagnostic{
		// 			Range: protocol.Range{
		// 				Start: protocol.Position{
		// 					Line:      float64(er.Line()),
		// 					Character: float64(er.Column()),
		// 				},
		// 				End: protocol.Position{
		// 					Line:      float64(er.Line()),
		// 					Character: float64(er.Column() + er.Len()),
		// 				},
		// 			},
		// 			Severity:           protocol.SeverityError,
		// 			Code:               er.Text(),
		// 			Source:             "ADL-TREE-PARSER",
		// 			Message:            er.Msg,
		// 			Tags:               []protocol.DiagnosticTag{},
		// 			RelatedInformation: []protocol.DiagnosticRelatedInformation{},
		// 		}
		// 		dss = append(dss, ds)
		// 	}
		// 	q.Q("Syntax Errors")
		// 	for i, er := range err3.SyntaxErr {
		// 		q.Q(i, er)
		// 		ds := protocol.Diagnostic{
		// 			Range: protocol.Range{
		// 				Start: protocol.Position{
		// 					Line:      float64(er.Line()),
		// 					Character: float64(er.Column()),
		// 				},
		// 				End: protocol.Position{
		// 					Line:      float64(er.Line()),
		// 					Character: float64(er.Column() + er.Len()),
		// 				},
		// 			},
		// 			Severity:           protocol.SeverityError,
		// 			Code:               er.Text(),
		// 			Source:             "ADL-TREE-SYNTAX",
		// 			Message:            er.Message(),
		// 			Tags:               []protocol.DiagnosticTag{},
		// 			RelatedInformation: []protocol.DiagnosticRelatedInformation{},
		// 		}
		// 		dss = append(dss, ds)
		// 	}
		// 	q.Q("Error Nodes")
		// 	for i, er := range errC.errs {
		// 		q.Q(i, er.GetSymbol())
		// 		ds := protocol.Diagnostic{
		// 			Range: protocol.Range{
		// 				Start: protocol.Position{
		// 					Line:      float64(er.GetSymbol().GetLine() - 1),
		// 					Character: float64(er.GetSymbol().GetColumn()),
		// 				},
		// 				End: protocol.Position{
		// 					Line:      float64(er.GetSymbol().GetLine() - 1),
		// 					Character: float64(er.GetSymbol().GetColumn() + len(er.GetText())),
		// 				},
		// 			},
		// 			Severity:           protocol.SeverityError,
		// 			Code:               er.GetText(),
		// 			Source:             "ADL-TREE-WALK",
		// 			Message:            fmt.Sprintf("error at token %v", er),
		// 			Tags:               []protocol.DiagnosticTag{},
		// 			RelatedInformation: []protocol.DiagnosticRelatedInformation{},
		// 		}
		// 		dss = append(dss, ds)
		// 		if i > 9 {
		// 			q.Q("  ... total errs ", len(errC.errs))
		// 			break
		// 		}
		// 	}
		// }
	}
	// q.Q(dss)
	svr.client.PublishDiagnostics(ctx, &protocol.PublishDiagnosticsParams{
		Diagnostics: dss,
		URI:         fname,
	})
}
