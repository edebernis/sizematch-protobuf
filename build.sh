#!/bin/sh

if [ -z "$1" ]
  then
    echo "No version supplied"
    exit 1
fi

VERSION=$1
GIT="${GITBIN:-git}"
PROTOC="${PROTOCBIN:-protoc}"
GITPKG="${GITPKGBIN:-gitpkg}"

# Python
$PROTOC -I=defs --python_out=build/python/sizematch/protobuf defs/*/*.proto
$GIT tag $VERSION
$GIT push origin $VERSION

# JS
$PROTOC -I=defs --js_out=import_style=commonjs,binary:build/js defs/*/*.proto
( cd build/js && $GITPKG publish )

# GO
$PROTOC -I=defs --go_opt=paths=source_relative --go_out=build/go defs/*/*.proto
$GIT tag build/go/$VERSION
$GIT push origin build/go/$VERSION
