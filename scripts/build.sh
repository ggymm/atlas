#!/bin/sh

cd ../cmd || exit
export CGO_ENABLED=1
export GOOS=windows
export GOARCH=amd64
#go clean -cache

cd api || exit
go build -ldflags="-s -w" -o ../../dist/atlas-api.exe
cd ..

cd task || exit
go build -ldflags="-s -w" -o ../../dist/atlas-task.exe
cd ..

go build -ldflags="-s -w -H=windowsgui" -o ../dist/atlas.exe
