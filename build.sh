#!/bin/bash

path=/tmp/$(uuidgen)

echo Using $path

mkdir $path

export GOPATH=$path

echo Go Getting github.com/DHowett/go-plist
go get "github.com/DHowett/go-plist"

mkdir -p build

echo Building
go build -o build/clientidentifier clientidentifier.go


echo Cleaning Up
rm -Rf $path
