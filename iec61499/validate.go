package iec61499

import (
	"errors"
	"fmt"
	"strings"
)

//A ValidateError is returned when validation of a network of function block fails
type ValidateError struct {
	Err error
	Arg string
	DebugInfo
}

func (v ValidateError) Error() string {
	s := fmt.Sprintf("Error (File %v, Line %v): %s", v.SourceFile, v.SourceLine, v.Err.Error())
	s = strings.Replace(s, "{{arg}}", v.Arg, -1)
	return s
}

func newValidateError(err error, arg string, debug DebugInfo) *ValidateError {
	return &ValidateError{
		Err:       err,
		Arg:       arg,
		DebugInfo: debug,
	}
}

var (
	//ErrFbTypeNameNotUnique is used to indicate when a FBType name isn't unique in the set
	ErrFbTypeNameNotUnique = errors.New("FBType Name '{{arg}}' is not unique in the set")

	//ErrInterfaceNameNotUnique is used to indicate when an interface name isn't unique in a block
	ErrInterfaceNameNotUnique = errors.New("Interface name '{{arg}}' is not unique in the FB")

	//ErrUndefinedAssociatedDataPort is used when an associated data port (to an event) can't be located
	ErrUndefinedAssociatedDataPort = errors.New("The associated data port '{{arg}}' can't be found")

	//ErrInterfaceDataTypeNotValid is used to indicate when an interface data type isn't valid
	ErrInterfaceDataTypeNotValid = errors.New("Interface data type '{{arg}}' is not valid")

	//ErrCompositeBFBsDontGetEventDataAssociations is used to indicate when a CFB block thas been given event/data associations
	ErrCompositeBFBsDontGetEventDataAssociations = errors.New("Composite FBs can't have event/data associations")

	//ErrUndefinedAlgorithm is used to indicate that an algorithm was referenced that can't be found (so probably a typo has occured)
	ErrUndefinedAlgorithm = errors.New("Can't find Algorithm with name '{{arg}}'")

	//ErrUndefinedState is used to indicate that a state was referenced that can't be found (so probably a typo has occured)
	ErrUndefinedState = errors.New("Can't find State with name '{{arg}}'")

	//ErrOnlyBasicFBsGetTriggers is returned when triggers (event/data associations) are put on anything other than a BasicFB
	ErrOnlyBasicFBsGetTriggers = errors.New("Only Basic FBs can have event/data associations/triggers")

	//ErrStateNameNotUnique is returned when a state name in a bfb is not unique
	ErrStateNameNotUnique = errors.New("The state name '{{arg}}' is not unique in the FB")

	//ErrFbReferenceInvalidType is returned when an embedded child type in a CFB can't be found
	ErrFbReferenceInvalidType = errors.New("The FBType of the FBReference (the instance type of an embedded FB) '{{arg}}' can't be found")

	//ErrFbReferenceNameNotUnique is returned when an instance in the CFB has a non-unique name
	ErrFbReferenceNameNotUnique = errors.New("The FBReference Name '{{arg}}' of the instance is not unique in the CFB")
)

//ValidateFBs will take a slice of FBs and check that they are valid
//It will check for the following, in this order:
// ALL FBs:
// - Type name unique in set
// - Interface names are unique in fb
// - Interface data types are valid
// - Interface associations (if present) are valid in fb (ie only bfbs get associations)
// BASIC FBs:
// - Internal data names don't conflict with interface names
// - Internal data types are valid
// COMPOSITE FBs:
// - CFB instance types exist
// - CFB instance names unique
// - CFB event/data connections valid based on instances
func ValidateFBs(fbs []FB) *ValidateError {
	for i := 0; i < len(fbs); i++ {
		oFBs := make([]FB, len(fbs))
		copy(oFBs, fbs)
		fb := fbs[i]
		if err := validateFB(fb, append(oFBs[0:i], oFBs[i+1:]...)); err != nil {
			return err
		}
	}
	return nil
}

func validateFB(fb FB, otherFbs []FB) *ValidateError {
	if err := fbTypeNamesUnique(fb, otherFbs); err != nil {
		return err
	}

	if err := fbInterfaceNamesUniqueAndValidTypesAndAssociationsValid(fb, otherFbs); err != nil {
		return err
	}

	if err := fbBFBStateMachineValid(fb, otherFbs); err != nil {
		return err
	}

	if err := fbCFBNetworkValid(fb, otherFbs); err != nil {
		return err
	}

	//check that the type name is not in the otherFbs
	return nil
}

//fbTypeNamesUnique makes sure each
func fbTypeNamesUnique(fb FB, otherFbs []FB) *ValidateError {
	for _, oFb := range otherFbs {
		if oFb.Name == fb.Name {
			return newValidateError(ErrFbTypeNameNotUnique, fb.Name, fb.DebugInfo)
		}
	}
	return nil
}

//fbInterfaceNamesUniqueAndValidTypesAndAssociationsValid checks
// -that interface names don't overlap within a fb
// -that types are valid
// -that no associations are present except in bfbs
// -that any associations are with existing inputs
// (we don't use otherFbs in this function)
func fbInterfaceNamesUniqueAndValidTypesAndAssociationsValid(fb FB, otherFbs []FB) *ValidateError {
	for _, e := range fb.EventInputs {
		if fb.NameUseCounter(e.Name) > 1 {
			return newValidateError(ErrInterfaceNameNotUnique, e.Name, e.DebugInfo)
		}
		if len(e.With) > 0 && fb.CompositeFB != nil {
			return newValidateError(ErrCompositeBFBsDontGetEventDataAssociations, e.Name, e.DebugInfo)
		}
		for _, w := range e.With {
			found := false
			for _, v := range fb.InputVars {
				if w.Var == v.Name {
					found = true
					break
				}
			}
			if found == false {
				return newValidateError(ErrUndefinedAssociatedDataPort, e.Name, e.DebugInfo)
			}
		}
	}
	for _, e := range fb.EventOutputs {
		if fb.NameUseCounter(e.Name) > 1 {
			return newValidateError(ErrInterfaceNameNotUnique, e.Name, e.DebugInfo)
		}
		if len(e.With) > 0 && fb.CompositeFB != nil {
			return newValidateError(ErrCompositeBFBsDontGetEventDataAssociations, e.Name, e.DebugInfo)
		}
		for _, w := range e.With {
			found := false
			for _, v := range fb.OutputVars {
				if w.Var == v.Name {
					found = true
					break
				}
			}
			if found == false {
				return newValidateError(ErrUndefinedAssociatedDataPort, e.Name, e.DebugInfo)
			}
		}
	}

	for _, v := range append(fb.InputVars, fb.OutputVars...) {
		if fb.NameUseCounter(v.Name) > 1 {
			return newValidateError(ErrInterfaceNameNotUnique, v.Name, v.DebugInfo)
		}
		if !isValidDataType(v.Type) {
			return newValidateError(ErrInterfaceDataTypeNotValid, v.Type, v.DebugInfo)
		}
	}

	return nil
}

func fbBFBStateMachineValid(fb FB, otherFBs []FB) *ValidateError {
	if fb.BasicFB == nil { //blocks that aren't basicFBs aren't going to have invalid state machines!
		return nil
	}

	for _, s := range fb.BasicFB.States {
		//make sure all state names are unique
		if fb.NameUseCounter(s.Name) > 1 {
			return newValidateError(ErrStateNameNotUnique, s.Name, s.DebugInfo)
		}

		//make sure the ecActions are valid
		for _, a := range s.ECActions {
			if a.Algorithm != "" {
				found := false
				for _, al := range fb.BasicFB.Algorithms {
					if al.Name == a.Algorithm {
						found = true
						break
					}
				}
				if found == false {
					return newValidateError(ErrUndefinedAlgorithm, a.Algorithm, a.DebugInfo)
				}

			}

			if a.Output != "" {
				found := false
				for _, eo := range fb.EventOutputs {
					if eo.Name == a.Output {
						found = true
						break
					}
				}
				if found == false {
					return newValidateError(ErrUndefinedEvent, a.Output, a.DebugInfo)
				}
			}
		}
	}

	//make sure all transitions can be mapped
	for _, t := range fb.BasicFB.Transitions {
		//find their source and destination states
		foundSource := false
		foundDest := false
		for _, s := range fb.BasicFB.States {
			if s.Name == t.Source {
				foundSource = true
			}
			if s.Name == t.Destination {
				foundDest = true
			}
			if foundSource && foundDest {
				break
			}
		}
		if foundSource == false {
			return newValidateError(ErrUndefinedState, t.Source, t.DebugInfo)
		}
		if foundDest == false {
			return newValidateError(ErrUndefinedState, t.Destination, t.DebugInfo)
		}

		//TODO: make sure condition components are valid??

	}

	return nil
}

//isValidDataType returns true if string s is one of the valid IEC61499 event/data types
func isValidDataType(s string) bool {
	s = strings.ToLower(s)
	if s == "bool" ||
		s == "byte" ||
		s == "word" ||
		s == "dword" ||
		s == "lword" ||
		s == "sint" ||
		s == "usint" ||
		s == "int" ||
		s == "uint" ||
		s == "dint" ||
		s == "udint" ||
		s == "lint" ||
		s == "ulint" ||
		s == "real" ||
		s == "lreal" ||
		s == "time" ||
		s == "any" {
		return true
	}
	return false
}

func fbCFBNetworkValid(fb FB, otherFBs []FB) *ValidateError {
	if fb.CompositeFB == nil { //blocks that aren't compositeFBs aren't going to have invalid state machines!
		return nil
	}

	cfb := fb.CompositeFB
	//make sure all instances can be found in otherFBs
	for i := 0; i < len(cfb.FBs); i++ {
		found := false
		for j := 0; j < len(otherFBs); j++ {
			if otherFBs[j].Name == cfb.FBs[i].Type {
				found = true
				break
			}
		}
		if found == false {
			return newValidateError(ErrFbReferenceInvalidType, cfb.FBs[i].Type, cfb.FBs[i].DebugInfo)
		}
	}

	//make sure each instance has a unique name
	for i := 0; i < len(cfb.FBs); i++ {
		for j := 0; j < len(cfb.FBs); j++ {
			if i == j {
				break
			}

			if cfb.FBs[i].Name == cfb.FBs[j].Name {
				return newValidateError(ErrFbReferenceNameNotUnique, cfb.FBs[i].Name, cfb.FBs[i].DebugInfo)
			}

		}
	}

	//TODO: make sure each source/destination port of each {events, data connections} can be found

	return nil
}
