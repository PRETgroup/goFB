REM save current working directory
SET WD=%cd%
REM compile fbt to c and compile the different versions
cd ..\..\..
goFB -i=%WD%\fbt -o=%WD%\c -l=c -t=topFLAT
gcc ./examples/c/testgoFB/c/*.c -o ./examples/c/testgoFB/testgoFB_topFLAT.exe

goFB -i=%WD%\fbt -o=%WD%\c -l=c -t=topFLAT -f
gcc ./examples/c/testgoFB/c/*.c -o ./examples/c/testgoFB/testgoFB_topFLAT_F.exe

goFB -i=%WD%\fbt -o=%WD%\c -l=c -t=topCFB1
gcc ./examples/c/testgoFB/c/*.c -o ./examples/c/testgoFB/testgoFB_topCFB1.exe

goFB -i=%WD%\fbt -o=%WD%\c -l=c -t=topCFB1 -f
gcc ./examples/c/testgoFB/c/*.c -o ./examples/c/testgoFB/testgoFB_topCFB1_F.exe

goFB -i=%WD%\fbt -o=%WD%\c -l=c -t=topMANY
gcc ./examples/c/testgoFB/c/*.c -o ./examples/c/testgoFB/testgoFB_topMANY.exe

goFB -i=%WD%\fbt -o=%WD%\c -l=c -t=topMANY -f
gcc ./examples/c/testgoFB/c/*.c -o ./examples/c/testgoFB/testgoFB_topMANY_F.exe

cd %WD%
