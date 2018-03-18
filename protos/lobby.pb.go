// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/Stymphalian/ikuaki/protos/lobby.proto

package Stymphalian_ikuaki

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

type ServerTypeEnum int32

const (
	ServerTypeEnum_UNKNOWN_SERVER ServerTypeEnum = 0
	ServerTypeEnum_WORLD_SERVER   ServerTypeEnum = 1
	ServerTypeEnum_AGENT_SERVER   ServerTypeEnum = 2
)

var ServerTypeEnum_name = map[int32]string{
	0: "UNKNOWN_SERVER",
	1: "WORLD_SERVER",
	2: "AGENT_SERVER",
}
var ServerTypeEnum_value = map[string]int32{
	"UNKNOWN_SERVER": 0,
	"WORLD_SERVER":   1,
	"AGENT_SERVER":   2,
}

func (x ServerTypeEnum) String() string {
	return proto.EnumName(ServerTypeEnum_name, int32(x))
}
func (ServerTypeEnum) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

// Create
type CreateReq struct {
	BinaryPath string         `protobuf:"bytes,1,opt,name=binary_path,json=binaryPath" json:"binary_path,omitempty"`
	ServerType ServerTypeEnum `protobuf:"varint,2,opt,name=server_type,json=serverType,enum=Stymphalian.ikuaki.ServerTypeEnum" json:"server_type,omitempty"`
}

func (m *CreateReq) Reset()                    { *m = CreateReq{} }
func (m *CreateReq) String() string            { return proto.CompactTextString(m) }
func (*CreateReq) ProtoMessage()               {}
func (*CreateReq) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *CreateReq) GetBinaryPath() string {
	if m != nil {
		return m.BinaryPath
	}
	return ""
}

func (m *CreateReq) GetServerType() ServerTypeEnum {
	if m != nil {
		return m.ServerType
	}
	return ServerTypeEnum_UNKNOWN_SERVER
}

type CreateRes struct {
	Id   string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Addr *Addr  `protobuf:"bytes,2,opt,name=addr" json:"addr,omitempty"`
}

func (m *CreateRes) Reset()                    { *m = CreateRes{} }
func (m *CreateRes) String() string            { return proto.CompactTextString(m) }
func (*CreateRes) ProtoMessage()               {}
func (*CreateRes) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *CreateRes) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *CreateRes) GetAddr() *Addr {
	if m != nil {
		return m.Addr
	}
	return nil
}

// Destroy
type DestroyReq struct {
	Id         string         `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	ServerType ServerTypeEnum `protobuf:"varint,2,opt,name=server_type,json=serverType,enum=Stymphalian.ikuaki.ServerTypeEnum" json:"server_type,omitempty"`
}

func (m *DestroyReq) Reset()                    { *m = DestroyReq{} }
func (m *DestroyReq) String() string            { return proto.CompactTextString(m) }
func (*DestroyReq) ProtoMessage()               {}
func (*DestroyReq) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *DestroyReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *DestroyReq) GetServerType() ServerTypeEnum {
	if m != nil {
		return m.ServerType
	}
	return ServerTypeEnum_UNKNOWN_SERVER
}

type DestroyRes struct {
}

func (m *DestroyRes) Reset()                    { *m = DestroyRes{} }
func (m *DestroyRes) String() string            { return proto.CompactTextString(m) }
func (*DestroyRes) ProtoMessage()               {}
func (*DestroyRes) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

// Listing
type ListReq struct {
	ServerType ServerTypeEnum `protobuf:"varint,1,opt,name=server_type,json=serverType,enum=Stymphalian.ikuaki.ServerTypeEnum" json:"server_type,omitempty"`
}

func (m *ListReq) Reset()                    { *m = ListReq{} }
func (m *ListReq) String() string            { return proto.CompactTextString(m) }
func (*ListReq) ProtoMessage()               {}
func (*ListReq) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func (m *ListReq) GetServerType() ServerTypeEnum {
	if m != nil {
		return m.ServerType
	}
	return ServerTypeEnum_UNKNOWN_SERVER
}

type ListRes struct {
	// A map[server_id]Addr
	// maps the server id to the hostport you can find the server on
	Servers map[string]*Addr `protobuf:"bytes,1,rep,name=servers" json:"servers,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *ListRes) Reset()                    { *m = ListRes{} }
func (m *ListRes) String() string            { return proto.CompactTextString(m) }
func (*ListRes) ProtoMessage()               {}
func (*ListRes) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{5} }

func (m *ListRes) GetServers() map[string]*Addr {
	if m != nil {
		return m.Servers
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateReq)(nil), "Stymphalian.ikuaki.CreateReq")
	proto.RegisterType((*CreateRes)(nil), "Stymphalian.ikuaki.CreateRes")
	proto.RegisterType((*DestroyReq)(nil), "Stymphalian.ikuaki.DestroyReq")
	proto.RegisterType((*DestroyRes)(nil), "Stymphalian.ikuaki.DestroyRes")
	proto.RegisterType((*ListReq)(nil), "Stymphalian.ikuaki.ListReq")
	proto.RegisterType((*ListRes)(nil), "Stymphalian.ikuaki.ListRes")
	proto.RegisterEnum("Stymphalian.ikuaki.ServerTypeEnum", ServerTypeEnum_name, ServerTypeEnum_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Lobby service

type LobbyClient interface {
	Create(ctx context.Context, in *CreateReq, opts ...grpc.CallOption) (*CreateRes, error)
	Destroy(ctx context.Context, in *DestroyReq, opts ...grpc.CallOption) (*DestroyRes, error)
	List(ctx context.Context, in *ListReq, opts ...grpc.CallOption) (*ListRes, error)
}

type lobbyClient struct {
	cc *grpc.ClientConn
}

func NewLobbyClient(cc *grpc.ClientConn) LobbyClient {
	return &lobbyClient{cc}
}

func (c *lobbyClient) Create(ctx context.Context, in *CreateReq, opts ...grpc.CallOption) (*CreateRes, error) {
	out := new(CreateRes)
	err := grpc.Invoke(ctx, "/Stymphalian.ikuaki.Lobby/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lobbyClient) Destroy(ctx context.Context, in *DestroyReq, opts ...grpc.CallOption) (*DestroyRes, error) {
	out := new(DestroyRes)
	err := grpc.Invoke(ctx, "/Stymphalian.ikuaki.Lobby/Destroy", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lobbyClient) List(ctx context.Context, in *ListReq, opts ...grpc.CallOption) (*ListRes, error) {
	out := new(ListRes)
	err := grpc.Invoke(ctx, "/Stymphalian.ikuaki.Lobby/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Lobby service

type LobbyServer interface {
	Create(context.Context, *CreateReq) (*CreateRes, error)
	Destroy(context.Context, *DestroyReq) (*DestroyRes, error)
	List(context.Context, *ListReq) (*ListRes, error)
}

func RegisterLobbyServer(s *grpc.Server, srv LobbyServer) {
	s.RegisterService(&_Lobby_serviceDesc, srv)
}

func _Lobby_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LobbyServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Stymphalian.ikuaki.Lobby/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LobbyServer).Create(ctx, req.(*CreateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Lobby_Destroy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DestroyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LobbyServer).Destroy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Stymphalian.ikuaki.Lobby/Destroy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LobbyServer).Destroy(ctx, req.(*DestroyReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Lobby_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LobbyServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Stymphalian.ikuaki.Lobby/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LobbyServer).List(ctx, req.(*ListReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Lobby_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Stymphalian.ikuaki.Lobby",
	HandlerType: (*LobbyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Lobby_Create_Handler,
		},
		{
			MethodName: "Destroy",
			Handler:    _Lobby_Destroy_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Lobby_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/Stymphalian/ikuaki/protos/lobby.proto",
}

func init() { proto.RegisterFile("github.com/Stymphalian/ikuaki/protos/lobby.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 433 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0xcb, 0x6e, 0xd3, 0x40,
	0x14, 0x86, 0x33, 0xee, 0x25, 0xea, 0x49, 0x14, 0x45, 0x67, 0x15, 0x19, 0x01, 0x91, 0x57, 0x11,
	0x42, 0x0e, 0x84, 0x05, 0xb7, 0x55, 0xd3, 0x5a, 0x14, 0x35, 0x72, 0xd1, 0x24, 0x50, 0x89, 0x4d,
	0x34, 0xae, 0x47, 0x64, 0x94, 0xf8, 0x92, 0x99, 0x49, 0xa5, 0x79, 0x1e, 0xde, 0x8c, 0x27, 0x41,
	0xbe, 0xb5, 0xa1, 0x35, 0xa1, 0x42, 0xdd, 0x79, 0xce, 0xff, 0xfb, 0x3b, 0xc7, 0xc7, 0xff, 0xc0,
	0xab, 0x1f, 0x42, 0x2f, 0x36, 0x81, 0x7b, 0x95, 0x44, 0xc3, 0xa9, 0x36, 0x51, 0xba, 0x60, 0x2b,
	0xc1, 0xe2, 0xa1, 0x58, 0x6e, 0xd8, 0x52, 0x0c, 0x53, 0x99, 0xe8, 0x44, 0x0d, 0x57, 0x49, 0x10,
	0x18, 0x37, 0x3f, 0x20, 0x6e, 0xd9, 0xdc, 0xc2, 0x66, 0xbf, 0x7e, 0x10, 0xe5, 0x2a, 0x89, 0xa2,
	0x24, 0x2e, 0x30, 0xce, 0x1a, 0x8e, 0x4e, 0x24, 0x67, 0x9a, 0x53, 0xbe, 0xc6, 0xe7, 0xd0, 0x0a,
	0x44, 0xcc, 0xa4, 0x99, 0xa7, 0x4c, 0x2f, 0x7a, 0xa4, 0x4f, 0x06, 0x47, 0x14, 0x8a, 0xd2, 0x17,
	0xa6, 0x17, 0x78, 0x02, 0x2d, 0xc5, 0xe5, 0x35, 0x97, 0x73, 0x6d, 0x52, 0xde, 0xb3, 0xfa, 0x64,
	0xd0, 0x19, 0x39, 0xee, 0xfd, 0x51, 0xdc, 0x69, 0x6e, 0x9b, 0x99, 0x94, 0x7b, 0xf1, 0x26, 0xa2,
	0xa0, 0x6e, 0xce, 0xce, 0xe7, 0xdb, 0x96, 0x0a, 0x3b, 0x60, 0x89, 0xb0, 0xec, 0x64, 0x89, 0x10,
	0x5f, 0xc2, 0x3e, 0x0b, 0x43, 0x99, 0xa3, 0x5b, 0xa3, 0x5e, 0x1d, 0xfa, 0x38, 0x0c, 0x25, 0xcd,
	0x5d, 0x0e, 0x03, 0x38, 0xe5, 0x4a, 0xcb, 0xc4, 0x64, 0xe3, 0xdf, 0x65, 0x3d, 0xca, 0xb4, 0xed,
	0xad, 0x16, 0xca, 0xf1, 0xa1, 0x39, 0x11, 0x4a, 0x67, 0xdd, 0xee, 0xd0, 0xc9, 0x7f, 0xd1, 0x7f,
	0x92, 0x0a, 0xa8, 0x70, 0x0c, 0xcd, 0x42, 0x51, 0x3d, 0xd2, 0xdf, 0x1b, 0xb4, 0x46, 0x83, 0x3a,
	0x58, 0xe9, 0x2e, 0xa1, 0xca, 0x8b, 0xb5, 0x34, 0xb4, 0x7a, 0xd1, 0x9e, 0x41, 0x7b, 0x5b, 0xc0,
	0x2e, 0xec, 0x2d, 0xb9, 0x29, 0x77, 0x92, 0x3d, 0xa2, 0x0b, 0x07, 0xd7, 0x6c, 0xb5, 0xe1, 0xff,
	0xdc, 0x70, 0x61, 0xfb, 0x60, 0xbd, 0x23, 0x2f, 0xce, 0xa0, 0xf3, 0xe7, 0x37, 0x20, 0x42, 0xe7,
	0xab, 0x7f, 0xee, 0x5f, 0x5c, 0xfa, 0xf3, 0xa9, 0x47, 0xbf, 0x79, 0xb4, 0xdb, 0xc0, 0x2e, 0xb4,
	0x2f, 0x2f, 0xe8, 0xe4, 0xb4, 0xaa, 0x90, 0xac, 0x72, 0xfc, 0xc9, 0xf3, 0x67, 0x55, 0xc5, 0x1a,
	0xfd, 0x22, 0x70, 0x30, 0xc9, 0x52, 0x8c, 0x67, 0x70, 0x58, 0xa4, 0x00, 0x9f, 0xd6, 0x8d, 0x70,
	0x13, 0x4a, 0x7b, 0xa7, 0xac, 0x9c, 0x06, 0x9e, 0x43, 0xb3, 0xfc, 0x43, 0xf8, 0xac, 0xce, 0x7b,
	0x9b, 0x10, 0x7b, 0xb7, 0x9e, 0xc1, 0xc6, 0xb0, 0x9f, 0x6d, 0x18, 0x9f, 0xfc, 0x7d, 0xf7, 0x6b,
	0x7b, 0x87, 0xa8, 0x9c, 0xc6, 0xf8, 0xfd, 0xf7, 0xb7, 0x0f, 0xb9, 0x88, 0x1f, 0xef, 0x43, 0x82,
	0xc3, 0x5c, 0x7a, 0xf3, 0x3b, 0x00, 0x00, 0xff, 0xff, 0xed, 0x38, 0x89, 0x94, 0x10, 0x04, 0x00,
	0x00,
}