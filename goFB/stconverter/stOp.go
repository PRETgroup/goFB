package stconverter

import (
	"github.com/PRETgroup/goFB/iec61499/postfix"
)

type stOp struct {
	token       string
	precedence  int
	numOperands int
	association postfix.Association
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

func (s stOp) GetAssociation() postfix.Association {
	return s.association
}

func (s stOp) LeftAssociative() bool {
	return s.association == postfix.AssociationLeft
}

func (s stOp) MarshalJSON() ([]byte, error) {
	return []byte("\"" + s.token + "\""), nil
}

var stOps = []postfix.Operator{
	stOp{stExit, 0, 0, postfix.AssociationNone},   //not technically operators but they might appear in expressions (with no other things around them)
	stOp{stReturn, 0, 0, postfix.AssociationNone}, //not technically operators but they might appear in expressions (with no other things around them)
	stOp{stNot, 0, 1, postfix.AssociationRight},
	stOp{stNegative, 0, 1, postfix.AssociationRight},
	stOp{stExponentiation, 1, 2, postfix.AssociationRight},
	stOp{stMultiply, 2, 2, postfix.AssociationLeft},
	stOp{stDivide, 2, 2, postfix.AssociationLeft},
	stOp{stModulo, 2, 2, postfix.AssociationLeft},
	stOp{stAdd, 3, 2, postfix.AssociationLeft},
	stOp{stSubtract, 3, 2, postfix.AssociationLeft},
	stOp{stLessThan, 4, 2, postfix.AssociationLeft},
	stOp{stGreaterThan, 4, 2, postfix.AssociationLeft},
	stOp{stLessThanEqualTo, 4, 2, postfix.AssociationLeft},
	stOp{stGreaterThanEqualTo, 4, 2, postfix.AssociationLeft},
	stOp{stEqual, 4, 2, postfix.AssociationLeft},
	stOp{stInequal, 4, 2, postfix.AssociationLeft},
	stOp{stAnd, 5, 2, postfix.AssociationLeft},
	stOp{stExlusiveOr, 5, 2, postfix.AssociationLeft},
	stOp{stOr, 5, 2, postfix.AssociationLeft},
	stOp{stAssignment, 6, 2, postfix.AssociationLeft},
}

//FindOp finds a given operator for a given token
func FindOp(op string) postfix.Operator {
	for i := 0; i < len(stOps); i++ {
		if stOps[i].GetToken() == op {
			return stOps[i]
		}
	}
	//still here? might be a function
	if is, fn := postfix.IsFunction(op); is {
		return fn
	}
	//still here? not an operator
	return nil
}

//OpTokenIsComparison takes a given operator token and returns a true if it
//is of a comparison type (e.g. ">=")
func OpTokenIsComparison(opTok string) bool {
	if opTok == stGreaterThan ||
		opTok == stGreaterThanEqualTo ||
		opTok == stLessThan ||
		opTok == stLessThanEqualTo ||
		opTok == stEqual ||
		opTok == stInequal {

		return true
	}
	return false
}

//OpTokenIsCombinator takes a given operator token and returns a true if it
//is of a combination type (e.g. "and", "or")
func OpTokenIsCombinator(opTok string) bool {
	if opTok == stAnd ||
		opTok == stOr ||
		opTok == stExlusiveOr {

		return true
	}
	return false
}
