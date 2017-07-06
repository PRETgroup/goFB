package tfbparser

import (
	"strconv"
	"strings"

	"github.com/kiwih/goFB/iec61499"
)

//parseHFBarchitecture shall only be called once we have already parsed the
// "architecture of [blockname]" part of the definition
// so, we are up to the brace
func (t *tfbParse) parseHFBarchitecture(fbIndex int) *ParseError {
	if s := t.pop(); s != pOpenBrace {
		return t.errorUnexpectedWithExpected(s, pOpenBrace)
	}
	//we now have several things that could be in here
	//internal | internals | location | locations | algorithm | algorithms | closeBrace

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
			//we're cheeky here, and take advantage of the fact that BFB and HFB internals are identical
			//the function parseBFBInternal can parse HFB internals as well
			if err := t.parsePossibleArrayInto(fbIndex, (*tfbParse).parseBFBInternal); err != nil {
				return err
			}
		} else if s == pAlgorithm || s == pAlgorithms {
			if err := t.parsePossibleArrayInto(fbIndex, (*tfbParse).parseHFBAlgorithm); err != nil {
				return err
			}
		} else if s == pLocation || s == pLocations {
			if err := t.parsePossibleArrayInto(fbIndex, (*tfbParse).parseHFBLocation); err != nil {
				return err
			}
		}
	}

	return nil
}

//parseHFBLocation parses a single location and adds it to fb identified by fbIndex
// remember that HFBs will be translated into BFBs as soon as they are parsed, hence
// most things in this function are validated later in the iec61499 package
func (t *tfbParse) parseHFBLocation(fbIndex int) *ParseError {
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
	var invariants []iec61499.HFBInvariant

	for {
		s := t.pop()
		if s == "" {
			return t.error(ErrUnexpectedEOF)
		}
		if s == pCloseBrace {
			break
		}
		if s == pInvariant {
			//capture any number of comma-separated invariants (invariants are surrounded by backticks)
			//should terminate with semicolon
			for {
				inv := t.pop()

				if len(inv) == 0 {
					return t.error(ErrUnexpectedEOF)
				}

				if inv[0] == '`' && inv[len(inv)-1] == '`' {
					inv = strings.Trim(inv, "`")
				} else {
					return t.errorWithReason(ErrUnexpectedValue, "Invariants should be surrounded by backticks")
				}

				invariants = append(invariants, iec61499.HFBInvariant{Invariant: inv, DebugInfo: t.getCurrentDebugInfo()})

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
		if s == pRun {
			//capture any number of comma-separated algorithms to run
			for {
				r := t.pop()

				if len(r) == 0 {
					return t.error(ErrUnexpectedEOF)
				}

				if r[0] == '`' && r[len(r)-1] == '`' {
					//we have a program, an anonymous algorithm here (identified by backtick surrounds)
					prog := strings.Trim(r, "`")
					anonAlgName := name + "_alg" + strconv.Itoa(anonAlgIndex)
					anonAlgIndex++

					//we're done here, add the algorithm to the block
					fb.HybridFB.AddAlgorithm(anonAlgName, prog, debug)

					//replace program with the name of the new anon algorithm so it can be joined properly
					r = anonAlgName

				} else if r[0] == '`' || r[len(r)-1] == '`' {
					//only one side of the program has a backtick
					return t.errorWithReason(ErrUnexpectedValue, "Language program should be surrounded by backticks")
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
					condComponents = append(condComponents, s)

				}
				//TODO: emitting/running things on transitions?!
			}
			if len(condComponents) == 0 { //put in a default condition if no condition exists
				condComponents = append(condComponents, "true")
			}

			//save the transition
			fb.BasicFB.AddTransition(name, destState, strings.Join(condComponents, " "), debug)
		}
	}

	//everything is parsed, add it to the state machine
	fb.HybridFB.AddLocation(name, invariants, append(emits, runs...), debug)

	return nil
}

//parseHFBAlgorithm parses a single algorithm and adds it to fb identified by fbIndex
//Notice that HFB algorithms differ to BFB algorithms and don't allow specification of language
func (t *tfbParse) parseHFBAlgorithm(fbIndex int) *ParseError {
	fb := &t.fbs[fbIndex]

	//next word is algorithm name
	name := t.pop()
	debug := t.getCurrentDebugInfo()

	//next item should be the program surrounded by backticks
	prog := t.pop()
	if len(prog) == 0 {
		return t.error(ErrUnexpectedEOF)
	}
	if prog[0] != '`' || prog[len(prog)-1] != '`' {
		return t.errorWithReason(ErrUnexpectedValue, "Language program should be surrounded by backticks")
	}
	prog = strings.Trim(prog, "`")

	//next item should be semicolon
	if s := t.pop(); s != pSemicolon {
		return t.errorUnexpectedWithExpected(s, pSemicolon)
	}

	//we're done here, add the algorithm to the block
	fb.HybridFB.AddAlgorithm(name, prog, debug)
	return nil
}
