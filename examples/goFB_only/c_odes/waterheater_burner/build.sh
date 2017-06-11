#!/bin/bash
#save current working directory
WD=$(pwd)
#compile fbt to c
cd ../../..
./goFB -i=$WD/fbt -o=$WD/c -l=c -t=top -cvode
cd $WD
#compile
#clang c/*.c -o top.out
gcc -o0 -Wall c/*.c /home/hammond/cvode/inst/lib/libsundials_cvode.a /home/hammond/cvode/inst/lib/libsundials_nvecserial.a -I/home/hammond/cvode/inst/include -lm -o ha.out -Wfatal-errors
#run
#./top.out


