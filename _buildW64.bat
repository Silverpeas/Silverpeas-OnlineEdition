set GOOS=windows
set GOARCH=amd64
go build -i -o out\onlineEditing64.exe onlineEditing_common.go onlineEditing_windows.go

"%INNOSETUP%"\iscc installer_setup64.iss