// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/Stymphalian/ikuaki/athrun/protos/athrun.proto

/*
Package Stymphalian_ikuaki_athrun is a generated protocol buffer package.

It is generated from these files:
	github.com/Stymphalian/ikuaki/athrun/protos/athrun.proto

It has these top-level messages:
	BuildRequest
	BuildResponse
	EnvRequest
	EnvResponse
*/
package Stymphalian_ikuaki_athrun

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type BuildRequest struct {
	// REQUIRED
	// A filepath relative to the GOPATH which specified the package to build.
	// The package should point to a binary to be built.
	Filepath string `protobuf:"bytes,1,opt,name=filepath" json:"filepath,omitempty"`
	// REQUIRED
	// The name you want to give the binary
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	// OPTIONAL
	// Additional 'go build' arguments to pass to the build
	Args []string `protobuf:"bytes,3,rep,name=args" json:"args,omitempty"`
	// OPTIONAL
	// If binary with the given name already exists are we allowed to overwrite?
	Overwrite bool `protobuf:"varint,4,opt,name=overwrite" json:"overwrite,omitempty"`
}

func (m *BuildRequest) Reset()                    { *m = BuildRequest{} }
func (m *BuildRequest) String() string            { return proto.CompactTextString(m) }
func (*BuildRequest) ProtoMessage()               {}
func (*BuildRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *BuildRequest) GetFilepath() string {
	if m != nil {
		return m.Filepath
	}
	return ""
}

func (m *BuildRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *BuildRequest) GetArgs() []string {
	if m != nil {
		return m.Args
	}
	return nil
}

func (m *BuildRequest) GetOverwrite() bool {
	if m != nil {
		return m.Overwrite
	}
	return false
}

type BuildResponse struct {
	// The full filepath for where you can find the built binary
	OutputFilepath string `protobuf:"bytes,1,opt,name=output_filepath,json=outputFilepath" json:"output_filepath,omitempty"`
	// The stdout from the running the command
	Stdout string `protobuf:"bytes,2,opt,name=stdout" json:"stdout,omitempty"`
	// The command which was run to build the binary.
	Command []string `protobuf:"bytes,3,rep,name=command" json:"command,omitempty"`
}

func (m *BuildResponse) Reset()                    { *m = BuildResponse{} }
func (m *BuildResponse) String() string            { return proto.CompactTextString(m) }
func (*BuildResponse) ProtoMessage()               {}
func (*BuildResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *BuildResponse) GetOutputFilepath() string {
	if m != nil {
		return m.OutputFilepath
	}
	return ""
}

func (m *BuildResponse) GetStdout() string {
	if m != nil {
		return m.Stdout
	}
	return ""
}

func (m *BuildResponse) GetCommand() []string {
	if m != nil {
		return m.Command
	}
	return nil
}

type EnvRequest struct {
	// OPTIONAL
	// The specific arguments to returns. Leave as empty to receive all the
	// environemnt variables from the server
	// Valid flags to request:
	//   ADDR
	//   GOPATH
	//   OUTPUT_DIR
	Args []string `protobuf:"bytes,1,rep,name=args" json:"args,omitempty"`
}

func (m *EnvRequest) Reset()                    { *m = EnvRequest{} }
func (m *EnvRequest) String() string            { return proto.CompactTextString(m) }
func (*EnvRequest) ProtoMessage()               {}
func (*EnvRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *EnvRequest) GetArgs() []string {
	if m != nil {
		return m.Args
	}
	return nil
}

type EnvResponse struct {
	// Map from environment name to value.
	Env map[string]string `protobuf:"bytes,1,rep,name=env" json:"env,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *EnvResponse) Reset()                    { *m = EnvResponse{} }
func (m *EnvResponse) String() string            { return proto.CompactTextString(m) }
func (*EnvResponse) ProtoMessage()               {}
func (*EnvResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *EnvResponse) GetEnv() map[string]string {
	if m != nil {
		return m.Env
	}
	return nil
}

func init() {
	proto.RegisterType((*BuildRequest)(nil), "Stymphalian.ikuaki.athrun.BuildRequest")
	proto.RegisterType((*BuildResponse)(nil), "Stymphalian.ikuaki.athrun.BuildResponse")
	proto.RegisterType((*EnvRequest)(nil), "Stymphalian.ikuaki.athrun.EnvRequest")
	proto.RegisterType((*EnvResponse)(nil), "Stymphalian.ikuaki.athrun.EnvResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Athrun service

type AthrunClient interface {
	// An RPC call to do a 'go build' for a single binary
	Build(ctx context.Context, in *BuildRequest, opts ...grpc.CallOption) (*BuildResponse, error)
	// Query the server for the environment in which it is building.
	// Things like the GOPATH, and the OUTPUT directory
	Env(ctx context.Context, in *EnvRequest, opts ...grpc.CallOption) (*EnvResponse, error)
}

type athrunClient struct {
	cc *grpc.ClientConn
}

func NewAthrunClient(cc *grpc.ClientConn) AthrunClient {
	return &athrunClient{cc}
}

func (c *athrunClient) Build(ctx context.Context, in *BuildRequest, opts ...grpc.CallOption) (*BuildResponse, error) {
	out := new(BuildResponse)
	err := grpc.Invoke(ctx, "/Stymphalian.ikuaki.athrun.Athrun/Build", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *athrunClient) Env(ctx context.Context, in *EnvRequest, opts ...grpc.CallOption) (*EnvResponse, error) {
	out := new(EnvResponse)
	err := grpc.Invoke(ctx, "/Stymphalian.ikuaki.athrun.Athrun/Env", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Athrun service

type AthrunServer interface {
	// An RPC call to do a 'go build' for a single binary
	Build(context.Context, *BuildRequest) (*BuildResponse, error)
	// Query the server for the environment in which it is building.
	// Things like the GOPATH, and the OUTPUT directory
	Env(context.Context, *EnvRequest) (*EnvResponse, error)
}

func RegisterAthrunServer(s *grpc.Server, srv AthrunServer) {
	s.RegisterService(&_Athrun_serviceDesc, srv)
}

func _Athrun_Build_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BuildRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AthrunServer).Build(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Stymphalian.ikuaki.athrun.Athrun/Build",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AthrunServer).Build(ctx, req.(*BuildRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Athrun_Env_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EnvRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AthrunServer).Env(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Stymphalian.ikuaki.athrun.Athrun/Env",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AthrunServer).Env(ctx, req.(*EnvRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Athrun_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Stymphalian.ikuaki.athrun.Athrun",
	HandlerType: (*AthrunServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Build",
			Handler:    _Athrun_Build_Handler,
		},
		{
			MethodName: "Env",
			Handler:    _Athrun_Env_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/Stymphalian/ikuaki/athrun/protos/athrun.proto",
}

func init() {
	proto.RegisterFile("github.com/Stymphalian/ikuaki/athrun/protos/athrun.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 362 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0x4f, 0x4b, 0xfb, 0x40,
	0x14, 0x6c, 0x9a, 0xb6, 0xbf, 0xf6, 0xf5, 0xe7, 0x1f, 0x16, 0x91, 0x58, 0x3c, 0x84, 0x80, 0xb6,
	0xa7, 0x04, 0x2a, 0x48, 0xd1, 0x53, 0x0b, 0xd1, 0x7b, 0x04, 0x0f, 0x22, 0xc8, 0xb6, 0x5d, 0x9b,
	0xb5, 0xc9, 0x6e, 0x4c, 0x76, 0x23, 0xfd, 0x06, 0x7e, 0x2a, 0x3f, 0x9b, 0x64, 0xb3, 0xb1, 0x45,
	0xa8, 0xd5, 0xdb, 0xcc, 0x30, 0xfb, 0xde, 0xbc, 0x49, 0x60, 0xb4, 0xa0, 0x22, 0x94, 0x53, 0x77,
	0xc6, 0x63, 0xef, 0x4e, 0xac, 0xe2, 0x24, 0xc4, 0x11, 0xc5, 0xcc, 0xa3, 0x4b, 0x89, 0x97, 0xd4,
	0xc3, 0x22, 0x4c, 0x25, 0xf3, 0x92, 0x94, 0x0b, 0x9e, 0x69, 0xe6, 0x2a, 0x86, 0x4e, 0x36, 0xec,
	0x6e, 0x69, 0x77, 0x4b, 0x83, 0x93, 0xc0, 0xff, 0x89, 0xa4, 0xd1, 0x3c, 0x20, 0xaf, 0x92, 0x64,
	0x02, 0xf5, 0xa0, 0xfd, 0x4c, 0x23, 0x92, 0x60, 0x11, 0x5a, 0x86, 0x6d, 0x0c, 0x3a, 0xc1, 0x17,
	0x47, 0x08, 0x1a, 0x0c, 0xc7, 0xc4, 0xaa, 0x2b, 0x5d, 0xe1, 0x42, 0xc3, 0xe9, 0x22, 0xb3, 0x4c,
	0xdb, 0x2c, 0xb4, 0x02, 0xa3, 0x53, 0xe8, 0xf0, 0x9c, 0xa4, 0x6f, 0x29, 0x15, 0xc4, 0x6a, 0xd8,
	0xc6, 0xa0, 0x1d, 0xac, 0x05, 0xe7, 0x05, 0xf6, 0xf4, 0xc6, 0x2c, 0xe1, 0x2c, 0x23, 0xa8, 0x0f,
	0x07, 0x5c, 0x8a, 0x44, 0x8a, 0xa7, 0x6f, 0x9b, 0xf7, 0x4b, 0xf9, 0xa6, 0xda, 0x7f, 0x0c, 0xad,
	0x4c, 0xcc, 0xb9, 0x14, 0x3a, 0x81, 0x66, 0xc8, 0x82, 0x7f, 0x33, 0x1e, 0xc7, 0x98, 0xcd, 0x75,
	0x8c, 0x8a, 0x3a, 0x36, 0x80, 0xcf, 0xf2, 0xea, 0xb6, 0x2a, 0xab, 0xb1, 0xce, 0xea, 0xbc, 0x1b,
	0xd0, 0x55, 0x16, 0x1d, 0x66, 0x0c, 0x26, 0x61, 0xb9, 0xb2, 0x74, 0x87, 0x9e, 0xbb, 0xb5, 0x38,
	0x77, 0xe3, 0x51, 0x81, 0x7d, 0x26, 0xd2, 0x55, 0x50, 0xbc, 0xed, 0x5d, 0x42, 0xbb, 0x12, 0xd0,
	0x21, 0x98, 0x4b, 0xb2, 0xd2, 0xf7, 0x14, 0x10, 0x1d, 0x41, 0x33, 0xc7, 0x91, 0xac, 0x5a, 0x2c,
	0xc9, 0x55, 0x7d, 0x64, 0x0c, 0x3f, 0x0c, 0x68, 0x8d, 0xd5, 0x70, 0xf4, 0x08, 0x4d, 0xd5, 0x11,
	0xea, 0xff, 0x90, 0x60, 0xf3, 0xbb, 0xf5, 0x06, 0xbb, 0x8d, 0x65, 0x58, 0xa7, 0x86, 0xee, 0xc1,
	0xf4, 0x59, 0x8e, 0xce, 0x76, 0x5d, 0x57, 0x4e, 0x3e, 0xff, 0x5d, 0x09, 0x4e, 0x6d, 0x72, 0xfb,
	0xe0, 0xff, 0xe1, 0x17, 0xbd, 0xde, 0x3a, 0x76, 0xda, 0x52, 0x8e, 0x8b, 0xcf, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x18, 0xfa, 0x0a, 0x31, 0xf2, 0x02, 0x00, 0x00,
}
