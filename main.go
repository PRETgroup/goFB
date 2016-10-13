package main

import (
	"flag"
	"fmt"
	"io/ioutil"

	"github.com/kiwih/go-iec61499-vhdl/iec61499vhdlconverter"
)

var (
	inFileName  = flag.String("i", "", "Specifies the name of the source file to be assembled")
	outFileName = flag.String("o", "out.txt", "Specifies the name of the output file")
)

func main() {
	flag.Parse()

	if len(*inFileName) == 0 {
		fmt.Println("The source file name must be specified using the -i=[name] flag")
		return
	}

	sourceFile, err := ioutil.ReadFile(*inFileName)
	if err != nil {
		fmt.Println("Error reading file:", err.Error())
		return
	}

	vhdl, err := iec61499vhdlconverter.IEC61499ToVHDL(sourceFile)
	if err != nil {
		fmt.Println("Error during conversion:", err.Error())
		return
	}

	fmt.Printf("\n%s\n\n", vhdl)
}
