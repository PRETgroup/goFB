package iec61499

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/PRETgroup/goFB/goFB/stconverter"
)

var aeiFB = *Must(
	NewBasicFB("AEIPolicy").
		AddEventInputNames([]string{"AS", "VS"}, d).
		AddEventOutputNames([]string{"AP", "VP"}, d).
		AddDataInputs([]string{"AEI_ns"}, []string{}, "ulint", "", "900000000", d)).
	AddPolicy("AEI").
	AddPFBState("s1", d).
	AddPFBState("s2", d).
	AddPFBDataInternals([]string{"tAEI"}, "DTIMER", "", "", d).
	AddPFBTransition("s1", "s2", "( VS or VP )", []PFBExpression{{VarName: "tAEI", Value: "0"}}, d).
	AddPFBTransition("s2", "s1", "( AS or AP )", nil, d).
	AddPFBTransition("s2", "violation", "( tAEI > AEI_ns )", nil, d)

// func TestPoliciesSTGuards(t *testing.T) {
// 	stGuards, err := aeiFB.Policies[0].GetPFBSTTransitions()
// 	if err != nil {
// 		t.Fatalf("Got an error and shouldn't have: %v", err.Error())
// 	}
// 	for i, g := range stGuards {
// 		rets := traverse(g.STGuard)
// 		bytes, _ := json.MarshalIndent(rets, "", "\t")
// 		ioutil.WriteFile(fmt.Sprintf("test_actual.%v.out.json", i), bytes, 0644)
// 	}

// }

func TestTraverse1(t *testing.T) {
	guard, err := FBECCGuardToSTExpression("test 1", "a and (b or c)")
	testOut := []stconverter.STExpression{
		stconverter.STExpressionOperator{
			Operator: stconverter.FindOp("and"),
			Arguments: []stconverter.STExpression{
				stconverter.STExpressionValue{Value: "c"},
				stconverter.STExpressionValue{Value: "a"},
			},
		},
		stconverter.STExpressionOperator{
			Operator: stconverter.FindOp("and"),
			Arguments: []stconverter.STExpression{
				stconverter.STExpressionValue{Value: "b"},
				stconverter.STExpressionValue{Value: "a"},
			},
		},
	}
	if err != nil {
		t.Fatalf("Got an error and shouldn't have: %v", err.Error())
	}
	expr, ok := guard[0].(stconverter.STExpression)
	if !ok {
		t.Fatalf("Got an error and shouldn't have: couldn't cast expression")
	}
	rets := traverse(stconverter.STExpression(expr))

	if !reflect.DeepEqual(rets, testOut) {
		expected, _ := json.MarshalIndent(testOut, "\t", "\t")
		received, _ := json.MarshalIndent(rets, "\t", "\t")
		t.Errorf("Test PARSING FAIL.\nExpected:\n\t%s\n\nReceived:\n\t%s\n\n", expected, received)
	}
}

func TestTraverse2(t *testing.T) {
	guard, err := FBECCGuardToSTExpression("test 1", "(a or c) and (b or d)")
	var testOut []stconverter.STExpression
	testOut = []stconverter.STExpression{
		stconverter.STExpressionOperator{
			Operator: stconverter.FindOp("and"),
			Arguments: []stconverter.STExpression{
				stconverter.STExpressionValue{Value: "d"},
				stconverter.STExpressionValue{Value: "c"},
			},
		},
		stconverter.STExpressionOperator{
			Operator: stconverter.FindOp("and"),
			Arguments: []stconverter.STExpression{
				stconverter.STExpressionValue{Value: "d"},
				stconverter.STExpressionValue{Value: "a"},
			},
		},
		stconverter.STExpressionOperator{
			Operator: stconverter.FindOp("and"),
			Arguments: []stconverter.STExpression{
				stconverter.STExpressionValue{Value: "b"},
				stconverter.STExpressionValue{Value: "c"},
			},
		},
		stconverter.STExpressionOperator{
			Operator: stconverter.FindOp("and"),
			Arguments: []stconverter.STExpression{
				stconverter.STExpressionValue{Value: "b"},
				stconverter.STExpressionValue{Value: "a"},
			},
		},
	}
	if err != nil {
		t.Fatalf("Got an error and shouldn't have: %v", err.Error())
	}
	expr, ok := guard[0].(stconverter.STExpression)
	if !ok {
		t.Fatalf("Got an error and shouldn't have: couldn't cast expression")
	}
	rets := traverse(stconverter.STExpression(expr))

	if !reflect.DeepEqual(rets, testOut) {
		expected, _ := json.MarshalIndent(testOut, "\t", "\t")
		received, _ := json.MarshalIndent(rets, "\t", "\t")
		t.Errorf("Test PARSING FAIL.\nExpected:\n\t%s\n\nReceived:\n\t%s\n\n", expected, received)
	}
}

func TestTraverse3(t *testing.T) {
	guard, err := FBECCGuardToSTExpression("test 1", "a and c and (b or d)")
	var testOut []stconverter.STExpression
	testOut = []stconverter.STExpression{
		stconverter.STExpressionOperator{
			Operator: stconverter.FindOp("and"),
			Arguments: []stconverter.STExpression{
				stconverter.STExpressionValue{Value: "d"},
				stconverter.STExpressionOperator{
					Operator: stconverter.FindOp("and"),
					Arguments: []stconverter.STExpression{
						stconverter.STExpressionValue{Value: "c"},
						stconverter.STExpressionValue{Value: "a"},
					},
				},
			},
		},
		stconverter.STExpressionOperator{
			Operator: stconverter.FindOp("and"),
			Arguments: []stconverter.STExpression{
				stconverter.STExpressionValue{Value: "b"},
				stconverter.STExpressionOperator{
					Operator: stconverter.FindOp("and"),
					Arguments: []stconverter.STExpression{
						stconverter.STExpressionValue{Value: "c"},
						stconverter.STExpressionValue{Value: "a"},
					},
				},
			},
		},
	}
	if err != nil {
		t.Fatalf("Got an error and shouldn't have: %v", err.Error())
	}
	expr, ok := guard[0].(stconverter.STExpression)
	if !ok {
		t.Fatalf("Got an error and shouldn't have: couldn't cast expression")
	}
	rets := traverse(stconverter.STExpression(expr))

	if !reflect.DeepEqual(rets, testOut) {
		expected, _ := json.MarshalIndent(testOut, "\t", "\t")
		received, _ := json.MarshalIndent(rets, "\t", "\t")
		t.Errorf("Test PARSING FAIL.\nExpected:\n\t%s\n\nReceived:\n\t%s\n\n", expected, received)
	}
}
