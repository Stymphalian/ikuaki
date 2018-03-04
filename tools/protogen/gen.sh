#!/usr/bin/env bash

protoc \
  --proto_path=$GOPATH/src/google.golang.org/grpc/reflection/ \
  --proto_path=$GOPATH/src \
  --go_out=plugins=grpc,\
Mgrpc_reflection_v1alpha/reflection.proto=google.golang.org/grpc/reflection/grpc_reflection_v1alpha,\
:$GOPATH/src \
  github.com/Stymphalian/ikuaki/api/protos/ikuaki.proto

protoc \
  --proto_path=$GOPATH/src \
  --go_out=plugins=grpc:$GOPATH/src \
  github.com/Stymphalian/ikuaki/api/protosfat/fattyghost.proto

protoc \
  --proto_path=$GOPATH/src \
  --go_out=plugins=grpc:$GOPATH/src \
  github.com/Stymphalian/ikuaki/tools/grpcc/testing/server/server.proto