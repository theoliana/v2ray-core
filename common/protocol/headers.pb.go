// Code generated by protoc-gen-go.
// source: v2ray.com/core/common/protocol/headers.proto
// DO NOT EDIT!

/*
Package protocol is a generated protocol buffer package.

It is generated from these files:
	v2ray.com/core/common/protocol/headers.proto
	v2ray.com/core/common/protocol/server_spec.proto
	v2ray.com/core/common/protocol/user.proto

It has these top-level messages:
	SecurityConfig
	ServerEndpoint
	User
*/
package protocol

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

type SecurityType int32

const (
	SecurityType_NONE              SecurityType = 0
	SecurityType_LEGACY            SecurityType = 1
	SecurityType_AUTO              SecurityType = 2
	SecurityType_AES128_GCM        SecurityType = 3
	SecurityType_CHACHA20_POLY1305 SecurityType = 4
)

var SecurityType_name = map[int32]string{
	0: "NONE",
	1: "LEGACY",
	2: "AUTO",
	3: "AES128_GCM",
	4: "CHACHA20_POLY1305",
}
var SecurityType_value = map[string]int32{
	"NONE":              0,
	"LEGACY":            1,
	"AUTO":              2,
	"AES128_GCM":        3,
	"CHACHA20_POLY1305": 4,
}

func (x SecurityType) String() string {
	return proto.EnumName(SecurityType_name, int32(x))
}
func (SecurityType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type SecurityConfig struct {
	Type SecurityType `protobuf:"varint,1,opt,name=type,enum=v2ray.core.common.protocol.SecurityType" json:"type,omitempty"`
}

func (m *SecurityConfig) Reset()                    { *m = SecurityConfig{} }
func (m *SecurityConfig) String() string            { return proto.CompactTextString(m) }
func (*SecurityConfig) ProtoMessage()               {}
func (*SecurityConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func init() {
	proto.RegisterType((*SecurityConfig)(nil), "v2ray.core.common.protocol.SecurityConfig")
	proto.RegisterEnum("v2ray.core.common.protocol.SecurityType", SecurityType_name, SecurityType_value)
}

func init() { proto.RegisterFile("v2ray.com/core/common/protocol/headers.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 227 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xd2, 0x29, 0x33, 0x2a, 0x4a,
	0xac, 0xd4, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xce, 0x2f, 0x4a, 0xd5, 0x4f, 0xce, 0xcf, 0xcd, 0xcd,
	0xcf, 0xd3, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0x4f, 0xce, 0xcf, 0xd1, 0xcf, 0x48, 0x4d, 0x4c, 0x49,
	0x2d, 0x2a, 0xd6, 0x03, 0x0b, 0x08, 0x49, 0xc1, 0x54, 0x17, 0xa5, 0xea, 0x41, 0x54, 0xea, 0xc1,
	0x54, 0x2a, 0xf9, 0x71, 0xf1, 0x05, 0xa7, 0x26, 0x97, 0x16, 0x65, 0x96, 0x54, 0x3a, 0xe7, 0xe7,
	0xa5, 0x65, 0xa6, 0x0b, 0xd9, 0x70, 0xb1, 0x94, 0x54, 0x16, 0xa4, 0x4a, 0x30, 0x2a, 0x30, 0x6a,
	0xf0, 0x19, 0x69, 0xe8, 0xe1, 0xd6, 0xac, 0x07, 0xd3, 0x19, 0x52, 0x59, 0x90, 0x1a, 0x04, 0xd6,
	0xa5, 0x15, 0xca, 0xc5, 0x83, 0x2c, 0x2a, 0xc4, 0xc1, 0xc5, 0xe2, 0xe7, 0xef, 0xe7, 0x2a, 0xc0,
	0x20, 0xc4, 0xc5, 0xc5, 0xe6, 0xe3, 0xea, 0xee, 0xe8, 0x1c, 0x29, 0xc0, 0x08, 0x12, 0x75, 0x0c,
	0x0d, 0xf1, 0x17, 0x60, 0x12, 0xe2, 0xe3, 0xe2, 0x72, 0x74, 0x0d, 0x36, 0x34, 0xb2, 0x88, 0x77,
	0x77, 0xf6, 0x15, 0x60, 0x16, 0x12, 0xe5, 0x12, 0x74, 0xf6, 0x70, 0x74, 0xf6, 0x70, 0x34, 0x32,
	0x88, 0x0f, 0xf0, 0xf7, 0x89, 0x34, 0x34, 0x36, 0x30, 0x15, 0x60, 0x71, 0xb2, 0xe0, 0x92, 0x4b,
	0xce, 0xcf, 0xc5, 0xe3, 0x16, 0x27, 0x1e, 0x0f, 0x88, 0x9f, 0x03, 0x40, 0x02, 0x51, 0x1c, 0x30,
	0xf1, 0x24, 0x36, 0x30, 0xcb, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xc4, 0x1e, 0x65, 0x93, 0x33,
	0x01, 0x00, 0x00,
}
