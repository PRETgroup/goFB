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
	{
		name:       "if/then 1",
		progString: "if y > x then y := x; end_if;",
		prog: []STInstruction{
			STIfElsIfElse{
				IfThens: []STIfThen{
					{
						IfExpression: STExpressionOperator{
							Operator: findOp(stGreaterThan),
							Arguments: []STExpression{
								STExpressionValue{"x"},
								STExpressionValue{"y"},
							},
						},
						ThenSequence: []STInstruction{
							STExpressionOperator{
								Operator: findOp(stAssignment),
								Arguments: []STExpression{
									STExpressionValue{"x"},
									STExpressionValue{"y"},
								},
							},
						},
					},
				},
			},
		},
	},
	{
		name: "if/elsif/else 1",
		progString: "" +
			"if y > x then" +
			"	y := x;\n" +
			"	print(\"hello\");\n" +
			"elsif x > y then\n" +
			"	a := 1 + 2 * 3;\n" +
			"else\n" +
			"	print(\"hi\"); \n" +
			"	print(\"yes\"); \n" +
			"end_if;",
		prog: []STInstruction{
			STIfElsIfElse{
				IfThens: []STIfThen{
					{
						IfExpression: STExpressionOperator{
							Operator: findOp(stGreaterThan),
							Arguments: []STExpression{
								STExpressionValue{"x"},
								STExpressionValue{"y"},
							},
						},
						ThenSequence: []STInstruction{
							STExpressionOperator{
								Operator: findOp(stAssignment),
								Arguments: []STExpression{
									STExpressionValue{"x"},
									STExpressionValue{"y"},
								},
							},
							STExpressionOperator{
								Operator: findOp("print<1>"),
								Arguments: []STExpression{
									STExpressionValue{"\"hello\""},
								},
							},
						},
					},
					{
						IfExpression: STExpressionOperator{
							Operator: findOp(stGreaterThan),
							Arguments: []STExpression{
								STExpressionValue{"y"},
								STExpressionValue{"x"},
							},
						},
						ThenSequence: []STInstruction{
							STExpressionOperator{
								Operator: findOp(stAssignment),
								Arguments: []STExpression{
									STExpressionOperator{
										Operator: findOp(stAdd),
										Arguments: []STExpression{
											STExpressionOperator{
												Operator: findOp(stMultiply),
												Arguments: []STExpression{
													STExpressionValue{"3"},
													STExpressionValue{"2"},
												},
											},
											STExpressionValue{"1"},
										},
									},
									STExpressionValue{"a"},
								},
							},
						},
					},
				},
				ElseSequence: []STInstruction{
					STExpressionOperator{
						Operator: findOp("print<1>"),
						Arguments: []STExpression{
							STExpressionValue{"\"hi\""},
						},
					},
					STExpressionOperator{
						Operator: findOp("print<1>"),
						Arguments: []STExpression{
							STExpressionValue{"\"yes\""},
						},
					},
				},
			},
		},
	},
	{
		name: "switchcase 1",
		progString: "" +
			"case x + 1 of \n" +
			"1:	print(\"hello\"); \n" +
			"	y := 2; \n" +
			"2, 3: print(\"many\");\n" +
			"else\n" +
			"	z := 2 + 2;" +
			"end_case;",
		prog: []STInstruction{
			STSwitchCase{
				SwitchOn: STExpressionOperator{
					Operator: findOp(stAdd),
					Arguments: []STExpression{
						STExpressionValue{"1"},
						STExpressionValue{"x"},
					},
				},
				Cases: []STCase{
					STCase{
						CaseValues: []string{"1"},
						Sequence: []STInstruction{
							STExpressionOperator{
								Operator: findOp("print<1>"),
								Arguments: []STExpression{
									STExpressionValue{"\"hello\""},
								},
							},
							STExpressionOperator{
								Operator: findOp(stAssignment),
								Arguments: []STExpression{
									STExpressionValue{"2"},
									STExpressionValue{"y"},
								},
							},
						},
					},
					STCase{
						CaseValues: []string{"2", "3"},
						Sequence: []STInstruction{
							STExpressionOperator{
								Operator: findOp("print<1>"),
								Arguments: []STExpression{
									STExpressionValue{"\"many\""},
								},
							},
						},
					},
				},
				ElseSequence: []STInstruction{
					STExpressionOperator{
						Operator: findOp(stAssignment),
						Arguments: []STExpression{
							STExpressionOperator{
								Operator: findOp(stAdd),
								Arguments: []STExpression{
									STExpressionValue{"2"},
									STExpressionValue{"2"},
								},
							},
							STExpressionValue{"z"},
						},
					},
				},
			},
		},
	},
	{
		name: "switchcase 2",
		progString: "" +
			"case (x + 1) of \n" +
			"1: print(\"hello\"); \n" +
			"	y := 2; \n" +
			"2: \n" +
			"3: print(\"many\");\n" +
			"else\n" +
			"	z := 2 + 2;\n" +
			"end_case;",
		prog: []STInstruction{
			STSwitchCase{
				SwitchOn: STExpressionOperator{
					Operator: findOp(stAdd),
					Arguments: []STExpression{
						STExpressionValue{"1"},
						STExpressionValue{"x"},
					},
				},
				Cases: []STCase{
					STCase{
						CaseValues: []string{"1"},
						Sequence: []STInstruction{
							STExpressionOperator{
								Operator: findOp("print<1>"),
								Arguments: []STExpression{
									STExpressionValue{"\"hello\""},
								},
							},
							STExpressionOperator{
								Operator: findOp(stAssignment),
								Arguments: []STExpression{
									STExpressionValue{"2"},
									STExpressionValue{"y"},
								},
							},
						},
					},
					STCase{
						CaseValues: []string{"2"},
					},
					STCase{
						CaseValues: []string{"3"},
						Sequence: []STInstruction{
							STExpressionOperator{
								Operator: findOp("print<1>"),
								Arguments: []STExpression{
									STExpressionValue{"\"many\""},
								},
							},
						},
					},
				},
				ElseSequence: []STInstruction{
					STExpressionOperator{
						Operator: findOp(stAssignment),
						Arguments: []STExpression{
							STExpressionOperator{
								Operator: findOp(stAdd),
								Arguments: []STExpression{
									STExpressionValue{"2"},
									STExpressionValue{"2"},
								},
							},
							STExpressionValue{"z"},
						},
					},
				},
			},
		},
	},
	{
		name: "for loop 1",
		progString: "" +
			"for i := 1 to 10 by 2 do\n" +
			"	print(i);\n" +
			"end_for;\n",
		prog: []STInstruction{
			STForLoop{
				ForAssignment: STExpressionOperator{
					Operator: findOp(stAssignment),
					Arguments: []STExpression{
						STExpressionValue{"1"},
						STExpressionValue{"i"},
					},
				},
				ToValue:     STExpressionValue{"10"},
				ByIncrement: STExpressionValue{"2"},
				Sequence: []STInstruction{
					STExpressionOperator{
						Operator: findOp("print<1>"),
						Arguments: []STExpression{
							STExpressionValue{"i"},
						},
					},
				},
			},
		},
	},
	{
		name: "for loop 2",
		progString: "" +
			"for i := 1 to (2+10)*5 do\n" +
			"	print(i);\n" +
			"end_for;\n",
		prog: []STInstruction{
			STForLoop{
				ForAssignment: STExpressionOperator{
					Operator: findOp(stAssignment),
					Arguments: []STExpression{
						STExpressionValue{"1"},
						STExpressionValue{"i"},
					},
				},
				ToValue: STExpressionOperator{
					Operator: findOp(stMultiply),
					Arguments: []STExpression{
						STExpressionValue{"5"},
						STExpressionOperator{
							Operator: findOp(stAdd),
							Arguments: []STExpression{
								STExpressionValue{"10"},
								STExpressionValue{"2"},
							},
						},
					},
				},
				Sequence: []STInstruction{
					STExpressionOperator{
						Operator: findOp("print<1>"),
						Arguments: []STExpression{
							STExpressionValue{"i"},
						},
					},
				},
			},
		},
	},
	{
		name: "while loop 1",
		progString: "" +
			"while i >= 1 do\n" +
			"	print(i);\n" +
			"	i := i-1;\n" +
			"	exit;\n" +
			"end_while;\n",
		prog: []STInstruction{
			STWhileLoop{
				WhileExpression: STExpressionOperator{
					Operator: findOp(stGreaterThanEqualTo),
					Arguments: []STExpression{
						STExpressionValue{"1"},
						STExpressionValue{"i"},
					},
				},
				Sequence: []STInstruction{
					STExpressionOperator{
						Operator: findOp("print<1>"),
						Arguments: []STExpression{
							STExpressionValue{"i"},
						},
					},
					STExpressionOperator{
						Operator: findOp(stAssignment),
						Arguments: []STExpression{
							STExpressionOperator{
								Operator: findOp(stSubtract),
								Arguments: []STExpression{
									STExpressionValue{"1"},
									STExpressionValue{"i"},
								},
							},
							STExpressionValue{"i"},
						},
					},
					STExpressionOperator{
						Operator: findOp("exit"),
					},
				},
			},
		},
	},
	{
		name: "bad while loop 1",
		progString: "" +
			"while i >= 1 do\n" +
			"	print(i);\n" +
			"	i := i-1;\n" +
			"end_for;\n",
		err: ErrUnexpectedToken,
	},
	{
		name: "bad while loop 2",
		progString: "" +
			"while i >= 1 do\n" +
			"	print(i);\n" +
			"	i := i-1;\n",
		err: ErrUnexpectedEOF,
	},
}

func TestCases(t *testing.T) {
	for i := 0; i < len(stTestCases); i++ {
		prog, err := ParseString(stTestCases[i].name, stTestCases[i].progString)
		if err != nil && stTestCases[i].err != nil {
			if stTestCases[i].err.Error() != err.Err.Error() {
				t.Errorf("Test %d (%s) FAIL.\nError mismatch. Expecting %s, but received:%s", i, stTestCases[i].name, stTestCases[i].err.Error(), err.Err.Error())
			}
			continue
		}
		if err != nil {
			t.Errorf("Test %d (%s) FAIL.\nNot expecting error, but received:%s", i, stTestCases[i].name, err.Error())
			continue
		}
		if stTestCases[i].err != nil {
			t.Errorf("Test %d (%s) FAIL.\nWas expecting error, but did not receive.", i, stTestCases[i].name)
		}
		if !reflect.DeepEqual(prog, stTestCases[i].prog) {
			expected, _ := json.MarshalIndent(stTestCases[i].prog, "\t", "\t")
			received, _ := json.MarshalIndent(prog, "\t", "\t")
			t.Errorf("Test %d (%s) FAIL.\n:Expected:\n\t%s\n\nReceived:\n\t%s\n\n", i, stTestCases[i].name, expected, received)
		}
	}
}
