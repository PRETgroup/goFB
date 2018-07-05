package stconverter

import (
	"strings"
)

//verilogTemplate templates make the following assumptions:
//1) All variables are VHDL "Variables", not "signals"
//2) All variables are integer types
//3) Everything completes in a single cycle
//4) loops aren't yet supported
const verilogTemplate = `
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
		*/}}{{if $i}} else {{end}}if ({{template "expression" $ifThen.IfExpression}}) begin
			{{compileSequence $ifThen.ThenSequence}}
		end{{end}}{{/*
	*/}}{{if .ElseSequence}} else begin 
		{{compileSequence .ElseSequence}}
	end{{end}}{{/*
*/}}{{end}}

{{define "switchcase"}}{{/*
	*/}}case({{template "expression" .SwitchOn}}) {{range $ci, $case := .Cases}}
	{{range $cii, $casev := $case.CaseValues}}{{if $cii}}, {{end}}{{$casev}}{{end}}: begin	
		{{compileSequence $case.Sequence}}
	end{{end}}
	{{if .ElseSequence}}default: begin
		{{compileSequence .ElseSequence}}
	end{{end}}
	endcase {{/*
*/}}{{end}}

{{define "forloop"}} --for loops not supported in VHDL
{{end}}

{{define "whileloop"}} --while loops not supported in VHDL
{{end}}

{{define "repeatloop"}} --repeat loops not supported in VHDL
{{end}}`

func verilogTokenIsFunctionCall(token string) bool {
	first := strings.Index(token, "<")
	if len(token) > 2 && token[len(token)-1:] == ">" && first != -1 {
		return true
	}
	return false
}

func verilogTranslateOperatorToken(token string) string {
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
		return "~"
	case stNegative:
		return "-"
	case stExponentiation:
		return "**"
	case stMultiply:
		return "*"
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
		return "=="
	case stInequal:
		return "!="
	case stAnd:
		return "&&"
	case stExlusiveOr:
		return "^"
	case stOr:
		return "||"
	case stAssignment:
		return "<="
	}
	//still here? panic
	panic("unsupported token " + token)
	return ""
}
