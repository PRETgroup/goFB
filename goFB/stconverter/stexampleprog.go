package stconverter

var exampleSTProg = []STInstruction{
	//x := 1;
	STAssignment{
		AValue: "x",
		Assigned: STExpression{
			AValue: "1",
		},
	},
	//y := x + 2;
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
	//TODO: if, for, loops, switch, etc
}
