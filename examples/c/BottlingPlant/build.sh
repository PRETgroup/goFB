#!/bin/bash

#compile fbt to c
goFB -i=./fbt -o=./c -l=c -t=FlexPRET

#compile
clang c/*.c -o BottlingPlant.out
#run
./BottlingPlant.out
