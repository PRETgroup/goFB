package stconverter

import (
	"encoding/json"
	"reflect"
	"strings"
	"testing"
)

type stTestCase struct {
	name           string
	progString     string
	prog           []STInstruction
	compC          string
	compVhdl       string
	compVerilog    string
	expressionOnly bool
	err            error
	knownNames     []string
}

var stTestCases = []stTestCase{
	{
		name:       "basic 1",
		progString: "1",
		prog: []STInstruction{
			STExpressionValue{"1"},
		},
		compC:          "1",
		compVhdl:       "1",
		compVerilog:    "1",
		expressionOnly: true,
	},
	{
		name:       "assignment 1",
		progString: "x := 1;",
		prog: []STInstruction{
			STExpressionOperator{
				Operator: FindOp(stAssignment),
				Arguments: []STExpression{
					STExpressionValue{"1"},
					STExpressionValue{"x"},
				},
			},
		},
		compC:       "x = 1;",
		compVhdl:    "x := 1;",
		compVerilog: "x = 1;",
	},
	{
		name:       "assignment 2",
		progString: "y := x + 2;",
		prog: []STInstruction{
			STExpressionOperator{
				Operator: FindOp(stAssignment),
				Arguments: []STExpression{
					STExpressionOperator{
						Operator: FindOp(stAdd),
						Arguments: []STExpression{
							STExpressionValue{"2"},
							STExpressionValue{"x"},
						},
					},
					STExpressionValue{"y"},
				},
			},
		},
		compC:       "y = x + 2;",
		compVhdl:    "y := x + 2;",
		compVerilog: "y = x + 2;",
	},
	{ //strictly speaking this might be invalid ST, not sure if they use ! as NOT
		name:       "assignment 2!",
		progString: "y := !x + 2;",
		prog: []STInstruction{
			STExpressionOperator{
				Operator: FindOp(stAssignment),
				Arguments: []STExpression{
					STExpressionOperator{
						Operator: FindOp(stAdd),
						Arguments: []STExpression{
							STExpressionValue{"2"},
							STExpressionOperator{
								Operator: FindOp(stNot),
								Arguments: []STExpression{
									STExpressionValue{"x"},
								},
							},
						},
					},
					STExpressionValue{"y"},
				},
			},
		},
	},
	/*	{ 
			name:       "assignment 3",
			progString: "y := (a or b) and (c or d);",
			prog: []STInstruction{
				STExpressionOperator{
					Operator: FindOp(stAssignment),
					Arguments: []STExpression{
						STExpressionOperator{
							Operator: FindOp(stAdd),
							Arguments: []STExpression{
								STExpressionValue{"2"},
								STExpressionOperator{
									Operator: FindOp(stNot),
									Arguments: []STExpression{
										STExpressionValue{"x"},
									},
								},
							},
						},
						STExpressionValue{"y"},
					},
				},
			},
		compC:       "y = !x + 2;",
		compVhdl:    "y := not(x) + 2;",
		compVerilog: "y = ~x + 2;",
	}, */
	{ 
		name:       "issue #1",
		progString: "y := ((a = 1 or b = 3) and c = 0);",
		prog: []STInstruction{
			STExpressionOperator{
				Operator: FindOp(stAssignment),
				Arguments: []STExpression{
					STExpressionOperator{
						Operator: FindOp(stAnd),
						Arguments: []STExpression{
							STExpressionOperator{
								Operator: FindOp(stEqual),
								Arguments: []STExpression{
									STExpressionValue{"0"},
									STExpressionValue{"c"},
								},
							},
							STExpressionOperator{
								Operator: FindOp(stOr),
								Arguments: []STExpression{
									STExpressionOperator{
										Operator: FindOp(stEqual),
										Arguments: []STExpression{
											STExpressionValue{"3"},
											STExpressionValue{"b"},
										},
									},
									STExpressionOperator{
										Operator: FindOp(stEqual),
										Arguments: []STExpression{
											STExpressionValue{"1"},
											STExpressionValue{"a"},
										},
									},
								},
							},
						},
					},
					STExpressionValue{"y"},
				},
			},
		},
		compC:       "y = (a == 1 || b == 3) && c == 0;",
		compVhdl:    "y := (a = 1 or b = 3) and c = 0;",
		compVerilog: "y = (a == 1 || b == 3) && c == 0;",
	},
	{
		name:       "assignment 3",
		progString: "integrationError := -WindupGuard;",
		prog: []STInstruction{
			STExpressionOperator{
				Operator: FindOp(stAssignment),
				Arguments: []STExpression{
					STExpressionOperator{
						Operator: FindOp(stNegative),
						Arguments: []STExpression{
							STExpressionValue{"WindupGuard"},
						},
					},
					STExpressionValue{"integrationError"},
				},
			},
		},
		compC:       "integrationError = -WindupGuard;",
		compVhdl:    "integrationError := -WindupGuard;",
		compVerilog: "integrationError = -WindupGuard;",
	},
	{
		name:       "basic function call",
		progString: "print(\"hi\");",
		prog: []STInstruction{
			STExpressionOperator{
				Operator: FindOp("print<1>"),
				Arguments: []STExpression{
					STExpressionValue{"\"hi\""},
				},
			},
		},
		compC: "print(\"hi\");",
	},
	{
		name:       "basic function call 2",
		progString: `printf("%d\n", i);`,
		prog: []STInstruction{
			STExpressionOperator{
				Operator: FindOp("printf<2>"),
				Arguments: []STExpression{
					STExpressionValue{`i`},
					STExpressionValue{`"%d\n"`},
				},
			},
		},
		compC: `printf("%d\n", i);`,
	},
	{
		name:       "if/then 1",
		progString: "if y > x then y := x; end_if;",
		prog: []STInstruction{
			STIfElsIfElse{
				IfThens: []STIfThen{
					{
						IfExpression: STExpressionOperator{
							Operator: FindOp(stGreaterThan),
							Arguments: []STExpression{
								STExpressionValue{"x"},
								STExpressionValue{"y"},
							},
						},
						ThenSequence: []STInstruction{
							STExpressionOperator{
								Operator: FindOp(stAssignment),
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
		compC:       "if(y > x) { y = x; }",
		compVhdl:    "if (y > x) then y := x; end if;",
		compVerilog: "if (y > x) begin y = x; end",
	},
	{
		name:       "if/then 2",
		progString: "if y < -x then y := -x; end_if;",
		prog: []STInstruction{
			STIfElsIfElse{
				IfThens: []STIfThen{
					{
						IfExpression: STExpressionOperator{
							Operator: FindOp(stLessThan),
							Arguments: []STExpression{
								STExpressionOperator{
									Operator: FindOp(stNegative),
									Arguments: []STExpression{
										STExpressionValue{"x"},
									},
								},
								STExpressionValue{"y"},
							},
						},
						ThenSequence: []STInstruction{
							STExpressionOperator{
								Operator: FindOp(stAssignment),
								Arguments: []STExpression{
									STExpressionOperator{
										Operator: FindOp(stNegative),
										Arguments: []STExpression{
											STExpressionValue{"x"},
										},
									},
									STExpressionValue{"y"},
								},
							},
						},
					},
				},
			},
		},
		compC:       "if(y < -x) { y = -x; }",
		compVhdl:    "if (y < -x) then y := -x; end if;",
		compVerilog: "if (y < -x) begin y = -x; end",
	},
	{
		name: "if/elsif/else 1",
		progString: "" +
			"if y > x then" +
			"	y := x;\n" +
			"	print(\"hello\");\n" +
			"elsif y <= x then\n" +
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
							Operator: FindOp(stGreaterThan),
							Arguments: []STExpression{
								STExpressionValue{"x"},
								STExpressionValue{"y"},
							},
						},
						ThenSequence: []STInstruction{
							STExpressionOperator{
								Operator: FindOp(stAssignment),
								Arguments: []STExpression{
									STExpressionValue{"x"},
									STExpressionValue{"y"},
								},
							},
							STExpressionOperator{
								Operator: FindOp("print<1>"),
								Arguments: []STExpression{
									STExpressionValue{"\"hello\""},
								},
							},
						},
					},
					{
						IfExpression: STExpressionOperator{
							Operator: FindOp(stLessThanEqualTo),
							Arguments: []STExpression{
								STExpressionValue{"x"},
								STExpressionValue{"y"},
							},
						},
						ThenSequence: []STInstruction{
							STExpressionOperator{
								Operator: FindOp(stAssignment),
								Arguments: []STExpression{
									STExpressionOperator{
										Operator: FindOp(stAdd),
										Arguments: []STExpression{
											STExpressionOperator{
												Operator: FindOp(stMultiply),
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
						Operator: FindOp("print<1>"),
						Arguments: []STExpression{
							STExpressionValue{"\"hi\""},
						},
					},
					STExpressionOperator{
						Operator: FindOp("print<1>"),
						Arguments: []STExpression{
							STExpressionValue{"\"yes\""},
						},
					},
				},
			},
		},
		compC: `
			if(y > x) {
				y = x;
				print("hello");
			} else if(y <= x) {
				a = 1 + 2 * 3;
			} else {
				print("hi");
				print("yes");
			}`,
		compVhdl: `
			if (y > x) then
				y := x;
				print("hello");
			elsif (y <= x) then
				a := 1 + 2 * 3;
			else
				print("hi");
				print("yes");
			end if;`,
		compVerilog: `
			if (y > x) begin
				y = x;
				print("hello");
			end else if (y <= x) begin
				a = 1 + 2 * 3;
			end else begin
				print("hi");
				print("yes");
			end`,
	},
	{
		name: "if/elsif/else 2",
		progString: `
		if integrationError < -WindupGuard then
			integrationError := -WindupGuard;
		elsif integrationError > WindupGuard then
			integrationError := WindupGuard;
		end_if;
		`,
		prog: []STInstruction{
			STIfElsIfElse{
				IfThens: []STIfThen{
					{
						IfExpression: STExpressionOperator{
							Operator: FindOp(stLessThan),
							Arguments: []STExpression{
								STExpressionOperator{
									Operator: FindOp(stNegative),
									Arguments: []STExpression{
										STExpressionValue{"WindupGuard"},
									},
								},
								STExpressionValue{"integrationError"},
							},
						},
						ThenSequence: []STInstruction{
							STExpressionOperator{
								Operator: FindOp(stAssignment),
								Arguments: []STExpression{
									STExpressionOperator{
										Operator: FindOp(stNegative),
										Arguments: []STExpression{
											STExpressionValue{"WindupGuard"},
										},
									},
									STExpressionValue{"integrationError"},
								},
							},
						},
					},
					{
						IfExpression: STExpressionOperator{
							Operator: FindOp(stGreaterThan),
							Arguments: []STExpression{
								STExpressionValue{"WindupGuard"},
								STExpressionValue{"integrationError"},
							},
						},
						ThenSequence: []STInstruction{
							STExpressionOperator{
								Operator: FindOp(stAssignment),
								Arguments: []STExpression{
									STExpressionValue{"WindupGuard"},
									STExpressionValue{"integrationError"},
								},
							},
						},
					},
				},
			},
		},
		compC: `
		if(integrationError < -WindupGuard) {
			integrationError = -WindupGuard;
		} else if(integrationError > WindupGuard) {
			integrationError = WindupGuard;
		}
		`,
		compVhdl: `
		if (integrationError < -WindupGuard) then
			integrationError := -WindupGuard;
		elsif (integrationError > WindupGuard) then
			integrationError := WindupGuard;
		end if;
		`,
		compVerilog: `
		if (integrationError < -WindupGuard) begin
			integrationError = -WindupGuard;
		end else if (integrationError > WindupGuard) begin
			integrationError = WindupGuard;
		end
		`,
	},
	{
		name: "switchcase 1",
		progString: `
			case x + 1 of 
			1:	print("hello");
				y := 2;
			2, 3: print("many");
			else
				z := 2 + 2;
			end_case;`,
		prog: []STInstruction{
			STSwitchCase{
				SwitchOn: STExpressionOperator{
					Operator: FindOp(stAdd),
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
								Operator: FindOp("print<1>"),
								Arguments: []STExpression{
									STExpressionValue{"\"hello\""},
								},
							},
							STExpressionOperator{
								Operator: FindOp(stAssignment),
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
								Operator: FindOp("print<1>"),
								Arguments: []STExpression{
									STExpressionValue{"\"many\""},
								},
							},
						},
					},
				},
				ElseSequence: []STInstruction{
					STExpressionOperator{
						Operator: FindOp(stAssignment),
						Arguments: []STExpression{
							STExpressionOperator{
								Operator: FindOp(stAdd),
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
		compC: `
		switch(x + 1) {
			case 1:
				print("hello");
				y = 2;
				break;
			case 2:
			case 3:
				print("many");
				break;
			default:
				z = 2 + 2;
				break;
		}`,
		compVhdl: `
		case (x + 1) is
			when 1 =>
				print("hello");
				y := 2;
			when 2 | 3 =>
				print("many");
			when others =>
				z := 2 + 2;
		end case;
		`,
		compVerilog: `
		case(x + 1)
			1: begin
				print("hello");
				y = 2;
			end
			2, 3: begin
				print("many");
			end
			default: begin
				z = 2 + 2;
			end
		endcase
		`,
	},
	{
		name: "switchcase 2",
		progString: `
			case (x + 1) of 
			1: print("hello");
				y := 2; 
			2: 
			3: print("many");
			else
				z := 2 + 2;
			end_case;`,
		prog: []STInstruction{
			STSwitchCase{
				SwitchOn: STExpressionOperator{
					Operator: FindOp(stAdd),
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
								Operator: FindOp("print<1>"),
								Arguments: []STExpression{
									STExpressionValue{"\"hello\""},
								},
							},
							STExpressionOperator{
								Operator: FindOp(stAssignment),
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
								Operator: FindOp("print<1>"),
								Arguments: []STExpression{
									STExpressionValue{"\"many\""},
								},
							},
						},
					},
				},
				ElseSequence: []STInstruction{
					STExpressionOperator{
						Operator: FindOp(stAssignment),
						Arguments: []STExpression{
							STExpressionOperator{
								Operator: FindOp(stAdd),
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
		compC: `
		switch(x + 1) {
			case 1:
				print("hello");
				y = 2;
				break;
			case 2:
				break;
			case 3:
				print("many");
				break;
			default:
				z = 2 + 2;
				break;
		}`,
		compVhdl: `
		case (x + 1) is
			when 1 =>
				print("hello");
				y := 2;
			when 2 =>
			when 3 =>
				print("many");
			when others =>
				z := 2 + 2;
		end case;`,
		compVerilog: `
		case(x + 1)
			1: begin
				print("hello");
				y = 2;
				end
			2: begin
				end
			3: begin
				print("many");
				end
			default: begin
				z = 2 + 2;
				end
		endcase`,
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
					Operator: FindOp(stAssignment),
					Arguments: []STExpression{
						STExpressionValue{"1"},
						STExpressionValue{"i"},
					},
				},
				ToValue:     STExpressionValue{"10"},
				ByIncrement: STExpressionValue{"2"},
				Sequence: []STInstruction{
					STExpressionOperator{
						Operator: FindOp("print<1>"),
						Arguments: []STExpression{
							STExpressionValue{"i"},
						},
					},
				},
			},
		},
		compC: "for(i = 1; i <= 10; i += 2) { print(i); }",
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
					Operator: FindOp(stAssignment),
					Arguments: []STExpression{
						STExpressionValue{"1"},
						STExpressionValue{"i"},
					},
				},
				ToValue: STExpressionOperator{
					Operator: FindOp(stMultiply),
					Arguments: []STExpression{
						STExpressionValue{"5"},
						STExpressionOperator{
							Operator: FindOp(stAdd),
							Arguments: []STExpression{
								STExpressionValue{"10"},
								STExpressionValue{"2"},
							},
						},
					},
				},
				Sequence: []STInstruction{
					STExpressionOperator{
						Operator: FindOp("print<1>"),
						Arguments: []STExpression{
							STExpressionValue{"i"},
						},
					},
				},
			},
		},
		compC: "for(i = 1; i <= (2 + 10) * 5; i++) { print(i); }",
		// compVhdl: ` //for loops not yet supported in VHDL
		// 	for i in 1 to (2 + 10) * 5 loop
		// 		print(i);
		// 	end loop;`,
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
					Operator: FindOp(stGreaterThanEqualTo),
					Arguments: []STExpression{
						STExpressionValue{"1"},
						STExpressionValue{"i"},
					},
				},
				Sequence: []STInstruction{
					STExpressionOperator{
						Operator: FindOp("print<1>"),
						Arguments: []STExpression{
							STExpressionValue{"i"},
						},
					},
					STExpressionOperator{
						Operator: FindOp(stAssignment),
						Arguments: []STExpression{
							STExpressionOperator{
								Operator: FindOp(stSubtract),
								Arguments: []STExpression{
									STExpressionValue{"1"},
									STExpressionValue{"i"},
								},
							},
							STExpressionValue{"i"},
						},
					},
					STExpressionOperator{
						Operator: FindOp(stExit),
					},
				},
			},
		},
		compC: `
		while(i >= 1) {
			print(i);
			i = i - 1;
			break;
		}`,
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
	{
		name: "repeat loop 1",
		progString: "" +
			"repeat\n" +
			"	print(i);\n" +
			"	i := i-1;\n" +
			"until i <> 5\n" +
			"end_repeat;",
		prog: []STInstruction{
			STRepeatLoop{
				UntilExpression: STExpressionOperator{
					Operator: FindOp(stInequal),
					Arguments: []STExpression{
						STExpressionValue{"5"},
						STExpressionValue{"i"},
					},
				},
				Sequence: []STInstruction{
					STExpressionOperator{
						Operator: FindOp("print<1>"),
						Arguments: []STExpression{
							STExpressionValue{"i"},
						},
					},
					STExpressionOperator{
						Operator: FindOp(stAssignment),
						Arguments: []STExpression{
							STExpressionOperator{
								Operator: FindOp(stSubtract),
								Arguments: []STExpression{
									STExpressionValue{"1"},
									STExpressionValue{"i"},
								},
							},
							STExpressionValue{"i"},
						},
					},
				},
			},
		},
		compC: `
		do {
			print(i);
			i = i - 1;
		} while(!(i != 5));`,
	},
	/*{ //test removed: "until" statement is now compulsory
		name: "repeat loop 2",
		progString: "" +
			"repeat\n" +
			"	printf(\"%d\", i);\n" +
			"	exit;\n" +
			"end_repeat;",
		prog: []STInstruction{
			STRepeatLoop{
				Sequence: []STInstruction{
					STExpressionOperator{
						Operator: FindOp("printf<2>"),
						Arguments: []STExpression{
							STExpressionValue{"\"%d\n\""},
							STExpressionValue{"i"},
						},
					},
					STExpressionOperator{
						Operator: FindOp(stExit),
					},
				},
			},
		},
		compC: `
		do {
			printf("%d", i);
			break;
		} while(1);
		`,
	},*/
	{
		name:       "known names test",
		progString: "if y > 5 then y := x; end_if;",
		prog: []STInstruction{
			STIfElsIfElse{
				IfThens: []STIfThen{
					{
						IfExpression: STExpressionOperator{
							Operator: FindOp(stGreaterThan),
							Arguments: []STExpression{
								STExpressionValue{"5"},
								STExpressionValue{"y"},
							},
						},
						ThenSequence: []STInstruction{
							STExpressionOperator{
								Operator: FindOp(stAssignment),
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
		compC:      "if(me->y > 5) { me->y = me->x; }",
		knownNames: []string{"x", "y"},
	},
}

func TestCases(t *testing.T) {
	for i := 0; i < len(stTestCases); i++ {
		//for i := 5; i < 6; i++ {
		t.Logf("Running test %d (%s)\n", i, stTestCases[i].name)

		SetKnownVarNames(stTestCases[i].knownNames)
		prog, err := ParseString(stTestCases[i].name, stTestCases[i].progString)
		if err != nil && stTestCases[i].err != nil {
			if stTestCases[i].err.Error() != err.Err.Error() {
				t.Errorf("Test %d (%s) PARSING FAIL.\nError mismatch. Expecting %s, but received:%s", i, stTestCases[i].name, stTestCases[i].err.Error(), err.Err.Error())
			}
			continue
		}
		if err != nil {
			t.Errorf("Test %d (%s) PARSING FAIL.\nNot expecting error, but received:%s", i, stTestCases[i].name, err.Error())
			continue
		}
		if stTestCases[i].err != nil {
			t.Errorf("Test %d (%s) PARSING FAIL.\nWas expecting error, but did not receive.", i, stTestCases[i].name)
		}
		if !reflect.DeepEqual(prog, stTestCases[i].prog) {
			expected, _ := json.MarshalIndent(stTestCases[i].prog, "\t", "\t")
			received, _ := json.MarshalIndent(prog, "\t", "\t")
			t.Errorf("Test %d (%s) PARSING FAIL.\nExpected:\n\t%s\n\nReceived:\n\t%s\n\n", i, stTestCases[i].name, expected, received)
		}

		if stTestCases[i].compC != "" {
			//now check if the compiled version matches
			var recvProg string
			if stTestCases[i].expressionOnly {
				recvProg = standardizeSpaces(CCompileExpression(prog[0].(STExpression)))
			} else {
				recvProg = standardizeSpaces(CCompileSequence(prog))
			}

			//convert to have equivalent whitespaces
			desrProg := standardizeSpaces(stTestCases[i].compC)

			if recvProg != desrProg {
				t.Errorf("Test %d (%s) C COMPILATION FAIL.\nExpected:\n\t%s\n\nReceived:\n\t%s\n\n", i, stTestCases[i].name, desrProg, recvProg)
			}
		}
		if stTestCases[i].compVhdl != "" {
			//now check if the compiled version matches
			var recvProg string
			if stTestCases[i].expressionOnly {
				recvProg = standardizeSpaces(VhdlCompileExpression(prog[0].(STExpression)))
			} else {
				recvProg = standardizeSpaces(VhdlCompileSequence(prog))
			}

			//convert to have equivalent whitespaces
			desrProg := standardizeSpaces(stTestCases[i].compVhdl)

			if recvProg != desrProg {
				t.Errorf("Test %d (%s) VHDL COMPILATION FAIL.\nExpected:\n\t%s\n\nReceived:\n\t%s\n\n", i, stTestCases[i].name, desrProg, recvProg)
			}
		}
		if stTestCases[i].compVerilog != "" {
			//now check if the compiled version matches
			var recvProg string
			if stTestCases[i].expressionOnly {
				recvProg = standardizeSpaces(VerilogCompileExpression(prog[0].(STExpression)))
			} else {
				recvProg = standardizeSpaces(VerilogCompileSequence(prog))
			}

			//convert to have equivalent whitespaces
			desrProg := standardizeSpaces(stTestCases[i].compVerilog)

			if recvProg != desrProg {
				t.Errorf("Test %d (%s) VERILOG COMPILATION FAIL.\nExpected:\n\t%s\n\nReceived:\n\t%s\n\n", i, stTestCases[i].name, desrProg, recvProg)
			}
		}
	}
}

//standardizeSpaces makes all whitespace gaps in a given string become a single space character
func standardizeSpaces(s string) string {
	return strings.Join(strings.Fields(s), " ")
}
