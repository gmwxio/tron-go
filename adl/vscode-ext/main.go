package main

import (
	"github.com/jpillora/opts"
	"github.com/wxio/tron-go/adl/lsp"
)

var (
	Version string
	Data    string
	Commit  string
)

type root struct{}
type adl struct{}

func main() {
	r := root{}
	opts.New(&r).Name("adl-lsp").
		EmbedGlobalFlagSet().
		Version(Version).
		AddCommand(lsp.NewTcp()).
		AddCommand(lsp.NewConsole()).
		Parse().
		RunFatal()
}
