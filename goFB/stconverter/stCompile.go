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

var vhdlTemplateFuncMap = template.FuncMap{
	"translateOperatorToken": vhdlTranslateOperatorToken,
	"tokenIsFunctionCall":    vhdlTokenIsFunctionCall,
	"compileSequence":        VhdlCompileSequence,
	"isKnownVar":             isKnownVar,
	"reverseArgs":            reverseArgs,
}

var stTemplateFuncMap = template.FuncMap{
	"translateOperatorToken": stTranslateOperatorToken,
	"tokenIsFunctionCall":    stTokenIsFunctionCall,
	//	"compileSequence":        STCompileSequence,
	"gte":         gte,
	"isKnownVar":  isKnownVar,
	"reverseArgs": reverseArgs,
}

var (
	cTemplates    *template.Template
	stTemplates   *template.Template
	vhdlTemplates *template.Template
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
	stTemplates = template.Must(template.New("").Funcs(stTemplateFuncMap).Parse(stTemplate))
	vhdlTemplates = template.Must(template.New("").Funcs(vhdlTemplateFuncMap).Parse(vhdlTemplate))
}

func gte(a, b int) bool {
	return a >= b
}

//SetKnownVarNames sets the names of known variables for the compiler
func SetKnownVarNames(varNames []string) {
	knownVarNames = varNames
}

func reverseArgs(s []STExpression) []STExpression {
	r := make([]STExpression, len(s))
	for i := 0; i < len(s); i++ {
		r[i] = s[len(s)-1-i]
	}
	return r
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
			_, err := output.WriteString(CCompileExpression(inst)) //we have a special function for CCompileExpression because we might want to call it separately for 61499 guards
			panicOnErr(err)
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

//VhdlCompileSequence will take a sequence of STInstructions and compile them to their equivalent VHDL codes using the
//	vhdl templates stored in vhdlTemplates
func VhdlCompileSequence(sequence []STInstruction) string {
	output := &bytes.Buffer{}
	for _, untypedInst := range sequence {
		switch inst := untypedInst.(type) {
		case STExpression:
			_, err := output.WriteString(VhdlCompileExpression(inst)) //we have a special function for CCompileExpression because we might want to call it separately for 61499 guards
			panicOnErr(err)
			panicOnErr(output.WriteByte(';'))
			panicOnErr(output.WriteByte('\n'))
		case STIfElsIfElse:
			panicOnErr(vhdlTemplates.ExecuteTemplate(output, "ifelsifelse", inst))
		case STSwitchCase:
			panicOnErr(vhdlTemplates.ExecuteTemplate(output, "switchcase", inst))
		case STForLoop:
			panicOnErr(vhdlTemplates.ExecuteTemplate(output, "forloop", inst))
		case STWhileLoop:
			panicOnErr(vhdlTemplates.ExecuteTemplate(output, "whileloop", inst))
		case STRepeatLoop:
			panicOnErr(vhdlTemplates.ExecuteTemplate(output, "repeatloop", inst))
		}
	}
	return output.String()
}

//CCompileExpression will compile an STExpression to its equivalent C codes using the
//	c templates stored in cTemplates
func CCompileExpression(expr STExpression) string {
	output := &bytes.Buffer{}
	panicOnErr(cTemplates.ExecuteTemplate(output, "expression", expr))

	return output.String()
}

//VhdlCompileExpression will compile an STExpression to its equivalent VHDL codes using the
//	vhdl templates stored in vhdlTemplates
func VhdlCompileExpression(expr STExpression) string {
	output := &bytes.Buffer{}
	panicOnErr(vhdlTemplates.ExecuteTemplate(output, "expression", expr))

	return output.String()
}

//STCompileExpression will compile an STExpression to its equivalent C codes using the
//	c templates stored in cTemplates
func STCompileExpression(expr STExpression) string {
	output := &bytes.Buffer{}
	panicOnErr(stTemplates.ExecuteTemplate(output, "expression", expr))

	return output.String()
}
