package adl

import (
	antlr "github.com/wxio/goantlr"
	walker "github.com/wxio/tron-go/internal/adlwi"
	walkerWo "github.com/wxio/tron-go/internal/adlwo"
	"github.com/wxio/tron-go/internal/ctree"
)

func WalkADL(tr ctree.Tree, list antlr.ParseTreeListener) errs {
	var tttype *TTType
	var tts antlr.TokenStream = ctree.NewTreeTokenSource(tr, tttype)
	p := walker.NewAdlWi(tts)
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

func WalkADLWo(tr ctree.Tree, list antlr.ParseTreeListener) errs {
	var tttype *TTType
	var tts antlr.TokenStream = ctree.NewTreeTokenSource(tr, tttype)
	p := walkerWo.NewAdlWo(tts)
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

func VisitADL(tr ctree.Tree, vi antlr.ParseTreeVisitor) errs {
	var tttype *TTType
	var tts antlr.TokenStream = ctree.NewTreeTokenSource(tr, tttype)
	p := walker.NewAdlWi(tts)
	// debugTreeToken(tts, p)
	p.SetTokenStream(tts)
	p.RemoveErrorListeners()
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	el := &parseErr{}
	p.AddErrorListener(el)
	ctx := p.Adl()
	ctx.Visit(vi)
	errs := errs{
		ParseErr:      el.ParseErr,
		SyntaxWarning: el.SyntaxWarning,
	}
	return errs
}

func VisitAdlWo(tr ctree.Tree, vi antlr.ParseTreeVisitor) errs {
	var tttype *TTType
	var tts antlr.TokenStream = ctree.NewTreeTokenSource(tr, tttype)
	p := walkerWo.NewAdlWo(tts)
	// debugTreeToken(tts, p)
	p.SetTokenStream(tts)
	p.RemoveErrorListeners()
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	el := &parseErr{}
	p.AddErrorListener(el)
	ctx := p.Adl()
	ctx.Visit(vi)
	errs := errs{
		ParseErr:      el.ParseErr,
		SyntaxWarning: el.SyntaxWarning,
	}
	return errs
}
