// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/health_checker/redis/v2/redis.proto

package envoy_config_health_checker_redis_v2

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Redis struct {
	// If set, optionally perform ``EXISTS <key>`` instead of ``PING``. A return value
	// from Redis of 0 (does not exist) is considered a passing healthcheck. A return value other
	// than 0 is considered a failure. This allows the user to mark a Redis instance for maintenance
	// by setting the specified key to any value and waiting for traffic to drain.
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Redis) Reset()         { *m = Redis{} }
func (m *Redis) String() string { return proto.CompactTextString(m) }
func (*Redis) ProtoMessage()    {}
func (*Redis) Descriptor() ([]byte, []int) {
	return fileDescriptor_055a998fcb839d64, []int{0}
}

func (m *Redis) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Redis.Unmarshal(m, b)
}
func (m *Redis) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Redis.Marshal(b, m, deterministic)
}
func (m *Redis) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Redis.Merge(m, src)
}
func (m *Redis) XXX_Size() int {
	return xxx_messageInfo_Redis.Size(m)
}
func (m *Redis) XXX_DiscardUnknown() {
	xxx_messageInfo_Redis.DiscardUnknown(m)
}

var xxx_messageInfo_Redis proto.InternalMessageInfo

func (m *Redis) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func init() {
	proto.RegisterType((*Redis)(nil), "envoy.config.health_checker.redis.v2.Redis")
}

func init() {
	proto.RegisterFile("envoy/config/health_checker/redis/v2/redis.proto", fileDescriptor_055a998fcb839d64)
}

var fileDescriptor_055a998fcb839d64 = []byte{
	// 140 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x48, 0xcd, 0x2b, 0xcb,
	0xaf, 0xd4, 0x4f, 0xce, 0xcf, 0x4b, 0xcb, 0x4c, 0xd7, 0xcf, 0x48, 0x4d, 0xcc, 0x29, 0xc9, 0x88,
	0x4f, 0xce, 0x48, 0x4d, 0xce, 0x4e, 0x2d, 0xd2, 0x2f, 0x4a, 0x4d, 0xc9, 0x2c, 0xd6, 0x2f, 0x33,
	0x82, 0x30, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x54, 0xc0, 0x3a, 0xf4, 0x20, 0x3a, 0xf4,
	0x50, 0x75, 0xe8, 0x41, 0x14, 0x96, 0x19, 0x29, 0x49, 0x72, 0xb1, 0x06, 0x81, 0xd8, 0x42, 0x02,
	0x5c, 0xcc, 0xd9, 0xa9, 0x95, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x20, 0xa6, 0x93, 0x13,
	0x97, 0x51, 0x66, 0xbe, 0x1e, 0xd8, 0x94, 0x82, 0xa2, 0xfc, 0x8a, 0x4a, 0x3d, 0x62, 0x0c, 0x74,
	0xe2, 0x02, 0x1b, 0x17, 0x00, 0x72, 0x42, 0x00, 0x63, 0x12, 0x1b, 0xd8, 0x2d, 0xc6, 0x80, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xff, 0x59, 0x26, 0xf5, 0xbf, 0x00, 0x00, 0x00,
}
