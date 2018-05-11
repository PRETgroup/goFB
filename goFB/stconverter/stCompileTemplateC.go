package stconverter

import (
	"strings"
)

const cTemplate = `
{{define "expression"}}{{$value := .HasValue}}{{$operator := .HasOperator}}{{/*
	*/}}{{if $value}}{{/*
		*/}}{{if isKnownVar $value}}me->{{end}}{{$value}}{{/*
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
		*/}}{{if $i}} else {{end}}if({{template "expression" $ifThen.IfExpression}}) {
			{{compileSequence $ifThen.ThenSequence}}
		}{{end}}{{/*
	*/}}{{if .ElseSequence}} else {
		{{compileSequence .ElseSequence}}
	}{{end}}{{/*
*/}}{{end}}

{{define "switchcase"}}{{/*
	*/}}switch({{template "expression" .SwitchOn}}) { {{range $ci, $case := .Cases}}
	{{range $cii, $casev := $case.CaseValues}}case {{$casev}}:
	{{end}}	{{compileSequence $case.Sequence}}
		break;{{end}}{{/*we need a break because ST case is not fall-through*/}}
	{{if .ElseSequence}}default:
		{{compileSequence .ElseSequence}}
		break;
	{{end}}
	} {{/*
*/}}{{end}}

{{define "forloop"}}{{$fc := .FindCounterName}}{{/*
	*/}}for({{if .ForAssignment}}{{template "expression" .ForAssignment}}{{end}}; {{/*
		*/}}{{if .ToValue}}{{if $fc}}{{if isKnownVar $fc}}me->{{end}}{{$fc}} <= {{end}}{{template "expression" .ToValue}}{{end}}; {{/*
		*/}}{{if .ByIncrement}}{{if $fc}}{{if isKnownVar $fc}}me->{{end}}{{$fc}} += {{end}}{{template "expression" .ByIncrement}}{{else}}{{if $fc}}{{$fc}}++{{end}}{{end}}) { 
		{{compileSequence .Sequence}}
	} {{/*
*/}}{{end}}

{{define "whileloop"}}{{/*
	*/}}while({{template "expression" .WhileExpression}}) {
		{{compileSequence .Sequence}}
	} {{/*
*/}}{{end}}

{{define "repeatloop"}}{{/*
	*/}}do {
		{{compileSequence .Sequence}}
	} while({{if .UntilExpression}}!({{template "expression" .UntilExpression}}){{else}}1{{end}}); {{/*
*/}}{{end}}`

func cTokenIsFunctionCall(token string) bool {
	first := strings.Index(token, "<")
	if len(token) > 2 && token[len(token)-1:] == ">" && first != -1 {
		return true
	}
	return false
}

func cTranslateOperatorToken(token string) string {
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
		return "!"
	case stNegative:
		return "-"
	case stExponentiation:
		//todo: we need to roll a custom exponentiation function
		panic("exponentiation not supported in C!")
		return ""
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
		return "="
	}
	//still here? panic
	panic("unsupported token " + token)
	return ""
}
