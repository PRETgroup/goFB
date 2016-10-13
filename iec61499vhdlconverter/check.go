package iec61499vhdlconverter

import (
	"errors"

	"github.com/kiwih/go-iec61499-vhdl/iec61499vhdlconverter/iec61499"
)

func checkFB(fb *iec61499.FB) error {
	if fb.BasicFB != nil {
		for i := 0; i < len(fb.BasicFB.Algorithms); i++ {
			if fb.BasicFB.Algorithms[i].Other.Language != "VHDL" {
				return errors.New("Algorithm " + fb.BasicFB.Algorithms[i].Name + " in block " + fb.Name + " is not written in VHDL")
			}
		}
	}
	return nil
}
