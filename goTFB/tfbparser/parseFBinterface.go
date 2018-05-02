package tfbparser

//parseFBinterface will add an interface to an existing internal FB
func (t *tfbParse) parseFBinterface() *ParseError {
	var s string
	//first word should be of
	s = t.pop()
	if s != pOf {
		return t.errorUnexpectedWithExpected(s, pOf)
	}

	//second word is fb name
	s = t.pop()
	fbIndex := t.getFBIndexFromName(s)
	if fbIndex == -1 {
		return t.errorWithArg(ErrUndefinedFB, s)
	}

	//third should be open brace
	s = t.pop()
	if s != pOpenBrace {
		return t.errorUnexpectedWithExpected(s, pOpenBrace)
	}

	//now we run until closed brace
	for {
		//inside the interface, we have
		//[enforce] in|out event|bool|byte|word|dword|lword|sint|usint|int|uint|dint|udint|lint|ulint|real|lreal|time|any (name[, name]) [on (name[, name]) - non-event types only]; //(comments)
		//repeated over and over again
		s = t.pop()
		if s == "" {
			return t.error(ErrUnexpectedEOF)
		}
		if s == pCloseBrace {
			return nil //we're done here
		}

		if s == pIn || s == pOut {
			if err := t.addFBio(s == pIn, fbIndex); err != nil {
				return err
			}
		}
	}
}

//addFBio adds a line of in/out event/data [with arraysize, triggers, default] to the FB interface
// some error checking is done
//isInput: set to TRUE if you want to add this to the Inputs rather than the Outputs
//fbIndex: the index of the FB we are working on inside t
func (t *tfbParse) addFBio(isInput bool, fbIndex int) *ParseError {
	fb := &t.fbs[fbIndex]

	//next s is type
	typ := t.pop()
	if !isValidType(typ) {
		return t.errorWithArgAndReason(ErrInvalidType, typ, "Expected valid type")
	}

	var intNames []string
	var intTriggers []string

	//there might be an array size next
	size := ""
	if t.peek() == pOpenBracket {
		if typ == pEvent { //events can't have sizes
			return t.errorWithArgAndReason(ErrInvalidIOMeta, "[", "Events cannot be array sized")
		}
		t.pop() // get rid of open bracket
		size = t.pop()
		if s := t.peek(); s != pCloseBracket {
			return t.errorUnexpectedWithExpected(s, pCloseBracket)
		}
		t.pop() //get rid of close bracket
	}

	//this could be an array of names, so we'll loop while we are finding commas
	for {
		name := t.pop()

		intNames = append(intNames, name)
		if t.peek() == pComma {
			t.pop() //get rid of the pComma
			continue
		}
		break
	}

	if t.peek() == pWith { //this input has an update association
		t.pop() //get rid of the on

		if typ == pEvent { //input conditions aren't for events
			return t.errorWithArgAndReason(ErrUnexpectedAssociation, "with", "Events are always updated")
		}

		//this could be an array of names, so we'll loop while we are finding commas
		for {
			intTriggers = append(intTriggers, t.pop())
			if t.peek() == pComma {
				t.pop() //get rid of the pComma
				continue
			}
			break
		}
	}

	//there might be a default value next
	initialValue := ""
	if t.peek() == pInitEq {
		if typ == pEvent { //events can't have sizes
			return t.errorWithReason(ErrInvalidIOMeta, "Events cannot have initial values")
		}
		t.pop() //get rid of pInitial

		s := t.pop()           //this might be an openbracket
		if s == pOpenBracket { //for arrays
			initialValue += s //we need to keep the brackets in
			for {
				s := t.pop()
				if s == "" {
					return t.error(ErrUnexpectedEOF)
				}
				if s == pSemicolon {
					return t.errorUnexpectedWithExpected(s, pOpenBracket)
				}
				if s == pCloseBracket {
					initialValue += s //we need to keep the brackets in
					break
				}
				initialValue += s
			}
		} else { //wasn't an open bracket, must just be value
			initialValue = s
		}
	}

	//clear out last semicolon
	if s := t.pop(); s != pSemicolon {
		return t.errorUnexpectedWithExpected(s, pSemicolon)
	}

	//we now have everything we need to add the io to the interface
	//was it an event?
	if typ == pEvent {
		// HAMMOND: THIS WAS REMOVED (ENFORCEFBS BECAME POLICYFBS), BUT WILL LEAVE IT HERE FOR NOW
		//ENFORCE io works differently, we need to create both an input and an output for it
		//(the in/out specifier now specifies what direction the data is coming from)
		//IN means that the data is going from PLANT to CONTROLLER
		//OUT means the data is going from CONTROLLER to PLANT
		// if isEnforced {
		// 	fb.AddEventInputNames(intNames, t.getCurrentDebugInfo())
		// 	emitNames := make([]string, len(intNames))
		// 	for i, n := range intNames {
		// 		if isInput {
		// 			emitNames[i] = n + "_to_controller"
		// 		} else {
		// 			emitNames[i] = n + "_to_plant"
		// 		}
		// 	}
		// 	fb.AddEventOutputNames(emitNames, t.getCurrentDebugInfo())
		// }

		if isInput { //was it an input?
			fb.AddEventInputNames(intNames, t.getCurrentDebugInfo())
		} else {
			fb.AddEventOutputNames(intNames, t.getCurrentDebugInfo())
		}
		return nil
	}
	//no, it was a data connection
	if isInput { //was it an input?
		//TODO: enforced data lines

		if _, err := fb.AddDataInputs(intNames, intTriggers, typ, size, initialValue, t.getCurrentDebugInfo()); err != nil {
			return t.errorWithArg(ErrUndefinedEvent, err.Arg)
		}

	} else {
		if _, err := fb.AddDataOutputs(intNames, intTriggers, typ, size, initialValue, t.getCurrentDebugInfo()); err != nil {
			return t.errorWithArg(ErrUndefinedEvent, err.Arg)
		}
	}
	return nil
}
