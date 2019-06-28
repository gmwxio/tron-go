package main

import (
	"github.com/jpillora/opts"
	"github.com/wxio/tron-go/cmd"
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
	opts.New(&r).Name("tron-go").
		EmbedGlobalFlagSet().
		Complete().
		Version(Version).
		AddCommand(opts.New(&adl{}).
			AddCommand(cmd.NewLoadAdlAst())).
		Parse().
		RunFatal()
}

func (r root) Run() {
}
