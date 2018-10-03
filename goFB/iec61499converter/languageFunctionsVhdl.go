package iec61499converter

import (
	"regexp"
	"strings"
)

const (
	vhdl_TOPIO_IN  = "TOPIO_IN" //if either of the TOPIO_ strings are in an event, var, or internal variable comment, it means these should be passed up to the top level file and used as global IO (used in VHDL only)
	vhdl_TOPIO_OUT = "TOPIO_OUT"
)

//getVhdlType returns the VHDL type to use with respect to an IEC61499 type
func getVhdlType(iec61499type string) string {
	vhdlType := ""
	switch strings.ToLower(iec61499type) {
	//IEC61499 types
	case "bool":
		vhdlType = "std_logic"
	case "byte":
		vhdlType = "std_logic_vector(7 downto 0)"
	case "word":
		vhdlType = "std_logic_vector(15 downto 0)"
	case "dword":
		vhdlType = "std_logic_vector(31 downto 0)"
	case "lword":
		vhdlType = "std_logic_vector(63 downto 0)"
	case "sint":
		vhdlType = "signed(7 downto 0)"
	case "usint":
		vhdlType = "unsigned(7 downto 0)"
	case "int":
		vhdlType = "signed(15 downto 0)"
	case "uint":
		vhdlType = "unsigned(15 downto 0)"
	case "dint":
		vhdlType = "signed(31 downto 0)"
	case "udint":
		vhdlType = "unsigned(31 downto 0)"
	case "lint":
		vhdlType = "signed(63 downto 0)"
	case "ulint":
		vhdlType = "unsigned(63 downto 0)"
	case "real":
		panic("Real type not allowed in conversion")
	case "lreal":
		panic("Lreal type not allowed in conversion")
	case "time":
		vhdlType = "unsigned(63 downto 0)"
	case "any":
		panic("Any type not allowed in conversion")
	//C types
	case "uint32_t":
		vhdlType = "unsigned(31 downto 0)"
	case "int32_t":
		vhdlType = "signed(31 downto 0)"
	case "float":
		panic("Float type not allowed in conversion")
	case "double":
		panic("Double type not allowed in conversion")
	case "string":
		panic("String type not allowed in conversion")
	case "char":
		vhdlType = "unsigned(7 downto 0)"
	default:
		panic("Unknown IEC61499 type: " + iec61499type)
	}

	return vhdlType
}

//getVerilogECCTransitionCondition returns the VHDL "if" condition to use in state machine next state logic
func getVerilogECCTransitionCondition(iec61499trans string) string {
	re := regexp.MustCompile("([a-zA-Z_<>=]+)")
	retVal := iec61499trans
	retVal = strings.Replace(retVal, "!", "not ", -1)
	retVal = strings.Replace(retVal, "AND", "and", -1)
	retVal = strings.Replace(retVal, "OR", "or", -1)
	retVal = re.ReplaceAllStringFunc(retVal, addTrueCheck)
	return retVal
}
