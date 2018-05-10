package stconverter

import (
	"strings"
)

const stTemplate = `
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
`

func stTokenIsFunctionCall(token string) bool {
	first := strings.Index(token, "<")
	if len(token) > 2 && token[len(token)-1:] == ">" && first != -1 {
		return true
	}
	return false
}

func stTranslateOperatorToken(token string) string {
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
	if token == stNot {
		return token + " "
	}
	return token
}
