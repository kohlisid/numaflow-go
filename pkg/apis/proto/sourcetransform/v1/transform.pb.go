// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.1
// source: pkg/apis/proto/sourcetransform/v1/transform.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// *
// SourceTransformerRequest represents a request element.
type SourceTransformRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keys      []string               `protobuf:"bytes,1,rep,name=keys,proto3" json:"keys,omitempty"`
	Value     []byte                 `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	EventTime *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=event_time,json=eventTime,proto3" json:"event_time,omitempty"`
	Watermark *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=watermark,proto3" json:"watermark,omitempty"`
	Headers   map[string]string      `protobuf:"bytes,5,rep,name=headers,proto3" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *SourceTransformRequest) Reset() {
	*x = SourceTransformRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_apis_proto_sourcetransform_v1_transform_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SourceTransformRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SourceTransformRequest) ProtoMessage() {}

func (x *SourceTransformRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_apis_proto_sourcetransform_v1_transform_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SourceTransformRequest.ProtoReflect.Descriptor instead.
func (*SourceTransformRequest) Descriptor() ([]byte, []int) {
	return file_pkg_apis_proto_sourcetransform_v1_transform_proto_rawDescGZIP(), []int{0}
}

func (x *SourceTransformRequest) GetKeys() []string {
	if x != nil {
		return x.Keys
	}
	return nil
}

func (x *SourceTransformRequest) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *SourceTransformRequest) GetEventTime() *timestamppb.Timestamp {
	if x != nil {
		return x.EventTime
	}
	return nil
}

func (x *SourceTransformRequest) GetWatermark() *timestamppb.Timestamp {
	if x != nil {
		return x.Watermark
	}
	return nil
}

func (x *SourceTransformRequest) GetHeaders() map[string]string {
	if x != nil {
		return x.Headers
	}
	return nil
}

// *
// SourceTransformerResponse represents a response element.
type SourceTransformResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results []*SourceTransformResponse_Result `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
}

func (x *SourceTransformResponse) Reset() {
	*x = SourceTransformResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_apis_proto_sourcetransform_v1_transform_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SourceTransformResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SourceTransformResponse) ProtoMessage() {}

func (x *SourceTransformResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_apis_proto_sourcetransform_v1_transform_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SourceTransformResponse.ProtoReflect.Descriptor instead.
func (*SourceTransformResponse) Descriptor() ([]byte, []int) {
	return file_pkg_apis_proto_sourcetransform_v1_transform_proto_rawDescGZIP(), []int{1}
}

func (x *SourceTransformResponse) GetResults() []*SourceTransformResponse_Result {
	if x != nil {
		return x.Results
	}
	return nil
}

// *
// ReadyResponse is the health check result.
type ReadyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ready bool `protobuf:"varint,1,opt,name=ready,proto3" json:"ready,omitempty"`
}

func (x *ReadyResponse) Reset() {
	*x = ReadyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_apis_proto_sourcetransform_v1_transform_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadyResponse) ProtoMessage() {}

func (x *ReadyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_apis_proto_sourcetransform_v1_transform_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadyResponse.ProtoReflect.Descriptor instead.
func (*ReadyResponse) Descriptor() ([]byte, []int) {
	return file_pkg_apis_proto_sourcetransform_v1_transform_proto_rawDescGZIP(), []int{2}
}

func (x *ReadyResponse) GetReady() bool {
	if x != nil {
		return x.Ready
	}
	return false
}

type SourceTransformResponse_Result struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keys      []string               `protobuf:"bytes,1,rep,name=keys,proto3" json:"keys,omitempty"`
	Value     []byte                 `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	EventTime *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=event_time,json=eventTime,proto3" json:"event_time,omitempty"`
	Tags      []string               `protobuf:"bytes,4,rep,name=tags,proto3" json:"tags,omitempty"`
}

func (x *SourceTransformResponse_Result) Reset() {
	*x = SourceTransformResponse_Result{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_apis_proto_sourcetransform_v1_transform_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SourceTransformResponse_Result) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SourceTransformResponse_Result) ProtoMessage() {}

func (x *SourceTransformResponse_Result) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_apis_proto_sourcetransform_v1_transform_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SourceTransformResponse_Result.ProtoReflect.Descriptor instead.
func (*SourceTransformResponse_Result) Descriptor() ([]byte, []int) {
	return file_pkg_apis_proto_sourcetransform_v1_transform_proto_rawDescGZIP(), []int{1, 0}
}

func (x *SourceTransformResponse_Result) GetKeys() []string {
	if x != nil {
		return x.Keys
	}
	return nil
}

func (x *SourceTransformResponse_Result) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *SourceTransformResponse_Result) GetEventTime() *timestamppb.Timestamp {
	if x != nil {
		return x.EventTime
	}
	return nil
}

func (x *SourceTransformResponse_Result) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

var File_pkg_apis_proto_sourcetransform_v1_transform_proto protoreflect.FileDescriptor

var file_pkg_apis_proto_sourcetransform_v1_transform_proto_rawDesc = []byte{
	0x0a, 0x31, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d,
	0x2f, 0x76, 0x31, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x14, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x66, 0x6f, 0x72, 0x6d, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc8, 0x02, 0x0a, 0x16, 0x53, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x39, 0x0a, 0x0a,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x38, 0x0a, 0x09, 0x77, 0x61, 0x74, 0x65, 0x72,
	0x6d, 0x61, 0x72, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x77, 0x61, 0x74, 0x65, 0x72, 0x6d, 0x61, 0x72,
	0x6b, 0x12, 0x53, 0x0a, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x18, 0x05, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x39, 0x2e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x66, 0x6f, 0x72, 0x6d, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x68,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x1a, 0x3a, 0x0a, 0x0c, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x22, 0xed, 0x01, 0x0a, 0x17, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4e,
	0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x34, 0x2e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72,
	0x6d, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x1a, 0x81,
	0x01, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x65, 0x79,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61,
	0x67, 0x73, 0x22, 0x25, 0x0a, 0x0d, 0x52, 0x65, 0x61, 0x64, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x65, 0x61, 0x64, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x05, 0x72, 0x65, 0x61, 0x64, 0x79, 0x32, 0xcb, 0x01, 0x0a, 0x0f, 0x53, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x70, 0x0a,
	0x11, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d,
	0x46, 0x6e, 0x12, 0x2c, 0x2e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x66, 0x6f, 0x72, 0x6d, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x2d, 0x2e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f,
	0x72, 0x6d, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x46, 0x0a, 0x07, 0x49, 0x73, 0x52, 0x65, 0x61, 0x64, 0x79, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x1a, 0x23, 0x2e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x66, 0x6f, 0x72, 0x6d, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x43, 0x5a, 0x41, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x75, 0x6d, 0x61, 0x70, 0x72, 0x6f, 0x6a, 0x2f, 0x6e,
	0x75, 0x6d, 0x61, 0x66, 0x6c, 0x6f, 0x77, 0x2d, 0x67, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_apis_proto_sourcetransform_v1_transform_proto_rawDescOnce sync.Once
	file_pkg_apis_proto_sourcetransform_v1_transform_proto_rawDescData = file_pkg_apis_proto_sourcetransform_v1_transform_proto_rawDesc
)

func file_pkg_apis_proto_sourcetransform_v1_transform_proto_rawDescGZIP() []byte {
	file_pkg_apis_proto_sourcetransform_v1_transform_proto_rawDescOnce.Do(func() {
		file_pkg_apis_proto_sourcetransform_v1_transform_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_apis_proto_sourcetransform_v1_transform_proto_rawDescData)
	})
	return file_pkg_apis_proto_sourcetransform_v1_transform_proto_rawDescData
}

var file_pkg_apis_proto_sourcetransform_v1_transform_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_pkg_apis_proto_sourcetransform_v1_transform_proto_goTypes = []interface{}{
	(*SourceTransformRequest)(nil),         // 0: sourcetransformer.v1.SourceTransformRequest
	(*SourceTransformResponse)(nil),        // 1: sourcetransformer.v1.SourceTransformResponse
	(*ReadyResponse)(nil),                  // 2: sourcetransformer.v1.ReadyResponse
	nil,                                    // 3: sourcetransformer.v1.SourceTransformRequest.HeadersEntry
	(*SourceTransformResponse_Result)(nil), // 4: sourcetransformer.v1.SourceTransformResponse.Result
	(*timestamppb.Timestamp)(nil),          // 5: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),                  // 6: google.protobuf.Empty
}
var file_pkg_apis_proto_sourcetransform_v1_transform_proto_depIdxs = []int32{
	5, // 0: sourcetransformer.v1.SourceTransformRequest.event_time:type_name -> google.protobuf.Timestamp
	5, // 1: sourcetransformer.v1.SourceTransformRequest.watermark:type_name -> google.protobuf.Timestamp
	3, // 2: sourcetransformer.v1.SourceTransformRequest.headers:type_name -> sourcetransformer.v1.SourceTransformRequest.HeadersEntry
	4, // 3: sourcetransformer.v1.SourceTransformResponse.results:type_name -> sourcetransformer.v1.SourceTransformResponse.Result
	5, // 4: sourcetransformer.v1.SourceTransformResponse.Result.event_time:type_name -> google.protobuf.Timestamp
	0, // 5: sourcetransformer.v1.SourceTransform.SourceTransformFn:input_type -> sourcetransformer.v1.SourceTransformRequest
	6, // 6: sourcetransformer.v1.SourceTransform.IsReady:input_type -> google.protobuf.Empty
	1, // 7: sourcetransformer.v1.SourceTransform.SourceTransformFn:output_type -> sourcetransformer.v1.SourceTransformResponse
	2, // 8: sourcetransformer.v1.SourceTransform.IsReady:output_type -> sourcetransformer.v1.ReadyResponse
	7, // [7:9] is the sub-list for method output_type
	5, // [5:7] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_pkg_apis_proto_sourcetransform_v1_transform_proto_init() }
func file_pkg_apis_proto_sourcetransform_v1_transform_proto_init() {
	if File_pkg_apis_proto_sourcetransform_v1_transform_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_apis_proto_sourcetransform_v1_transform_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SourceTransformRequest); i {
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
		file_pkg_apis_proto_sourcetransform_v1_transform_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SourceTransformResponse); i {
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
		file_pkg_apis_proto_sourcetransform_v1_transform_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadyResponse); i {
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
		file_pkg_apis_proto_sourcetransform_v1_transform_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SourceTransformResponse_Result); i {
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
			RawDescriptor: file_pkg_apis_proto_sourcetransform_v1_transform_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_apis_proto_sourcetransform_v1_transform_proto_goTypes,
		DependencyIndexes: file_pkg_apis_proto_sourcetransform_v1_transform_proto_depIdxs,
		MessageInfos:      file_pkg_apis_proto_sourcetransform_v1_transform_proto_msgTypes,
	}.Build()
	File_pkg_apis_proto_sourcetransform_v1_transform_proto = out.File
	file_pkg_apis_proto_sourcetransform_v1_transform_proto_rawDesc = nil
	file_pkg_apis_proto_sourcetransform_v1_transform_proto_goTypes = nil
	file_pkg_apis_proto_sourcetransform_v1_transform_proto_depIdxs = nil
}
