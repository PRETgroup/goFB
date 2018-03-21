package postfix

import (
	"fmt"
	"strconv"
	"strings"
)

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

//functionOp is just so we can make our arbitrary function calls
// (that are in tokens)
// behave like operators
type functionOp struct {
	token string
}

//GetToken returns the token string and fulfils the Operator interface
func (f functionOp) GetToken() string {
	return f.token
}

//GetPrecedence returns the precedence (which is the highest priority value, 0, for functions) and fulfils the Operator interface
func (f functionOp) GetPrecedence() int {
	return 0
}

//GetNumOperands returns the number of operands used in this function
func (f functionOp) GetNumOperands() int {
	first := strings.Index(f.token, "<")
	if first == -1 {
		return 0
	}
	ops := f.token[first+1 : len(f.token)-1]
	opsInt, err := strconv.Atoi(ops)
	if err != nil {
		return 0
	}
	return opsInt
}

//GetAssociation returns the Association (which is always Right for functions) and fulfils the Operator interface
func (f functionOp) GetAssociation() Association {
	return AssociationRight
}

func (f functionOp) MarshalJSON() ([]byte, error) {
	return []byte("\"" + f.token + "\""), nil
}

//NewConverter creates a converter for a list of operators
func NewConverter(Operators []Operator) Converter {
	return Converter{Operators}
}

//IsOperator checks if a given token is an operator, and if so, returns it
func (c *Converter) IsOperator(tok string) (bool, Operator) {
	if c.IsFunction(tok) { //functions are just a special kind of operator
		return true, functionOp{tok}
	}
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
	is, _ := IsFunction(tok)
	return is
}

//IsFunction returns if this is a function call or not
func IsFunction(tok string) (bool, Operator) {
	if (len(tok) > 1 && tok[len(tok)-1:] == "(") ||
		(len(tok) > 2 && tok[len(tok)-1:] == ">") {
		return true, functionOp{tok}
	}
	return false, nil
}

//IsPossibleFunctionName returns if a given string _could_ be a function
//i.e. if it begins with a zero, it cannot be a function
func IsPossibleFunctionName(tok string) bool {
	if len(tok) == 0 {
		return false
	}
	c := byte(tok[0])
	if c >= '0' && c <= '9' {
		return false
	}
	if c >= 'a' && c <= 'z' {
		return true
	}
	if c >= 'A' && c <= 'Z' {
		return true
	}
	return false
}

func (c *Converter) changeFunctionCalls(tokens []string) []string {
	outp := make([]string, 0)
	for i := 0; i < len(tokens); i++ {
		token := tokens[i]
		//check if it is in list of operators
		isOp, _ := c.IsOperator(token)
		if isOp {
			outp = append(outp, token)
			continue
		}

		//check if it has a valid name
		if !IsPossibleFunctionName(token) {
			outp = append(outp, token)
			continue
		}

		//check if the next thing is an bracket
		if i+1 < len(tokens) && tokens[i+1] == "(" {
			//we have found a function
			outp = append(outp, token+"(")
			i++
		} else {
			//not a bracket: not a function
			outp = append(outp, token)
		}

	}
	return outp
}

//ToPostfix converts a string slice of tokens in infix format to postfix
//this is an implementation of the Shunting-Yard algorithm
//https://en.wikipedia.org/wiki/Shunting-yard_algorithm
//	It has been extended to support arbitrary function calls.
//	These are treated as right-aligned operators (i.e. function(2, 3) becomes "2 3 function<2>")
func (c *Converter) ToPostfix(nftokens []string) []string {
	tokens := c.changeFunctionCalls(nftokens)
	var stack Stack

	postfix := make([]string, 0, len(tokens))

	length := len(tokens)

	for i := 0; i < length; i++ {

		token := tokens[i]
		if token == "(" {
			stack.Push(token)
		} else if token == ")" {
			for !stack.Empty() {
				str, _ := stack.Top().(string)
				if str == "(" {
					break
				}
				//extension to algorithm: if the stack contains a function call, then we need to push that to the output
				//	(open brackets are normally just tossed away)
				if c.IsFunction(str) {
					postfix = append(postfix, str)
					break
				}
				postfix = append(postfix, str)
				stack.Pop()
			}
			stack.Pop()
		} else if token == "," {
			//do nothing, token is not important here
			//	commas are only used when counting them during function call checks in if c.IsFunction() {} earlier
		} else if isOp, _ := c.IsOperator(token); !isOp && !c.IsFunction(token) {
			//if the token is not an operator, and the token is not a function
			//	then, the token is an operand, and it should be pushed onto the output.
			postfix = append(postfix, token)
		} else {
			// the token is an operator or a function call

			//extension to algorithm: if it is a function call, then create an appropriately-sized operand and rewrite the token
			if c.IsFunction(token) {
				// count the number of arguments to the function call
				numArgs := 0
				numOpenBrackets := 0 //we only want to count the number of commas at our bracket nesting level, this keeps track of nesting level
				for j := i + 1; j < length; j++ {
					tmpToken := tokens[j]
					if tmpToken == ")" {
						if numOpenBrackets == 0 {
							break
						}
						numOpenBrackets--
					}
					if numArgs == 0 {
						numArgs++
					}

					//we only want to count the number of commas at our bracket level
					//otherwise, nested function calls will mess things up
					if tmpToken == "," && numOpenBrackets == 0 {
						numArgs++
					}

					//keep track of any open brackets or additional function calls
					if tmpToken == "(" || c.IsFunction(tmpToken) {
						numOpenBrackets++
					}
				}
				token = fmt.Sprintf("%s<%d>", token[:len(token)-1], numArgs)
			}

			// keep grabbing operators until either you get open bracket, or have a precedence problem
			for !stack.Empty() {
				top, _ := stack.Top().(string)
				if top == "(" || c.IsFunction(top) {
					break
				}
				_, topOp := c.IsOperator(top)
				_, tokOp := c.IsOperator(token)

				if topOp.GetPrecedence() > tokOp.GetPrecedence() ||
					(topOp.GetPrecedence() == tokOp.GetPrecedence() && topOp.GetAssociation() != AssociationLeft) {
					break
				}
				//fmt.Printf("appending\n")
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

	//TODO: scan the completed postfix and check for errors

	return postfix
}
