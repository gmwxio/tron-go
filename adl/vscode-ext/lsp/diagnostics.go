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
	errs []adl.DiagMessage
}

type errorNode struct {
	antlr.ErrorNode
}

func (er errorNode) Line() int {
	return er.GetSymbol().GetLine() - 1
}
func (er errorNode) Column() int {
	return er.GetSymbol().GetColumn()
}
func (er errorNode) Message() string {
	return fmt.Sprintf("error at token %v", er.ErrorNode)
}
func (er errorNode) Len() int {
	return len(er.GetText())
}
func (er errorNode) Text() string {
	return er.GetText()
}

func (s *errColl) VisitTerminal(node antlr.TerminalNode)      {}
func (s *errColl) EnterEveryRule(ctx antlr.ParserRuleContext) {}
func (s *errColl) ExitEveryRule(ctx antlr.ParserRuleContext)  {}
func (s *errColl) VisitErrorNode(node antlr.ErrorNode) {

	s.errs = append(s.errs, errorNode{node})
}

// type diagColl struct {
// 	module  adl.Module
// 	imports []adl.Import
// }

// func (s *diagColl) VisitTerminal(node antlr.TerminalNode)      {}
// func (s *diagColl) EnterEveryRule(ctx antlr.ParserRuleContext) {}
// func (s *diagColl) ExitEveryRule(ctx antlr.ParserRuleContext)  {}
// func (s *diagColl) VisitErrorNode(node antlr.ErrorNode) {
// 	// s.errs = append(s.errs, node)
// }
// func (s *diagColl) EnterModule(ctx *adlwi.ModuleContext) {
// 	q.Q(ctx)
// 	s.module = ctx.GetTok().(*ctree.TreeNode).Val.(adl.Module)
// }
// func (s *diagColl) EnterImportModule(ctx *adlwi.ImportModuleContext) {
// 	q.Q(ctx)
// 	s.imports = append(s.imports, ctx.GetStart().(*ctree.TreeNode).Val.(adl.Import))
// }
// func (s *diagColl) EnterImportScopedModule(ctx *adlwi.ImportScopedModuleContext) {
// 	q.Q(ctx)
// 	s.imports = append(s.imports, ctx.GetStart().(*ctree.TreeNode).Val.(adl.Import))
// }

func (svr *server) diag(ctx context.Context, text string) []protocol.Diagnostic {
	defer func() {
		if r := recover(); r != nil {
			q.Q(r)
			qstack()
		}
	}()
	dss := []protocol.Diagnostic{}
	tr, atr, bl, ts, err1 := adl.BuildAdlAST(text)
	_, _, _, _ = tr, atr, bl, ts
	// adl.QTreeToken(ts, bl)
	if err1.Error() != nil {
		// q.Q("%v", tr.TreeString())
		errC := &errColl{}
		antlr.ParseTreeWalkerDefault.Walk(errC, atr)
		collectionDiag(err1.LexErr, "ADL-LEXER", -1, &dss)
		collectionDiag(err1.ParseErr, "ADL-PARSER", -1, &dss)
		collectionDiag(err1.SyntaxErr, "ADL-SYNTAX", -1, &dss)
		collectionDiag(errC.errs, "ADL-WALK", -9, &dss)
		if tr != nil {
			errC = &errColl{}
			_, err3 := adl.WalkADLWo(tr, errC)
			collectionDiag(err3.ParseErr, "ADL-TREE-PARSER", -1, &dss)
			collectionDiag(err3.SyntaxErr, "ADL-TREE-SYNTAX", -1, &dss)
			collectionDiag(errC.errs, "ADL-TREE-WALK", 9, &dss)
		}
	}
	//
	// dc := &diagColl{}
	// adl.WalkADL(tr, dc)
	// q.Q(dc.module)
	// q.Q(dc.imports)
	return dss
}

func collectionDiag(errs []adl.DiagMessage, src string, max int, dss *[]protocol.Diagnostic) {
	for i, er := range errs {
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
			Source:             src,
			Message:            er.Message(),
			Tags:               []protocol.DiagnosticTag{},
			RelatedInformation: []protocol.DiagnosticRelatedInformation{},
		}
		*dss = append(*dss, ds)
		if max != -1 && i > max {
			q.Q("  ... total errs ", len(errs))
			break
		}
	}
}
