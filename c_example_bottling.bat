cls
go-iec61499-vhdl.exe -i=.\examples\c\BottlingPlant\fbt -o=.\examples\c\BottlingPlant\c -l=c -t=FlexPRET
gcc examples\c\BottlingPlant\c\*.c -o examples\c\BottlingPlant\c\BottlingPlant.exe
.\examples\c\BottlingPlant\c\BottlingPlant.exe