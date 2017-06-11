#!/bin/bash
#save current working directory
WD=$(pwd)

# build goTFB
go build -o ./bin/goTFB -i ./goTFB

# build goFB
go build -o ./bin/goFB -i ./goFB 

