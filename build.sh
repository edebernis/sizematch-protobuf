#!/bin/sh

if [ -z "$1" ]
  then
    echo "No current version supplied. Semantic version format : vX.Y.Z"
    exit 1
fi
if [ -z "$2" ]
  then
    echo "No new version supplied. Semantic version format : vX.Y.Z"
    exit 1
fi

CURRENT_VERSION=$1
NEW_VERSION=$2

GIT="${GITBIN:-git}"
BUMP2VERSION="${BUMP2VERSIONBIN:-bump2version}"
NPM="${NPMBIN:-npm}"
GITPKG="${GITPKGBIN:-gitpkg}"
PROTOC="${PROTOCBIN:-protoc}"

SRC_DIR="defs"
DST_DIR="build"

# Generate protobuf code
$PROTOC -I=$SRC_DIR --python_out=$DST_DIR/python/sizematch/protobuf $SRC_DIR/*/*.proto
$PROTOC -I=$SRC_DIR --js_out=import_style=commonjs,binary:$DST_DIR/js $SRC_DIR/*/*.proto
$PROTOC -I=$SRC_DIR --go_opt=paths=source_relative --go_out=$DST_DIR/go $SRC_DIR/*/*.proto

# Bump versions
$BUMP2VERSION --current-version $CURRENT_VERSION --new-version $NEW_VERSION patch $DST_DIR/python/setup.py
( cd $DST_DIR/js && $NPM version $NEW_VERSION )
$GIT add $DST_DIR/*
$GIT commit -m "Bump version: ${CURRENT_VERSION} to ${NEW_VERSION}"

# Tag using language-specific packaging syntax
$GIT tag $NEW_VERSION
$GIT tag $DST_DIR/go/$NEW_VERSION

# Publish everything
$GIT push && $GIT push --tags
( cd $DST_DIR/js && $GITPKG publish )
