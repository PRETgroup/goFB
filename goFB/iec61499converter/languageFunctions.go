package iec61499converter

import (
	"strings"
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

func strToUpper(s string) string {
	return strings.ToUpper(s)
}
