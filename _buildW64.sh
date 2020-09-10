#!/usr/bin/env bash

GOOS=windows GOARCH=amd64 go build -i -o out/onlineEditing64.exe onlineEditing_common.go onlineEditing_windows.go
