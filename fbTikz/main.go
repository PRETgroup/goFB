package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var (
	inFileName     = flag.String("i", "", "Specifies the name of the source file or directory of fbt-type files to be compiled.")
	outLocation    = flag.String("o", "", "Specifies the name of the directory to put output files. If blank, uses current directory")
	drawInternalOf = flag.String("internalOf", "", "Specifies the name of a function block to render the internals of. If blank, will draw all I/O files instead.")
)

func main() {
	flag.Parse()

	*inFileName = strings.TrimSuffix(*inFileName, "/")
	*inFileName = strings.TrimSuffix(*inFileName, "\\")

	*outLocation = strings.TrimSuffix(*outLocation, "/")
	*outLocation = strings.TrimSuffix(*outLocation, "\\")

	if *outLocation == "" {
		*outLocation = "."
	}

	if *inFileName == "" {
		fmt.Println("You need to specify a file or directory name to compile! Check out -help for options")
		return

	}

	fileInfo, err := os.Stat(*inFileName)
	if err != nil {
		fmt.Println("Error reading file statistics:", err.Error())
		return
	}

	var fbtFileNames []string
	var otherFileNames []string

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
				fbtFileNames = append(fbtFileNames, fmt.Sprintf("%s%c%s", *inFileName, os.PathSeparator, name))
			} else {
				//only copy the "extra" files if they don't begin with a '.'
				if name[0] != '.' {
					otherFileNames = append(otherFileNames, name)
				}
			}
		}
	} else {
		fmt.Println("Running in Single mode")
		fbtFileNames = []string{*inFileName}
	}

	fbTikz := new(FBTikz)

	for _, name := range fbtFileNames {
		sourceFile, err := ioutil.ReadFile(name)
		if err != nil {
			fmt.Printf("Error reading file '%s' for conversion: %s\n", name, err.Error())
			return
		}

		err = fbTikz.AddBlock(sourceFile)
		if err != nil {
			fmt.Printf("Error during adding file '%s' for conversion: %s\n", name, err.Error())
			return
		}
	}

	var outputs []OutputFile

	if *drawInternalOf == "" {
		fmt.Println("Drawing all I/O file(s)")
		outputs, err = fbTikz.ConvertAll()
		if err != nil {
			fmt.Println("Error during conversion:", err.Error())
			return
		}
	} else {
		fmt.Println("Drawing internals of FB called '" + *drawInternalOf + "'")
		outputs, err = fbTikz.ConvertInternal(*drawInternalOf)
		if err != nil {
			fmt.Println("Error during conversion:", err.Error())
			return
		}
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
