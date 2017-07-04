package tfbparser

//parsePossibleArrayInto will parse either a single item or an array of items into a single-item function
func (t *tfbParse) parsePossibleArrayInto(fbIndex int, singleFn func(*tfbParse, int) *ParseError) *ParseError {

	//if the next argument is a brace, we are going to be looping and creating many singles
	s := t.peek()
	if s == pOpenBrace {
		t.pop() //get rid of the open brace
		for {
			if err := singleFn(t, fbIndex); err != nil {
				return err
			}
			if s := t.peek(); s == pCloseBrace {
				t.pop() //get rid of the close brace
				break
			}
		}
		return nil
	}

	return singleFn(t, fbIndex)
}
