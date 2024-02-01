#!/bin/bash

SRC_DIR=./src
OUT_DIR=../bin

cd $SRC_DIR

mkdir -p $OUT_DIR/macos
GOOS=darwin GOARCH=amd64 go build -o $OUT_DIR/macos/gomove ./main

mkdir -p $OUT_DIR/linux
GOOS=linux GOARCH=amd64 go build -o $OUT_DIR/linux/gomove ./main
