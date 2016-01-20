#!/bin/bash

export APPNAME="GetMovieInfos"
export GOPATH="$(cd "$(dirname "$1")"; pwd)/$(basename "$1")"

echo Set GOPATH to ${GOPATH}

if [ ! -d ${GOPATH}/bin ]
then
    mkdir ${GOPATH}/bin
fi

if [ ! -d ${GOPATH}/pkg ]
then
    mkdir ${GOPATH}/pkg
fi

cd ${GOPATH}/src
echo go get
go get
PACKAGES=`find . -type d | grep -v  "\/\." | tr -s '\n' ' '`
echo go fmt ${PACKAGES}
go fmt ${PACKAGES} 2> /dev/null
echo go build -o ${GOPATH}/bin/$APPNAME
go build -o ${GOPATH}/bin/$APPNAME