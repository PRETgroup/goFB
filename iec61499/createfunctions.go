package iec61499

import (
	"encoding/xml"
	"errors"
	"strings"
)

//FBError is used to pass helpful error messages out to clients
type FBError struct {
	Arg string
	Err error
}

func (f FBError) error() string {
	if f.Arg != "" {
		return f.Err.Error() + " (" + f.Arg + ")"
	}
	return f.Err.Error()
}

var (
	//ErrUndefinedEvent is used to indicate an event can't be found
	ErrUndefinedEvent = errors.New("Unknown event")

	//ErrNoEvents is used to indicate no events are present
	ErrNoEvents = errors.New("No Events Present")

	//ErrWrongBlockType is used to indicate the wrong type of block has been used in an operation (ie adding an internal var to a compositefb)
	ErrWrongBlockType = errors.New("Wrong Block Type for this operation")

	//ErrFbInstNameUndefined is used to indicate that an fb block instance with that name can't be found
	ErrFbInstNameUndefined = errors.New("FB instance name undefined")
)

func newError(err error, arg string) *FBError {
	return &FBError{
		Arg: arg,
		Err: err,
	}
}

//NewBasicFB returns a BasicFB with default fields filled
func NewBasicFB(name string) *FB {
	fb := FB{}
	fb.Name = name
	fb.Identification.Standard = "61499-2"
	fb.BasicFB = new(BasicFB)
	fb.SetXMLName()
	return &fb
}

//NewCompositeFB returns a CompositeFB with default fields filled
func NewCompositeFB(name string) *FB {
	fb := FB{}
	fb.Name = name
	fb.Identification.Standard = "61499-2"
	fb.CompositeFB = new(CompositeFB)
	fb.SetXMLName()
	return &fb
}

//NewServiceFB returns a ServiceFB with default fields filled
func NewServiceFB(name string) *FB {
	fb := FB{}
	fb.Name = name
	fb.Identification.Standard = "61499-2"
	fb.ServiceFB = new(ServiceFB)
	fb.SetXMLName()
	return &fb
}

//NewHybridFB returns a HybridFB with default fields filled
func NewHybridFB(name string) *FB {
	fb := FB{}
	fb.Name = name
	fb.Identification.Standard = "61499-2"
	fb.HybridFB = new(HybridFB)
	fb.SetXMLName()
	return &fb
}

//SetXMLName sets an appropriate name for the xml block type
func (fb *FB) SetXMLName() {
	fb.XMLName = xml.Name{Space: "", Local: "FBType"}
}

//NameUseCounter will search all name fields {events,data,states,etc} of a FB to check to see if a given name is used anywhere and increment each time it is
func (fb *FB) NameUseCounter(name string) int {
	usages := 0
	//check top level name
	if fb.Name == name {
		usages++
	}

	//check all of the interface names

	for _, e := range fb.EventInputs {
		if e.Name == name {
			usages++
		}
	}

	for _, e := range fb.EventOutputs {
		if e.Name == name {
			usages++
		}
	}

	for _, e := range fb.InputVars {
		if e.Name == name {
			usages++
		}
	}

	for _, e := range fb.OutputVars {
		if e.Name == name {
			usages++
		}
	}

	//if it is BFB, check all internals,states, and algorithms
	if fb.BasicFB != nil {

		for _, e := range fb.BasicFB.InternalVars {
			if e.Name == name {
				usages++
			}
		}

		for _, s := range fb.BasicFB.States {
			if s.Name == name {
				usages++
			}
		}
		for _, a := range fb.BasicFB.Algorithms {
			if a.Name == name {
				usages++
			}
		}
	}

	//if it is CFB, check all instances
	if fb.CompositeFB != nil {
		for _, e := range fb.CompositeFB.FBs {
			if e.Name == name {
				usages++
			}
		}
	}

	//if it is Device, check all resources
	for _, e := range fb.Resources {
		if e.Name == name {
			usages++
		}
	}
	return usages
}

//AddEventInputNames adds event input names to a FB
func (fb *FB) AddEventInputNames(names []string, debug DebugInfo) *FB {

	for _, name := range names {
		fb.EventInputs = append(fb.EventInputs, Event{Name: name, DebugInfo: debug})
	}
	return fb
}

//AddEventOutputNames adds event output names to a FB
func (fb *FB) AddEventOutputNames(names []string, debug DebugInfo) *FB {

	for _, name := range names {
		fb.EventOutputs = append(fb.EventOutputs, Event{Name: name, DebugInfo: debug})
	}
	return fb
}

//Must is used to panic if an error occurs when chaining functions
func Must(f *FB, err *FBError) *FB {
	if err != nil {
		panic(err)
	}
	return f
}

//AddDataInputs adds data inputs to an FB
// it will return an error message if a trigger can't be found
func (fb *FB) AddDataInputs(intNames []string, intTriggers []string, typ string, size string, initialValue string, debug DebugInfo) (*FB, *FBError) {
	typ = strings.ToUpper(typ)

	for _, iname := range intNames {
		fb.InputVars = append(fb.InputVars, Variable{Name: iname, Type: typ, ArraySize: size, InitialValue: initialValue, DebugInfo: debug})

	inTrig:
		for _, tname := range intTriggers {
			for i := 0; i < len(fb.EventInputs); i++ {
				if fb.EventInputs[i].Name == tname {
					fb.EventInputs[i].With = append(fb.EventInputs[i].With, With{Var: iname})
					continue inTrig
				}
			}
			return nil, newError(ErrUndefinedEvent, tname)
		}
	}

	return fb, nil
}

//AddDataOutputs adds data inputs to an FB
// it will return an error message if a trigger can't be found
func (fb *FB) AddDataOutputs(intNames []string, intTriggers []string, typ string, size string, initialValue string, debug DebugInfo) (*FB, *FBError) {
	typ = strings.ToUpper(typ)

	for _, iname := range intNames {
		fb.OutputVars = append(fb.OutputVars, Variable{Name: iname, Type: typ, ArraySize: size, InitialValue: initialValue, DebugInfo: debug})
		if len(intTriggers) != 0 && fb.EventOutputs == nil {
			return nil, newError(ErrNoEvents, fb.Name+" has no output events to associate with")
		}
	outTrig:
		for _, tname := range intTriggers {
			for i := 0; i < len(fb.EventOutputs); i++ {
				if fb.EventOutputs[i].Name == tname {
					fb.EventOutputs[i].With = append(fb.EventOutputs[i].With, With{Var: iname})
					continue outTrig
				}
			}
			return nil, newError(ErrUndefinedEvent, tname)
		}
	}
	return fb, nil
}

//AddxFBDataInternals adds data internals to an fb's bfb OR hfb
//if a block is neither an hfb or a bfb then it returns an error
func (fb *FB) AddxFBDataInternals(intNames []string, typ string, size string, initialValue string, debug DebugInfo) (*FB, error) {
	typ = strings.ToUpper(typ)

	if fb.BasicFB != nil {
		fb.BasicFB.AddDataInternals(intNames, typ, size, initialValue, debug)
		return fb, nil
	}
	if fb.HybridFB != nil {
		fb.HybridFB.AddDataInternals(intNames, typ, size, initialValue, debug)
		return fb, nil
	}
	return nil, errors.New("AddxFBDataInternals may only be called on HFBs or BFBs")
}

//AddBFBDataInternals adds data internals to an fb's bfb without performing error checking
func (fb *FB) AddBFBDataInternals(intNames []string, typ string, size string, initialValue string, debug DebugInfo) *FB {
	fb.BasicFB.AddDataInternals(intNames, typ, size, initialValue, debug)
	return fb
}

//AddDataInternals adds data internals to a bfb, and adds the InternalVars section if it is nil
func (bfb *BasicFB) AddDataInternals(intNames []string, typ string, size string, initialValue string, debug DebugInfo) *BasicFB {

	for _, iname := range intNames {
		bfb.InternalVars = append(bfb.InternalVars, Variable{Name: iname, Type: typ, ArraySize: size, InitialValue: initialValue, DebugInfo: debug})
	}
	return bfb
}

//AddBFBAlgorithm adds algorithm to an fb's bfb without performing error checking
func (fb *FB) AddBFBAlgorithm(name string, lang string, prog string, debug DebugInfo) *FB {
	fb.BasicFB.AddAlgorithm(name, lang, prog, debug)
	return fb
}

//AddAlgorithm adds an algorithm to a bfb
// it will return an error message if the block is not a basicFB
func (bfb *BasicFB) AddAlgorithm(name string, lang string, prog string, debug DebugInfo) *BasicFB {

	bfb.Algorithms = append(bfb.Algorithms, Algorithm{
		Name: name,
		Other: OtherLanguage{
			Language: lang,
			Text:     prog,
		},
		DebugInfo: debug,
	})

	return bfb
}

//AddBFBState adds state to an fb's bfb without performing error checking
func (fb *FB) AddBFBState(name string, actions []Action, debug DebugInfo) *FB {
	fb.BasicFB.AddState(name, actions, debug)
	return fb
}

//AddState adds a state to a bfb
func (bfb *BasicFB) AddState(name string, actions []Action, debug DebugInfo) *BasicFB {
	bfb.States = append(bfb.States, ECState{
		Name:      name,
		ECActions: actions,
		DebugInfo: debug,
	})
	return bfb
}

//AddBFBTransition adds a transition to an fb's bfb without performing error checking
func (fb *FB) AddBFBTransition(source string, dest string, cond string, debug DebugInfo) *FB {
	fb.BasicFB.AddTransition(source, dest, cond, debug)
	return fb
}

//AddTransition adds a state transition to a bfb
func (bfb *BasicFB) AddTransition(source string, dest string, cond string, debug DebugInfo) *BasicFB {
	bfb.Transitions = append(bfb.Transitions, ECTransition{
		Source:      source,
		Destination: dest,
		Condition:   cond,

		DebugInfo: debug,
	})
	return bfb
}

//AddCFBInstances adds an instance to a fb cfb without performing error checking
func (fb *FB) AddCFBInstances(fbTypeName string, fbInstNames []string, debug DebugInfo) *FB {
	fb.CompositeFB.AddInstances(fbTypeName, fbInstNames, debug)
	return fb
}

//AddInstances adds an instance to a cfb
func (cfb *CompositeFB) AddInstances(fbTypeName string, fbInstNames []string, debug DebugInfo) *CompositeFB {
	for _, fbInstName := range fbInstNames {
		cfb.FBs = append(cfb.FBs, FBReference{
			Name:      fbInstName,
			Type:      fbTypeName,
			DebugInfo: debug,
		})
	}
	return cfb
}

//AddCFBNetworkEventConns adds event network connections from many sources to a single destination in a fb cfb without performing error checking
func (fb *FB) AddCFBNetworkEventConns(sources []string, destination string, debug DebugInfo) *FB {
	fb.CompositeFB.AddNetworkEventConns(sources, destination, debug)
	return fb
}

//AddNetworkEventConns adds event network connections from many sources to a single destination in a cfb
func (cfb *CompositeFB) AddNetworkEventConns(sources []string, destination string, debug DebugInfo) *CompositeFB {
	for _, source := range sources {
		cfb.EventConnections = append(cfb.EventConnections, Connection{
			Source:      source,
			Destination: destination,
			DebugInfo:   debug,
		})
	}
	return cfb
}

//AddCFBNetworkDataConn adds data network conenctions in a fb cfb without performing error checking
func (fb *FB) AddCFBNetworkDataConn(source string, destination string, debug DebugInfo) *FB {
	fb.CompositeFB.AddNetworkDataConn(source, destination, debug)
	return fb
}

//AddNetworkDataConn adds data network connections from one source to one destination
func (cfb *CompositeFB) AddNetworkDataConn(source string, destination string, debug DebugInfo) *CompositeFB {
	cfb.DataConnections = append(cfb.DataConnections, Connection{
		Source:      source,
		Destination: destination,
		DebugInfo:   debug,
	})
	return cfb
}

//AddCFBNetworkParameter adds adds a parameter to a cfb network without performing error checking
func (fb *FB) AddCFBNetworkParameter(param string, fbInstName string, portName string, debug DebugInfo) *FB {
	fb.CompositeFB.AddNetworkParameter(param, fbInstName, portName, debug)
	return fb
}

//AddNetworkParameter adds a parameter to a cfb network. It returns an error if the fbInstName is undefined
// (as it needs it to store the parameter)
func (cfb *CompositeFB) AddNetworkParameter(param string, fbInstName string, portName string, debug DebugInfo) (*CompositeFB, *FBError) {
	for i := 0; i < len(cfb.FBs); i++ {
		if cfb.FBs[i].Name == fbInstName {
			cfb.FBs[i].Parameter = append(cfb.FBs[i].Parameter, Parameter{
				Name:  portName,
				Value: param,

				DebugInfo: debug,
			})
			return nil, nil
		}
	}
	return nil, newError(ErrFbInstNameUndefined, fbInstName)
}

//AddSIFBParams adds all parameters to an SIFB. It returns no error
func (fb *FB) AddSIFBParams(lang string, arbitrary string, inStruct string, preInit string, init string, run string, shutdown string, debug DebugInfo) *FB {
	fb.ServiceFB.AddParams(lang, arbitrary, inStruct, preInit, init, run, shutdown, debug)
	return fb
}

//AddParams adds all parameters to an autogenerating SIFB
func (sifb *ServiceFB) AddParams(lang string, arbitrary string, inStruct string, preInit string, init string, run string, shutdown string, debug DebugInfo) *ServiceFB {
	sifb.Autogenerate = new(ServiceAutogenerateCode)
	sifb.Autogenerate.Language = lang
	sifb.Autogenerate.ArbitraryText = arbitrary
	sifb.Autogenerate.InStructText = inStruct
	sifb.Autogenerate.PreInitText = preInit
	sifb.Autogenerate.InitText = init
	sifb.Autogenerate.RunText = run
	sifb.Autogenerate.ShutdownText = shutdown
	sifb.Autogenerate.DebugInfo = debug
	return sifb
}

//AddHFBDataInternals adds data internals to an fb's bfb without performing error checking
func (fb *FB) AddHFBDataInternals(intNames []string, typ string, size string, initialValue string, debug DebugInfo) *FB {
	fb.HybridFB.AddDataInternals(intNames, typ, size, initialValue, debug)
	return fb
}

//AddDataInternals adds data internals to a bfb, and adds the InternalVars section if it is nil
func (hfb *HybridFB) AddDataInternals(intNames []string, typ string, size string, initialValue string, debug DebugInfo) *HybridFB {

	for _, iname := range intNames {
		hfb.InternalVars = append(hfb.InternalVars, Variable{Name: iname, Type: typ, ArraySize: size, InitialValue: initialValue, DebugInfo: debug})
	}
	return hfb
}

//AddHFBAlgorithm adds algorithm to an fb's hfb without performing error checking
func (fb *FB) AddHFBAlgorithm(name string, prog string, debug DebugInfo) *FB {
	fb.HybridFB.AddAlgorithm(name, prog, debug)
	return fb
}

//AddAlgorithm adds an algorithm to a hfb
func (hfb *HybridFB) AddAlgorithm(name string, prog string, debug DebugInfo) *HybridFB {

	hfb.Algorithms = append(hfb.Algorithms, Algorithm{
		Name: name,
		Other: OtherLanguage{
			Language: "ODE",
			Text:     prog,
		},
		DebugInfo: debug,
	})

	return hfb
}

//AddHFBLocation adds location to an fb's hfb without performing error checking
func (fb *FB) AddHFBLocation(name string, invariants []HFBInvariant, actions []Action, debug DebugInfo) *FB {
	fb.HybridFB.AddLocation(name, invariants, actions, debug)
	return fb
}

//AddLocation adds a location to a hfb
func (hfb *HybridFB) AddLocation(name string, invariants []HFBInvariant, actions []Action, debug DebugInfo) *HybridFB {
	hfb.Locations = append(hfb.Locations, HFBLocation{
		Name:       name,
		Invariants: invariants,
		ECActions:  actions,
		DebugInfo:  debug,
	})
	return hfb
}

//AddHFBTransition adds a transition to an fb's bfb without performing error checking
func (fb *FB) AddHFBTransition(source string, dest string, cond string, debug DebugInfo) *FB {
	fb.HybridFB.AddTransition(source, dest, cond, debug)
	return fb
}

//AddTransition adds a state transition to a bfb
func (hfb *HybridFB) AddTransition(source string, dest string, cond string, debug DebugInfo) *HybridFB {
	hfb.Transitions = append(hfb.Transitions, ECTransition{
		Source:      source,
		Destination: dest,
		Condition:   cond,

		DebugInfo: debug,
	})
	return hfb
}
