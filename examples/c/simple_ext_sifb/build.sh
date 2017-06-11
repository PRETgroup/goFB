#!/bin/bash

#compile fbt to c

goFB -i=./fbt -o=./c -l=c -t=Netw

#compile
gcc c/*.c -o simple.out
#run
#./simple.out
