package tfbparser

import (
	"strings"
)

//networkConn is a connection between any event or data line in
// a cfb network
//we use it to store preliminary parsing until we validate once
//the whole network is in
type networkConn struct {
	line   int
	source string
	dest   string
}

//parseCFBarchitecture shall only be called once we have already parsed the
// "architecture of [blockname]" part of the definition
// so, we are up to the brace
func (t *tfbParse) parseCFBarchitecture(fbIndex int) *ParseError {
	if s := t.pop(); s != pOpenBrace {
		return t.errorUnexpectedWithExpected(s, pOpenBrace)
	}
	//we now have several things that could be in here
	//instance | instances | event | events | datum | data | closeBrace

	//unlike in an interface, the various things that are in an architecture can be presented out of order
	//this has consequences for both events and data, which we'll only verify
	//once we have encountered all instances and reached the closebracket
	//this is a very similar process to the one followed in parseBFBarchitecture

	for {
		s := t.pop()
		if s == "" {
			return t.error(ErrUnexpectedEOF)
		} else if s == pCloseBrace {
			//this is the end of the architecture
			break
		} else if s == pInstance || s == pInstances {
			if err := t.parsePossibleArrayInto(fbIndex, (*tfbParse).parseCFBInstance); err != nil {
				return err
			}
		} else if s == pEvent || s == pEvents {
			if err := t.parsePossibleArrayInto(fbIndex, (*tfbParse).parseCFBEventConnection); err != nil {
				return err
			}
		} else if s == pData || s == pDatum {
			if err := t.parsePossibleArrayInto(fbIndex, (*tfbParse).parseCFBDataConnection); err != nil {
				return err
			}
		}
	}

	return nil
}

//parseCFBInstance parses a single instance and adds it to fb identified by fbIndex
func (t *tfbParse) parseCFBInstance(fbIndex int) *ParseError {
	fb := &t.fbs[fbIndex]

	//next is the type being instantiated (will be name of other fb)
	fbTypeName := t.pop()
	debug := t.getCurrentDebugInfo()

	//next is the name of the instance of that block

	instNames := []string{}

	//this could be an array of names, so we'll loop while we are finding commas
	for {
		name := t.pop()
		instNames = append(instNames, name)
		if t.peek() == pComma {
			t.pop() //get rid of the pComma
			continue
		}
		break
	}

	//next should be a semicolon
	if s := t.pop(); s != pSemicolon {
		return t.errorUnexpectedWithExpected(s, pSemicolon)
	}

	fb.CompositeFB.AddInstances(fbTypeName, instNames, debug)
	return nil
}

//parseCFBEventConnection parses a single event connection between instances/parents
// and adds it to fb identified by fbIndex
func (t *tfbParse) parseCFBEventConnection(fbIndex int) *ParseError {
	fb := &t.fbs[fbIndex]

	//names are string.string or just string

	//next is destination
	dest := t.pop()
	debug := t.getCurrentDebugInfo()

	if t.peek() == pPeriod {
		dest += t.pop() //add period to dest
		dest += t.pop() //add next part of name to dest
	}

	//next should be pConn
	if s := t.pop(); s != pConn {
		return t.errorUnexpectedWithExpected(s, pConn)
	}

	//next is source, which could be comma separated (this is only true of events as these can have multiple sources)
	var sources []string
	for {
		s := t.pop()
		if s == "" {
			return t.error(ErrUnexpectedEOF)
		}
		if s == pSemicolon {
			break
		}
		if s == pComma {
			continue
		}
		source := s
		if t.peek() == pPeriod {
			source += t.pop() //add period to source
			source += t.pop() //add next part of name to source
		}
		sources = append(sources, source)
	}

	fb.CompositeFB.AddNetworkEventConns(sources, dest, debug)
	return nil
}

func (t *tfbParse) parseCFBDataConnection(fbIndex int) *ParseError {
	fb := &t.fbs[fbIndex]

	//names are string.string or just string
	dest := t.pop()
	debug := t.getCurrentDebugInfo()

	if t.peek() == pPeriod {
		dest += t.pop() //add period to dest
		dest += t.pop() //add next part of name to dest
	}

	//next should be pConn
	if s := t.pop(); s != pConn {
		return t.errorUnexpectedWithExpected(s, pConn)
	}

	//next is source, which could be surrounded by backticks if it is a parameter
	s := t.pop()
	if len(s) >= 0 && s[0] == '`' && s[len(s)-1] == '`' {
		//the source of this data is a constant, not a connection
		//this makes this a parameter, not a data connection
		param := strings.Trim(s, "`")
		destParts := strings.Split(dest, ".")
		if len(destParts) != 2 {
			return t.errorWithArg(ErrOnlyInstancesGetParameters, dest)
		}
		if _, err := fb.CompositeFB.AddNetworkParameter(param, destParts[0], destParts[1], debug); err != nil {
			return t.errorWithArg(err.Err, err.Arg)
		}

		//next must be semicolon
		if s := t.pop(); s != pSemicolon {
			return t.errorUnexpectedWithExpected(s, pSemicolon)
		}

		return nil
	}

	//not backticks, but it might have two parts
	source := s
	if t.peek() == pPeriod {
		source += t.pop() //add period to source
		source += t.pop() //add next part of name to source
	}

	//next must be semicolon
	if s := t.pop(); s != pSemicolon {
		return t.errorUnexpectedWithExpected(s, pSemicolon)
	}

	fb.CompositeFB.AddNetworkDataConn(source, dest, debug)
	return nil
}
