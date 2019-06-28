package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/jpillora/opts"
	"github.com/wxio/tron-go/adl"
)

func NewLoadAdlAst() opts.Opts {
	return opts.New(&readast{}).Name("load_ast")
}

type readast struct {
	File string `type:"arg" help:"adl ast file" predict:"files"`
}

func (et *readast) Run() error {
	by, err := ioutil.ReadFile(et.File)
	if err != nil {
		return err
	}
	m := make(map[string]adl.Module)
	err = json.Unmarshal(by, &m)
	if err != nil {
		return err
	}
	// fmt.Printf("%v\n", m)
	b2, err := json.MarshalIndent(m, "", "    ")
	if err != nil {
		return err
	}
	fmt.Printf("%v\n", string(b2))
	return nil
}
