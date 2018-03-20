package stconverter

import "testing"

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
			STAssignment{
				AValue: "x",
				Assigned: STExpression{
					AValue: "1",
				},
			},
		},
	},
	{
		name:       "assignment 2",
		progString: "y := x + 2;",
		prog: []STInstruction{
			STAssignment{
				AValue: "y",
				Assigned: STExpression{
					AValue:   "x",
					Operator: stAdd,
					B: &STExpression{
						AValue: "2",
					},
				},
			},
		},
	},
	{
		name:       "if/then 1",
		progString: "if y > x then y := x; end_if;",
		prog: []STInstruction{
			STIfElsIfElse{
				IfThens: []STIfThen{
					{
						IfExpression: STExpression{
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
	},
}

func TestCases(t *testing.T) {
	for i := 0; i < len(stTestCases); i++ {

	}
}
