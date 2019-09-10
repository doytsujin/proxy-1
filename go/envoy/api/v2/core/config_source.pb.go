// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/api/v2/core/config_source.proto

package envoy_api_v2_core

import (
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

// APIs may be fetched via either REST or gRPC.
type ApiConfigSource_ApiType int32

const (
	// Ideally this would be 'reserved 0' but one can't reserve the default
	// value. Instead we throw an exception if this is ever used.
	ApiConfigSource_UNSUPPORTED_REST_LEGACY ApiConfigSource_ApiType = 0 // Deprecated: Do not use.
	// REST-JSON v2 API. The `canonical JSON encoding
	// <https://developers.google.com/protocol-buffers/docs/proto3#json>`_ for
	// the v2 protos is used.
	ApiConfigSource_REST ApiConfigSource_ApiType = 1
	// gRPC v2 API.
	ApiConfigSource_GRPC ApiConfigSource_ApiType = 2
	// Using the delta xDS gRPC service, i.e. DeltaDiscovery{Request,Response}
	// rather than Discovery{Request,Response}. Rather than sending Envoy the entire state
	// with every update, the xDS server only sends what has changed since the last update.
	//
	// DELTA_GRPC is not yet entirely implemented! Initially, only CDS is available.
	// Do not use for other xDSes. TODO(fredlas) update/remove this warning when appropriate.
	ApiConfigSource_DELTA_GRPC ApiConfigSource_ApiType = 3
)

var ApiConfigSource_ApiType_name = map[int32]string{
	0: "UNSUPPORTED_REST_LEGACY",
	1: "REST",
	2: "GRPC",
	3: "DELTA_GRPC",
}

var ApiConfigSource_ApiType_value = map[string]int32{
	"UNSUPPORTED_REST_LEGACY": 0,
	"REST":                    1,
	"GRPC":                    2,
	"DELTA_GRPC":              3,
}

func (x ApiConfigSource_ApiType) String() string {
	return proto.EnumName(ApiConfigSource_ApiType_name, int32(x))
}

func (ApiConfigSource_ApiType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1ffcc55cf4c30535, []int{0, 0}
}

// API configuration source. This identifies the API type and cluster that Envoy
// will use to fetch an xDS API.
// [#next-free-field: 8]
type ApiConfigSource struct {
	ApiType ApiConfigSource_ApiType `protobuf:"varint,1,opt,name=api_type,json=apiType,proto3,enum=envoy.api.v2.core.ApiConfigSource_ApiType" json:"api_type,omitempty"`
	// Cluster names should be used only with REST. If > 1
	// cluster is defined, clusters will be cycled through if any kind of failure
	// occurs.
	//
	// .. note::
	//
	//  The cluster with name ``cluster_name`` must be statically defined and its
	//  type must not be ``EDS``.
	ClusterNames []string `protobuf:"bytes,2,rep,name=cluster_names,json=clusterNames,proto3" json:"cluster_names,omitempty"`
	// Multiple gRPC services be provided for GRPC. If > 1 cluster is defined,
	// services will be cycled through if any kind of failure occurs.
	GrpcServices []*GrpcService `protobuf:"bytes,4,rep,name=grpc_services,json=grpcServices,proto3" json:"grpc_services,omitempty"`
	// For REST APIs, the delay between successive polls.
	RefreshDelay *duration.Duration `protobuf:"bytes,3,opt,name=refresh_delay,json=refreshDelay,proto3" json:"refresh_delay,omitempty"`
	// For REST APIs, the request timeout. If not set, a default value of 1s will be used.
	RequestTimeout *duration.Duration `protobuf:"bytes,5,opt,name=request_timeout,json=requestTimeout,proto3" json:"request_timeout,omitempty"`
	// For GRPC APIs, the rate limit settings. If present, discovery requests made by Envoy will be
	// rate limited.
	RateLimitSettings *RateLimitSettings `protobuf:"bytes,6,opt,name=rate_limit_settings,json=rateLimitSettings,proto3" json:"rate_limit_settings,omitempty"`
	// Skip the node identifier in subsequent discovery requests for streaming gRPC config types.
	SetNodeOnFirstMessageOnly bool     `protobuf:"varint,7,opt,name=set_node_on_first_message_only,json=setNodeOnFirstMessageOnly,proto3" json:"set_node_on_first_message_only,omitempty"`
	XXX_NoUnkeyedLiteral      struct{} `json:"-"`
	XXX_unrecognized          []byte   `json:"-"`
	XXX_sizecache             int32    `json:"-"`
}

func (m *ApiConfigSource) Reset()         { *m = ApiConfigSource{} }
func (m *ApiConfigSource) String() string { return proto.CompactTextString(m) }
func (*ApiConfigSource) ProtoMessage()    {}
func (*ApiConfigSource) Descriptor() ([]byte, []int) {
	return fileDescriptor_1ffcc55cf4c30535, []int{0}
}

func (m *ApiConfigSource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApiConfigSource.Unmarshal(m, b)
}
func (m *ApiConfigSource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApiConfigSource.Marshal(b, m, deterministic)
}
func (m *ApiConfigSource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApiConfigSource.Merge(m, src)
}
func (m *ApiConfigSource) XXX_Size() int {
	return xxx_messageInfo_ApiConfigSource.Size(m)
}
func (m *ApiConfigSource) XXX_DiscardUnknown() {
	xxx_messageInfo_ApiConfigSource.DiscardUnknown(m)
}

var xxx_messageInfo_ApiConfigSource proto.InternalMessageInfo

func (m *ApiConfigSource) GetApiType() ApiConfigSource_ApiType {
	if m != nil {
		return m.ApiType
	}
	return ApiConfigSource_UNSUPPORTED_REST_LEGACY
}

func (m *ApiConfigSource) GetClusterNames() []string {
	if m != nil {
		return m.ClusterNames
	}
	return nil
}

func (m *ApiConfigSource) GetGrpcServices() []*GrpcService {
	if m != nil {
		return m.GrpcServices
	}
	return nil
}

func (m *ApiConfigSource) GetRefreshDelay() *duration.Duration {
	if m != nil {
		return m.RefreshDelay
	}
	return nil
}

func (m *ApiConfigSource) GetRequestTimeout() *duration.Duration {
	if m != nil {
		return m.RequestTimeout
	}
	return nil
}

func (m *ApiConfigSource) GetRateLimitSettings() *RateLimitSettings {
	if m != nil {
		return m.RateLimitSettings
	}
	return nil
}

func (m *ApiConfigSource) GetSetNodeOnFirstMessageOnly() bool {
	if m != nil {
		return m.SetNodeOnFirstMessageOnly
	}
	return false
}

// Aggregated Discovery Service (ADS) options. This is currently empty, but when
// set in :ref:`ConfigSource <envoy_api_msg_core.ConfigSource>` can be used to
// specify that ADS is to be used.
type AggregatedConfigSource struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AggregatedConfigSource) Reset()         { *m = AggregatedConfigSource{} }
func (m *AggregatedConfigSource) String() string { return proto.CompactTextString(m) }
func (*AggregatedConfigSource) ProtoMessage()    {}
func (*AggregatedConfigSource) Descriptor() ([]byte, []int) {
	return fileDescriptor_1ffcc55cf4c30535, []int{1}
}

func (m *AggregatedConfigSource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AggregatedConfigSource.Unmarshal(m, b)
}
func (m *AggregatedConfigSource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AggregatedConfigSource.Marshal(b, m, deterministic)
}
func (m *AggregatedConfigSource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AggregatedConfigSource.Merge(m, src)
}
func (m *AggregatedConfigSource) XXX_Size() int {
	return xxx_messageInfo_AggregatedConfigSource.Size(m)
}
func (m *AggregatedConfigSource) XXX_DiscardUnknown() {
	xxx_messageInfo_AggregatedConfigSource.DiscardUnknown(m)
}

var xxx_messageInfo_AggregatedConfigSource proto.InternalMessageInfo

// [#not-implemented-hide:]
// Self-referencing config source options. This is currently empty, but when
// set in :ref:`ConfigSource <envoy_api_msg_core.ConfigSource>` can be used to
// specify that other data can be obtained from the same server.
type SelfConfigSource struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SelfConfigSource) Reset()         { *m = SelfConfigSource{} }
func (m *SelfConfigSource) String() string { return proto.CompactTextString(m) }
func (*SelfConfigSource) ProtoMessage()    {}
func (*SelfConfigSource) Descriptor() ([]byte, []int) {
	return fileDescriptor_1ffcc55cf4c30535, []int{2}
}

func (m *SelfConfigSource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SelfConfigSource.Unmarshal(m, b)
}
func (m *SelfConfigSource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SelfConfigSource.Marshal(b, m, deterministic)
}
func (m *SelfConfigSource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SelfConfigSource.Merge(m, src)
}
func (m *SelfConfigSource) XXX_Size() int {
	return xxx_messageInfo_SelfConfigSource.Size(m)
}
func (m *SelfConfigSource) XXX_DiscardUnknown() {
	xxx_messageInfo_SelfConfigSource.DiscardUnknown(m)
}

var xxx_messageInfo_SelfConfigSource proto.InternalMessageInfo

// Rate Limit settings to be applied for discovery requests made by Envoy.
type RateLimitSettings struct {
	// Maximum number of tokens to be used for rate limiting discovery request calls. If not set, a
	// default value of 100 will be used.
	MaxTokens *wrappers.UInt32Value `protobuf:"bytes,1,opt,name=max_tokens,json=maxTokens,proto3" json:"max_tokens,omitempty"`
	// Rate at which tokens will be filled per second. If not set, a default fill rate of 10 tokens
	// per second will be used.
	FillRate             *wrappers.DoubleValue `protobuf:"bytes,2,opt,name=fill_rate,json=fillRate,proto3" json:"fill_rate,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *RateLimitSettings) Reset()         { *m = RateLimitSettings{} }
func (m *RateLimitSettings) String() string { return proto.CompactTextString(m) }
func (*RateLimitSettings) ProtoMessage()    {}
func (*RateLimitSettings) Descriptor() ([]byte, []int) {
	return fileDescriptor_1ffcc55cf4c30535, []int{3}
}

func (m *RateLimitSettings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RateLimitSettings.Unmarshal(m, b)
}
func (m *RateLimitSettings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RateLimitSettings.Marshal(b, m, deterministic)
}
func (m *RateLimitSettings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RateLimitSettings.Merge(m, src)
}
func (m *RateLimitSettings) XXX_Size() int {
	return xxx_messageInfo_RateLimitSettings.Size(m)
}
func (m *RateLimitSettings) XXX_DiscardUnknown() {
	xxx_messageInfo_RateLimitSettings.DiscardUnknown(m)
}

var xxx_messageInfo_RateLimitSettings proto.InternalMessageInfo

func (m *RateLimitSettings) GetMaxTokens() *wrappers.UInt32Value {
	if m != nil {
		return m.MaxTokens
	}
	return nil
}

func (m *RateLimitSettings) GetFillRate() *wrappers.DoubleValue {
	if m != nil {
		return m.FillRate
	}
	return nil
}

// Configuration for :ref:`listeners <config_listeners>`, :ref:`clusters
// <config_cluster_manager>`, :ref:`routes
// <envoy_api_msg_RouteConfiguration>`, :ref:`endpoints
// <arch_overview_service_discovery>` etc. may either be sourced from the
// filesystem or from an xDS API source. Filesystem configs are watched with
// inotify for updates.
// [#next-free-field: 6]
type ConfigSource struct {
	// Types that are valid to be assigned to ConfigSourceSpecifier:
	//	*ConfigSource_Path
	//	*ConfigSource_ApiConfigSource
	//	*ConfigSource_Ads
	//	*ConfigSource_Self
	ConfigSourceSpecifier isConfigSource_ConfigSourceSpecifier `protobuf_oneof:"config_source_specifier"`
	// When this timeout is specified, Envoy will wait no longer than the specified time for first
	// config response on this xDS subscription during the :ref:`initialization process
	// <arch_overview_initialization>`. After reaching the timeout, Envoy will move to the next
	// initialization phase, even if the first config is not delivered yet. The timer is activated
	// when the xDS API subscription starts, and is disarmed on first config update or on error. 0
	// means no timeout - Envoy will wait indefinitely for the first xDS config (unless another
	// timeout applies). The default is 15s.
	InitialFetchTimeout  *duration.Duration `protobuf:"bytes,4,opt,name=initial_fetch_timeout,json=initialFetchTimeout,proto3" json:"initial_fetch_timeout,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *ConfigSource) Reset()         { *m = ConfigSource{} }
func (m *ConfigSource) String() string { return proto.CompactTextString(m) }
func (*ConfigSource) ProtoMessage()    {}
func (*ConfigSource) Descriptor() ([]byte, []int) {
	return fileDescriptor_1ffcc55cf4c30535, []int{4}
}

func (m *ConfigSource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigSource.Unmarshal(m, b)
}
func (m *ConfigSource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigSource.Marshal(b, m, deterministic)
}
func (m *ConfigSource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigSource.Merge(m, src)
}
func (m *ConfigSource) XXX_Size() int {
	return xxx_messageInfo_ConfigSource.Size(m)
}
func (m *ConfigSource) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigSource.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigSource proto.InternalMessageInfo

type isConfigSource_ConfigSourceSpecifier interface {
	isConfigSource_ConfigSourceSpecifier()
}

type ConfigSource_Path struct {
	Path string `protobuf:"bytes,1,opt,name=path,proto3,oneof"`
}

type ConfigSource_ApiConfigSource struct {
	ApiConfigSource *ApiConfigSource `protobuf:"bytes,2,opt,name=api_config_source,json=apiConfigSource,proto3,oneof"`
}

type ConfigSource_Ads struct {
	Ads *AggregatedConfigSource `protobuf:"bytes,3,opt,name=ads,proto3,oneof"`
}

type ConfigSource_Self struct {
	Self *SelfConfigSource `protobuf:"bytes,5,opt,name=self,proto3,oneof"`
}

func (*ConfigSource_Path) isConfigSource_ConfigSourceSpecifier() {}

func (*ConfigSource_ApiConfigSource) isConfigSource_ConfigSourceSpecifier() {}

func (*ConfigSource_Ads) isConfigSource_ConfigSourceSpecifier() {}

func (*ConfigSource_Self) isConfigSource_ConfigSourceSpecifier() {}

func (m *ConfigSource) GetConfigSourceSpecifier() isConfigSource_ConfigSourceSpecifier {
	if m != nil {
		return m.ConfigSourceSpecifier
	}
	return nil
}

func (m *ConfigSource) GetPath() string {
	if x, ok := m.GetConfigSourceSpecifier().(*ConfigSource_Path); ok {
		return x.Path
	}
	return ""
}

func (m *ConfigSource) GetApiConfigSource() *ApiConfigSource {
	if x, ok := m.GetConfigSourceSpecifier().(*ConfigSource_ApiConfigSource); ok {
		return x.ApiConfigSource
	}
	return nil
}

func (m *ConfigSource) GetAds() *AggregatedConfigSource {
	if x, ok := m.GetConfigSourceSpecifier().(*ConfigSource_Ads); ok {
		return x.Ads
	}
	return nil
}

func (m *ConfigSource) GetSelf() *SelfConfigSource {
	if x, ok := m.GetConfigSourceSpecifier().(*ConfigSource_Self); ok {
		return x.Self
	}
	return nil
}

func (m *ConfigSource) GetInitialFetchTimeout() *duration.Duration {
	if m != nil {
		return m.InitialFetchTimeout
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ConfigSource) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ConfigSource_Path)(nil),
		(*ConfigSource_ApiConfigSource)(nil),
		(*ConfigSource_Ads)(nil),
		(*ConfigSource_Self)(nil),
	}
}

func init() {
	proto.RegisterEnum("envoy.api.v2.core.ApiConfigSource_ApiType", ApiConfigSource_ApiType_name, ApiConfigSource_ApiType_value)
	proto.RegisterType((*ApiConfigSource)(nil), "envoy.api.v2.core.ApiConfigSource")
	proto.RegisterType((*AggregatedConfigSource)(nil), "envoy.api.v2.core.AggregatedConfigSource")
	proto.RegisterType((*SelfConfigSource)(nil), "envoy.api.v2.core.SelfConfigSource")
	proto.RegisterType((*RateLimitSettings)(nil), "envoy.api.v2.core.RateLimitSettings")
	proto.RegisterType((*ConfigSource)(nil), "envoy.api.v2.core.ConfigSource")
}

func init() {
	proto.RegisterFile("envoy/api/v2/core/config_source.proto", fileDescriptor_1ffcc55cf4c30535)
}

var fileDescriptor_1ffcc55cf4c30535 = []byte{
	// 727 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x53, 0xdd, 0x6e, 0xe2, 0x46,
	0x14, 0xc6, 0xc0, 0x6e, 0x60, 0x20, 0x04, 0x66, 0xdb, 0xc6, 0x1b, 0x55, 0x2c, 0x65, 0xb7, 0x12,
	0xdd, 0x0b, 0x23, 0xb1, 0x52, 0xa5, 0xaa, 0x6a, 0x25, 0x0c, 0x6c, 0xb2, 0x12, 0x0b, 0x68, 0x20,
	0x95, 0x7a, 0x35, 0x9a, 0x98, 0x63, 0x67, 0x54, 0xe3, 0x71, 0x67, 0xc6, 0x34, 0xdc, 0xf6, 0x35,
	0xfa, 0x06, 0x7d, 0x8d, 0x3e, 0x48, 0x5f, 0xa3, 0xda, 0xab, 0x6a, 0x8c, 0xd3, 0x86, 0x9f, 0x28,
	0xbe, 0xf2, 0x39, 0xf3, 0x7d, 0x73, 0xbe, 0xf3, 0xcd, 0x39, 0xe8, 0x6b, 0x88, 0xd6, 0x62, 0xd3,
	0x65, 0x31, 0xef, 0xae, 0x7b, 0x5d, 0x4f, 0x48, 0xe8, 0x7a, 0x22, 0xf2, 0x79, 0x40, 0x95, 0x48,
	0xa4, 0x07, 0x4e, 0x2c, 0x85, 0x16, 0xb8, 0x91, 0xc2, 0x1c, 0x16, 0x73, 0x67, 0xdd, 0x73, 0x0c,
	0xec, 0xe2, 0xcd, 0x21, 0x33, 0x90, 0xb1, 0x47, 0x15, 0xc8, 0x35, 0xbf, 0x27, 0x5e, 0x34, 0x03,
	0x21, 0x82, 0x10, 0xba, 0x69, 0x74, 0x93, 0xf8, 0xdd, 0x65, 0x22, 0x99, 0xe6, 0x22, 0x7a, 0xec,
	0xfc, 0x37, 0xc9, 0xe2, 0x18, 0xa4, 0xca, 0xce, 0xcf, 0xd7, 0x2c, 0xe4, 0x4b, 0xa6, 0xa1, 0x7b,
	0xff, 0xb3, 0x3d, 0x68, 0xff, 0x55, 0x44, 0x67, 0xfd, 0x98, 0x0f, 0x52, 0xb1, 0xf3, 0x54, 0x2b,
	0x9e, 0xa2, 0x12, 0x8b, 0x39, 0xd5, 0x9b, 0x18, 0x6c, 0xab, 0x65, 0x75, 0x6a, 0xbd, 0xb7, 0xce,
	0x81, 0x70, 0x67, 0x8f, 0x65, 0xe2, 0xc5, 0x26, 0x06, 0xb7, 0xf4, 0xc9, 0x7d, 0xf6, 0xbb, 0x95,
	0xaf, 0x5b, 0xe4, 0x84, 0x6d, 0x53, 0xf8, 0x35, 0x3a, 0xf5, 0xc2, 0x44, 0x69, 0x90, 0x34, 0x62,
	0x2b, 0x50, 0x76, 0xbe, 0x55, 0xe8, 0x94, 0x49, 0x35, 0x4b, 0x4e, 0x4c, 0x0e, 0x0f, 0xd0, 0xe9,
	0xc3, 0xc6, 0x95, 0x5d, 0x6c, 0x15, 0x3a, 0x95, 0x5e, 0xf3, 0x48, 0xe9, 0x4b, 0x19, 0x7b, 0xf3,
	0x2d, 0x8c, 0x54, 0x83, 0xff, 0x03, 0x85, 0x7f, 0x44, 0xa7, 0x12, 0x7c, 0x09, 0xea, 0x96, 0x2e,
	0x21, 0x64, 0x1b, 0xbb, 0xd0, 0xb2, 0x3a, 0x95, 0xde, 0x4b, 0x67, 0xeb, 0x8f, 0x73, 0xef, 0x8f,
	0x33, 0xcc, 0xfc, 0x23, 0xd5, 0x0c, 0x3f, 0x34, 0x70, 0x3c, 0x46, 0x67, 0x12, 0x7e, 0x4d, 0x40,
	0x69, 0xaa, 0xf9, 0x0a, 0x44, 0xa2, 0xed, 0x67, 0x4f, 0xdc, 0x90, 0x36, 0xfc, 0xa7, 0x95, 0x7f,
	0x9b, 0x23, 0xb5, 0x8c, 0xbb, 0xd8, 0x52, 0xf1, 0x02, 0xbd, 0x90, 0x4c, 0x03, 0x0d, 0xf9, 0x8a,
	0x6b, 0xaa, 0x40, 0x6b, 0x1e, 0x05, 0xca, 0x7e, 0x9e, 0xde, 0xf8, 0xe6, 0x48, 0x63, 0x84, 0x69,
	0x18, 0x1b, 0xf0, 0x3c, 0xc3, 0x92, 0x86, 0xdc, 0x4f, 0xe1, 0x3e, 0x6a, 0x2a, 0xd0, 0x34, 0x12,
	0x4b, 0xa0, 0x22, 0xa2, 0x3e, 0x97, 0x4a, 0xd3, 0x15, 0x28, 0xc5, 0x02, 0x93, 0x08, 0x37, 0xf6,
	0x49, 0xcb, 0xea, 0x94, 0xc8, 0x4b, 0x05, 0x7a, 0x22, 0x96, 0x30, 0x8d, 0xde, 0x1b, 0xc8, 0xc7,
	0x2d, 0x62, 0x1a, 0x85, 0x9b, 0xf6, 0x04, 0x9d, 0x64, 0xcf, 0x85, 0x5f, 0xa1, 0xf3, 0xeb, 0xc9,
	0xfc, 0x7a, 0x36, 0x9b, 0x92, 0xc5, 0x68, 0x48, 0xc9, 0x68, 0xbe, 0xa0, 0xe3, 0xd1, 0x65, 0x7f,
	0xf0, 0x73, 0x3d, 0x77, 0x91, 0x2f, 0x59, 0xb8, 0x84, 0x8a, 0x26, 0x59, 0x4f, 0xff, 0x2e, 0xc9,
	0x6c, 0x50, 0xcf, 0xe3, 0x1a, 0x42, 0xc3, 0xd1, 0x78, 0xd1, 0xa7, 0x69, 0x5c, 0x68, 0xdb, 0xe8,
	0x8b, 0x7e, 0x10, 0x48, 0x08, 0x98, 0x86, 0xe5, 0xc3, 0xa9, 0x68, 0x63, 0x54, 0x9f, 0x43, 0xe8,
	0xef, 0xe4, 0xfe, 0xb0, 0x50, 0xe3, 0xa0, 0x53, 0xfc, 0x3d, 0x42, 0x2b, 0x76, 0x47, 0xb5, 0xf8,
	0x05, 0x22, 0x95, 0xce, 0x5d, 0xa5, 0xf7, 0xe5, 0x81, 0xeb, 0xd7, 0x1f, 0x22, 0xfd, 0xae, 0xf7,
	0x13, 0x0b, 0x13, 0x20, 0xe5, 0x15, 0xbb, 0x5b, 0xa4, 0x70, 0xfc, 0x01, 0x95, 0x7d, 0x1e, 0x86,
	0xd4, 0xb8, 0x65, 0xe7, 0x1f, 0xe1, 0x0e, 0x45, 0x72, 0x13, 0x42, 0xca, 0x75, 0x6b, 0x9f, 0xdc,
	0x0a, 0x2e, 0x7f, 0x95, 0xcb, 0x3e, 0x52, 0x32, 0x74, 0x23, 0xaa, 0xfd, 0x77, 0x1e, 0x55, 0x77,
	0xd6, 0xe1, 0x33, 0x54, 0x8c, 0x99, 0xbe, 0x4d, 0x25, 0x95, 0xaf, 0x72, 0x24, 0x8d, 0xf0, 0x0c,
	0x35, 0xcc, 0x92, 0xec, 0x6c, 0x79, 0x56, 0xb9, 0xfd, 0xf4, 0xb6, 0x5c, 0xe5, 0xc8, 0x19, 0xdb,
	0x5b, 0xbb, 0x1f, 0x50, 0x81, 0x2d, 0x55, 0x36, 0xb1, 0xdf, 0x1c, 0xbb, 0xe3, 0xa8, 0xc5, 0x57,
	0x39, 0x62, 0x78, 0xf8, 0x3b, 0x54, 0x54, 0x10, 0xfa, 0xd9, 0xbc, 0xbe, 0x3e, 0xc2, 0xdf, 0x7f,
	0x08, 0xd3, 0x8b, 0xa1, 0xe0, 0x8f, 0xe8, 0x73, 0x1e, 0x71, 0xcd, 0x59, 0x48, 0x7d, 0xd0, 0xde,
	0xed, 0x7f, 0xb3, 0x5f, 0x7c, 0x6a, 0x7b, 0x5e, 0x64, 0xbc, 0xf7, 0x86, 0x96, 0x8d, 0xbd, 0xdb,
	0x44, 0xe7, 0x3b, 0xb6, 0x50, 0x15, 0x83, 0xc7, 0x7d, 0x0e, 0x12, 0x17, 0xfe, 0x71, 0x2d, 0xf7,
	0x5b, 0xf4, 0x8a, 0x8b, 0xad, 0xbe, 0x58, 0x8a, 0xbb, 0xcd, 0xa1, 0x54, 0xb7, 0xf1, 0x50, 0xe7,
	0xcc, 0x94, 0x9d, 0x59, 0x37, 0xcf, 0xd3, 0xfa, 0xef, 0xfe, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x04,
	0x3b, 0xef, 0x94, 0x6d, 0x05, 0x00, 0x00,
}
