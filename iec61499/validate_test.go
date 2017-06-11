package iec61499

import "testing"

func TestBFBValid(t *testing.T) {
	if err := ValidateFBs([]FB{testBfb}); err != nil {
		t.Error("testBfb invalid when it should be valid")
	}
}
