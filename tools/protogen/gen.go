package protogen

//go:generate  ./gen.sh

// protosfat/protos gen.go
// go:generate  protoc --proto_path=$GOPATH/src/github.com  --go_out=plugins=grpc:$GOPATH/src Stymphalian/ikuaki/api/protosfat/fattyghost.proto

// api/protos gen.go
// go:generate  protoc --proto_path=$GOPATH/src/google.golang.org/grpc/reflection/ --proto_path=$GOPATH/src/github.com --proto_path=$GOPATH/src --go_out=plugins=grpc,Mgrpc_reflection_v1alpha/reflection.proto=google.golang.org/grpc/reflection/grpc_reflection_v1alpha,Mfattyghost=github.com/Stymphalian/ikuaki/api/protosfat:$GOPATH/src Stymphalian/ikuaki/api/protos/ikuaki.proto
//  go:generate  protoc --proto_path=$GOPATH/src/google.golang.org/grpc/reflection/ --proto_path=$GOPATH/src/github.com --proto_path=$GOPATH/src --go_out=plugins=grpc,Mgrpc_reflection_v1alpha/reflection.proto=google.golang.org/grpc/reflection/grpc_reflection_v1alpha,MStymphalian.ikuaki.api.protosfat=github.com/Stymphalian/ikuaki/api/protosfat:$GOPATH/src Stymphalian/ikuaki/api/protos/ikuaki.proto
//  go:generate  protoc --proto_path=$GOPATH/src/google.golang.org/grpc/reflection/ --proto_path=$GOPATH/src/github.com --proto_path=$GOPATH/src --go_out=plugins=grpc,Mgrpc_reflection_v1alpha/reflection.proto=google.golang.org/grpc/reflection/grpc_reflection_v1alpha,MStymphalian.ikuaki.api.protosfat=github.com/Stymphalian.ikuaki.api.protosfat:$GOPATH/src Stymphalian/ikuaki/api/protos/ikuaki.proto

// --proto_path=$GOPATH/src/google.golang.org/grpc/reflection/
// --proto_path=$GOPATH/src/github.com
// --proto_path=$GOPATH/src
// --go_out=plugins=grpc,
// Mgrpc_reflection_v1alpha/reflection.proto=google.golang.org/grpc/reflection/grpc_reflection_v1alpha,
// MStymphalian.ikuaki.api.protosfat=github.com/Stymphalian.ikuaki.api.protosfat:$GOPATH/src
// $GOPATH/src/github.com/Stymphalian/ikuaki/api/protos/ikuaki.proto

//protoc
// --proto_path=$GOPATH/src/google.golang.org/grpc/reflection/
// --proto_path=$GOPATH/src/github.com  <-- right here we need to find the proto
//                                          so that file name registered matches
//                                          import paths.
// --proto_path=$GOPATH/src             <-- This is misc so that we can find
//                                          any other protos.
// --go_out=plugins=grpc
//   ,Mgrpc_reflection_v1alpha/reflection.proto=google.golang.org/grpc/reflection/grpc_reflection_v1alpha
//   ,MStymphalian.ikuaki.api.protosfat=github.com/Stymphalian.ikuaki.api.protosfat
//   :$GOPATH/src
//  Stymphalian/ikuaki/api/protos/ikuaki.proto
