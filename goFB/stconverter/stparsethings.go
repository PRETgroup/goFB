package stconverter

import (
	"errors"
	"fmt"

	"github.com/PRETgroup/goFB/goFB/postfix"
)

func (t *stParse) parseNext() *STParseError {
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
func (t *stParse) parseIfElsifElse() *STParseError {
	//the first word should be if
	s := t.pop()
	if s != stIf {
		return t.errorUnexpectedWithExpected(s, stIf)
	}

	//now we should get an expression terminated with then
	// expr, err := t.parseExpressionTerminatesWith(stThen)
	// if err != nil {
	// 	return err
	// }

	return nil

}

func (t *stParse) parseSwitchCase() *STParseError {
	return t.error(errors.New("not yet implemented"))
}

func (t *stParse) parseForLoop() *STParseError {
	return t.error(errors.New("not yet implemented"))
}

func (t *stParse) parseWhileLoop() *STParseError {
	return t.error(errors.New("not yet implemented"))
}

func (t *stParse) parseRepeatLoop() *STParseError {
	return t.error(errors.New("not yet implemented"))
}

func (t *stParse) parseAssignment() *STParseError {
	ass, err := t.parseExpressionTerminatesWith(stSemicolon)
	if err != nil {
		fmt.Println("error", err)
		return err
	}
	t.instructions = append(t.instructions, ass)
	return nil
}
