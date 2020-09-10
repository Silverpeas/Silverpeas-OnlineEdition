#!/usr/bin/env bash

test -d out || mkdir out
mkdir -p out/usr/share/applications/
mkdir out/usr/bin
cp spwebdav.desktop out/usr/share/applications
GOOS=linux GOARCH=amd64 go build -i -o out/usr/bin/onlineEditing onlineEditing_common.go onlineEditing_linux.go
pushd out || return
tar zcvf spwebdav.tgz usr
popd || return
