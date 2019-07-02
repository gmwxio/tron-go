package main

import (
	"errors"
	"fmt"

	"github.com/jpillora/opts"
	"github.com/wxio/tron-go/adl/lsp"
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
	opts.New(&r).Name("adl-lsp").
		EmbedGlobalFlagSet().
		Version(Version).
		AddCommand(lsp.NewTcp()).
		AddCommand(lsp.NewConsole()).
		Parse().
		RunFatal()
}

func (*root) Run() error {
	fmt.Printf("version: %v\ncommit: %v\n date: %v\n", Version, Commit, Date)
	return errors.New("")
}
