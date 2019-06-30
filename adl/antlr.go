package adl

import (
	"fmt"
	"strconv"

	antlr "github.com/wxio/goantlr"
	walker "github.com/wxio/tron-go/internal/adlwi"
)

type TTType struct{}

func (*TTType) Eof() int  { return walker.AdlWiEOF }
func (*TTType) Down() int { return walker.AdlWiDOWN }
func (*TTType) Up() int   { return walker.AdlWiUP }

func debugTreeToken(tts antlr.TokenStream, p antlr.Recognizer) {
	i := 1
	for {
		to := tts.Get(i)
		if -1 == to.GetTokenType() {
			break
		}
		fmt.Printf("%d ", to.GetTokenType())
		fmt.Printf("%d : %v\t\t%v\n", i,
			p.GetSymbolicNames()[to.GetTokenType()],
			to.GetLine(),
		)
		i++
	}
}

func tokens2strings(arr []antlr.Token) []string {
	name := make([]string, len(arr))
	for i, tks := range arr {
		name[i] = tks.GetText()
	}
	return name
}

type Error struct {
	Expected    []string
	Received    string
	Annotations `json:"annotations"`
}

type antlrEL struct {
	err     error
	warning error
}

func (d *antlrEL) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	d.err = fmt.Errorf("line " + strconv.Itoa(line) + ":" + strconv.Itoa(column) + " " + msg)
}

func (d *antlrEL) ReportAmbiguity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, exact bool, ambigAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
	sta := recognizer.GetTokenStream().Get(startIndex)
	sto := recognizer.GetTokenStream().Get(stopIndex)
	d.warning = fmt.Errorf("%v:%v ReportAmbiguity rec:%v dfs:%v start:%d stop:%d, exact:%v, ambigAlts:%v config:%v\n", sta, sto, recognizer, dfa, startIndex, stopIndex, exact, ambigAlts, configs)
}

func (d *antlrEL) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
	d.warning = fmt.Errorf("ReportAttemptingFullContext rec:%v dfs:%v start:%d stop:%d, conflictingAlts:%v config:%v\n", recognizer, dfa, startIndex, stopIndex, conflictingAlts, configs)
}

func (d *antlrEL) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex, prediction int, configs antlr.ATNConfigSet) {
	d.warning = fmt.Errorf("ReportContextSensitivity rec:%v dfs:%v start:%d stop:%d, config:%v\n", recognizer, dfa, startIndex, stopIndex, configs)
}
