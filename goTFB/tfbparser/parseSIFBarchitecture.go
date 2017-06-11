package tfbparser

import "strings"

//parseSIFBarchitecture shall only be called once we have already parsed the
// "architecture of [blockname]" part of the definition
// so, we are up to the brace
func (t *tfbParse) parseSIFBarchitecture(fbIndex int) *ParseError {
	if s := t.pop(); s != pOpenBrace {
		return t.errorUnexpectedWithExpected(s, pOpenBrace)
	}

	//the next tag must be the pIn, which preceeds the language
	if s := t.pop(); s != pIn {
		return t.errorUnexpectedWithExpected(s, pIn)
	}

	lang := t.pop()
	lang = strings.Trim(lang, "\"")

	//the next tag must be the semicolon
	if s := t.pop(); s != pSemicolon {
		return t.errorUnexpectedWithExpected(s, pSemicolon)
	}

	debug := t.getCurrentDebugInfo()

	//we now have several things that could be in here
	//pre_init | init | run | shutdown

	//unlike in an interface, the various things that are in an architecture can be presented out of order
	inStruct := ""
	preInit := ""
	init := ""
	run := ""
	shutdown := ""

	for {
		s := t.pop()
		if s == "" {
			return t.error(ErrUnexpectedEOF)
		} else if s == pCloseBrace {
			//this is the end of the architecture
			break
		} else if s == pInStruct {
			inStruct = t.pop()
			if inStruct[0] != '`' || inStruct[len(inStruct)-1] != '`' {
				return t.errorWithReason(ErrUnexpectedValue, "Language program should be surrounded by backticks")
			}
			inStruct = strings.Trim(inStruct, "`")
			if s := t.pop(); s != pSemicolon {
				return t.errorUnexpectedWithExpected(s, pSemicolon)
			}
		} else if s == pPreInit {
			preInit = t.pop()
			if preInit[0] != '`' || preInit[len(preInit)-1] != '`' {
				return t.errorWithReason(ErrUnexpectedValue, "Language program should be surrounded by backticks")
			}
			preInit = strings.Trim(preInit, "`")
			if s := t.pop(); s != pSemicolon {
				return t.errorUnexpectedWithExpected(s, pSemicolon)
			}
		} else if s == pInit {
			init = t.pop()
			if init[0] != '`' || init[len(init)-1] != '`' {
				return t.errorWithReason(ErrUnexpectedValue, "Language program should be surrounded by backticks")
			}
			init = strings.Trim(init, "`")
			if s := t.pop(); s != pSemicolon {
				return t.errorUnexpectedWithExpected(s, pSemicolon)
			}
		} else if s == pRun {
			run = t.pop()
			if run[0] != '`' || run[len(run)-1] != '`' {
				return t.errorWithReason(ErrUnexpectedValue, "Language program should be surrounded by backticks")
			}
			run = strings.Trim(run, "`")
			if s := t.pop(); s != pSemicolon {
				return t.errorUnexpectedWithExpected(s, pSemicolon)
			}
		} else if s == pShutdown {
			shutdown = t.pop()
			if shutdown[0] != '`' || shutdown[len(shutdown)-1] != '`' {
				return t.errorWithReason(ErrUnexpectedValue, "Language program should be surrounded by backticks")
			}
			shutdown = strings.Trim(shutdown, "`")
			if s := t.pop(); s != pSemicolon {
				return t.errorUnexpectedWithExpected(s, pSemicolon)
			}
		} else {
			return t.error(ErrUnexpectedValue)
		}
	}

	t.fbs[fbIndex].ServiceFB.AddParams(lang, inStruct, preInit, init, run, shutdown, debug)

	return nil
}
