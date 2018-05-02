## goTFB: an IEC 61499 compliant transpiler for Textual Function Blocks

### goTFB is not ready for use yet! It is still under heavy development.

IEC 61499 is the newest standard for the design of industrial automation systems. While many tools for IEC61499 have been developed, they are primarily GUI and WYSIWYG based.

While in and of itself this is not a bad thing, there are three main drawbacks to this approach: 
1. When rendering complex systems in WYSIWYGs, the visual representations naturally tend to become complex, encapsulating lots of minute details. These rapidly can become difficult to parse and understand.
2. GUI-based tools tend to be more complex, less platform-independent, more intensive, and less interoperable than command-line based alternatives.
3. Visual representations tend to have complex filetypes, which may not play nicely with version control software. Text-based representations of FBs allows for simpler and clearer integrations.

Just as Structured Text (ST) is a valid alternative for representing Ladder Logic and/or Function Block control systems in the IEC 61131-3 standard, Textual Function Blocks seek to become a valid alternative for those working with the newer IEC 61499 standard.  

### This is not a compiler! 

Please note that goTFB is *not* a compiler. Rather, it transpiles the TFB language into industry standard IEC 61499 XML, which can then be compiled using a suitable tool of your choice. 

All examples in this repository can be compiled using the [goFB](https://github.com/kiwih/goFB) IEC61499 Function Block compiler.

### Language guide

A full language guide is available at [The TFB Language Specification](docs/TFB/README.md)

## A Quick Example

A FB that simply echos what it receives back out again:

```
basicFB echo;

interface of echo {
	in event rx;
	out event tx;
	
	in with rx byte rxDat;
	out with tx byte txDat;
}

architecture of echo {
	state await {
		emit tx;
		run in "C" `me->txDat = me->rxDat`;
		-> await on rx;
	}
}
```

## Usage

### Building

Your computer will need to be set up to develop in the Go programming language.

```
go get github.com/kiwih/goTFB
cd goTFB
go build
```

### Using

```
goTFB -i=examples/[project]/tfbs -o=examples/[project]/tfbs
```

## Currently Supports

- [x] Basic Function Blocks
- [x] Service Interface Function Blocks
- [x] Composite Function Blocks
- [ ] Hybrid Function Blocks
- [ ] Resources
- [ ] Devices
- [ ] Systems

## Testing

The main (root) package is just a wrapper for the tfbparser library to give it access to files. 
You can test the parser by navigating to the tfbparser folder:

```
cd tfbparser
go test -coverprofile=cover.out
go tool cover -html=cover.out
```

Also included is a very thorough IEC61499 XML package.
```
cd iec61499
go test -coverprofile=cover.out
go tool cover -html=cover.out
```
