#!/usr/bin/env bash

protoc \
  --proto_path=$GOPATH/src \
  --go_out=plugins=grpc:$GOPATH/src \
  github.com/Stymphalian/ikuaki/tools/grpcc/testing/server/server.proto

protoc \
  --proto_path=$GOPATH/src \
  --go_out=plugins=grpc,:$GOPATH/src \
  github.com/Stymphalian/ikuaki/protos/common.proto \
  github.com/Stymphalian/ikuaki/protos/lobby.proto \
  github.com/Stymphalian/ikuaki/protos/world.proto \
  github.com/Stymphalian/ikuaki/protos/agent.proto

protoc \
  --proto_path=$GOPATH/src \
  --go_out=plugins=grpc:$GOPATH/src \
  github.com/Stymphalian/ikuaki/athrun/protos/athrun.proto
