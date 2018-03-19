package tfbparser

import (
	"errors"
	"strings"
	"text/scanner"

	"github.com/PRETgroup/goFB/iec61499"
)

const (
	pNewline = "\n"

	pBasicFB     = "basicFB"
	pCompositeFB = "compositeFB"
	pServiceFB   = "serviceFB"
	pHybridFB    = "hybridFB"

	pCompilerInfoHeader = "compileheader"

	pArbitrary = "arbitrary" //used in SIFB autogeneration for goFB
	pInStruct  = "in_struct"
	pPreInit   = "pre_init"
	pInit      = "init"
	pShutdown  = "shutdown"

	pOpenBrace    = "{"
	pCloseBrace   = "}"
	pOpenBracket  = "["
	pCloseBracket = "]"
	pComma        = ","
	pSemicolon    = ";"
	pInitial      = "initial"
	pInitEq       = ":="

	pFBinterface    = "interface"
	pFBarchitecture = "architecture"
	pOf             = "of"

	pIn  = "in"
	pOut = "out"

	pWith = "with"

	pRun   = "run"
	pEmit  = "emit"
	pTrans = "->"
	pOn    = "on"

	pInvariant = "invariant"

	pInternal   = "internal"
	pInternals  = "internals"
	pState      = "state"
	pStates     = "states"
	pAlgorithm  = "algorithm"
	pAlgorithms = "algorithms"

	pLocation  = "location"
	pLocations = "locations"

	pInstance  = "instance"
	pInstances = "instances"
	pEvent     = "event"
	pEvents    = "events"
	pDatum     = "datum"
	pData      = "data"

	pConn   = "<-"
	pPeriod = "."
)

//ParseString takes an input string (i.e. filename) and input and returns all FBs in that string
func ParseString(name string, input string) ([]iec61499.FB, *ParseError) {
	//break up input string into all of its parts
	items := scanString(name, input)

	//now parse the items
	return parseItems(name, items)
}

func scanString(name string, input string) []string {
	var s scanner.Scanner

	s.Filename = name
	s.Init(strings.NewReader(input))

	//we don't want to ignore \n characters (we want to know what line we're on)
	s.Whitespace = 1<<'\t' | 0<<'\n' | 1<<'\r' | 1<<' '
	//we don't want scanner.ScanChars
	s.Mode = scanner.ScanIdents | scanner.ScanFloats | scanner.ScanStrings | scanner.ScanRawStrings | scanner.ScanComments | scanner.SkipComments

	//TODO: think about the scanner error function. Maybe we should provide one, that when an error occurs, halts scanning?

	var tok rune
	var items []string
	for tok != scanner.EOF {
		tok = s.Scan()
		items = append(items, s.TokenText())
	}

	//combine -> and <- and := operators

	for i := 0; i < len(items)-1; i++ {
		if items[i] == "<" && items[i+1] == "-" {
			items[i] = "<-"
			items = append(items[:i+1], items[i+2:]...)
		}

		if items[i] == "-" && items[i+1] == ">" {
			items[i] = "->"
			items = append(items[:i+1], items[i+2:]...)
		}

		if items[i] == ":" && items[i+1] == "=" {
			items[i] = ":="
			items = append(items[:i+1], items[i+2:]...)
		}

		if items[i] == "=" && items[i+1] == "=" {
			items[i] = "=="
			items = append(items[:i+1], items[i+2:]...)
		}

		if items[i] == "!" && items[i+1] == "=" {
			items[i] = "!="
			items = append(items[:i+1], items[i+2:]...)
		}

		if items[i] == "&" && items[i+1] == "&" {
			items[i] = "&&"
			items = append(items[:i+1], items[i+2:]...)
		}

		if items[i] == "|" && items[i+1] == "|" {
			items[i] = "||"
			items = append(items[:i+1], items[i+2:]...)
		}
	}

	return items
}

//parseItems creates and runs a tfbparse struct
func parseItems(name string, items []string) ([]iec61499.FB, *ParseError) {
	t := tfbParse{items: items, currentLine: 1, currentFile: name}

	for !t.done() {
		s := t.pop()
		if t.done() {
			break
		}
		//have we defined a basicFB or compositeFB
		if s == pBasicFB || s == pCompositeFB || s == pServiceFB {
			if err := t.parseFB(s); err != nil {
				return nil, err
			}
			continue
		}

		//is this defining an interface for an fb
		if s == pFBinterface {
			if err := t.parseFBinterface(); err != nil {
				return nil, err
			}
			continue
		}

		//is this defining an architecture for an fb
		if s == pFBarchitecture {
			if err := t.parseFBarchitecture(); err != nil {
				return nil, err
			}
			continue
		}
		return nil, t.errorWithArg(ErrUnexpectedValue, s)
	}

	return t.fbs, nil
}

//isValidType returns true if string s is one of the valid IEC61499 event/data types
func isValidType(s string) bool {
	if s == "event" ||
		s == "bool" ||
		s == "byte" ||
		s == "word" ||
		s == "dword" ||
		s == "lword" ||
		s == "sint" ||
		s == "usint" ||
		s == "int" ||
		s == "uint" ||
		s == "dint" ||
		s == "udint" ||
		s == "lint" ||
		s == "ulint" ||
		s == "real" ||
		s == "lreal" ||
		s == "time" ||
		s == "any" {
		return true
	}
	return false
}

//parseFB will create a new basicFB and add it to the list of internal FBs
func (t *tfbParse) parseFB(fbType string) *ParseError {
	var fbs []iec61499.FB
	for {
		name := t.pop()
		if !t.isFBNameUnused(name) {
			return t.errorWithArg(ErrNameAlreadyInUse, name)
		}

		if fbType == pBasicFB {
			fbs = append(fbs, *iec61499.NewBasicFB(name))
		} else if fbType == pCompositeFB {
			fbs = append(fbs, *iec61499.NewCompositeFB(name))
		} else if fbType == pServiceFB {
			fbs = append(fbs, *iec61499.NewServiceFB(name))
		} else if fbType == pHybridFB {
			fbs = append(fbs, *iec61499.NewHybridFB(name))
		} else {
			return t.errorWithReason(ErrInternal, "I can't parse fbType "+fbType)
		}

		if t.peek() == pComma {
			t.pop() //get rid of comma
			continue
		}
		break
	}

	s := t.pop()
	if s == pCompilerInfoHeader {
		header := t.pop()
		header = strings.Trim(header, "\"")
		for i := 0; i < len(fbs); i++ {
			fbs[i].CompilerInfo.Header = header
		}

		s = t.pop()
	}

	if s != pSemicolon {
		return t.errorUnexpectedWithExpected(s, pSemicolon+" or "+pCompilerInfoHeader)
	}

	t.fbs = append(t.fbs, fbs...)

	return nil
}

func (t *tfbParse) parseFBarchitecture() *ParseError {
	var s string
	//first word should be of
	s = t.pop()
	if s != pOf {
		return t.errorUnexpectedWithExpected(s, pOf)
	}

	//second word is fb name
	s = t.pop()
	fbIndex := t.getFBIndexFromName(s)
	if fbIndex == -1 {
		return t.errorWithArg(ErrUndefinedFB, s)
	}

	//detect type of FB and parse as appropriate
	if t.fbs[fbIndex].BasicFB != nil {
		return t.parseBFBarchitecture(fbIndex)
	}

	if t.fbs[fbIndex].CompositeFB != nil {
		return t.parseCFBarchitecture(fbIndex)
	}

	if t.fbs[fbIndex].ServiceFB != nil {
		return t.parseSIFBarchitecture(fbIndex)
	}

	if t.fbs[fbIndex].HybridFB != nil {
		return t.parseHFBarchitecture(fbIndex)
	}

	return t.error(errors.New("Can't parse unknown architecture type"))
}
