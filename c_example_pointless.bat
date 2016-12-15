cls
go-iec61499-vhdl.exe -i=.\examples\c\pointless\fbt -o=.\examples\c\pointless\c -l=c -t=test2
gcc examples\c\pointless\c\*.c -o examples\c\pointless\c\pointless.exe
.\examples\c\pointless\c\pointless.exe