package main

import (
	"bytes"
	"encoding/xml"
	"errors"

	"github.com/PRETgroup/goFB/iec61499"
)

//FBTikz is the structure we use to hold everything
type FBTikz struct {
	Blocks []iec61499.FB
}

//OutputFile is used when returning the converted data from the iec61499
type OutputFile struct {
	Name      string
	Extension string
	Contents  []byte
}

//AddBlock should be called for each block in the network
func (f *FBTikz) AddBlock(iec61499bytes []byte) error {
	FB := iec61499.FB{}
	if err := xml.Unmarshal(iec61499bytes, &FB); err != nil {
		return errors.New("Couldn't unmarshal iec61499 xml: " + err.Error())
	}

	f.Blocks = append(f.Blocks, FB)

	return nil
}

//ConvertAll will compile all the FBs to Tikz files
func (f *FBTikz) ConvertAll() ([]OutputFile, error) {
	finishedConversions := make([]OutputFile, 0, len(f.Blocks))

	for i := 0; i < len(f.Blocks); i++ {

		output := &bytes.Buffer{}

		if err := tikzTemplate.ExecuteTemplate(output, "", FBTikzHelper(f.Blocks[i])); err != nil {
			return nil, errors.New("Couldn't format template (fb) of" + f.Blocks[i].Name + ": " + err.Error())
		}

		finishedConversions = append(finishedConversions, OutputFile{Name: f.Blocks[i].Name + "-tikz", Extension: "tex", Contents: output.Bytes()})
	}

	return finishedConversions, nil
}

//FBTikzHelper overleads the iec61499.FB type with extra helper functions
type FBTikzHelper iec61499.FB

//GetVarsSize returns the height of the vars area that needs to be drawen
func (f FBTikzHelper) GetVarsSize() int {
	if len(f.InterfaceList.InputVars) > len(f.InterfaceList.OutputVars) {
		return len(f.InterfaceList.InputVars)
	}
	return len(f.InterfaceList.OutputVars)
}

//GetEventsSize returns the height of the vars area that needs to be drawen
func (f FBTikzHelper) GetEventsSize() int {
	if len(f.InterfaceList.EventInputs) > len(f.InterfaceList.EventOutputs) {
		return len(f.InterfaceList.EventInputs)
	}
	return len(f.InterfaceList.EventOutputs)
}

//FBTikzIONames is used for rendering matched IO
type FBTikzIONames struct {
	Input  string
	Output string
}

//GetMatchedVars returns the matched names of the variable area for a FB
func (f FBTikzHelper) GetMatchedVars() []FBTikzIONames {
	matchedNames := make([]FBTikzIONames, f.GetVarsSize())
	for i := 0; i < f.GetVarsSize(); i++ {
		if i < len(f.InterfaceList.InputVars) {
			matchedNames[i].Input = f.InterfaceList.InputVars[i].Name
		}
		if i < len(f.InterfaceList.OutputVars) {
			matchedNames[i].Output = f.InterfaceList.OutputVars[i].Name
		}
	}
	return matchedNames
}

//GetMatchedEvents returns the matched names of the event area for a FB
func (f FBTikzHelper) GetMatchedEvents() []FBTikzIONames {
	matchedNames := make([]FBTikzIONames, f.GetEventsSize())
	for i := 0; i < f.GetEventsSize(); i++ {
		if i < len(f.InterfaceList.EventInputs) {
			matchedNames[i].Input = f.InterfaceList.EventInputs[i].Name
		}
		if i < len(f.InterfaceList.EventOutputs) {
			matchedNames[i].Output = f.InterfaceList.EventOutputs[i].Name
		}
	}
	return matchedNames
}
