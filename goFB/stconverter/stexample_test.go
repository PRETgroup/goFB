package stconverter

import (
	"encoding/json"
	"reflect"
	"testing"
)

type stTestCase struct {
	name       string
	progString string
	prog       []STInstruction
	err        error
}

var stTestCases = []stTestCase{
	{
		name:       "assignment 1",
		progString: "x := 1;",
		prog: []STInstruction{
			STExpressionOperator{
				Operator: findOp(stAssignment),
				Arguments: []STExpression{
					STExpressionValue{"1"},
					STExpressionValue{"x"},
				},
			},
		},
	},
	{
		name:       "assignment 2",
		progString: "y := x + 2;",
		prog: []STInstruction{
			STExpressionOperator{
				Operator: findOp(stAssignment),
				Arguments: []STExpression{
					STExpressionOperator{
						Operator: findOp(stAdd),
						Arguments: []STExpression{
							STExpressionValue{"2"},
							STExpressionValue{"x"},
						},
					},
					STExpressionValue{"y"},
				},
			},
		},
	},
	/*{
		name:       "if/then 1",
		progString: "if y > x then y := x; end_if;",
		prog: []STInstruction{
			STIfElsIfElse{
				IfThens: []STIfThen{
					{
						IfExpression: STExpressionOperator{
							AValue:   "y",
							Operator: stGreaterThan,
							B: &STExpression{
								AValue: "x",
							},
						},
						ThenSequence: []STInstruction{
							STAssignment{
								AValue: "y",
								Assigned: STExpression{
									AValue: "x",
								},
							},
						},
					},
				},
			},
		},
	},*/
}

func TestCases(t *testing.T) {
	for i := 0; i < len(stTestCases); i++ {
		prog, err := ParseString(stTestCases[i].name, stTestCases[i].progString)
		if err != nil && stTestCases[i].err != nil {
			//TODO check if errors are the same
			continue
		}
		if !reflect.DeepEqual(prog, stTestCases[i].prog) {
			expected, _ := json.MarshalIndent(stTestCases[i].prog, "\t", "\t")
			received, _ := json.MarshalIndent(prog, "\t", "\t")
			t.Errorf("Test %d (%s) FAIL.\n:Expected:\n\t%s\n\nReceived:\n\t%s\n\n", i, stTestCases[i].name, expected, received)
		}
	}
}
