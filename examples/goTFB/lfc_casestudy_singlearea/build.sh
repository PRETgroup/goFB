#!/bin/bash

#compile tfb to fbt
echo "Compiling TFBs to FBTs..."
goTFB -i=tfb -o=fbt

#compile fbt to c
echo "Compiling FBTs to C..."
goFB -i=./fbt -o=./c -l=c -t=Top

#compile
echo "Compiling C to binary..."
gcc c/*.c -o lfc_casestudy_singlearea.out -DMAX_TICKS=300001
#run
#./simple.out
