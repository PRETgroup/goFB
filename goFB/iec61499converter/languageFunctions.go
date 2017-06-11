package iec61499converter

import (
	"strings"

	"github.com/kiwih/goFB/iec61499"
)

//connChildSourceOnly is used in templates for getting rid of prefix stuff on connections
func connChildSourceOnly(in string) string {
	splitName := strings.Split(in, ".")
	return splitName[len(splitName)-1]
}

//connChildNameOnly is used in templates for getting rid of suffix stuff on connections
func connChildNameOnly(in string) string {
	splitName := strings.Split(in, ".")
	if len(splitName) != 2 {
		return ""
	}
	return splitName[0]
}

//connChildNameMatches is used in templates for location matching
func connChildNameMatches(in string, name string) bool {
	return strings.HasPrefix(in, name+".")
}

//used to check if an iec61499.Connection's .Source or .Destination (send in appropriate string) are going to a parent's port
func connIsOnParent(connName string) bool {
	return !strings.Contains(connName, ".")
}

func div(a int, b int) int {
	return a / b
}

func add(a int, b int) int {
	return a + b
}

func mod(a int, b int) int {
	return a % b
}

func count(a int) []int {
	b := make([]int, a)
	for i := 0; i < a; i++ {
		b[i] = i
	}
	return b
}

func findBlockDefinitionForType(bs []iec61499.FB, t string) *iec61499.FB {
	for _, b := range bs {
		if b.Name == t {
			return &b
		}
	}
	return nil
}

func strToUpper(s string) string {
	return strings.ToUpper(s)
}

func findVarDefinitionForName(b iec61499.FB, n string) *iec61499.Variable {
	if b.InputVars != nil {
		for _, varD := range b.InputVars {
			if varD.Name == n {
				return &varD
			}
		}
	}

	if b.OutputVars != nil {
		for _, varD := range b.OutputVars {
			if varD.Name == n {
				return &varD
			}
		}
	}

	return nil
}
