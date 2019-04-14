# c_tcrest examples

These examples were used in the ISORC 2018 publication "Faster Function Blocks for Precision Timed Industrial Automation".

It should be relatively straightforward to run these examples yourself, using the Makefile in this directory.

`make build PROJECT=[project]` is the syntax to build an *.elf for download. It will load it into the t-crest/patmos/tmp directory so that you can run `make download APP=[project]` from there to put it on the processor.

`make wcet PROJECT=[project] FUNCTION=[function] T_DELAY_CYCLES=[cycles]` is the syntax for building a worst-case execution time (WCET) report. All the existing reports are attached.

For cycle numbers, use 83 for the T-CREST four-core platform, and 21 for the Patmos single-core platform. These values were derived from the info in the T-CREST manual Section 6.8.

## A note when running the designs:

For all projects, T-CREST and Patmos were slightly customised:
Three memory mapped registers were created (on T-CREST, there is three for each core, for Patmos, just three in total).

The addresses and purposes are described here:
```
#define LED 		( *( ( volatile _IODEV unsigned * )	0xF0090000 ) )
#define HEX 		( *( ( volatile _IODEV unsigned * )	0xF0070000 ) )
#define SWITCHES ( *( ( volatile _IODEV unsigned * ) 0xF0060000 ) )
```

The project uses the DE2-115 board. For the T-CREST four-core, they were each given two 7-segment HEX displays, four LEDs, and four SWITCHES.
For Patmos, all were mapped to the single core.

