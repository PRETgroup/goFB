# goFB
## An IEC61499 Function Block transpiler
This program is a transpiler (i.e. source-to-source) for IEC61499 Function Blocks. 
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
```

Two example networks for use with the C transpiler are presented. You can try them by running 
`c_example_bottling.bat` which presents an example bottling plant,
or
`c_example_pointless.bat` which presents two function blocks passing 4-element arrays between themselves.

