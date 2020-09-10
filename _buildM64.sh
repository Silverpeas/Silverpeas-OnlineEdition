#!/usr/bin/env bash

GOOS=darwin GOARCH=amd64 go build -i -o out/onlineEditing onlineEditing_common.go onlineEditing_darwin.go
