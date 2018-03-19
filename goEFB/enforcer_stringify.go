package main

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"text/template"
	"time"
)

//VFuncMap provides helpful functions to use in the templates
var VFuncMap = template.FuncMap{
	"curdatetime":      curdatetime,
	"underscoreString": underscoreString,
	"policyIO":         policyIO,
	"pair":             pair,
	"quad":             quad,
	"classifyVar":      classifyVar,
	"opString":         opString,
	"notOpString":      notOpString,
	"add":              add,
}

func curdatetime() string {
	return time.Now().Format(time.RFC3339)
}

func pair(x, y interface{}) interface{} {
	return struct{ First, Second interface{} }{x, y}
}

func quad(a, b, c, d interface{}) interface{} {
	return struct{ First, Second, Third, Forth interface{} }{a, b, c, d}
}

//underscoreString replaces spaces with underscores in a string
func underscoreString(s string) string {
	return strings.Replace(s, " ", "_", -1)
}

//policyIO is used to convert IO for policies into something usable in templates
func policyIO(e Enforcer, pIndex int) IO {
	inp, enf, _, err := e.GetIO(e.Policies[pIndex])
	if err != nil {
		fmt.Println("error:", err.Error())
		return IO{}
	}
	return IO{
		inp, enf,
	}
}

//classifyVar is used to classify Var for enforcers into something usable in templates
//if 0, it is unknown
//if 1, it is an input
//if 2, it is an enforce line
//if 3, it is a trigger name
func classifyVar(e Enforcer, pIndex int, v string) int {

	inp, enf, tmr, err := e.GetIO(e.Policies[pIndex])
	if err != nil {
		fmt.Println("error:", err.Error())
		return 0
	}
	for i := 0; i < len(inp); i++ {
		if inp[i].Name == v {
			return 1
		}
	}
	for i := 0; i < len(enf); i++ {
		if enf[i].Name == v {
			return 2
		}
	}
	for i := 0; i < len(tmr); i++ {
		if tmr[i] == v {
			return 3
		}
	}
	fmt.Println("error: couldn't find name'", v, "' in enforcer policy pIndex", pIndex)
	return 0
}

func notOpString(e Enforcer, p int, op Operation) string {
	return opStringWithNotAndTime(e, p, op, true, false)
}

func opString(e Enforcer, p int, op Operation) string {
	return opStringWithNotAndTime(e, p, op, false, false)
}

//opStringWithNotAndTime converts an operation with some metadata to a string for synthesis
//e : the Enforcer we're operating over
//p : the index of the Policy from which the operation draws from
//op: the Operation
//not: whether or not the Operation needs to have the keyword "not" in front of it (for use when checking requirements)
//inTime: whether or not a time clause has begun, as behaviour changes depending on this case
func opStringWithNotAndTime(e Enforcer, p int, op Operation, not bool, inTime bool) string {
	if op.IsValue() {
		return op.Value
	}
	if op.IsVariable() {
		varType := classifyVar(e, p, op.Value)
		if varType == 2 {
			return "q_enf." + op.Value
		}
		if varType == 3 {
			return "trigger_" + op.Value + "_time"
		}
		return op.Value
	}

	if !inTime {
		s := ""
		if not {
			s = "not"
		}
		if op.IsAfter() {
			if op.A != nil {
				return s + "(" + opStringWithNotAndTime(e, p, *op.A, not, inTime) + ") and (t < " + opStringWithNotAndTime(e, p, *op.B, not, true) + ")"
			}
			return "(t > " + opStringWithNotAndTime(e, p, *op.B, not, true) + ")"
		}
		if op.IsUntil() {
			return s + "(" + opStringWithNotAndTime(e, p, *op.A, not, inTime) + ") and (t < " + opStringWithNotAndTime(e, p, *op.B, not, true) + ")"
		}
		if op.IsBefore() {
			return s + "(" + opStringWithNotAndTime(e, p, *op.A, not, inTime) + ") and (t > " + opStringWithNotAndTime(e, p, *op.B, not, true) + ")"
		}
	} else {
		if op.IsAfter() {
			return "(" + opStringWithNotAndTime(e, p, *op.A, not, inTime) + " + " + opStringWithNotAndTime(e, p, *op.B, not, inTime) + ")"
		}
		if op.IsUntil() || op.IsBefore() {
			return "(syntax error: Only `after` can be used inside time clauses)"
		}
	}

	if op.IsSet() {
		return opStringWithNotAndTime(e, p, *op.A, not, inTime) + " := " + opStringWithNotAndTime(e, p, *op.B, not, inTime)
	}

	return "(" + opStringWithNotAndTime(e, p, *op.A, not, inTime) + " " + op.Type + " " + opStringWithNotAndTime(e, p, *op.B, not, inTime) + ")"
}

func add(a, b int) int {
	return a + b
}

//EnforcerString is a container for a stringified version of an enforcer in VHDL
type EnforcerString struct {
	Name     string
	Contents []byte
}

//Stringify converts an enforcer to an enforcerString (a VHDL representation of its policies etc)
func (e *Enforcer) Stringify(enforcerTpls *template.Template) ([]EnforcerString, error) {
	enforcerStrings := make([]EnforcerString, len(e.Policies)+2)

	enforcerTypeContents := &bytes.Buffer{}
	if err := enforcerTpls.ExecuteTemplate(enforcerTypeContents, "enforcerTypes", e); err != nil {
		return nil, err
	}

	enforcerStrings[0] = EnforcerString{
		Name:     "enforcement_types_" + underscoreString(e.Name) + ".vhdl",
		Contents: enforcerTypeContents.Bytes(),
	}

	enforcerTopContents := &bytes.Buffer{}
	if err := enforcerTpls.ExecuteTemplate(enforcerTopContents, "top", e); err != nil {
		return nil, err
	}

	enforcerStrings[1] = EnforcerString{
		Name:     "enforcement_top_" + underscoreString(e.Name) + ".vhdl",
		Contents: enforcerTopContents.Bytes(),
	}

	for i := 0; i < len(e.Policies); i++ {
		dat := struct {
			E        *Enforcer
			PolIndex int
		}{
			E:        e,
			PolIndex: i,
		}

		enforcerContents := &bytes.Buffer{}
		if err := enforcerTpls.ExecuteTemplate(enforcerContents, "enforcer", dat); err != nil {
			return nil, errors.New("(index " + strconv.Itoa(i) + "): " + err.Error())
		}
		enforcerStrings[i+2] = EnforcerString{
			//enforcer_{{underscoreString .E.Name}}_{{underscoreString $policy.Name}}
			Name:     "enforcer_" + underscoreString(e.Name) + "_" + underscoreString(e.Policies[i].Name) + ".vhdl",
			Contents: enforcerContents.Bytes(),
		}
		enforcerContents.Reset()
	}

	return enforcerStrings, nil
}
