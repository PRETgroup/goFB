package iec61499converter

import (
	"regexp"
	"strings"

	"github.com/kiwih/goFB/iec61499converter/iec61499"
)

//getVhdlType returns the VHDL type to use with respect to an IEC61499 type
func getVhdlType(iec61499type string) string {
	vhdlType := ""
	switch strings.ToLower(iec61499type) {
	case "unsigned":
		vhdlType = "unsigned(31 downto 0)"
	case "int":
		vhdlType = "signed(31 downto 0)"
	case "float":
		panic("Float type not allowed in conversion")
	case "double":
		panic("Double type not allowed in conversion")
	case "bool":
		vhdlType = "std_logic"
	case "string":
		panic("String type not allowed in conversion")
	case "byte":
		vhdlType = "unsigned(7 downto 0)"
	case "any":
		panic("Any type not allowed in conversion")
	default:
		panic("Unknown IEC61499 type: " + iec61499type)
	}

	return vhdlType
}

//getVhdlECCTransitionCondition returns the VHDL "if" condition to use in state machine next state logic
func getVhdlECCTransitionCondition(iec61499trans string) string {
	re := regexp.MustCompile("([a-zA-Z_<>=]+)")
	retVal := iec61499trans
	retVal = strings.Replace(retVal, "!", "not ", -1)
	retVal = strings.Replace(retVal, "AND", "and", -1)
	retVal = strings.Replace(retVal, "OR", "or", -1)
	retVal = re.ReplaceAllStringFunc(retVal, addTrueCheck)
	return retVal
}

//addTrueCheck is used in conjunction with getVhdlECCTransitionCondition to format the ECC transition in a VHDL-friendly manner
//it is responsible for converting things such as "if variable and variable2" to "if variable = '1' and variable2 = '1'"
func addTrueCheck(in string) string {
	if strings.ToLower(in) == "and" || strings.ToLower(in) == "or" || strings.ToLower(in) == "not" || strings.ContainsAny(in, "<>=") || strings.ToLower(in) == "true" || strings.ToLower(in) == "false" {
		return in
	}

	return in + " = '1'"
}

//in our algorithms, DONE needs to be turned into the correct signal name
func renameDoneSignal(in string, name string) string {
	return strings.Replace(in, "DONE", name+"_alg_done", -1)
}

//renameConnSignal is used in templates to make a consistent and friendly name for the connections
func renameConnSignal(in string) string {
	return strings.Replace(in, ".", "_", -1) + "_conn" // + "_to_" + strings.Replace(c.Destination, ".", "_", -1)
}

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

//getCECCTransitionCondition returns the C "if" condition to use in state machine next state logic
func getCECCTransitionCondition(block iec61499.FB, iec61499trans string) string {
	re1 := regexp.MustCompile("([<>=!]+)")
	re2 := regexp.MustCompile("([a-zA-Z_<>=]+)")

	retVal := iec61499trans

	//rename AND and OR
	retVal = strings.Replace(retVal, "AND", "&&", -1)
	retVal = strings.Replace(retVal, "OR", "||", -1)

	//re1: add whitespace around operators
	retVal = re1.ReplaceAllStringFunc(retVal, func(in string) string {
		return " " + in + " "
	})

	//re2: add "me->" where appropriate
	retVal = re2.ReplaceAllStringFunc(retVal, func(in string) string {
		if strings.ToLower(in) == "and" || strings.ToLower(in) == "or" || strings.ContainsAny(in, "!><=") || strings.ToLower(in) == "true" || strings.ToLower(in) == "false" {
			//no need to make changes, these aren't variables or events
			return in
		}
		//check to see if it is an input event
		if block.EventInputs != nil {
			for _, event := range block.EventInputs.Events {
				if event.Name == in {
					return "me->inputEvents.event." + event.Name //FUTURE WORK: Consider use of pointers and offsets to minimise memory footprint for events? i.e. "*(me->inputEvents." + in + " + ev_offset)"
				}
			}
		}
		//else, return it assuming input data
		return "me->" + in
	})

	//tidy the whitespace
	retVal = strings.Replace(retVal, "  ", " ", -1)

	return retVal
}

//locationType = 0 (Var)
//locationType = 1 (destination)
//locationType = 2 (source)
//don't use this function, use one of the helper functions
func renameCLocation(in string, locationType int) string {
	if strings.Contains(in, ".") {
		//it comes from a child FB
		if locationType == 1 {
			return strings.Replace(in, ".", ".inputEvents.event.", 1) //events are in a sub struct
		} else if locationType == 2 {
			return strings.Replace(in, ".", ".outputEvents.event.", 1) //events are in a sub struct
		}

		return in

	}
	//it comes from the parent FB, which means input/output is swapped
	if locationType == 1 {
		return "outputEvents.event." + in
	}
	if locationType == 2 {
		return "inputEvents.event." + in
	}
	return in

}

func renameCEventDestinationLocation(in string) string {
	return renameCLocation(in, 1)
}

func renameCEventSourceLocation(in string) string {
	return renameCLocation(in, 2)
}

//This finds a data connection source based on a destination source (destinations can only ever have one source for data connections in iec61499)
func findSourceDataName(conns []iec61499.Connection, destChildName string, destVarName string) string {
	for _, conn := range conns {
		if destChildName != "" {
			if conn.Destination == destChildName+"."+destVarName {
				return renameCLocation(conn.Source, 0)
			}
		} else {
			if conn.Destination == destVarName {
				return renameCLocation(conn.Source, 0)
			}
		}
	}
	return "0"
}

//Ths finds event connection source(s) based on a destination source (events can have multiple sources for event connections in iec61499)
func findSourcesEventName(conns []iec61499.Connection, destChildName string, destEventName string) []string {
	sources := make([]string, 0)
	for _, conn := range conns {
		if destChildName != "" {
			if conn.Destination == destChildName+"."+destEventName {
				sources = append(sources, renameCLocation(conn.Source, 2))
			}
		} else {
			if conn.Destination == destEventName {
				sources = append(sources, renameCLocation(conn.Source, 2))
			}
		}
	}
	return sources
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
		for _, varD := range b.InputVars.Variables {
			if varD.Name == n {
				return &varD
			}
		}
	}

	if b.OutputVars != nil {
		for _, varD := range b.OutputVars.Variables {
			if varD.Name == n {
				return &varD
			}
		}
	}

	return nil
}
