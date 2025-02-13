#!/bin/sh

cd ../cmd || exit
export CGO_ENABLED=1
export GOOS=windows
export GOARCH=amd64

go build -ldflags="-s -w" -o ../dist/atlas.exe
