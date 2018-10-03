package iec61499converter

import (
	"regexp"
	"strings"

	"github.com/PRETgroup/goFB/iec61499"
)

//getVerilogType returns the VHDL type to use with respect to an IEC61499 type
func getVerilogType(iec61499type string) string {
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
	return in
	//return in + " = '1'"
}

//in our algorithms, DONE needs to be turned into the correct signal name
func renameDoneSignal(in string, name string) string {
	return strings.Replace(in, "DONE", name+"_alg_done", -1)
}

//renameConnSignal is used in templates to make a consistent and friendly name for the connections
func renameConnSignal(in string) string {
	return strings.Replace(in, ".", "_", -1) + "_conn" // + "_to_" + strings.Replace(c.Destination, ".", "_", -1)
}

//variableIsTOPIO_OUT used in templates
func variableIsTOPIO_OUT(v iec61499.Variable) bool {
	return v.Comment == vhdl_TOPIO_OUT
}

//variableIsTOPIO_IN used in templates
func variableIsTOPIO_IN(v iec61499.Variable) bool {
	return v.Comment == vhdl_TOPIO_IN
}

//eventIsTOPIO_OUT used in templates
func eventIsTOPIO_OUT(e iec61499.Event) bool {
	return e.Comment == vhdl_TOPIO_OUT
}

//eventIsTOPIO_IN used in templates
func eventIsTOPIO_IN(e iec61499.Event) bool {
	return e.Comment == vhdl_TOPIO_IN
}

//SpecialIO is used to store internal variables that are "special" (i.e. exported because they are for debugging or service interfaces)
type SpecialIO struct {
	//Perhaps in future we will have special []Event and []Variable for normal event and data API
	InternalVars []iec61499.Variable
}

//getSpecialIOForRef returns all SpecialIO for a given FBReference
func getSpecialIOForRef(fr iec61499.FBReference, otherBlocks []iec61499.FB) SpecialIO {
	for j := 0; j < len(otherBlocks); j++ {
		if otherBlocks[j].Name == fr.Type {
			return getSpecialIO(otherBlocks[j], otherBlocks)
		}
	}
	return SpecialIO{}
}

//getSpecialIO is used for service interface blocks and those blocks that contain service interface blocks
func getSpecialIO(f iec61499.FB, otherBlocks []iec61499.FB) SpecialIO {
	s := SpecialIO{
		InternalVars: make([]iec61499.Variable, 0),
	}

	if f.BasicFB != nil {
		if f.BasicFB.InternalVars != nil {
			for i := 0; i < len(f.BasicFB.InternalVars); i++ {
				if variableIsTOPIO_IN(f.BasicFB.InternalVars[i]) || variableIsTOPIO_OUT(f.BasicFB.InternalVars[i]) {
					s.InternalVars = append(s.InternalVars, f.BasicFB.InternalVars[i])
				}
			}
		}
	} else if f.CompositeFB != nil {
		for i := 0; i < len(f.CompositeFB.FBs); i++ {
			for j := 0; j < len(otherBlocks); j++ {
				if otherBlocks[j].Name == f.CompositeFB.FBs[i].Type {
					os := getSpecialIO(otherBlocks[j], otherBlocks)
					s.InternalVars = append(s.InternalVars, os.InternalVars...)
					continue
				}
			}
		}
	}

	return s
}
