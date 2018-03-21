package stconverter

import (
	"errors"
	"fmt"

	"github.com/PRETgroup/goFB/goFB/postfix"
)

func (t *stParse) parseNext() ([]STInstruction, *STParseError) {
	s := t.peek()
	//decide which parser to call
	//are we beginning a block {if, case, for, while, repeat}
	if s == stIf {
		return t.parseIfElsifElse()
	}

	if s == stCase {
		return t.parseSwitchCase()
	}

	if s == stFor {
		return t.parseForLoop()
	}

	if s == stWhile {
		return t.parseWhileLoop()
	}

	if s == stRepeat {
		return t.parseRepeatLoop()
	}

	//if still here, parse an assignment (the only valid remaining option)
	return t.parseAssignment()
}

func (t *stParse) parseExpressionTerminatesWith(terminates string) (STExpression, *STParseError) {
	//step 1, convert from infix to postfix
	//load the string tokens up to a match with "terminates"
	infixExprString := make([]string, 0)
	for {
		if t.done() {
			return nil, t.error(ErrUnexpectedEOF)
		}
		s := t.pop()
		//determine what s is
		if s == terminates {
			break
		}
		infixExprString = append(infixExprString, s)
	}

	//convert to postfix notation
	postfixConverter := postfix.NewConverter(stOps)
	postfixExprString := postfixConverter.ToPostfix(infixExprString)

	//now go through the postfix expression and convert to function tree
	//postfixExprString could look something like this: []string{"x", "y", "2", "z", "*", "max<2>", ">="},
	var stack []STExpression
	for i := 0; i < len(postfixExprString); i++ {
		token := postfixExprString[i]
		op := findOp(token)
		if op == nil {
			stack = append(stack, STExpressionValue{token})
			continue
		}
		//if op is not nil, then we use it (it is an operator)

		//create an stExpressionOperator
		stEOp := STExpressionOperator{}
		var e STExpression
		stEOp.Operator = op
		for j := 0; j < op.GetNumOperands(); j++ {
			e, stack = stack[len(stack)-1], stack[:len(stack)-1]
			stEOp.Arguments = append(stEOp.Arguments, e)
		}
		stack = append(stack, stEOp)
	}
	//now we're done!
	if len(stack) != 1 {
		fmt.Println(stack)
		return nil, t.error(ErrBadExpression)
	}
	s := stack[0]
	return s, nil
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
END_IF ; */
func (t *stParse) parseIfElsifElse() ([]STInstruction, *STParseError) {
	//the first word should be if
	s := t.pop()
	if s != stIf {
		return nil, t.errorUnexpectedWithExpected(s, stIf)
	}

	ifte := STIfElsIfElse{}

	//now we should get an expression terminated with "then"
	//("then" is consumed in this process)
	ifExpr, err := t.parseExpressionTerminatesWith(stThen)
	if err != nil {
		return nil, err
	}

	it := STIfThen{
		IfExpression: ifExpr,
	}

	//now we should get a then sequence terminated by either end_if or elsif
	for t.peek() != stElsif && t.peek() != stElse && t.peek() != stEndIf && !t.done() {
		seq, err := t.parseNext()
		if err != nil {
			return nil, err
		}
		it.ThenSequence = append(it.ThenSequence, seq...)
	}

	ifte.IfThens = append(ifte.IfThens, it)

	//if we have an elsif...
	for t.peek() == stElsif {
		t.pop()
		//terminate at then
		elsIfExpr, err := t.parseExpressionTerminatesWith(stThen)
		if err != nil {
			return nil, err
		}
		eit := STIfThen{
			IfExpression: elsIfExpr,
		}
		//now we should get a then sequence terminated by either end_if or elsif
		for t.peek() != stElsif && t.peek() != stElse && t.peek() != stEndIf && !t.done() {
			seq, err := t.parseNext()
			if err != nil {
				return nil, err
			}
			eit.ThenSequence = append(eit.ThenSequence, seq...)
		}
		ifte.IfThens = append(ifte.IfThens, eit)
	}

	//if we have an else
	if t.peek() == stElse {
		t.pop()
		for t.peek() != stEndIf && !t.done() {
			seq, err := t.parseNext()
			if err != nil {
				return nil, err
			}
			ifte.ElseSequence = append(ifte.ElseSequence, seq...)
		}
	}

	//now consume the stEndIf (we've only peeked at it until now)
	s = t.pop()
	if s != stEndIf {
		return nil, t.errorUnexpectedWithExpected(s, stEndIf)
	}

	//now consume the stSemicolon
	s = t.pop()
	if s != stSemicolon {
		return nil, t.errorUnexpectedWithExpected(s, stSemicolon)
	}

	return []STInstruction{ifte}, nil

}

func (t *stParse) parseSwitchCase() ([]STInstruction, *STParseError) {
	return nil, t.error(errors.New("not yet implemented"))
}

func (t *stParse) parseForLoop() ([]STInstruction, *STParseError) {
	return nil, t.error(errors.New("not yet implemented"))
}

func (t *stParse) parseWhileLoop() ([]STInstruction, *STParseError) {
	return nil, t.error(errors.New("not yet implemented"))
}

func (t *stParse) parseRepeatLoop() ([]STInstruction, *STParseError) {
	return nil, t.error(errors.New("not yet implemented"))
}

func (t *stParse) parseAssignment() ([]STInstruction, *STParseError) {
	//consumes stSemicolon
	ass, err := t.parseExpressionTerminatesWith(stSemicolon)
	if err != nil {
		fmt.Println("error", err)
		return nil, err
	}
	return []STInstruction{ass}, nil
}
