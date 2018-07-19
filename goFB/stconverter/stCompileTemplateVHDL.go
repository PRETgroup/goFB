package stconverter

import (
	"strings"
)

//vhdlTemplate templates make the following assumptions:
//1) All variables are VHDL "Variables", not "signals"
//2) All variables are integer types
//3) Everything completes in a single cycle
//4) loops aren't yet supported
const vhdlTemplate = `
{{define "expression"}}{{$value := .HasValue}}{{$operator := .HasOperator}}{{/*
	*/}}{{if $value}}{{/*
		*/}}{{if isKnownVar $value}}{{$value}}{{else}}{{$value}}{{end}}{{/*
	*/}}{{else if $operator}}{{/* //then we need to determine how to print this operator
		*/}}{{if $operator.LeftAssociative}}{{/* //print first argument, operator string, then remaining arguments
			*/}}{{$args := reverseArgs .GetArguments}}{{$a := index $args 0}}{{$b := index $args 1}}{{$curPrec := $operator.GetPrecedence}}{{$aop := $a.HasOperator}}{{$bop := $b.HasOperator}}{{/*
			*/}}{{if $aop}}{{if gt $aop.GetPrecedence $curPrec}}({{end}}{{end}}{{template "expression" $a}}{{if $aop}}{{if gt $aop.GetPrecedence $curPrec}}){{end}}{{end}} {{translateOperatorToken $operator.GetToken}} {{if $bop}}{{if gt $bop.GetPrecedence $curPrec}}({{end}}{{end}}{{template "expression" $b}}{{if $bop}}{{if gt $bop.GetPrecedence $curPrec}}){{end}}{{end}}{{/*
		*/}}{{else}}{{/* //print name, opening bracket, then arguments separated by commas
			*/}}{{translateOperatorToken $operator.GetToken}}{{/*
			*/}}{{range $ind, $arg := reverseArgs .GetArguments}}{{if $ind}}, {{end}}{{template "expression" $arg}}{{end}}{{if tokenIsFunctionCall $operator.GetToken}}){{end}}{{/*
		*/}}{{end}}{{/*
	*/}}{{end}}{{/*}}
*/}}{{end}}

{{define "ifelsifelse"}}{{/*
	*/}}{{range $i, $ifThen := .IfThens}}{{/* //cycle through all the ifThens
		*/}}{{if $i}} els{{end}}if ({{template "expression" $ifThen.IfExpression}}) then
			{{compileSequence $ifThen.ThenSequence}}
		{{end}}{{/*
	*/}}{{if .ElseSequence}} else 
		{{compileSequence .ElseSequence}}
	{{end}}end if;{{/*
*/}}{{end}}

{{define "switchcase"}}{{/*
	*/}}case ({{template "expression" .SwitchOn}}) is {{range $ci, $case := .Cases}}
	when {{range $cii, $casev := $case.CaseValues}}{{if $cii}} | {{end}}{{$casev}}{{end}} =>	
		{{compileSequence $case.Sequence}}
	{{end}}
	{{if .ElseSequence}}when others =>
		{{compileSequence .ElseSequence}}
	{{end}}
	end case; {{/*
*/}}{{end}}

{{define "forloop"}} --for loops not supported in VHDL
{{end}}

{{define "whileloop"}} --while loops not supported in VHDL
{{end}}

{{define "repeatloop"}} --repeat loops not supported in VHDL
{{end}}`

func vhdlTokenIsFunctionCall(token string) bool {
	if token == "not" {
		return true
	}
	first := strings.Index(token, "<")
	if len(token) > 2 && token[len(token)-1:] == ">" && first != -1 {
		return true
	}
	return false
}

func vhdlTranslateOperatorToken(token string) string {
	//is it our function<n> syntax?
	first := strings.Index(token, "<")
	if len(token) > 2 && token[len(token)-1:] == ">" && first != -1 {
		//we on't need the following... we don't need the number of arguments??
		// ops := token[first+1 : len(token)-1]
		// opsInt, err := strconv.Atoi(ops)
		// if err != nil {
		// 	return token[:first-1] + "("
		// }
		return token[:first] + "("
	}
	//ok, not a function, so it's one of the st Operators
	switch token {
	case stExit:
		return "break"
	case stReturn:
		return "return"
	case stNot:
		return "not("
	case stNegative:
		return "-"
	case stExponentiation:
		//todo: we need to roll a custom exponentiation function
		panic("exponentiation not supported in VHDL!")
		return ""
	case stMultiply:
		return "*"
	case stDivide:
		return "/"
	case stModulo:
		return "%"
	case stAdd:
		return "+"
	case stSubtract:
		return "-"
	case stLessThan:
		return "<"
	case stGreaterThan:
		return ">"
	case stLessThanEqualTo:
		return "<="
	case stGreaterThanEqualTo:
		return ">="
	case stEqual:
		return "="
	case stInequal:
		return "/="
	case stAnd:
		return "and"
	case stExlusiveOr:
		return "xor"
	case stOr:
		return "or"
	case stAssignment:
		return ":="
	}
	//still here? panic
	panic("unsupported token " + token)
	return ""
}
