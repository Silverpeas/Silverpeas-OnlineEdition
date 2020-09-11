@echo off

rmdir /q /s out\usr
del /q out\spwebdav_linux.zip

echo Building onlineEditing into out\usr\bin directory
mkdir out\usr\share\applications
mkdir out\usr\bin
copy spwebdav.desktop out\usr\share\applications /Y
set GOOS=linux
set GOARCH=amd64
go build -i -o out\usr\bin\onlineEditing onlineEditing_common.go onlineEditing_linux.go
if errorlevel 1 goto end

echo Making an archive into out directory
if not "%ZIP_HOME%" == "" (
  set "ZIP_EXE=%ZIP_HOME%\7z"
) else (
  set "ZIP_EXE=7z"
)
if "%ZIP_EXE%" == "\7z" set "%ZIP_EXE%=7z"
cd out
"%ZIP_EXE%" a -tzip spwebdav_linux.zip usr\
if errorlevel 1 (
  cd ..
  goto end
) else (
  cd ..
)

echo Building all successfully.
:end
