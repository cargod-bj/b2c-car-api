// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0
// 	protoc        v3.9.0
// source: carEnumProto/carEnum.proto

package carEnumProto

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

// 枚举类型，支持："Color"、"BodyType"、"Transmission"、"Seat"、"RegistrationType"、"CarState"、"InventoryStatus"、"ReconditionPointType"等等
// 最新的支持枚举以car_utils.go中定义的为准
type EnumTypeDto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 枚举类型，支持："Color"、"BodyType"、"Transmission"、"Seat"、"RegistrationType"、"CarState"、"InventoryStatus"、"ReconditionPointType"等等
	// 最新的支持枚举以car_utils.go中定义的为准
	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	// 语言类型：en
	Local string `protobuf:"bytes,2,opt,name=local,proto3" json:"local,omitempty"`
	// 此参数和type互斥，查询多个分组，只有在GetCarEnums方法中可用
	Types string `protobuf:"bytes,3,opt,name=types,proto3" json:"types,omitempty"`
}

func (x *EnumTypeDto) Reset() {
	*x = EnumTypeDto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_carEnumProto_carEnum_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnumTypeDto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnumTypeDto) ProtoMessage() {}

func (x *EnumTypeDto) ProtoReflect() protoreflect.Message {
	mi := &file_carEnumProto_carEnum_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnumTypeDto.ProtoReflect.Descriptor instead.
func (*EnumTypeDto) Descriptor() ([]byte, []int) {
	return file_carEnumProto_carEnum_proto_rawDescGZIP(), []int{0}
}

func (x *EnumTypeDto) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *EnumTypeDto) GetLocal() string {
	if x != nil {
		return x.Local
	}
	return ""
}

func (x *EnumTypeDto) GetTypes() string {
	if x != nil {
		return x.Types
	}
	return ""
}

// 查询多个枚举类型，支持："Color"、"BodyType"、"Transmission"、"Seat"、"RegistrationType"、"CarState"、"InventoryStatus"、"ReconditionPointType"等等
// 最新的支持枚举以car_utils.go中定义的为准
type EnumTypesDto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 枚举类型，支持："Color"、"BodyType"、"Transmission"、"Seat"、"RegistrationType"、"CarState"、"InventoryStatus"、"ReconditionPointType"等等
	// 最新的支持枚举以car_utils.go中定义的为准
	Types []string `protobuf:"bytes,1,rep,name=types,proto3" json:"types,omitempty"`
	// 语言类型：en
	Local string `protobuf:"bytes,2,opt,name=local,proto3" json:"local,omitempty"`
}

func (x *EnumTypesDto) Reset() {
	*x = EnumTypesDto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_carEnumProto_carEnum_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnumTypesDto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnumTypesDto) ProtoMessage() {}

func (x *EnumTypesDto) ProtoReflect() protoreflect.Message {
	mi := &file_carEnumProto_carEnum_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnumTypesDto.ProtoReflect.Descriptor instead.
func (*EnumTypesDto) Descriptor() ([]byte, []int) {
	return file_carEnumProto_carEnum_proto_rawDescGZIP(), []int{1}
}

func (x *EnumTypesDto) GetTypes() []string {
	if x != nil {
		return x.Types
	}
	return nil
}

func (x *EnumTypesDto) GetLocal() string {
	if x != nil {
		return x.Local
	}
	return ""
}

// 多个查询
type EnumsDto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Enums map[string]*KVMap `protobuf:"bytes,1,rep,name=enums,proto3" json:"enums,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *EnumsDto) Reset() {
	*x = EnumsDto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_carEnumProto_carEnum_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnumsDto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnumsDto) ProtoMessage() {}

func (x *EnumsDto) ProtoReflect() protoreflect.Message {
	mi := &file_carEnumProto_carEnum_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnumsDto.ProtoReflect.Descriptor instead.
func (*EnumsDto) Descriptor() ([]byte, []int) {
	return file_carEnumProto_carEnum_proto_rawDescGZIP(), []int{2}
}

func (x *EnumsDto) GetEnums() map[string]*KVMap {
	if x != nil {
		return x.Enums
	}
	return nil
}

type KVMap struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Kvs map[string]uint32 `protobuf:"bytes,1,rep,name=kvs,proto3" json:"kvs,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (x *KVMap) Reset() {
	*x = KVMap{}
	if protoimpl.UnsafeEnabled {
		mi := &file_carEnumProto_carEnum_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KVMap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KVMap) ProtoMessage() {}

func (x *KVMap) ProtoReflect() protoreflect.Message {
	mi := &file_carEnumProto_carEnum_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KVMap.ProtoReflect.Descriptor instead.
func (*KVMap) Descriptor() ([]byte, []int) {
	return file_carEnumProto_carEnum_proto_rawDescGZIP(), []int{3}
}

func (x *KVMap) GetKvs() map[string]uint32 {
	if x != nil {
		return x.Kvs
	}
	return nil
}

// 根据EnumTypeDto获取出来的key value值
type KeyValueDto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// key值，如：Red
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// value值，如：1
	Value uint32 `protobuf:"varint,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *KeyValueDto) Reset() {
	*x = KeyValueDto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_carEnumProto_carEnum_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeyValueDto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyValueDto) ProtoMessage() {}

func (x *KeyValueDto) ProtoReflect() protoreflect.Message {
	mi := &file_carEnumProto_carEnum_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyValueDto.ProtoReflect.Descriptor instead.
func (*KeyValueDto) Descriptor() ([]byte, []int) {
	return file_carEnumProto_carEnum_proto_rawDescGZIP(), []int{4}
}

func (x *KeyValueDto) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *KeyValueDto) GetValue() uint32 {
	if x != nil {
		return x.Value
	}
	return 0
}

var File_carEnumProto_carEnum_proto protoreflect.FileDescriptor

var file_carEnumProto_carEnum_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x63, 0x61, 0x72, 0x45, 0x6e, 0x75, 0x6d, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63,
	0x61, 0x72, 0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x63, 0x61,
	0x72, 0x45, 0x6e, 0x75, 0x6d, 0x1a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x63, 0x61, 0x72, 0x67, 0x6f, 0x64, 0x2d, 0x62, 0x6a, 0x2f, 0x62, 0x32, 0x63, 0x2d,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x4d, 0x0a, 0x0b, 0x45, 0x6e, 0x75, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x44, 0x74, 0x6f, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x79, 0x70, 0x65, 0x73, 0x22,
	0x3a, 0x0a, 0x0c, 0x45, 0x6e, 0x75, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x73, 0x44, 0x74, 0x6f, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x79, 0x70, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x22, 0x88, 0x01, 0x0a, 0x08,
	0x45, 0x6e, 0x75, 0x6d, 0x73, 0x44, 0x74, 0x6f, 0x12, 0x32, 0x0a, 0x05, 0x65, 0x6e, 0x75, 0x6d,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x63, 0x61, 0x72, 0x45, 0x6e, 0x75,
	0x6d, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x44, 0x74, 0x6f, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x1a, 0x48, 0x0a, 0x0a,
	0x45, 0x6e, 0x75, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x24, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x63, 0x61,
	0x72, 0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x4b, 0x56, 0x4d, 0x61, 0x70, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x6a, 0x0a, 0x05, 0x4b, 0x56, 0x4d, 0x61, 0x70, 0x12,
	0x29, 0x0a, 0x03, 0x6b, 0x76, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x63,
	0x61, 0x72, 0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x4b, 0x56, 0x4d, 0x61, 0x70, 0x2e, 0x4b, 0x76, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x03, 0x6b, 0x76, 0x73, 0x1a, 0x36, 0x0a, 0x08, 0x4b, 0x76,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x22, 0x35, 0x0a, 0x0b, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x44, 0x74,
	0x6f, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x32, 0xb9, 0x01, 0x0a, 0x07, 0x43, 0x61,
	0x72, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x34, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x45,
	0x6e, 0x75, 0x6d, 0x12, 0x14, 0x2e, 0x63, 0x61, 0x72, 0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x45, 0x6e,
	0x75, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x44, 0x74, 0x6f, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x0b, 0x47,
	0x65, 0x74, 0x43, 0x61, 0x72, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x12, 0x15, 0x2e, 0x63, 0x61, 0x72,
	0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x73, 0x44, 0x74,
	0x6f, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x45, 0x6e, 0x75,
	0x6d, 0x73, 0x46, 0x6f, 0x72, 0x57, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x12, 0x15, 0x2e, 0x63,
	0x61, 0x72, 0x45, 0x6e, 0x75, 0x6d, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x73,
	0x44, 0x74, 0x6f, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x61, 0x72, 0x67, 0x6f, 0x64, 0x2d, 0x62, 0x6a, 0x2f, 0x62, 0x32,
	0x63, 0x2d, 0x63, 0x61, 0x72, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x61, 0x72, 0x45, 0x6e, 0x75,
	0x6d, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_carEnumProto_carEnum_proto_rawDescOnce sync.Once
	file_carEnumProto_carEnum_proto_rawDescData = file_carEnumProto_carEnum_proto_rawDesc
)

func file_carEnumProto_carEnum_proto_rawDescGZIP() []byte {
	file_carEnumProto_carEnum_proto_rawDescOnce.Do(func() {
		file_carEnumProto_carEnum_proto_rawDescData = protoimpl.X.CompressGZIP(file_carEnumProto_carEnum_proto_rawDescData)
	})
	return file_carEnumProto_carEnum_proto_rawDescData
}

var file_carEnumProto_carEnum_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_carEnumProto_carEnum_proto_goTypes = []interface{}{
	(*EnumTypeDto)(nil),     // 0: carEnum.EnumTypeDto
	(*EnumTypesDto)(nil),    // 1: carEnum.EnumTypesDto
	(*EnumsDto)(nil),        // 2: carEnum.EnumsDto
	(*KVMap)(nil),           // 3: carEnum.KVMap
	(*KeyValueDto)(nil),     // 4: carEnum.KeyValueDto
	nil,                     // 5: carEnum.EnumsDto.EnumsEntry
	nil,                     // 6: carEnum.KVMap.KvsEntry
	(*common.Response)(nil), // 7: common.Response
}
var file_carEnumProto_carEnum_proto_depIdxs = []int32{
	5, // 0: carEnum.EnumsDto.enums:type_name -> carEnum.EnumsDto.EnumsEntry
	6, // 1: carEnum.KVMap.kvs:type_name -> carEnum.KVMap.KvsEntry
	3, // 2: carEnum.EnumsDto.EnumsEntry.value:type_name -> carEnum.KVMap
	0, // 3: carEnum.CarEnum.GetCarEnum:input_type -> carEnum.EnumTypeDto
	1, // 4: carEnum.CarEnum.GetCarEnums:input_type -> carEnum.EnumTypesDto
	1, // 5: carEnum.CarEnum.GetCarEnumsForWebsite:input_type -> carEnum.EnumTypesDto
	7, // 6: carEnum.CarEnum.GetCarEnum:output_type -> common.Response
	7, // 7: carEnum.CarEnum.GetCarEnums:output_type -> common.Response
	7, // 8: carEnum.CarEnum.GetCarEnumsForWebsite:output_type -> common.Response
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_carEnumProto_carEnum_proto_init() }
func file_carEnumProto_carEnum_proto_init() {
	if File_carEnumProto_carEnum_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_carEnumProto_carEnum_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnumTypeDto); i {
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
		file_carEnumProto_carEnum_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnumTypesDto); i {
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
		file_carEnumProto_carEnum_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnumsDto); i {
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
		file_carEnumProto_carEnum_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KVMap); i {
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
		file_carEnumProto_carEnum_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeyValueDto); i {
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
			RawDescriptor: file_carEnumProto_carEnum_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_carEnumProto_carEnum_proto_goTypes,
		DependencyIndexes: file_carEnumProto_carEnum_proto_depIdxs,
		MessageInfos:      file_carEnumProto_carEnum_proto_msgTypes,
	}.Build()
	File_carEnumProto_carEnum_proto = out.File
	file_carEnumProto_carEnum_proto_rawDesc = nil
	file_carEnumProto_carEnum_proto_goTypes = nil
	file_carEnumProto_carEnum_proto_depIdxs = nil
}
