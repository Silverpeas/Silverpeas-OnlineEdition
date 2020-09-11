#!/usr/bin/env bash

7z &> /dev/null
if [ $? -eq 0 ]; then
  _7Z="7z"
else
  7zr &> /dev/null
  if [ $? -eq 0 ]; then
    _7Z="7zr"
  else
    echo "Missing 7z program. Cannot build a Linux distribution of the OnlineEdition tool"
    exit 1
  fi
fi

test -d out || mkdir out
test -d out/usr && rm -r out/usr &> /dev/null

echo "Building onlineEditing into out/usr/bin directory"
mkdir -p out/usr/share/applications/
mkdir out/usr/bin
cp spwebdav.desktop out/usr/share/applications
GOOS=linux GOARCH=amd64 go build -i -o out/usr/bin/onlineEditing onlineEditing_common.go onlineEditing_linux.go

echo "Making an archive into out directory"

pushd out || return
$_7Z a -r spwebdav.7z usr
popd || return
