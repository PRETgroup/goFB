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

//FBTikzIO will hold the events and data slices for rendering purposes
type FBTikzIO struct {
	Events []FBTikzIONames
	Data   []FBTikzIONames
}

//FBTikzIONames is used for rendering matched IO on a horizontal level
//You can use the InputAssocPos to link between events and data, by using the same index
//and input/output-ness
type FBTikzIONames struct {
	Input       string
	InputAssoc  FBTikzIOAssocPos
	Output      string
	OutputAssoc FBTikzIOAssocPos
}

//FBTikzIOAssocPos is used for data-variable association
type FBTikzIOAssocPos struct {
	PosX     int //if 0, not linked
	PosEvent int //if 0, this is the event
}

//GetTikzIO will return the IO of the block in a helpful and easy-to-render structre
func (f FBTikzHelper) GetTikzIO() FBTikzIO {
	IO := FBTikzIO{
		Events: make([]FBTikzIONames, f.GetEventsSize()),
		Data:   make([]FBTikzIONames, f.GetVarsSize()),
	}

	//sort out the event area for a FB
	inputEventsPos := 0
	outputEventsPos := 0

	inputAssocPos := make(map[string]FBTikzIOAssocPos)
	outputAssocPos := make(map[string]FBTikzIOAssocPos)

	for i := 0; i < len(IO.Events); i++ {
		if i < len(f.InterfaceList.EventInputs) {
			IO.Events[i].Input = f.InterfaceList.EventInputs[i].Name
			if len(f.InterfaceList.EventInputs[i].With) > 0 {
				inputEventsPos++
				IO.Events[i].InputAssoc = FBTikzIOAssocPos{PosX: inputEventsPos, PosEvent: 0}
				for _, with := range f.InterfaceList.EventInputs[i].With {
					inputAssocPos[with.Var] = FBTikzIOAssocPos{PosX: inputEventsPos, PosEvent: i}
				}
			}
		}
		if i < len(f.InterfaceList.EventOutputs) {
			IO.Events[i].Output = f.InterfaceList.EventOutputs[i].Name
			if len(f.InterfaceList.EventOutputs[i].With) > 0 {
				outputEventsPos++
				IO.Events[i].OutputAssoc = FBTikzIOAssocPos{PosX: outputEventsPos, PosEvent: 0}
				for _, with := range f.InterfaceList.EventOutputs[i].With {
					outputAssocPos[with.Var] = FBTikzIOAssocPos{PosX: outputEventsPos, PosEvent: i}
				}
			}
		}
	}

	//fmt.Printf("%+v\r\n", IO.Events)

	//var names
	for i := 0; i < len(IO.Data); i++ {
		if i < len(f.InterfaceList.InputVars) {
			IO.Data[i].Input = f.InterfaceList.InputVars[i].Name
			if pos, ok := inputAssocPos[f.InterfaceList.InputVars[i].Name]; ok {
				IO.Data[i].InputAssoc = pos
			}
		}
		if i < len(f.InterfaceList.OutputVars) {
			IO.Data[i].Output = f.InterfaceList.OutputVars[i].Name
			if pos, ok := outputAssocPos[f.InterfaceList.OutputVars[i].Name]; ok {
				IO.Data[i].OutputAssoc = pos
			}
		}
	}

	return IO
}
