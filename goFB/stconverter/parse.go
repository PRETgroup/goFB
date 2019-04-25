package stconverter

import (
	"strings"
	"text/scanner"
)

const (
	stNewline   = "\n"
	stSemicolon = ";"

	stOpenBracket  = "("
	stCloseBracket = ")"

	stAssignment = ":="

	//stOpValue              = ""
	stNot                = "not"
	stExclamationNot     = "!"
	stNegative           = "`" //this is not actually in code, but we convert "negation" operators to it (i.e. "3 + -4 = -1" would become 3 + `4 = `1)
	stExponentiation     = "**"
	stMultiply           = "*"
	stDivide             = "/"
	stModulo             = "MOD"
	stAdd                = "+"
	stSubtract           = "-"
	stLessThan           = "<"
	stGreaterThan        = ">"
	stLessThanEqualTo    = "<="
	stGreaterThanEqualTo = ">="
	stEqual              = "="
	stInequal            = "<>"
	stAnd                = "and"
	stExlusiveOr         = "xor"
	stOr                 = "or"

	stExit   = "exit"
	stReturn = "return"

	stIf    = "if" //block beginner
	stThen  = "then"
	stElsif = "elsif"
	stElse  = "else"
	stEndIf = "end_if"

	stCase    = "case" //block beginner
	stOf      = "of"
	stComma   = ","
	stColon   = ":"
	stEndCase = "end_case"

	stFor    = "for" //block beginner
	stTo     = "to"
	stBy     = "by"
	stDo     = "do"
	stEndFor = "end_for"

	stWhile = "while" //block beginner
	//stWhile also uses stDo
	stEndWhile = "end_while"

	stRepeat    = "repeat" //block beginner
	stUntil     = "until"
	stEndRepeat = "end_repeat"
)

//ParseString takes an input string (i.e. filename) and input and returns all ST instructions in that string
func ParseString(name string, input string) ([]STInstruction, *STParseError) {
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

	for i := 0; i < len(items)-1; i++ {
		if items[i] == ":" && items[i+1] == "=" {
			items[i] = ":="
			items = append(items[:i+1], items[i+2:]...)
		}

		if items[i] == ">" && items[i+1] == "=" {
			items[i] = ">="
			items = append(items[:i+1], items[i+2:]...)
		}

		if items[i] == "<" && items[i+1] == "=" {
			items[i] = "<="
			items = append(items[:i+1], items[i+2:]...)
		}

		if items[i] == "<" && items[i+1] == ">" {
			items[i] = "<>"
			items = append(items[:i+1], items[i+2:]...)
		}

		if items[i] == stExclamationNot {
			items[i] = stNot
		}

		if items[i] == "|" && items[i+1] == "|" {
			items[i] = stOr
			items = append(items[:i+1], items[i+2:]...)
		}

		if items[i] == "&" && items[i+1] == "&" {
			items[i] = stAnd
			items = append(items[:i+1], items[i+2:]...)
		}

		if items[i] == "*" && items[i+1] == "*" {
			items[i] = stExponentiation
			items = append(items[:i+1], items[i+2:]...)
		}

		if items[i] == "AND" {
			items[i] = "and"
		}

		if items[i] == "OR" {
			items[i] = "or"
		}

		if items[i] == "XOR" {
			items[i] = "xor"
		}

		if items[i] == "'" && len(items) > i+2 && items[i+2] == "'" { //single-quoted constants become a single term
			items[i] = "'" + items[i+1] + "'"
			items = append(items[:i+1], items[i+3:]...)
		}
	}

	return items
}

//parseItems creates and runs a tfbparse struct
func parseItems(name string, items []string) ([]STInstruction, *STParseError) {
	t := stParse{items: items, currentLine: 1, currentFile: name}

	for !t.done() {
		s := t.peek()
		if s == "" { //don't do anything with empty strings
			t.pop()
			continue
		}
		//continue parsing...
		instr, err := t.parseNext()
		if err != nil {
			return nil, err
		}
		t.instructions = append(t.instructions, instr)
	}

	return t.instructions, nil
}
