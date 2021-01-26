// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: envoy/api/v2/endpoint.proto

package envoy_api_v2

import (
	endpoint "github.com/cilium/proxy/go/envoy/api/v2/endpoint"
	_type "github.com/cilium/proxy/go/envoy/type"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
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

// Each route from RDS will map to a single cluster or traffic split across
// clusters using weights expressed in the RDS WeightedCluster.
//
// With EDS, each cluster is treated independently from a LB perspective, with
// LB taking place between the Localities within a cluster and at a finer
// granularity between the hosts within a locality. The percentage of traffic
// for each endpoint is determined by both its load_balancing_weight, and the
// load_balancing_weight of its locality. First, a locality will be selected,
// then an endpoint within that locality will be chose based on its weight.
// [#next-free-field: 6]
type ClusterLoadAssignment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the cluster. This will be the :ref:`service_name
	// <envoy_api_field_Cluster.EdsClusterConfig.service_name>` value if specified
	// in the cluster :ref:`EdsClusterConfig
	// <envoy_api_msg_Cluster.EdsClusterConfig>`.
	ClusterName string `protobuf:"bytes,1,opt,name=cluster_name,json=clusterName,proto3" json:"cluster_name,omitempty"`
	// List of endpoints to load balance to.
	Endpoints []*endpoint.LocalityLbEndpoints `protobuf:"bytes,2,rep,name=endpoints,proto3" json:"endpoints,omitempty"`
	// Map of named endpoints that can be referenced in LocalityLbEndpoints.
	// [#not-implemented-hide:]
	NamedEndpoints map[string]*endpoint.Endpoint `protobuf:"bytes,5,rep,name=named_endpoints,json=namedEndpoints,proto3" json:"named_endpoints,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Load balancing policy settings.
	Policy *ClusterLoadAssignment_Policy `protobuf:"bytes,4,opt,name=policy,proto3" json:"policy,omitempty"`
}

func (x *ClusterLoadAssignment) Reset() {
	*x = ClusterLoadAssignment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_api_v2_endpoint_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClusterLoadAssignment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterLoadAssignment) ProtoMessage() {}

func (x *ClusterLoadAssignment) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_api_v2_endpoint_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClusterLoadAssignment.ProtoReflect.Descriptor instead.
func (*ClusterLoadAssignment) Descriptor() ([]byte, []int) {
	return file_envoy_api_v2_endpoint_proto_rawDescGZIP(), []int{0}
}

func (x *ClusterLoadAssignment) GetClusterName() string {
	if x != nil {
		return x.ClusterName
	}
	return ""
}

func (x *ClusterLoadAssignment) GetEndpoints() []*endpoint.LocalityLbEndpoints {
	if x != nil {
		return x.Endpoints
	}
	return nil
}

func (x *ClusterLoadAssignment) GetNamedEndpoints() map[string]*endpoint.Endpoint {
	if x != nil {
		return x.NamedEndpoints
	}
	return nil
}

func (x *ClusterLoadAssignment) GetPolicy() *ClusterLoadAssignment_Policy {
	if x != nil {
		return x.Policy
	}
	return nil
}

// Load balancing policy settings.
// [#next-free-field: 6]
type ClusterLoadAssignment_Policy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Action to trim the overall incoming traffic to protect the upstream
	// hosts. This action allows protection in case the hosts are unable to
	// recover from an outage, or unable to autoscale or unable to handle
	// incoming traffic volume for any reason.
	//
	// At the client each category is applied one after the other to generate
	// the 'actual' drop percentage on all outgoing traffic. For example:
	//
	// .. code-block:: json
	//
	//  { "drop_overloads": [
	//      { "category": "throttle", "drop_percentage": 60 }
	//      { "category": "lb", "drop_percentage": 50 }
	//  ]}
	//
	// The actual drop percentages applied to the traffic at the clients will be
	//    "throttle"_drop = 60%
	//    "lb"_drop = 20%  // 50% of the remaining 'actual' load, which is 40%.
	//    actual_outgoing_load = 20% // remaining after applying all categories.
	// [#not-implemented-hide:]
	DropOverloads []*ClusterLoadAssignment_Policy_DropOverload `protobuf:"bytes,2,rep,name=drop_overloads,json=dropOverloads,proto3" json:"drop_overloads,omitempty"`
	// Priority levels and localities are considered overprovisioned with this
	// factor (in percentage). This means that we don't consider a priority
	// level or locality unhealthy until the percentage of healthy hosts
	// multiplied by the overprovisioning factor drops below 100.
	// With the default value 140(1.4), Envoy doesn't consider a priority level
	// or a locality unhealthy until their percentage of healthy hosts drops
	// below 72%. For example:
	//
	// .. code-block:: json
	//
	//  { "overprovisioning_factor": 100 }
	//
	// Read more at :ref:`priority levels <arch_overview_load_balancing_priority_levels>` and
	// :ref:`localities <arch_overview_load_balancing_locality_weighted_lb>`.
	OverprovisioningFactor *wrapperspb.UInt32Value `protobuf:"bytes,3,opt,name=overprovisioning_factor,json=overprovisioningFactor,proto3" json:"overprovisioning_factor,omitempty"`
	// The max time until which the endpoints from this assignment can be used.
	// If no new assignments are received before this time expires the endpoints
	// are considered stale and should be marked unhealthy.
	// Defaults to 0 which means endpoints never go stale.
	EndpointStaleAfter *durationpb.Duration `protobuf:"bytes,4,opt,name=endpoint_stale_after,json=endpointStaleAfter,proto3" json:"endpoint_stale_after,omitempty"`
	// The flag to disable overprovisioning. If it is set to true,
	// :ref:`overprovisioning factor
	// <arch_overview_load_balancing_overprovisioning_factor>` will be ignored
	// and Envoy will not perform graceful failover between priority levels or
	// localities as endpoints become unhealthy. Otherwise Envoy will perform
	// graceful failover as :ref:`overprovisioning factor
	// <arch_overview_load_balancing_overprovisioning_factor>` suggests.
	// [#not-implemented-hide:]
	//
	// Deprecated: Do not use.
	DisableOverprovisioning bool `protobuf:"varint,5,opt,name=disable_overprovisioning,json=disableOverprovisioning,proto3" json:"disable_overprovisioning,omitempty"`
}

func (x *ClusterLoadAssignment_Policy) Reset() {
	*x = ClusterLoadAssignment_Policy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_api_v2_endpoint_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClusterLoadAssignment_Policy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterLoadAssignment_Policy) ProtoMessage() {}

func (x *ClusterLoadAssignment_Policy) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_api_v2_endpoint_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClusterLoadAssignment_Policy.ProtoReflect.Descriptor instead.
func (*ClusterLoadAssignment_Policy) Descriptor() ([]byte, []int) {
	return file_envoy_api_v2_endpoint_proto_rawDescGZIP(), []int{0, 0}
}

func (x *ClusterLoadAssignment_Policy) GetDropOverloads() []*ClusterLoadAssignment_Policy_DropOverload {
	if x != nil {
		return x.DropOverloads
	}
	return nil
}

func (x *ClusterLoadAssignment_Policy) GetOverprovisioningFactor() *wrapperspb.UInt32Value {
	if x != nil {
		return x.OverprovisioningFactor
	}
	return nil
}

func (x *ClusterLoadAssignment_Policy) GetEndpointStaleAfter() *durationpb.Duration {
	if x != nil {
		return x.EndpointStaleAfter
	}
	return nil
}

// Deprecated: Do not use.
func (x *ClusterLoadAssignment_Policy) GetDisableOverprovisioning() bool {
	if x != nil {
		return x.DisableOverprovisioning
	}
	return false
}

// [#not-implemented-hide:]
type ClusterLoadAssignment_Policy_DropOverload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Identifier for the policy specifying the drop.
	Category string `protobuf:"bytes,1,opt,name=category,proto3" json:"category,omitempty"`
	// Percentage of traffic that should be dropped for the category.
	DropPercentage *_type.FractionalPercent `protobuf:"bytes,2,opt,name=drop_percentage,json=dropPercentage,proto3" json:"drop_percentage,omitempty"`
}

func (x *ClusterLoadAssignment_Policy_DropOverload) Reset() {
	*x = ClusterLoadAssignment_Policy_DropOverload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_api_v2_endpoint_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClusterLoadAssignment_Policy_DropOverload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterLoadAssignment_Policy_DropOverload) ProtoMessage() {}

func (x *ClusterLoadAssignment_Policy_DropOverload) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_api_v2_endpoint_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClusterLoadAssignment_Policy_DropOverload.ProtoReflect.Descriptor instead.
func (*ClusterLoadAssignment_Policy_DropOverload) Descriptor() ([]byte, []int) {
	return file_envoy_api_v2_endpoint_proto_rawDescGZIP(), []int{0, 0, 0}
}

func (x *ClusterLoadAssignment_Policy_DropOverload) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *ClusterLoadAssignment_Policy_DropOverload) GetDropPercentage() *_type.FractionalPercent {
	if x != nil {
		return x.DropPercentage
	}
	return nil
}

var File_envoy_api_v2_endpoint_proto protoreflect.FileDescriptor

var file_envoy_api_v2_endpoint_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x2f, 0x65,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x1a, 0x2f, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x2f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x2f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x63, 0x6f, 0x6d, 0x70,
	0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x6d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfb, 0x06, 0x0a,
	0x15, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4c, 0x6f, 0x61, 0x64, 0x41, 0x73, 0x73, 0x69,
	0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x2a, 0x0a, 0x0c, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42,
	0x04, 0x72, 0x02, 0x20, 0x01, 0x52, 0x0b, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x48, 0x0a, 0x09, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x32, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x4c, 0x6f,
	0x63, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x4c, 0x62, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x73, 0x52, 0x09, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x60, 0x0a, 0x0f,
	0x6e, 0x61, 0x6d, 0x65, 0x64, 0x5f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18,
	0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x37, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x32, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4c, 0x6f, 0x61, 0x64,
	0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x64,
	0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0e,
	0x6e, 0x61, 0x6d, 0x65, 0x64, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x42,
	0x0a, 0x06, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a,
	0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x43, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x4c, 0x6f, 0x61, 0x64, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d,
	0x65, 0x6e, 0x74, 0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x06, 0x70, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x1a, 0xe1, 0x03, 0x0a, 0x06, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x5e, 0x0a,
	0x0e, 0x64, 0x72, 0x6f, 0x70, 0x5f, 0x6f, 0x76, 0x65, 0x72, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x37, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x32, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4c, 0x6f, 0x61, 0x64,
	0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x2e, 0x44, 0x72, 0x6f, 0x70, 0x4f, 0x76, 0x65, 0x72, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x0d,
	0x64, 0x72, 0x6f, 0x70, 0x4f, 0x76, 0x65, 0x72, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x12, 0x5e, 0x0a,
	0x17, 0x6f, 0x76, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e,
	0x67, 0x5f, 0x66, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x07, 0xfa, 0x42,
	0x04, 0x2a, 0x02, 0x20, 0x00, 0x52, 0x16, 0x6f, 0x76, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x76, 0x69,
	0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x46, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x55, 0x0a,
	0x14, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x6c, 0x65, 0x5f,
	0x61, 0x66, 0x74, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x08, 0xfa, 0x42, 0x05, 0xaa, 0x01, 0x02, 0x2a, 0x00,
	0x52, 0x12, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x6c, 0x65, 0x41,
	0x66, 0x74, 0x65, 0x72, 0x12, 0x3d, 0x0a, 0x18, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x5f,
	0x6f, 0x76, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x42, 0x02, 0x18, 0x01, 0x52, 0x17, 0x64, 0x69, 0x73, 0x61,
	0x62, 0x6c, 0x65, 0x4f, 0x76, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e,
	0x69, 0x6e, 0x67, 0x1a, 0x7b, 0x0a, 0x0c, 0x44, 0x72, 0x6f, 0x70, 0x4f, 0x76, 0x65, 0x72, 0x6c,
	0x6f, 0x61, 0x64, 0x12, 0x23, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x20, 0x01, 0x52, 0x08,
	0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x46, 0x0a, 0x0f, 0x64, 0x72, 0x6f, 0x70,
	0x5f, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1d, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x46,
	0x72, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74,
	0x52, 0x0e, 0x64, 0x72, 0x6f, 0x70, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65,
	0x4a, 0x04, 0x08, 0x01, 0x10, 0x02, 0x1a, 0x62, 0x0a, 0x13, 0x4e, 0x61, 0x6d, 0x65, 0x64, 0x45,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x35, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f,
	0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x65, 0x6e,
	0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x55, 0x0a, 0x1a, 0x69, 0x6f,
	0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x42, 0x0d, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0xf2, 0x98, 0xfe, 0x8f, 0x05, 0x1a, 0x12,
	0x18, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x65, 0x6e,
	0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x76, 0x33, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10,
	0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_api_v2_endpoint_proto_rawDescOnce sync.Once
	file_envoy_api_v2_endpoint_proto_rawDescData = file_envoy_api_v2_endpoint_proto_rawDesc
)

func file_envoy_api_v2_endpoint_proto_rawDescGZIP() []byte {
	file_envoy_api_v2_endpoint_proto_rawDescOnce.Do(func() {
		file_envoy_api_v2_endpoint_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_api_v2_endpoint_proto_rawDescData)
	})
	return file_envoy_api_v2_endpoint_proto_rawDescData
}

var file_envoy_api_v2_endpoint_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_envoy_api_v2_endpoint_proto_goTypes = []interface{}{
	(*ClusterLoadAssignment)(nil),        // 0: envoy.api.v2.ClusterLoadAssignment
	(*ClusterLoadAssignment_Policy)(nil), // 1: envoy.api.v2.ClusterLoadAssignment.Policy
	nil,                                  // 2: envoy.api.v2.ClusterLoadAssignment.NamedEndpointsEntry
	(*ClusterLoadAssignment_Policy_DropOverload)(nil), // 3: envoy.api.v2.ClusterLoadAssignment.Policy.DropOverload
	(*endpoint.LocalityLbEndpoints)(nil),              // 4: envoy.api.v2.endpoint.LocalityLbEndpoints
	(*wrapperspb.UInt32Value)(nil),                    // 5: google.protobuf.UInt32Value
	(*durationpb.Duration)(nil),                       // 6: google.protobuf.Duration
	(*endpoint.Endpoint)(nil),                         // 7: envoy.api.v2.endpoint.Endpoint
	(*_type.FractionalPercent)(nil),                   // 8: envoy.type.FractionalPercent
}
var file_envoy_api_v2_endpoint_proto_depIdxs = []int32{
	4, // 0: envoy.api.v2.ClusterLoadAssignment.endpoints:type_name -> envoy.api.v2.endpoint.LocalityLbEndpoints
	2, // 1: envoy.api.v2.ClusterLoadAssignment.named_endpoints:type_name -> envoy.api.v2.ClusterLoadAssignment.NamedEndpointsEntry
	1, // 2: envoy.api.v2.ClusterLoadAssignment.policy:type_name -> envoy.api.v2.ClusterLoadAssignment.Policy
	3, // 3: envoy.api.v2.ClusterLoadAssignment.Policy.drop_overloads:type_name -> envoy.api.v2.ClusterLoadAssignment.Policy.DropOverload
	5, // 4: envoy.api.v2.ClusterLoadAssignment.Policy.overprovisioning_factor:type_name -> google.protobuf.UInt32Value
	6, // 5: envoy.api.v2.ClusterLoadAssignment.Policy.endpoint_stale_after:type_name -> google.protobuf.Duration
	7, // 6: envoy.api.v2.ClusterLoadAssignment.NamedEndpointsEntry.value:type_name -> envoy.api.v2.endpoint.Endpoint
	8, // 7: envoy.api.v2.ClusterLoadAssignment.Policy.DropOverload.drop_percentage:type_name -> envoy.type.FractionalPercent
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_envoy_api_v2_endpoint_proto_init() }
func file_envoy_api_v2_endpoint_proto_init() {
	if File_envoy_api_v2_endpoint_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_api_v2_endpoint_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClusterLoadAssignment); i {
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
		file_envoy_api_v2_endpoint_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClusterLoadAssignment_Policy); i {
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
		file_envoy_api_v2_endpoint_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClusterLoadAssignment_Policy_DropOverload); i {
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
			RawDescriptor: file_envoy_api_v2_endpoint_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_api_v2_endpoint_proto_goTypes,
		DependencyIndexes: file_envoy_api_v2_endpoint_proto_depIdxs,
		MessageInfos:      file_envoy_api_v2_endpoint_proto_msgTypes,
	}.Build()
	File_envoy_api_v2_endpoint_proto = out.File
	file_envoy_api_v2_endpoint_proto_rawDesc = nil
	file_envoy_api_v2_endpoint_proto_goTypes = nil
	file_envoy_api_v2_endpoint_proto_depIdxs = nil
}
