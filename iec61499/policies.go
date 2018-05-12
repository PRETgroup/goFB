package iec61499

import (
	"fmt"
	"reflect"
	"strings"

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

//GetDTimers returns all DTIMERS in a PFBEnforcerPolicy
func (pol PFBEnforcerPolicy) GetDTimers() []Variable {
	dTimers := make([]Variable, 0)
	for _, v := range pol.InternalVars {
		if strings.ToLower(v.Type) == "dtimer" {
			dTimers = append(dTimers, v)
		}
	}
	return dTimers
}

//GetViolationTransitions returns a slice of all transitions in this PFBEnforcerPolicy
//that have their destinations set to "violation", ie. are violation transitions
func (pol PFBEnforcerPolicy) GetViolationTransitions() []PFBSTTransition {
	violTrans := make([]PFBSTTransition, 0)
	for _, tr := range pol.Transitions {
		if tr.Destination == "violation" {
			violTrans = append(violTrans, tr)
		}
	}
	return violTrans
}

//GetNonViolationTransitions returns a slice of all transitions in this PFBEnforcerPolicy
//that have their destinations not set to "violation", ie. are not violation transitions
func (pol PFBEnforcerPolicy) GetNonViolationTransitions() []PFBSTTransition {
	nviolTrans := make([]PFBSTTransition, 0)
	for _, tr := range pol.Transitions {
		if tr.Destination != "violation" {
			nviolTrans = append(nviolTrans, tr)
		}
	}
	return nviolTrans
}

//DoesExpressionInvolveTime returns true if a given expression uses time
func (pol PFBEnforcerPolicy) DoesExpressionInvolveTime(expr stconverter.STExpression) bool {
	op := expr.HasOperator()
	if op == nil {
		return VariablesContain(pol.GetDTimers(), expr.HasValue())
	}
	for _, arg := range expr.GetArguments() {
		if pol.DoesExpressionInvolveTime(arg) {
			return true
		}
	}
	return false
}

//A PFBEnforcer will store a given input and output policy and can derive the enforcers required to uphold them
type PFBEnforcer struct {
	interfaceList InterfaceList
	Name          string
	OutputPolicy  PFBEnforcerPolicy
	InputPolicy   PFBEnforcerPolicy
}

//MakePFBEnforcer will convert a given policy to an enforcer for that policy
func MakePFBEnforcer(il InterfaceList, p PolicyFB) (*PFBEnforcer, error) {
	//make the enforcer
	enf := &PFBEnforcer{interfaceList: il, Name: p.Name}
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
	enf.OutputPolicy.RemoveDuplicateTransitions()

	enf.InputPolicy = DeriveInputEnforcerPolicy(il, enf.OutputPolicy)
	enf.InputPolicy.RemoveDuplicateTransitions()

	return enf, nil
}

//RemoveDuplicateTransitions will do a search through a policies transitions and remove any that are simple duplicates
//(i.e. every field the same and in the same order).
func (pol *PFBEnforcerPolicy) RemoveDuplicateTransitions() {
	for i := 0; i < len(pol.Transitions); i++ {
		for j := i + 1; j < len(pol.Transitions); j++ {
			if reflect.DeepEqual(pol.Transitions[i], pol.Transitions[j]) {
				pol.Transitions = append(pol.Transitions[:j], pol.Transitions[j+1:]...)
				j--
			}
		}
	}
}

//STExpressionSolution stores a solution to a violation transition
type STExpressionSolution struct {
	Expression string
	Comment    string
}

//SolveViolationTransition will attempt to solve a given transition
//TODO: consider where people have been too explicit with their time variables, and have got non-violating time-based transitions
//1. Check to see if there is a non-violating transition with an equivalent guard to the violating transition
//2. Select first solution
func (pol *PFBEnforcerPolicy) SolveViolationTransition(tr PFBSTTransition) STExpressionSolution {
	posSolTrs := make([]PFBSTTransition, 0) //possible Solution Transitions
	for _, propTr := range pol.Transitions {
		if propTr.Destination == "violation" {
			continue
		}
		if propTr.Source != tr.Source {
			continue
		}
		if pol.DoesExpressionInvolveTime(propTr.STGuard) {
			continue
		}
		posSolTrs = append(posSolTrs, propTr)
	}

	// Make sure there's at least one solution
	if len(posSolTrs) == 0 {
		return STExpressionSolution{Expression: "", Comment: "No possible solutions!"}
	}

	//1. Check to see if there is a non-violating transition with an equivalent guard to the violating transition
	for _, posSolTr := range posSolTrs {
		if reflect.DeepEqual(tr.STGuard, posSolTr.STGuard) {
			return STExpressionSolution{Expression: "", Comment: fmt.Sprintf("Selected non-violation transition \"%s -> %s on %s\" which has an equivalent guard, so no action is required", posSolTr.Source, posSolTr.Destination, posSolTr.Condition)}
		}
	}

	//2. Select first solution
	posSolTr := posSolTrs[0]
	solution := SolveSTExpression(posSolTr.STGuard)

	return STExpressionSolution{Expression: stconverter.CCompileExpression(solution), Comment: fmt.Sprintf("Selected non-violation transition \"%s -> %s on %s\" and action is required", posSolTr.Source, posSolTr.Destination, posSolTr.Condition)}

}

//DeriveInputEnforcerPolicy will derive an Input Policy from a given Output Policy
func DeriveInputEnforcerPolicy(il InterfaceList, outPol PFBEnforcerPolicy) PFBEnforcerPolicy {
	inpEnf := PFBEnforcerPolicy{
		States: outPol.States,
	}

	//inpEnf.InternalVars = nil
	//just realised that theres no internalVars that can't be managed by externalVars?
	inpEnf.InternalVars = make([]Variable, len(outPol.InternalVars))
	// for i := 0; i < len(outPol.InternalVars; i++) {}
	copy(inpEnf.InternalVars, outPol.InternalVars)

	//convert transitions and internal var names in transitions
	for i := 0; i < len(outPol.Transitions); i++ {
		inpEnf.Transitions = append(inpEnf.Transitions, ConvertPFBSTTransitionForInputPolicy(il, inpEnf.InternalVars, outPol.Transitions[i]))
	}

	// //convert internal var names on enforcer policy
	// for i := 0; i < len(inpEnf.InternalVars); i++ {
	// 	inpEnf.InternalVars[i].Name = inpEnf.InternalVars[i].Name + "_i"
	// }

	return inpEnf
}

//ConvertPFBSTTransitionForInputPolicy will convert a single PFBSTTransition from an Output Policy to its Input Policy Deriviation
func ConvertPFBSTTransitionForInputPolicy(il InterfaceList, intl []Variable, outpTrans PFBSTTransition) PFBSTTransition {
	retSTGuard := ConvertSTExpressionForInputPolicy(il, intl, outpTrans.STGuard)
	retTrans := outpTrans
	retTrans.STGuard = retSTGuard
	retTrans.Condition = stconverter.CCompileExpression(retSTGuard)
	return retTrans
}

//VariablesContain returns true if a list of variables contains a given name
func VariablesContain(vars []Variable, name string) bool {
	for i := 0; i < len(vars); i++ {
		if vars[i].Name == name {
			return true
		}
	}
	return false
}

//ConvertSTExpressionForInputPolicy will convert a single STExpression from an Output Policy transition guard to its Input Policy transition guard's Deriviation
//a == input
//b == output
//"a" becomes "a"
//"b" becomes "true" (technically becomes "true or not true")
//"a and b" becomes "a"
//"func(a, b)" becomes "func(a, true)"
//"!b" becomes "true" (technically becomes "not(true or not true)")
//TODO: a transition based only on time becomes nil?
func ConvertSTExpressionForInputPolicy(il InterfaceList, intl []Variable, expr stconverter.STExpression) stconverter.STExpression {
	//options
	//1. It is just a value
	//	  --if input or value, return
	//    --if output, return true
	//2. It is an operator
	//    Foreach arg
	//	      If arg

	//consider not(b)
	//op == not, args == []{b}
	//should return "true"

	op := expr.HasOperator()
	if op == nil { //if it's just a value, return if that value
		if il.HasOutput(expr.HasValue()) {
			return nil
		}
		// if VariablesContain(intl, expr.HasValue()) {
		// 	return stconverter.STExpressionValue{Value: expr.HasValue() + "_i"}
		// }
		return expr
	}

	args := expr.GetArguments()
	acceptableArgIs := make([]bool, 0)
	numAcceptable := 0
	acceptableArgs := make([]stconverter.STExpression, 0)
	//for each argument, we want to check if it is "acceptable", which here means
	//"is not a value that is an output var"
	//and
	//"if it is an operator, convert it via this function, and see if it is acceptable"
	for i := 0; i < len(args); i++ {
		arg := args[i]
		argOp := arg.HasOperator()

		if argOp == nil {
			//it is a value
			argV := stconverter.STExpressionValue{Value: arg.HasValue()}
			//see if it is acceptable
			if il.HasOutput(argV.HasValue()) {
				acceptableArgIs = append(acceptableArgIs, false)
				acceptableArgs = append(acceptableArgs, nil)
			} else {
				// if VariablesContain(intl, argV.HasValue()) {
				// 	argV.Value = argV.Value + "_i"
				// }
				acceptableArgIs = append(acceptableArgIs, true)
				acceptableArgs = append(acceptableArgs, argV)
				numAcceptable++
			}
			continue
		} else {
			//it is an operator, run the operator through this function and see if it is acceptable
			convArg := ConvertSTExpressionForInputPolicy(il, intl, args[i])
			if convArg != nil {
				acceptableArgIs = append(acceptableArgIs, true)
				acceptableArgs = append(acceptableArgs, convArg)
				numAcceptable++
			} else {
				acceptableArgIs = append(acceptableArgIs, false)
				acceptableArgs = append(acceptableArgs, nil)
			}
		}
	}

	//now we need to come up with a new STExpression to represent this expression and its arguments

	if numAcceptable < len(args) {
		//if less than the total args are acceptable, and only one argument is acceptable, then it is easy,
		//we can just return that one argument as an independent value
		//e.g. "(a and b)" becomes "a"
		if numAcceptable == 1 {
			for i := 0; i < len(acceptableArgIs); i++ {
				if acceptableArgIs[i] == true {
					return acceptableArgs[i]
				}
			}
		}
	}
	if numAcceptable == 0 {
		//if nothing at all is acceptable then it is easy, we just return nil
		return nil
	}

	//if we are still here, then it means that there is no easy answer, so we'll just make a new
	//STExpressionOperator, which has the same operator as we're currently examining
	//then, all unacceptable (i.e. nil) arguments should be replaced with simple value "true"
	actualArgs := make([]stconverter.STExpression, len(acceptableArgs))
	for i := 0; i < len(actualArgs); i++ {
		if acceptableArgs[i] != nil {
			actualArgs[i] = acceptableArgs[i]
		} else {
			actualArgs[i] = stconverter.STExpressionValue{Value: "true"}
		}
	}

	ret := stconverter.STExpressionOperator{
		Operator:  op,
		Arguments: actualArgs,
	}

	return ret
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
			newTrans.PFBTransition.Condition = stconverter.CCompileExpression(splitTrans[len(splitTrans)-j-1])
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

//SolveSTExpression will solve simple STExpressions
//The top level should be one of the following
//if VARIABLE ONLY, 			return VARIABLE = 1
//if NOT(VARIABLE) ONLY, 		return VARIABLE = 0
//if VARIABLE == EXPRESSION, 	return VARIABLE = VARIABLE
//if VARIABLE > EXPRESSION, 	return VARIABLE = EXPRESSION + 1
//if VARIABLE >= EXPRESSION, 	return VARIABLE = EXPRESSION
//if VARIABLE < EXPRESSION, 	return VARIABLE = EXPRESSION - 1
//if VARIABLE <= EXPRESSION, 	return VARIABLE = EXPRESSION
//if VARIABLE != EXPRESSION,	return VARIABLE = EXPRESSION + 1
//otherwise, return nil (can't solve)
func SolveSTExpression(problem stconverter.STExpression) stconverter.STExpression {
	op := problem.HasOperator()
	if op == nil { //if VARIABLE ONLY, 			return VARIABLE = 1
		return stconverter.STExpressionOperator{
			Operator: stconverter.FindOp(":="),
			Arguments: []stconverter.STExpression{
				stconverter.STExpressionValue{Value: "1"},
				stconverter.STExpressionValue{Value: problem.HasValue()},
			}}
	}
	args := problem.GetArguments()

	if op.GetToken() == "not" && len(args) == 1 { //if NOT(VARIABLE) ONLY, 		return VARIABLE = 1
		return stconverter.STExpressionOperator{
			Operator: stconverter.FindOp(":="),
			Arguments: []stconverter.STExpression{
				stconverter.STExpressionValue{Value: "0"},
				stconverter.STExpressionValue{Value: args[0].HasValue()},
			}}
	}

	return nil
}
