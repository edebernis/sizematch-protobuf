#!/bin/sh

PROTOC="${PROTOCBIN:-protoc}"
GITPKG="${GITPKGBIN:-gitpkg}"

# Python
$PROTOC -I=defs --python_out=build/python/sizematch/protobuf defs/*.proto

# JS
$PROTOC -I=defs --js_out=import_style=commonjs,binary:build/js defs/*.proto
( cd build/js && $GITPKG publish )

# GO
#$PROTOC -I=defs --go_out=build/go defs/*.proto
