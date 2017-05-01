package iec61499converter

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"strings"
	"text/template"

	"github.com/kiwih/goFB/iec61499converter/iec61499"
)

//Converter is the struct we use to store all blocks for conversion (and what we operate from)
type Converter struct {
	Blocks  []iec61499.FB
	topName string

	ConverterSettings

	outputLanguage language
	templates      *template.Template
}

//New returns a new instance of a Converter based on the provided language
func New(language string) (*Converter, error) {
	if strings.ToLower(language) == "vhdl" {
		return &Converter{Blocks: make([]iec61499.FB, 0), outputLanguage: languageVHDL, templates: vhdlTemplates}, nil
	} else if strings.ToLower(language) == "c" {
		return &Converter{Blocks: make([]iec61499.FB, 0), outputLanguage: languageC, templates: cTemplates}, nil
	} else if strings.ToLower(language) == "eventc" {
		return &Converter{Blocks: make([]iec61499.FB, 0), outputLanguage: languageEventC, templates: eventCTemplates}, nil
	}

	return nil, errors.New("Unknown language " + language + "for converter")
}

//DisableAlgorithmLanguageChecks prevents checking for compatible languages inside algorithms
func (c *Converter) DisableAlgorithmLanguageChecks() {
	c.IgnoreAlgorithmLanguages = true
}

//SetTcrestUsingSPM sets a flag that the output should be formatted for the TCREST architecture, specifically,
// it should put the FBs into SPM memory
func (c *Converter) SetTcrestUsingSPM() {
	c.TcrestUsingSPM = true
}

//SetTcrestIncludes sets that the output fbtypes.h should include the TCREST headers
// than set and forget
func (c *Converter) SetTcrestIncludes() {
	c.TcrestIncludes = true
}

//SetTcrestSmartSPM sets that the output should be formatted for T-CREST architecture and use the SPMs for BFB execution
func (c *Converter) SetTcrestSmartSPM() {
	c.TcrestSmartSPM = true
}

//CvodeEnable enables using C CVODE library from SUNDIALS to solve algorithms with 'ODE' and 'ODE_init' in comment fields
func (c *Converter) CvodeEnable() error {
	if c.outputLanguage != languageC {
		return errors.New("Can't enable cvode unless output language is C")
	}
	c.CvodeEnabled = true
	return nil
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

//ConverterSettings holds the settings for this conversion
type ConverterSettings struct {
	TcrestUsingSPM           bool
	TcrestSmartSPM           bool
	TcrestIncludes           bool
	IgnoreAlgorithmLanguages bool
	CvodeEnabled             bool
}

//OutputFile is used when returning the converted data from the iec61499
type OutputFile struct {
	Name      string
	Extension string
	Contents  []byte
}

//TemplateData is the structure used to hold data being passed into the templating engine
type TemplateData struct {
	ConverterSettings

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

//Flatten flattens all CFBs from a network (excluding the parent)
func (c *Converter) Flatten() error {
	if c.topName == "" {
		return errors.New("Top Block needs to be set before Flatten() can be called")
	}

	//find the parent
	var topFB *iec61499.FB
	for i := 0; i < len(c.Blocks); i++ {
		if c.Blocks[i].Name == c.topName {
			topFB = &c.Blocks[i]
		}
	}
	if topFB == nil {
		return errors.New("Can't find top block when flattening")
	}

	if topFB.CompositeFB != nil {
		//top block is composite or resource
		return c.flattenFromCFB(topFB)
	}

	//top block is probably device
	if len(topFB.Resources) > 0 {
		//yup, it is. Flatten each of the resource blocks independently
		for i := 0; i < len(topFB.Resources); i++ {
			for j := 0; j < len(c.Blocks); j++ {
				if topFB.Resources[i].Type == c.Blocks[j].Name {
					if err := c.flattenFromCFB(&c.Blocks[j]); err != nil {
						return err
					}
				}
			}
		}
		//done
		return nil

	}

	return errors.New("Can't flatten top block, is neither composite/resource nor device")
}

//flattenFromCFB "flattens out" all the CFBs inside a parent CFB (so that it only contains BFBs)
func (c *Converter) flattenFromCFB(parentCFB *iec61499.FB) error {
	//for each child element, check to see if it is a compositeFB. If so, then we need to put its children in the parent and join up the connections

	if parentCFB.CompositeFB == nil {
		return errors.New("Needs to be an RFB or a CFB")
	}

	for i := 0; i < len(parentCFB.CompositeFB.FBs); i++ {
		t := parentCFB.CompositeFB.FBs[i].Type
		childFB := findBlockDefinitionForType(c.Blocks, t)
		if childFB.CompositeFB != nil {
			c.extractChildrenFromCFBChild(parentCFB, childFB, parentCFB.CompositeFB.FBs[i])
			i--
			//fmt.Printf("fin: \n%+v\n", parentCFB.CompositeFB)
		}
	}
	return nil
}

//extractChildrenFromCFBChild does several things
//1. it appends all children of childCFB to parentCFB (while ensuring their names remain unique)
//2. it corrects all connections that go from blocks in parentCFB to those children in childCFB and vice versa, and parameters
//3. it appends all connections that go from grandchild to grandchild to the new blocks
//4. it removes the reference to childCFB from parentCFB
func (c *Converter) extractChildrenFromCFBChild(parentCFB *iec61499.FB, childCFB *iec61499.FB, childRef iec61499.FBReference) {
	fmt.Printf("Extracting children from '%s' in '%s' type '%s'\n", childRef.Name, parentCFB.Name, childCFB.Name)
	//1. Add grandchildren to parent
	grandchildrenFBs := make([]iec61499.FBReference, len(childCFB.CompositeFB.FBs))
	for i := 0; i < len(grandchildrenFBs); i++ {
		grandchildrenFBs[i] = childCFB.CompositeFB.FBs[i]                                        //copy out the ref
		grandchildrenFBs[i].Name = "Flattened_" + childRef.Name + "_" + grandchildrenFBs[i].Name //fix the name
	}
	parentCFB.CompositeFB.FBs = append(parentCFB.CompositeFB.FBs, grandchildrenFBs...) //append them to parent

	//2a). Fix connections that went from grandchild to Parent

	addChildToParentConns := func(childConns []iec61499.Connection, parentConns *[]iec61499.Connection) {
		for i := 0; i < len(childConns); i++ {
			destParts := strings.Split(childConns[i].Destination, ".")

			if len(destParts) == 1 { //this connection terminated on a source node for the child (i.e. it terminates on something external)
				//fmt.Printf("conn %+v goes to parent\n", childConns[i])
				childSourceName := destParts[0]
				for j := 0; j < len(*parentConns); j++ {

					if (*parentConns)[j].Source == childRef.Name+"."+childSourceName {
						//fmt.Printf("Matched to parent link %+v\n", parentConns[j])
						//bingo, we have a match between a grandchild and another child block
						newConn := childConns[i]
						newConn.Source = "Flattened_" + childRef.Name + "_" + newConn.Source
						newConn.Destination = (*parentConns)[j].Destination
						//fmt.Printf("Appending %+v\n", newConn)
						*parentConns = append(*parentConns, newConn)
					}
				}
			}
		}
	}

	addChildToParentConns(childCFB.CompositeFB.EventConnections, &parentCFB.CompositeFB.EventConnections)
	addChildToParentConns(childCFB.CompositeFB.DataConnections, &parentCFB.CompositeFB.DataConnections)

	//2b). Fix connections that went from Parent to grandchild

	addParentToChildConns := func(childConns []iec61499.Connection, parentConns *[]iec61499.Connection) {
		for i := 0; i < len(*parentConns); i++ {
			destParts := strings.Split((*parentConns)[i].Destination, ".")

			if len(destParts) == 2 && destParts[0] == childRef.Name { //this connection terminated on a source node for the child (i.e. it terminates on something external)
				//fmt.Printf("conn %+v goes to child\n", childConns[i])
				childSourceName := destParts[1]
				for j := 0; j < len(childConns); j++ {

					if childConns[j].Source == childSourceName {
						//fmt.Printf("Matched to parent link %+v\n", parentConns[j])
						//bingo, we have a match between a child block and a grandchild
						newConn := childConns[j]
						newConn.Destination = "Flattened_" + childRef.Name + "_" + newConn.Destination
						newConn.Source = (*parentConns)[i].Source
						//fmt.Printf("Appending %+v\n", newConn)
						*parentConns = append(*parentConns, newConn)
					}
				}
			}
		}
	}

	addParentToChildConns(childCFB.CompositeFB.EventConnections, &parentCFB.CompositeFB.EventConnections)
	addParentToChildConns(childCFB.CompositeFB.DataConnections, &parentCFB.CompositeFB.DataConnections)

	//2c). Move relevant parameters from child to former grandchildren
	for i := 0; i < len(childRef.Parameter); i++ {
		param := childRef.Parameter[i]

		for j := 0; j < len(childCFB.CompositeFB.DataConnections); j++ {
			sourceParts := strings.Split(childCFB.CompositeFB.DataConnections[j].Source, ".")
			if len(sourceParts) == 1 && sourceParts[0] == param.Name {
				destParts := strings.Split(childCFB.CompositeFB.DataConnections[j].Destination, ".")
				if len(destParts) == 2 {
					destFBRefName := destParts[0]
					for k := 0; k < len(parentCFB.CompositeFB.FBs); k++ {
						if parentCFB.CompositeFB.FBs[k].Name == "Flattened_"+childRef.Name+"_"+destFBRefName {
							param.Name = destParts[1]
							parentCFB.CompositeFB.FBs[k].Parameter = append(parentCFB.CompositeFB.FBs[k].Parameter, param)
						}
					}
				}
			}
		}
	}

	//2d). Remove connections that 2a) and 2b) replace

	removeChildConnsFromParent := func(parentConns *[]iec61499.Connection) {
		for i := 0; i < len(*parentConns); i++ {
			sourceParts := strings.Split((*parentConns)[i].Source, ".")
			destParts := strings.Split((*parentConns)[i].Destination, ".")
			//fmt.Printf("Removing %+v ? ", (*parentConns)[i])
			if (len(sourceParts) == 2 && sourceParts[0] == childRef.Name) ||
				(len(destParts) == 2 && destParts[0] == childRef.Name) {

				//fmt.Printf("Yes\n")
				*parentConns = append((*parentConns)[:i], (*parentConns)[i+1:]...)
				i--
			} else {
				//fmt.Printf("No\n")
			}
		}
	}

	removeChildConnsFromParent(&parentCFB.CompositeFB.EventConnections)
	removeChildConnsFromParent(&parentCFB.CompositeFB.DataConnections)

	//3. append all connections that go from grandchild to grandchild to the new blocks in parent cfb
	addChildToChildConns := func(childConns []iec61499.Connection, parentConns *[]iec61499.Connection) {
		for i := 0; i < len(childConns); i++ {
			sourceParts := strings.Split(childConns[i].Source, ".")
			destParts := strings.Split(childConns[i].Destination, ".")

			if len(sourceParts) == 2 && len(destParts) == 2 { //these links don't go to the parent
				newConn := childConns[i]
				newConn.Source = "Flattened_" + childRef.Name + "_" + newConn.Source
				newConn.Destination = "Flattened_" + childRef.Name + "_" + newConn.Destination
				//fmt.Printf("Appending %+v\n", newConn)
				*parentConns = append(*parentConns, newConn)
			}
		}
	}

	addChildToChildConns(childCFB.CompositeFB.EventConnections, &parentCFB.CompositeFB.EventConnections)
	addChildToChildConns(childCFB.CompositeFB.DataConnections, &parentCFB.CompositeFB.DataConnections)

	//4. remove the reference to childCFB from parentCFB

	for i := 0; i < len(parentCFB.CompositeFB.FBs); i++ {
		if parentCFB.CompositeFB.FBs[i].Name == childRef.Name {
			parentCFB.CompositeFB.FBs = append(parentCFB.CompositeFB.FBs[:i], parentCFB.CompositeFB.FBs[i+1:]...)
		}
	}

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
			//there is a possibility that this BFB is infact an odeFB (which have different semantics)
			if c.ConverterSettings.CvodeEnabled && blockNeedsCvode(c.Blocks[i]) {
				templateName = "odeFB"
			}
		} else if c.Blocks[i].Resources != nil {
			templateName = "deviceFB"
		} else {
			return nil, errors.New("Can't determine type of FB of " + c.Blocks[i].Name)
		}

		if err := c.templates.ExecuteTemplate(output, templateName, TemplateData{BlockIndex: i, Blocks: c.Blocks, ConverterSettings: c.ConverterSettings}); err != nil {
			return nil, errors.New("Couldn't format template (fb) of" + c.Blocks[i].Name + ": " + err.Error())
		}

		finishedConversions = append(finishedConversions, OutputFile{Name: "FB_" + c.Blocks[i].Name, Extension: c.outputLanguage.getExtension(), Contents: output.Bytes()})

		if c.outputLanguage.hasHeaders() {
			output := &bytes.Buffer{}
			templateName := "FBheader"

			if err := c.templates.ExecuteTemplate(output, templateName, TemplateData{BlockIndex: i, Blocks: c.Blocks, ConverterSettings: c.ConverterSettings}); err != nil {
				return nil, errors.New("Couldn't format template (fb header) of" + c.Blocks[i].Name + ": " + err.Error())
			}

			finishedConversions = append(finishedConversions, OutputFile{Name: "FB_" + c.Blocks[i].Name, Extension: c.outputLanguage.getHeaderExtension(), Contents: output.Bytes()})
		}
	}

	//interface with the top file if it is present
	if topIndex != -1 {
		output := &bytes.Buffer{}

		if err := c.templates.ExecuteTemplate(output, "top", TemplateData{BlockIndex: topIndex, Blocks: c.Blocks, ConverterSettings: c.ConverterSettings}); err != nil {
			return nil, errors.New("Couldn't format template (top) of" + c.Blocks[topIndex].Name + ": " + err.Error())
		}

		finishedConversions = append(finishedConversions, OutputFile{Name: "top", Extension: c.outputLanguage.getExtension(), Contents: output.Bytes()})

	}

	//convert any supporting files
	for _, st := range c.outputLanguage.supportFileTemplates() {
		output := &bytes.Buffer{}

		if err := c.templates.ExecuteTemplate(output, st.templateName, TemplateData{Blocks: c.Blocks, ConverterSettings: c.ConverterSettings}); err != nil {
			return nil, errors.New("Couldn't format template (support) of" + c.Blocks[topIndex].Name + ": " + err.Error())
		}

		finishedConversions = append(finishedConversions, OutputFile{Name: st.fileName, Extension: st.extension, Contents: output.Bytes()})
	}

	return finishedConversions, nil
}
