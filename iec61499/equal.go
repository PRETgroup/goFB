package iec61499

import (
	"fmt"
	"reflect"
)

//d is used to overwrite existing debug info when testing for equality (line numbers/source file names don't matter when testing equality)
var d = DebugInfo{}

//FBsEqual checks two function blocks for equality.
//It only checks the important fields.
func FBsEqual(a FB, b FB) bool {
	if a.Name != b.Name {
		return false
	}

	if a.Identification.Standard != b.Identification.Standard {
		return false
	}

	if same := blockInterfacesSame(a, b); same != true {
		return false
	}

	if a.BasicFB != nil && b.BasicFB != nil {
		return basicFbsSame(*a.BasicFB, *b.BasicFB)
	}

	if a.CompositeFB != nil && b.CompositeFB != nil {
		return compositeFbsSame(*a.CompositeFB, *b.CompositeFB)
	}

	if a.ServiceFB != nil && b.ServiceFB != nil {
		return siFbsSame(*a.ServiceFB, *b.ServiceFB)
	}

	if a.HybridFB != nil && b.HybridFB != nil {
		fmt.Printf("I don't know how to test hybridfb equivalence")
		return false
	}

	return false //don't know how to test equivalents on other types of block
}

func blockInterfacesSame(a FB, b FB) bool {
	if len(a.EventInputs) != len(b.EventInputs) {
		return false
	}
	for i := 0; i < len(a.EventInputs); i++ {
		if a.EventInputs[i].Name != b.EventInputs[i].Name {
			return false
		}
	}

	if len(a.EventOutputs) != len(b.EventOutputs) {
		return false
	}
	for i := 0; i < len(a.EventOutputs); i++ {
		if a.EventOutputs[i].Name != b.EventOutputs[i].Name {
			return false
		}
	}

	if len(a.InputVars) != len(b.InputVars) {
		return false
	}
	for i := 0; i < len(a.InputVars); i++ {
		a.InputVars[i].DebugInfo = d
		b.InputVars[i].DebugInfo = d
		if a.InputVars[i] != b.InputVars[i] {
			return false
		}
	}

	if len(a.OutputVars) != len(b.OutputVars) {
		return false
	}
	for i := 0; i < len(a.OutputVars); i++ {
		a.OutputVars[i].DebugInfo = d
		b.OutputVars[i].DebugInfo = d
		if a.OutputVars[i] != b.OutputVars[i] {
			return false
		}
	}

	return true
}

func basicFbsSame(a BasicFB, b BasicFB) bool {
	if len(a.Algorithms) != len(b.Algorithms) {
		return false
	}
	for i := 0; i < len(a.Algorithms); i++ {
		a.Algorithms[i].DebugInfo = d
		b.Algorithms[i].DebugInfo = d
		if a.Algorithms[i] != b.Algorithms[i] {
			return false
		}
	}
	for i := 0; i < len(a.Transitions); i++ {
		a.Transitions[i].DebugInfo = d
		b.Transitions[i].DebugInfo = d
		if a.Transitions[i] != b.Transitions[i] {
			return false
		}
	}
	for i := 0; i < len(a.States); i++ {
		if a.States[i].Name != b.States[i].Name {
			return false
		}
		if len(a.States[i].ECActions) != len(b.States[i].ECActions) {
			return false
		}
		for j := 0; j < len(a.States[i].ECActions); j++ {
			a.States[i].ECActions[j].DebugInfo = d
			b.States[i].ECActions[j].DebugInfo = d
			if a.States[i].ECActions[j] != b.States[i].ECActions[j] {
				return false
			}
		}
	}
	return true
}

func compositeFbsSame(a CompositeFB, b CompositeFB) bool {
	if len(a.FBs) != len(b.FBs) {
		return false
	}
	if len(a.DataConnections) != len(b.DataConnections) {
		return false
	}
	if len(a.EventConnections) != len(b.EventConnections) {
		return false
	}

	for i := 0; i < len(a.FBs); i++ {
		a.FBs[i].DebugInfo = d
		b.FBs[i].DebugInfo = d

		if len(a.FBs[i].Parameter) != len(b.FBs[i].Parameter) {
			return false
		}
		for j := 0; j < len(a.FBs[i].Parameter); j++ {
			a.FBs[i].Parameter[j].DebugInfo = d
			b.FBs[i].Parameter[j].DebugInfo = d
		}
		if !reflect.DeepEqual(a.FBs[i], b.FBs[i]) {
			return false
		}
	}

	for i := 0; i < len(a.DataConnections); i++ {
		a.DataConnections[i].DebugInfo = d
		b.DataConnections[i].DebugInfo = d
		if a.DataConnections[i] != b.DataConnections[i] {
			return false
		}
	}

	for i := 0; i < len(a.EventConnections); i++ {
		a.EventConnections[i].DebugInfo = d
		b.EventConnections[i].DebugInfo = d
		if a.EventConnections[i] != b.EventConnections[i] {
			return false
		}
	}

	return true
}

func siFbsSame(a ServiceFB, b ServiceFB) bool {
	a.Autogenerate.DebugInfo = d
	b.Autogenerate.DebugInfo = d

	return reflect.DeepEqual(a, b)
}
