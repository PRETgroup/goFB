package stconverter

import (
	"bytes"
	"text/template"
)

var cTemplateFuncMap = template.FuncMap{
	"translateOperatorToken": cTranslateOperatorToken,
	"tokenIsFunctionCall":    cTokenIsFunctionCall,
	"compileSequence":        CCompileSequence,
	"isKnownVar":             isKnownVar,
	"reverseArgs":            reverseArgs,
}

var (
	cTemplates    *template.Template
	knownVarNames []string
)

func panicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}

//init runs to initialise the package
func init() {
	cTemplates = template.Must(template.New("").Funcs(cTemplateFuncMap).Parse(cTemplate))
}

//SetKnownVarNames sets the names of known variables for the compiler
func SetKnownVarNames(varNames []string) {
	knownVarNames = varNames
}

func reverseArgs(s []STExpression) []STExpression {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func isKnownVar(name string) bool {
	for _, n := range knownVarNames {
		if n == name {
			return true
		}
	}
	return false
}

//CCompileSequence will take a sequence of STInstructions and compile them to their equivalent C codes using the
//	c templates stored in cTemplates
func CCompileSequence(sequence []STInstruction) string {
	output := &bytes.Buffer{}
	for _, untypedInst := range sequence {
		switch inst := untypedInst.(type) {
		case STExpression:
			panicOnErr(cTemplates.ExecuteTemplate(output, "expression", inst))
			panicOnErr(output.WriteByte(';'))
			panicOnErr(output.WriteByte('\n'))
		case STIfElsIfElse:
			panicOnErr(cTemplates.ExecuteTemplate(output, "ifelsifelse", inst))
		case STSwitchCase:
			panicOnErr(cTemplates.ExecuteTemplate(output, "switchcase", inst))
		case STForLoop:
			panicOnErr(cTemplates.ExecuteTemplate(output, "forloop", inst))
		case STWhileLoop:
			panicOnErr(cTemplates.ExecuteTemplate(output, "whileloop", inst))
		case STRepeatLoop:
			panicOnErr(cTemplates.ExecuteTemplate(output, "repeatloop", inst))
		}
	}
	return output.String()
}
