// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/common/tap/v2alpha/common.proto

package envoy_config_common_tap_v2alpha

import (
	fmt "fmt"
	core "github.com/cilium/proxy/go/envoy/api/v2/core"
	v2alpha "github.com/cilium/proxy/go/envoy/service/tap/v2alpha"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

// Common configuration for all tap extensions.
type CommonExtensionConfig struct {
	// Types that are valid to be assigned to ConfigType:
	//	*CommonExtensionConfig_AdminConfig
	//	*CommonExtensionConfig_StaticConfig
	//	*CommonExtensionConfig_TapdsConfig
	ConfigType           isCommonExtensionConfig_ConfigType `protobuf_oneof:"config_type"`
	XXX_NoUnkeyedLiteral struct{}                           `json:"-"`
	XXX_unrecognized     []byte                             `json:"-"`
	XXX_sizecache        int32                              `json:"-"`
}

func (m *CommonExtensionConfig) Reset()         { *m = CommonExtensionConfig{} }
func (m *CommonExtensionConfig) String() string { return proto.CompactTextString(m) }
func (*CommonExtensionConfig) ProtoMessage()    {}
func (*CommonExtensionConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_79cf139d98a2fe3f, []int{0}
}

func (m *CommonExtensionConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommonExtensionConfig.Unmarshal(m, b)
}
func (m *CommonExtensionConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommonExtensionConfig.Marshal(b, m, deterministic)
}
func (m *CommonExtensionConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommonExtensionConfig.Merge(m, src)
}
func (m *CommonExtensionConfig) XXX_Size() int {
	return xxx_messageInfo_CommonExtensionConfig.Size(m)
}
func (m *CommonExtensionConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_CommonExtensionConfig.DiscardUnknown(m)
}

var xxx_messageInfo_CommonExtensionConfig proto.InternalMessageInfo

type isCommonExtensionConfig_ConfigType interface {
	isCommonExtensionConfig_ConfigType()
}

type CommonExtensionConfig_AdminConfig struct {
	AdminConfig *AdminConfig `protobuf:"bytes,1,opt,name=admin_config,json=adminConfig,proto3,oneof"`
}

type CommonExtensionConfig_StaticConfig struct {
	StaticConfig *v2alpha.TapConfig `protobuf:"bytes,2,opt,name=static_config,json=staticConfig,proto3,oneof"`
}

type CommonExtensionConfig_TapdsConfig struct {
	TapdsConfig *CommonExtensionConfig_TapDSConfig `protobuf:"bytes,3,opt,name=tapds_config,json=tapdsConfig,proto3,oneof"`
}

func (*CommonExtensionConfig_AdminConfig) isCommonExtensionConfig_ConfigType() {}

func (*CommonExtensionConfig_StaticConfig) isCommonExtensionConfig_ConfigType() {}

func (*CommonExtensionConfig_TapdsConfig) isCommonExtensionConfig_ConfigType() {}

func (m *CommonExtensionConfig) GetConfigType() isCommonExtensionConfig_ConfigType {
	if m != nil {
		return m.ConfigType
	}
	return nil
}

func (m *CommonExtensionConfig) GetAdminConfig() *AdminConfig {
	if x, ok := m.GetConfigType().(*CommonExtensionConfig_AdminConfig); ok {
		return x.AdminConfig
	}
	return nil
}

func (m *CommonExtensionConfig) GetStaticConfig() *v2alpha.TapConfig {
	if x, ok := m.GetConfigType().(*CommonExtensionConfig_StaticConfig); ok {
		return x.StaticConfig
	}
	return nil
}

func (m *CommonExtensionConfig) GetTapdsConfig() *CommonExtensionConfig_TapDSConfig {
	if x, ok := m.GetConfigType().(*CommonExtensionConfig_TapdsConfig); ok {
		return x.TapdsConfig
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*CommonExtensionConfig) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*CommonExtensionConfig_AdminConfig)(nil),
		(*CommonExtensionConfig_StaticConfig)(nil),
		(*CommonExtensionConfig_TapdsConfig)(nil),
	}
}

// [#not-implemented-hide:]
type CommonExtensionConfig_TapDSConfig struct {
	// Configuration for the source of TapDS updates for this Cluster.
	ConfigSource *core.ConfigSource `protobuf:"bytes,1,opt,name=config_source,json=configSource,proto3" json:"config_source,omitempty"`
	// Tap config to request from XDS server.
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CommonExtensionConfig_TapDSConfig) Reset()         { *m = CommonExtensionConfig_TapDSConfig{} }
func (m *CommonExtensionConfig_TapDSConfig) String() string { return proto.CompactTextString(m) }
func (*CommonExtensionConfig_TapDSConfig) ProtoMessage()    {}
func (*CommonExtensionConfig_TapDSConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_79cf139d98a2fe3f, []int{0, 0}
}

func (m *CommonExtensionConfig_TapDSConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommonExtensionConfig_TapDSConfig.Unmarshal(m, b)
}
func (m *CommonExtensionConfig_TapDSConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommonExtensionConfig_TapDSConfig.Marshal(b, m, deterministic)
}
func (m *CommonExtensionConfig_TapDSConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommonExtensionConfig_TapDSConfig.Merge(m, src)
}
func (m *CommonExtensionConfig_TapDSConfig) XXX_Size() int {
	return xxx_messageInfo_CommonExtensionConfig_TapDSConfig.Size(m)
}
func (m *CommonExtensionConfig_TapDSConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_CommonExtensionConfig_TapDSConfig.DiscardUnknown(m)
}

var xxx_messageInfo_CommonExtensionConfig_TapDSConfig proto.InternalMessageInfo

func (m *CommonExtensionConfig_TapDSConfig) GetConfigSource() *core.ConfigSource {
	if m != nil {
		return m.ConfigSource
	}
	return nil
}

func (m *CommonExtensionConfig_TapDSConfig) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// Configuration for the admin handler. See :ref:`here <config_http_filters_tap_admin_handler>` for
// more information.
type AdminConfig struct {
	// Opaque configuration ID. When requests are made to the admin handler, the passed opaque ID is
	// matched to the configured filter opaque ID to determine which filter to configure.
	ConfigId             string   `protobuf:"bytes,1,opt,name=config_id,json=configId,proto3" json:"config_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AdminConfig) Reset()         { *m = AdminConfig{} }
func (m *AdminConfig) String() string { return proto.CompactTextString(m) }
func (*AdminConfig) ProtoMessage()    {}
func (*AdminConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_79cf139d98a2fe3f, []int{1}
}

func (m *AdminConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdminConfig.Unmarshal(m, b)
}
func (m *AdminConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdminConfig.Marshal(b, m, deterministic)
}
func (m *AdminConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdminConfig.Merge(m, src)
}
func (m *AdminConfig) XXX_Size() int {
	return xxx_messageInfo_AdminConfig.Size(m)
}
func (m *AdminConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_AdminConfig.DiscardUnknown(m)
}

var xxx_messageInfo_AdminConfig proto.InternalMessageInfo

func (m *AdminConfig) GetConfigId() string {
	if m != nil {
		return m.ConfigId
	}
	return ""
}

func init() {
	proto.RegisterType((*CommonExtensionConfig)(nil), "envoy.config.common.tap.v2alpha.CommonExtensionConfig")
	proto.RegisterType((*CommonExtensionConfig_TapDSConfig)(nil), "envoy.config.common.tap.v2alpha.CommonExtensionConfig.TapDSConfig")
	proto.RegisterType((*AdminConfig)(nil), "envoy.config.common.tap.v2alpha.AdminConfig")
}

func init() {
	proto.RegisterFile("envoy/config/common/tap/v2alpha/common.proto", fileDescriptor_79cf139d98a2fe3f)
}

var fileDescriptor_79cf139d98a2fe3f = []byte{
	// 378 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0x51, 0x4b, 0xeb, 0x30,
	0x14, 0xc7, 0x6f, 0xb7, 0xdd, 0x7b, 0xb7, 0xb4, 0x83, 0x4b, 0xe0, 0xa2, 0xd4, 0x87, 0x8d, 0x31,
	0xc5, 0x87, 0x99, 0x42, 0xf7, 0x2e, 0x98, 0x29, 0x28, 0x82, 0xcc, 0xcd, 0xf7, 0x71, 0x6c, 0xe3,
	0x0c, 0xac, 0x4d, 0x68, 0x63, 0xd9, 0xfc, 0x08, 0x7e, 0x1a, 0x3f, 0x9e, 0xec, 0x49, 0x9a, 0x64,
	0xae, 0x03, 0x65, 0x6f, 0xed, 0x39, 0xff, 0xf3, 0xfb, 0x9f, 0xfc, 0x13, 0x34, 0x60, 0x69, 0x21,
	0x56, 0x41, 0x24, 0xd2, 0x27, 0x3e, 0x0f, 0x22, 0x91, 0x24, 0x22, 0x0d, 0x14, 0xc8, 0xa0, 0x08,
	0x61, 0x21, 0x9f, 0xc1, 0x96, 0x88, 0xcc, 0x84, 0x12, 0xb8, 0xa3, 0xd5, 0xc4, 0xa8, 0x89, 0x6d,
	0x29, 0x90, 0xc4, 0xaa, 0xfd, 0x63, 0x83, 0x03, 0xc9, 0x83, 0x22, 0x0c, 0x22, 0x91, 0x31, 0x8b,
	0x9e, 0xe5, 0xe2, 0x25, 0x8b, 0x98, 0xe1, 0xf8, 0x27, 0x46, 0x96, 0xb3, 0xac, 0xe0, 0x11, 0xfb,
	0xd1, 0xcf, 0x3f, 0x28, 0x60, 0xc1, 0x63, 0x50, 0x2c, 0xd8, 0x7c, 0x98, 0x46, 0xef, 0xbd, 0x8e,
	0xfe, 0x8f, 0xb4, 0xf2, 0x6a, 0xa9, 0x58, 0x9a, 0x73, 0x91, 0x8e, 0xb4, 0x0f, 0xbe, 0x47, 0x1e,
	0xc4, 0x09, 0x4f, 0x67, 0xc6, 0xf7, 0xd0, 0xe9, 0x3a, 0xa7, 0x6e, 0x38, 0x20, 0x7b, 0x36, 0x27,
	0x17, 0xe5, 0x90, 0x61, 0x5c, 0xff, 0x9a, 0xb8, 0xb0, 0xfd, 0xc5, 0xb7, 0xa8, 0x9d, 0x2b, 0x50,
	0x3c, 0xda, 0x30, 0x6b, 0x9a, 0xd9, 0xb7, 0x4c, 0x7b, 0x8a, 0x1d, 0xda, 0x03, 0xc8, 0x2f, 0x96,
	0x67, 0x86, 0x2d, 0x6c, 0x8e, 0x3c, 0x05, 0x32, 0xce, 0x37, 0xac, 0xba, 0x66, 0xd1, 0xbd, 0xfb,
	0x7d, 0x7b, 0xda, 0xd2, 0xe7, 0x72, 0xba, 0xdd, 0x5a, 0x93, 0xcd, 0xaf, 0xff, 0x8a, 0xdc, 0x4a,
	0x17, 0xdf, 0xa1, 0xf6, 0xce, 0x4d, 0xd8, 0x60, 0x3a, 0xd6, 0x18, 0x24, 0x27, 0x45, 0x48, 0xca,
	0x1b, 0x23, 0x66, 0x62, 0xaa, 0x65, 0xb4, 0xb9, 0xa6, 0xbf, 0xdf, 0x9c, 0xda, 0x3f, 0x67, 0xe2,
	0x45, 0x95, 0x3a, 0x3e, 0x42, 0x8d, 0x14, 0x12, 0xa6, 0xb3, 0x68, 0xd1, 0xbf, 0x6b, 0xda, 0xc8,
	0x6a, 0x5d, 0x67, 0xa2, 0x8b, 0x14, 0x23, 0xd7, 0x9a, 0xa9, 0x95, 0x64, 0xb8, 0xfe, 0x41, 0x9d,
	0xde, 0x10, 0xb9, 0x95, 0x8c, 0x71, 0x1f, 0xb5, 0xac, 0x84, 0xc7, 0x7a, 0x97, 0x0a, 0xa4, 0x69,
	0x3a, 0x37, 0x31, 0x3d, 0x47, 0x67, 0x5c, 0x98, 0x15, 0x65, 0x26, 0x96, 0xab, 0x7d, 0x31, 0x51,
	0xd7, 0xe4, 0x34, 0x2e, 0x5f, 0xc9, 0xd8, 0x79, 0xfc, 0xa3, 0x9f, 0xcb, 0xf0, 0x33, 0x00, 0x00,
	0xff, 0xff, 0x93, 0xc8, 0x8a, 0x58, 0xe7, 0x02, 0x00, 0x00,
}
