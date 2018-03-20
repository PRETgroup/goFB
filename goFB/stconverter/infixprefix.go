package stconverter

func isOperand(s string) (bool, operand) {

}

func convertInfixExpressionToPrefix(infix []string) ([]string, *STParseError) {
	//reverse the string

	operands := make([]operand, 0)
	operators := make([]string, 0)

	for len(infix) > 0 {
		//pop first element from infix
		in, infix := infix[0], infix[1:]

		//check if is operand
		if _, op := isOperand(in); op != nil {
			operands = append(operands, op)
		}

	}

}
