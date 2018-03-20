package stconverter

import (
	"github.com/PRETgroup/goFB/goFB/postfix"
)

//STExpression is an interface defining an assignments and comparison function tree
//E.g. A := (2 + y) can be defined using STExpressions
type STExpression interface {
	IsValue() (bool, string) //IsValue reflects that this STExpression is an STExpressionValue, and is just a single variable or value
	IsOperator() (bool, postfix.Operator)
	GetArguments() []STExpression
	IsInstruction() bool
}

//STExpressionValue is a type of STExpression that is just a single variable or value (i.e. an operand)
type STExpressionValue struct {
	Value string
}

//IsValue returns the internal value
func (s STExpressionValue) IsValue() (bool, string) {
	return true, s.Value
}

//IsOperator returns nothing, as an STExpressionValue is not an operator
func (s STExpressionValue) IsOperator() (bool, postfix.Operator) {
	return false, nil
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

//IsValue returns nothing, as STExpressionOperator is not a value
func (s STExpressionOperator) IsValue() (bool, string) {
	return false, ""
}

//IsOperator returns the internal operator
func (s STExpressionOperator) IsOperator() (bool, postfix.Operator) {
	return true, s.Operator
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
	CaseValue string
	Sequence  []STInstruction
}

//STSwitchCase is used for the switch... case... case... else sequence
//examples:
/*
CASE [numeric expression] OF
    result1: <statement>;
    resultN: <statemtent>;
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
	ForAssigment STExpression
	ToValue      STExpression
	ByIncrement  string
	Sequence     []STInstruction
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
