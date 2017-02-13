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

	ignoreAlgorithmLanguages bool
	tcrestUsingSPM           bool

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

func (c *Converter) SetTcrestUsingSPM() {
	c.tcrestUsingSPM = true
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
	TcrestUsingSPM bool
	BlockIndex     int
	Blocks         []iec61499.FB
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

	return c.flattenFromCFB(topFB)
}

//flattenFromCFB "flattens out" all the CFBs inside a parent CFB (so that it only contains BFBs)
func (c *Converter) flattenFromCFB(parentCFB *iec61499.FB) error {
	//for each child element, check to see if it is a compositeFB. If so, then we need to put its children in the parent and join up the connections

	if parentCFB.CompositeFB == nil {
		return errors.New("Needs to be a CFB")
	}

	innerFBRefs := parentCFB.CompositeFB.FBs

	for i := 0; i < len(innerFBRefs); i++ {
		t := innerFBRefs[i].Type
		childFB := findBlockDefinitionForType(c.Blocks, t)
		if childFB.CompositeFB != nil {
			c.extractChildrenFromCFBChild(parentCFB, childFB, innerFBRefs[i])
			//fmt.Printf("fin: \n%+v\n", parentCFB.CompositeFB)
		}
	}

	return errors.New("Not yet implemented")
}

//extractChildrenFromCFBChild does several things
//1. it appends all children of childCFB to parentCFB (while ensuring their names remain unique)
//2. it corrects all connections that go from blocks in parentCFB to those children in childCFB and vice versa
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
	for i := 0; i < len(childCFB.CompositeFB.EventConnections); i++ {
		destParts := strings.Split(childCFB.CompositeFB.EventConnections[i].Destination, ".")
		//sourceParts := strings.Split(parentCFB.CompositeFB.EventConnections[i].Source, ".")

		if len(destParts) == 1 { //this connection terminated on a source node for the child (i.e. it terminates on something external)
			fmt.Printf("conn %+v goes to parent\n", childCFB.CompositeFB.EventConnections[i])
			childSourceName := destParts[0]
			for j := 0; j < len(parentCFB.CompositeFB.EventConnections); j++ {
				//sourceParts := strings.Split(, ".")
				if parentCFB.CompositeFB.EventConnections[j].Source == childRef.Name+"."+childSourceName { //len(sourceParts) == 2 && sourceParts[1] == childSourceName {
					fmt.Printf("Matched to parent link %+v\n", parentCFB.CompositeFB.EventConnections[j])
					//bingo, we have a match between a grandchild and another child block
					newConn := childCFB.CompositeFB.EventConnections[i]
					newConn.Source = "Flattened_" + childRef.Name + "_" + newConn.Source
					newConn.Destination = parentCFB.CompositeFB.EventConnections[j].Destination

					fmt.Printf("Appending %+v\n", newConn)

					parentCFB.CompositeFB.EventConnections = append(parentCFB.CompositeFB.EventConnections, newConn)
				}
			}
		}

		// if len(destParts) == 2 && destParts[0] == childRef.Name {
		// 	//this goes to the original child, need to amend the name (i.e. find the inner destinations that this goes to)
		// 	var innerDests []iec61499.Connection
		// 	for j := 0; j < len(childCFB.CompositeFB.EventConnections); j++ {

		// 	}
		// }

		// if len(sourceParts) == 2 && sourceParts[0] == childRef.Name {
		// 	//this comes from the original child, need to amend the name (i.e. find the other child that this connection goes to)
		// 	for j := 0; j < len(childC)
		// }
	}
	//2b). Fix connections that went from Parent to grandchild
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

		if err := c.templates.ExecuteTemplate(output, templateName, TemplateData{BlockIndex: i, Blocks: c.Blocks, TcrestUsingSPM: c.tcrestUsingSPM}); err != nil {
			return nil, errors.New("Couldn't format template (fb) of" + c.Blocks[i].Name + ": " + err.Error())
		}

		finishedConversions = append(finishedConversions, OutputFile{Name: c.Blocks[i].Name, Extension: c.outputLanguage.getExtension(), Contents: output.Bytes()})

		if c.outputLanguage.hasHeaders() {
			output := &bytes.Buffer{}
			templateName := "FBheader"

			if err := c.templates.ExecuteTemplate(output, templateName, TemplateData{BlockIndex: i, Blocks: c.Blocks, TcrestUsingSPM: c.tcrestUsingSPM}); err != nil {
				return nil, errors.New("Couldn't format template (fb header) of" + c.Blocks[i].Name + ": " + err.Error())
			}

			finishedConversions = append(finishedConversions, OutputFile{Name: c.Blocks[i].Name, Extension: c.outputLanguage.getHeaderExtension(), Contents: output.Bytes()})
		}
	}

	//interface with the top file if it is present
	if topIndex != -1 {
		output := &bytes.Buffer{}

		if err := c.templates.ExecuteTemplate(output, "top", TemplateData{BlockIndex: topIndex, Blocks: c.Blocks, TcrestUsingSPM: c.tcrestUsingSPM}); err != nil {
			return nil, errors.New("Couldn't format template (top) of" + c.Blocks[topIndex].Name + ": " + err.Error())
		}

		finishedConversions = append(finishedConversions, OutputFile{Name: "iec61499_network_top", Extension: c.outputLanguage.getExtension(), Contents: output.Bytes()})

	}

	//convert any supporting files
	for _, st := range c.outputLanguage.supportFileTemplates() {
		output := &bytes.Buffer{}

		if err := c.templates.ExecuteTemplate(output, st.templateName, TemplateData{Blocks: c.Blocks, TcrestUsingSPM: c.tcrestUsingSPM}); err != nil {
			return nil, errors.New("Couldn't format template (support) of" + c.Blocks[topIndex].Name + ": " + err.Error())
		}

		finishedConversions = append(finishedConversions, OutputFile{Name: st.fileName, Extension: st.extension, Contents: output.Bytes()})
	}

	return finishedConversions, nil
}
