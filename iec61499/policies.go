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

//A PFBEnforcerPolicy is what goes inside a PFBEnforcer, it is derived from a Policy
type PFBEnforcerPolicy struct {
	InternalVars []Variable
	States       []PFBState
	Transitions  []PFBSTTransition
}

//A PFBEnforcer will store a given input and output policy and can derive the enforcers required to uphold them
type PFBEnforcer struct {
	interfaceList InterfaceList
	Name          string
	OutputPolicy  PFBEnforcerPolicy
	InputPolicy   PFBEnforcerPolicy
}

//MakePFBEnforcer will convert a given policy to an enforcer for that policy
func MakePFBEnforcer(i InterfaceList, p PolicyFB) (*PFBEnforcer, error) {
	//make the enforcer
	enf := &PFBEnforcer{interfaceList: i, Name: p.Name}
	//first, convert policy transitions
	outpTr, err := p.GetPFBSTTransitions()
	if err != nil {
		return nil, err
	}
	splOutpTr := SplitPFBSTTransitions(outpTr)
	enf.OutputPolicy = PFBEnforcerPolicy{
		InternalVars: p.InternalVars,
		States:       p.States,
		Transitions:  splOutpTr,
	}

	return enf, nil
}

//GetPFBSTTransitions will convert all internal PFBTransitions into PFBSTTransitions (i.e. PFBTransitions with a ST symbolic tree condition)
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
//it relies on the SplitExpressionsOnOr function below
func SplitPFBSTTransitions(cTrans []PFBSTTransition) []PFBSTTransition {
	brTrans := make([]PFBSTTransition, 0)

	for i := 0; i < len(cTrans); i++ {
		cTran := cTrans[i]
		splitTrans := SplitExpressionsOnOr(cTran.STGuard)
		for j := 0; j < len(splitTrans); j++ {
			newTrans := PFBSTTransition{
				PFBTransition: cTran.PFBTransition,
			}
			//recompile the condition
			newTrans.PFBTransition.Condition = stconverter.STCompileExpression(splitTrans[len(splitTrans)-j-1])
			newTrans.STGuard = splitTrans[len(splitTrans)-j-1]

			brTrans = append(brTrans, newTrans)
		}
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
