package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/kiwih/goFB/goTFB/tfbparser"
	"github.com/kiwih/goFB/iec61499"
)

var (
	inFileName  = flag.String("i", "", "Specifies the name of the source file or directory of fbt-type files to be compiled.")
	outLocation = flag.String("o", "", "Specifies the name of the directory to put output files. If blank, uses current directory")
)

var (
	xmlHeader          = []byte(`<?xml version="1.0" encoding="UTF-8"?>` + "\n")
	fbTypeHeader       = []byte(`<!DOCTYPE FBType SYSTEM "http://www.holobloc.com/xml/LibraryElement.dtd" >` + "\n")
	resourceTypeHeader = []byte(`<!DOCTYPE ResourceType SYSTEM "http://www.holobloc.com/xml/LibraryElement.dtd" >` + "\n")
	deviceTypeHeader   = []byte(`<!DOCTYPE DeviceType SYSTEM "http://www.holobloc.com/xml/LibraryElement.dtd" >` + "\n")
)

func main() {
	flag.Parse()

	*inFileName = strings.TrimSuffix(*inFileName, "/")
	*inFileName = strings.TrimSuffix(*inFileName, "\\")

	*outLocation = strings.TrimSuffix(*outLocation, "/")
	*outLocation = strings.TrimSuffix(*outLocation, "\\")

	if *inFileName == "" {
		fmt.Println("You need to specify a file or directory name to compile! Check out -help for options")
		return
	}

	fileInfo, err := os.Stat(*inFileName)
	if err != nil {
		fmt.Println("Error reading file statistics:", err.Error())
		return
	}

	var fileNames []string

	if fileInfo.IsDir() {
		//Running in Dir mode
		files, err := ioutil.ReadDir(*inFileName)
		if err != nil {
			log.Fatal(err)
		}

		for _, file := range files {
			//only read the .tfb files
			name := file.Name()
			nameComponents := strings.Split(name, ".")
			extension := nameComponents[len(nameComponents)-1]
			if extension == "tfb" {
				fileNames = append(fileNames, name)
			}
		}
	} else {
		//Running in Single mode
		fileNames = append(fileNames, *inFileName)
	}

	var fbs []iec61499.FB

	for _, name := range fileNames {
		sourceFile, err := ioutil.ReadFile(fmt.Sprintf("%s%c%s", *inFileName, os.PathSeparator, name))
		if err != nil {
			fmt.Printf("Error reading file '%s' for conversion: %s\n", name, err.Error())
			return
		}

		mfbs, parseErr := tfbparser.ParseString(name, string(sourceFile))
		if parseErr != nil {
			fmt.Printf("Error during parsing file '%s': %s\n", name, parseErr.Error())
			return
		}

		//we need to translate any hybridFBs to BFBs before we try export them
		for i := 0; i < len(mfbs); i++ {
			if mfbs[i].HybridFB != nil {
				if err := mfbs[i].TranslateHFBtoBFB(); err != nil {
					fmt.Printf("Error during hybridFB translation in file '%s': %s\n", name, parseErr.Error())
					return
				}
			}
		}
		fbs = append(fbs, mfbs...)

	}

	validateErr := iec61499.ValidateFBs(fbs)
	if validateErr != nil {
		fmt.Println("Error during validation:", validateErr.Error())
		return
	}

	for _, fb := range fbs {
		name := fb.Name
		extn := "fbt"
		//TODO: work out what extension to use based on the fb.XMLname field
		bytes, err := xml.MarshalIndent(fb, "", "\t")
		if err != nil {
			fmt.Println("Error during marshal:", err.Error())
			return
		}
		output := append(xmlHeader, fbTypeHeader...)
		output = append(output, bytes...)

		fmt.Printf("Writing %s.%s\n", name, extn)
		err = ioutil.WriteFile(fmt.Sprintf("%s%c%s.%s", *outLocation, os.PathSeparator, name, extn), output, 0644)
		if err != nil {
			fmt.Println("Error during file write:", err.Error())
			return
		}
	}

}
