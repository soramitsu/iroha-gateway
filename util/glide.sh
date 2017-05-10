#!/bin/bash

ARCH=`uname | tr '[:upper:]' '[:lower:]'`-amd64
VERSION=${1}

if type glide >/dev/null 2>&1; then
    echo "already installed"
    exit 0
fi


curl -L -O -s https://github.com/Masterminds/glide/releases/download/${VERSION}/glide-${VERSION}-${ARCH}.tar.gz
tar xfz glide-${VERSION}-${ARCH}.tar.gz -C /tmp
mkdir -p ${GOPATH}/bin
cp /tmp/${ARCH}/glide ${GOPATH}/bin/
rm glide-${VERSION}-${ARCH}.tar.gz
glide -v
