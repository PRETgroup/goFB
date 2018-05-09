package iec61499

import (
	"fmt"

	"github.com/PRETgroup/goFB/goFB/stconverter"
)

//FBECCGuardToSTExpression converts a given FB's guard into a STExpression parsetree
func FBECCGuardToSTExpression(pName, guard string) ([]stconverter.STInstruction, *stconverter.STParseError) {
	return stconverter.ParseString(pName, guard)
}

//PFBSTTransition is a container struct for a PFBTransition and its ST translated guard
type PFBSTTransition struct {
	PFBTransition
	STGuard stconverter.STExpression
}

//GetPFBSTTransitions will scan a PolicyFB for violation transitions where the guard has multiple clauses separated by ORs
//Where found, it will break them into two transitions
func (p *PolicyFB) GetPFBSTTransitions() ([]PFBSTTransition, error) {
	stTrans := make([]PFBSTTransition, len(p.Transitions))
	for i := 0; i < len(p.Transitions); i++ {
		stguard, err := FBECCGuardToSTExpression(p.Name, p.Transitions[i].Condition)
		if err != nil {
			return nil, err
		}
		if len(stguard) != 1 {
			return nil, fmt.Errorf("Incompatible policy guard (wrong number of expressions), File:%v, Line:%v", p.Transitions[i].DebugInfo.SourceFile, p.Transitions[i].DebugInfo.SourceLine)
		}
		expr, ok := stguard[0].(stconverter.STExpression)
		if !ok {
			return nil, fmt.Errorf("Incompatible policy guard (not an expression), File:%v, Line:%v", p.Transitions[i].DebugInfo.SourceFile, p.Transitions[i].DebugInfo.SourceLine)
		}
		stTrans[i] = PFBSTTransition{
			PFBTransition: p.Transitions[i],
			STGuard:       expr,
		}
	}
	return stTrans, nil
}

//SplitPFBSTTransitions will take a slice of PFBSTTRansitions and then split transitions which have OR clauses
//into multiple transitions
func SplitPFBSTTransitions(cTrans []PFBSTTransition) []PFBSTTransition {
	brTrans := make([]PFBSTTransition, 0)
	//"a or (b and c)" should become "a", "(b and c)"
	//"a and (b or c)" should become "a and b", "a and c"
	//[and a, [or b, c]]
	//ROOT = and
	for i := 0; i < len(cTrans); i++ {
		cTran := cTrans[i]
		createGuards := make([]stconverter.STExpression, 0)
		createGuards = append(createGuards, cTran.STGuard)

		// for i := 0; i < len(createGuards); i++ {
		// 	guard := createGuards[i]
		// 	op := guard.HasOperator()
		// 	for {
		// 		if op == nil {
		// 			continue
		// 		}
		// 		if guard.GetArguments()

		// 		if op.GetToken() == "or" { //todo: this should be const defined somewhere
		// 			//break either side of the or
		// 			arg1
		// 		}
		// 	}
		// }
		//
		// for the expression
		//	if value, done
		//	if not or, enter the subentry
		//	if or, break and copy
	}

	//reformat all the guards based off the transactions
	return brTrans
}

//[a] should become [a]
//[or a b] should become [a] [b]
//[or a [b and c]] should become [a] [b and c]
func traverse(expr stconverter.STExpression) []stconverter.STExpression {
	//IF IS OR
	//	BREAK APART
	//IF IS VALUE
	//	RETURN CURRENT
	//IF IS OTHER OPERATOR
	//	MARK LOCATION AND RECURSE

	// broken := breakIfOr(expr)
	// if len(broken) == 1 {
	// 	return broken
	// }

	op := expr.HasOperator()
	if op == nil {
		return []stconverter.STExpression{expr}
	}
	if op.GetToken() == "or" {
		return expr.GetArguments()
	}
	//otherwise, things are more interesting

	//the thing we're returning
	rets := make([]stconverter.STExpressionOperator, 0)

	//build a new expression
	var nExpr stconverter.STExpressionOperator

	//operator is op, arguments are args
	nExpr.Operator = op
	args := expr.GetArguments()
	nExpr.Arguments = make([]stconverter.STExpression, len(args))

	rets = append(rets, nExpr)

	//each argument should be only one value, and we can check that by calling traverse again
	for i, arg := range args {
		argT := traverse(arg)
		//if argT has more than one value, it indicates that this argument was "broken", and we should return two nExpr, one with each argument
		//if we need more rets
		for len(argT) > len(rets) {
			rets = append(rets, rets[len(rets)-1])
		}
		c := 0 //keep track of how many times we updated
		for j, at := range argT {
			fmt.Printf("at:%v, i:%v, len(rets):%v\n", at, i, len(rets))
			fmt.Printf("appending %v\n", at)
			rets[j].Arguments[i] = at
			c++
		}
		//fill in any remaining spaces with the first element of argT
		for j := c; j < len(rets); j++ {
			rets[j].Arguments[i] = argT[len(argT)-1]
		}
	}

	//conversion for returning
	actualRets := make([]stconverter.STExpression, len(rets))
	for i := 0; i < len(rets); i++ {
		actualRets[i] = rets[i]
	}
	return actualRets

}

//breakIfOr: the goal of this is to recreate the expression exactly, but if it has an "or", break it into two
func breakIfOr(expr stconverter.STExpression) []stconverter.STExpression {
	op := expr.HasOperator()
	if op == nil {
		return []stconverter.STExpression{expr}
	}
	if op.GetToken() != "or" {
		return []stconverter.STExpression{expr}
	}
	return expr.GetArguments()
}
