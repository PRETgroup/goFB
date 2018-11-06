#!/bin/bash

#compile fbt to c
goFB -i=./fbt -o=./verilog -l=verilog -t=VVI_Pacemaker

#compile
#clang c/*.c -o BottlingPlant.out
#run
#./BottlingPlant.out
