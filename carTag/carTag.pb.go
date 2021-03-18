// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.15.5
// source: carTag/carTag.proto

package carTag

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

// 新增Tagdto
type CarTagDto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 标签id，如果为新增，则此字段为0
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	//tag名称
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// tag类型id
	TypeId uint32 `protobuf:"varint,3,opt,name=type_id,json=typeId,proto3" json:"type_id,omitempty"`
	// tag是否可用
	Status uint32 `protobuf:"varint,4,opt,name=status,proto3" json:"status,omitempty"`
	//tag新增时间
	CreateTime uint64 `protobuf:"varint,5,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// tag的更新时间，如果为新增，或更新，则此字段无效
	UpdateTime uint64 `protobuf:"varint,6,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	TypeName   string `protobuf:"bytes,7,opt,name=type_name,json=typeName,proto3" json:"type_name,omitempty"`
	UpdateBy   string `protobuf:"bytes,8,opt,name=update_by,json=updateBy,proto3" json:"update_by,omitempty"`
	TagType    uint32 `protobuf:"varint,9,opt,name=tag_type,json=tagType,proto3" json:"tag_type,omitempty"`
}

func (x *CarTagDto) Reset() {
	*x = CarTagDto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_carTag_carTag_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CarTagDto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CarTagDto) ProtoMessage() {}

func (x *CarTagDto) ProtoReflect() protoreflect.Message {
	mi := &file_carTag_carTag_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CarTagDto.ProtoReflect.Descriptor instead.
func (*CarTagDto) Descriptor() ([]byte, []int) {
	return file_carTag_carTag_proto_rawDescGZIP(), []int{0}
}

func (x *CarTagDto) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CarTagDto) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CarTagDto) GetTypeId() uint32 {
	if x != nil {
		return x.TypeId
	}
	return 0
}

func (x *CarTagDto) GetStatus() uint32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *CarTagDto) GetCreateTime() uint64 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

func (x *CarTagDto) GetUpdateTime() uint64 {
	if x != nil {
		return x.UpdateTime
	}
	return 0
}

func (x *CarTagDto) GetTypeName() string {
	if x != nil {
		return x.TypeName
	}
	return ""
}

func (x *CarTagDto) GetUpdateBy() string {
	if x != nil {
		return x.UpdateBy
	}
	return ""
}

func (x *CarTagDto) GetTagType() uint32 {
	if x != nil {
		return x.TagType
	}
	return 0
}

// Tag搜索条件
type CarTagConditionDto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// tag的id
	Id       uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	TypeId   uint32 `protobuf:"varint,3,opt,name=type_id,json=typeId,proto3" json:"type_id,omitempty"`
	TypeName string `protobuf:"bytes,4,opt,name=type_name,json=typeName,proto3" json:"type_name,omitempty"`
	Page     uint32 `protobuf:"varint,5,opt,name=page,proto3" json:"page,omitempty"`
	Limit    uint32 `protobuf:"varint,6,opt,name=limit,proto3" json:"limit,omitempty"`
	TagType  uint32 `protobuf:"varint,7,opt,name=tag_type,json=tagType,proto3" json:"tag_type,omitempty"`
	Status   uint32 `protobuf:"varint,8,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *CarTagConditionDto) Reset() {
	*x = CarTagConditionDto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_carTag_carTag_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CarTagConditionDto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CarTagConditionDto) ProtoMessage() {}

func (x *CarTagConditionDto) ProtoReflect() protoreflect.Message {
	mi := &file_carTag_carTag_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CarTagConditionDto.ProtoReflect.Descriptor instead.
func (*CarTagConditionDto) Descriptor() ([]byte, []int) {
	return file_carTag_carTag_proto_rawDescGZIP(), []int{1}
}

func (x *CarTagConditionDto) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CarTagConditionDto) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CarTagConditionDto) GetTypeId() uint32 {
	if x != nil {
		return x.TypeId
	}
	return 0
}

func (x *CarTagConditionDto) GetTypeName() string {
	if x != nil {
		return x.TypeName
	}
	return ""
}

func (x *CarTagConditionDto) GetPage() uint32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *CarTagConditionDto) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *CarTagConditionDto) GetTagType() uint32 {
	if x != nil {
		return x.TagType
	}
	return 0
}

func (x *CarTagConditionDto) GetStatus() uint32 {
	if x != nil {
		return x.Status
	}
	return 0
}

//新建加的   car_tag 枚举
type CarTagsType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CarTagsType) Reset() {
	*x = CarTagsType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_carTag_carTag_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CarTagsType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CarTagsType) ProtoMessage() {}

func (x *CarTagsType) ProtoReflect() protoreflect.Message {
	mi := &file_carTag_carTag_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CarTagsType.ProtoReflect.Descriptor instead.
func (*CarTagsType) Descriptor() ([]byte, []int) {
	return file_carTag_carTag_proto_rawDescGZIP(), []int{2}
}

func (x *CarTagsType) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CarTagsType) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CarTagsTypes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CarTagsType []*CarTagsType `protobuf:"bytes,1,rep,name=carTagsType,proto3" json:"carTagsType,omitempty"`
}

func (x *CarTagsTypes) Reset() {
	*x = CarTagsTypes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_carTag_carTag_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CarTagsTypes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CarTagsTypes) ProtoMessage() {}

func (x *CarTagsTypes) ProtoReflect() protoreflect.Message {
	mi := &file_carTag_carTag_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CarTagsTypes.ProtoReflect.Descriptor instead.
func (*CarTagsTypes) Descriptor() ([]byte, []int) {
	return file_carTag_carTag_proto_rawDescGZIP(), []int{3}
}

func (x *CarTagsTypes) GetCarTagsType() []*CarTagsType {
	if x != nil {
		return x.CarTagsType
	}
	return nil
}

var File_carTag_carTag_proto protoreflect.FileDescriptor

var file_carTag_carTag_proto_rawDesc = []byte{
	0x0a, 0x13, 0x63, 0x61, 0x72, 0x54, 0x61, 0x67, 0x2f, 0x63, 0x61, 0x72, 0x54, 0x61, 0x67, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x63, 0x61, 0x72, 0x1a, 0x39, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x61, 0x72, 0x67, 0x6f, 0x64, 0x2d, 0x62, 0x6a,
	0x2f, 0x62, 0x32, 0x63, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf7, 0x01, 0x0a, 0x09, 0x43, 0x61, 0x72, 0x54, 0x61, 0x67,
	0x44, 0x74, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x79, 0x70, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x74, 0x79, 0x70, 0x65, 0x49, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x79,
	0x70, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74,
	0x79, 0x70, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x5f, 0x62, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x42, 0x79, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x61, 0x67, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x74, 0x61, 0x67, 0x54, 0x79, 0x70, 0x65, 0x22,
	0xcb, 0x01, 0x0a, 0x12, 0x43, 0x61, 0x72, 0x54, 0x61, 0x67, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x44, 0x74, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x79,
	0x70, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x74, 0x79, 0x70,
	0x65, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x79, 0x70, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04,
	0x70, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x61,
	0x67, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x74, 0x61,
	0x67, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x31, 0x0a,
	0x0b, 0x43, 0x61, 0x72, 0x54, 0x61, 0x67, 0x73, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x22, 0x42, 0x0a, 0x0c, 0x43, 0x61, 0x72, 0x54, 0x61, 0x67, 0x73, 0x54, 0x79, 0x70, 0x65, 0x73,
	0x12, 0x32, 0x0a, 0x0b, 0x63, 0x61, 0x72, 0x54, 0x61, 0x67, 0x73, 0x54, 0x79, 0x70, 0x65, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63, 0x61, 0x72, 0x2e, 0x43, 0x61, 0x72, 0x54,
	0x61, 0x67, 0x73, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b, 0x63, 0x61, 0x72, 0x54, 0x61, 0x67, 0x73,
	0x54, 0x79, 0x70, 0x65, 0x32, 0xff, 0x01, 0x0a, 0x06, 0x43, 0x61, 0x72, 0x54, 0x61, 0x67, 0x12,
	0x29, 0x0a, 0x03, 0x41, 0x64, 0x64, 0x12, 0x0e, 0x2e, 0x63, 0x61, 0x72, 0x2e, 0x43, 0x61, 0x72,
	0x54, 0x61, 0x67, 0x44, 0x74, 0x6f, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x35, 0x0a, 0x06, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x12, 0x17, 0x2e, 0x63, 0x61, 0x72, 0x2e, 0x43, 0x61, 0x72, 0x54, 0x61,
	0x67, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x74, 0x6f, 0x1a, 0x10, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x2c, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x0e, 0x2e, 0x63, 0x61,
	0x72, 0x2e, 0x43, 0x61, 0x72, 0x54, 0x61, 0x67, 0x44, 0x74, 0x6f, 0x1a, 0x10, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x33, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x17, 0x2e, 0x63, 0x61, 0x72, 0x2e, 0x43, 0x61,
	0x72, 0x54, 0x61, 0x67, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x74, 0x6f,
	0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x30, 0x0a, 0x08, 0x54, 0x61, 0x67, 0x73, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x10, 0x2e, 0x63, 0x61, 0x72, 0x2e, 0x43, 0x61, 0x72, 0x54, 0x61, 0x67, 0x73, 0x54, 0x79,
	0x70, 0x65, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x61, 0x72, 0x67, 0x6f, 0x64, 0x2d, 0x62, 0x6a, 0x2f, 0x62,
	0x32, 0x63, 0x2d, 0x63, 0x61, 0x72, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x61, 0x72, 0x54, 0x61,
	0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_carTag_carTag_proto_rawDescOnce sync.Once
	file_carTag_carTag_proto_rawDescData = file_carTag_carTag_proto_rawDesc
)

func file_carTag_carTag_proto_rawDescGZIP() []byte {
	file_carTag_carTag_proto_rawDescOnce.Do(func() {
		file_carTag_carTag_proto_rawDescData = protoimpl.X.CompressGZIP(file_carTag_carTag_proto_rawDescData)
	})
	return file_carTag_carTag_proto_rawDescData
}

var file_carTag_carTag_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_carTag_carTag_proto_goTypes = []interface{}{
	(*CarTagDto)(nil),          // 0: car.CarTagDto
	(*CarTagConditionDto)(nil), // 1: car.CarTagConditionDto
	(*CarTagsType)(nil),        // 2: car.CarTagsType
	(*CarTagsTypes)(nil),       // 3: car.CarTagsTypes
	(*common.Response)(nil),    // 4: common.Response
}
var file_carTag_carTag_proto_depIdxs = []int32{
	2, // 0: car.CarTagsTypes.carTagsType:type_name -> car.CarTagsType
	0, // 1: car.CarTag.Add:input_type -> car.CarTagDto
	1, // 2: car.CarTag.Delete:input_type -> car.CarTagConditionDto
	0, // 3: car.CarTag.Update:input_type -> car.CarTagDto
	1, // 4: car.CarTag.List:input_type -> car.CarTagConditionDto
	2, // 5: car.CarTag.TagsType:input_type -> car.CarTagsType
	4, // 6: car.CarTag.Add:output_type -> common.Response
	4, // 7: car.CarTag.Delete:output_type -> common.Response
	4, // 8: car.CarTag.Update:output_type -> common.Response
	4, // 9: car.CarTag.List:output_type -> common.Response
	4, // 10: car.CarTag.TagsType:output_type -> common.Response
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_carTag_carTag_proto_init() }
func file_carTag_carTag_proto_init() {
	if File_carTag_carTag_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_carTag_carTag_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CarTagDto); i {
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
		file_carTag_carTag_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CarTagConditionDto); i {
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
		file_carTag_carTag_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CarTagsType); i {
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
		file_carTag_carTag_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CarTagsTypes); i {
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
			RawDescriptor: file_carTag_carTag_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_carTag_carTag_proto_goTypes,
		DependencyIndexes: file_carTag_carTag_proto_depIdxs,
		MessageInfos:      file_carTag_carTag_proto_msgTypes,
	}.Build()
	File_carTag_carTag_proto = out.File
	file_carTag_carTag_proto_rawDesc = nil
	file_carTag_carTag_proto_goTypes = nil
	file_carTag_carTag_proto_depIdxs = nil
}
