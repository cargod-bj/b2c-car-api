// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0
// 	protoc        v3.9.0
// source: brandModel/year_engine_trans.proto

package brandModelProto

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

// key dto
type KeyDto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// key
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *KeyDto) Reset() {
	*x = KeyDto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_brandModel_year_engine_trans_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeyDto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyDto) ProtoMessage() {}

func (x *KeyDto) ProtoReflect() protoreflect.Message {
	mi := &file_brandModel_year_engine_trans_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyDto.ProtoReflect.Descriptor instead.
func (*KeyDto) Descriptor() ([]byte, []int) {
	return file_brandModel_year_engine_trans_proto_rawDescGZIP(), []int{0}
}

func (x *KeyDto) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

// model和year的入参
type ModelYearReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Year    string `protobuf:"bytes,1,opt,name=year,proto3" json:"year,omitempty"`
	ModelId uint32 `protobuf:"varint,2,opt,name=modelId,proto3" json:"modelId,omitempty"`
}

func (x *ModelYearReq) Reset() {
	*x = ModelYearReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_brandModel_year_engine_trans_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModelYearReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModelYearReq) ProtoMessage() {}

func (x *ModelYearReq) ProtoReflect() protoreflect.Message {
	mi := &file_brandModel_year_engine_trans_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModelYearReq.ProtoReflect.Descriptor instead.
func (*ModelYearReq) Descriptor() ([]byte, []int) {
	return file_brandModel_year_engine_trans_proto_rawDescGZIP(), []int{1}
}

func (x *ModelYearReq) GetYear() string {
	if x != nil {
		return x.Year
	}
	return ""
}

func (x *ModelYearReq) GetModelId() uint32 {
	if x != nil {
		return x.ModelId
	}
	return 0
}

// model、year和variant的入参
type ModelYearVariantReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Year      string `protobuf:"bytes,1,opt,name=year,proto3" json:"year,omitempty"`
	ModelId   uint32 `protobuf:"varint,2,opt,name=modelId,proto3" json:"modelId,omitempty"`
	VariantId int64  `protobuf:"varint,3,opt,name=variantId,proto3" json:"variantId,omitempty"`
}

func (x *ModelYearVariantReq) Reset() {
	*x = ModelYearVariantReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_brandModel_year_engine_trans_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModelYearVariantReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModelYearVariantReq) ProtoMessage() {}

func (x *ModelYearVariantReq) ProtoReflect() protoreflect.Message {
	mi := &file_brandModel_year_engine_trans_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModelYearVariantReq.ProtoReflect.Descriptor instead.
func (*ModelYearVariantReq) Descriptor() ([]byte, []int) {
	return file_brandModel_year_engine_trans_proto_rawDescGZIP(), []int{2}
}

func (x *ModelYearVariantReq) GetYear() string {
	if x != nil {
		return x.Year
	}
	return ""
}

func (x *ModelYearVariantReq) GetModelId() uint32 {
	if x != nil {
		return x.ModelId
	}
	return 0
}

func (x *ModelYearVariantReq) GetVariantId() int64 {
	if x != nil {
		return x.VariantId
	}
	return 0
}

// brand、model、year、variant、engine的入参
type ModelYearVariantEngineReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Year      string `protobuf:"bytes,1,opt,name=year,proto3" json:"year,omitempty"`
	ModelId   uint32 `protobuf:"varint,2,opt,name=modelId,proto3" json:"modelId,omitempty"`
	VariantId int64  `protobuf:"varint,3,opt,name=variantId,proto3" json:"variantId,omitempty"`
	Engine    string `protobuf:"bytes,4,opt,name=engine,proto3" json:"engine,omitempty"`
	BrandId   uint32 `protobuf:"varint,5,opt,name=brandId,proto3" json:"brandId,omitempty"`
}

func (x *ModelYearVariantEngineReq) Reset() {
	*x = ModelYearVariantEngineReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_brandModel_year_engine_trans_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModelYearVariantEngineReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModelYearVariantEngineReq) ProtoMessage() {}

func (x *ModelYearVariantEngineReq) ProtoReflect() protoreflect.Message {
	mi := &file_brandModel_year_engine_trans_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModelYearVariantEngineReq.ProtoReflect.Descriptor instead.
func (*ModelYearVariantEngineReq) Descriptor() ([]byte, []int) {
	return file_brandModel_year_engine_trans_proto_rawDescGZIP(), []int{3}
}

func (x *ModelYearVariantEngineReq) GetYear() string {
	if x != nil {
		return x.Year
	}
	return ""
}

func (x *ModelYearVariantEngineReq) GetModelId() uint32 {
	if x != nil {
		return x.ModelId
	}
	return 0
}

func (x *ModelYearVariantEngineReq) GetVariantId() int64 {
	if x != nil {
		return x.VariantId
	}
	return 0
}

func (x *ModelYearVariantEngineReq) GetEngine() string {
	if x != nil {
		return x.Engine
	}
	return ""
}

func (x *ModelYearVariantEngineReq) GetBrandId() uint32 {
	if x != nil {
		return x.BrandId
	}
	return 0
}

// model和year的入参
type VariantEngineReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Engine    string `protobuf:"bytes,1,opt,name=engine,proto3" json:"engine,omitempty"`
	VariantId int64  `protobuf:"varint,2,opt,name=variantId,proto3" json:"variantId,omitempty"`
}

func (x *VariantEngineReq) Reset() {
	*x = VariantEngineReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_brandModel_year_engine_trans_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VariantEngineReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VariantEngineReq) ProtoMessage() {}

func (x *VariantEngineReq) ProtoReflect() protoreflect.Message {
	mi := &file_brandModel_year_engine_trans_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VariantEngineReq.ProtoReflect.Descriptor instead.
func (*VariantEngineReq) Descriptor() ([]byte, []int) {
	return file_brandModel_year_engine_trans_proto_rawDescGZIP(), []int{4}
}

func (x *VariantEngineReq) GetEngine() string {
	if x != nil {
		return x.Engine
	}
	return ""
}

func (x *VariantEngineReq) GetVariantId() int64 {
	if x != nil {
		return x.VariantId
	}
	return 0
}

// variant req
type FindEngineByVariantHasCarReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Year      string `protobuf:"bytes,1,opt,name=year,proto3" json:"year,omitempty"`
	ModelId   uint32 `protobuf:"varint,2,opt,name=modelId,proto3" json:"modelId,omitempty"`
	VariantId int64  `protobuf:"varint,3,opt,name=variantId,proto3" json:"variantId,omitempty"`
}

func (x *FindEngineByVariantHasCarReq) Reset() {
	*x = FindEngineByVariantHasCarReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_brandModel_year_engine_trans_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindEngineByVariantHasCarReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindEngineByVariantHasCarReq) ProtoMessage() {}

func (x *FindEngineByVariantHasCarReq) ProtoReflect() protoreflect.Message {
	mi := &file_brandModel_year_engine_trans_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindEngineByVariantHasCarReq.ProtoReflect.Descriptor instead.
func (*FindEngineByVariantHasCarReq) Descriptor() ([]byte, []int) {
	return file_brandModel_year_engine_trans_proto_rawDescGZIP(), []int{5}
}

func (x *FindEngineByVariantHasCarReq) GetYear() string {
	if x != nil {
		return x.Year
	}
	return ""
}

func (x *FindEngineByVariantHasCarReq) GetModelId() uint32 {
	if x != nil {
		return x.ModelId
	}
	return 0
}

func (x *FindEngineByVariantHasCarReq) GetVariantId() int64 {
	if x != nil {
		return x.VariantId
	}
	return 0
}

type IdNameDto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *IdNameDto) Reset() {
	*x = IdNameDto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_brandModel_year_engine_trans_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IdNameDto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdNameDto) ProtoMessage() {}

func (x *IdNameDto) ProtoReflect() protoreflect.Message {
	mi := &file_brandModel_year_engine_trans_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdNameDto.ProtoReflect.Descriptor instead.
func (*IdNameDto) Descriptor() ([]byte, []int) {
	return file_brandModel_year_engine_trans_proto_rawDescGZIP(), []int{6}
}

func (x *IdNameDto) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *IdNameDto) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_brandModel_year_engine_trans_proto protoreflect.FileDescriptor

var file_brandModel_year_engine_trans_proto_rawDesc = []byte{
	0x0a, 0x22, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x79, 0x65, 0x61,
	0x72, 0x5f, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x4d, 0x6f, 0x64, 0x65, 0x6c,
	0x1a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x61, 0x72,
	0x67, 0x6f, 0x64, 0x2d, 0x62, 0x6a, 0x2f, 0x62, 0x32, 0x63, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2d, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1a, 0x0a, 0x06, 0x4b,
	0x65, 0x79, 0x44, 0x74, 0x6f, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x3c, 0x0a, 0x0c, 0x4d, 0x6f, 0x64, 0x65, 0x6c,
	0x59, 0x65, 0x61, 0x72, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x79, 0x65, 0x61, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x79, 0x65, 0x61, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x49, 0x64, 0x22, 0x61, 0x0a, 0x13, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x59, 0x65,
	0x61, 0x72, 0x56, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04,
	0x79, 0x65, 0x61, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x79, 0x65, 0x61, 0x72,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x07, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x76, 0x61,
	0x72, 0x69, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x76,
	0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x99, 0x01, 0x0a, 0x19, 0x4d, 0x6f, 0x64,
	0x65, 0x6c, 0x59, 0x65, 0x61, 0x72, 0x56, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x45, 0x6e, 0x67,
	0x69, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x79, 0x65, 0x61, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x79, 0x65, 0x61, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x76, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x49,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x76, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74,
	0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x72,
	0x61, 0x6e, 0x64, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x62, 0x72, 0x61,
	0x6e, 0x64, 0x49, 0x64, 0x22, 0x48, 0x0a, 0x10, 0x56, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x45,
	0x6e, 0x67, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x6e, 0x67, 0x69,
	0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65,
	0x12, 0x1c, 0x0a, 0x09, 0x76, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x76, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x6a,
	0x0a, 0x1c, 0x46, 0x69, 0x6e, 0x64, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x42, 0x79, 0x56, 0x61,
	0x72, 0x69, 0x61, 0x6e, 0x74, 0x48, 0x61, 0x73, 0x43, 0x61, 0x72, 0x52, 0x65, 0x71, 0x12, 0x12,
	0x0a, 0x04, 0x79, 0x65, 0x61, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x79, 0x65,
	0x61, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x49, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x07, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09,
	0x76, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x76, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x2f, 0x0a, 0x09, 0x49, 0x64,
	0x4e, 0x61, 0x6d, 0x65, 0x44, 0x74, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x32, 0xe7, 0x04, 0x0a, 0x0f,
	0x59, 0x65, 0x61, 0x72, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x12,
	0x34, 0x0a, 0x0f, 0x46, 0x69, 0x6e, 0x64, 0x59, 0x65, 0x61, 0x72, 0x42, 0x79, 0x4d, 0x6f, 0x64,
	0x65, 0x6c, 0x12, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x49, 0x64, 0x44, 0x74,
	0x6f, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x11, 0x46, 0x69, 0x6e, 0x64, 0x56, 0x61, 0x72,
	0x69, 0x61, 0x6e, 0x74, 0x42, 0x79, 0x59, 0x65, 0x61, 0x72, 0x12, 0x18, 0x2e, 0x62, 0x72, 0x61,
	0x6e, 0x64, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x59, 0x65, 0x61,
	0x72, 0x52, 0x65, 0x71, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4a, 0x0a, 0x13, 0x46, 0x69, 0x6e, 0x64,
	0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x42, 0x79, 0x56, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x12,
	0x1f, 0x2e, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x4d, 0x6f, 0x64,
	0x65, 0x6c, 0x59, 0x65, 0x61, 0x72, 0x56, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x71,
	0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x55, 0x0a, 0x18, 0x46, 0x69, 0x6e, 0x64, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x79, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65,
	0x12, 0x25, 0x2e, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x4d, 0x6f,
	0x64, 0x65, 0x6c, 0x59, 0x65, 0x61, 0x72, 0x56, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x45, 0x6e,
	0x67, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x15, 0x46,
	0x69, 0x6e, 0x64, 0x59, 0x65, 0x61, 0x72, 0x42, 0x79, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x48, 0x61,
	0x73, 0x43, 0x61, 0x72, 0x12, 0x12, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x49, 0x64,
	0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x44, 0x54, 0x4f, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x47, 0x0a, 0x17,
	0x46, 0x69, 0x6e, 0x64, 0x56, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x42, 0x79, 0x59, 0x65, 0x61,
	0x72, 0x48, 0x61, 0x73, 0x43, 0x61, 0x72, 0x12, 0x18, 0x2e, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x4d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x59, 0x65, 0x61, 0x72, 0x52, 0x65,
	0x71, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x59, 0x0a, 0x19, 0x46, 0x69, 0x6e, 0x64, 0x45, 0x6e, 0x67,
	0x69, 0x6e, 0x65, 0x42, 0x79, 0x56, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x48, 0x61, 0x73, 0x43,
	0x61, 0x72, 0x12, 0x28, 0x2e, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x2e,
	0x46, 0x69, 0x6e, 0x64, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x42, 0x79, 0x56, 0x61, 0x72, 0x69,
	0x61, 0x6e, 0x74, 0x48, 0x61, 0x73, 0x43, 0x61, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x10, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x53, 0x0a, 0x16, 0x46, 0x69, 0x6e, 0x64, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x48, 0x61, 0x73, 0x43, 0x61, 0x72, 0x12, 0x25, 0x2e, 0x62, 0x72, 0x61,
	0x6e, 0x64, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x59, 0x65, 0x61,
	0x72, 0x56, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x52, 0x65,
	0x71, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x32, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x61, 0x72, 0x67, 0x6f, 0x64, 0x2d, 0x62, 0x6a, 0x2f, 0x62, 0x32,
	0x63, 0x2d, 0x63, 0x61, 0x72, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x4d,
	0x6f, 0x64, 0x65, 0x6c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_brandModel_year_engine_trans_proto_rawDescOnce sync.Once
	file_brandModel_year_engine_trans_proto_rawDescData = file_brandModel_year_engine_trans_proto_rawDesc
)

func file_brandModel_year_engine_trans_proto_rawDescGZIP() []byte {
	file_brandModel_year_engine_trans_proto_rawDescOnce.Do(func() {
		file_brandModel_year_engine_trans_proto_rawDescData = protoimpl.X.CompressGZIP(file_brandModel_year_engine_trans_proto_rawDescData)
	})
	return file_brandModel_year_engine_trans_proto_rawDescData
}

var file_brandModel_year_engine_trans_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_brandModel_year_engine_trans_proto_goTypes = []interface{}{
	(*KeyDto)(nil),                       // 0: brandModel.KeyDto
	(*ModelYearReq)(nil),                 // 1: brandModel.ModelYearReq
	(*ModelYearVariantReq)(nil),          // 2: brandModel.ModelYearVariantReq
	(*ModelYearVariantEngineReq)(nil),    // 3: brandModel.ModelYearVariantEngineReq
	(*VariantEngineReq)(nil),             // 4: brandModel.VariantEngineReq
	(*FindEngineByVariantHasCarReq)(nil), // 5: brandModel.FindEngineByVariantHasCarReq
	(*IdNameDto)(nil),                    // 6: brandModel.IdNameDto
	(*common.IdDto)(nil),                 // 7: common.IdDto
	(*common.IdLocalDTO)(nil),            // 8: common.IdLocalDTO
	(*common.Response)(nil),              // 9: common.Response
}
var file_brandModel_year_engine_trans_proto_depIdxs = []int32{
	7, // 0: brandModel.YearEngineTrans.FindYearByModel:input_type -> common.IdDto
	1, // 1: brandModel.YearEngineTrans.FindVariantByYear:input_type -> brandModel.ModelYearReq
	2, // 2: brandModel.YearEngineTrans.FindEngineByVariant:input_type -> brandModel.ModelYearVariantReq
	3, // 3: brandModel.YearEngineTrans.FindTransmissionByEngine:input_type -> brandModel.ModelYearVariantEngineReq
	8, // 4: brandModel.YearEngineTrans.FindYearByModelHasCar:input_type -> common.IdLocalDTO
	1, // 5: brandModel.YearEngineTrans.FindVariantByYearHasCar:input_type -> brandModel.ModelYearReq
	5, // 6: brandModel.YearEngineTrans.FindEngineByVariantHasCar:input_type -> brandModel.FindEngineByVariantHasCarReq
	3, // 7: brandModel.YearEngineTrans.FindTransmissionHasCar:input_type -> brandModel.ModelYearVariantEngineReq
	9, // 8: brandModel.YearEngineTrans.FindYearByModel:output_type -> common.Response
	9, // 9: brandModel.YearEngineTrans.FindVariantByYear:output_type -> common.Response
	9, // 10: brandModel.YearEngineTrans.FindEngineByVariant:output_type -> common.Response
	9, // 11: brandModel.YearEngineTrans.FindTransmissionByEngine:output_type -> common.Response
	9, // 12: brandModel.YearEngineTrans.FindYearByModelHasCar:output_type -> common.Response
	9, // 13: brandModel.YearEngineTrans.FindVariantByYearHasCar:output_type -> common.Response
	9, // 14: brandModel.YearEngineTrans.FindEngineByVariantHasCar:output_type -> common.Response
	9, // 15: brandModel.YearEngineTrans.FindTransmissionHasCar:output_type -> common.Response
	8, // [8:16] is the sub-list for method output_type
	0, // [0:8] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_brandModel_year_engine_trans_proto_init() }
func file_brandModel_year_engine_trans_proto_init() {
	if File_brandModel_year_engine_trans_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_brandModel_year_engine_trans_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeyDto); i {
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
		file_brandModel_year_engine_trans_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModelYearReq); i {
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
		file_brandModel_year_engine_trans_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModelYearVariantReq); i {
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
		file_brandModel_year_engine_trans_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModelYearVariantEngineReq); i {
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
		file_brandModel_year_engine_trans_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VariantEngineReq); i {
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
		file_brandModel_year_engine_trans_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindEngineByVariantHasCarReq); i {
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
		file_brandModel_year_engine_trans_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IdNameDto); i {
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
			RawDescriptor: file_brandModel_year_engine_trans_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_brandModel_year_engine_trans_proto_goTypes,
		DependencyIndexes: file_brandModel_year_engine_trans_proto_depIdxs,
		MessageInfos:      file_brandModel_year_engine_trans_proto_msgTypes,
	}.Build()
	File_brandModel_year_engine_trans_proto = out.File
	file_brandModel_year_engine_trans_proto_rawDesc = nil
	file_brandModel_year_engine_trans_proto_goTypes = nil
	file_brandModel_year_engine_trans_proto_depIdxs = nil
}
