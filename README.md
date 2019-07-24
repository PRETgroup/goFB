# goFB, an IEC61499 Function Block toolchain

This repository provides all source code for the goFB project, which aims to supply a complete toolchain for IEC61499 design, validation, and compilation. It typically generates C code from your Function Block networks which executes synchronously (the default behaviour), but also has experimental support for a more-standard event-driven network. Both approaches generate code amenable to static timing analysis given suitable target architectures.

Furthermore, there is yet additional experimental support for compiling Function Block networks to Verilog for synchronous execution in reconfigurable hardware.

goFB was inspired by [FBC](https://www.researchgate.net/publication/224453746_Efficient_implementation_of_IEC_61499_function_blocks), [BlokIDE](http://pretzel.ece.auckland.ac.nz/#!research?project=iec61499), and Structured Text.

There are two major components to this toolchain:
* *goFB/goTFB* provides a commandline text-based methodology for design of IEC61499 function blocks and networks, and
* *goFB/goFB* provides the compiler goFB for IEC61499 to C.

goTFB creates intermediary IEC61499-compliant XML files as its output, and goFB uses those same intermediary IEC61499-compliant XML files as its input.
This means that goTFB may be paired with any suitable compiler (e.g. FBC), and goFB may be paired with any other IEC61499 design tool (e.g. BlokIDE).

## Support

The goFB toolchain currently supports the following aspects of the IEC61499 standard
- [x] Basic Function Blocks
- [x] Service Interface Function Blocks
- [x] Composite Function Blocks
- [x] Resources *(partial - goFB but not goTFB)*
- [ ] Devices
- [ ] Systems

There is also some work into IEC61499 extensions:
- [x] Hybrid Function Blocks *(partial - goFB but not goTFB)*
- [ ] Enforcer Function Blocks

## Examples

Examples can be found in the `examples` directory. 

For complete usage of the toolchain, refer to the examples in the `examples/goTFB` directory. `examples/goFB_only` has examples that utilize the compiler in different ways, such as adding CVODE, and building for the T-CREST platform.

## Output languages

Primarily, this tool is designed to compile IEC61499 to C. 

There is also Verilog support for compiling BFBs/CFBs for FPGAs. 

## Build

The goFB toolchain is designed using the [Go](https://golang.org) programming language. 

To compile it, firstly you will need to install/set up Go version 1.8+.

Then, you can `go get` this repository, and run `bash build.sh`.

You will need to add the `bin` directory of the repostory to your `$PATH`, for instance, in `~/.profile`.

## Publications

The goFB toolchain was used in the following publications:

* ACM-IEEE MEMOCODE 2017 publication [Simulation of cyber-physical systems using IEC61499](https://dl.acm.org/citation.cfm?id=3127052). 
  * The examples used in this paper can be found under `examples/goFB_only/c_odes`.

* IEEE ISORC 2018 paper [Faster Function Blocks for Precision Timed Industrial Automation](https://ieeexplore.ieee.org/abstract/document/8421148/). 
  * The examples for this paper can be found under `examples/goFB_only/c_tcrest`.

* IEIE-IEEE ICEIC 2019 paper "Synthesizing IEC61499 Function Blocks to Hardware" (IEEEXplore link not yet available).
  * The examples for this paper can be found under `examples/hw_compilation`. 

For the purposes of replicating results, these and other examples are kept pre-compiled within the repository for your convenience.
