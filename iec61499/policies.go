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

//SplitExpressionsOnOr will take a given STExpression and return a slice of STExpressions which are
//split over the "or" operators, e.g.
//[a] should become [a]
//[or a b] should become [a] [b]
//[or a [b and c]] should become [a] [b and c]
//[[a or b] and [c or d]] should become [a and c] [a and d] [b and c] [b and d]
func SplitExpressionsOnOr(expr stconverter.STExpression) []stconverter.STExpression {
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
	if op == nil { //if it's just a value, return
		return []stconverter.STExpression{expr}
	}
	if op.GetToken() == "or" { //if it's an "or", return the arguments
		rets := make([]stconverter.STExpression, 0)
		args := expr.GetArguments()
		for i := 0; i < len(args); i++ { //for each argument of the "or", return it, unless it is itself an "or" (in which case, expand further)
			arg := args[i]
			argOp := arg.HasOperator()
			if argOp == nil || argOp.GetToken() != "or" {
				rets = append(rets, arg)
				continue
			}
			args = append(args, arg.GetArguments()...)
		}
		return rets
	}

	//otherwise, things are more interesting

	//make the thing we're returning
	rets := make([]stconverter.STExpressionOperator, 0)

	//build a new expression
	var nExpr stconverter.STExpressionOperator

	//operator is op, arguments are args
	nExpr.Operator = op
	args := expr.GetArguments()
	nExpr.Arguments = make([]stconverter.STExpression, len(args))

	rets = append(rets, nExpr)
	//for each argument in the expression operator
	for i, arg := range args {
		//get arguments to operator by calling SplitExpressionsOnOr again
		argT := SplitExpressionsOnOr(arg)
		//if argT has more than one value, it indicates that this argument was "split", and we should return two nExpr, one with each argument
		//we will increase the size of rets by a multiplyFactor, which is the size of argT
		//i.e. if we receive two arguments, and we already had two elements in rets, it indicates we need to return 4 values
		//for instance, if our original command was "(a or b) and (c or d)" we'd need to return 4 elements (a and c) (a and d) (b and c) (b and d)
		multiplyFactor := len(argT)
		//for each factor in multiplyFactor, duplicate rets[n]
		//e.g. multiplyFactor 2 on [1 2 3] becomes [1 1 2 2 3 3]
		//e.g. multiplyFactor 3 on [1 2 3] becomes [1 1 1 2 2 2 3 3 3]
		for y := 0; y < len(rets); y++ {
			for z := 1; z < multiplyFactor; z++ {

				var newElem stconverter.STExpressionOperator
				copyElem := rets[y]
				newElem.Operator = copyElem.Operator
				newElem.Arguments = make([]stconverter.STExpression, len(copyElem.Arguments))
				copy(newElem.Arguments, copyElem.Arguments)

				rets = append(rets, stconverter.STExpressionOperator{})
				copy(rets[y+1:], rets[y:])
				rets[y] = newElem
				y++
			}
		}

		//for each argument, copy it into the return elements at the appropriate locations
		//(if we have multiple arguments, they will be chosen in a round-robin fashion)
		for j := 0; j < len(argT); j++ {
			at := argT[j]
			for k := j; k < len(rets); k += len(argT) {
				rets[k].Arguments[i] = at
			}
		}

		//expected, _ := json.MarshalIndent(rets, "\t", "\t")
		//fmt.Printf("Current:\n\t%s\n\n", expected)
	}

	//conversion for returning
	actualRets := make([]stconverter.STExpression, len(rets))
	for i := 0; i < len(rets); i++ {
		actualRets[i] = rets[i]
	}
	return actualRets

}
