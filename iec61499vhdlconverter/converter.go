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
	vhdlTemplateFuncMap = template.FuncMap{
		"getVhdlType":                   getVhdlType,
		"getVhdlECCTransitionCondition": getVhdlECCTransitionCondition,
		"renameDoneSignal":              renameDoneSignal,
		"renameConnSignal":              renameConnSignal,
		"connChildSourceOnly":           connChildSourceOnly,
		"connChildNameMatches":          connChildNameMatches,
	}
	vhdlTemplates = template.Must(template.New("").Funcs(vhdlTemplateFuncMap).ParseGlob("./vhdltemplates/*"))
)

//Converter is the struct we use to store all blocks for conversion (and what we operate from)
type Converter struct {
	Blocks  []iec61499.FB
	topName string
}

//New returns a new instance of a Converter
func New() (*Converter, error) {
	return &Converter{Blocks: make([]iec61499.FB, 0)}, nil
}

//AddBlock should be called for each block in the network
func (c *Converter) AddBlock(iec61499bytes []byte) error {
	FB := iec61499.FB{}
	if err := xml.Unmarshal(iec61499bytes, &FB); err != nil {
		return errors.New("Couldn't unmarshal iec61499 xml: " + err.Error())
	}

	if err := checkFB(&FB); err != nil {
		return errors.New("FB is not suitable for conversion to VHDL: " + err.Error())
	}
	c.Blocks = append(c.Blocks, FB)

	return nil
}

//VHDLOutput is used when returning the converted vhdl from the iec61499
type VHDLOutput struct {
	Name string
	VHDL []byte
}

//TemplateData is the structure used to hold data being passed into the templating engine
type TemplateData struct {
	BlockIndex int
	Blocks     []iec61499.FB
}

//SetTopName sets the IEC61499 top level entity to the name provided
//This checks to ensure a valid name
func (c *Converter) SetTopName(name string) error {
	if name == "" { //no name provided is valid
		return nil
	}

	found := false
	for i := 0; i < len(c.Blocks); i++ {
		if c.Blocks[i].Name == name {
			found = true
			break
		}
	}

	if found == false {
		return errors.New("Can't find provided top-level name '" + name + "'")
	}

	c.topName = name
	return nil
}

//AllToVHDL converts iec61499 xml (stored as []FB) into vhdl []byte for each block (becomes []VHDLOutput struct)
//Returns nil error on success
func (c *Converter) AllToVHDL() ([]VHDLOutput, error) {

	finishedConversions := make([]VHDLOutput, 0, len(c.Blocks))

	for i := 0; i < len(c.Blocks); i++ {
		output := &bytes.Buffer{}
		templateName := "basicFB"
		if c.Blocks[i].BasicFB == nil {
			templateName = "compositeFB"
		}

		if err := vhdlTemplates.ExecuteTemplate(output, templateName, TemplateData{BlockIndex: i, Blocks: c.Blocks}); err != nil {
			return nil, errors.New("Couldn't format template: " + err.Error())
		}

		finishedConversions = append(finishedConversions, VHDLOutput{Name: c.Blocks[i].Name, VHDL: output.Bytes()})
	}

	if c.topName != "" {
		output := &bytes.Buffer{}
		topIndex := -1
		for i := 0; i < len(c.Blocks); i++ {
			if c.Blocks[i].Name == c.topName {
				topIndex = i
				break
			}
		}

		if topIndex == -1 {
			return nil, errors.New("Can't find provided top-level name '" + c.topName + "'")
		}

		if err := vhdlTemplates.ExecuteTemplate(output, "top", TemplateData{BlockIndex: topIndex, Blocks: c.Blocks}); err != nil {
			return nil, errors.New("Couldn't format template: " + err.Error())
		}

		finishedConversions = append(finishedConversions, VHDLOutput{Name: "iec61499_network_top", VHDL: output.Bytes()})

	}

	return finishedConversions, nil
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
	if strings.ToLower(in) == "and" || strings.ToLower(in) == "or" || strings.ToLower(in) == "not" || strings.ContainsAny(in, "<>=") || strings.ToLower(in) == "true" {
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

//connChildNameMatches is used in templates for location matching
func connChildNameMatches(in string, name string) bool {
	return strings.HasPrefix(in, name+".")
}
