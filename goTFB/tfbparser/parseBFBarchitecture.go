package tfbparser

import (
	"strconv"
	"strings"

	"github.com/PRETgroup/goFB/iec61499"
)

//parseBFBarchitecture shall only be called once we have already parsed the
// "architecture of [blockname]" part of the definition
// so, we are up to the brace
func (t *tfbParse) parseBFBarchitecture(fbIndex int) *ParseError {
	if s := t.pop(); s != pOpenBrace {
		return t.errorUnexpectedWithExpected(s, pOpenBrace)
	}
	//we now have several things that could be in here
	//internal | internals | state | states | algorithm | algorithms | closeBrace

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
			if err := t.parsePossibleArrayInto(fbIndex, (*tfbParse).parseBFBInternal); err != nil {
				return err
			}
		} else if s == pAlgorithm || s == pAlgorithms {
			if err := t.parsePossibleArrayInto(fbIndex, (*tfbParse).parseBFBAlgorithm); err != nil {
				return err
			}
		} else if s == pState || s == pStates {
			if err := t.parsePossibleArrayInto(fbIndex, (*tfbParse).parseBFBState); err != nil {
				return err
			}
		}
	}

	return nil
}

//parseBFBInternal parses a single internal and adds it to fb identified by fbIndex
func (t *tfbParse) parseBFBInternal(fbIndex int) *ParseError {
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

	//there might be a default value next
	initialValue := ""
	if t.peek() == pInitial {
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

	for {
		name := t.pop()

		intNames = append(intNames, name)
		if t.peek() == pComma {
			t.pop() //get rid of the pComma
			continue
		}
		break
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

//parseBFBState parses a single state and adds it to fb identified by fbIndex
// most things in this function are validated later in the iec61499 package
func (t *tfbParse) parseBFBState(fbIndex int) *ParseError {
	fb := &t.fbs[fbIndex]

	//next is name of state
	name := t.pop()
	anonAlgIndex := 0 //used to count anonymous algorithms (to give them unique names)

	for _, st := range fb.BasicFB.States {
		if st.Name == name {
			return t.errorWithArg(ErrNameAlreadyInUse, name)
		}
	}
	debug := t.getCurrentDebugInfo()

	//next should be open brace
	if s := t.pop(); s != pOpenBrace {
		return t.errorUnexpectedWithExpected(s, pOpenBrace)
	}

	//now we have an unknown number of runs, emits, and ->s

	var runs []iec61499.Action
	var emits []iec61499.Action

	for {
		s := t.pop()
		if s == "" {
			return t.error(ErrUnexpectedEOF)
		}
		if s == pCloseBrace {
			break
		}
		if s == pRun {
			//capture any number of comma-separated algorithms to run
			for {
				r := t.pop()

				if r == pIn {
					//we have an anonymous algorithm here (introduced by the keyword pIn)
					anonAlgName := name + "_alg" + strconv.Itoa(anonAlgIndex)
					anonAlgIndex++
					lang, prog, err := t.getAlgorithmParts()
					if err != nil {
						return err
					}
					//we're done here, add the algorithm to the block
					fb.BasicFB.AddAlgorithm(anonAlgName, lang, prog, debug)
					//replace keyword pIn with the name of the new anon algorithm so it can be joined properly
					r = anonAlgName
				}

				runs = append(runs, iec61499.Action{Algorithm: r, DebugInfo: t.getCurrentDebugInfo()})
				if t.peek() == pComma {
					t.pop()
					continue
				}
				break

			}
			if s := t.pop(); s != pSemicolon {
				return t.errorUnexpectedWithExpected(s, pSemicolon)
			}
		}
		if s == pEmit {
			//capture any number of comma-separated events to emit
			for {
				e := t.pop()
				emits = append(emits, iec61499.Action{Output: e, DebugInfo: t.getCurrentDebugInfo()})
				if t.peek() == pComma {
					t.pop()
					continue
				}
				break
			}
			if s := t.pop(); s != pSemicolon {
				return t.errorUnexpectedWithExpected(s, pSemicolon)
			}
		}
		if s == pTrans {

			//next is dest state
			destState := t.pop()
			debug := t.getCurrentDebugInfo()

			var condComponents []string

			//next is on if we have a condition
			if t.peek() == pOn {
				if s := t.pop(); s != pOn {
					return t.errorUnexpectedWithExpected(s, pOn)
				}

				//now we have an unknown number of condition components, terminated by a semicolon
				for {
					s := t.pop()
					if s == "" {
						return t.error(ErrUnexpectedEOF)
					}
					if s == pSemicolon {
						break
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

			//save the transition
			fb.BasicFB.AddTransition(name, destState, strings.Join(condComponents, " "), debug)
		}
	}

	//everything is parsed, add it to the state machine
	fb.BasicFB.AddState(name, append(emits, runs...), debug)

	return nil
}

//parseBFBAlgorithm parses a single algorithm and adds it to fb identified by fbIndex
func (t *tfbParse) parseBFBAlgorithm(fbIndex int) *ParseError {
	fb := &t.fbs[fbIndex]

	//next word is algorithm name
	name := t.pop()
	debug := t.getCurrentDebugInfo()

	//next word should be "in"
	if s := t.pop(); s != pIn {
		return t.errorUnexpectedWithExpected(s, pIn)
	}

	lang, prog, err := t.getAlgorithmParts()
	if err != nil {
		return err
	}

	//next item should be semicolon
	if s := t.pop(); s != pSemicolon {
		return t.errorUnexpectedWithExpected(s, pSemicolon)
	}

	//we're done here, add the algorithm to the block
	fb.BasicFB.AddAlgorithm(name, lang, prog, debug)
	return nil
}

func (t *tfbParse) getAlgorithmParts() (string, string, *ParseError) {
	//next item is the language surrounded by single quotes
	lang := t.pop()
	if len(lang) == 0 {
		return "", "", t.error(ErrUnexpectedEOF)
	}
	if lang[0] != '"' || lang[len(lang)-1] != '"' {
		return "", "", t.errorWithReason(ErrUnexpectedValue, "Language argument should be surrounded by double quotes")
	}
	lang = strings.Trim(lang, "\"")

	//next item should be the program surrounded by backticks
	prog := t.pop()
	if len(prog) == 0 {
		return "", "", t.error(ErrUnexpectedEOF)
	}
	if prog[0] != '`' || prog[len(prog)-1] != '`' {
		return "", "", t.errorWithReason(ErrUnexpectedValue, "Language program should be surrounded by backticks")
	}
	prog = strings.Trim(prog, "`")

	return lang, prog, nil
}
