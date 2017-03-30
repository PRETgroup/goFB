package iec61499

import (
	"encoding/xml"
	"errors"
	"strconv"
	"strings"
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

	ResourceVars []Variable    `xml:"VarDeclaration"` //used in resource I/O
	Resources    []FBReference `xml:"Resource"`       //used in devices

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
	Name         string `xml:"Name,attr"`
	Type         string `xml:"Type,attr"`
	ArraySize    string `xml:"ArraySize,attr,omitempty"`
	InitialValue string `xml:"InitialValue,attr,omitempty"`
	Comment      string `xml:"Comment,attr"`
}

//GetArraySize returns the array size as an integer if there is one that can be parsed, otherwise 0
func (v Variable) GetArraySize() int {
	size, err := strconv.Atoi(v.ArraySize)
	if err != nil {
		return 0
	}
	return size
}

//GetInitialArray returns a formatted initial array if there is one to do so
func (v Variable) GetInitialArray() []string {
	//if cannot parse an array size then give up
	_, err := strconv.Atoi(v.ArraySize)
	if err != nil {
		return nil
	}

	//remove everything except commas and values
	raw := v.InitialValue
	raw = strings.TrimPrefix(raw, "[")
	raw = strings.TrimSuffix(raw, "]")

	raws := strings.Split(raw, ",")
	for i := 0; i < len(raws); i++ {
		raws[i] = strings.Trim(raws[i], " ")
	}
	return raws
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

//IsLoadFor used in templates for triggering updates of input data
func (e *Event) IsLoadFor(v *Variable) bool {
	for i := 0; i < len(e.With); i++ {
		if e.With[i].Var == v.Name {
			return true
		}
	}
	return false
}

//GetUniqueEventConnSources is used to list all unique sources for event connections (useful when templating as we don't want duplicate signal wires)
func (c *CompositeFB) GetUniqueEventConnSources() []string {
	sources := make([]string, 0, len(c.EventConnections)) //preallocate for speed
nextConn:
	for i := 0; i < len(c.EventConnections); i++ {
		//check if source already found
		for j := 0; j < len(sources); j++ {
			if sources[j] == c.EventConnections[i].Source {
				continue nextConn
			}
		}
		sources = append(sources, c.EventConnections[i].Source)
	}
	return sources
}

//GetUniqueDataConnSources is used to list all unique sources for data connections (useful when templating as we don't want duplicate signal wires)
func (c *CompositeFB) GetUniqueDataConnSources() []string {
	sources := make([]string, 0, len(c.DataConnections)) //preallocate for speed
nextConn:
	for i := 0; i < len(c.DataConnections); i++ {
		//check if source already found
		for j := 0; j < len(sources); j++ {
			if sources[j] == c.DataConnections[i].Source {
				continue nextConn
			}
		}
		sources = append(sources, c.DataConnections[i].Source)
	}
	return sources
}

//ConnectionWithType struct used in cases when we want to know what a Connection's Data Type is (this can store it)
type ConnectionWithType struct {
	Connection
	Type string
}

//SourceAndType is used in similiar cases to ConnectionWithType, when we want to know both the source's name and the data type (this can store it)
type SourceAndType struct {
	Source string
	Type   string
}

//GetUniqueDataConnSourcesWithTypes is used to list all unique sources for data connections and their types (useful when templating as we don't want duplicate signal wires)
func (f FB) GetUniqueDataConnSourcesWithTypes(otherBlocks []FB) ([]SourceAndType, error) {
	cats, err := f.GetDataConnectionTypes(otherBlocks)
	if err != nil {
		return nil, err
	}

	sources := make([]SourceAndType, 0, len(cats)) //preallocate for speed
nextConn:
	for i := 0; i < len(cats); i++ {
		//check if source already found
		for j := 0; j < len(sources); j++ {
			if sources[j].Source == cats[i].Source {
				continue nextConn
			}
		}
		sources = append(sources, SourceAndType{Source: cats[i].Source, Type: cats[i].Type})
	}
	return sources, nil
}

//GetDataConnectionTypes is used to list all data connections in a given FB with their types as well (required when templating as we need to know what to define signals as)
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
				if otherBlocks[j].OutputVars != nil {
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
