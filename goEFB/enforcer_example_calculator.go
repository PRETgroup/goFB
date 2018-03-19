package main

//a policy begins with "require(condition)", which can be qualified with either
//until(time) or before(time), and must finished with a recover `code`;

//CalculatorEnforcer example of an Enforcer
var CalculatorEnforcer = Enforcer{
	Name: "CalculatorEnforcer",
	IO: IO{
		Enforce: []IOLine{
			// {
			// 	Name: "c1A",
			// 	Type: "unsigned(31 downto 0)",
			// },
			// {
			// 	Name: "c1B",
			// 	Type: "unsigned(31 downto 0)",
			// },
			// {
			// 	Name: "c1X",
			// 	Type: "unsigned(31 downto 0)",
			// },
			{
				Name: "c2A",
				Type: "unsigned(31 downto 0)",
			},
			{
				Name: "c2B",
				Type: "unsigned(31 downto 0)",
			},
			{
				Name: "c2X",
				Type: "unsigned(31 downto 0)",
			},
			{
				Name: "c2Start",
				Type: "std_logic",
			},
			{
				Name: "c2Done",
				Type: "std_logic",
			},
		},
	},
	Policies: []Policy{

		{ //P4: Calculator with done/start signals, c1X <= c1A + c1B with Done, which is no longer than 5ms from any change
			Name: "P1",
			Triggers: []Trigger{
				{
					Name: "tC",
					StartCondition: Operation{
						Type: OperatorEquals,
						A: &Operation{
							Type:  OperatorVariable,
							Value: "c2Start",
						},
						B: &Operation{
							Type:  OperatorValue,
							Value: "'1'",
						},
					},
					ResetCondition: nil,
				},
			},
			Requirements: []Requirement{
				{
					//with tC require (cDone) before (5ms after tC);
					With: []string{"tC"},
					Requires: []Operation{
						Operation{
							Type: OperatorBefore,
							A: &Operation{
								Type: OperatorEquals,
								A: &Operation{
									Type:  OperatorVariable,
									Value: "cDone",
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
									Value: "to_unsigned(5000000, 64)",
								},
								B: &Operation{
									Type:  OperatorVariable,
									Value: "tC",
								},
							},
						},
					},
					//VP = 1;
					Recover: []Operation{
						{
							Type: OperatorSet,
							A: &Operation{
								Type:  OperatorVariable,
								Value: "cDone",
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
	},
}
