package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/kiwih/goFB/goFB/iec61499converter"
)

var (
	inFileName             = flag.String("i", "", "Specifies the name of the source file or directory of fbt-type files to be compiled.")
	outLocation            = flag.String("o", "", "Specifies the name of the directory to put output files. If blank, uses current directory")
	topName                = flag.String("t", "", "Specifies the name of the top level fbt-type file. If blank, no top file will be generated.")
	outputLanguage         = flag.String("l", "c", "Specifies the output language for the program.")
	algorithmLanguageCheck = flag.Bool("alc", false, "Enable checking algorithm language compatibility with output language.")
	tcrestUsingSPM         = flag.Bool("tuspm", false, "(Experimental flag) When building for T-CREST processor, will put all FBs into _SPM memory. Also includes -ti")
	tcrestSmartSPM         = flag.Bool("tsspm", false, "(Experimental flag) When building for T-CREST processor, will put BFBs onto SPM before running/evicting them. Also includes -ti.")
	tcrestIncludes         = flag.Bool("ti", false, "(Experimental flag) Include the T-CREST header files in fbtypes.h")
	autoFlatten            = flag.Bool("af", false, "Automatically flatten out CFBs to save memory")
	cvodeEnable            = flag.Bool("cvode", false, "Enable cvode for solving algorithms with 'ODE' and 'ODE_init' in comment field")
	eventMoC               = flag.Bool("eventMoC", false, "Use event-driven MoC instead of synchronous MoC (makes it compliant with IEC61499 revision 2)")
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

	conv, err := iec61499converter.New(*outputLanguage)
	if err != nil {
		fmt.Println("Error creating converter:", err.Error())
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
				fbtFileNames = append(fbtFileNames, name)
			} else {
				//only copy the "extra" files if they don't begin with a '.'
				if name[0] != '.' {
					otherFileNames = append(otherFileNames, name)
				}
			}
		}
	} else {
		fmt.Println("Running in Single mode")
		fbtFileNames = append(fbtFileNames, *inFileName)
	}

	if *algorithmLanguageCheck == false {
		conv.DisableAlgorithmLanguageChecks()
	}

	if *tcrestUsingSPM == true {
		conv.SetTcrestUsingSPM()
		conv.SetTcrestIncludes()
	}

	if *tcrestSmartSPM == true {
		conv.SetTcrestSmartSPM()
		conv.SetTcrestIncludes()
	}

	if *tcrestIncludes == true {
		conv.SetTcrestIncludes()
	}

	if *eventMoC == true {
		conv.SetRunOnECC()
		conv.SetEventQueue()
	}

	if *cvodeEnable == true {
		if err := conv.CvodeEnable(); err != nil {
			fmt.Println("Error: " + err.Error())
			return
		}
	}

	for _, name := range fbtFileNames {
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

	if *autoFlatten {
		fmt.Printf("Flattening FBs...\n")
		if err := conv.Flatten(); err != nil {
			fmt.Printf("Error: %s\n", err.Error())
			return
		}
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

	for _, otherFileName := range otherFileNames {
		fmt.Printf("Duplicating %s\n", otherFileName)

		sourceFile, err := ioutil.ReadFile(fmt.Sprintf("%s%c%s", *inFileName, os.PathSeparator, otherFileName))
		if err != nil {
			fmt.Printf("Error reading file '%s' for duplication: %s\n", otherFileName, err.Error())
			return
		}

		err = ioutil.WriteFile(fmt.Sprintf("%s%c%s", *outLocation, os.PathSeparator, otherFileName), sourceFile, 0644)
		if err != nil {
			fmt.Println("Error during file copy:", err.Error())
			return
		}
	}

}
