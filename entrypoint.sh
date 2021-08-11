#!/bin/bash
cmd="main"
if [ $1 == $cmd ]
then
    cd ./goutils
    go build -o app main.go
fi
./app
