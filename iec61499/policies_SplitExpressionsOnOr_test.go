package iec61499

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/PRETgroup/goFB/goFB/stconverter"
)

func TestSplitExpressionsOnOr1(t *testing.T) {
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
	rets := SplitExpressionsOnOr(stconverter.STExpression(expr))

	if !reflect.DeepEqual(rets, testOut) {
		expected, _ := json.MarshalIndent(testOut, "\t", "\t")
		received, _ := json.MarshalIndent(rets, "\t", "\t")
		t.Errorf("Test PARSING FAIL.\nExpected:\n\t%s\n\nReceived:\n\t%s\n\n", expected, received)
	}
}

func TestSplitExpressionsOnOr2(t *testing.T) {
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
	rets := SplitExpressionsOnOr(stconverter.STExpression(expr))

	if !reflect.DeepEqual(rets, testOut) {
		expected, _ := json.MarshalIndent(testOut, "\t", "\t")
		received, _ := json.MarshalIndent(rets, "\t", "\t")
		t.Errorf("Test PARSING FAIL.\nExpected:\n\t%s\n\nReceived:\n\t%s\n\n", expected, received)
	}
}

func TestSplitExpressionsOnOr3(t *testing.T) {
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
	rets := SplitExpressionsOnOr(stconverter.STExpression(expr))

	if !reflect.DeepEqual(rets, testOut) {
		expected, _ := json.MarshalIndent(testOut, "\t", "\t")
		received, _ := json.MarshalIndent(rets, "\t", "\t")
		t.Errorf("Test PARSING FAIL.\nExpected:\n\t%s\n\nReceived:\n\t%s\n\n", expected, received)
	}
}

func TestSplitExpressionsOnOr4(t *testing.T) {
	guard, err := FBECCGuardToSTExpression("test 1", "a or b or c or d")
	var testOut []stconverter.STExpression
	testOut = []stconverter.STExpression{

		stconverter.STExpressionValue{Value: "d"},
		stconverter.STExpressionValue{Value: "c"},
		stconverter.STExpressionValue{Value: "b"},
		stconverter.STExpressionValue{Value: "a"},
	}
	if err != nil {
		t.Fatalf("Got an error and shouldn't have: %v", err.Error())
	}
	expr, ok := guard[0].(stconverter.STExpression)
	if !ok {
		t.Fatalf("Got an error and shouldn't have: couldn't cast expression")
	}
	rets := SplitExpressionsOnOr(stconverter.STExpression(expr))

	if !reflect.DeepEqual(rets, testOut) {
		expected, _ := json.MarshalIndent(testOut, "\t", "\t")
		received, _ := json.MarshalIndent(rets, "\t", "\t")
		t.Errorf("Test PARSING FAIL.\nExpected:\n\t%s\n\nReceived:\n\t%s\n\n", expected, received)
	}
}

func TestTraverse5(t *testing.T) {
	guard, err := FBECCGuardToSTExpression("test 1", "(a and b) or c")
	testOut := []stconverter.STExpression{
		stconverter.STExpressionValue{Value: "c"},
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
	rets := SplitExpressionsOnOr(stconverter.STExpression(expr))

	if !reflect.DeepEqual(rets, testOut) {
		expected, _ := json.MarshalIndent(testOut, "\t", "\t")
		received, _ := json.MarshalIndent(rets, "\t", "\t")
		t.Errorf("Test PARSING FAIL.\nExpected:\n\t%s\n\nReceived:\n\t%s\n\n", expected, received)
	}
}

func TestSplitExpressionsOnOr6(t *testing.T) {
	guard, err := FBECCGuardToSTExpression("test 1", "(a or b) and (c or d) and (e or f or g)")
	var testOut []stconverter.STExpression
	testOut = []stconverter.STExpression{
		stconverter.STExpressionOperator{
			Operator: stconverter.FindOp("and"),
			Arguments: []stconverter.STExpression{
				stconverter.STExpressionValue{Value: "g"},
				stconverter.STExpressionOperator{
					Operator: stconverter.FindOp("and"),
					Arguments: []stconverter.STExpression{
						stconverter.STExpressionValue{Value: "d"},
						stconverter.STExpressionValue{Value: "b"},
					},
				},
			},
		},
		stconverter.STExpressionOperator{
			Operator: stconverter.FindOp("and"),
			Arguments: []stconverter.STExpression{
				stconverter.STExpressionValue{Value: "g"},
				stconverter.STExpressionOperator{
					Operator: stconverter.FindOp("and"),
					Arguments: []stconverter.STExpression{
						stconverter.STExpressionValue{Value: "d"},
						stconverter.STExpressionValue{Value: "a"},
					},
				},
			},
		},
		stconverter.STExpressionOperator{
			Operator: stconverter.FindOp("and"),
			Arguments: []stconverter.STExpression{
				stconverter.STExpressionValue{Value: "g"},
				stconverter.STExpressionOperator{
					Operator: stconverter.FindOp("and"),
					Arguments: []stconverter.STExpression{
						stconverter.STExpressionValue{Value: "c"},
						stconverter.STExpressionValue{Value: "b"},
					},
				},
			},
		},
		stconverter.STExpressionOperator{
			Operator: stconverter.FindOp("and"),
			Arguments: []stconverter.STExpression{
				stconverter.STExpressionValue{Value: "g"},
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
				stconverter.STExpressionValue{Value: "f"},
				stconverter.STExpressionOperator{
					Operator: stconverter.FindOp("and"),
					Arguments: []stconverter.STExpression{
						stconverter.STExpressionValue{Value: "d"},
						stconverter.STExpressionValue{Value: "b"},
					},
				},
			},
		},
		stconverter.STExpressionOperator{
			Operator: stconverter.FindOp("and"),
			Arguments: []stconverter.STExpression{
				stconverter.STExpressionValue{Value: "f"},
				stconverter.STExpressionOperator{
					Operator: stconverter.FindOp("and"),
					Arguments: []stconverter.STExpression{
						stconverter.STExpressionValue{Value: "d"},
						stconverter.STExpressionValue{Value: "a"},
					},
				},
			},
		},
		stconverter.STExpressionOperator{
			Operator: stconverter.FindOp("and"),
			Arguments: []stconverter.STExpression{
				stconverter.STExpressionValue{Value: "f"},
				stconverter.STExpressionOperator{
					Operator: stconverter.FindOp("and"),
					Arguments: []stconverter.STExpression{
						stconverter.STExpressionValue{Value: "c"},
						stconverter.STExpressionValue{Value: "b"},
					},
				},
			},
		},
		stconverter.STExpressionOperator{
			Operator: stconverter.FindOp("and"),
			Arguments: []stconverter.STExpression{
				stconverter.STExpressionValue{Value: "f"},
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
				stconverter.STExpressionValue{Value: "e"},
				stconverter.STExpressionOperator{
					Operator: stconverter.FindOp("and"),
					Arguments: []stconverter.STExpression{
						stconverter.STExpressionValue{Value: "d"},
						stconverter.STExpressionValue{Value: "b"},
					},
				},
			},
		},
		stconverter.STExpressionOperator{
			Operator: stconverter.FindOp("and"),
			Arguments: []stconverter.STExpression{
				stconverter.STExpressionValue{Value: "e"},
				stconverter.STExpressionOperator{
					Operator: stconverter.FindOp("and"),
					Arguments: []stconverter.STExpression{
						stconverter.STExpressionValue{Value: "d"},
						stconverter.STExpressionValue{Value: "a"},
					},
				},
			},
		},
		stconverter.STExpressionOperator{
			Operator: stconverter.FindOp("and"),
			Arguments: []stconverter.STExpression{
				stconverter.STExpressionValue{Value: "e"},
				stconverter.STExpressionOperator{
					Operator: stconverter.FindOp("and"),
					Arguments: []stconverter.STExpression{
						stconverter.STExpressionValue{Value: "c"},
						stconverter.STExpressionValue{Value: "b"},
					},
				},
			},
		},
		stconverter.STExpressionOperator{
			Operator: stconverter.FindOp("and"),
			Arguments: []stconverter.STExpression{
				stconverter.STExpressionValue{Value: "e"},
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
	rets := SplitExpressionsOnOr(stconverter.STExpression(expr))

	if !reflect.DeepEqual(rets, testOut) {
		expected, _ := json.MarshalIndent(testOut, "\t", "\t")
		received, _ := json.MarshalIndent(rets, "\t", "\t")
		t.Errorf("Test PARSING FAIL.\nExpected:\n\t%s\n\nReceived:\n\t%s\n\n", expected, received)
	}
}

func TestSplitExpressionsOnOr7(t *testing.T) {
	guard, err := FBECCGuardToSTExpression("test 1", "a and (b or c or d)")
	var testOut []stconverter.STExpression
	testOut = []stconverter.STExpression{
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
	rets := SplitExpressionsOnOr(stconverter.STExpression(expr))

	if !reflect.DeepEqual(rets, testOut) {
		expected, _ := json.MarshalIndent(testOut, "\t", "\t")
		received, _ := json.MarshalIndent(rets, "\t", "\t")
		t.Errorf("Test PARSING FAIL.\nExpected:\n\t%s\n\nReceived:\n\t%s\n\n", expected, received)
	}
}
