#!/bin/bash
#save current working directory
WD=$(pwd)
#compile fbt to c

cd ../../..
./goFB -i=$WD/fbt -o=$WD/c -l=c -t=top -cvode
cd $WD

#compile
#clang c/*.c -o top.out
for mMAXTICKS in 5 10 15 20 25 30 35 40 45 50 55 60 65 70 75 80 85 90 95 100
do
	echo $(($mMAXTICKS*100000))
	gcc c/*.c /home/hammond/cvode/inst/lib/libsundials_cvode.a /home/hammond/cvode/inst/lib/libsundials_nvecserial.a -I/home/hammond/cvode/inst/include -lm -o ha.out -Wfatal-errors -DMAX_TICKS=$(($mMAXTICKS*100000))
	./ha.out
done
#run


