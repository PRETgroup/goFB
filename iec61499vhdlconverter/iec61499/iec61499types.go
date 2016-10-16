package iec61499

import (
	"encoding/xml"
	"errors"
	"strings"
)

const (
	TOPIO_IN  = "TOPIO_IN" //if either of the TOPIO_ strings are in an event, var, or internal variable comment, it means these should be passed up to the top level file and used as global IO
	TOPIO_OUT = "TOPIO_OUT"
)

//FB is the overall type for function blocks
type FB struct {
	Name           string `xml:"Name,attr"`
	Comment        string `xml:"Comment,attr"`
	Identification struct {
		Standard string `xml:"Standard,attr"`
	}
	VersionInfo struct {
		Organization string `xml:"Organization,attr"`
		Version      string `xml:"Version,attr"`
		Author       string `xml:"Author,attr"`
		Date         string `xml:"Date,attr"`
	}
	CompilerInfo struct {
		Header   string `xml:"header,attr"`
		Classdef string `xml:"classdef,attr"`
	}
	//interface
	EventInputs  *EventDeclare `xml:"InterfaceList>EventInputs,omitempty"`
	EventOutputs *EventDeclare `xml:"InterfaceList>EventOutputs,omitempty"`
	InputVars    *VarDeclare   `xml:"InterfaceList>InputVars,omitempty"`
	OutputVars   *VarDeclare   `xml:"InterfaceList>OutputVars,omitempty"`

	BasicFB     *BasicFB     `xml:",omitempty"`
	CompositeFB *CompositeFB `xml:",omitempty"`
}

//CompositeFB is the type for composition function blocks (should be embedded into FB)
type CompositeFB struct {
	XMLName          xml.Name      `xml:"FBNetwork"`
	FBs              []FBReference `xml:"FB"`
	EventConnections []Connection  `xml:"EventConnections>Connection,omitempty"`
	DataConnections  []Connection  `xml:"DataConnections>Connection,omitempty"`
}

//FBReference are used inside CompositeFBs as instances of other function blocks
type FBReference struct {
	Name      string      `xml:"Name,attr"`
	Type      string      `xml:"Type,attr"`
	X         string      `xml:"x,attr"`
	Y         string      `xml:"y,attr"`
	Parameter []Parameter `xml:",omitempty"`
}

//Parameter is a custom values associated with a FBReferences
type Parameter struct {
	Name  string `xml:"Name,attr"`
	Value string `xml:"Value,attr"`
}

//Connection is used to link either events or data lines together inside CompositeFBs
type Connection struct {
	Source      string `xml:"Source,attr"`
	Destination string `xml:"Destination,attr"`
	Dx1         string `xml:"dx1,attr"`
}

//BasicFB is used for basic function blocks (ie FB definitions)
type BasicFB struct {
	InternalVars *VarDeclare    `xml:"InternalVars,omitempty"`
	States       []ECState      `xml:"ECC>ECState"`
	Transistions []ECTransition `xml:"ECC>ECTransition,omitempty"`
	Algorithms   []Algorithm    `xml:"Algorithm,omitempty"`
}

//ECState is a state in the ECC (Execution control chart) of a BasicFB
type ECState struct {
	Name      string   `xml:"Name,attr"`
	ECActions []Action `xml:"ECAction,omitempty"`
	Comment   string   `xml:"Comment,attr"`
	X         string   `xml:"x,attr"`
	Y         string   `xml:"y,attr"`
}

//Action is a link between an ECState and an Algorithm
type Action struct {
	Algorithm string `xml:"Algorithm,attr,omitempty"`
	Output    string `xml:"Output,attr,omitempty"`
}

//ECTransition is a transition in the ECC of a BasicFB
type ECTransition struct {
	Source      string `xml:"Source,attr"`
	Destination string `xml:"Destination,attr"`
	Condition   string `xml:"Condition,attr"`
	X           string `xml:"x,attr"`
	Y           string `xml:"y,attr"`
}

//Algorithm is some code associated with a BasicFB
type Algorithm struct {
	Name    string        `xml:"Name,attr"`
	Comment string        `xml:"Comment,attr"`
	Other   OtherLanguage `xml:",omitEmpty"`
}

//OtherLanguage states what language an algorithm is written in
type OtherLanguage struct {
	XMLName  xml.Name `xml:"Other"`
	Language string   `xml:"Language,attr"`
	Text     string   `xml:"Text,attr"`
}

//VarDeclare is used to store variable declarations of BasicFBs and for special IO of interface FBs
type VarDeclare struct {
	Variables []Variable `xml:"VarDeclaration"`
}

//Variable is used as a variable in an algorithm of a BasicFB
type Variable struct {
	Name      string `xml:"Name,attr"`
	Type      string `xml:"Type,attr"`
	ArraySize string `xml:"ArraySize,attr,omitempty"`
	Comment   string `xml:"Comment,attr"`
}

//EventDeclare is used to store event declarations of BasicFBs
type EventDeclare struct {
	Events []Event `xml:"Event"`
}

//Event is used to store events in BasicFBs
type Event struct {
	Name    string `xml:"Name,attr"`
	Comment string `xml:"Comment,attr"`
	With    []With `xml:",omitEmpty"`
}

//With associates a variable and an event (as data only changes on events)
type With struct {
	Var string `xml:"Var,attr"`
}

//GetTransitionsForState returns only the transitions for a given source state
func (b *BasicFB) GetTransitionsForState(source string) []ECTransition {
	trans := make([]ECTransition, 0, len(b.Transistions))
	for i := 0; i < len(b.Transistions); i++ {
		if b.Transistions[i].Source == source {
			trans = append(trans, b.Transistions[i])
		}
	}
	return trans
}

//VhdlName is used in templates to make a consistent and friendly name for the connections
func (c *Connection) VhdlName() string {
	return strings.Replace(c.Source, ".", "_", -1) + "_to_" + strings.Replace(c.Destination, ".", "_", -1)
}

//FromName is used in templates for location matching
func (c *Connection) FromName(name string) bool {
	return strings.HasPrefix(c.Source, name)
}

//ToName is used in templates for location matching
func (c *Connection) ToName(name string) bool {
	return strings.HasPrefix(c.Destination, name)
}

//SourceApiNameOnly is used in templates for getting rid of prefix stuff
func (c *Connection) SourceApiNameOnly() string {
	splitName := strings.Split(c.Source, ".")
	return splitName[len(splitName)-1]
}

//DestApiNameOnly is used in templates for getting rid of prefix stuff
func (c *Connection) DestApiNameOnly() string {
	splitName := strings.Split(c.Destination, ".")
	return splitName[len(splitName)-1]
}

type ConnectionWithType struct {
	Connection
	Type string
}

//IsLoadFor used in templates for triggering updates of input data
func (e *Event) IsLoadFor(v *Variable) bool {
	for i := 0; i < len(e.With); i++ {
		if e.With[i].Var == v.Name {
			return true
		}
	}
	return false
}

//IsTOPIO_OUT used in templates
func (v *Variable) IsTOPIO_OUT() bool {
	return v.Comment == TOPIO_OUT
}

//IsTOPIO_IN used in templates
func (v *Variable) IsTOPIO_IN() bool {
	return v.Comment == TOPIO_IN
}

func (e *Event) IsTOPIO_OUT() bool {
	return e.Comment == TOPIO_OUT
}

func (e *Event) IsTOPIO_IN() bool {
	return e.Comment == TOPIO_IN
}

type SpecialIO struct {
	//Perhaps in future we will have special []Event and []Variable for normal event and data API
	InternalVars []Variable
}

func (fr FBReference) GetSpecialIO(otherBlocks []FB) SpecialIO {
	for j := 0; j < len(otherBlocks); j++ {
		if otherBlocks[j].Name == fr.Type {
			return otherBlocks[j].GetSpecialIO(otherBlocks)
		}
	}
	return SpecialIO{}
}

//GetSpecialIO is used for service interface blocks and those blocks that contain service interface blocks
func (f FB) GetSpecialIO(otherBlocks []FB) SpecialIO {
	s := SpecialIO{
		InternalVars: make([]Variable, 0),
	}

	if f.BasicFB != nil {
		if f.BasicFB.InternalVars != nil {
			for i := 0; i < len(f.BasicFB.InternalVars.Variables); i++ {
				if f.BasicFB.InternalVars.Variables[i].IsTOPIO_IN() || f.BasicFB.InternalVars.Variables[i].IsTOPIO_OUT() {
					s.InternalVars = append(s.InternalVars, f.BasicFB.InternalVars.Variables[i])
				}
			}
		}
	} else if f.CompositeFB != nil {
		for i := 0; i < len(f.CompositeFB.FBs); i++ {
			for j := 0; j < len(otherBlocks); j++ {
				if otherBlocks[j].Name == f.CompositeFB.FBs[i].Type {
					os := otherBlocks[j].GetSpecialIO(otherBlocks)
					s.InternalVars = append(s.InternalVars, os.InternalVars...)
					continue
				}
			}
		}
	}

	return s
}

func (f FB) GetDataConnectionTypes(otherBlocks []FB) ([]ConnectionWithType, error) {
	if f.CompositeFB == nil {
		return nil, nil //basic function blocks don't have dataconnections
	}

	c := f.CompositeFB
	connAndTypes := make([]ConnectionWithType, len(c.DataConnections))

conns:
	for i := 0; i < len(c.DataConnections); i++ {
		//store all connection data
		connAndTypes[i].Connection = c.DataConnections[i]

		conn := &connAndTypes[i].Connection

		found := false

		//find the type based only off the source
		if !strings.Contains(conn.Source, ".") {
			//source is from this block's parent port
			if f.InputVars != nil {
				for j := 0; j < len(f.InputVars.Variables); j++ {
					if f.InputVars.Variables[j].Name == conn.Source {
						found = true
						connAndTypes[i].Type = f.InputVars.Variables[j].Type
						continue conns
					}
				}
			}
		}

		//still here? source must be from a child block's output port
		splitSourceName := strings.Split(conn.Source, ".")
		if len(splitSourceName) != 2 {
			return nil, errors.New("Source of dataconnection '" + conn.Source + "' has an incorrect number of periods (should be 0 or 1).")
		}
		childName := splitSourceName[0]
		sourceName := splitSourceName[1]
		childType := ""

		//find the child's real block name
		for j := 0; j < len(c.FBs); j++ {
			if c.FBs[j].Name == childName {
				childType = c.FBs[j].Type
			}
		}
		if childType == "" {
			return nil, errors.New("Could not find source of dataconnection '" + conn.Source + "' as child block can't be found.")
		}

		//scan through all blocks trying to find correct API type
		for j := 0; j < len(otherBlocks); j++ {
			if otherBlocks[j].Name == childType { //matched, now scan their API
				if otherBlocks[j].InputVars != nil {
					for k := 0; k < len(otherBlocks[j].OutputVars.Variables); k++ {
						if otherBlocks[j].OutputVars.Variables[k].Name == sourceName {
							found = true
							connAndTypes[i].Type = otherBlocks[j].OutputVars.Variables[k].Type
							continue conns
						}
					}
				} else {
					return nil, errors.New("Source of dataconnection '" + conn.Source + "' has no output vars!")
				}
			}
		}

		if found == false {
			return nil, errors.New("Could not find source of dataconnection '" + conn.Source + "' in any included file.")
		}
	}

	return connAndTypes, nil
}
