#!/bin/bash
#save current working directory
WD=$(pwd)
#compile fbt to c
goFB -i=$WD/fbt -o=$WD/c -l=c -t=_CB_TCREST -ti
cd $WD
rm ./c/top.c
