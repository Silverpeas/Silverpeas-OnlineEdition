@echo off

echo Building onlineEditing32.exe into out directory
set GOOS=windows
set GOARCH=386
go build -i -o out\onlineEditing32.exe onlineEditing_common.go onlineEditing_windows.go
if errorlevel 1 goto end

echo Building setup32.exe into out directory
if not "%INNO_SETUP_HOME%" == "" (
  set "IS_EXE=%INNO_SETUP_HOME%\iscc"
) else (
  set "IS_EXE=iscc"
)
if "%IS_EXE%" == "\iscc" set "%IS_EXE%=iscc"
"%IS_EXE%" installer_setup32.iss
if errorlevel 1 goto end

echo Building all successfully.
:end