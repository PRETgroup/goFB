# goFB
## An IEC61499 Function Block compiler
This program is a compiler for IEC61499 Function Blocks. 
Currently, there is support for IEC61499 to C, and experimental support for IEC61499 to VHDL.

goFB was inspired by the [FBC](https://www.researchgate.net/publication/224453746_Efficient_implementation_of_IEC_61499_function_blocks) compiler, and internally functions in a similiar way.

The goal of the goFB project is to create a complete toolchain for IEC61499 construction, validation, and compilation. For construction purposes [goTFB](https://github.com/kiwih/goTFB) is under development, and for compilation purposes this project is under development.

## Support

goFB currently supports the following aspects of the IEC61499 standard
- [x] Basic Function Blocks
- [x] Composite Function Blocks
- [x] Resources
- [x] Devices
- [ ] Systems

### Special Extensions

It also supports an extension known as Hybrid Function Blocks (HFBs), which is currently pending publication. To examine support for HFBs, check out the `examples/c_odes` folder. The `-cvode` flag for goFB enables support for CVODE integration when compiling *ODE Algorithm Language*.

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
  -cvode
    	Enable cvode for solving algorithms with 'ODE' and 'ODE_init' in comment field
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
  -tsspm
    	(Experimental flag) When building for T-CREST processor, will put BFBs onto SPM before running/evicting them. Also includes -ti.
  -tuspm
    	(Experimental flag) When building for T-CREST processor, will put all FBs into _SPM memory. Also includes -ti

```

## Examples

The `examples` directory provides support for creating your own networks with goFB.

Under the `examples/c` directory are the straightforward and simple examples that can be used immediately. Three example networks {BottlingPlant, Pointless, and testgoFB} are provided here.

### Other examples

There are other examples, not recommended for casual use:

`examples/c_odes` provides networks used in testing HFB functionality.

`examples/c_tcrest` provides networks used in testing the support for the T-CREST embedded time-predictable platform.

`examples/vhdl` provides networks used in the experimentation for IEC61499-to-vhdl compilation.



