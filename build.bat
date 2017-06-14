@echo off

rem To install prerequisites, run the following commands once:
rem set GOARCH=amd64
rem set GOOS=linux
rem go tool dist install -v pkg/runtime >nul
rem go install -v -a std >nul

if not exist bin mkdir bin
set VERSION="0.12"
set NAME="shinshop"

echo Packing Data
cd webserver
go-bindata-assetfs -ignore=\\.DS_Store -pkg template templates/...
move bindata_assetfs.go template/templatedata.go >nul
go-bindata-assetfs -ignore=\\.DS_Store -pkg webserver web/...
move bindata_assetfs.go webdata.go >nul
go-bindata-assetfs -ignore=\\.DS_Store -pkg rest rest/map/...
move bindata_assetfs.go rest/mapdata.go >nul
cd ..

echo Building Linux
set GOOS=linux
set GOARCH=amd64
go build main.go
move main bin/%NAME%-%VERSION%-linux-x64 >nul

rem To build x86 linux you will need to install mingw and configure those prerequisites
rem set GOOS=linux
rem set GOARCH=386 go build main.go
rem move main bin/%NAME%-%VERSION%-linux-x86 >nul

echo Building Windows
set GOOS=windows
set GOARCH=amd64
go build main.go
move main.exe bin/%NAME%-%VERSION%-windows-x64.exe >nul
set GOOS=windows
set GOARCH=386
go build main.go
move main.exe bin/%NAME%-%VERSION%-windows-x86.exe >nul

echo Building OSX
set GOOS=darwin
set GOARCH=amd64
go build main.go
move main bin/%NAME%-%VERSION%-osx-x64 >nul
