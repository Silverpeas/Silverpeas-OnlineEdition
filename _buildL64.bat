set GOOS=linux
set GOARCH=amd64
go build -i -o out\onlineEditing_linux onlineEditing_common.go onlineEditing_linux.go
