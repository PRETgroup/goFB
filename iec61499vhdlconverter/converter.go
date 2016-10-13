package iec61499vhdlconverter

import (
	"encoding/xml"
	"errors"

	"github.com/kiwih/go-iec61499-vhdl/iec61499vhdlconverter/iec61499"
)

//IEC61499ToVHDL converts iec61499 xml-formatted []byte into vhdl []byte
//Returns nil error on success
func IEC61499ToVHDL(iec61499bytes []byte) ([]byte, error) {
	FB := iec61499.FB{}
	if err := xml.Unmarshal(iec61499bytes, &FB); err != nil {
		return nil, errors.New("Couldn't unmarshal iec61499 xml: " + err.Error())
	}

	return nil, errors.New("Not yet implemented")
}

//IEC61499BasicFBVHDL is used as template string for making ECC
const IEC61499BasicFBVHDL = `

`
