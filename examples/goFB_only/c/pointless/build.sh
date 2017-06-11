#!/bin/bash

#compile fbt to c
goFB -i=./fbt -o=./c -l=c -t=test2

#compile
clang c/*.c -o pointless.out
#run
./pointless.out