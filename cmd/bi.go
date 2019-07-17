package cmd

import (
	"fmt"
	"io/ioutil"

	"github.com/jpillora/opts"
	antlr "github.com/wxio/goantlr"
	"github.com/wxio/tron-go/adl"
)

func BuildAdlAst() opts.Opts {
	return opts.New(&buildAdlAst{}).Name("build_ast")
}

type buildAdlAst struct {
	File string `type:"arg" help:"adl file" predict:"files"`
}

func (cm *buildAdlAst) Run() error {
	by, err := ioutil.ReadFile(cm.File)
	if err != nil {
		return err
	}
	tr, atr, bl, ts, err1 := adl.BuildAdlAST(string(by))
	_, _, _, _ = tr, atr, bl, ts
	if err1.Error() != nil {
		// fmt.Printf("%v\n", tr.TreeString())
		errColl := &errColl{}
		antlr.ParseTreeWalkerDefault.Walk(errColl, atr)
		fmt.Printf("Lex Errors\n")
		for i, er := range err1.LexErr {
			fmt.Printf("  %d %v\n", i, er)
		}
		fmt.Printf("Parse Errors\n")
		for i, er := range err1.ParseErr {
			fmt.Printf("  %d %v\n", i, er)
		}
		fmt.Printf("Syntax Errors\n")
		for i, er := range err1.SyntaxErr {
			fmt.Printf("  %d %v\n", i, er)
		}
		fmt.Printf("Error Nodes\n")
		for i, er := range errColl.errs {
			fmt.Printf("  %d %v\n", i, er.GetSymbol())
			if i > 9 {
				fmt.Printf("  ... total errs %d\n", len(errColl.errs))
				break
			}
		}
		return fmt.Errorf("build err '%v' errorNodes'%v' ", err1.Error(), errColl.errs)
	}
	// fmt.Printf("%v\n", tr)
	// if err != nil {
	// 	return err
	// }
	// fmt.Printf("%v\n", tr.TreeString())
	_, err2 := adl.WalkADLWi(tr, &AdlWiListener{})
	// fmt.Printf("%v\n", bl)
	// fmt.Printf("%v\n", ts)
	if err2.Error() != nil {
		return fmt.Errorf("walk err '%v'", err2.Error())
	}
	return nil
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

type AdlWiListener struct {
	intend   string
	tokCount int
}

func (s *AdlWiListener) VisitTerminal(node antlr.TerminalNode) {
	s.tokCount++
	// fmt.Printf("%d%s  >>%v\n", s.tokCount, s.intend, node)
}
func (s *AdlWiListener) VisitErrorNode(node antlr.ErrorNode) {
	s.tokCount++
	// fmt.Printf("%d Error %v\n", s.tokCount, node)
}
func (s *AdlWiListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	s.tokCount++
	// fmt.Printf("%d%s>>%T\n", s.tokCount, s.intend, ctx)
	s.intend += "\t"
}
func (s *AdlWiListener) ExitEveryRule(ctx antlr.ParserRuleContext) {
	s.intend = s.intend[1:]
}
