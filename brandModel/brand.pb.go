// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0-devel
// 	protoc        v3.12.1
// source: brandModel/brand.proto

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

// 品牌dto
type BrandDto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id，如果为新增，则此字段为0，否则为id
	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// 创建时间，如果为新增，或更新，则此字段无效
	CreateTime uint64 `protobuf:"varint,2,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// 更新时间，如果为新增，或更新，则此字段无效
	UpdateTime uint64 `protobuf:"varint,3,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	// 当前key的描述信息, 如果为更新，则此字段无效
	Name   string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Status uint32 `protobuf:"varint,5,opt,name=status,proto3" json:"status,omitempty"`
	Order  uint32 `protobuf:"varint,6,opt,name=order,proto3" json:"order,omitempty"`
}

func (x *BrandDto) Reset() {
	*x = BrandDto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_brandModel_brand_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BrandDto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BrandDto) ProtoMessage() {}

func (x *BrandDto) ProtoReflect() protoreflect.Message {
	mi := &file_brandModel_brand_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BrandDto.ProtoReflect.Descriptor instead.
func (*BrandDto) Descriptor() ([]byte, []int) {
	return file_brandModel_brand_proto_rawDescGZIP(), []int{0}
}

func (x *BrandDto) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *BrandDto) GetCreateTime() uint64 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

func (x *BrandDto) GetUpdateTime() uint64 {
	if x != nil {
		return x.UpdateTime
	}
	return 0
}

func (x *BrandDto) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *BrandDto) GetStatus() uint32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *BrandDto) GetOrder() uint32 {
	if x != nil {
		return x.Order
	}
	return 0
}

type NoticeCarMasterDTO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CarMasterId       uint32 `protobuf:"varint,1,opt,name=CarMasterId,proto3" json:"CarMasterId,omitempty"`
	CarBrandId        uint32 `protobuf:"varint,2,opt,name=CarBrandId,proto3" json:"CarBrandId,omitempty"`
	CarModelId        uint32 `protobuf:"varint,3,opt,name=CarModelId,proto3" json:"CarModelId,omitempty"`
	CarVariantId      uint32 `protobuf:"varint,4,opt,name=CarVariantId,proto3" json:"CarVariantId,omitempty"`
	CarYearId         uint32 `protobuf:"varint,5,opt,name=CarYearId,proto3" json:"CarYearId,omitempty"`
	CarEngineId       uint32 `protobuf:"varint,6,opt,name=CarEngineId,proto3" json:"CarEngineId,omitempty"`
	CarTransmissionId uint32 `protobuf:"varint,7,opt,name=CarTransmissionId,proto3" json:"CarTransmissionId,omitempty"`
	Year              uint32 `protobuf:"varint,8,opt,name=Year,proto3" json:"Year,omitempty"`
	Transmission      string `protobuf:"bytes,9,opt,name=Transmission,proto3" json:"Transmission,omitempty"`
	BrandName         string `protobuf:"bytes,10,opt,name=BrandName,proto3" json:"BrandName,omitempty"`
	ModelName         string `protobuf:"bytes,11,opt,name=ModelName,proto3" json:"ModelName,omitempty"`
	VariantName       string `protobuf:"bytes,12,opt,name=VariantName,proto3" json:"VariantName,omitempty"`
	Engine            string `protobuf:"bytes,13,opt,name=Engine,proto3" json:"Engine,omitempty"`
	BrandLogo         string `protobuf:"bytes,14,opt,name=BrandLogo,proto3" json:"BrandLogo,omitempty"`
}

func (x *NoticeCarMasterDTO) Reset() {
	*x = NoticeCarMasterDTO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_brandModel_brand_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NoticeCarMasterDTO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NoticeCarMasterDTO) ProtoMessage() {}

func (x *NoticeCarMasterDTO) ProtoReflect() protoreflect.Message {
	mi := &file_brandModel_brand_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NoticeCarMasterDTO.ProtoReflect.Descriptor instead.
func (*NoticeCarMasterDTO) Descriptor() ([]byte, []int) {
	return file_brandModel_brand_proto_rawDescGZIP(), []int{1}
}

func (x *NoticeCarMasterDTO) GetCarMasterId() uint32 {
	if x != nil {
		return x.CarMasterId
	}
	return 0
}

func (x *NoticeCarMasterDTO) GetCarBrandId() uint32 {
	if x != nil {
		return x.CarBrandId
	}
	return 0
}

func (x *NoticeCarMasterDTO) GetCarModelId() uint32 {
	if x != nil {
		return x.CarModelId
	}
	return 0
}

func (x *NoticeCarMasterDTO) GetCarVariantId() uint32 {
	if x != nil {
		return x.CarVariantId
	}
	return 0
}

func (x *NoticeCarMasterDTO) GetCarYearId() uint32 {
	if x != nil {
		return x.CarYearId
	}
	return 0
}

func (x *NoticeCarMasterDTO) GetCarEngineId() uint32 {
	if x != nil {
		return x.CarEngineId
	}
	return 0
}

func (x *NoticeCarMasterDTO) GetCarTransmissionId() uint32 {
	if x != nil {
		return x.CarTransmissionId
	}
	return 0
}

func (x *NoticeCarMasterDTO) GetYear() uint32 {
	if x != nil {
		return x.Year
	}
	return 0
}

func (x *NoticeCarMasterDTO) GetTransmission() string {
	if x != nil {
		return x.Transmission
	}
	return ""
}

func (x *NoticeCarMasterDTO) GetBrandName() string {
	if x != nil {
		return x.BrandName
	}
	return ""
}

func (x *NoticeCarMasterDTO) GetModelName() string {
	if x != nil {
		return x.ModelName
	}
	return ""
}

func (x *NoticeCarMasterDTO) GetVariantName() string {
	if x != nil {
		return x.VariantName
	}
	return ""
}

func (x *NoticeCarMasterDTO) GetEngine() string {
	if x != nil {
		return x.Engine
	}
	return ""
}

func (x *NoticeCarMasterDTO) GetBrandLogo() string {
	if x != nil {
		return x.BrandLogo
	}
	return ""
}

var File_brandModel_brand_proto protoreflect.FileDescriptor

var file_brandModel_brand_proto_rawDesc = []byte{
	0x0a, 0x16, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x62, 0x72, 0x61,
	0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x4d,
	0x6f, 0x64, 0x65, 0x6c, 0x1a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x63, 0x61, 0x72, 0x67, 0x6f, 0x64, 0x2d, 0x62, 0x6a, 0x2f, 0x62, 0x32, 0x63, 0x2d, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x9e, 0x01, 0x0a, 0x08, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x44, 0x74, 0x6f, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1f, 0x0a,
	0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x22, 0xd4, 0x03, 0x0a, 0x12, 0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x43, 0x61, 0x72, 0x4d, 0x61,
	0x73, 0x74, 0x65, 0x72, 0x44, 0x54, 0x4f, 0x12, 0x20, 0x0a, 0x0b, 0x43, 0x61, 0x72, 0x4d, 0x61,
	0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x43, 0x61,
	0x72, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x61, 0x72,
	0x42, 0x72, 0x61, 0x6e, 0x64, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x43,
	0x61, 0x72, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x61, 0x72,
	0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x43,
	0x61, 0x72, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x43, 0x61, 0x72,
	0x56, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0c, 0x43, 0x61, 0x72, 0x56, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1c, 0x0a,
	0x09, 0x43, 0x61, 0x72, 0x59, 0x65, 0x61, 0x72, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x09, 0x43, 0x61, 0x72, 0x59, 0x65, 0x61, 0x72, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x43,
	0x61, 0x72, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x49, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0b, 0x43, 0x61, 0x72, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x49, 0x64, 0x12, 0x2c, 0x0a,
	0x11, 0x43, 0x61, 0x72, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x49, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x11, 0x43, 0x61, 0x72, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x59,
	0x65, 0x61, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x59, 0x65, 0x61, 0x72, 0x12,
	0x22, 0x0a, 0x0c, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x1c, 0x0a, 0x09, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x56, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x56, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x42, 0x72, 0x61,
	0x6e, 0x64, 0x4c, 0x6f, 0x67, 0x6f, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x42, 0x72,
	0x61, 0x6e, 0x64, 0x4c, 0x6f, 0x67, 0x6f, 0x32, 0xf4, 0x02, 0x0a, 0x05, 0x42, 0x72, 0x61, 0x6e,
	0x64, 0x12, 0x2f, 0x0a, 0x03, 0x41, 0x64, 0x64, 0x12, 0x14, 0x2e, 0x62, 0x72, 0x61, 0x6e, 0x64,
	0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x44, 0x74, 0x6f, 0x1a, 0x10,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x30, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x12, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x49, 0x64, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x44, 0x54, 0x4f,
	0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x32, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x14,
	0x2e, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x42, 0x72, 0x61, 0x6e,
	0x64, 0x44, 0x74, 0x6f, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2d, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12,
	0x12, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x49, 0x64, 0x4c, 0x6f, 0x63, 0x61, 0x6c,
	0x44, 0x54, 0x4f, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2d, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x11, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x50, 0x61,
	0x67, 0x65, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x72,
	0x61, 0x6e, 0x64, 0x48, 0x61, 0x73, 0x43, 0x61, 0x72, 0x12, 0x11, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x50, 0x61, 0x67, 0x65, 0x1a, 0x10, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x3c, 0x0a, 0x06, 0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x12, 0x1e, 0x2e, 0x62, 0x72, 0x61,
	0x6e, 0x64, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x43, 0x61,
	0x72, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x44, 0x54, 0x4f, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x32,
	0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x61, 0x72,
	0x67, 0x6f, 0x64, 0x2d, 0x62, 0x6a, 0x2f, 0x62, 0x32, 0x63, 0x2d, 0x63, 0x61, 0x72, 0x2d, 0x61,
	0x70, 0x69, 0x2f, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_brandModel_brand_proto_rawDescOnce sync.Once
	file_brandModel_brand_proto_rawDescData = file_brandModel_brand_proto_rawDesc
)

func file_brandModel_brand_proto_rawDescGZIP() []byte {
	file_brandModel_brand_proto_rawDescOnce.Do(func() {
		file_brandModel_brand_proto_rawDescData = protoimpl.X.CompressGZIP(file_brandModel_brand_proto_rawDescData)
	})
	return file_brandModel_brand_proto_rawDescData
}

var file_brandModel_brand_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_brandModel_brand_proto_goTypes = []interface{}{
	(*BrandDto)(nil),           // 0: brandModel.BrandDto
	(*NoticeCarMasterDTO)(nil), // 1: brandModel.NoticeCarMasterDTO
	(*common.IdLocalDTO)(nil),  // 2: common.IdLocalDTO
	(*common.LocalPage)(nil),   // 3: common.LocalPage
	(*common.Response)(nil),    // 4: common.Response
}
var file_brandModel_brand_proto_depIdxs = []int32{
	0, // 0: brandModel.Brand.Add:input_type -> brandModel.BrandDto
	2, // 1: brandModel.Brand.Delete:input_type -> common.IdLocalDTO
	0, // 2: brandModel.Brand.Update:input_type -> brandModel.BrandDto
	2, // 3: brandModel.Brand.Get:input_type -> common.IdLocalDTO
	3, // 4: brandModel.Brand.List:input_type -> common.LocalPage
	3, // 5: brandModel.Brand.ListBrandHasCar:input_type -> common.LocalPage
	1, // 6: brandModel.Brand.Notice:input_type -> brandModel.NoticeCarMasterDTO
	4, // 7: brandModel.Brand.Add:output_type -> common.Response
	4, // 8: brandModel.Brand.Delete:output_type -> common.Response
	4, // 9: brandModel.Brand.Update:output_type -> common.Response
	4, // 10: brandModel.Brand.Get:output_type -> common.Response
	4, // 11: brandModel.Brand.List:output_type -> common.Response
	4, // 12: brandModel.Brand.ListBrandHasCar:output_type -> common.Response
	4, // 13: brandModel.Brand.Notice:output_type -> common.Response
	7, // [7:14] is the sub-list for method output_type
	0, // [0:7] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_brandModel_brand_proto_init() }
func file_brandModel_brand_proto_init() {
	if File_brandModel_brand_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_brandModel_brand_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BrandDto); i {
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
		file_brandModel_brand_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NoticeCarMasterDTO); i {
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
			RawDescriptor: file_brandModel_brand_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_brandModel_brand_proto_goTypes,
		DependencyIndexes: file_brandModel_brand_proto_depIdxs,
		MessageInfos:      file_brandModel_brand_proto_msgTypes,
	}.Build()
	File_brandModel_brand_proto = out.File
	file_brandModel_brand_proto_rawDesc = nil
	file_brandModel_brand_proto_goTypes = nil
	file_brandModel_brand_proto_depIdxs = nil
}
