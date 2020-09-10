set GOOS=darwin
set GOARCH=amd64
go build -i -o out\onlineEditing_darwin onlineEditing_common.go onlineEditing_darwin.go
