#!/bin/bash
cmd="main"
if [ $1 == $cmd ]
then
    go build -o app main.go
fi
./app
