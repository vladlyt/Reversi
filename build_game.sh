#!/bin/bash
export CC=x86_64-w64-mingw32-gcc
export CXX=x86_64-w64-mingw32-g++
export GOOS=windows
export GOARCH=amd64
export CGO_ENABLED=1

go build -i -o ${1%.*}_${GOOS}_${GOARCH}.exe $1
