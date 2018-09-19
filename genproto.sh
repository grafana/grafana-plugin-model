#!/usr/bin/env bash
#
# Generate all protobuf bindings.
# Run from repository root.
#
# Initial script taken from etcd under the Apache 2.0 license
# File: https://github.com/coreos/etcd/blob/78a5eb79b510eb497deddd1a76f5153bc4b202d2/scripts/genproto.sh

set -e
set -u

if ! [[ $(protoc --version) =~ "3.6.1" ]]; then
        echo "could not find protoc 3.6.1 is it installed + in PATH?"
        exit 255
fi


protoc -I. \
    -I=$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
    --grpc-gateway_out=logtostderr=true:. \
    --swagger_out=logtostderr=true:. \
    --go_out=plugins=grpc:go/datasource/. \
    datasource.proto


protoc -I. \
    -I=$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
    --grpc-gateway_out=logtostderr=true:. \
    --swagger_out=logtostderr=true:. \
    --go_out=plugins=grpc:go/renderer/. \
    renderer.proto

goimports -w go/**/*.pb.go
