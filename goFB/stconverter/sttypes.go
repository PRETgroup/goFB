package stconverter

import (
	"github.com/PRETgroup/goFB/iec61499/postfix"
)

//STExpression is an interface defining an assignments and comparison function tree
//E.g. A := (2 + y) can be defined using STExpressions
type STExpression interface {
	HasValue() string //IsValue reflects that this STExpression is an STExpressionValue, and is just a single variable or value
	HasOperator() postfix.Operator
	GetArguments() []STExpression
	IsInstruction() bool
}

//STExpressionValue is a type of STExpression that is just a single variable or value (i.e. an operand)
type STExpressionValue struct {
	Value string
}

//HasValue returns the internal value
func (s STExpressionValue) HasValue() string {
	return s.Value
}

//HasOperator returns nothing, as an STExpressionValue is not an operator
func (s STExpressionValue) HasOperator() postfix.Operator {
	return nil
}

//GetArguments returns the internal value as a single-element slice
func (s STExpressionValue) GetArguments() []STExpression {
	return []STExpression{s}
}

//STExpressionOperator is a type of STExpression that is an operator or a function with a list of arguments
type STExpressionOperator struct {
	Operator  postfix.Operator
	Arguments []STExpression
}

//HasValue returns nothing, as STExpressionOperator is not a value
func (s STExpressionOperator) HasValue() string {
	return ""
}

//HasOperator returns the internal operator
func (s STExpressionOperator) HasOperator() postfix.Operator {
	return s.Operator
}

//GetArguments returns the list of arguments
func (s STExpressionOperator) GetArguments() []STExpression {
	return s.Arguments
}

//STIfThen is used as part of an STIfElsIfElse
//  it holds the if [boolean expression] then <statements...>; part
type STIfThen struct {
	IfExpression STExpression
	ThenSequence []STInstruction
}

//STIfElsIfElse is used to make up the full if... elsif... elsif.... else... sequence
//  the ifThens are evaluated in order
//example:
/*
IF [boolean expression] THEN
<statement>;
ELSIF [boolean expression] THEN
    <statement>;
ELSE
    <statement>;
END_IF ;

A := 0;
IF A = 0 THEN
    B := 0;
END_IF ;
*/
type STIfElsIfElse struct {
	IfThens      []STIfThen
	ElseSequence []STInstruction
}

//STCase is used inside STSwitchCase to store the different case options
type STCase struct {
	CaseValues []string
	Sequence   []STInstruction
}

//STSwitchCase is used for the switch... case... case... else sequence
//examples:
/*
CASE [numeric expression] OF
    result1, result2: <statement>;
    resultN[, resultN]: <statemtent>;
ELSE
    <statement>;
END_CASE;

CASE StateMachine OF
	1:
		StateMachine := 2;
	2:
		StateMachine := 1;
ELSE
	StateMachine := 1;
END_CASE;
*/
type STSwitchCase struct {
	SwitchOn     STExpression
	Cases        []STCase
	ElseSequence []STInstruction
}

//STForLoop is used for for loops
//Example:
/*
FOR count := initial_value TO final_value BY increment DO
    <statement>;
END_FOR;
*/
type STForLoop struct {
	ForAssignment STExpression
	ToValue       STExpression
	ByIncrement   STExpression
	Sequence      []STInstruction
}

//FindCounterName will return the variable name assigned to in the ForAssignment, if one can easily be found
func (f STForLoop) FindCounterName() string {
	if f.ForAssignment == nil {
		return ""
	}
	faop := f.ForAssignment.HasOperator()
	if faop == nil {
		return ""
	}
	if faop.GetToken() != stAssignment {
		return ""
	}
	if len(f.ForAssignment.GetArguments()) < 2 {
		return ""
	}
	return f.ForAssignment.GetArguments()[1].HasValue()
}

//STWhileLoop is used for while loops
//Example:
/*
WHILE [boolean expression] DO
    <statement>;
END_WHILE;
*/
type STWhileLoop struct {
	WhileExpression STExpression
	Sequence        []STInstruction
}

//STRepeatLoop is used for Repeat....Until loops
//Example:
/*
REPEAT
    <statement>;
UNTIL [boolean expression]
END_REPEAT;
*/
type STRepeatLoop struct {
	UntilExpression STExpression
	Sequence        []STInstruction
}
