set GOOS=windows
set GOARCH=amd64
go build -i -o onlineEditing64.exe onlineEditing_common.go onlineEditing_windows.go