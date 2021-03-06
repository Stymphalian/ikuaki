// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/Stymphalian/ikuaki/tools/grpcc/testing/server/server.proto

/*
Package server is a generated protocol buffer package.

It is generated from these files:
	github.com/Stymphalian/ikuaki/tools/grpcc/testing/server/server.proto

It has these top-level messages:
	CreateWorldReq
	CreateWorldRes
*/
package server

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

type CreateWorldReq struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *CreateWorldReq) Reset()                    { *m = CreateWorldReq{} }
func (m *CreateWorldReq) String() string            { return proto.CompactTextString(m) }
func (*CreateWorldReq) ProtoMessage()               {}
func (*CreateWorldReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CreateWorldReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type CreateWorldRes struct {
	Addr string `protobuf:"bytes,1,opt,name=addr" json:"addr,omitempty"`
}

func (m *CreateWorldRes) Reset()                    { *m = CreateWorldRes{} }
func (m *CreateWorldRes) String() string            { return proto.CompactTextString(m) }
func (*CreateWorldRes) ProtoMessage()               {}
func (*CreateWorldRes) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CreateWorldRes) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

func init() {
	proto.RegisterType((*CreateWorldReq)(nil), "server.CreateWorldReq")
	proto.RegisterType((*CreateWorldRes)(nil), "server.CreateWorldRes")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Lobby service

type LobbyClient interface {
	CreateWorld(ctx context.Context, in *CreateWorldReq, opts ...grpc.CallOption) (*CreateWorldRes, error)
	ClientStream(ctx context.Context, opts ...grpc.CallOption) (Lobby_ClientStreamClient, error)
	ServerStream(ctx context.Context, in *CreateWorldReq, opts ...grpc.CallOption) (Lobby_ServerStreamClient, error)
	BidiStream(ctx context.Context, opts ...grpc.CallOption) (Lobby_BidiStreamClient, error)
}

type lobbyClient struct {
	cc *grpc.ClientConn
}

func NewLobbyClient(cc *grpc.ClientConn) LobbyClient {
	return &lobbyClient{cc}
}

func (c *lobbyClient) CreateWorld(ctx context.Context, in *CreateWorldReq, opts ...grpc.CallOption) (*CreateWorldRes, error) {
	out := new(CreateWorldRes)
	err := grpc.Invoke(ctx, "/server.Lobby/CreateWorld", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lobbyClient) ClientStream(ctx context.Context, opts ...grpc.CallOption) (Lobby_ClientStreamClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Lobby_serviceDesc.Streams[0], c.cc, "/server.Lobby/ClientStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &lobbyClientStreamClient{stream}
	return x, nil
}

type Lobby_ClientStreamClient interface {
	Send(*CreateWorldReq) error
	CloseAndRecv() (*CreateWorldRes, error)
	grpc.ClientStream
}

type lobbyClientStreamClient struct {
	grpc.ClientStream
}

func (x *lobbyClientStreamClient) Send(m *CreateWorldReq) error {
	return x.ClientStream.SendMsg(m)
}

func (x *lobbyClientStreamClient) CloseAndRecv() (*CreateWorldRes, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(CreateWorldRes)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *lobbyClient) ServerStream(ctx context.Context, in *CreateWorldReq, opts ...grpc.CallOption) (Lobby_ServerStreamClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Lobby_serviceDesc.Streams[1], c.cc, "/server.Lobby/ServerStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &lobbyServerStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Lobby_ServerStreamClient interface {
	Recv() (*CreateWorldRes, error)
	grpc.ClientStream
}

type lobbyServerStreamClient struct {
	grpc.ClientStream
}

func (x *lobbyServerStreamClient) Recv() (*CreateWorldRes, error) {
	m := new(CreateWorldRes)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *lobbyClient) BidiStream(ctx context.Context, opts ...grpc.CallOption) (Lobby_BidiStreamClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Lobby_serviceDesc.Streams[2], c.cc, "/server.Lobby/BidiStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &lobbyBidiStreamClient{stream}
	return x, nil
}

type Lobby_BidiStreamClient interface {
	Send(*CreateWorldReq) error
	Recv() (*CreateWorldRes, error)
	grpc.ClientStream
}

type lobbyBidiStreamClient struct {
	grpc.ClientStream
}

func (x *lobbyBidiStreamClient) Send(m *CreateWorldReq) error {
	return x.ClientStream.SendMsg(m)
}

func (x *lobbyBidiStreamClient) Recv() (*CreateWorldRes, error) {
	m := new(CreateWorldRes)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Lobby service

type LobbyServer interface {
	CreateWorld(context.Context, *CreateWorldReq) (*CreateWorldRes, error)
	ClientStream(Lobby_ClientStreamServer) error
	ServerStream(*CreateWorldReq, Lobby_ServerStreamServer) error
	BidiStream(Lobby_BidiStreamServer) error
}

func RegisterLobbyServer(s *grpc.Server, srv LobbyServer) {
	s.RegisterService(&_Lobby_serviceDesc, srv)
}

func _Lobby_CreateWorld_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateWorldReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LobbyServer).CreateWorld(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.Lobby/CreateWorld",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LobbyServer).CreateWorld(ctx, req.(*CreateWorldReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Lobby_ClientStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(LobbyServer).ClientStream(&lobbyClientStreamServer{stream})
}

type Lobby_ClientStreamServer interface {
	SendAndClose(*CreateWorldRes) error
	Recv() (*CreateWorldReq, error)
	grpc.ServerStream
}

type lobbyClientStreamServer struct {
	grpc.ServerStream
}

func (x *lobbyClientStreamServer) SendAndClose(m *CreateWorldRes) error {
	return x.ServerStream.SendMsg(m)
}

func (x *lobbyClientStreamServer) Recv() (*CreateWorldReq, error) {
	m := new(CreateWorldReq)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Lobby_ServerStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(CreateWorldReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(LobbyServer).ServerStream(m, &lobbyServerStreamServer{stream})
}

type Lobby_ServerStreamServer interface {
	Send(*CreateWorldRes) error
	grpc.ServerStream
}

type lobbyServerStreamServer struct {
	grpc.ServerStream
}

func (x *lobbyServerStreamServer) Send(m *CreateWorldRes) error {
	return x.ServerStream.SendMsg(m)
}

func _Lobby_BidiStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(LobbyServer).BidiStream(&lobbyBidiStreamServer{stream})
}

type Lobby_BidiStreamServer interface {
	Send(*CreateWorldRes) error
	Recv() (*CreateWorldReq, error)
	grpc.ServerStream
}

type lobbyBidiStreamServer struct {
	grpc.ServerStream
}

func (x *lobbyBidiStreamServer) Send(m *CreateWorldRes) error {
	return x.ServerStream.SendMsg(m)
}

func (x *lobbyBidiStreamServer) Recv() (*CreateWorldReq, error) {
	m := new(CreateWorldReq)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Lobby_serviceDesc = grpc.ServiceDesc{
	ServiceName: "server.Lobby",
	HandlerType: (*LobbyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateWorld",
			Handler:    _Lobby_CreateWorld_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ClientStream",
			Handler:       _Lobby_ClientStream_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "ServerStream",
			Handler:       _Lobby_ServerStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "BidiStream",
			Handler:       _Lobby_BidiStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "github.com/Stymphalian/ikuaki/tools/grpcc/testing/server/server.proto",
}

func init() {
	proto.RegisterFile("github.com/Stymphalian/ikuaki/tools/grpcc/testing/server/server.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 226 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x72, 0x4d, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x0f, 0x2e, 0xa9, 0xcc, 0x2d, 0xc8, 0x48, 0xcc, 0xc9,
	0x4c, 0xcc, 0xd3, 0xcf, 0xcc, 0x2e, 0x4d, 0xcc, 0xce, 0xd4, 0x2f, 0xc9, 0xcf, 0xcf, 0x29, 0xd6,
	0x4f, 0x2f, 0x2a, 0x48, 0x4e, 0xd6, 0x2f, 0x49, 0x2d, 0x2e, 0xc9, 0xcc, 0x4b, 0xd7, 0x2f, 0x4e,
	0x2d, 0x2a, 0x4b, 0x2d, 0x82, 0x52, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x6c, 0x10, 0x9e,
	0x92, 0x0a, 0x17, 0x9f, 0x73, 0x51, 0x6a, 0x62, 0x49, 0x6a, 0x78, 0x7e, 0x51, 0x4e, 0x4a, 0x50,
	0x6a, 0xa1, 0x90, 0x10, 0x17, 0x4b, 0x5e, 0x62, 0x6e, 0xaa, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67,
	0x10, 0x98, 0x8d, 0xa1, 0xaa, 0x18, 0xa4, 0x2a, 0x31, 0x25, 0xa5, 0x08, 0xa6, 0x0a, 0xc4, 0x36,
	0x9a, 0xc2, 0xc4, 0xc5, 0xea, 0x93, 0x9f, 0x94, 0x54, 0x29, 0x64, 0xcf, 0xc5, 0x8d, 0xa4, 0x5e,
	0x48, 0x4c, 0x0f, 0x6a, 0x37, 0xaa, 0x55, 0x52, 0xd8, 0xc5, 0x8b, 0x95, 0x18, 0x84, 0x9c, 0xb8,
	0x78, 0x9c, 0x73, 0x32, 0x53, 0xf3, 0x4a, 0x82, 0x4b, 0x8a, 0x52, 0x13, 0x73, 0x49, 0x37, 0x41,
	0x83, 0x11, 0x64, 0x46, 0x30, 0x58, 0x92, 0x5c, 0x33, 0x0c, 0x40, 0x66, 0x70, 0x39, 0x65, 0xa6,
	0x64, 0x92, 0xef, 0x0a, 0x03, 0x46, 0x27, 0xc7, 0x28, 0x7b, 0x72, 0xe3, 0xcc, 0x1a, 0x42, 0x25,
	0xb1, 0x81, 0x23, 0xcd, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xf7, 0xc5, 0x69, 0x22, 0xfd, 0x01,
	0x00, 0x00,
}
