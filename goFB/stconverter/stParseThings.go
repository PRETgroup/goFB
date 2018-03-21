package stconverter

import (
	"errors"
	"fmt"

	"github.com/PRETgroup/goFB/goFB/postfix"
)

func (t *stParse) parseNext() (STInstruction, *STParseError) {
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

//pareExpressionTerminatesWith will run through the parse until it reaches a given termination value
//then it will convert that into a STExpression
func (t *stParse) parseExpressionTerminatesWith(terminates ...string) (STExpression, *STParseError) {
	//step 1, convert from infix to postfix
	//load the string tokens up to a match with "terminates"
	infixExprString := make([]string, 0)
out:
	for {
		if t.done() {
			return nil, t.error(ErrUnexpectedEOF)
		}
		s := t.peek()
		//determine if s terminates
		for _, t := range terminates {
			if s == t {
				break out
			}
		}
		t.pop()
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
func (t *stParse) parseIfElsifElse() (STInstruction, *STParseError) {
	//the first word should be if
	s := t.pop()
	if s != stIf {
		return nil, t.errorUnexpectedWithExpected(s, stIf)
	}

	ifte := STIfElsIfElse{}

	//now we should get an expression terminated with "then"
	ifExpr, err := t.parseExpressionTerminatesWith(stThen)
	if err != nil {
		return nil, err
	}
	t.pop() //consume then

	it := STIfThen{
		IfExpression: ifExpr,
	}

	//now we should get a then sequence terminated by either end_if or elsif
	for t.peek() != stElsif && t.peek() != stElse && t.peek() != stEndIf && !t.done() {
		seq, err := t.parseNext()
		if err != nil {
			return nil, err
		}
		it.ThenSequence = append(it.ThenSequence, seq)
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
		t.pop() //consume then
		eit := STIfThen{
			IfExpression: elsIfExpr,
		}
		//now we should get a then sequence terminated by either end_if or elsif
		for t.peek() != stElsif && t.peek() != stElse && t.peek() != stEndIf && !t.done() {
			seq, err := t.parseNext()
			if err != nil {
				return nil, err
			}
			eit.ThenSequence = append(eit.ThenSequence, seq)
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
			ifte.ElseSequence = append(ifte.ElseSequence, seq)
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

	return ifte, nil

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
*/
func (t *stParse) parseSwitchCase() (STInstruction, *STParseError) {
	//the first word should be case
	s := t.pop()
	if s != stCase {
		return nil, t.errorUnexpectedWithExpected(s, stCase)
	}

	sc := STSwitchCase{}

	//now we should get an expression terminated with "of"
	ofExpr, err := t.parseExpressionTerminatesWith(stOf)
	if err != nil {
		return nil, err
	}
	t.pop() //consume of
	sc.SwitchOn = ofExpr

	//now we have a number of cases marked as
	//[value][, value]: <statement>; <statement>; etc
cases:
	for {
		//to begin, get the case values
		scase := STCase{}
		for {
			scase.CaseValues = append(scase.CaseValues, t.pop())
			if t.peek() == stComma {
				t.pop()
				continue
			}
			break
		}
		//now we should have a colon
		if colon := t.pop(); colon != stColon {
			return nil, t.errorUnexpectedWithExpected(colon, stColon)
		}

		//now we have a sequence of instructions, terminated by the next case or terminated by else
	caseseq:
		for {
			//is the next instruction an "else" or an "end_case"?
			if t.peek() == stElse || t.peek() == stEndCase || t.done() {
				break caseseq
			}

			//is the next instruction the beginning of the next case?
			//the only way to tell this is to peek far ahead using the deepPeek instruction
			//to see which is first, a semicolon (indicating that it's a normal instruction)
			//or a colon, indicating that it's a case
			i := 0
		search:
			for {
				if t.deepPeek(i) == stColon {
					//the next instruction is the beginning of the next case
					break caseseq
				}
				if t.deepPeek(i) == stSemicolon {
					//the next instruction is a general instruction
					break search
				}

				//this ain't no infinite loop
				if t.itemIndex+i > len(t.items) {
					return nil, t.error(ErrUnexpectedEOF)
				}
				i++
			}
			//the next instruction must be a general instruction
			seq, err := t.parseNext()
			if err != nil {
				return nil, err
			}
			scase.Sequence = append(scase.Sequence, seq)
		}
		sc.Cases = append(sc.Cases, scase)
		if t.peek() == stElse || t.peek() == stEndCase || t.done() {
			break cases
		}
	}

	//if we have an else
	if t.peek() == stElse {
		t.pop()
		for t.peek() != stEndCase && !t.done() {
			seq, err := t.parseNext()
			if err != nil {
				return nil, err
			}
			sc.ElseSequence = append(sc.ElseSequence, seq)
		}
	}

	//now consume the stEndCase (we've only peeked at it until now)
	s = t.pop()
	if s != stEndCase {
		return nil, t.errorUnexpectedWithExpected(s, stEndCase)
	}

	//now consume the stSemicolon
	s = t.pop()
	if s != stSemicolon {
		return nil, t.errorUnexpectedWithExpected(s, stSemicolon)
	}

	return sc, nil
}

//STForLoop is used for for loops
//Example:
/*
FOR count := initial_value TO final_value BY increment DO
    <statement>;
END_FOR;
*/
func (t *stParse) parseForLoop() (STInstruction, *STParseError) {
	//the first word should be for
	s := t.pop()
	if s != stFor {
		return nil, t.errorUnexpectedWithExpected(s, stFor)
	}

	fl := STForLoop{}

	//now we should get an expression terminated with "to"
	forAss, err := t.parseExpressionTerminatesWith(stTo)
	if err != nil {
		return nil, err
	}
	t.pop() //consume "to"
	fl.ForAssignment = forAss

	//now we should get an expression terminated with "by" or "do"
	//(consumed in this process)
	toVal, err := t.parseExpressionTerminatesWith(stBy, stDo)
	if err != nil {
		return nil, err
	}

	fl.ToValue = toVal

	//if the next value is "by" then we need to do increment
	if t.pop() == stBy {
		incrVal, err := t.parseExpressionTerminatesWith(stDo)
		if err != nil {
			return nil, err
		}
		t.pop() //consume "do"
		fl.ByIncrement = incrVal
	}

	//now we should get a sequence terminated by end_for
	for t.peek() != stEndFor && !t.done() {
		seq, err := t.parseNext()
		if err != nil {
			return nil, err
		}
		fl.Sequence = append(fl.Sequence, seq)
	}

	//now consume the stEndIf (we've only peeked at it until now)
	s = t.pop()
	if s != stEndFor {
		return nil, t.errorUnexpectedWithExpected(s, stEndFor)
	}

	//now consume the stSemicolon
	s = t.pop()
	if s != stSemicolon {
		return nil, t.errorUnexpectedWithExpected(s, stSemicolon)
	}

	return fl, nil
}

//STWhileLoop is used for while loops
//Example:
/*
WHILE [boolean expression] DO
    <statement>;
END_WHILE;
*/
func (t *stParse) parseWhileLoop() (STInstruction, *STParseError) {
	//the first word should be while
	s := t.pop()
	if s != stWhile {
		return nil, t.errorUnexpectedWithExpected(s, stWhile)
	}

	wl := STWhileLoop{}

	//now we should get an expression terminated with "do"
	wExpr, err := t.parseExpressionTerminatesWith(stDo)
	if err != nil {
		return nil, err
	}
	t.pop() //consume "do"
	wl.WhileExpression = wExpr

	//now we should get a sequence terminated by end_while
	for t.peek() != stEndWhile && !t.done() {
		seq, err := t.parseNext()
		if err != nil {
			return nil, err
		}
		wl.Sequence = append(wl.Sequence, seq)
	}

	//now consume the stEndWhile (we've only peeked at it until now)
	s = t.pop()
	if s != stEndWhile {
		return nil, t.errorUnexpectedWithExpected(s, stEndWhile)
	}

	//now consume the stSemicolon
	s = t.pop()
	if s != stSemicolon {
		return nil, t.errorUnexpectedWithExpected(s, stSemicolon)
	}

	return wl, nil
}

//STRepeatLoop is used for Repeat....Until loops
//Example:
/*
REPEAT
    <statement>;
UNTIL [boolean expression]
END_REPEAT;
*/
func (t *stParse) parseRepeatLoop() (STInstruction, *STParseError) {
	return nil, t.error(errors.New("not yet implemented"))
}

func (t *stParse) parseAssignment() (STInstruction, *STParseError) {
	//consumes stSemicolon
	ass, err := t.parseExpressionTerminatesWith(stSemicolon)
	if err != nil {
		fmt.Println("error", err)
		return nil, err
	}
	t.pop() //consume semicolon
	return ass, nil
}
