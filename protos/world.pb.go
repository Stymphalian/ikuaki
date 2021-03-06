// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/Stymphalian/ikuaki/protos/world.proto

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

type EnterReq struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *EnterReq) Reset()                    { *m = EnterReq{} }
func (m *EnterReq) String() string            { return proto.CompactTextString(m) }
func (*EnterReq) ProtoMessage()               {}
func (*EnterReq) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *EnterReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type EnterRes struct {
}

func (m *EnterRes) Reset()                    { *m = EnterRes{} }
func (m *EnterRes) String() string            { return proto.CompactTextString(m) }
func (*EnterRes) ProtoMessage()               {}
func (*EnterRes) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

type ExitReq struct {
}

func (m *ExitReq) Reset()                    { *m = ExitReq{} }
func (m *ExitReq) String() string            { return proto.CompactTextString(m) }
func (*ExitReq) ProtoMessage()               {}
func (*ExitReq) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

type ExitRes struct {
}

func (m *ExitRes) Reset()                    { *m = ExitRes{} }
func (m *ExitRes) String() string            { return proto.CompactTextString(m) }
func (*ExitRes) ProtoMessage()               {}
func (*ExitRes) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{3} }

type InformReq struct {
	AgentName string `protobuf:"bytes,1,opt,name=AgentName" json:"AgentName,omitempty"`
	Text      string `protobuf:"bytes,2,opt,name=Text" json:"Text,omitempty"`
}

func (m *InformReq) Reset()                    { *m = InformReq{} }
func (m *InformReq) String() string            { return proto.CompactTextString(m) }
func (*InformReq) ProtoMessage()               {}
func (*InformReq) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{4} }

func (m *InformReq) GetAgentName() string {
	if m != nil {
		return m.AgentName
	}
	return ""
}

func (m *InformReq) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

type InformRes struct {
	Text string `protobuf:"bytes,1,opt,name=Text" json:"Text,omitempty"`
}

func (m *InformRes) Reset()                    { *m = InformRes{} }
func (m *InformRes) String() string            { return proto.CompactTextString(m) }
func (*InformRes) ProtoMessage()               {}
func (*InformRes) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{5} }

func (m *InformRes) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func init() {
	proto.RegisterType((*EnterReq)(nil), "Stymphalian.ikuaki.EnterReq")
	proto.RegisterType((*EnterRes)(nil), "Stymphalian.ikuaki.EnterRes")
	proto.RegisterType((*ExitReq)(nil), "Stymphalian.ikuaki.ExitReq")
	proto.RegisterType((*ExitRes)(nil), "Stymphalian.ikuaki.ExitRes")
	proto.RegisterType((*InformReq)(nil), "Stymphalian.ikuaki.InformReq")
	proto.RegisterType((*InformRes)(nil), "Stymphalian.ikuaki.InformRes")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for World service

type WorldClient interface {
	Enter(ctx context.Context, in *EnterReq, opts ...grpc.CallOption) (*EnterRes, error)
	Exit(ctx context.Context, in *ExitReq, opts ...grpc.CallOption) (*ExitRes, error)
	Inform(ctx context.Context, opts ...grpc.CallOption) (World_InformClient, error)
}

type worldClient struct {
	cc *grpc.ClientConn
}

func NewWorldClient(cc *grpc.ClientConn) WorldClient {
	return &worldClient{cc}
}

func (c *worldClient) Enter(ctx context.Context, in *EnterReq, opts ...grpc.CallOption) (*EnterRes, error) {
	out := new(EnterRes)
	err := grpc.Invoke(ctx, "/Stymphalian.ikuaki.World/Enter", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *worldClient) Exit(ctx context.Context, in *ExitReq, opts ...grpc.CallOption) (*ExitRes, error) {
	out := new(ExitRes)
	err := grpc.Invoke(ctx, "/Stymphalian.ikuaki.World/Exit", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *worldClient) Inform(ctx context.Context, opts ...grpc.CallOption) (World_InformClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_World_serviceDesc.Streams[0], c.cc, "/Stymphalian.ikuaki.World/Inform", opts...)
	if err != nil {
		return nil, err
	}
	x := &worldInformClient{stream}
	return x, nil
}

type World_InformClient interface {
	Send(*InformReq) error
	Recv() (*InformRes, error)
	grpc.ClientStream
}

type worldInformClient struct {
	grpc.ClientStream
}

func (x *worldInformClient) Send(m *InformReq) error {
	return x.ClientStream.SendMsg(m)
}

func (x *worldInformClient) Recv() (*InformRes, error) {
	m := new(InformRes)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for World service

type WorldServer interface {
	Enter(context.Context, *EnterReq) (*EnterRes, error)
	Exit(context.Context, *ExitReq) (*ExitRes, error)
	Inform(World_InformServer) error
}

func RegisterWorldServer(s *grpc.Server, srv WorldServer) {
	s.RegisterService(&_World_serviceDesc, srv)
}

func _World_Enter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EnterReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorldServer).Enter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Stymphalian.ikuaki.World/Enter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorldServer).Enter(ctx, req.(*EnterReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _World_Exit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExitReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorldServer).Exit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Stymphalian.ikuaki.World/Exit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorldServer).Exit(ctx, req.(*ExitReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _World_Inform_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(WorldServer).Inform(&worldInformServer{stream})
}

type World_InformServer interface {
	Send(*InformRes) error
	Recv() (*InformReq, error)
	grpc.ServerStream
}

type worldInformServer struct {
	grpc.ServerStream
}

func (x *worldInformServer) Send(m *InformRes) error {
	return x.ServerStream.SendMsg(m)
}

func (x *worldInformServer) Recv() (*InformReq, error) {
	m := new(InformReq)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _World_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Stymphalian.ikuaki.World",
	HandlerType: (*WorldServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Enter",
			Handler:    _World_Enter_Handler,
		},
		{
			MethodName: "Exit",
			Handler:    _World_Exit_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Inform",
			Handler:       _World_Inform_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "github.com/Stymphalian/ikuaki/protos/world.proto",
}

func init() { proto.RegisterFile("github.com/Stymphalian/ikuaki/protos/world.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 269 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x48, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x0f, 0x2e, 0xa9, 0xcc, 0x2d, 0xc8, 0x48, 0xcc, 0xc9,
	0x4c, 0xcc, 0xd3, 0xcf, 0xcc, 0x2e, 0x4d, 0xcc, 0xce, 0xd4, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0x2f,
	0xd6, 0x2f, 0xcf, 0x2f, 0xca, 0x49, 0xd1, 0x03, 0x73, 0x84, 0x84, 0x90, 0x94, 0xe9, 0x41, 0x94,
	0x49, 0x19, 0x12, 0x65, 0x4a, 0x72, 0x7e, 0x6e, 0x6e, 0x7e, 0x1e, 0xc4, 0x18, 0x25, 0x39, 0x2e,
	0x0e, 0xd7, 0xbc, 0x92, 0xd4, 0xa2, 0xa0, 0xd4, 0x42, 0x21, 0x21, 0x2e, 0x96, 0xbc, 0xc4, 0xdc,
	0x54, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x30, 0x5b, 0x89, 0x0b, 0x2e, 0x5f, 0xac, 0xc4,
	0xc9, 0xc5, 0xee, 0x5a, 0x91, 0x59, 0x12, 0x94, 0x5a, 0x88, 0x60, 0x16, 0x2b, 0xd9, 0x72, 0x71,
	0x7a, 0xe6, 0xa5, 0xe5, 0x17, 0xe5, 0x82, 0x8c, 0x90, 0xe1, 0xe2, 0x74, 0x4c, 0x4f, 0xcd, 0x2b,
	0xf1, 0x43, 0x98, 0x83, 0x10, 0x00, 0x59, 0x10, 0x92, 0x5a, 0x51, 0x22, 0xc1, 0x04, 0xb1, 0x00,
	0xc4, 0x56, 0x92, 0x47, 0x68, 0x2f, 0x86, 0x2b, 0x60, 0x44, 0x28, 0x30, 0x7a, 0xc0, 0xc8, 0xc5,
	0x1a, 0x0e, 0xf2, 0xb8, 0x90, 0x2b, 0x17, 0x2b, 0xd8, 0x2d, 0x42, 0x32, 0x7a, 0x98, 0x9e, 0xd7,
	0x83, 0x79, 0x43, 0x0a, 0x9f, 0x6c, 0xb1, 0x12, 0x83, 0x90, 0x13, 0x17, 0x0b, 0xc8, 0xed, 0x42,
	0xd2, 0x58, 0xd5, 0x41, 0x3c, 0x28, 0x85, 0x47, 0x12, 0x64, 0x86, 0x0f, 0x17, 0x1b, 0xc4, 0xd5,
	0x42, 0xb2, 0xd8, 0x14, 0xc2, 0x03, 0x44, 0x0a, 0xaf, 0x74, 0xb1, 0x12, 0x83, 0x06, 0xa3, 0x01,
	0xa3, 0x93, 0x65, 0x94, 0x39, 0x31, 0x31, 0x67, 0x8d, 0x69, 0x54, 0x12, 0x1b, 0x58, 0xca, 0x18,
	0x10, 0x00, 0x00, 0xff, 0xff, 0x1b, 0x42, 0xc3, 0xec, 0x41, 0x02, 0x00, 0x00,
}
