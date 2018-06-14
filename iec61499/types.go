package iec61499

import (
	"encoding/xml"
)

//DebugInfo is used when valid so that if something is invalid
//we can let the user know
type DebugInfo struct {
	SourceLine int
	SourceFile string
}

//FB is the overall type for function blocks
type FB struct {
	XMLName xml.Name
	Name    string `xml:"Name,attr"`
	Comment string `xml:"Comment,attr"`

	Identification Identification

	VersionInfo VersionInfo

	CompilerInfo CompilerInfo
	//interface
	InterfaceList

	ResourceVars []Variable    `xml:"VarDeclaration"` //used in resource I/O
	Resources    []FBReference `xml:"Resource"`       //used in devices

	BasicFB     *BasicFB     `xml:",omitempty"`
	CompositeFB *CompositeFB `xml:",omitempty"`
	ServiceFB   *ServiceFB   `xml:",omitempty"`
	HybridFB    *HybridFB    `xml:"-"`          //don't ever export HybridFBs, convert them to BFBs first
	Policies    []PolicyFB   `xml:",omitempty"` //PolicyFB can wrap the execution of a given block

	NumChildren int `xml:"-"` //this is useful when using the event queue as we use it to assign unique blockInstanceIDs, it stores the recursive number of children a block has
	DebugInfo   `xml:"-"`
}

//InterfaceList is a container for all the IO ports of a Function Block
type InterfaceList struct {
	EventInputs  []Event    `xml:"InterfaceList>EventInputs>Event,omitempty"`
	EventOutputs []Event    `xml:"InterfaceList>EventOutputs>Event,omitempty"`
	InputVars    []Variable `xml:"InterfaceList>InputVars>VarDeclaration,omitempty"`
	OutputVars   []Variable `xml:"InterfaceList>OutputVars>VarDeclaration,omitempty"`
}

//HasIONamed will check a given InterfaceList to see if it has an output of that name
func (il InterfaceList) HasIONamed(input bool, s string) bool {
	if il.HasEventNamed(input, s) {
		return true
	}
	if il.HasIODataNamed(input, s) {
		return true
	}
	return false
}

//HasEventNamed will check a given InterfaceList to see if it has an event of that name
func (il InterfaceList) HasEventNamed(input bool, s string) bool {
	if input {
		for i := 0; i < len(il.EventInputs); i++ {
			if il.EventInputs[i].Name == s {
				return true
			}
		}
		return false
	}
	for i := 0; i < len(il.EventOutputs); i++ {
		if il.EventOutputs[i].Name == s {
			return true
		}
	}
	return false
}

//HasIODataNamed will check a given InterfaceList to see if it has an IO data var of that name
func (il InterfaceList) HasIODataNamed(input bool, s string) bool {
	if input {
		for i := 0; i < len(il.InputVars); i++ {
			if il.InputVars[i].Name == s {
				return true
			}
		}
		return false
	}
	for i := 0; i < len(il.OutputVars); i++ {
		if il.OutputVars[i].Name == s {
			return true
		}
	}
	return false
}

//IsSIFB returns true if the FB is an interface only (i.e. is a SIFB)
func (f FB) IsSIFB() bool {
	if f.BasicFB == nil && f.CompositeFB == nil && len(f.ResourceVars) == 0 && len(f.Resources) == 0 {
		return true
	}
	return false
}

//Identification identifies what version of the standard is being used
type Identification struct {
	Standard string `xml:"Standard,attr"`
}

//CompilerInfo is used when compiling by some 61499 compilers (but not goFB)
type CompilerInfo struct {
	Header   string `xml:"header,attr"`
	Classdef string `xml:"classdef,attr"`
}

//VersionInfo is used to provide metadata about the fb's version
type VersionInfo struct {
	Organization string `xml:"Organization,attr"`
	Version      string `xml:"Version,attr"`
	Author       string `xml:"Author,attr"`
	Date         string `xml:"Date,attr"`
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

	DebugInfo `xml:"-"`
}

//Parameter is a custom values associated with a FBReferences
type Parameter struct {
	Name  string `xml:"Name,attr"`
	Value string `xml:"Value,attr"`

	DebugInfo `xml:"-"`
}

//Connection is used to link either events or data lines together inside CompositeFBs
type Connection struct {
	Source      string `xml:"Source,attr"`
	Destination string `xml:"Destination,attr"`
	Dx1         string `xml:"dx1,attr"`

	DebugInfo `xml:"-"`
}

//BasicFB is used for basic function blocks (ie FB definitions)
type BasicFB struct {
	InternalVars []Variable     `xml:"InternalVars>VarDeclaration,omitempty"`
	States       []ECState      `xml:"ECC>ECState"`
	Transitions  []ECTransition `xml:"ECC>ECTransition,omitempty"`
	Algorithms   []Algorithm    `xml:"Algorithm,omitempty"`
}

//HybridFB is used for Hybrid Function Blocks, which are translated to Basic Function Blocks before export
type HybridFB struct {
	InternalVars []Variable     `xml:"-"`
	Locations    []HFBLocation  `xml:"-"`
	Transitions  []ECTransition `xml:"-"`
	Algorithms   []Algorithm    `xml:"-"`
}

//HFBLocation is a location in the HFB HA (hybrid automata state machine) of a HybridFB
type HFBLocation struct {
	Name       string
	ECActions  []Action
	Comment    string
	Invariants []HFBInvariant

	DebugInfo
}

//HFBInvariant is used to store invariant conditions in HFB locations
//The invariant will be in the form "variable [operand] value [and/or] ..."
type HFBInvariant struct {
	Invariant string

	DebugInfo
}

//PolicyFB is used for specifying policies that must be kept
type PolicyFB struct {
	Name         string
	InternalVars []Variable      `xml:"InternalVars>VarDeclaration,omitempty"`
	States       []PFBState      `xml:"ECC>ECState"`
	Transitions  []PFBTransition `xml:"ECC>ECTransition,omitempty"`
}

//PFBState is a state in the policy specification of an enforcerFB
type PFBState struct {
	Name      string
	DebugInfo `xml:"-"`
}

//PFBTransition is a transition between PFBStates in an PolicyFB (mealy machine transitions)
type PFBTransition struct {
	ECTransition
	Expressions []PFBExpression //expressions associated with this transition
}

//PFBExpression is used to assign a var a value based on a PFBTransitions
type PFBExpression struct {
	VarName string
	Value   string
}

//ECState is a state in the ECC (Execution control chart) of a BasicFB
type ECState struct {
	Name      string   `xml:"Name,attr"`
	ECActions []Action `xml:"ECAction,omitempty"`
	Comment   string   `xml:"Comment,attr"`
	X         string   `xml:"x,attr"`
	Y         string   `xml:"y,attr"`

	DebugInfo `xml:"-"`
}

//Action is a link between an ECState and an Algorithm
type Action struct {
	Algorithm string `xml:"Algorithm,attr,omitempty"`
	Output    string `xml:"Output,attr,omitempty"`

	DebugInfo `xml:"-"`
}

//ECTransition is a transition in the ECC of a BasicFB
type ECTransition struct {
	Source      string `xml:"Source,attr"`
	Destination string `xml:"Destination,attr"`
	Condition   string `xml:"Condition,attr"`
	X           string `xml:"x,attr"`
	Y           string `xml:"y,attr"`

	DebugInfo `xml:"-"`
}

//Algorithm is some code associated with a BasicFB
type Algorithm struct {
	Name    string        `xml:"Name,attr"`
	Comment string        `xml:"Comment,attr"`
	Other   OtherLanguage `xml:",omitEmpty"`

	DebugInfo `xml:"-"`
}

//OtherLanguage states what language an algorithm is written in
type OtherLanguage struct {
	XMLName  xml.Name `xml:"Other"`
	Language string   `xml:"Language,attr"`
	Text     string   `xml:"Text,attr"`
}

//Variable is used as a variable in an algorithm of a BasicFB
type Variable struct {
	Name         string `xml:"Name,attr"`
	Type         string `xml:"Type,attr"`
	ArraySize    string `xml:"ArraySize,attr,omitempty"`
	InitialValue string `xml:"InitialValue,attr,omitempty"`
	Comment      string `xml:"Comment,attr"`

	DebugInfo `xml:"-"`
}

//Event is used to store events in BasicFBs
type Event struct {
	Name    string `xml:"Name,attr"`
	Comment string `xml:"Comment,attr"`
	With    []With `xml:",omitEmpty"`

	DebugInfo `xml:"-"`
}

//With associates a variable and an event (as data only changes on events)
type With struct {
	Var string `xml:"Var,attr"`
}

//ServiceFB is used for compatible SIFBs
type ServiceFB struct {
	XMLName      xml.Name                 `xml:"Service"`
	Autogenerate *ServiceAutogenerateCode `xml:",omitempty"` //set fields in this we want to autogenerate files that is goFB compatible using the below *Texts
}

//ServiceAutogenerateCode is used to store autogenerating code information for ServiceFBs
// it is not technically part of the standard, but should be safely ignored by tools that don't use this package
type ServiceAutogenerateCode struct {
	Language string `xml:",attr"` //Language of all *Texts below

	ArbitraryText string `xml:",attr"` //When autogenerating, put this outside of functions/structs at the beginning of the .c file
	InStructText  string `xml:",attr"` //When autogenerating, put this in the struct
	PreInitText   string `xml:",attr"` //When autogenerating, put this in pre_init function
	InitText      string `xml:",attr"` //When autogenerating, put this in init function
	RunText       string `xml:",attr"` //When autogenerating, put this in run function
	ShutdownText  string `xml:",attr"` //When autogenerating, put this in shutdown function

	DebugInfo `xml:"-"`
}
