#!/bin/bash
#save current working directory
WD=$(pwd)
#compile fbt to c
cd ../../..
./goFB -i=$WD/fbt -o=$WD/c -l=c -t=top -cvode
cd $WD
#compile
#clang c/*.c -o top.out
#run
#./top.out
