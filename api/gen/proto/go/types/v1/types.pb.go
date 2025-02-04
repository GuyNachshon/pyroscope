// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: types/v1/types.proto

package typesv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type LabelPair struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *LabelPair) Reset() {
	*x = LabelPair{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_v1_types_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LabelPair) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LabelPair) ProtoMessage() {}

func (x *LabelPair) ProtoReflect() protoreflect.Message {
	mi := &file_types_v1_types_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LabelPair.ProtoReflect.Descriptor instead.
func (*LabelPair) Descriptor() ([]byte, []int) {
	return file_types_v1_types_proto_rawDescGZIP(), []int{0}
}

func (x *LabelPair) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *LabelPair) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type ProfileType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID         string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name       string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	SampleType string `protobuf:"bytes,4,opt,name=sample_type,json=sampleType,proto3" json:"sample_type,omitempty"`
	SampleUnit string `protobuf:"bytes,5,opt,name=sample_unit,json=sampleUnit,proto3" json:"sample_unit,omitempty"`
	PeriodType string `protobuf:"bytes,6,opt,name=period_type,json=periodType,proto3" json:"period_type,omitempty"`
	PeriodUnit string `protobuf:"bytes,7,opt,name=period_unit,json=periodUnit,proto3" json:"period_unit,omitempty"`
}

func (x *ProfileType) Reset() {
	*x = ProfileType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_v1_types_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProfileType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProfileType) ProtoMessage() {}

func (x *ProfileType) ProtoReflect() protoreflect.Message {
	mi := &file_types_v1_types_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProfileType.ProtoReflect.Descriptor instead.
func (*ProfileType) Descriptor() ([]byte, []int) {
	return file_types_v1_types_proto_rawDescGZIP(), []int{1}
}

func (x *ProfileType) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *ProfileType) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ProfileType) GetSampleType() string {
	if x != nil {
		return x.SampleType
	}
	return ""
}

func (x *ProfileType) GetSampleUnit() string {
	if x != nil {
		return x.SampleUnit
	}
	return ""
}

func (x *ProfileType) GetPeriodType() string {
	if x != nil {
		return x.PeriodType
	}
	return ""
}

func (x *ProfileType) GetPeriodUnit() string {
	if x != nil {
		return x.PeriodUnit
	}
	return ""
}

type Labels struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// LabelPair is the key value pairs to identify the corresponding profile
	Labels []*LabelPair `protobuf:"bytes,1,rep,name=labels,proto3" json:"labels,omitempty"`
}

func (x *Labels) Reset() {
	*x = Labels{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_v1_types_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Labels) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Labels) ProtoMessage() {}

func (x *Labels) ProtoReflect() protoreflect.Message {
	mi := &file_types_v1_types_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Labels.ProtoReflect.Descriptor instead.
func (*Labels) Descriptor() ([]byte, []int) {
	return file_types_v1_types_proto_rawDescGZIP(), []int{2}
}

func (x *Labels) GetLabels() []*LabelPair {
	if x != nil {
		return x.Labels
	}
	return nil
}

type Series struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Labels []*LabelPair `protobuf:"bytes,1,rep,name=labels,proto3" json:"labels,omitempty"`
	Points []*Point     `protobuf:"bytes,2,rep,name=points,proto3" json:"points,omitempty"`
}

func (x *Series) Reset() {
	*x = Series{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_v1_types_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Series) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Series) ProtoMessage() {}

func (x *Series) ProtoReflect() protoreflect.Message {
	mi := &file_types_v1_types_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Series.ProtoReflect.Descriptor instead.
func (*Series) Descriptor() ([]byte, []int) {
	return file_types_v1_types_proto_rawDescGZIP(), []int{3}
}

func (x *Series) GetLabels() []*LabelPair {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *Series) GetPoints() []*Point {
	if x != nil {
		return x.Points
	}
	return nil
}

type Point struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value float64 `protobuf:"fixed64,1,opt,name=value,proto3" json:"value,omitempty"`
	// Milliseconds unix timestamp
	Timestamp int64 `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *Point) Reset() {
	*x = Point{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_v1_types_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Point) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Point) ProtoMessage() {}

func (x *Point) ProtoReflect() protoreflect.Message {
	mi := &file_types_v1_types_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Point.ProtoReflect.Descriptor instead.
func (*Point) Descriptor() ([]byte, []int) {
	return file_types_v1_types_proto_rawDescGZIP(), []int{4}
}

func (x *Point) GetValue() float64 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *Point) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

type LabelValuesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Matchers []string `protobuf:"bytes,2,rep,name=matchers,proto3" json:"matchers,omitempty"`
}

func (x *LabelValuesRequest) Reset() {
	*x = LabelValuesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_v1_types_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LabelValuesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LabelValuesRequest) ProtoMessage() {}

func (x *LabelValuesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_types_v1_types_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LabelValuesRequest.ProtoReflect.Descriptor instead.
func (*LabelValuesRequest) Descriptor() ([]byte, []int) {
	return file_types_v1_types_proto_rawDescGZIP(), []int{5}
}

func (x *LabelValuesRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *LabelValuesRequest) GetMatchers() []string {
	if x != nil {
		return x.Matchers
	}
	return nil
}

type LabelValuesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Names []string `protobuf:"bytes,1,rep,name=names,proto3" json:"names,omitempty"`
}

func (x *LabelValuesResponse) Reset() {
	*x = LabelValuesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_v1_types_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LabelValuesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LabelValuesResponse) ProtoMessage() {}

func (x *LabelValuesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_types_v1_types_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LabelValuesResponse.ProtoReflect.Descriptor instead.
func (*LabelValuesResponse) Descriptor() ([]byte, []int) {
	return file_types_v1_types_proto_rawDescGZIP(), []int{6}
}

func (x *LabelValuesResponse) GetNames() []string {
	if x != nil {
		return x.Names
	}
	return nil
}

type LabelNamesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Matchers []string `protobuf:"bytes,1,rep,name=matchers,proto3" json:"matchers,omitempty"`
}

func (x *LabelNamesRequest) Reset() {
	*x = LabelNamesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_v1_types_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LabelNamesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LabelNamesRequest) ProtoMessage() {}

func (x *LabelNamesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_types_v1_types_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LabelNamesRequest.ProtoReflect.Descriptor instead.
func (*LabelNamesRequest) Descriptor() ([]byte, []int) {
	return file_types_v1_types_proto_rawDescGZIP(), []int{7}
}

func (x *LabelNamesRequest) GetMatchers() []string {
	if x != nil {
		return x.Matchers
	}
	return nil
}

type LabelNamesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Names []string `protobuf:"bytes,1,rep,name=names,proto3" json:"names,omitempty"`
}

func (x *LabelNamesResponse) Reset() {
	*x = LabelNamesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_v1_types_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LabelNamesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LabelNamesResponse) ProtoMessage() {}

func (x *LabelNamesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_types_v1_types_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LabelNamesResponse.ProtoReflect.Descriptor instead.
func (*LabelNamesResponse) Descriptor() ([]byte, []int) {
	return file_types_v1_types_proto_rawDescGZIP(), []int{8}
}

func (x *LabelNamesResponse) GetNames() []string {
	if x != nil {
		return x.Names
	}
	return nil
}

var File_types_v1_types_proto protoreflect.FileDescriptor

var file_types_v1_types_proto_rawDesc = []byte{
	0x0a, 0x14, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31,
	0x22, 0x35, 0x0a, 0x09, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x50, 0x61, 0x69, 0x72, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0xb5, 0x01, 0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1f, 0x0a, 0x0b,
	0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x75, 0x6e, 0x69, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x55, 0x6e, 0x69, 0x74, 0x12, 0x1f, 0x0a,
	0x0b, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1f,
	0x0a, 0x0b, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x5f, 0x75, 0x6e, 0x69, 0x74, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x55, 0x6e, 0x69, 0x74, 0x22,
	0x35, 0x0a, 0x06, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x2b, 0x0a, 0x06, 0x6c, 0x61, 0x62,
	0x65, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x50, 0x61, 0x69, 0x72, 0x52, 0x06,
	0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x22, 0x5e, 0x0a, 0x06, 0x53, 0x65, 0x72, 0x69, 0x65, 0x73,
	0x12, 0x2b, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x13, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x61, 0x62, 0x65,
	0x6c, 0x50, 0x61, 0x69, 0x72, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x27, 0x0a,
	0x06, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x06,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x22, 0x3b, 0x0a, 0x05, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x22, 0x44, 0x0a, 0x12, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x08, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x73, 0x22, 0x2b, 0x0a, 0x13, 0x4c, 0x61, 0x62,
	0x65, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x05, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x22, 0x2f, 0x0a, 0x11, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x4e,
	0x61, 0x6d, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6d,
	0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x6d,
	0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x73, 0x22, 0x2a, 0x0a, 0x12, 0x4c, 0x61, 0x62, 0x65, 0x6c,
	0x4e, 0x61, 0x6d, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x61,
	0x6d, 0x65, 0x73, 0x42, 0x9b, 0x01, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x76, 0x31, 0x42, 0x0a, 0x54, 0x79, 0x70, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67,
	0x72, 0x61, 0x66, 0x61, 0x6e, 0x61, 0x2f, 0x70, 0x79, 0x72, 0x6f, 0x73, 0x63, 0x6f, 0x70, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67,
	0x6f, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x76, 0x31, 0xa2, 0x02, 0x03, 0x54, 0x58, 0x58, 0xaa, 0x02, 0x08, 0x54, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x56, 0x31, 0xca, 0x02, 0x08, 0x54, 0x79, 0x70, 0x65, 0x73, 0x5c, 0x56, 0x31, 0xe2, 0x02,
	0x14, 0x54, 0x79, 0x70, 0x65, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x09, 0x54, 0x79, 0x70, 0x65, 0x73, 0x3a, 0x3a, 0x56,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_types_v1_types_proto_rawDescOnce sync.Once
	file_types_v1_types_proto_rawDescData = file_types_v1_types_proto_rawDesc
)

func file_types_v1_types_proto_rawDescGZIP() []byte {
	file_types_v1_types_proto_rawDescOnce.Do(func() {
		file_types_v1_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_types_v1_types_proto_rawDescData)
	})
	return file_types_v1_types_proto_rawDescData
}

var file_types_v1_types_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_types_v1_types_proto_goTypes = []interface{}{
	(*LabelPair)(nil),           // 0: types.v1.LabelPair
	(*ProfileType)(nil),         // 1: types.v1.ProfileType
	(*Labels)(nil),              // 2: types.v1.Labels
	(*Series)(nil),              // 3: types.v1.Series
	(*Point)(nil),               // 4: types.v1.Point
	(*LabelValuesRequest)(nil),  // 5: types.v1.LabelValuesRequest
	(*LabelValuesResponse)(nil), // 6: types.v1.LabelValuesResponse
	(*LabelNamesRequest)(nil),   // 7: types.v1.LabelNamesRequest
	(*LabelNamesResponse)(nil),  // 8: types.v1.LabelNamesResponse
}
var file_types_v1_types_proto_depIdxs = []int32{
	0, // 0: types.v1.Labels.labels:type_name -> types.v1.LabelPair
	0, // 1: types.v1.Series.labels:type_name -> types.v1.LabelPair
	4, // 2: types.v1.Series.points:type_name -> types.v1.Point
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_types_v1_types_proto_init() }
func file_types_v1_types_proto_init() {
	if File_types_v1_types_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_types_v1_types_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LabelPair); i {
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
		file_types_v1_types_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProfileType); i {
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
		file_types_v1_types_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Labels); i {
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
		file_types_v1_types_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Series); i {
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
		file_types_v1_types_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Point); i {
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
		file_types_v1_types_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LabelValuesRequest); i {
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
		file_types_v1_types_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LabelValuesResponse); i {
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
		file_types_v1_types_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LabelNamesRequest); i {
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
		file_types_v1_types_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LabelNamesResponse); i {
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
			RawDescriptor: file_types_v1_types_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_types_v1_types_proto_goTypes,
		DependencyIndexes: file_types_v1_types_proto_depIdxs,
		MessageInfos:      file_types_v1_types_proto_msgTypes,
	}.Build()
	File_types_v1_types_proto = out.File
	file_types_v1_types_proto_rawDesc = nil
	file_types_v1_types_proto_goTypes = nil
	file_types_v1_types_proto_depIdxs = nil
}
