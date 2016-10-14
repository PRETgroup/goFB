package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/kiwih/go-iec61499-vhdl/iec61499vhdlconverter"
)

var (
	inFileName  = flag.String("i", "", "Specifies the name of the source file or directory of files to be compiled")
	outLocation = flag.String("o", "", "Specifies the name of the directory to put output vhdl files")
)

func main() {
	flag.Parse()

	if len(*inFileName) == 0 {
		fmt.Println("The source file or directory name must be specified using the -i=[name] flag")
		return
	}

	*inFileName = strings.TrimSuffix(*inFileName, "/")
	*inFileName = strings.TrimSuffix(*inFileName, "\\")

	*outLocation = strings.TrimSuffix(*outLocation, "/")
	*outLocation = strings.TrimSuffix(*outLocation, "\\")

	fmt.Println("i=", *inFileName)

	fileInfo, err := os.Stat(*inFileName)
	if err != nil {
		fmt.Println("Error reading file statistics:", err.Error())
		return
	}

	fileNames := make([]string, 0)

	if fileInfo.IsDir() {
		files, err := ioutil.ReadDir(*inFileName)
		if err != nil {
			log.Fatal(err)
		}

		for _, file := range files {
			//only read the .fbt and .res files
			name := file.Name()
			nameComponents := strings.Split(name, ".")
			if nameComponents[len(nameComponents)-1] == "fbt" || nameComponents[len(nameComponents)-1] == "res" {
				fileNames = append(fileNames, name)
			} else {
				fmt.Println("Didn't add, the extn was", nameComponents[len(nameComponents)-1])
			}
		}
	} else {
		fileNames = append(fileNames, *inFileName)
	}

	for _, name := range fileNames {
		sourceFile, err := ioutil.ReadFile(fmt.Sprintf("%s%c%s", *inFileName, os.PathSeparator, name))
		if err != nil {
			fmt.Println("Error reading file:", err.Error())
			return
		}

		vhdl, err := iec61499vhdlconverter.IEC61499ToVHDL(sourceFile)
		if err != nil {
			fmt.Println("Error during conversion:", err.Error())
			return
		}

		nameComponents := strings.Split(name, ".")
		nameComponents[len(nameComponents)-1] = "vhd" //change the lastmost extension to vhd
		name = strings.Join(nameComponents, ".")

		fmt.Println("Writing", name)

		err = ioutil.WriteFile(fmt.Sprintf("%s%c%s", *outLocation, os.PathSeparator, name), vhdl, 0644)
		if err != nil {
			fmt.Println("Error during file write:", err.Error())
			return
		}
	}

}
