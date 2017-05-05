package iec61499

import "encoding/xml"

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
