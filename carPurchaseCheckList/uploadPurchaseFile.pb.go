// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.17.0
// source: carPurchaseCheckList/uploadPurchaseFile.proto

package CarPurCheckList

import (
	common "github.com/cargod-bj/b2c-proto-common/common"
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type PurchaseFileDto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//purchasefile urls
	UrlsInfo []*CarPurchaseFileDto `protobuf:"bytes,1,rep,name=urls_info,json=urlsInfo,proto3" json:"urls_info,omitempty"`
}

func (x *PurchaseFileDto) Reset() {
	*x = PurchaseFileDto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_carPurchaseCheckList_uploadPurchaseFile_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PurchaseFileDto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PurchaseFileDto) ProtoMessage() {}

func (x *PurchaseFileDto) ProtoReflect() protoreflect.Message {
	mi := &file_carPurchaseCheckList_uploadPurchaseFile_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PurchaseFileDto.ProtoReflect.Descriptor instead.
func (*PurchaseFileDto) Descriptor() ([]byte, []int) {
	return file_carPurchaseCheckList_uploadPurchaseFile_proto_rawDescGZIP(), []int{0}
}

func (x *PurchaseFileDto) GetUrlsInfo() []*CarPurchaseFileDto {
	if x != nil {
		return x.UrlsInfo
	}
	return nil
}

type CarPurchaseFileDto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Type uint64 `protobuf:"varint,2,opt,name=type,proto3" json:"type,omitempty"`
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Url  string `protobuf:"bytes,4,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *CarPurchaseFileDto) Reset() {
	*x = CarPurchaseFileDto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_carPurchaseCheckList_uploadPurchaseFile_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CarPurchaseFileDto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CarPurchaseFileDto) ProtoMessage() {}

func (x *CarPurchaseFileDto) ProtoReflect() protoreflect.Message {
	mi := &file_carPurchaseCheckList_uploadPurchaseFile_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CarPurchaseFileDto.ProtoReflect.Descriptor instead.
func (*CarPurchaseFileDto) Descriptor() ([]byte, []int) {
	return file_carPurchaseCheckList_uploadPurchaseFile_proto_rawDescGZIP(), []int{1}
}

func (x *CarPurchaseFileDto) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CarPurchaseFileDto) GetType() uint64 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *CarPurchaseFileDto) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CarPurchaseFileDto) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

//create checklist request params
type CreateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PurchaseId uint64 `protobuf:"varint,1,opt,name=purchaseId,proto3" json:"purchaseId,omitempty"`
	Name       string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Code       string `protobuf:"bytes,3,opt,name=code,proto3" json:"code,omitempty"`
	Required   uint32 `protobuf:"varint,4,opt,name=required,proto3" json:"required,omitempty"`
	Sequence   uint32 `protobuf:"varint,5,opt,name=sequence,proto3" json:"sequence,omitempty"`
	Status     uint32 `protobuf:"varint,6,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *CreateReq) Reset() {
	*x = CreateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_carPurchaseCheckList_uploadPurchaseFile_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateReq) ProtoMessage() {}

func (x *CreateReq) ProtoReflect() protoreflect.Message {
	mi := &file_carPurchaseCheckList_uploadPurchaseFile_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateReq.ProtoReflect.Descriptor instead.
func (*CreateReq) Descriptor() ([]byte, []int) {
	return file_carPurchaseCheckList_uploadPurchaseFile_proto_rawDescGZIP(), []int{2}
}

func (x *CreateReq) GetPurchaseId() uint64 {
	if x != nil {
		return x.PurchaseId
	}
	return 0
}

func (x *CreateReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateReq) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *CreateReq) GetRequired() uint32 {
	if x != nil {
		return x.Required
	}
	return 0
}

func (x *CreateReq) GetSequence() uint32 {
	if x != nil {
		return x.Sequence
	}
	return 0
}

func (x *CreateReq) GetStatus() uint32 {
	if x != nil {
		return x.Status
	}
	return 0
}

//batch create checklist
type BatchCreate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*CreateReq `protobuf:"bytes,1,rep,name=Infos,proto3" json:"Infos,omitempty"`
}

func (x *BatchCreate) Reset() {
	*x = BatchCreate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_carPurchaseCheckList_uploadPurchaseFile_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchCreate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchCreate) ProtoMessage() {}

func (x *BatchCreate) ProtoReflect() protoreflect.Message {
	mi := &file_carPurchaseCheckList_uploadPurchaseFile_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchCreate.ProtoReflect.Descriptor instead.
func (*BatchCreate) Descriptor() ([]byte, []int) {
	return file_carPurchaseCheckList_uploadPurchaseFile_proto_rawDescGZIP(), []int{3}
}

func (x *BatchCreate) GetInfos() []*CreateReq {
	if x != nil {
		return x.Infos
	}
	return nil
}

type UpdatesReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos      []*UpdateReq `protobuf:"bytes,1,rep,name=Infos,proto3" json:"Infos,omitempty"`
	PurchaseId uint64       `protobuf:"varint,2,opt,name=purchaseId,proto3" json:"purchaseId,omitempty"`
}

func (x *UpdatesReq) Reset() {
	*x = UpdatesReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_carPurchaseCheckList_uploadPurchaseFile_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatesReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatesReq) ProtoMessage() {}

func (x *UpdatesReq) ProtoReflect() protoreflect.Message {
	mi := &file_carPurchaseCheckList_uploadPurchaseFile_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatesReq.ProtoReflect.Descriptor instead.
func (*UpdatesReq) Descriptor() ([]byte, []int) {
	return file_carPurchaseCheckList_uploadPurchaseFile_proto_rawDescGZIP(), []int{4}
}

func (x *UpdatesReq) GetInfos() []*UpdateReq {
	if x != nil {
		return x.Infos
	}
	return nil
}

func (x *UpdatesReq) GetPurchaseId() uint64 {
	if x != nil {
		return x.PurchaseId
	}
	return 0
}

//update checklist request params
type UpdateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name     string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Code     string `protobuf:"bytes,4,opt,name=code,proto3" json:"code,omitempty"`
	Required uint32 `protobuf:"varint,5,opt,name=required,proto3" json:"required,omitempty"`
	Sequence uint32 `protobuf:"varint,6,opt,name=sequence,proto3" json:"sequence,omitempty"`
	Status   uint32 `protobuf:"varint,7,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *UpdateReq) Reset() {
	*x = UpdateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_carPurchaseCheckList_uploadPurchaseFile_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateReq) ProtoMessage() {}

func (x *UpdateReq) ProtoReflect() protoreflect.Message {
	mi := &file_carPurchaseCheckList_uploadPurchaseFile_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateReq.ProtoReflect.Descriptor instead.
func (*UpdateReq) Descriptor() ([]byte, []int) {
	return file_carPurchaseCheckList_uploadPurchaseFile_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateReq) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateReq) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *UpdateReq) GetRequired() uint32 {
	if x != nil {
		return x.Required
	}
	return 0
}

func (x *UpdateReq) GetSequence() uint32 {
	if x != nil {
		return x.Sequence
	}
	return 0
}

func (x *UpdateReq) GetStatus() uint32 {
	if x != nil {
		return x.Status
	}
	return 0
}

//carPurchaseId query ckecklist req params
type CarPurchaseIdReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CarPurchaseId uint64 `protobuf:"varint,1,opt,name=carPurchaseId,proto3" json:"carPurchaseId,omitempty"`
}

func (x *CarPurchaseIdReq) Reset() {
	*x = CarPurchaseIdReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_carPurchaseCheckList_uploadPurchaseFile_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CarPurchaseIdReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CarPurchaseIdReq) ProtoMessage() {}

func (x *CarPurchaseIdReq) ProtoReflect() protoreflect.Message {
	mi := &file_carPurchaseCheckList_uploadPurchaseFile_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CarPurchaseIdReq.ProtoReflect.Descriptor instead.
func (*CarPurchaseIdReq) Descriptor() ([]byte, []int) {
	return file_carPurchaseCheckList_uploadPurchaseFile_proto_rawDescGZIP(), []int{6}
}

func (x *CarPurchaseIdReq) GetCarPurchaseId() uint64 {
	if x != nil {
		return x.CarPurchaseId
	}
	return 0
}

//IdReq update or delete checklist req params
type IdReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CarPurchaseId uint64 `protobuf:"varint,1,opt,name=carPurchaseId,proto3" json:"carPurchaseId,omitempty"`
}

func (x *IdReq) Reset() {
	*x = IdReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_carPurchaseCheckList_uploadPurchaseFile_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IdReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdReq) ProtoMessage() {}

func (x *IdReq) ProtoReflect() protoreflect.Message {
	mi := &file_carPurchaseCheckList_uploadPurchaseFile_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdReq.ProtoReflect.Descriptor instead.
func (*IdReq) Descriptor() ([]byte, []int) {
	return file_carPurchaseCheckList_uploadPurchaseFile_proto_rawDescGZIP(), []int{7}
}

func (x *IdReq) GetCarPurchaseId() uint64 {
	if x != nil {
		return x.CarPurchaseId
	}
	return 0
}

type ChecklistDto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	PurchaseId uint64 `protobuf:"varint,2,opt,name=purchaseId,proto3" json:"purchaseId,omitempty"`
	Name       string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Code       string `protobuf:"bytes,4,opt,name=code,proto3" json:"code,omitempty"`
	Required   uint32 `protobuf:"varint,5,opt,name=required,proto3" json:"required,omitempty"`
	Sequence   uint32 `protobuf:"varint,6,opt,name=sequence,proto3" json:"sequence,omitempty"`
	Status     uint32 `protobuf:"varint,7,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *ChecklistDto) Reset() {
	*x = ChecklistDto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_carPurchaseCheckList_uploadPurchaseFile_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChecklistDto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChecklistDto) ProtoMessage() {}

func (x *ChecklistDto) ProtoReflect() protoreflect.Message {
	mi := &file_carPurchaseCheckList_uploadPurchaseFile_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChecklistDto.ProtoReflect.Descriptor instead.
func (*ChecklistDto) Descriptor() ([]byte, []int) {
	return file_carPurchaseCheckList_uploadPurchaseFile_proto_rawDescGZIP(), []int{8}
}

func (x *ChecklistDto) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ChecklistDto) GetPurchaseId() uint64 {
	if x != nil {
		return x.PurchaseId
	}
	return 0
}

func (x *ChecklistDto) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ChecklistDto) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *ChecklistDto) GetRequired() uint32 {
	if x != nil {
		return x.Required
	}
	return 0
}

func (x *ChecklistDto) GetSequence() uint32 {
	if x != nil {
		return x.Sequence
	}
	return 0
}

func (x *ChecklistDto) GetStatus() uint32 {
	if x != nil {
		return x.Status
	}
	return 0
}

var File_carPurchaseCheckList_uploadPurchaseFile_proto protoreflect.FileDescriptor

var file_carPurchaseCheckList_uploadPurchaseFile_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x63, 0x61, 0x72, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x2f, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x50, 0x75, 0x72,
	0x63, 0x68, 0x61, 0x73, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0f, 0x43, 0x61, 0x72, 0x50, 0x75, 0x72, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x4c, 0x69, 0x73, 0x74,
	0x1a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x61, 0x72,
	0x67, 0x6f, 0x64, 0x2d, 0x62, 0x6a, 0x2f, 0x62, 0x32, 0x63, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2d, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x53, 0x0a, 0x0f, 0x50,
	0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x44, 0x74, 0x6f, 0x12, 0x40,
	0x0a, 0x09, 0x75, 0x72, 0x6c, 0x73, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x23, 0x2e, 0x43, 0x61, 0x72, 0x50, 0x75, 0x72, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x4c,
	0x69, 0x73, 0x74, 0x2e, 0x43, 0x61, 0x72, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x46,
	0x69, 0x6c, 0x65, 0x44, 0x74, 0x6f, 0x52, 0x08, 0x75, 0x72, 0x6c, 0x73, 0x49, 0x6e, 0x66, 0x6f,
	0x22, 0x5e, 0x0a, 0x12, 0x43, 0x61, 0x72, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x46,
	0x69, 0x6c, 0x65, 0x44, 0x74, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c,
	0x22, 0xa3, 0x01, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x12, 0x1e,
	0x0a, 0x0a, 0x70, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x0a, 0x70, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x49, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72,
	0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72,
	0x65, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x3f, 0x0a, 0x0b, 0x42, 0x61, 0x74, 0x63, 0x68, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x30, 0x0a, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x43, 0x61, 0x72, 0x50, 0x75, 0x72, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x52, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x22, 0x5e, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x73, 0x52, 0x65, 0x71, 0x12, 0x30, 0x0a, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x43, 0x61, 0x72, 0x50, 0x75, 0x72, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x52, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x75, 0x72, 0x63, 0x68,
	0x61, 0x73, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x70, 0x75, 0x72,
	0x63, 0x68, 0x61, 0x73, 0x65, 0x49, 0x64, 0x22, 0x93, 0x01, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x08, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x71,
	0x75, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x73, 0x65, 0x71,
	0x75, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x38, 0x0a,
	0x10, 0x43, 0x61, 0x72, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x49, 0x64, 0x52, 0x65,
	0x71, 0x12, 0x24, 0x0a, 0x0d, 0x63, 0x61, 0x72, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0d, 0x63, 0x61, 0x72, 0x50, 0x75, 0x72,
	0x63, 0x68, 0x61, 0x73, 0x65, 0x49, 0x64, 0x22, 0x2d, 0x0a, 0x05, 0x49, 0x64, 0x52, 0x65, 0x71,
	0x12, 0x24, 0x0a, 0x0d, 0x63, 0x61, 0x72, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0d, 0x63, 0x61, 0x72, 0x50, 0x75, 0x72, 0x63,
	0x68, 0x61, 0x73, 0x65, 0x49, 0x64, 0x22, 0xb6, 0x01, 0x0a, 0x0c, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x6c, 0x69, 0x73, 0x74, 0x44, 0x74, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x75, 0x72, 0x63, 0x68,
	0x61, 0x73, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x70, 0x75, 0x72,
	0x63, 0x68, 0x61, 0x73, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x08, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x73,
	0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x73,
	0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32,
	0xfc, 0x03, 0x0a, 0x10, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x42, 0x0a, 0x0a, 0x53, 0x61, 0x76, 0x65, 0x41, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x12, 0x20, 0x2e, 0x43, 0x61, 0x72, 0x50, 0x75, 0x72, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x4c, 0x69, 0x73, 0x74, 0x2e, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x46, 0x69, 0x6c,
	0x65, 0x44, 0x74, 0x6f, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x11, 0x47, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x16, 0x2e,
	0x43, 0x61, 0x72, 0x50, 0x75, 0x72, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x2e,
	0x49, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x07, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x73, 0x12, 0x1b, 0x2e, 0x43, 0x61, 0x72, 0x50, 0x75, 0x72, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65,
	0x71, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12,
	0x1a, 0x2e, 0x43, 0x61, 0x72, 0x50, 0x75, 0x72, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x4c, 0x69, 0x73,
	0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x10, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x3b, 0x0a, 0x07, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x73, 0x12, 0x1c, 0x2e, 0x43, 0x61, 0x72,
	0x50, 0x75, 0x72, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x2e, 0x42, 0x61, 0x74,
	0x63, 0x68, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x07,
	0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x21, 0x2e, 0x43, 0x61, 0x72, 0x50, 0x75, 0x72,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x2e, 0x43, 0x61, 0x72, 0x50, 0x75, 0x72,
	0x63, 0x68, 0x61, 0x73, 0x65, 0x49, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x38,
	0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1a, 0x2e, 0x43, 0x61, 0x72, 0x50, 0x75,
	0x72, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x12, 0x16, 0x2e, 0x43, 0x61, 0x72, 0x50, 0x75, 0x72, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x4c, 0x69, 0x73, 0x74, 0x2e, 0x49, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x32,
	0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x61, 0x72,
	0x67, 0x6f, 0x64, 0x2d, 0x62, 0x6a, 0x2f, 0x62, 0x32, 0x63, 0x2d, 0x63, 0x61, 0x72, 0x2d, 0x61,
	0x70, 0x69, 0x2f, 0x43, 0x61, 0x72, 0x50, 0x75, 0x72, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x4c, 0x69,
	0x73, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_carPurchaseCheckList_uploadPurchaseFile_proto_rawDescOnce sync.Once
	file_carPurchaseCheckList_uploadPurchaseFile_proto_rawDescData = file_carPurchaseCheckList_uploadPurchaseFile_proto_rawDesc
)

func file_carPurchaseCheckList_uploadPurchaseFile_proto_rawDescGZIP() []byte {
	file_carPurchaseCheckList_uploadPurchaseFile_proto_rawDescOnce.Do(func() {
		file_carPurchaseCheckList_uploadPurchaseFile_proto_rawDescData = protoimpl.X.CompressGZIP(file_carPurchaseCheckList_uploadPurchaseFile_proto_rawDescData)
	})
	return file_carPurchaseCheckList_uploadPurchaseFile_proto_rawDescData
}

var file_carPurchaseCheckList_uploadPurchaseFile_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_carPurchaseCheckList_uploadPurchaseFile_proto_goTypes = []interface{}{
	(*PurchaseFileDto)(nil),    // 0: CarPurCheckList.PurchaseFileDto
	(*CarPurchaseFileDto)(nil), // 1: CarPurCheckList.CarPurchaseFileDto
	(*CreateReq)(nil),          // 2: CarPurCheckList.CreateReq
	(*BatchCreate)(nil),        // 3: CarPurCheckList.BatchCreate
	(*UpdatesReq)(nil),         // 4: CarPurCheckList.UpdatesReq
	(*UpdateReq)(nil),          // 5: CarPurCheckList.UpdateReq
	(*CarPurchaseIdReq)(nil),   // 6: CarPurCheckList.CarPurchaseIdReq
	(*IdReq)(nil),              // 7: CarPurCheckList.IdReq
	(*ChecklistDto)(nil),       // 8: CarPurCheckList.ChecklistDto
	(*common.Response)(nil),    // 9: common.Response
}
var file_carPurchaseCheckList_uploadPurchaseFile_proto_depIdxs = []int32{
	1,  // 0: CarPurCheckList.PurchaseFileDto.urls_info:type_name -> CarPurCheckList.CarPurchaseFileDto
	2,  // 1: CarPurCheckList.BatchCreate.Infos:type_name -> CarPurCheckList.CreateReq
	5,  // 2: CarPurCheckList.UpdatesReq.Infos:type_name -> CarPurCheckList.UpdateReq
	0,  // 3: CarPurCheckList.CheckListService.SaveAdress:input_type -> CarPurCheckList.PurchaseFileDto
	7,  // 4: CarPurCheckList.CheckListService.GenerateCheckList:input_type -> CarPurCheckList.IdReq
	4,  // 5: CarPurCheckList.CheckListService.Updates:input_type -> CarPurCheckList.UpdatesReq
	2,  // 6: CarPurCheckList.CheckListService.Create:input_type -> CarPurCheckList.CreateReq
	3,  // 7: CarPurCheckList.CheckListService.Creates:input_type -> CarPurCheckList.BatchCreate
	6,  // 8: CarPurCheckList.CheckListService.GetList:input_type -> CarPurCheckList.CarPurchaseIdReq
	5,  // 9: CarPurCheckList.CheckListService.Update:input_type -> CarPurCheckList.UpdateReq
	7,  // 10: CarPurCheckList.CheckListService.Delete:input_type -> CarPurCheckList.IdReq
	9,  // 11: CarPurCheckList.CheckListService.SaveAdress:output_type -> common.Response
	9,  // 12: CarPurCheckList.CheckListService.GenerateCheckList:output_type -> common.Response
	9,  // 13: CarPurCheckList.CheckListService.Updates:output_type -> common.Response
	9,  // 14: CarPurCheckList.CheckListService.Create:output_type -> common.Response
	9,  // 15: CarPurCheckList.CheckListService.Creates:output_type -> common.Response
	9,  // 16: CarPurCheckList.CheckListService.GetList:output_type -> common.Response
	9,  // 17: CarPurCheckList.CheckListService.Update:output_type -> common.Response
	9,  // 18: CarPurCheckList.CheckListService.Delete:output_type -> common.Response
	11, // [11:19] is the sub-list for method output_type
	3,  // [3:11] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_carPurchaseCheckList_uploadPurchaseFile_proto_init() }
func file_carPurchaseCheckList_uploadPurchaseFile_proto_init() {
	if File_carPurchaseCheckList_uploadPurchaseFile_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_carPurchaseCheckList_uploadPurchaseFile_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PurchaseFileDto); i {
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
		file_carPurchaseCheckList_uploadPurchaseFile_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CarPurchaseFileDto); i {
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
		file_carPurchaseCheckList_uploadPurchaseFile_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateReq); i {
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
		file_carPurchaseCheckList_uploadPurchaseFile_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchCreate); i {
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
		file_carPurchaseCheckList_uploadPurchaseFile_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatesReq); i {
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
		file_carPurchaseCheckList_uploadPurchaseFile_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateReq); i {
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
		file_carPurchaseCheckList_uploadPurchaseFile_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CarPurchaseIdReq); i {
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
		file_carPurchaseCheckList_uploadPurchaseFile_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IdReq); i {
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
		file_carPurchaseCheckList_uploadPurchaseFile_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChecklistDto); i {
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
			RawDescriptor: file_carPurchaseCheckList_uploadPurchaseFile_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_carPurchaseCheckList_uploadPurchaseFile_proto_goTypes,
		DependencyIndexes: file_carPurchaseCheckList_uploadPurchaseFile_proto_depIdxs,
		MessageInfos:      file_carPurchaseCheckList_uploadPurchaseFile_proto_msgTypes,
	}.Build()
	File_carPurchaseCheckList_uploadPurchaseFile_proto = out.File
	file_carPurchaseCheckList_uploadPurchaseFile_proto_rawDesc = nil
	file_carPurchaseCheckList_uploadPurchaseFile_proto_goTypes = nil
	file_carPurchaseCheckList_uploadPurchaseFile_proto_depIdxs = nil
}
