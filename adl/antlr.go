package adl

import (
	"fmt"
	"strings"

	"github.com/golangq/q"
	antlr "github.com/wxio/goantlr"
	walker "github.com/wxio/tron-go/internal/adlwi"
	"github.com/wxio/tron-go/internal/ctree"
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

func QTreeToken(tts antlr.TokenStream, p antlr.Recognizer) {
	i := 1
	for {
		to := tts.Get(i)
		if -1 == to.GetTokenType() {
			break
		}
		q.Q(fmt.Sprintf("tt:%v %d : %v\t\t%d:%d", to.GetTokenType(), i,
			p.GetSymbolicNames()[to.GetTokenType()],
			to.GetLine(), to.GetColumn(),
		))
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

type DiagMessage interface {
	Line() int
	Column() int
	Message() string
	Len() int
	Text() string
}

type Error struct {
	ctree.TreeNode `json:"-"`
	Start, Stop    antlr.Token
	Expected       []string
	Received       string
	Annotations    `json:"annotations"`
}

func (er *Error) SetTreeNode(t ctree.TreeNode) { er.TreeNode = t }
func (er Error) Line() int {
	return er.Start.GetLine() - 1
}
func (er Error) Column() int {
	return er.Start.GetColumn()
}
func (er Error) Message() string {
	return fmt.Sprintf("expected %v received %v", er.Expected, er.Received)
}
func (er Error) Len() int {
	return len(er.Start.GetText())
}
func (er Error) Text() string {
	return er.Start.GetText()
}

type lexErrMsg struct {
	OffendingSymbol interface{}
	OffendingToken  antlr.Token
	line, column    int
	Msg             string
}

func (er *lexErrMsg) Line() int {
	return er.line
}
func (er *lexErrMsg) Column() int {
	return er.column
}
func (er *lexErrMsg) Message() string {
	return er.Msg
}
func (lem *lexErrMsg) Len() int {
	if lem.OffendingToken != nil {
		return len(lem.OffendingToken.GetText())
	}
	if t, ok := lem.OffendingSymbol.(antlr.Token); ok {
		return len(t.GetText())
	}
	return 1
}
func (lem *lexErrMsg) Text() string {
	if lem.OffendingToken != nil {
		return lem.OffendingToken.GetText()
	}
	if t, ok := lem.OffendingSymbol.(antlr.Token); ok {
		return t.GetText()
	}
	return "<unknown token>"
}

type lexErr struct {
	err     []DiagMessage
	warning []interface{}
}

func (d *lexErr) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	var t antlr.Token
	if e != nil {
		t = e.GetOffendingToken()
	}
	d.err = append(d.err, &lexErrMsg{
		OffendingSymbol: offendingSymbol,
		OffendingToken:  t,
		line:            line,
		column:          column,
		Msg:             msg,
	})
}

func (d *lexErr) ReportAmbiguity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, exact bool, ambigAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
	sta := recognizer.GetTokenStream().Get(startIndex)
	sto := recognizer.GetTokenStream().Get(stopIndex)
	d.warning = append(d.warning, fmt.Errorf("%v:%v ReportAmbiguity rec:%v dfs:%v start:%d stop:%d, exact:%v, ambigAlts:%v config:%v\n", sta, sto, recognizer, dfa, startIndex, stopIndex, exact, ambigAlts, configs))
}

func (d *lexErr) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
	d.warning = append(d.warning, fmt.Errorf("ReportAttemptingFullContext rec:%v dfs:%v start:%d stop:%d, conflictingAlts:%v config:%v\n", recognizer, dfa, startIndex, stopIndex, conflictingAlts, configs))
}

func (d *lexErr) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex, prediction int, configs antlr.ATNConfigSet) {
	d.warning = append(d.warning, fmt.Errorf("ReportContextSensitivity rec:%v dfs:%v start:%d stop:%d, config:%v\n", recognizer, dfa, startIndex, stopIndex, configs))
}

type parseErrMsg struct {
	offendingSymbol interface{}
	offendingToken  antlr.Token
	line, column    int
	Msg             string
}

func (er parseErrMsg) Line() int {
	return er.line - 1
}
func (er parseErrMsg) Column() int {
	return er.column
}
func (er parseErrMsg) Message() string {
	return er.Msg
}
func (er parseErrMsg) Len() int {
	if er.offendingToken != nil {
		return len(er.offendingToken.GetText())
	}
	if t, ok := er.offendingSymbol.(antlr.Token); ok {
		return len(t.GetText())
	}
	return 1
}
func (er parseErrMsg) Text() string {
	if er.offendingToken != nil {
		return er.offendingToken.GetText()
	}
	if t, ok := er.offendingSymbol.(antlr.Token); ok {
		return t.GetText()
	}
	return "<nil>"
}

type parseErr struct {
	ParseErr []DiagMessage
	// SyntaxErr []interface{}
	SyntaxWarning []interface{}
}

func (v *parseErr) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	if strings.HasPrefix(msg, "reportAttemptingFullContext") { // TODO remove NewDiagnosticErrorListener and move warning to ReportAmbiguity etc. when getDecisionDescription is make public
		v.SyntaxWarning = append(v.SyntaxWarning, fmt.Sprintf("At %d:%d <%s>", line, column, msg))
		return
	}
	var t antlr.Token
	if e != nil {
		t = e.GetOffendingToken()
	}
	// t, ok := offendingSymbol.(antlr.Token)
	// if !ok && e != nil {
	// 	t = e.GetOffendingToken()
	// }
	v.ParseErr = append(v.ParseErr, parseErrMsg{
		offendingSymbol: offendingSymbol,
		offendingToken:  t,
		line:            line,
		column:          column,
		Msg:             msg,
	})
}
func (v *parseErr) ReportAmbiguity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, exact bool, ambigAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
	sta := recognizer.GetTokenStream().Get(startIndex)
	sto := recognizer.GetTokenStream().Get(stopIndex)
	v.SyntaxWarning = append(v.SyntaxWarning, fmt.Sprintf("%v:%v ReportAmbiguity rec:%v dfs:%v start:%d stop:%d, exact:%v, ambigAlts:%v config:%v\n", sta, sto, recognizer, dfa, startIndex, stopIndex, exact, ambigAlts, configs))
}
func (v *parseErr) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
	v.SyntaxWarning = append(v.SyntaxWarning, fmt.Sprintf("ReportAttemptingFullContext rec:%v dfs:%v start:%d stop:%d, conflictingAlts:%v config:%v\n", recognizer, dfa, startIndex, stopIndex, conflictingAlts, configs))
}
func (v *parseErr) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex, prediction int, configs antlr.ATNConfigSet) {
	v.SyntaxWarning = append(v.SyntaxWarning, fmt.Sprintf("ReportContextSensitivity rec:%v dfs:%v start:%d stop:%d, config:%v\n", recognizer, dfa, startIndex, stopIndex, configs))
}
