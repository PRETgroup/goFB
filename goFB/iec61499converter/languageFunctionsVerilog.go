package iec61499converter

import (
	"regexp"
	"strings"

	"github.com/PRETgroup/goFB/goFB/stconverter"
	"github.com/PRETgroup/goFB/iec61499"
)

func verilogCompileTransition(block iec61499.FB, trans string) string {
	//transitions are all ST type
	stconverter.SetKnownVarNames(block.GetAllVarNames())
	instrs, err := stconverter.ParseString(block.Name+"(transition)", trans)
	if err != nil {
		panic(err)
	}
	comp := stconverter.VerilogCompileSequence(instrs, false)
	return comp
}

func verilogCompileAlgorithm(block iec61499.FB, algorithm iec61499.Algorithm) string {
	//if it's ST we know how to compile that! :)
	if algorithm.Other.Language == "ST" {
		stconverter.SetKnownVarNames(block.GetAllVarNames())
		instrs, err := stconverter.ParseString(block.Name+"_"+algorithm.Name, algorithm.Other.Text)
		if err != nil {
			panic(err)
		}
		comp := stconverter.VerilogCompileSequence(instrs, true)
		return comp
	}
	//can't do much otherwise...
	return algorithm.Other.Text
}

//getVerilogSize returns the Verilog size to use with respect to an IEC61499 type
func getVerilogSize(iec61499type string) string {
	verilogType := ""
	switch strings.ToLower(iec61499type) {
	//IEC61499 types
	case "bool":
		verilogType = ""
	case "byte":
		verilogType = "[7:0]"
	case "word":
		verilogType = "[15:0]"
	case "dword":
		verilogType = "[31:0]"
	case "lword":
		verilogType = "[63:0]"
	case "sint":
		verilogType = "signed [7:0]"
	case "usint":
		verilogType = "unsigned [7:0]"
	case "int":
		verilogType = "signed [15:0]"
	case "uint":
		verilogType = "unsigned [15:0]"
	case "dint":
		verilogType = "signed [31:0]"
	case "udint":
		verilogType = "unsigned [31:0]"
	case "lint":
		verilogType = "signed [63:0]"
	case "ulint":
		verilogType = "unsigned [63:0]"
	case "real":
		panic("Real type not allowed in conversion")
	case "lreal":
		panic("Lreal type not allowed in conversion")
	case "time":
		verilogType = "unsigned [63:0]"
	case "any":
		panic("Any type not allowed in conversion")
	//C types
	case "uint32_t":
		verilogType = "unsigned [31:0]"
	case "int32_t":
		verilogType = "signed [31:0]"
	case "float":
		panic("Float type not allowed in conversion")
	case "double":
		panic("Double type not allowed in conversion")
	case "string":
		panic("String type not allowed in conversion")
	case "char":
		verilogType = "unsigned [7:0]"
	default:
		panic("Unknown IEC61499 type: " + iec61499type)
	}

	return verilogType
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
