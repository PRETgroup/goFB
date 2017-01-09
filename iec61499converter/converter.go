package iec61499converter

import (
	"bytes"
	"encoding/xml"
	"errors"
	"strings"
	"text/template"

	"github.com/kiwih/goFB/iec61499converter/iec61499"
)

//Converter is the struct we use to store all blocks for conversion (and what we operate from)
type Converter struct {
	Blocks  []iec61499.FB
	topName string

	ignoreAlgorithmLanguages bool

	outputLanguage language
	templates      *template.Template
}

//New returns a new instance of a Converter based on the provided language
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

//OutputFile is used when returning the converted data from the iec61499
type OutputFile struct {
	Name      string
	Extension string
	Contents  []byte
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

//ConvertAll converts iec61499 xml (stored as []FB) into vhdl []byte for each block (becomes []VHDLOutput struct)
//Returns nil error on success
func (c *Converter) ConvertAll() ([]OutputFile, error) {
	finishedConversions := make([]OutputFile, 0, len(c.Blocks))

	//if a top block is present
	topIndex := -1
	if c.topName != "" {
		for i := 0; i < len(c.Blocks); i++ {
			if c.Blocks[i].Name == c.topName {
				topIndex = i
				break
			}
		}

		if topIndex == -1 {
			return nil, errors.New("Can't find provided top-level name '" + c.topName + "'")
		}
	}

	//convert all function blocks
	for i := 0; i < len(c.Blocks); i++ {
		if c.Blocks[i].Comment == "goFB_ignore" {
			//if they have this comment, for whatever reason, don't generate any files for these blocks (but they may still be referenced)
			continue
		}
		output := &bytes.Buffer{}
		templateName := ""
		if c.Blocks[i].CompositeFB != nil {
			templateName = "compositeFB"
		} else if c.Blocks[i].BasicFB != nil {
			templateName = "basicFB"
		} else if c.Blocks[i].Resources != nil {
			templateName = "deviceFB"
		} else {
			return nil, errors.New("Can't determine type of FB of " + c.Blocks[i].Name)
		}

		if err := c.templates.ExecuteTemplate(output, templateName, TemplateData{BlockIndex: i, Blocks: c.Blocks}); err != nil {
			return nil, errors.New("Couldn't format template (fb) of" + c.Blocks[i].Name + ": " + err.Error())
		}

		finishedConversions = append(finishedConversions, OutputFile{Name: c.Blocks[i].Name, Extension: c.outputLanguage.getExtension(), Contents: output.Bytes()})

		if c.outputLanguage.hasHeaders() {
			output := &bytes.Buffer{}
			templateName := "FBheader"

			if err := c.templates.ExecuteTemplate(output, templateName, TemplateData{BlockIndex: i, Blocks: c.Blocks}); err != nil {
				return nil, errors.New("Couldn't format template (fb header) of" + c.Blocks[i].Name + ": " + err.Error())
			}

			finishedConversions = append(finishedConversions, OutputFile{Name: c.Blocks[i].Name, Extension: c.outputLanguage.getHeaderExtension(), Contents: output.Bytes()})
		}
	}

	//interface with the top file if it is present
	if topIndex != -1 {
		output := &bytes.Buffer{}

		if err := c.templates.ExecuteTemplate(output, "top", TemplateData{BlockIndex: topIndex, Blocks: c.Blocks}); err != nil {
			return nil, errors.New("Couldn't format template (top) of" + c.Blocks[topIndex].Name + ": " + err.Error())
		}

		finishedConversions = append(finishedConversions, OutputFile{Name: "iec61499_network_top", Extension: c.outputLanguage.getExtension(), Contents: output.Bytes()})

	}

	//convert any supporting files
	for _, st := range c.outputLanguage.supportFileTemplates() {
		output := &bytes.Buffer{}

		if err := c.templates.ExecuteTemplate(output, st.templateName, TemplateData{Blocks: c.Blocks}); err != nil {
			return nil, errors.New("Couldn't format template (support) of" + c.Blocks[topIndex].Name + ": " + err.Error())
		}

		finishedConversions = append(finishedConversions, OutputFile{Name: st.fileName, Extension: st.extension, Contents: output.Bytes()})
	}

	return finishedConversions, nil
}
