package adl

import (
	antlr "github.com/wxio/goantlr"
	walker "github.com/wxio/tron-go/internal/adlwi"
	"github.com/wxio/tron-go/internal/ctree"
)

func WalkADL(tr ctree.Tree, list antlr.ParseTreeListener) error {
	var tttype *TTType
	var tts antlr.TokenStream = ctree.NewTreeTokenSource(tr, tttype)
	p := walker.NewAdlWi(tts)
	// debugTreeToken(tts, p)
	p.SetTokenStream(tts)
	p.RemoveErrorListeners()
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	el := &antlrEL{}
	p.AddErrorListener(el)
	p.BuildParseTrees = true
	jv := p.Adl()
	antlr.ParseTreeWalkerDefault.Walk(list, jv)
	return el.err
}

func VisitADL(tr ctree.Tree, vi antlr.ParseTreeVisitor) error {
	var tttype *TTType
	var tts antlr.TokenStream = ctree.NewTreeTokenSource(tr, tttype)
	p := walker.NewAdlWi(tts)
	debugTreeToken(tts, p)
	p.SetTokenStream(tts)
	p.RemoveErrorListeners()
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	el := &antlrEL{}
	p.AddErrorListener(el)
	ctx := p.Adl()
	ctx.Visit(vi)
	return el.err
}
