package main

import (
	"github.com/jpillora/opts"
	"github.com/wxio/tron-go/adl/lsp"
	"github.com/wxio/tron-go/cmd"
)

var (
	Version string
	Date    string
	Commit  string
)

type root struct{}
type adl struct{}

func main() {
	r := root{}
	opts.New(&r).Name("tron-go").
		EmbedGlobalFlagSet().
		Complete().
		Version(Version).
		AddCommand(opts.New(&adl{}).
			AddCommand(cmd.NewLoadAdlAst()).
			AddCommand(cmd.BuildAdlAst()).
			AddCommand(lsp.NewTcp(Version)).
			AddCommand(lsp.NewConsole(Version))).
		Parse().
		RunFatal()
}
