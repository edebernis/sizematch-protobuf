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

SRC_DIR="defs"
DST_DIR="build"

# Python
$PROTOC -I=$SRC_DIR --python_out=$DST_DIR/python/sizematch/protobuf $SRC_DIR/*/*.proto
$GIT tag $VERSION
$GIT push origin $VERSION

# JS
$PROTOC -I=$SRC_DIR --js_out=import_style=commonjs,binary:$DST_DIR/js $SRC_DIR/*/*.proto
( cd $DST_DIR/js && $GITPKG publish )

# GO
$PROTOC -I=$SRC_DIR --go_opt=paths=source_relative --go_out=$DST_DIR/go $SRC_DIR/*/*.proto
$GIT tag $DST_DIR/go/$VERSION
$GIT push origin $DST_DIR/go/$VERSION
