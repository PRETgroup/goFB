package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/kiwih/goFB/iec61499converter"
)

var (
	inFileName             = flag.String("i", "", "Specifies the name of the source file or directory of fbt-type files to be compiled.")
	outLocation            = flag.String("o", "", "Specifies the name of the directory to put output files. If blank, uses current directory")
	topName                = flag.String("t", "", "Specifies the name of the top level fbt-type file. If blank, no top file will be generated.")
	outputLanguage         = flag.String("l", "c", "Specifies the output language for the program.")
	algorithmLanguageCheck = flag.Bool("alc", false, "Enable checking algorithm language compatibility with output language.")
)

func main() {
	flag.Parse()

	*inFileName = strings.TrimSuffix(*inFileName, "/")
	*inFileName = strings.TrimSuffix(*inFileName, "\\")

	*outLocation = strings.TrimSuffix(*outLocation, "/")
	*outLocation = strings.TrimSuffix(*outLocation, "\\")

	if *inFileName == "" {
		fmt.Println("You need to specify a file or directory name to transpile! Check out -help for options")
		return
	}

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
			extension := nameComponents[len(nameComponents)-1]
			if extension == "fbt" || extension == "res" || extension == "dev" {
				fileNames = append(fileNames, name)
			} else {
				fmt.Println("Didn't add, the extn was", extension)
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

	outputs, err := conv.ConvertAll()
	if err != nil {
		fmt.Println("Error during conversion:", err.Error())
		return
	}

	for _, output := range outputs {
		fmt.Printf("Writing %s.%s\n", output.Name, output.Extension)

		err = ioutil.WriteFile(fmt.Sprintf("%s%c%s.%s", *outLocation, os.PathSeparator, output.Name, output.Extension), output.Contents, 0644)
		if err != nil {
			fmt.Println("Error during file write:", err.Error())
			return
		}
	}

}
