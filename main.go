package main

import (
	"github.com/jpillora/opts"
	"github.com/wxio/tron-go/cmd"
	"github.com/wxio/tron-go/tools"
)

var (
	Version string
	Date    string
	Commit  string
)

type root struct{}
type build struct{}
type adl struct{}

func main() {
	r := root{}
	opts.New(&r).Name("tron-go").
		EmbedGlobalFlagSet().
		Complete().
		Version(Version).
		AddCommand(opts.New(&build{}).
			AddCommand(tools.NewAntlr()).
			AddCommand(tools.NewAntlrs().
				ConfigPath(".antlr.build.json"))).
		AddCommand(opts.New(&adl{}).
			AddCommand(cmd.NewLoadAdlAst()).
			AddCommand(cmd.BuildAdlAst())).
		Parse().
		RunFatal()
}
