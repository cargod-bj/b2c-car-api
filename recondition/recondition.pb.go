// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0-devel
// 	protoc        v3.12.1
// source: recondition/recondition.proto

package recondition

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

// 整备信息
type ReconditionDto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         uint64 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	CarId      uint64 `protobuf:"varint,2,opt,name=CarId,proto3" json:"CarId,omitempty"`
	StoreId    uint64 `protobuf:"varint,3,opt,name=StoreId,proto3" json:"StoreId,omitempty"`
	Status     uint32 `protobuf:"varint,4,opt,name=Status,proto3" json:"Status,omitempty"`
	StartTime  int64  `protobuf:"varint,5,opt,name=StartTime,proto3" json:"StartTime,omitempty"`
	EndTime    int64  `protobuf:"varint,6,opt,name=EndTime,proto3" json:"EndTime,omitempty"`
	CreateTime int64  `protobuf:"varint,7,opt,name=CreateTime,proto3" json:"CreateTime,omitempty"`
	UpdateTime int64  `protobuf:"varint,8,opt,name=UpdateTime,proto3" json:"UpdateTime,omitempty"`
}

func (x *ReconditionDto) Reset() {
	*x = ReconditionDto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_recondition_recondition_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReconditionDto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReconditionDto) ProtoMessage() {}

func (x *ReconditionDto) ProtoReflect() protoreflect.Message {
	mi := &file_recondition_recondition_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReconditionDto.ProtoReflect.Descriptor instead.
func (*ReconditionDto) Descriptor() ([]byte, []int) {
	return file_recondition_recondition_proto_rawDescGZIP(), []int{0}
}

func (x *ReconditionDto) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ReconditionDto) GetCarId() uint64 {
	if x != nil {
		return x.CarId
	}
	return 0
}

func (x *ReconditionDto) GetStoreId() uint64 {
	if x != nil {
		return x.StoreId
	}
	return 0
}

func (x *ReconditionDto) GetStatus() uint32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *ReconditionDto) GetStartTime() int64 {
	if x != nil {
		return x.StartTime
	}
	return 0
}

func (x *ReconditionDto) GetEndTime() int64 {
	if x != nil {
		return x.EndTime
	}
	return 0
}

func (x *ReconditionDto) GetCreateTime() int64 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

func (x *ReconditionDto) GetUpdateTime() int64 {
	if x != nil {
		return x.UpdateTime
	}
	return 0
}

// damage的顶层类别，每个顶层类别可能包含多个 子类别
type ReportDamageCategoryDto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 当前category的id
	Id uint64 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	// 当前category的Name
	Name uint64 `protobuf:"varint,2,opt,name=Name,proto3" json:"Name,omitempty"`
	// 当前category的子项，如果没有子项则为nil
	Children []*ReportDamageSubCategoryDto `protobuf:"bytes,3,rep,name=children,proto3" json:"children,omitempty"`
}

func (x *ReportDamageCategoryDto) Reset() {
	*x = ReportDamageCategoryDto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_recondition_recondition_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportDamageCategoryDto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportDamageCategoryDto) ProtoMessage() {}

func (x *ReportDamageCategoryDto) ProtoReflect() protoreflect.Message {
	mi := &file_recondition_recondition_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportDamageCategoryDto.ProtoReflect.Descriptor instead.
func (*ReportDamageCategoryDto) Descriptor() ([]byte, []int) {
	return file_recondition_recondition_proto_rawDescGZIP(), []int{1}
}

func (x *ReportDamageCategoryDto) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ReportDamageCategoryDto) GetName() uint64 {
	if x != nil {
		return x.Name
	}
	return 0
}

func (x *ReportDamageCategoryDto) GetChildren() []*ReportDamageSubCategoryDto {
	if x != nil {
		return x.Children
	}
	return nil
}

// damage的子类别，每个子类别可能包含多个 监测点
type ReportDamageSubCategoryDto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 当前category的id
	Id uint64 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	// 当前category的Name
	Name uint64 `protobuf:"varint,2,opt,name=Name,proto3" json:"Name,omitempty"`
	// 当前category的子项，如果没有子项则为nil
	Children []*ReportDamagePointDto `protobuf:"bytes,3,rep,name=children,proto3" json:"children,omitempty"`
}

func (x *ReportDamageSubCategoryDto) Reset() {
	*x = ReportDamageSubCategoryDto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_recondition_recondition_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportDamageSubCategoryDto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportDamageSubCategoryDto) ProtoMessage() {}

func (x *ReportDamageSubCategoryDto) ProtoReflect() protoreflect.Message {
	mi := &file_recondition_recondition_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportDamageSubCategoryDto.ProtoReflect.Descriptor instead.
func (*ReportDamageSubCategoryDto) Descriptor() ([]byte, []int) {
	return file_recondition_recondition_proto_rawDescGZIP(), []int{2}
}

func (x *ReportDamageSubCategoryDto) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ReportDamageSubCategoryDto) GetName() uint64 {
	if x != nil {
		return x.Name
	}
	return 0
}

func (x *ReportDamageSubCategoryDto) GetChildren() []*ReportDamagePointDto {
	if x != nil {
		return x.Children
	}
	return nil
}

// damage的监测点，每个监测点可能包含多个 tag点，多个 损坏照片
type ReportDamagePointDto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 当前point的id
	Id uint64 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	// 当前point的Name
	Name uint64 `protobuf:"varint,2,opt,name=Name,proto3" json:"Name,omitempty"`
	// 当前point上的损坏点列表
	Tags []string `protobuf:"bytes,3,rep,name=tags,proto3" json:"tags,omitempty"`
	// 当前point上的损坏点图片
	Photos []string `protobuf:"bytes,4,rep,name=photos,proto3" json:"photos,omitempty"`
}

func (x *ReportDamagePointDto) Reset() {
	*x = ReportDamagePointDto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_recondition_recondition_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportDamagePointDto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportDamagePointDto) ProtoMessage() {}

func (x *ReportDamagePointDto) ProtoReflect() protoreflect.Message {
	mi := &file_recondition_recondition_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportDamagePointDto.ProtoReflect.Descriptor instead.
func (*ReportDamagePointDto) Descriptor() ([]byte, []int) {
	return file_recondition_recondition_proto_rawDescGZIP(), []int{3}
}

func (x *ReportDamagePointDto) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ReportDamagePointDto) GetName() uint64 {
	if x != nil {
		return x.Name
	}
	return 0
}

func (x *ReportDamagePointDto) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *ReportDamagePointDto) GetPhotos() []string {
	if x != nil {
		return x.Photos
	}
	return nil
}

var File_recondition_recondition_proto protoreflect.FileDescriptor

var file_recondition_recondition_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x72, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x72, 0x65,
	0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0b, 0x72, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x39, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x61, 0x72, 0x67, 0x6f, 0x64, 0x2d,
	0x62, 0x6a, 0x2f, 0x62, 0x32, 0x63, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe0, 0x01, 0x0a, 0x0e, 0x52, 0x65, 0x63, 0x6f,
	0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x74, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x43, 0x61,
	0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x43, 0x61, 0x72, 0x49, 0x64,
	0x12, 0x18, 0x0a, 0x07, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x07, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x45, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x07, 0x45, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x82, 0x01, 0x0a, 0x17, 0x52,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x44, 0x61, 0x6d, 0x61, 0x67, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x44, 0x74, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x02, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x43, 0x0a, 0x08, 0x63, 0x68,
	0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x72,
	0x65, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x44, 0x61, 0x6d, 0x61, 0x67, 0x65, 0x53, 0x75, 0x62, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x44, 0x74, 0x6f, 0x52, 0x08, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x22,
	0x7f, 0x0a, 0x1a, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x44, 0x61, 0x6d, 0x61, 0x67, 0x65, 0x53,
	0x75, 0x62, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x44, 0x74, 0x6f, 0x12, 0x0e, 0x0a,
	0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x49, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x3d, 0x0a, 0x08, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x44, 0x61, 0x6d, 0x61, 0x67, 0x65, 0x50, 0x6f,
	0x69, 0x6e, 0x74, 0x44, 0x74, 0x6f, 0x52, 0x08, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e,
	0x22, 0x66, 0x0a, 0x14, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x44, 0x61, 0x6d, 0x61, 0x67, 0x65,
	0x50, 0x6f, 0x69, 0x6e, 0x74, 0x44, 0x74, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x61, 0x67, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73,
	0x12, 0x16, 0x0a, 0x06, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x06, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x73, 0x32, 0xd7, 0x02, 0x0a, 0x0b, 0x52, 0x65, 0x63,
	0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x36, 0x0a, 0x03, 0x41, 0x64, 0x64, 0x12,
	0x1b, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65,
	0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x74, 0x6f, 0x1a, 0x10, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x39, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x1b, 0x2e, 0x72, 0x65, 0x63,
	0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x44, 0x74, 0x6f, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x06, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1b, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x44,
	0x74, 0x6f, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x1b, 0x2e,
	0x72, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x63, 0x6f,
	0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x74, 0x6f, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x28,
	0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x50, 0x61, 0x67, 0x65, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x52,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x44, 0x61, 0x6d, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x49, 0x64, 0x44, 0x74, 0x6f, 0x1a, 0x10,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x63, 0x61, 0x72, 0x67, 0x6f, 0x64, 0x2d, 0x62, 0x6a, 0x2f, 0x62, 0x32, 0x63, 0x2d, 0x63,
	0x61, 0x72, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_recondition_recondition_proto_rawDescOnce sync.Once
	file_recondition_recondition_proto_rawDescData = file_recondition_recondition_proto_rawDesc
)

func file_recondition_recondition_proto_rawDescGZIP() []byte {
	file_recondition_recondition_proto_rawDescOnce.Do(func() {
		file_recondition_recondition_proto_rawDescData = protoimpl.X.CompressGZIP(file_recondition_recondition_proto_rawDescData)
	})
	return file_recondition_recondition_proto_rawDescData
}

var file_recondition_recondition_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_recondition_recondition_proto_goTypes = []interface{}{
	(*ReconditionDto)(nil),             // 0: recondition.ReconditionDto
	(*ReportDamageCategoryDto)(nil),    // 1: recondition.ReportDamageCategoryDto
	(*ReportDamageSubCategoryDto)(nil), // 2: recondition.ReportDamageSubCategoryDto
	(*ReportDamagePointDto)(nil),       // 3: recondition.ReportDamagePointDto
	(*common.Page)(nil),                // 4: common.Page
	(*common.IdDto)(nil),               // 5: common.IdDto
	(*common.Response)(nil),            // 6: common.Response
}
var file_recondition_recondition_proto_depIdxs = []int32{
	2, // 0: recondition.ReportDamageCategoryDto.children:type_name -> recondition.ReportDamageSubCategoryDto
	3, // 1: recondition.ReportDamageSubCategoryDto.children:type_name -> recondition.ReportDamagePointDto
	0, // 2: recondition.Recondition.Add:input_type -> recondition.ReconditionDto
	0, // 3: recondition.Recondition.Delete:input_type -> recondition.ReconditionDto
	0, // 4: recondition.Recondition.Update:input_type -> recondition.ReconditionDto
	0, // 5: recondition.Recondition.Get:input_type -> recondition.ReconditionDto
	4, // 6: recondition.Recondition.List:input_type -> common.Page
	5, // 7: recondition.Recondition.GetReportDamageInfo:input_type -> common.IdDto
	6, // 8: recondition.Recondition.Add:output_type -> common.Response
	6, // 9: recondition.Recondition.Delete:output_type -> common.Response
	6, // 10: recondition.Recondition.Update:output_type -> common.Response
	6, // 11: recondition.Recondition.Get:output_type -> common.Response
	6, // 12: recondition.Recondition.List:output_type -> common.Response
	6, // 13: recondition.Recondition.GetReportDamageInfo:output_type -> common.Response
	8, // [8:14] is the sub-list for method output_type
	2, // [2:8] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_recondition_recondition_proto_init() }
func file_recondition_recondition_proto_init() {
	if File_recondition_recondition_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_recondition_recondition_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReconditionDto); i {
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
		file_recondition_recondition_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReportDamageCategoryDto); i {
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
		file_recondition_recondition_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReportDamageSubCategoryDto); i {
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
		file_recondition_recondition_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReportDamagePointDto); i {
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
			RawDescriptor: file_recondition_recondition_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_recondition_recondition_proto_goTypes,
		DependencyIndexes: file_recondition_recondition_proto_depIdxs,
		MessageInfos:      file_recondition_recondition_proto_msgTypes,
	}.Build()
	File_recondition_recondition_proto = out.File
	file_recondition_recondition_proto_rawDesc = nil
	file_recondition_recondition_proto_goTypes = nil
	file_recondition_recondition_proto_depIdxs = nil
}
