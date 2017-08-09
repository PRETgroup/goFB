package tfbparser

import (
	"errors"
	"fmt"
)

var (
	//ErrInternal means something went wrong and it's the transpiler's fault
	ErrInternal = errors.New("An internal error occured")

	//ErrUnexpectedEOF means the document ended unexpectedly
	ErrUnexpectedEOF = errors.New("Unexpected EOF")

	//ErrUnexpectedValue is used to indicate a parsed value was not what was expected (i.e. a word instead of a semicolon)
	ErrUnexpectedValue = errors.New("Unexpected value")

	//ErrUndefinedFB is used to indicate a FB was referenced that can't be found (so probably a typo has occured)
	ErrUndefinedFB = errors.New("Can't find FB with name")

	//ErrUndefinedEvent is used to indicate that an event was referenced that can't be found (so probably a typo has occured)
	ErrUndefinedEvent = errors.New("Can't find Event with name")

	//ErrUnexpectedAssociation is used when an association statement ("with") is used on a data line in an interface inappropriately (i.e. in a compositeFB, or trying to trigger an event)
	ErrUnexpectedAssociation = errors.New("Unexpected association")

	//ErrInvalidType is used when the type of an event or data variable is bad
	ErrInvalidType = errors.New("Invalid or missing data/event type")

	//ErrInvalidIOMeta is used when metadata for an I/O line is bad
	ErrInvalidIOMeta = errors.New("Invalid metadata for data/event line")

	//ErrNameAlreadyInUse is returned whenever something is named but the name is already in use elsewhere
	ErrNameAlreadyInUse = errors.New("This name is already defined elsewhere")

	//ErrOnlyInstancesGetParameters is returned when a constant parameter is attempt-assigned to a cfb output
	ErrOnlyInstancesGetParameters = errors.New("Constant parameters can only be provided to fb instances")
)

//ParseError is used to contain a helpful error message when parsing fails
type ParseError struct {
	LineNumber int
	Argument   string
	Reason     string
	Err        error
}

//Error makes ParseError fulfill error interface
func (p ParseError) Error() string {
	s := fmt.Sprintf("Error (Line %v): %s", p.LineNumber, p.Err.Error())
	if p.Argument != "" {
		s += " '" + p.Argument + "'"
	}
	if p.Reason != "" {
		s += ", (" + p.Reason + ")"
	}
	return s
}

// helper functions to help construct helpful error messages

func (t *tfbParse) errorWithArg(err error, arg string) *ParseError {
	return &ParseError{LineNumber: t.currentLine, Argument: arg, Reason: "", Err: err}
}

func (t *tfbParse) errorWithArgAndLineNumber(err error, arg string, line int) *ParseError {
	return &ParseError{LineNumber: line, Argument: arg, Reason: "", Err: err}
}

func (t *tfbParse) errorWithReason(err error, reason string) *ParseError {
	return &ParseError{LineNumber: t.currentLine, Argument: "", Reason: reason, Err: err}
}

func (t *tfbParse) error(err error) *ParseError {
	return &ParseError{LineNumber: t.currentLine, Argument: "", Reason: "", Err: err}
}

func (t *tfbParse) errorWithArgAndReason(err error, arg string, reason string) *ParseError {
	return &ParseError{LineNumber: t.currentLine, Argument: arg, Reason: reason, Err: err}
}

func (t *tfbParse) errorUnexpectedWithExpected(unexpected string, expected string) *ParseError {
	return &ParseError{LineNumber: t.currentLine, Argument: unexpected, Reason: "Expected: " + expected, Err: ErrUnexpectedValue}
}
