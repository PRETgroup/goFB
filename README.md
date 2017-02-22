# goFB
## An IEC61499 Function Block compiler
This program is a compiler for IEC61499 Function Blocks. 
Currently, there is support for IEC61499 to C, and experimental support for IEC61499 to VHDL.

## Build instructions

Build support is for the Windows environment only, as the IEC61499 IDE that is recommended for use with this project is [BlokIDE](http://timeme.io), which only runs on Windows.

Firstly, you will need to have [golang](http://golang.org/dl) installed. This readme does not discuss how to set up your Go workspace.

Then, you will need to have some kind of gcc application installed - I recommend [mingw-64](http://sourceforge.net/projects/mingw-w64/).

Then, navigate to your Go workspace and `go get github.com/kiwih/goFB` this project. 

Once you have acquired it, run `go get -u` then `go build` from within the project directory.

Usage:
```
Usage of goFB:
  -af
    	Automatically flatten out CFBs to save memory
  -alc
    	Enable checking algorithm language compatibility with output language.
  -i string
    	Specifies the name of the source file or directory of fbt-type files to be compiled.
  -l string
    	Specifies the output language for the program. (default "c")
  -o string
    	Specifies the name of the directory to put output files. If blank, uses current directory
  -t string
    	Specifies the name of the top level fbt-type file. If blank, no top file will be generated.
  -ti
    	(Experimental flag) Include the T-CREST header files in fbtypes.h
  -tuspm
    	(Experimental flag) When building for T-CREST processor, will put FBs into _SPM memory, also includes -ti
```

Three example networks {BottlingPlant, Pointless, and testgoFB} can be used with the C compiler. 

There are also experimental networks for use on the multi-core T-CREST architecture. Using these has a more complicated setup and is not recommended for casual use at this time.

