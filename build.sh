#!/bin/sh

PROTOC="${PROTOCBIN:-protoc}"

$PROTOC -I=defs --python_out=build/python/sizematch/protobuf defs/items.proto
#$PROTOC -I=defs --js_out=build/js defs/items.proto
#$PROTOC -I=defs --go_out=build/go defs/items.proto
