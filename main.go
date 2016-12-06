package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/kiwih/go-iec61499-vhdl/iec61499converter"
)

var (
	inFileName             = flag.String("i", "", "Specifies the name of the source file or directory of files to be compiled. If blank, uses this directory")
	outLocation            = flag.String("o", "", "Specifies the name of the directory to put output vhdl files. If blank, uses this directory")
	topName                = flag.String("t", "", "Specifies the name of the top level fbt file. If blank, no top file will be generated.")
	outputLanguage         = flag.String("l", "vhdl", "Specifies the output language for the program.")
	algorithmLanguageCheck = flag.Bool("alc", false, "Sets flag for checking algorithm language compatibility with output language.")
)

func main() {
	flag.Parse()

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
		fmt.Println("Running in Dir mode")
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
		fmt.Println("Running in Single mode")
		fileNames = append(fileNames, *inFileName)
	}

	conv, err := iec61499converter.New(*outputLanguage)
	if err != nil {
		fmt.Println("Error creating converter:", err.Error())
		return
	}

	if *algorithmLanguageCheck == false {
		conv.DisableAlgorithmLanguageChecks()
	}

	for _, name := range fileNames {
		sourceFile, err := ioutil.ReadFile(fmt.Sprintf("%s%c%s", *inFileName, os.PathSeparator, name))
		if err != nil {
			fmt.Printf("Error reading file '%s' for conversion: %s\n", name, err.Error())
			return
		}

		err = conv.AddBlock(sourceFile)
		if err != nil {
			fmt.Printf("Error during adding file '%s' for conversion: %s\n", name, err.Error())
			return
		}
	}

	fmt.Println("Found", len(conv.Blocks), "blocks")

	if err := conv.SetTopName(*topName); err != nil {
		fmt.Printf("Error with provided top name:%s\n", err.Error())
		return
	}

	outputs, err := conv.AllToVHDL()
	if err != nil {
		fmt.Println("Error during conversion:", err.Error())
		return
	}

	for _, output := range outputs {
		fmt.Println("Writing", output.Name)

		err = ioutil.WriteFile(fmt.Sprintf("%s%c%s.vhd", *outLocation, os.PathSeparator, output.Name), output.VHDL, 0644)
		if err != nil {
			fmt.Println("Error during file write:", err.Error())
			return
		}
	}

}
