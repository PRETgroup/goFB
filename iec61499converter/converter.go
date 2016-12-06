package iec61499converter

import (
	"bytes"
	"encoding/xml"
	"errors"
	"strings"
	"text/template"

	"github.com/kiwih/go-iec61499-vhdl/iec61499converter/iec61499"
)

const (
	languageVHDL = "vhdl"
	languageC    = "c"
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
	vhdlTemplates = template.Must(template.New("").Funcs(vhdlTemplateFuncMap).ParseGlob("./templates/vhdl/*"))

	cTemplateFuncMap = template.FuncMap{}
	cTemplates       = template.Must(template.New("").Funcs(cTemplateFuncMap).ParseGlob("./templates/c/*"))
)

//Converter is the struct we use to store all blocks for conversion (and what we operate from)
type Converter struct {
	Blocks  []iec61499.FB
	topName string

	ignoreAlgorithmLanguages bool

	outputLanguage string
	templates      *template.Template
}

//New returns a new instance of a Converter
func New(language string) (*Converter, error) {
	if strings.ToLower(language) == "vhdl" {
		return &Converter{Blocks: make([]iec61499.FB, 0), outputLanguage: languageVHDL, templates: vhdlTemplates}, nil
	} else if strings.ToLower(language) == "c" {
		return &Converter{Blocks: make([]iec61499.FB, 0), outputLanguage: languageC, templates: cTemplates}, nil
	} else {
		return nil, errors.New("Unknown language for converter")
	}

}

//DisableAlgorithmLanguageChecks prevents checking for compatible languages and assumes VHDL
func (c *Converter) DisableAlgorithmLanguageChecks() {
	c.ignoreAlgorithmLanguages = true
}

//AddBlock should be called for each block in the network
func (c *Converter) AddBlock(iec61499bytes []byte) error {
	FB := iec61499.FB{}
	if err := xml.Unmarshal(iec61499bytes, &FB); err != nil {
		return errors.New("Couldn't unmarshal iec61499 xml: " + err.Error())
	}

	if err := c.checkFB(&FB); err != nil {
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
