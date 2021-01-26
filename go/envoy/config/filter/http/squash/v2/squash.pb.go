// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: envoy/config/filter/http/squash/v2/squash.proto

package envoy_config_filter_http_squash_v2

import (
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// [#next-free-field: 6]
type Squash struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the cluster that hosts the Squash server.
	Cluster string `protobuf:"bytes,1,opt,name=cluster,proto3" json:"cluster,omitempty"`
	// When the filter requests the Squash server to create a DebugAttachment, it will use this
	// structure as template for the body of the request. It can contain reference to environment
	// variables in the form of '{{ ENV_VAR_NAME }}'. These can be used to provide the Squash server
	// with more information to find the process to attach the debugger to. For example, in a
	// Istio/k8s environment, this will contain information on the pod:
	//
	// .. code-block:: json
	//
	//  {
	//    "spec": {
	//      "attachment": {
	//        "pod": "{{ POD_NAME }}",
	//        "namespace": "{{ POD_NAMESPACE }}"
	//      },
	//      "match_request": true
	//    }
	//  }
	//
	// (where POD_NAME, POD_NAMESPACE are configured in the pod via the Downward API)
	AttachmentTemplate *structpb.Struct `protobuf:"bytes,2,opt,name=attachment_template,json=attachmentTemplate,proto3" json:"attachment_template,omitempty"`
	// The timeout for individual requests sent to the Squash cluster. Defaults to 1 second.
	RequestTimeout *durationpb.Duration `protobuf:"bytes,3,opt,name=request_timeout,json=requestTimeout,proto3" json:"request_timeout,omitempty"`
	// The total timeout Squash will delay a request and wait for it to be attached. Defaults to 60
	// seconds.
	AttachmentTimeout *durationpb.Duration `protobuf:"bytes,4,opt,name=attachment_timeout,json=attachmentTimeout,proto3" json:"attachment_timeout,omitempty"`
	// Amount of time to poll for the status of the attachment object in the Squash server
	// (to check if has been attached). Defaults to 1 second.
	AttachmentPollPeriod *durationpb.Duration `protobuf:"bytes,5,opt,name=attachment_poll_period,json=attachmentPollPeriod,proto3" json:"attachment_poll_period,omitempty"`
}

func (x *Squash) Reset() {
	*x = Squash{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_config_filter_http_squash_v2_squash_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Squash) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Squash) ProtoMessage() {}

func (x *Squash) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_config_filter_http_squash_v2_squash_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Squash.ProtoReflect.Descriptor instead.
func (*Squash) Descriptor() ([]byte, []int) {
	return file_envoy_config_filter_http_squash_v2_squash_proto_rawDescGZIP(), []int{0}
}

func (x *Squash) GetCluster() string {
	if x != nil {
		return x.Cluster
	}
	return ""
}

func (x *Squash) GetAttachmentTemplate() *structpb.Struct {
	if x != nil {
		return x.AttachmentTemplate
	}
	return nil
}

func (x *Squash) GetRequestTimeout() *durationpb.Duration {
	if x != nil {
		return x.RequestTimeout
	}
	return nil
}

func (x *Squash) GetAttachmentTimeout() *durationpb.Duration {
	if x != nil {
		return x.AttachmentTimeout
	}
	return nil
}

func (x *Squash) GetAttachmentPollPeriod() *durationpb.Duration {
	if x != nil {
		return x.AttachmentPollPeriod
	}
	return nil
}

var File_envoy_config_filter_http_squash_v2_squash_proto protoreflect.FileDescriptor

var file_envoy_config_filter_http_squash_v2_squash_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x66,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f, 0x73, 0x71, 0x75, 0x61, 0x73,
	0x68, 0x2f, 0x76, 0x32, 0x2f, 0x73, 0x71, 0x75, 0x61, 0x73, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x22, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x73, 0x71, 0x75, 0x61,
	0x73, 0x68, 0x2e, 0x76, 0x32, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd4, 0x02, 0x0a, 0x06,
	0x53, 0x71, 0x75, 0x61, 0x73, 0x68, 0x12, 0x21, 0x0a, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x20, 0x01,
	0x52, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x48, 0x0a, 0x13, 0x61, 0x74, 0x74,
	0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52,
	0x12, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x12, 0x42, 0x0a, 0x0f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x48, 0x0a, 0x12, 0x61, 0x74, 0x74, 0x61, 0x63,
	0x68, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x11,
	0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75,
	0x74, 0x12, 0x4f, 0x0a, 0x16, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x5f,
	0x70, 0x6f, 0x6c, 0x6c, 0x5f, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x14, 0x61, 0x74,
	0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x50, 0x6f, 0x6c, 0x6c, 0x50, 0x65, 0x72, 0x69,
	0x6f, 0x64, 0x42, 0x78, 0x0a, 0x30, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72,
	0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x73, 0x71, 0x75,
	0x61, 0x73, 0x68, 0x2e, 0x76, 0x32, 0x42, 0x0b, 0x53, 0x71, 0x75, 0x61, 0x73, 0x68, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0xf2, 0x98, 0xfe, 0x8f, 0x05, 0x29, 0x12, 0x27, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x73, 0x71, 0x75, 0x61, 0x73,
	0x68, 0x2e, 0x76, 0x33, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x01, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_config_filter_http_squash_v2_squash_proto_rawDescOnce sync.Once
	file_envoy_config_filter_http_squash_v2_squash_proto_rawDescData = file_envoy_config_filter_http_squash_v2_squash_proto_rawDesc
)

func file_envoy_config_filter_http_squash_v2_squash_proto_rawDescGZIP() []byte {
	file_envoy_config_filter_http_squash_v2_squash_proto_rawDescOnce.Do(func() {
		file_envoy_config_filter_http_squash_v2_squash_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_config_filter_http_squash_v2_squash_proto_rawDescData)
	})
	return file_envoy_config_filter_http_squash_v2_squash_proto_rawDescData
}

var file_envoy_config_filter_http_squash_v2_squash_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_envoy_config_filter_http_squash_v2_squash_proto_goTypes = []interface{}{
	(*Squash)(nil),              // 0: envoy.config.filter.http.squash.v2.Squash
	(*structpb.Struct)(nil),     // 1: google.protobuf.Struct
	(*durationpb.Duration)(nil), // 2: google.protobuf.Duration
}
var file_envoy_config_filter_http_squash_v2_squash_proto_depIdxs = []int32{
	1, // 0: envoy.config.filter.http.squash.v2.Squash.attachment_template:type_name -> google.protobuf.Struct
	2, // 1: envoy.config.filter.http.squash.v2.Squash.request_timeout:type_name -> google.protobuf.Duration
	2, // 2: envoy.config.filter.http.squash.v2.Squash.attachment_timeout:type_name -> google.protobuf.Duration
	2, // 3: envoy.config.filter.http.squash.v2.Squash.attachment_poll_period:type_name -> google.protobuf.Duration
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_envoy_config_filter_http_squash_v2_squash_proto_init() }
func file_envoy_config_filter_http_squash_v2_squash_proto_init() {
	if File_envoy_config_filter_http_squash_v2_squash_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_config_filter_http_squash_v2_squash_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Squash); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_envoy_config_filter_http_squash_v2_squash_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_config_filter_http_squash_v2_squash_proto_goTypes,
		DependencyIndexes: file_envoy_config_filter_http_squash_v2_squash_proto_depIdxs,
		MessageInfos:      file_envoy_config_filter_http_squash_v2_squash_proto_msgTypes,
	}.Build()
	File_envoy_config_filter_http_squash_v2_squash_proto = out.File
	file_envoy_config_filter_http_squash_v2_squash_proto_rawDesc = nil
	file_envoy_config_filter_http_squash_v2_squash_proto_goTypes = nil
	file_envoy_config_filter_http_squash_v2_squash_proto_depIdxs = nil
}
