package iec61499vhdlconverter

import (
	"errors"

	"github.com/kiwih/go-iec61499-vhdl/iec61499vhdlconverter/iec61499"
)

//checkFB is used internally to ensure that a given inputted function block is amenable to this conversion tool
func checkFB(fb *iec61499.FB, ignoreLanguages bool) error {
	//1. Make sure all algorithms are written in VHDL
	if ignoreLanguages == false {
		if fb.BasicFB != nil {
			for i := 0; i < len(fb.BasicFB.Algorithms); i++ {
				if fb.BasicFB.Algorithms[i].Other.Language != "VHDL" {
					return errors.New("Algorithm " + fb.BasicFB.Algorithms[i].Name + " in block " + fb.Name + " is not written in VHDL")
				}
			}
		}
	}

	//2. Ensure no use of STRING or ANY types in data or internal variables TODO
	return nil
}
