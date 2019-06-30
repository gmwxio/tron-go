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
	tr, bl, ts, err := adl.BuildAdlAST(string(by))
	_, _, _ = tr, bl, ts
	fmt.Printf("%v\n", tr)
	adl.WalkADL(tr, &AdlWiListener{})
	// fmt.Printf("%v\n", bl)
	// fmt.Printf("%v\n", ts)
	return err
}

type AdlWiListener struct {
	intend   string
	tokCount int
}

func (s *AdlWiListener) VisitTerminal(node antlr.TerminalNode) {
	s.tokCount++
	fmt.Printf("%d%s  >>%v\n", s.tokCount, s.intend, node)
}
func (s *AdlWiListener) VisitErrorNode(node antlr.ErrorNode) {
	s.tokCount++
	fmt.Printf("%d Error %v\n", s.tokCount, node)
}
func (s *AdlWiListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	s.tokCount++
	fmt.Printf("%d%s>>%T\n", s.tokCount, s.intend, ctx)
	s.intend += "\t"
}
func (s *AdlWiListener) ExitEveryRule(ctx antlr.ParserRuleContext) {
	s.intend = s.intend[1:]
}
