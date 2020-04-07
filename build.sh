#!/bin/sh

VERSION=$1
PROTOC="${PROTOCBIN:-protoc}"
GITPKG="${GITPKGBIN:-gitpkg}"

# Python
$PROTOC -I=defs --python_out=build/python/sizematch/protobuf defs/*/*.proto
git tag $VERSION
git push origin $VERSION

# JS
$PROTOC -I=defs --js_out=import_style=commonjs,binary:build/js defs/*/*.proto
( cd build/js && $GITPKG publish )

# GO
$PROTOC -I=defs --go_opt=paths=source_relative --go_out=build/go defs/*/*.proto
git tag build/go/$VERSION
git push origin build/go/$VERSION
