package adl

import (
	"fmt"
	"testing"

	antlr "github.com/wxio/goantlr"
	parser "github.com/wxio/tron-go/internal/adllp"
	"github.com/wxio/tron-go/internal/ctree"
)

func Test_EmptyPredict(t *testing.T) {
	var tttype *TTType
	tr := ctree.NewBuild("ADL", tttype, nil, parser.AdlPADL, nil).
		Add(&ctree.TreeToken{
			TType:      tttype.Down(),
			Start:      0,
			Stop:       0,
			TokenIndex: 2,
		}).
		Add(&ctree.TreeToken{
			TType:      tttype.Up(),
			Start:      0,
			Stop:       0,
			TokenIndex: 3,
		}).
		Build()
	ec := &errColl{}
	errs := WalkAdlWo(tr, ec)
	fmt.Printf("%+v\n", errs)
	for _, x := range ec.errs {
		fmt.Printf("%T\n", x)
	}
	fmt.Printf("%+v\n", ec.errs)
}

type errColl struct {
	errs []antlr.ErrorNode
}

func (s *errColl) VisitTerminal(node antlr.TerminalNode)      {}
func (s *errColl) EnterEveryRule(ctx antlr.ParserRuleContext) {}
func (s *errColl) ExitEveryRule(ctx antlr.ParserRuleContext)  {}
func (s *errColl) VisitErrorNode(node antlr.ErrorNode) {
	s.errs = append(s.errs, node)
}
