// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/Stymphalian/ikuaki/protos/common.proto

/*
Package Stymphalian_ikuaki is a generated protocol buffer package.

It is generated from these files:
	github.com/Stymphalian/ikuaki/protos/common.proto
	github.com/Stymphalian/ikuaki/protos/lobby.proto
	github.com/Stymphalian/ikuaki/protos/world.proto
	github.com/Stymphalian/ikuaki/protos/agent.proto

It has these top-level messages:
	AgentId
	Addr
	CreateReq
	CreateRes
	DestroyReq
	DestroyRes
	ListReq
	ListRes
	EnterReq
	EnterRes
	ExitReq
	ExitRes
	InformReq
	InformRes
	UpdateRes
	UpdateReq
*/
package Stymphalian_ikuaki

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type AgentId struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *AgentId) Reset()                    { *m = AgentId{} }
func (m *AgentId) String() string            { return proto.CompactTextString(m) }
func (*AgentId) ProtoMessage()               {}
func (*AgentId) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *AgentId) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type Addr struct {
	Hostport string `protobuf:"bytes,1,opt,name=hostport" json:"hostport,omitempty"`
}

func (m *Addr) Reset()                    { *m = Addr{} }
func (m *Addr) String() string            { return proto.CompactTextString(m) }
func (*Addr) ProtoMessage()               {}
func (*Addr) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Addr) GetHostport() string {
	if m != nil {
		return m.Hostport
	}
	return ""
}

func init() {
	proto.RegisterType((*AgentId)(nil), "Stymphalian.ikuaki.AgentId")
	proto.RegisterType((*Addr)(nil), "Stymphalian.ikuaki.Addr")
}

func init() { proto.RegisterFile("github.com/Stymphalian/ikuaki/protos/common.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 144 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x4c, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x0f, 0x2e, 0xa9, 0xcc, 0x2d, 0xc8, 0x48, 0xcc, 0xc9,
	0x4c, 0xcc, 0xd3, 0xcf, 0xcc, 0x2e, 0x4d, 0xcc, 0xce, 0xd4, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0x2f,
	0xd6, 0x4f, 0xce, 0xcf, 0xcd, 0xcd, 0xcf, 0xd3, 0x03, 0xf3, 0x84, 0x84, 0x90, 0xd4, 0xe9, 0x41,
	0xd4, 0x29, 0x49, 0x72, 0xb1, 0x3b, 0xa6, 0xa7, 0xe6, 0x95, 0x78, 0xa6, 0x08, 0xf1, 0x71, 0x31,
	0x65, 0xa6, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x31, 0x65, 0xa6, 0x28, 0x29, 0x71, 0xb1,
	0x38, 0xa6, 0xa4, 0x14, 0x09, 0x49, 0x71, 0x71, 0x64, 0xe4, 0x17, 0x97, 0x14, 0xe4, 0x17, 0x95,
	0x40, 0x65, 0xe1, 0x7c, 0x27, 0xcb, 0x28, 0x73, 0x62, 0xdc, 0x61, 0x8d, 0x69, 0x73, 0x12, 0x1b,
	0x58, 0xca, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x29, 0x2d, 0x43, 0x6c, 0xc9, 0x00, 0x00, 0x00,
}