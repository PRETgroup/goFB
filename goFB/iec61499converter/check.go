package iec61499converter

import (
	"errors"

	"github.com/PRETgroup/goFB/iec61499"
)

//checkFB is used internally to ensure that a given inputted function block is amenable to this conversion tool
func (c *Converter) checkFB(fb *iec61499.FB) error {
	//1. Make sure all algorithms are written in VHDL
	if c.IgnoreAlgorithmLanguages == false {
		if fb.BasicFB != nil {
			for i := 0; i < len(fb.BasicFB.Algorithms); i++ {
				if !c.outputLanguage.equals(fb.BasicFB.Algorithms[i].Other.Language) {
					return errors.New("Algorithm " + fb.BasicFB.Algorithms[i].Name + " in block " + fb.Name + " is not written in VHDL")
				}
			}
		}
	}

	//2. Ensure no use of STRING or ANY types in data or internal variables TODO
	return nil
}
