package iec61499

import "testing"

var testBfb = FB{
	Name: "testingFB",
	Identification: Identification{
		Standard: "61499-2",
	},
	EventInputs: []Event{
		{
			Name: "inEv1",
			With: []With{
				{
					Var: "inDat1",
				},
			},
		},
	},
	EventOutputs: []Event{
		{
			Name: "outEv2",
			With: []With{
				{
					Var: "outDat1",
				},
			},
		},
	},
	InputVars: []Variable{
		{
			Name:         "inDat1",
			Type:         "bool",
			ArraySize:    "2",
			InitialValue: "[0,0]",
		},
	},
	OutputVars: []Variable{
		{
			Name:         "outDat1",
			Type:         "bool",
			ArraySize:    "2",
			InitialValue: "[0,0]",
		},
	},
	BasicFB: &BasicFB{
		InternalVars: []Variable{
			{
				Name:         "inVar1",
				Type:         "bool",
				ArraySize:    "2",
				InitialValue: "[0,0]",
			},
		},
		States: []ECState{
			{
				Name: "s1",
				ECActions: []Action{
					{
						Algorithm: "alg1",
					},
					{
						Output: "outEv2",
					},
				},
			},
			{
				Name:      "s2",
				ECActions: []Action{},
			},
		},
		Transitions: []ECTransition{
			{
				Source:      "s1",
				Destination: "s2",
				Condition:   "inEv1",
			},
			{
				Source:      "s2",
				Destination: "s1",
				Condition:   "inEv1",
			},
		},
		Algorithms: []Algorithm{
			{
				Name: "alg1",
				Other: OtherLanguage{
					Language: "C",
					Text:     "//algorithm text here",
				},
			},
		},
	},
}

func TestFBsEqualBasics(t *testing.T) {
	A := testBfb
	B := testBfb

	if !FBsEqual(A, B) {
		t.Error("A and B not equal when they should be")
	}

	B.Name = "asdasdasdasd"
	if FBsEqual(A, B) {
		t.Error("A and B equal and they shouldn't be")
	}

	B = A
	B.Identification.Standard = "237923798234"
	if FBsEqual(A, B) {
		t.Error("A and B equal and they shouldn't be")
	}

	B = A
	B.VersionInfo.Author = "asdasdasd"
	if !FBsEqual(A, B) {
		t.Error("A and B not equal when they should be")
	}
}

func TestFBsEqualWithInterfaces(t *testing.T) {
	A := testBfb
	B := testBfb

	B.EventInputs = nil
	if FBsEqual(A, B) {
		t.Error("A and B equal and they shouldn't be")
	}

	B = A
	B.EventOutputs = nil
	if FBsEqual(A, B) {
		t.Error("A and B equal and they shouldn't be")
	}

	B = A
	B.InputVars = nil
	if FBsEqual(A, B) {
		t.Error("A and B equal and they shouldn't be")
	}

	B = A
	B.OutputVars = nil
	if FBsEqual(A, B) {
		t.Error("A and B equal and they shouldn't be")
	}
}

func TestBFBInternalVars(t *testing.T) {
	//TODO
}

func TestBFBStateMachine(t *testing.T) {
	//TODO
}

func TestBFBAlgorithms(t *testing.T) {
	//TODO
}
