package stconverter

import "github.com/PRETgroup/goFB/iec61499"

//stParse is the containing struct for the parsing code
type stParse struct {
	instructions []STInstruction

	items     []string
	itemIndex int

	currentLine int
	currentFile string
}

//getCurrentDebugInfo returns the debug info for the last popped item
func (t *stParse) getCurrentDebugInfo() iec61499.DebugInfo {
	return iec61499.DebugInfo{
		SourceLine: t.currentLine,
		SourceFile: t.currentFile,
	}
}

//pop gets the current element of the tfbparse internal items slice
// and increments the index
func (t *stParse) pop() string {
	if t.done() {
		return ""
	}
	s := t.items[t.itemIndex]
	t.itemIndex++

	if s == stNewline {
		t.currentLine++
		return t.pop()
	}
	return s
}

//peek gets the current element of the tfbparse internal items slice (or the next non-newline character)
// without incrementing the index
func (t *stParse) peek() string {
	if t.done() {
		return ""
	}
	for i := 0; i < len(t.items); i++ {
		if t.items[t.itemIndex+i] != stNewline {
			return t.items[t.itemIndex+i]
		}
	}
	return ""
}

//done checks to see if the tfbparse is completed (i.e. nothing left to parse)
func (t *stParse) done() bool {
	return t.itemIndex >= len(t.items)
}
