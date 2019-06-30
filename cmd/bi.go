package cmd

import (
	"fmt"
	"io/ioutil"

	"github.com/jpillora/opts"
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
	// fmt.Printf("%v\n", bl)
	// fmt.Printf("%v\n", ts)
	return err
}
