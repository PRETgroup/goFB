package main

//a policy begins with "require(condition)", which can be qualified with either
//until(time) or before(time), and must finished with a recover `code`;

//PaceEnforcer example of an Enforcer
var PaceEnforcer = Enforcer{
	Name: "PaceEnforcer",
	IO: IO{
		Inputs: []IOLine{
			{
				Name:    "AVI_time",
				Type:    "unsigned(63 downto 0)",
				Initial: "to_unsigned( 300000000, 64)",
			},
			{
				Name:    "AEI_time",
				Type:    "unsigned(63 downto 0)",
				Initial: "to_unsigned( 900000000, 64)",
			},
			{
				Name:    "URI_time",
				Type:    "unsigned(63 downto 0)",
				Initial: "to_unsigned(1000000000, 64)",
			},
			{
				Name:    "LRI_time",
				Type:    "unsigned(63 downto 0)",
				Initial: "to_unsigned(1100000000, 64)",
			},
			{
				Name:    "VPHigh_time",
				Type:    "unsigned(63 downto 0)",
				Initial: "to_unsigned(  50000000, 64)",
			},
			{
				Name:    "APHigh_time",
				Type:    "unsigned(63 downto 0)",
				Initial: "to_unsigned(  50000000, 64)",
			},
		},
		Enforce: []IOLine{
			{
				Name: "VP",
				Type: "std_logic",
			},
			{
				Name: "AP",
				Type: "std_logic",
			},
			{
				Name: "VS",
				Type: "std_logic",
			},
			{
				Name: "AS",
				Type: "std_logic",
			},
		},
	},
	Policies: []Policy{
		{
			Name:     "NO_AP_VP",
			Triggers: nil,
			Requirements: []Requirement{
				{
					//AP && VP == 0
					Requires: []Operation{
						Operation{
							Type: OperatorEquals,
							A: &Operation{
								Type: OperatorAnd,
								A: &Operation{
									Type:  OperatorVariable,
									Value: "AP",
								},
								B: &Operation{
									Type:  OperatorVariable,
									Value: "VP",
								},
							},
							B: &Operation{
								Type:  OperatorValue,
								Value: "'0'",
							},
						},
					},
					//VP = 0; AP = 0;
					Recover: []Operation{
						{
							Type: OperatorSet,
							A: &Operation{
								Type:  OperatorVariable,
								Value: "VP",
							},
							B: &Operation{
								Type:  OperatorValue,
								Value: "'0'",
							},
						},
						{
							Type: OperatorSet,
							A: &Operation{
								Type:  OperatorVariable,
								Value: "AP",
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
		{ //P2: VS or VP must be true within AVI after an atrial event AS or AP.
			Name: "AVI",
			Triggers: []Trigger{
				{
					Name: "tAVI",
					StartCondition: Operation{
						Type: OperatorEquals,
						A: &Operation{
							Type: OperatorOr,
							A: &Operation{
								Type:  OperatorVariable,
								Value: "AS",
							},
							B: &Operation{
								Type:  OperatorVariable,
								Value: "AP",
							},
						},
						B: &Operation{
							Type:  OperatorValue,
							Value: "'1'",
						},
					},
					ResetCondition: &Operation{
						Type: OperatorEquals,
						A: &Operation{
							Type: OperatorOr,
							A: &Operation{
								Type:  OperatorVariable,
								Value: "VS",
							},
							B: &Operation{
								Type:  OperatorVariable,
								Value: "VP",
							},
						},
						B: &Operation{
							Type:  OperatorValue,
							Value: "'1'",
						},
					},
				},
			},
			Requirements: []Requirement{
				{
					//with tAVI require ((VS or VP) before (AVI_time after tAVI)) recover {VP <= 1;};
					With: []string{"tAVI"},
					Requires: []Operation{
						Operation{
							Type: OperatorBefore,
							A: &Operation{
								Type: OperatorEquals,
								A: &Operation{
									Type: OperatorOr,
									A: &Operation{
										Type:  OperatorVariable,
										Value: "VS",
									},
									B: &Operation{
										Type:  OperatorVariable,
										Value: "VP",
									},
								},
								B: &Operation{
									Type:  OperatorValue,
									Value: "'1'",
								},
							},
							B: &Operation{ //before AVI_time from tAVI
								Type: OperatorAfter,
								A: &Operation{
									Type:  OperatorVariable,
									Value: "AVI_time",
								},
								B: &Operation{
									Type:  OperatorVariable,
									Value: "tAVI",
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
								Value: "VP",
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
		{ //P3: AS or AP must be true within AEI after a ventricular event VS or VP.
			Name: "AEI",
			Triggers: []Trigger{
				{
					Name: "tAEI",
					StartCondition: Operation{
						Type: OperatorEquals,
						A: &Operation{
							Type: OperatorOr,
							A: &Operation{
								Type:  OperatorVariable,
								Value: "VS",
							},
							B: &Operation{
								Type:  OperatorVariable,
								Value: "VP",
							},
						},
						B: &Operation{
							Type:  OperatorValue,
							Value: "'1'",
						},
					},
					ResetCondition: &Operation{
						Type: OperatorEquals,
						A: &Operation{
							Type: OperatorOr,
							A: &Operation{
								Type:  OperatorVariable,
								Value: "AS",
							},
							B: &Operation{
								Type:  OperatorVariable,
								Value: "AP",
							},
						},
						B: &Operation{
							Type:  OperatorValue,
							Value: "'1'",
						},
					},
				},
			},
			Requirements: []Requirement{
				{
					//with tAEI require ((AS or AP) before (AEI_time after tAEI)) recover {AP <= 1;};
					With: []string{"tAEI"},
					Requires: []Operation{
						Operation{
							Type: OperatorBefore,
							A: &Operation{
								Type: OperatorEquals,
								A: &Operation{
									Type: OperatorOr,
									A: &Operation{
										Type:  OperatorVariable,
										Value: "AS",
									},
									B: &Operation{
										Type:  OperatorVariable,
										Value: "AP",
									},
								},
								B: &Operation{
									Type:  OperatorValue,
									Value: "'1'",
								},
							},
							B: &Operation{ //before AVI_time from tAVI
								Type: OperatorAfter,
								A: &Operation{
									Type:  OperatorVariable,
									Value: "AEI_time",
								},
								B: &Operation{
									Type:  OperatorVariable,
									Value: "tAEI",
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
								Value: "AP",
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
		{ //P4: After a ventricular event, another ventricular event can happen only after URI.
			Name: "URI",
			Triggers: []Trigger{
				{
					Name: "tURI", // trigger tURI on (VS || VP) reset after (URI_time after tURI);
					StartCondition: Operation{
						Type: OperatorEquals,
						A: &Operation{
							Type: OperatorOr,
							A: &Operation{
								Type:  OperatorVariable,
								Value: "VS",
							},
							B: &Operation{
								Type:  OperatorVariable,
								Value: "VP",
							},
						},
						B: &Operation{
							Type:  OperatorValue,
							Value: "'1'",
						},
					},
					ResetCondition: &Operation{
						Type: OperatorAfter,
						A:    nil,
						B: &Operation{
							Type: OperatorAdd,
							A: &Operation{
								Type:  OperatorVariable,
								Value: "URI_time",
							},
							B: &Operation{
								Type:  OperatorVariable,
								Value: "tURI",
							},
						},
					},
				},
			},
			Requirements: []Requirement{
				{
					//with tURI require ((VP == 0) until (URI_time after tURI)) recover {VP <= 0;};
					With: []string{"tURI"},
					Requires: []Operation{
						Operation{
							Type: OperatorUntil,
							A: &Operation{
								Type: OperatorEquals,
								A: &Operation{
									Type:  OperatorVariable,
									Value: "VP",
								},
								B: &Operation{
									Type:  OperatorValue,
									Value: "'0'",
								},
							},
							B: &Operation{ //before AVI_time after tAVI
								Type: OperatorAfter,
								A: &Operation{
									Type:  OperatorVariable,
									Value: "URI_time",
								},
								B: &Operation{
									Type:  OperatorVariable,
									Value: "tURI",
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
								Value: "VP",
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
		{ //P5: After a ventricular event, another ventricular event should happen within LRI.
			Name: "LRI",
			Triggers: []Trigger{
				{
					Name: "tLRI",
					StartCondition: Operation{
						Type: OperatorEquals,
						A: &Operation{
							Type: OperatorOr,
							A: &Operation{
								Type:  OperatorVariable,
								Value: "VS",
							},
							B: &Operation{
								Type:  OperatorVariable,
								Value: "VP",
							},
						},
						B: &Operation{
							Type:  OperatorValue,
							Value: "'1'",
						},
					},
					ResetCondition: &Operation{
						Type: OperatorAfter,
						A:    nil,
						B: &Operation{
							Type: OperatorAdd,
							A: &Operation{
								Type:  OperatorVariable,
								Value: "LRI_time",
							},
							B: &Operation{
								Type:  OperatorVariable,
								Value: "tLRI",
							},
						},
					},
				},
			},
			Requirements: []Requirement{
				{
					// with tLRI require (VS || VP) before (LRI_time after tLRI) recover {VP <= '1';};
					With: []string{"tLRI"},
					Requires: []Operation{
						Operation{
							Type: OperatorBefore,
							A: &Operation{
								Type: OperatorEquals,
								A: &Operation{
									Type: OperatorOr,
									A: &Operation{
										Type:  OperatorVariable,
										Value: "VS",
									},
									B: &Operation{
										Type:  OperatorVariable,
										Value: "VP",
									},
								},
								B: &Operation{
									Type:  OperatorValue,
									Value: "'1'",
								},
							},
							B: &Operation{ //before AVI_time from tAVI
								Type: OperatorAdd,
								A: &Operation{
									Type:  OperatorVariable,
									Value: "LRI_time",
								},
								B: &Operation{
									Type:  OperatorVariable,
									Value: "tLRI",
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
								Value: "VP",
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
