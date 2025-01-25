#!/bin/sh

cd ../cmd || exit

export CGO_ENABLED=1
export GOOS=windows
export GOARCH=amd64
go clean -cache
go build --tags=fts5 -ldflags="-s -w" -o ../atlas.exe



