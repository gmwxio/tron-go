package ctree

import (
	"fmt"

	antlr "github.com/wxio/goantlr"
)

type WalkableBuilder interface {
	Add(n antlr.Token) WalkableBuilder
	AddNode(n antlr.Token, ttype int, val interface{}) WalkableBuilder
	Down() WalkableBuilder
	Up() WalkableBuilder
	Build() Tree
	Current() antlr.Token
}
type walkerbuilder struct {
	tree Tree
	ttt  TreeTokenTypes
	curr antlr.Token
	last antlr.Token
}

type TreeNode struct {
	antlr.Token
	TType  int
	Val    interface{}
	TNText *string
}

type TokenDown struct{ antlr.Token }
type TokenUp struct{ antlr.Token }
type TokenEof struct{ antlr.Token }

func (TokenDown) String() string { return "<DOWN>" }
func (TokenUp) String() string   { return "<UP>" }
func (TokenEof) String() string  { return "<EOF>" }

func (t *TreeNode) GetTokenType() int { return t.TType }
func (tn TreeNode) String() string {
	if tn.TNText != nil {
		return *tn.TNText
	}
	return fmt.Sprintf("%v", tn.Val)
}

func NewWalkableBuild(name string, ttt TreeTokenTypes, root antlr.Token) WalkableBuilder {
	t := NewTree(name, root)
	b := &walkerbuilder{t, ttt, root, nil}
	return b
}
func (b *walkerbuilder) Add(n antlr.Token) WalkableBuilder {
	ok, _ := b.tree.Add(b.curr, n)
	if !ok {
		panic("Can't replace in a walkerbuilder")
	}
	b.last = n
	return b
}

func NewBuild(tree_name string, ttt TreeTokenTypes, n antlr.Token, ttype int, val interface{}) WalkableBuilder {
	if _, ok := val.(antlr.Token); ok {
		panic("trying to add a token as a node - this is just to confusing to be allowed")
	}
	tn := &TreeNode{Token: n, TType: ttype, Val: val}
	t := NewTree(tree_name, tn)
	b := &walkerbuilder{t, ttt, tn, nil}
	return b
}

var (
	UP  = "UP"
	DN  = "DOWN"
	EOF = "EOF"
)

func (b *walkerbuilder) AddNode(an antlr.Token, ttype int, val interface{}) WalkableBuilder {
	if _, ok := val.(antlr.Token); ok {
		panic("trying to add a token as a node - this is just to confusing to be allowed")
	}
	tn := &TreeNode{Token: an, TType: ttype, Val: val}
	switch ttype {
	case b.ttt.Up():
		tn.Token = &TokenUp{Token: an}
	case b.ttt.Down():
		tn.TNText = &DN
	case b.ttt.Eof():
		tn.TNText = &EOF
	}
	b.tree.Add(b.curr, tn)
	b.last = tn
	return b
}
func (b *walkerbuilder) Current() antlr.Token {
	return b.last
}
func (b *walkerbuilder) Down() WalkableBuilder {
	//	if b.last != nil {
	b.curr = b.last
	//	}
	return b
}
func (b *walkerbuilder) Up() WalkableBuilder {
	b.curr = b.tree.Parent(b.curr).(antlr.Token)
	b.last = b.curr
	return b
}
func (b *walkerbuilder) Build() Tree {
	return b.tree
}

type TreeToken struct {
	TType      int
	Start      int // optional return -1 if not implemented.
	Stop       int // optional return -1 if not implemented.
	TokenIndex int
}

func (t *TreeToken) GetTokenType() int { return t.TType }

var _ antlr.Token = &TreeToken{}

func (tt *TreeToken) GetSource() *antlr.TokenSourceCharStreamPair {
	var source *antlr.TokenSourceCharStreamPair
	return source
}
func (tt *TreeToken) GetChannel() int                   { return antlr.TokenDefaultChannel }
func (tt *TreeToken) GetStart() int                     { return tt.Start }
func (tt *TreeToken) GetStop() int                      { return tt.Stop }
func (tt *TreeToken) GetLine() int                      { return 0 }
func (tt *TreeToken) GetColumn() int                    { return tt.Start }
func (tt *TreeToken) GetText() string                   { return fmt.Sprintf("%v", tt) }
func (tt *TreeToken) SetText(s string)                  { panic("") }
func (tt *TreeToken) GetTokenIndex() int                { return tt.TokenIndex }
func (tt *TreeToken) SetTokenIndex(v int)               { panic("") }
func (tt *TreeToken) GetTokenSource() antlr.TokenSource { panic("") }
func (tt *TreeToken) GetInputStream() antlr.CharStream  { panic("") }

type TreeTokenSource struct {
	tttype TreeTokenTypes

	tree     Tree
	iterator Iterator
	nodes    []INode

	index int
}

type TreeTokenTypes interface {
	Eof() int
	Down() int
	Up() int
}

func NewTreeTokenSource(tree Tree, tttype TreeTokenTypes) antlr.TokenStream {
	tts := &TreeTokenSource{}
	tts.tree = tree
	tts.tttype = tttype
	tts.nodes = make([]INode, 0, tree.Size())
	tts.iterator = NewPreOrderTreeIterator(tree, tree.Root())
	for tts.iterator.HasNext() {
		n := tts.iterator.Next()
		tts.nodes = append(tts.nodes, n)
	}
	return tts
}

func NewTreeTokenSourceFromStart(tree Tree, tttype TreeTokenTypes, start INode) antlr.TokenStream {
	tts := &TreeTokenSource{}
	tts.tree = tree
	tts.tttype = tttype
	tts.nodes = make([]INode, 0, tree.Size())
	tts.iterator = NewPreOrderTreeIterator(tree, start)
	for tts.iterator.HasNext() {
		n := tts.iterator.Next()
		tts.nodes = append(tts.nodes, n)
	}
	return tts
}

// same a common token
func (tts *TreeTokenSource) Mark() int                                         { return 0 }
func (tts *TreeTokenSource) Release(marker int)                                {}
func (tts *TreeTokenSource) Index() int                                        { return tts.index }
func (tts *TreeTokenSource) Seek(index int)                                    { tts.index = index }
func (tts *TreeTokenSource) Size() int                                         { return len(tts.nodes) }
func (tts *TreeTokenSource) GetSourceName() string                             { return tts.tree.Name() }
func (tts *TreeTokenSource) GetTokenSource() antlr.TokenSource                 { return nil }
func (tts *TreeTokenSource) SetTokenSource(antlr.TokenSource)                  {}
func (tts *TreeTokenSource) GetAllText() string                                { return "" }
func (tts *TreeTokenSource) GetTextFromInterval(*antlr.Interval) string        { return "" }
func (tts *TreeTokenSource) GetTextFromRuleContext(antlr.RuleContext) string   { return "" }
func (tts *TreeTokenSource) GetTextFromTokens(antlr.Token, antlr.Token) string { return "" }

func (tts *TreeTokenSource) Consume() {
	tts.index = tts.index + 1
}
func (tts *TreeTokenSource) LA(i int) int {
	lt := tts.LT(i)
	if lt == nil {
		return 0
	}
	return lt.GetTokenType()
}

func (tts *TreeTokenSource) LT(k int) (ret antlr.Token) {
	// defer func() {
	// 	fmt.Printf("LT %#+v\n", ret)
	// }()
	if k < 1 {
		return nil
	}
	if k+tts.index >= len(tts.nodes) {
		tt := &TreeToken{
			TType:      antlr.TokenEOF,
			Start:      k + tts.index,
			Stop:       k + tts.index,
			TokenIndex: k + tts.index,
		}
		return tt
	}

	n := tts.nodes[k+tts.index-1]

	var ttype int
	switch n := n.(type) {
	case antlr.Token:
		return n
	case *Eof:
		ttype = tts.tttype.Eof()
	case *Down:
		ttype = tts.tttype.Down()
	case *Up:
		ttype = tts.tttype.Up()
	default:
		panic(fmt.Errorf("%T %v", n, n))
	}

	tt := &TreeToken{
		TType:      ttype,
		Start:      k + tts.index,
		Stop:       k + tts.index,
		TokenIndex: k + tts.index,
	}
	return tt
}

func (tts *TreeTokenSource) Get(index int) antlr.Token {
	// if index > len(tts.nodes) {
	// 	return &TreeTokenEOF{}
	// }
	index = index - 1
	n := tts.nodes[index]
	var ttype int
	switch n := n.(type) {
	case antlr.Token:
		return n
	case *Eof:
		ttype = tts.tttype.Eof()
	case *Down:
		ttype = tts.tttype.Down()
	case *Up:
		ttype = tts.tttype.Up()
	default:
		panic("")
	}
	tt := &TreeToken{
		TType:      ttype,
		Start:      index,
		Stop:       index,
		TokenIndex: index,
	}
	return tt
}
