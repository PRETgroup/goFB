package tfbparser

import "github.com/kiwih/goFB/iec61499"

//tfbparse is the containing struct for the parsing code
type tfbParse struct {
	fbs []iec61499.FB

	items     []string
	itemIndex int

	currentLine int
	currentFile string
}

//getCurrentDebugInfo returns the debug info for the last popped item
func (t *tfbParse) getCurrentDebugInfo() iec61499.DebugInfo {
	return iec61499.DebugInfo{
		SourceLine: t.currentLine,
		SourceFile: t.currentFile,
	}
}

//isFBNameUnused will check all registered fbs to see if a name can be used for an FB (as they need to be unique)
func (t *tfbParse) isFBNameUnused(name string) bool {
	for i := 0; i < len(t.fbs); i++ {
		if t.fbs[i].Name == name {
			return false
		}
	}
	return true
}

//pop gets the current element of the tfbparse internal items slice
// and increments the index
func (t *tfbParse) pop() string {
	if t.done() {
		return ""
	}
	s := t.items[t.itemIndex]
	t.itemIndex++

	if s == pNewline {
		t.currentLine++
		return t.pop()
	}
	return s
}

//peek gets the current element of the tfbparse internal items slice (or the next non-newline character)
// without incrementing the index
func (t *tfbParse) peek() string {
	if t.done() {
		return ""
	}
	for i := 0; i < len(t.items); i++ {
		if t.items[t.itemIndex+i] != pNewline {
			return t.items[t.itemIndex+i]
		}
	}
	return ""
}

//done checks to see if the tfbparse is completed (i.e. nothing left to parse)
func (t *tfbParse) done() bool {
	return t.itemIndex >= len(t.items)
}

//getFBIndexFromName will search the tfbparse slice of FBs for one that matches
// the provided name and return the index if found
func (t *tfbParse) getFBIndexFromName(name string) int {
	for i := 0; i < len(t.fbs); i++ {
		if t.fbs[i].Name == name {
			return i
		}
	}
	return -1
}
