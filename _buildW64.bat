@echo off

echo Building onlineEditing64.exe into out directory
set GOOS=windows
set GOARCH=amd64
go build -i -o out\onlineEditing64.exe onlineEditing_common.go onlineEditing_windows.go
if errorlevel 1 goto end

echo Building setup64.exe into out directory
if not "%INNO_SETUP_HOME%" == "" (
  set "IS_EXE=%INNO_SETUP_HOME%\iscc"
) else (
  set "IS_EXE=iscc"
)
if "%IS_EXE%" == "\iscc" set "%IS_EXE%=iscc"
"%IS_EXE%" installer_setup64.iss
if errorlevel 1 goto end

echo Building all successfully.
:end