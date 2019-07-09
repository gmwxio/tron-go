package main

import (
	"errors"
	"fmt"

	"github.com/jpillora/opts"
	"github.com/wxio/tron-go/adl/vscode-ext/lsp"
)

var (
	Version = "dev"
	Date    string
	Commit  string
)

type root struct{}

func main() {
	r := root{}
	opts.New(&r).Name("adl-lsp").
		EmbedGlobalFlagSet().
		Version(Version).
		AddCommand(lsp.NewTcp(Version)).
		AddCommand(lsp.NewConsole(Version)).
		Parse().
		RunFatal()
}

func (*root) Run() error {
	fmt.Printf("version: %v\ncommit: %v\ndate: %v\n", Version, Commit, Date)
	return errors.New("")
}
