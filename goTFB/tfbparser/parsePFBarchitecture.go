package tfbparser

import (
	"strings"

	"github.com/PRETgroup/goFB/iec61499"
)

//parsePFBarchitecture shall only be called once we have already parsed the
// "architecture of [blockname]" part of the definition
// so, we are up to the brace
func (t *tfbParse) parsePFBarchitecture(fbIndex int) *ParseError {
	if s := t.pop(); s != pOpenBrace {
		return t.errorUnexpectedWithExpected(s, pOpenBrace)
	}
	//we now have several things that could be in here
	//internal | internals | state | states | closeBrace

	//unlike in an interface, the various things that are in an architecture can be presented out of order
	//this only has consequences with regards to states in the state machine
	//because we can't verify them "on-the-fly" (a state might point to a state we've not yet parsed)
	//Situations like this is the main reason most non-syntax parse-related validation is done in the iec61499 package

	for {
		s := t.pop()
		if s == "" {
			return t.error(ErrUnexpectedEOF)
		} else if s == pCloseBrace {
			//this is the end of the architecture
			break
		} else if s == pInternal || s == pInternals { //we actually care about { vs not-{, and so either internal or internals are valid prefixes for both situations
			if err := t.parsePossibleArrayInto(fbIndex, (*tfbParse).parsePFBInternal); err != nil {
				return err
			}
		} else if s == pState || s == pStates {
			if err := t.parsePossibleArrayInto(fbIndex, (*tfbParse).parsePFBState); err != nil {
				return err
			}
		}
	}

	return nil
}

//parsePFBInternal parses a single internal and adds it to fb identified by fbIndex
func (t *tfbParse) parsePFBInternal(fbIndex int) *ParseError {
	//the beginning of this is very similar to parseFBio, but different enough that it should be another function
	fb := &t.fbs[fbIndex]

	//next s is type
	typ := t.pop()
	debug := t.getCurrentDebugInfo()

	if !isValidType(typ) {
		return t.errorWithArgAndReason(ErrInvalidType, typ, "Expected valid type")
	}

	if typ == pEvent {
		return t.errorWithArgAndReason(ErrInvalidType, typ, "Internals cannot be event type")
	}

	var intNames []string

	//there might be an array size next
	size := ""
	if t.peek() == pOpenBracket {
		t.pop() // get rid of open bracket
		size = t.pop()
		if s := t.peek(); s != pCloseBracket {
			return t.errorUnexpectedWithExpected(s, pCloseBracket)
		}
		t.pop() //get rid of close bracket
	}

	for {
		name := t.pop()

		intNames = append(intNames, name)
		if t.peek() == pComma {
			t.pop() //get rid of the pComma
			continue
		}
		if t.peek() == pInitEq {
			break
		}
		break
	}

	//there might be a default value next
	initialValue := ""
	if t.peek() == pInitEq {
		t.pop() //get rid of pInitial

		s := t.pop()           //this might be an openbracket
		if s == pOpenBracket { //for arrays
			initialValue += s //we need to keep the brackets in
			for {
				s := t.pop()
				if s == "" {
					return t.error(ErrUnexpectedEOF)
				}
				if s == pSemicolon {
					return t.errorUnexpectedWithExpected(s, pOpenBracket)
				}
				if s == pCloseBracket {
					initialValue += s //we need to keep the brackets in
					break
				}
				initialValue += s
			}
		} else { //wasn't an open bracket, must just be value
			initialValue = s
		}
	}

	if t.peek() == pWith { //special error case to be helpful
		return t.errorWithArgAndReason(ErrUnexpectedAssociation, "with", "Internals cannot be associated with events")
	}

	//clear out last semicolon
	if s := t.pop(); s != pSemicolon {
		return t.errorUnexpectedWithExpected(s, pSemicolon)
	}

	//we now have everything we need to add the internal to the fb

	//while this can return an error,
	//the only permissible error is "wrong block type" and we have already ensured we are operating on a basicFB
	if _, err := fb.AddxFBDataInternals(intNames, typ, size, initialValue, debug); err != nil {
		return t.error(err)
	}

	return nil
}

//parsePFBState parses a single state and adds it to fb identified by fbIndex
// most things in this function are validated later in the iec61499 package
func (t *tfbParse) parsePFBState(fbIndex int) *ParseError {
	fb := &t.fbs[fbIndex]

	//next is name of state
	name := t.pop()

	for _, st := range fb.PolicyFB.States {
		if st.Name == name {
			return t.errorWithArg(ErrNameAlreadyInUse, name)
		}
	}
	debug := t.getCurrentDebugInfo()

	//next should be open brace
	if s := t.pop(); s != pOpenBrace {
		return t.errorUnexpectedWithExpected(s, pOpenBrace)
	}

	//now we have an unknown number of ->s
	// format is -> <destination> [on guard] [: output expression][, output expression...] ;
	for {
		s := t.pop()
		if s == "" {
			return t.error(ErrUnexpectedEOF)
		}
		if s == pCloseBrace {
			break
		}
		if s == pRun {
			//PFBs don't have algorithms
			return t.errorWithArgAndReason(ErrUnexpectedValue, pRun, "PFBs don't have algorithms")

		}
		if s == pEmit {
			//PFBs don't emit
			return t.errorWithArgAndReason(ErrUnexpectedValue, pEmit, "PFBs don't emit")
		}
		if s == pTrans {

			//next is dest state
			destState := t.pop()
			debug := t.getCurrentDebugInfo()

			var condComponents []string
			//next is on if we have a condition
			if t.peek() == pOn {
				t.pop() //clear the pOn

				//now we have an unknown number of condition components, terminated by a semicolon
				for {
					//pColon means that there are EXPRESSIONS that follow, but we're done here
					//pSemicolon means that there is NOTHING that follows, and we're done here
					if t.peek() == pColon || t.peek() == pSemicolon {
						break
					}

					s = t.pop()
					if s == "" {
						return t.error(ErrUnexpectedEOF)
					}

					//if any condComponent is "and" then turn it into &&
					if s == "and" {
						s = "&&"
					}
					//if any condComponint is "or" then turn it into ||
					if s == "or" {
						s = "||"
					}
					condComponents = append(condComponents, s)

				}
			}
			if len(condComponents) == 0 { //put in a default condition if no condition exists
				condComponents = append(condComponents, "true")
			}

			var expressions []iec61499.PFBExpression
			var expressionComponents []string
			var expressionVar string
			//if we broke on a colon, then we now have EXPRESSIONS to parse
			if t.peek() == pColon {
				t.pop() //clear the pColon
				//the format is
				// VARIABLE := EXPRESSION [, VARIABLE := EXPRESSION]
				expressionVar = ""
				for {
					if t.peek() == pSemicolon || t.peek() == pComma {
						//finish the previous expression (if possible, indicated by expressionVar) and start the next one (if available, indicated by a comma)
						if expressionVar != "" {
							expressions = append(expressions, iec61499.PFBExpression{
								VarName: expressionVar,
								Value:   strings.Join(expressionComponents, " "),
							})
							expressionVar = ""
						}
						if t.peek() == pComma {
							t.pop()
							continue
						}
						break
					}
					s = t.pop()
					if s == "" {
						return t.error(ErrUnexpectedEOF)
					}
					//we already dealt with case where it's a comma or a semicolon in the peek section above
					if expressionVar == "" { //we've not yet started the expression, so here's the "VARIABLE :=" part
						expressionVar = s
						s = t.pop()
						if s != pAssigment {
							return t.errorUnexpectedWithExpected(s, pAssigment)
						}
						continue
					} else {
						//now here's the condition components
						expressionComponents = append(expressionComponents, s)
					}
				}
			}

			if t.peek() != pSemicolon {
				return t.errorUnexpectedWithExpected(t.peek(), pSemicolon)
			}
			t.pop() //pop the pSemicolon
			//save the transition
			fb.PolicyFB.AddTransition(name, destState, strings.Join(condComponents, " "), expressions, debug)
		}
	}

	//everything is parsed, add it to the state machine
	fb.PolicyFB.AddState(name, debug)

	return nil
}
