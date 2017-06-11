#!/bin/bash
#save current working directory
WD=$(pwd)
#compile fbt to c
cd ../../..
./goFB -i=$WD/fbt -o=$WD/c -l=c -t=_TCREST -tuspm
cd $WD
