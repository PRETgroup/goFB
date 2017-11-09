# goFB, an IEC61499 Function Block toolchain

This repository provides all source code for the goFB project, which aims to supply acomplete toolchain for IEC61499 design, validation, and compilation.

goFB was inspired by [FBC](https://www.researchgate.net/publication/224453746_Efficient_implementation_of_IEC_61499_function_blocks), [BlokIDE](http://timeme.io), and Structured Text.

There are two major components to this toolchain:
* *goFB/goTFB* provides a commandline text-based methodology for design of IEC61499 function blocks and networks, and
* *goFB/goFB* provides the compiler goFB for IEC61499 to C.

goTFB creates intermediary IEC61499-compliant XML files as its output, and goFB uses those same intermediary IEC61499-compliant XML files as its input.
This means that goTFB can be paired with any suitable compiler, and goFB can be paired with any other IEC61499 design tool.

## Support

The goFB toolchain currently supports the following aspects of the IEC61499 standard
- [x] Basic Function Blocks
- [x] Service Interface Function Blocks
- [x] Composite Function Blocks
- [x] Resources *(partial - goFB but not goTFB)*
- [x] Hybrid Function Blocks *(partial - goFB but not goTFB)*
- [ ] Enforcer Function Blocks
- [ ] Devices 
- [ ] Systems

## Examples

Examples can be found in the `examples` directory. 

For complete usage of the toolchain, refer to the examples in the `examples/goTFB` directory. `examples/goFB_only` has examples that utilize the compiler in different ways, such as adding CVODE, and building for the T-CREST platform.

## Build

The goFB toolchain is designed using the [Go](https://golang.org) programming language. 

To compile it, firstly you will need to install/set up Go version 1.8+.

Then, you can `go get` this repository, and run `bash build.sh`.

You will need to add the `bin` directory of the repostory to your `$PATH`, for instance, in `~/.profile`.

## Publications

The goFB toolchain was used in the ACM-IEEE MEMOCODE publication [Simulation of cyber-physical systems using IEC61499](https://dl.acm.org/citation.cfm?id=3127052), and the examples used in this paper can be found under `examples/goFB_only/c_odes`.
