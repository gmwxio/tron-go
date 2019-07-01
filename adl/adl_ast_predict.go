package adl

import (
	antlr "github.com/wxio/goantlr"
	walker "github.com/wxio/tron-go/internal/adlwo"
	"github.com/wxio/tron-go/internal/ctree"
)

func WalkAdlWo(tr ctree.Tree, list antlr.ParseTreeListener) errs {
	var tttype *TTType
	var tts antlr.TokenStream = ctree.NewTreeTokenSource(tr, tttype)
	p := walker.NewAdlWo(tts)
	// debugTreeToken(tts, p)
	p.SetTokenStream(tts)
	p.RemoveErrorListeners()
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	el := &parseErr{}
	p.AddErrorListener(el)
	p.BuildParseTrees = true
	jv := p.Adl()
	antlr.ParseTreeWalkerDefault.Walk(list, jv)
	errs := errs{
		ParseErr:      el.ParseErr,
		SyntaxWarning: el.SyntaxWarning,
	}
	return errs
}