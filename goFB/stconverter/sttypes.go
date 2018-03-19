package stconverter

const (
	stOpenBracket  = "("
	stCloseBracket = ")"

	stNoFurtherOperations = ""

	stAssignment = ":="

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

//STExpression stores primitives
//E.g. A = 2
type STExpression struct {
	A        *STExpression
	AValue   string //if A is nil, then use AValue
	Operator string
	B        *STExpression
}

//STAssignment is used to assign a value to A
//examples:
//  x := 2;
//  ON_OFF := (ONS_Trig AND NOT ON_OFF) OR (ON_OFF AND NOT ONS_Trig);
type STAssignment struct {
	AValue   string
	Assigned STExpression
}

//STIfThen is used as part of an STIfElsIfElse
//  it holds the if [boolean expression] then <statements...>; part
type STIfThen struct {
	Expression   STExpression
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
	ForAssigment STAssignment
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
