#!/bin/bash
#save current working directory
WD=$(pwd)
#compile fbt to c
cd ../../..
./goFB -i=$WD/fbt -o=$WD/c -l=c -t=FlexPRET
cd $WD
#compile
#clang c/*.c -o BottlingPlant.out
#run
#./BottlingPlant.out
