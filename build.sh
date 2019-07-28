#! /bin/bash
# Copyright 2018 Kuei-chun Chen. All rights reserved.

# dep init

DEP=`which dep`

if [ "$DEP" == "" ]; then
    echo "dep command not found"
    exit
fi

$DEP ensure -update
export version="$(git symbolic-ref --short HEAD)-$(date "+%Y%m%d.%s")"
export ver="1.2.0"
export version="v${ver}-$(date "+%Y%m%d")"
mkdir -p build
env GOOS=linux GOARCH=amd64 go build -ldflags "-X main.version=$version" -o build/keyhole-linux-x64 keyhole.go
env GOOS=darwin GOARCH=amd64 go build -ldflags "-X main.version=$version" -o build/keyhole-osx-x64 keyhole.go
env GOOS=windows GOARCH=amd64 go build -ldflags "-X main.version=$version" -o build/keyhole-win-x64.exe keyhole.go

docker build -t simagix/keyhole:latest -t simagix/keyhole:${ver} .
