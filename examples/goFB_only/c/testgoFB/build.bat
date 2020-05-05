REM save current working directory
SET WD=%cd%
REM compile fbt to c and compile the different versions
cd ..\..\..\..\bin

goFB -i=%WD%\fbt -o=%WD%\c -l=c -t=topFLAT 
cd %WD%
gcc ./c/*.c -o "testgoFB_topFLAT.exe" -O1

REM goFB -i=%WD%\fbt -o=%WD%\c -l=c -t=topFLAT -af 
REM gcc ./examples/goFB_only/c/testgoFB/c/*.c -o ./examples/c/testgoFB/testgoFB_topFLAT_F.exe -O1

REM goFB -i=%WD%\fbt -o=%WD%\c -l=c -t=topCFB1
REM gcc ./examples/c/testgoFB/c/*.c -o ./examples/c/testgoFB/testgoFB_topCFB1.exe -O1

REM goFB -i=%WD%\fbt -o=%WD%\c -l=c -t=topCFB1 -af
REM gcc ./examples/c/testgoFB/c/*.c -o ./examples/c/testgoFB/testgoFB_topCFB1_F.exe -O1

REM goFB -i=%WD%\fbt -o=%WD%\c -l=c -t=topMANY
REM gcc ./examples/c/testgoFB/c/*.c -o ./examples/c/testgoFB/testgoFB_topMANY.exe -O1

REM goFB -i=%WD%\fbt -o=%WD%\c -l=c -t=topMANY -af 
REM gcc ./examples/c/testgoFB/c/*.c -o ./examples/c/testgoFB/testgoFB_topMANY_F.exe -O1


cd %WD%



