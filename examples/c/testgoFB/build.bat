REM save current working directory
SET WD=%cd%
REM compile fbt to c and compile the different versions
cd ..\..\..
goFB -i=%WD%\fbt -o=%WD%\c -l=c -t=topFLAT -iem
gcc ./examples/c/testgoFB/c/*.c -o ./examples/c/testgoFB/testgoFB_topFLAT.exe -o1

goFB -i=%WD%\fbt -o=%WD%\c -l=c -t=topFLAT -af -iem
gcc ./examples/c/testgoFB/c/*.c -o ./examples/c/testgoFB/testgoFB_topFLAT_F.exe -o1

goFB -i=%WD%\fbt -o=%WD%\c -l=c -t=topCFB1 -iem
gcc ./examples/c/testgoFB/c/*.c -o ./examples/c/testgoFB/testgoFB_topCFB1.exe -o1

goFB -i=%WD%\fbt -o=%WD%\c -l=c -t=topCFB1 -af -iem
gcc ./examples/c/testgoFB/c/*.c -o ./examples/c/testgoFB/testgoFB_topCFB1_F.exe -o1

goFB -i=%WD%\fbt -o=%WD%\c -l=c -t=topMANY -iem
gcc ./examples/c/testgoFB/c/*.c -o ./examples/c/testgoFB/testgoFB_topMANY.exe -o1

goFB -i=%WD%\fbt -o=%WD%\c -l=c -t=topMANY -af -iem
gcc ./examples/c/testgoFB/c/*.c -o ./examples/c/testgoFB/testgoFB_topMANY_F.exe -o1

cd %WD%
