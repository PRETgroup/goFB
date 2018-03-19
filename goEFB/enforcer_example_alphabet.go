package main

//a policy begins with "require(condition)", which can be qualified with either
//until(time) or before(time), and must finished with a recover `code`;

//AlphabetEnforcer example of an Enforcer
var AlphabetEnforcer = Enforcer{
	Name: "AlphabetEnforcer",
	IO: IO{
		Enforce: []IOLine{
			{
				Name: "A",
				Type: "std_logic",
			},
			{
				Name: "B",
				Type: "std_logic",
			},
			{
				Name: "C",
				Type: "std_logic",
			},
			{
				Name: "D",
				Type: "std_logic",
			},
			{
				Name: "L",
				Type: "std_logic",
			},
		},
	},
	Policies: []Policy{

		{ //P1: If B happens within 30ms of A, C must happen within 60ms of A, and D must happen within 10ms of B.
			Name: "P1",
			Triggers: []Trigger{
				{
					Name: "tA",
					StartCondition: Operation{
						Type: OperatorEquals,
						A: &Operation{
							Type:  OperatorVariable,
							Value: "A",
						},
						B: &Operation{
							Type:  OperatorValue,
							Value: "'1'",
						},
					},
					ResetCondition: nil,
				},
				{
					Name: "tB",
					StartCondition: Operation{
						Type: OperatorBefore,
						A: &Operation{
							Type: OperatorEquals,
							A: &Operation{
								Type:  OperatorVariable,
								Value: "B",
							},
							B: &Operation{
								Type:  OperatorValue,
								Value: "'1'",
							},
						},
						B: &Operation{
							Type: OperatorAfter,
							A: &Operation{
								Type:  OperatorValue,
								Value: "to_unsigned(30000000, 64)",
							},
							B: &Operation{
								Type:  OperatorVariable,
								Value: "tA",
							},
						},
					},
					ResetCondition: nil,
				},
			},
			Requirements: []Requirement{
				{
					//with tA, tB require ((C = 1) before (60ms after tA));
					With: []string{"tA", "tB"},
					Requires: []Operation{
						Operation{
							Type: OperatorBefore,
							A: &Operation{
								Type: OperatorEquals,
								A: &Operation{
									Type:  OperatorVariable,
									Value: "C",
								},
								B: &Operation{
									Type:  OperatorValue,
									Value: "'1'",
								},
							},
							B: &Operation{
								Type: OperatorAfter,
								A: &Operation{
									Type:  OperatorValue,
									Value: "to_unsigned(60000000, 64)",
								},
								B: &Operation{
									Type:  OperatorVariable,
									Value: "tA",
								},
							},
						},
					},
					//C = 1;
					Recover: []Operation{
						{
							Type: OperatorSet,
							A: &Operation{
								Type:  OperatorVariable,
								Value: "C",
							},
							B: &Operation{
								Type:  OperatorValue,
								Value: "'1'",
							},
						},
					},
				},
				{
					//with tB require ((D) before (10ms after tB));
					With: []string{"tB"},
					Requires: []Operation{
						Operation{
							Type: OperatorBefore,
							A: &Operation{
								Type: OperatorEquals,
								A: &Operation{
									Type:  OperatorVariable,
									Value: "D",
								},
								B: &Operation{
									Type:  OperatorValue,
									Value: "'1'",
								},
							},
							B: &Operation{
								Type: OperatorAfter,
								A: &Operation{
									Type:  OperatorValue,
									Value: "to_unsigned(10000000, 64)",
								},
								B: &Operation{
									Type:  OperatorVariable,
									Value: "tB",
								},
							},
						},
					},
					//D = 1;
					Recover: []Operation{
						{
							Type: OperatorSet,
							A: &Operation{
								Type:  OperatorVariable,
								Value: "D",
							},
							B: &Operation{
								Type:  OperatorValue,
								Value: "'1'",
							},
						},
					},
				},
			},
		},
		{ ////P2: Flashing light L. L needs to be on for between 500ms and 600ms, and then off for between 500ms and 600ms, then on again etc.
			Name: "P2",
			Triggers: []Trigger{
				{
					Name: "tL1",
					StartCondition: Operation{
						Type: OperatorEquals,
						A: &Operation{
							Type:  OperatorVariable,
							Value: "L",
						},
						B: &Operation{
							Type:  OperatorValue,
							Value: "'0'",
						},
					},
					ResetCondition: &Operation{
						Type: OperatorEquals,
						A: &Operation{
							Type:  OperatorVariable,
							Value: "L",
						},
						B: &Operation{
							Type:  OperatorValue,
							Value: "'1'",
						},
					},
				},
				{
					Name: "tL2",
					StartCondition: Operation{
						Type: OperatorEquals,
						A: &Operation{
							Type:  OperatorVariable,
							Value: "L",
						},
						B: &Operation{
							Type:  OperatorValue,
							Value: "'1'",
						},
					},
					ResetCondition: &Operation{
						Type: OperatorEquals,
						A: &Operation{
							Type:  OperatorVariable,
							Value: "L",
						},
						B: &Operation{
							Type:  OperatorValue,
							Value: "'0'",
						},
					},
				},
			},
			Requirements: []Requirement{
				{
					//with tL1 require ((L=0) until (500ms after tL1));
					With: []string{"tL1"},
					Requires: []Operation{
						Operation{
							Type: OperatorUntil,
							A: &Operation{
								Type: OperatorEquals,
								A: &Operation{
									Type:  OperatorVariable,
									Value: "L",
								},
								B: &Operation{
									Type:  OperatorValue,
									Value: "'0'",
								},
							},
							B: &Operation{
								Type: OperatorAfter,
								A: &Operation{
									Type:  OperatorValue,
									Value: "to_unsigned(500000000, 64)",
								},
								B: &Operation{
									Type:  OperatorVariable,
									Value: "tL1",
								},
							},
						},
					},
					Recover: []Operation{
						{
							Type: OperatorSet,
							A: &Operation{
								Type:  OperatorVariable,
								Value: "L",
							},
							B: &Operation{
								Type:  OperatorValue,
								Value: "'0'",
							},
						},
					},
				},
				{
					//with tL1 require ((L=1) before (600ms after tL1));
					With: []string{"tL1"},
					Requires: []Operation{
						Operation{
							Type: OperatorBefore,
							A: &Operation{
								Type: OperatorEquals,
								A: &Operation{
									Type:  OperatorVariable,
									Value: "L",
								},
								B: &Operation{
									Type:  OperatorValue,
									Value: "'1'",
								},
							},
							B: &Operation{
								Type: OperatorAfter,
								A: &Operation{
									Type:  OperatorValue,
									Value: "to_unsigned(600000000, 64)",
								},
								B: &Operation{
									Type:  OperatorVariable,
									Value: "tL1",
								},
							},
						},
					},
					Recover: []Operation{
						{
							Type: OperatorSet,
							A: &Operation{
								Type:  OperatorVariable,
								Value: "L",
							},
							B: &Operation{
								Type:  OperatorValue,
								Value: "'1'",
							},
						},
					},
				},
				{
					//with tL2 require ((L=1) until (500ms after tL2));
					With: []string{"tL2"},
					Requires: []Operation{
						Operation{
							Type: OperatorUntil,
							A: &Operation{
								Type: OperatorEquals,
								A: &Operation{
									Type:  OperatorVariable,
									Value: "L",
								},
								B: &Operation{
									Type:  OperatorValue,
									Value: "'1'",
								},
							},
							B: &Operation{
								Type: OperatorAfter,
								A: &Operation{
									Type:  OperatorValue,
									Value: "to_unsigned(500000000, 64)",
								},
								B: &Operation{
									Type:  OperatorVariable,
									Value: "tL2",
								},
							},
						},
					},
					Recover: []Operation{
						{
							Type: OperatorSet,
							A: &Operation{
								Type:  OperatorVariable,
								Value: "L",
							},
							B: &Operation{
								Type:  OperatorValue,
								Value: "'1'",
							},
						},
					},
				},
				{
					//with tL2 require ((L=0) before (600ms after tL2));
					With: []string{"tL2"},
					Requires: []Operation{
						Operation{
							Type: OperatorBefore,
							A: &Operation{
								Type: OperatorEquals,
								A: &Operation{
									Type:  OperatorVariable,
									Value: "L",
								},
								B: &Operation{
									Type:  OperatorValue,
									Value: "'0'",
								},
							},
							B: &Operation{
								Type: OperatorAfter,
								A: &Operation{
									Type:  OperatorValue,
									Value: "to_unsigned(600000000, 64)",
								},
								B: &Operation{
									Type:  OperatorVariable,
									Value: "tL2",
								},
							},
						},
					},
					Recover: []Operation{
						{
							Type: OperatorSet,
							A: &Operation{
								Type:  OperatorVariable,
								Value: "L",
							},
							B: &Operation{
								Type:  OperatorValue,
								Value: "'0'",
							},
						},
					},
				},
			},
		},
	},
}
