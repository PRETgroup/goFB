package postfix

import (
	"reflect"
	"testing"
)

type stOp struct {
	token       string
	precedence  int
	numOperands int
	association Association
}

func (s stOp) GetToken() string {
	return s.token
}

func (s stOp) GetPrecedence() int {
	return s.precedence
}

func (s stOp) GetNumOperands() int {
	return s.numOperands
}

func (s stOp) GetAssociation() Association {
	return s.association
}

const (
	stNot                = "not"
	stExponentiation     = "**"
	stMultiply           = "*"
	stDivide             = "/"
	stModulo             = "MOD"
	stAdd                = "+"
	stSubtract           = "-"
	stLessThan           = "<"
	stGreaterThan        = ">"
	stLessThanEqualTo    = "<="
	stGreaterThanEqualTo = ">="
	stEqual              = "="
	stInequal            = "<>"
	stAnd                = "and"
	stExlusiveOr         = "xor"
	stOr                 = "or"
)

var ops = []Operator{
	stOp{stNot, 0, 1, AssociationRight},
	stOp{stExponentiation, 1, 2, AssociationRight},
	stOp{stMultiply, 2, 2, AssociationLeft},
	stOp{stDivide, 2, 2, AssociationLeft},
	stOp{stModulo, 2, 2, AssociationLeft},
	stOp{stAdd, 3, 2, AssociationLeft},
	stOp{stSubtract, 3, 2, AssociationLeft},
	stOp{stLessThan, 4, 2, AssociationLeft},
	stOp{stGreaterThan, 4, 2, AssociationLeft},
	stOp{stLessThanEqualTo, 4, 2, AssociationLeft},
	stOp{stGreaterThanEqualTo, 4, 2, AssociationLeft},
	stOp{stEqual, 4, 2, AssociationLeft},
	stOp{stInequal, 4, 2, AssociationLeft},
	stOp{stAnd, 5, 2, AssociationLeft},
	stOp{stExlusiveOr, 5, 2, AssociationLeft},
	stOp{stOr, 5, 2, AssociationLeft},
}

type postfixTest struct {
	in, out []string
}

var tests = []postfixTest{
	{
		in:  []string{"1", "+", "2"},
		out: []string{"1", "2", "+"},
	},
	{
		in:  []string{"1", "+", "2", "*", "3"},
		out: []string{"1", "2", "3", "*", "+"},
	},
	{
		in:  []string{"1", "/", "2", "*", "3"},
		out: []string{"1", "2", "/", "3", "*"},
	},
	{
		in:  []string{"(", "1", "+", "(", "2", "+", "3", ")", "+", "4", ")", "*", "3"},
		out: []string{"1", "2", "3", "+", "+", "4", "+", "3", "*"},
	},
	{
		in:  []string{"3", "+", "4", "*", "2", "/", "(", "1", "-", "5", ")", "**", "2", "**", "3"},
		out: []string{"3", "4", "2", "*", "1", "5", "-", "2", "3", "**", "**", "/", "+"},
	},
	{
		in:  []string{"not", "1"},
		out: []string{"1", "not"},
	},
	{
		in:  []string{"1", "+", "not", "2"},
		out: []string{"1", "2", "not", "+"},
	},
	{
		in:  []string{"sin(", "max(", "2", "3", ")", "/", "3", "*", "3.1415"},
		out: []string{"2", "3", "max<2>", "3", "/", "3.1415", "*", "sin<1>"},
	},
}

func TestPostfix(t *testing.T) {
	c := NewConverter(ops)
	for i := 0; i < len(tests); i++ {
		pOut := c.ToPostfix(tests[i].in)
		if !reflect.DeepEqual(tests[i].out, pOut) {
			t.Errorf("failed test %d\nout:  %+v\npOut: %+v\n", i, tests[i].out, pOut)
		}
	}
}
