package main

import "errors"

//Enforcer is a container that will hold all things to enforce
type Enforcer struct {
	Name     string
	IO       IO
	Policies []Policy
}

//IO is a container for enforcers that holds the input parameters, as well
//as the definition of the data lines to enforce
type IO struct {
	Inputs  []IOLine
	Enforce []IOLine
}

//IOLine contains a single line of IO, either a parameter, or an enforcement line
type IOLine struct {
	Name    string
	Type    string
	Initial string
}

//A Policy is a top level ObserveNode with no entry condition (i.e. it runs all the time)
type Policy struct {
	Name         string
	Triggers     []Trigger
	Requirements []Requirement
}

//A Trigger represents the concept of timing something when considering timed requirements,
//it is implemented as a named register in HW that will be used to store the current time when a start condition is met
//there will be a start/stop flag that indicates whether a given Trigger should be considered "active"
type Trigger struct {
	Name           string
	StartCondition Operation
	ResetCondition *Operation
}

//What is the difference between `before`, `until`, and `after` for operations?
// * `before`: Requirement needs to only happen once, before or at the condition
// * `until`: Requirement needs to happen continuously, up to the condition
// * `after`: Requirement needs to be sustained permanently after the condition

//A Requirement is the concrete representation of a policy to be implemented
type Requirement struct {
	With     []string //can be empty, but must be the names of Triggers
	Requires []Operation
	Recover  []Operation
}

//FindIO is used on an enforcer to find and classify the type of IO for a given name
func (e Enforcer) FindIO(name string) (inputsIndex int, enforceIndex int) {
	inputsIndex = -1
	enforceIndex = -1
	for i := 0; i < len(e.IO.Enforce); i++ {
		if e.IO.Enforce[i].Name == name {
			enforceIndex = i
			return
		}
	}
	for i := 0; i < len(e.IO.Inputs); i++ {
		if e.IO.Inputs[i].Name == name {
			inputsIndex = i
			return
		}
	}
	return
}

//GetIO gets all the IO used in a given policy ObserveNode (which will be a subset of the enforcer IO in general)
func (e Enforcer) GetIO(p Policy) (inputs []IOLine, enforce []IOLine, Triggers []string, err error) {
	inputs = make([]IOLine, 0, 10)
	enforce = make([]IOLine, 0, 10)
	Triggers = make([]string, len(p.Triggers))
	for i := 0; i < len(p.Triggers); i++ {
		Triggers[i] = p.Triggers[i].Name
	}
	err = nil

	var vars []string

	//get the IO for all Triggers & policies in this child observations and append if it is unique
	for i := 0; i < len(p.Triggers); i++ {
		vars = append(vars, p.Triggers[i].StartCondition.GetVariables()...)
		vars = append(vars, p.Triggers[i].StartCondition.GetVariables()...)
	}

	for i := 0; i < len(p.Requirements); i++ {
		for j := 0; j < len(p.Requirements[i].Requires); j++ {
			vars = append(vars, p.Requirements[i].Requires[j].GetVariables()...)
		}

		for j := 0; j < len(p.Requirements[i].Recover); j++ {
			vars = append(vars, p.Requirements[i].Recover[j].GetVariables()...)
		}
	}

	for j := 0; j < len(vars); j++ {
		//first make sure it isn't a Trigger
		found := false
		for i := 0; i < len(p.Triggers); i++ {
			if vars[j] == p.Triggers[i].Name {
				found = true
			}
		}

		if found == false {
			//get type of IO
			inputsIndex, enforceIndex := e.FindIO(vars[j])
			if inputsIndex < 0 && enforceIndex < 0 {
				err = errors.New("Couldn't find IO " + vars[j])
				return
			}

			//append to appropriate place, if unique
			if inputsIndex >= 0 {
				input := e.IO.Inputs[inputsIndex]
				found := false
				for k := 0; k < len(inputs); k++ {
					if inputs[k].Name == input.Name {
						found = true
						break
					}
				}
				if found == false {
					inputs = append(inputs, input)
				}
			} else {
				enforc := e.IO.Enforce[enforceIndex]
				found := false
				for k := 0; k < len(enforce); k++ {
					if enforce[k].Name == enforc.Name {
						found = true
						break
					}
				}
				if found == false {
					enforce = append(enforce, enforc)
				}
			}
		}
	}

	return
}

//OperatorXXXX constants for use in classifying the possible types of operations
const (
	OperatorNothing  = ""
	OperatorEquals   = "="
	OperatorNEquals  = "/="
	OperatorValue    = "#"
	OperatorVariable = "$"
	OperatorAfter    = "after"
	OperatorBefore   = "before"
	OperatorUntil    = "until"
	OperatorSet      = ":="
	OperatorAdd      = "+"
	OperatorSubtract = "-"
	OperatorMultiply = "*"
	OperatorDivide   = "/"
	OperatorAnd      = "and"
	OperatorOr       = "or"
	OperatorGT       = ">"
	OperatorGTE      = ">="
	OperatorLT       = "<"
	OperatorLTE      = "<="
)

type Operation struct {
	Type  string
	Value string
	A     *Operation
	B     *Operation
}

func (o Operation) IsValue() bool {
	return o.Type == OperatorValue
}

func (o Operation) IsVariable() bool {
	return o.Type == OperatorVariable
}

func (o Operation) IsAfter() bool {
	return o.Type == OperatorAfter
}

func (o Operation) IsBefore() bool {
	return o.Type == OperatorBefore
}

func (o Operation) IsUntil() bool {
	return o.Type == OperatorUntil
}

func (o Operation) IsSet() bool {
	return o.Type == OperatorSet
}

//GetVariables finds all variables referred to in a given operation
//It traverses its tree recursively
func (o Operation) GetVariables() []string {
	if o.Type == OperatorVariable {
		return []string{o.Value}
	}
	if o.Type == OperatorValue {
		return nil
	}
	var varsA []string
	var varsB []string
	if o.A != nil {
		varsA = o.A.GetVariables()
	}
	if o.B != nil {
		varsB = o.B.GetVariables()
	}
	return append(varsA, varsB...)
}
