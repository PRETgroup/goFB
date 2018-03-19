package main

//a policy begins with "require(condition)", which can be qualified with either
//until(time) or before(time), and must finished with a recover `code`;

//WaterBoilerEnforcer example of an Enforcer
var WaterBoilerEnforcer = Enforcer{
	Name: "WaterBoilerEnforcer",
	IO: IO{
		Enforce: []IOLine{
			{
				Name: "Pboiler",
				Type: "unsigned(7 downto 0)",
			},
			{
				Name: "Fin",
				Type: "unsigned(7 downto 0)",
			},
			{
				Name: "Fout",
				Type: "unsigned(7 downto 0)",
			},
			{
				Name: "Fop",
				Type: "unsigned(7 downto 0)",
			},
			{
				Name: "Lboiler",
				Type: "unsigned(7 downto 0)",
			},
			{
				Name: "Hboiler",
				Type: "std_logic",
			},
			{
				Name: "Cin",
				Type: "std_logic",
			},
			{
				Name: "Vin",
				Type: "unsigned(7 downto 0)",
			},
			{
				Name: "Vop",
				Type: "unsigned(7 downto 0)",
			},
			{
				Name: "Vout",
				Type: "unsigned(7 downto 0)",
			},
			{
				Name: "Aop",
				Type: "std_logic",
			},
		},
	},
	Policies: []Policy{

		{ //P1: Vop must be opened within 10ms if Pboiler > 100
			Name: "P1",
			Triggers: []Trigger{
				{
					Name: "Top",
					StartCondition: Operation{
						Type: OperatorGTE,
						A: &Operation{
							Type:  OperatorVariable,
							Value: "Pboiler",
						},
						B: &Operation{
							Type:  OperatorValue,
							Value: "to_unsigned(100, 8)",
						},
					},
					ResetCondition: &Operation{
						Type: OperatorLT,
						A: &Operation{
							Type:  OperatorVariable,
							Value: "Pboiler",
						},
						B: &Operation{
							Type:  OperatorValue,
							Value: "to_unsigned(100, 8)",
						},
					},
				},
			},
			Requirements: []Requirement{
				{
					//with tOP require ((C = 1) before (10ms after tA));
					With: []string{"Top"},
					Requires: []Operation{
						Operation{
							Type: OperatorBefore,
							A: &Operation{
								Type: OperatorEquals,
								A: &Operation{
									Type:  OperatorVariable,
									Value: "Vop",
								},
								B: &Operation{
									Type:  OperatorValue,
									Value: "to_unsigned(255, 8)",
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
									Value: "Top",
								},
							},
						},
					},
					Recover: []Operation{
						{
							Type: OperatorSet,
							A: &Operation{
								Type:  OperatorVariable,
								Value: "Vop",
							},
							B: &Operation{
								Type:  OperatorValue,
								Value: "to_unsigned(255, 8)",
							},
						},
					},
				},
			},
		},
		{ //P2: Aop must sound if Pboiler > psafe
			Name: "P2",
			Triggers: []Trigger{
				{
					Name: "Top",
					StartCondition: Operation{
						Type: OperatorGTE,
						A: &Operation{
							Type:  OperatorVariable,
							Value: "Pboiler",
						},
						B: &Operation{
							Type:  OperatorValue,
							Value: "to_unsigned(100, 8)",
						},
					},
					ResetCondition: nil,
				},
			},
			Requirements: []Requirement{
				{
					//with Top require (Aop = 1);
					With: []string{"Top"},
					Requires: []Operation{
						Operation{
							Type: OperatorEquals,
							A: &Operation{
								Type:  OperatorVariable,
								Value: "Aop",
							},
							B: &Operation{
								Type:  OperatorValue,
								Value: "'1'",
							},
						},
					},
					Recover: []Operation{
						{
							Type: OperatorSet,
							A: &Operation{
								Type:  OperatorVariable,
								Value: "Aop",
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
