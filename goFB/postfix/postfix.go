package postfix

//Association is an enum for identifying if operators are left, right, or not associative
type Association int

//AssociativeXXXX define the possible Association types
const (
	AssociationLeft  Association = -1
	AssociationRight Association = 1
	AssociationNone  Association = 0
)

//Operator defines an operation (like AND or ADD or *)
//  which has a precedence and also a number of operands
type Operator interface {
	GetToken() string
	GetPrecedence() int
	GetNumOperands() int
	GetAssociation() Association
}

//A Converter is used to convert between infix and postfix
type Converter struct {
	Operators []Operator
}

//NewConverter creates a converter for a list of operators
func NewConverter(Operators []Operator) Converter {
	return Converter{Operators}
}

//IsOperator checks if a given token is an operator, and if so, returns it
func (c *Converter) IsOperator(tok string) (bool, Operator) {
	for _, op := range c.Operators {
		if op.GetToken() == tok {
			return true, op
		}
	}
	return false, nil
}

//IsOperand returns if a given token is an operand
func (c *Converter) IsOperand(tok string) bool {
	op, _ := c.IsOperator(tok)
	return !op
}

//IsFunction returns if this is a function call or not
func (c *Converter) IsFunction(tok string) bool {
	return len(tok) > 1 && tok[len(tok)-1:] == "("
}

//ToPostfix converts a string slice of tokens in infix format to postfix
func (c *Converter) ToPostfix(tokens []string) []string {

	var stack Stack

	postfix := make([]string, 0, len(tokens))

	length := len(tokens)

	for i := 0; i < length; i++ {

		token := tokens[i]

		if token == "(" {
			stack.Push(token)
		} else if c.IsFunction(token) {
			//TODO count the number of operands
			stack.Push(token)
		} else if token == ")" {
			for !stack.Empty() {
				str, _ := stack.Top().(string)
				if str == "(" {
					break
				}
				postfix = append(postfix, str)
				stack.Pop()
			}
			stack.Pop()
		} else if isOp, _ := c.IsOperator(token); !isOp {
			// If token is not an operator
			// Keep in mind it's just an operand
			postfix = append(postfix, token)
		} else {
			// it is an operator
			// keep grabbing operators until either you get open bracket, or have a precedence problem
			for !stack.Empty() {
				top, _ := stack.Top().(string)
				_, topOp := c.IsOperator(top)
				_, tokOp := c.IsOperator(token)
				if top == "(" ||
					topOp.GetPrecedence() > tokOp.GetPrecedence() ||
					(topOp.GetPrecedence() == tokOp.GetPrecedence() && topOp.GetAssociation() != AssociationLeft) {
					break
				}
				postfix = append(postfix, top)
				stack.Pop()
			}
			stack.Push(token)
		}
	}

	for !stack.Empty() {
		str, _ := stack.Pop().(string)
		postfix = append(postfix, str)
	}

	return postfix
}
