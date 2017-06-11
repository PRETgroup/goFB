#!/bin/bash
#save current working directory
WD=$(pwd)
#compile fbt to c
cd ../../..
./goFB -i=$WD/fbt -o=$WD/c -l=c -t=Netw
cd $WD
#compile
gcc c/*.c -o simple.out
#run
#./simple.out
