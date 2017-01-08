#!/bin/bash
#save current working directory
WD=$(pwd)
#compile fbt to c
cd ../../..
./goFB -i=$WD/fbt -o=$WD/c -l=c -t=test2
cd $WD
#compile
clang c/*.c -o pointless.out
#run
./pointless.out
