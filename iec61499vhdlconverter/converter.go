package iec61499vhdlconverter

import (
	"bytes"
	"encoding/xml"
	"errors"
	"regexp"
	"strings"
	"text/template"

	"github.com/kiwih/go-iec61499-vhdl/iec61499vhdlconverter/iec61499"
)

var (
	vhdlTemplateFuncMap = template.FuncMap{"getVhdlType": getVhdlType, "getVhdlECCTransitionCondition": getVhdlECCTransitionCondition}
	vhdlTemplates       = template.Must(template.New("").Funcs(vhdlTemplateFuncMap).ParseGlob("./vhdltemplates/*"))
)

//IEC61499ToVHDL converts iec61499 xml (stored as []byte) into vhdl []byte
//Returns nil error on success
func IEC61499ToVHDL(iec61499bytes []byte) ([]byte, error) {
	FB := iec61499.FB{}
	if err := xml.Unmarshal(iec61499bytes, &FB); err != nil {
		return nil, errors.New("Couldn't unmarshal iec61499 xml: " + err.Error())
	}

	output := &bytes.Buffer{}

	if err := checkFB(&FB); err != nil {
		return nil, errors.New("FB is not suitable for conversion to VHDL: " + err.Error())
	}

	if err := vhdlTemplates.ExecuteTemplate(output, "basicFB", FB); err != nil {
		return nil, errors.New("Couldn't format template: " + err.Error())
	}

	return output.Bytes(), nil
}

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
		vhdlType = "std_logic_vector(7 downto 0)"
	case "any":
		panic("Any type not allowed in conversion")
	default:
		panic("Unknown IEC61499 type: " + iec61499type)
	}

	return vhdlType
}

//getVhdlECCTransitionCondition returns the VHDL "if" condition to use in state machine next state logic
func getVhdlECCTransitionCondition(iec61499trans string) string {
	re := regexp.MustCompile("([a-zA-Z_]+)")
	retVal := iec61499trans
	retVal = strings.Replace(retVal, "!", "not ", -1)
	retVal = strings.Replace(retVal, "AND", "and", -1)
	retVal = strings.Replace(retVal, "OR", "or", -1)
	retVal = re.ReplaceAllStringFunc(retVal, addTrueCheck)
	return retVal
}

func addTrueCheck(in string) string {
	if strings.ToLower(in) == "and" || strings.ToLower(in) == "or" || strings.ToLower(in) == "not" {
		return in
	}
	return in + " = '1'"
}
