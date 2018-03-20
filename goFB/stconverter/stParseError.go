package stconverter

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

//STParseError is used to contain a helpful error message when parsing fails
type STParseError struct {
	LineNumber int
	Argument   string
	Reason     string
	Err        error
}

//Error makes ParseError fulfill error interface
func (p STParseError) Error() string {
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

func (t *stParse) errorWithArg(err error, arg string) *STParseError {
	return &STParseError{LineNumber: t.currentLine, Argument: arg, Reason: "", Err: err}
}

func (t *stParse) errorWithArgAndLineNumber(err error, arg string, line int) *STParseError {
	return &STParseError{LineNumber: line, Argument: arg, Reason: "", Err: err}
}

func (t *stParse) errorWithReason(err error, reason string) *STParseError {
	return &STParseError{LineNumber: t.currentLine, Argument: "", Reason: reason, Err: err}
}

func (t *stParse) error(err error) *STParseError {
	return &STParseError{LineNumber: t.currentLine, Argument: "", Reason: "", Err: err}
}

func (t *stParse) errorWithArgAndReason(err error, arg string, reason string) *STParseError {
	return &STParseError{LineNumber: t.currentLine, Argument: arg, Reason: reason, Err: err}
}

func (t *stParse) errorUnexpectedWithExpected(unexpected string, expected string) *STParseError {
	return &STParseError{LineNumber: t.currentLine, Argument: unexpected, Reason: "Expected: " + expected, Err: ErrUnexpectedValue}
}
