package stconverter

import "errors"

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

// func convertInfixExpressionToPrefix(infix []string) ([]string, *STParseError) {
// 	operands := make([]string, 0)
// 	operators := make([]string, 0)

// 	for len(infix) > 0 {
// 		in := infix[len]
// 	}
// 	/*//we first want an A argument
// 	if s == stOpenBracket {
// 		expr, err := t.parseExpressionTerminatesWith(stCloseBracket)
// 		if err != nil {
// 			return nil, err
// 		}
// 		curExpr.A = expr
// 	} else {
// 		curExpr.AValue = s
// 	}

// 	//pop the next thing, if we terminate, time to go home
// 	s = t.pop()
// 	if s == terminates {
// 		return &topExpr, nil
// 	}

// 	//now we are looking for an operator?
// 	//if lookingForOp {
// 	foundOp := false
// 	for _, op := range stOps {
// 		if s == op.token {
// 			foundOp := true
// 			break
// 			//the returned precedence should effect where the operation gets put

// 		}
// 	}

// 	if foundOp == false {
// 		//we were looking for an operator but didn't find one
// 		return nil, t.errorUnexpectedWithExpected(s, "[any valid operator]")
// 	}*/

// }

func (t *stParse) parseExpressionTerminatesWith(terminates string) (*STExpression, *STParseError) {
	//step 1, convert from infix to prefix
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
	prefixExprString, err := t.convertInfixExpressionToPrefix(infixExprString)
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
	expr, err := t.parseExpressionTerminatesWith(stThen)
	if err != nil {
		return err
	}

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
	return t.error(errors.New("not yet implemented"))
}
