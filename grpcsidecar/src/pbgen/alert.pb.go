// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.15.8
// source: alert.proto

package pbgen

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type URGENCY int32

const (
	URGENCY_LAX       URGENCY = 0
	URGENCY_IMMEDIATE URGENCY = 2
	URGENCY_IMPORTANT URGENCY = 1
)

// Enum value maps for URGENCY.
var (
	URGENCY_name = map[int32]string{
		0: "LAX",
		2: "IMMEDIATE",
		1: "IMPORTANT",
	}
	URGENCY_value = map[string]int32{
		"LAX":       0,
		"IMMEDIATE": 2,
		"IMPORTANT": 1,
	}
)

func (x URGENCY) Enum() *URGENCY {
	p := new(URGENCY)
	*p = x
	return p
}

func (x URGENCY) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (URGENCY) Descriptor() protoreflect.EnumDescriptor {
	return file_alert_proto_enumTypes[0].Descriptor()
}

func (URGENCY) Type() protoreflect.EnumType {
	return &file_alert_proto_enumTypes[0]
}

func (x URGENCY) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use URGENCY.Descriptor instead.
func (URGENCY) EnumDescriptor() ([]byte, []int) {
	return file_alert_proto_rawDescGZIP(), []int{0}
}

type AlertDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Updateflag  bool                   `protobuf:"varint,1,opt,name=updateflag,proto3" json:"updateflag,omitempty"`
	Alerttime   *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=alerttime,proto3" json:"alerttime,omitempty"`
	URLS        []string               `protobuf:"bytes,3,rep,name=URLS,proto3" json:"URLS,omitempty"`
	Callinghost string                 `protobuf:"bytes,5,opt,name=callinghost,proto3" json:"callinghost,omitempty"`
	Stallcap    URGENCY                `protobuf:"varint,6,opt,name=stallcap,proto3,enum=alerter.URGENCY" json:"stallcap,omitempty"`
	Name        string                 `protobuf:"bytes,7,opt,name=name,proto3" json:"name,omitempty"`
	K8Sorigin   *K8SInstance           `protobuf:"bytes,8,opt,name=k8sorigin,proto3" json:"k8sorigin,omitempty"`
}

func (x *AlertDetail) Reset() {
	*x = AlertDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alert_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AlertDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlertDetail) ProtoMessage() {}

func (x *AlertDetail) ProtoReflect() protoreflect.Message {
	mi := &file_alert_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlertDetail.ProtoReflect.Descriptor instead.
func (*AlertDetail) Descriptor() ([]byte, []int) {
	return file_alert_proto_rawDescGZIP(), []int{0}
}

func (x *AlertDetail) GetUpdateflag() bool {
	if x != nil {
		return x.Updateflag
	}
	return false
}

func (x *AlertDetail) GetAlerttime() *timestamppb.Timestamp {
	if x != nil {
		return x.Alerttime
	}
	return nil
}

func (x *AlertDetail) GetURLS() []string {
	if x != nil {
		return x.URLS
	}
	return nil
}

func (x *AlertDetail) GetCallinghost() string {
	if x != nil {
		return x.Callinghost
	}
	return ""
}

func (x *AlertDetail) GetStallcap() URGENCY {
	if x != nil {
		return x.Stallcap
	}
	return URGENCY_LAX
}

func (x *AlertDetail) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AlertDetail) GetK8Sorigin() *K8SInstance {
	if x != nil {
		return x.K8Sorigin
	}
	return nil
}

type AlertList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Alerts []*AlertDetail `protobuf:"bytes,1,rep,name=alerts,proto3" json:"alerts,omitempty"`
}

func (x *AlertList) Reset() {
	*x = AlertList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alert_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AlertList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlertList) ProtoMessage() {}

func (x *AlertList) ProtoReflect() protoreflect.Message {
	mi := &file_alert_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlertList.ProtoReflect.Descriptor instead.
func (*AlertList) Descriptor() ([]byte, []int) {
	return file_alert_proto_rawDescGZIP(), []int{1}
}

func (x *AlertList) GetAlerts() []*AlertDetail {
	if x != nil {
		return x.Alerts
	}
	return nil
}

type AlertRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Alertsbatch *AlertList `protobuf:"bytes,1,opt,name=alertsbatch,proto3" json:"alertsbatch,omitempty"`
	Cid         string     `protobuf:"bytes,2,opt,name=cid,proto3" json:"cid,omitempty"`
}

func (x *AlertRequest) Reset() {
	*x = AlertRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alert_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AlertRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlertRequest) ProtoMessage() {}

func (x *AlertRequest) ProtoReflect() protoreflect.Message {
	mi := &file_alert_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlertRequest.ProtoReflect.Descriptor instead.
func (*AlertRequest) Descriptor() ([]byte, []int) {
	return file_alert_proto_rawDescGZIP(), []int{2}
}

func (x *AlertRequest) GetAlertsbatch() *AlertList {
	if x != nil {
		return x.Alertsbatch
	}
	return nil
}

func (x *AlertRequest) GetCid() string {
	if x != nil {
		return x.Cid
	}
	return ""
}

type AlertResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Seen string `protobuf:"bytes,1,opt,name=seen,proto3" json:"seen,omitempty"`
}

func (x *AlertResponse) Reset() {
	*x = AlertResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alert_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AlertResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlertResponse) ProtoMessage() {}

func (x *AlertResponse) ProtoReflect() protoreflect.Message {
	mi := &file_alert_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlertResponse.ProtoReflect.Descriptor instead.
func (*AlertResponse) Descriptor() ([]byte, []int) {
	return file_alert_proto_rawDescGZIP(), []int{3}
}

func (x *AlertResponse) GetSeen() string {
	if x != nil {
		return x.Seen
	}
	return ""
}

type DescriptionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Giveme string `protobuf:"bytes,1,opt,name=giveme,proto3" json:"giveme,omitempty"`
}

func (x *DescriptionRequest) Reset() {
	*x = DescriptionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alert_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescriptionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescriptionRequest) ProtoMessage() {}

func (x *DescriptionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_alert_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescriptionRequest.ProtoReflect.Descriptor instead.
func (*DescriptionRequest) Descriptor() ([]byte, []int) {
	return file_alert_proto_rawDescGZIP(), []int{4}
}

func (x *DescriptionRequest) GetGiveme() string {
	if x != nil {
		return x.Giveme
	}
	return ""
}

type RetrieveAlertRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cid string `protobuf:"bytes,1,opt,name=cid,proto3" json:"cid,omitempty"`
}

func (x *RetrieveAlertRequest) Reset() {
	*x = RetrieveAlertRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alert_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RetrieveAlertRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RetrieveAlertRequest) ProtoMessage() {}

func (x *RetrieveAlertRequest) ProtoReflect() protoreflect.Message {
	mi := &file_alert_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RetrieveAlertRequest.ProtoReflect.Descriptor instead.
func (*RetrieveAlertRequest) Descriptor() ([]byte, []int) {
	return file_alert_proto_rawDescGZIP(), []int{5}
}

func (x *RetrieveAlertRequest) GetCid() string {
	if x != nil {
		return x.Cid
	}
	return ""
}

type AlertSchema struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Schema string `protobuf:"bytes,1,opt,name=schema,proto3" json:"schema,omitempty"`
}

func (x *AlertSchema) Reset() {
	*x = AlertSchema{}
	if protoimpl.UnsafeEnabled {
		mi := &file_alert_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AlertSchema) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlertSchema) ProtoMessage() {}

func (x *AlertSchema) ProtoReflect() protoreflect.Message {
	mi := &file_alert_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlertSchema.ProtoReflect.Descriptor instead.
func (*AlertSchema) Descriptor() ([]byte, []int) {
	return file_alert_proto_rawDescGZIP(), []int{6}
}

func (x *AlertSchema) GetSchema() string {
	if x != nil {
		return x.Schema
	}
	return ""
}

var File_alert_proto protoreflect.FileDescriptor

var file_alert_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x61,
	0x6c, 0x65, 0x72, 0x74, 0x65, 0x72, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x93, 0x02, 0x0a, 0x0b, 0x41, 0x6c, 0x65, 0x72,
	0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x1e, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x66, 0x6c, 0x61, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x66, 0x6c, 0x61, 0x67, 0x12, 0x38, 0x0a, 0x09, 0x61, 0x6c, 0x65, 0x72, 0x74,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x74, 0x69, 0x6d,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x55, 0x52, 0x4c, 0x53, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x04, 0x55, 0x52, 0x4c, 0x53, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x61, 0x6c, 0x6c, 0x69, 0x6e, 0x67,
	0x68, 0x6f, 0x73, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x61, 0x6c, 0x6c,
	0x69, 0x6e, 0x67, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x08, 0x73, 0x74, 0x61, 0x6c, 0x6c,
	0x63, 0x61, 0x70, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x61, 0x6c, 0x65, 0x72,
	0x74, 0x65, 0x72, 0x2e, 0x55, 0x52, 0x47, 0x45, 0x4e, 0x43, 0x59, 0x52, 0x08, 0x73, 0x74, 0x61,
	0x6c, 0x6c, 0x63, 0x61, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x32, 0x0a, 0x09, 0x6b, 0x38, 0x73,
	0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x61,
	0x6c, 0x65, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x4b, 0x38, 0x73, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x52, 0x09, 0x6b, 0x38, 0x73, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x22, 0x39, 0x0a,
	0x09, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x06, 0x61, 0x6c,
	0x65, 0x72, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x61, 0x6c, 0x65,
	0x72, 0x74, 0x65, 0x72, 0x2e, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x52, 0x06, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x22, 0x56, 0x0a, 0x0c, 0x41, 0x6c, 0x65, 0x72,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x34, 0x0a, 0x0b, 0x61, 0x6c, 0x65, 0x72,
	0x74, 0x73, 0x62, 0x61, 0x74, 0x63, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e,
	0x61, 0x6c, 0x65, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x0b, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x62, 0x61, 0x74, 0x63, 0x68, 0x12, 0x10,
	0x0a, 0x03, 0x63, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x63, 0x69, 0x64,
	0x22, 0x23, 0x0a, 0x0d, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x65, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x73, 0x65, 0x65, 0x6e, 0x22, 0x2c, 0x0a, 0x12, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x67,
	0x69, 0x76, 0x65, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x67, 0x69, 0x76,
	0x65, 0x6d, 0x65, 0x22, 0x28, 0x0a, 0x14, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x41,
	0x6c, 0x65, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x63,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x63, 0x69, 0x64, 0x22, 0x25, 0x0a,
	0x0b, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0x2a, 0x30, 0x0a, 0x07, 0x55, 0x52, 0x47, 0x45, 0x4e, 0x43, 0x59, 0x12,
	0x07, 0x0a, 0x03, 0x4c, 0x41, 0x58, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x49, 0x4d, 0x4d, 0x45,
	0x44, 0x49, 0x41, 0x54, 0x45, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x49, 0x4d, 0x50, 0x4f, 0x52,
	0x54, 0x41, 0x4e, 0x54, 0x10, 0x01, 0x32, 0xde, 0x01, 0x0a, 0x11, 0x43, 0x61, 0x74, 0x65, 0x72,
	0x41, 0x6c, 0x65, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3d, 0x0a, 0x0a,
	0x43, 0x61, 0x74, 0x65, 0x72, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x12, 0x15, 0x2e, 0x61, 0x6c, 0x65,
	0x72, 0x74, 0x65, 0x72, 0x2e, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x16, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x41, 0x6c, 0x65, 0x72,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x44, 0x0a, 0x0d, 0x44,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x12, 0x1b, 0x2e, 0x61,
	0x6c, 0x65, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x61, 0x6c, 0x65, 0x72,
	0x74, 0x65, 0x72, 0x2e, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x22,
	0x00, 0x12, 0x44, 0x0a, 0x0d, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x41, 0x6c, 0x65,
	0x72, 0x74, 0x12, 0x1d, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x74,
	0x72, 0x69, 0x65, 0x76, 0x65, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x12, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x41, 0x6c, 0x65, 0x72,
	0x74, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x62, 0x67,
	0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_alert_proto_rawDescOnce sync.Once
	file_alert_proto_rawDescData = file_alert_proto_rawDesc
)

func file_alert_proto_rawDescGZIP() []byte {
	file_alert_proto_rawDescOnce.Do(func() {
		file_alert_proto_rawDescData = protoimpl.X.CompressGZIP(file_alert_proto_rawDescData)
	})
	return file_alert_proto_rawDescData
}

var file_alert_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_alert_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_alert_proto_goTypes = []interface{}{
	(URGENCY)(0),                  // 0: alerter.URGENCY
	(*AlertDetail)(nil),           // 1: alerter.AlertDetail
	(*AlertList)(nil),             // 2: alerter.AlertList
	(*AlertRequest)(nil),          // 3: alerter.AlertRequest
	(*AlertResponse)(nil),         // 4: alerter.AlertResponse
	(*DescriptionRequest)(nil),    // 5: alerter.DescriptionRequest
	(*RetrieveAlertRequest)(nil),  // 6: alerter.RetrieveAlertRequest
	(*AlertSchema)(nil),           // 7: alerter.AlertSchema
	(*timestamppb.Timestamp)(nil), // 8: google.protobuf.Timestamp
	(*K8SInstance)(nil),           // 9: alerter.K8sInstance
}
var file_alert_proto_depIdxs = []int32{
	8, // 0: alerter.AlertDetail.alerttime:type_name -> google.protobuf.Timestamp
	0, // 1: alerter.AlertDetail.stallcap:type_name -> alerter.URGENCY
	9, // 2: alerter.AlertDetail.k8sorigin:type_name -> alerter.K8sInstance
	1, // 3: alerter.AlertList.alerts:type_name -> alerter.AlertDetail
	2, // 4: alerter.AlertRequest.alertsbatch:type_name -> alerter.AlertList
	3, // 5: alerter.CaterAlertRequest.CaterAlert:input_type -> alerter.AlertRequest
	5, // 6: alerter.CaterAlertRequest.DescribeAlert:input_type -> alerter.DescriptionRequest
	6, // 7: alerter.CaterAlertRequest.RetrieveAlert:input_type -> alerter.RetrieveAlertRequest
	4, // 8: alerter.CaterAlertRequest.CaterAlert:output_type -> alerter.AlertResponse
	7, // 9: alerter.CaterAlertRequest.DescribeAlert:output_type -> alerter.AlertSchema
	2, // 10: alerter.CaterAlertRequest.RetrieveAlert:output_type -> alerter.AlertList
	8, // [8:11] is the sub-list for method output_type
	5, // [5:8] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_alert_proto_init() }
func file_alert_proto_init() {
	if File_alert_proto != nil {
		return
	}
	file_instance_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_alert_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlertDetail); i {
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
		file_alert_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlertList); i {
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
		file_alert_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlertRequest); i {
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
		file_alert_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlertResponse); i {
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
		file_alert_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescriptionRequest); i {
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
		file_alert_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RetrieveAlertRequest); i {
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
		file_alert_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlertSchema); i {
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
			RawDescriptor: file_alert_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_alert_proto_goTypes,
		DependencyIndexes: file_alert_proto_depIdxs,
		EnumInfos:         file_alert_proto_enumTypes,
		MessageInfos:      file_alert_proto_msgTypes,
	}.Build()
	File_alert_proto = out.File
	file_alert_proto_rawDesc = nil
	file_alert_proto_goTypes = nil
	file_alert_proto_depIdxs = nil
}
