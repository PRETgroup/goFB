#!/bin/bash
#save current working directory
WD=$(pwd)
#compile fbt to c
cd ../../..
./goFB -i=$WD/fbt -o=$WD/c -l=c -t=FakePacemakerTop -ti
cd $WD
rm ./c/iec61499_network_top.c
