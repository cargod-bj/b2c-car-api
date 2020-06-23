// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0-devel
// 	protoc        v3.12.1
// source: report/report.proto

package reportProto

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

// 整备类型
type ReconditionType int32

const (
	// 没有整备过
	ReconditionType_NO ReconditionType = 0
	// 已修复
	ReconditionType_FIXED ReconditionType = 1
	// 已更新
	ReconditionType_REPLACED ReconditionType = 2
)

// Enum value maps for ReconditionType.
var (
	ReconditionType_name = map[int32]string{
		0: "NO",
		1: "FIXED",
		2: "REPLACED",
	}
	ReconditionType_value = map[string]int32{
		"NO":       0,
		"FIXED":    1,
		"REPLACED": 2,
	}
)

func (x ReconditionType) Enum() *ReconditionType {
	p := new(ReconditionType)
	*p = x
	return p
}

func (x ReconditionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ReconditionType) Descriptor() protoreflect.EnumDescriptor {
	return file_report_report_proto_enumTypes[0].Descriptor()
}

func (ReconditionType) Type() protoreflect.EnumType {
	return &file_report_report_proto_enumTypes[0]
}

func (x ReconditionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ReconditionType.Descriptor instead.
func (ReconditionType) EnumDescriptor() ([]byte, []int) {
	return file_report_report_proto_rawDescGZIP(), []int{0}
}

// damage的顶层类别，每个顶层类别可能包含多个 子类别
type ReportDamageCategoryDto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 当前category的id
	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// 当前category的Name
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// 当前category的子项，如果没有子项则为nil
	Children []*ReportDamageSubCategoryDto `protobuf:"bytes,3,rep,name=children,proto3" json:"children,omitempty"`
	// 包含多少个point点
	TotalPoint uint32 `protobuf:"varint,4,opt,name=totalPoint,proto3" json:"totalPoint,omitempty"`
	// 包含多少个没有修复的损坏点
	TotalDamage uint32 `protobuf:"varint,5,opt,name=totalDamage,proto3" json:"totalDamage,omitempty"`
}

func (x *ReportDamageCategoryDto) Reset() {
	*x = ReportDamageCategoryDto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_report_report_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportDamageCategoryDto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportDamageCategoryDto) ProtoMessage() {}

func (x *ReportDamageCategoryDto) ProtoReflect() protoreflect.Message {
	mi := &file_report_report_proto_msgTypes[0]
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
	return file_report_report_proto_rawDescGZIP(), []int{0}
}

func (x *ReportDamageCategoryDto) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ReportDamageCategoryDto) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ReportDamageCategoryDto) GetChildren() []*ReportDamageSubCategoryDto {
	if x != nil {
		return x.Children
	}
	return nil
}

func (x *ReportDamageCategoryDto) GetTotalPoint() uint32 {
	if x != nil {
		return x.TotalPoint
	}
	return 0
}

func (x *ReportDamageCategoryDto) GetTotalDamage() uint32 {
	if x != nil {
		return x.TotalDamage
	}
	return 0
}

// damage的子类别，每个子类别可能包含多个 监测点
type ReportDamageSubCategoryDto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 当前category的id
	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// 当前category的Name
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// 当前category的子项，如果没有子项则为nil
	Points []*ReportDamagePointDto `protobuf:"bytes,3,rep,name=points,proto3" json:"points,omitempty"`
	// 包含多少个point点
	TotalPoint uint32 `protobuf:"varint,4,opt,name=totalPoint,proto3" json:"totalPoint,omitempty"`
	// 包含多少个没有修复的损坏点
	TotalDamage uint32 `protobuf:"varint,5,opt,name=totalDamage,proto3" json:"totalDamage,omitempty"`
}

func (x *ReportDamageSubCategoryDto) Reset() {
	*x = ReportDamageSubCategoryDto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_report_report_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportDamageSubCategoryDto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportDamageSubCategoryDto) ProtoMessage() {}

func (x *ReportDamageSubCategoryDto) ProtoReflect() protoreflect.Message {
	mi := &file_report_report_proto_msgTypes[1]
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
	return file_report_report_proto_rawDescGZIP(), []int{1}
}

func (x *ReportDamageSubCategoryDto) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ReportDamageSubCategoryDto) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ReportDamageSubCategoryDto) GetPoints() []*ReportDamagePointDto {
	if x != nil {
		return x.Points
	}
	return nil
}

func (x *ReportDamageSubCategoryDto) GetTotalPoint() uint32 {
	if x != nil {
		return x.TotalPoint
	}
	return 0
}

func (x *ReportDamageSubCategoryDto) GetTotalDamage() uint32 {
	if x != nil {
		return x.TotalDamage
	}
	return 0
}

// damage的监测点，每个监测点可能包含多个 tag点，多个 损坏照片
type ReportDamagePointDto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 当前point的id
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// 当前point的Name
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// 当前point上的损坏点列表
	Tags []string `protobuf:"bytes,3,rep,name=tags,proto3" json:"tags,omitempty"`
	// 当前point上的损坏点图片
	Photos []string `protobuf:"bytes,4,rep,name=photos,proto3" json:"photos,omitempty"`
	// 是否已通过
	IsPassed bool `protobuf:"varint,5,opt,name=isPassed,proto3" json:"isPassed,omitempty"`
	// 是否为na
	IsNA bool `protobuf:"varint,6,opt,name=isNA,proto3" json:"isNA,omitempty"`
	// 当前监测点上的value值
	TagValue string `protobuf:"bytes,7,opt,name=tagValue,proto3" json:"tagValue,omitempty"`
	// 当前监测点上备注的信息
	Remark string `protobuf:"bytes,8,opt,name=remark,proto3" json:"remark,omitempty"`
	// 当前监测点的损坏照片
	DamagePhotos []string `protobuf:"bytes,9,rep,name=damagePhotos,proto3" json:"damagePhotos,omitempty"`
	// 当前监测点的整备照片
	ReconditionPhotos []string `protobuf:"bytes,10,rep,name=reconditionPhotos,proto3" json:"reconditionPhotos,omitempty"`
}

func (x *ReportDamagePointDto) Reset() {
	*x = ReportDamagePointDto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_report_report_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportDamagePointDto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportDamagePointDto) ProtoMessage() {}

func (x *ReportDamagePointDto) ProtoReflect() protoreflect.Message {
	mi := &file_report_report_proto_msgTypes[2]
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
	return file_report_report_proto_rawDescGZIP(), []int{2}
}

func (x *ReportDamagePointDto) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ReportDamagePointDto) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
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

func (x *ReportDamagePointDto) GetIsPassed() bool {
	if x != nil {
		return x.IsPassed
	}
	return false
}

func (x *ReportDamagePointDto) GetIsNA() bool {
	if x != nil {
		return x.IsNA
	}
	return false
}

func (x *ReportDamagePointDto) GetTagValue() string {
	if x != nil {
		return x.TagValue
	}
	return ""
}

func (x *ReportDamagePointDto) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

func (x *ReportDamagePointDto) GetDamagePhotos() []string {
	if x != nil {
		return x.DamagePhotos
	}
	return nil
}

func (x *ReportDamagePointDto) GetReconditionPhotos() []string {
	if x != nil {
		return x.ReconditionPhotos
	}
	return nil
}

// 做Damage的整备message
type PointReconditionDto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 当前point的id
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// 当前point的remark，可以为空
	Remark string `protobuf:"bytes,2,opt,name=remark,proto3" json:"remark,omitempty"`
	// 当前point的整备类型：0 没有整备过，1 已修复，2 已更新
	ReconditionType ReconditionType `protobuf:"varint,3,opt,name=reconditionType,proto3,enum=report.ReconditionType" json:"reconditionType,omitempty"`
}

func (x *PointReconditionDto) Reset() {
	*x = PointReconditionDto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_report_report_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PointReconditionDto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PointReconditionDto) ProtoMessage() {}

func (x *PointReconditionDto) ProtoReflect() protoreflect.Message {
	mi := &file_report_report_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PointReconditionDto.ProtoReflect.Descriptor instead.
func (*PointReconditionDto) Descriptor() ([]byte, []int) {
	return file_report_report_proto_rawDescGZIP(), []int{3}
}

func (x *PointReconditionDto) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *PointReconditionDto) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

func (x *PointReconditionDto) GetReconditionType() ReconditionType {
	if x != nil {
		return x.ReconditionType
	}
	return ReconditionType_NO
}

// DamagePhotos
type DamagePhotoDto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 当前point的Id
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// 要进行绑定的url或者解绑的url
	Urls []string `protobuf:"bytes,2,rep,name=urls,proto3" json:"urls,omitempty"`
}

func (x *DamagePhotoDto) Reset() {
	*x = DamagePhotoDto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_report_report_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DamagePhotoDto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DamagePhotoDto) ProtoMessage() {}

func (x *DamagePhotoDto) ProtoReflect() protoreflect.Message {
	mi := &file_report_report_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DamagePhotoDto.ProtoReflect.Descriptor instead.
func (*DamagePhotoDto) Descriptor() ([]byte, []int) {
	return file_report_report_proto_rawDescGZIP(), []int{4}
}

func (x *DamagePhotoDto) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *DamagePhotoDto) GetUrls() []string {
	if x != nil {
		return x.Urls
	}
	return nil
}

var File_report_report_proto protoreflect.FileDescriptor

var file_report_report_proto_rawDesc = []byte{
	0x0a, 0x13, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x1a, 0x39, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x61, 0x72, 0x67, 0x6f, 0x64,
	0x2d, 0x62, 0x6a, 0x2f, 0x62, 0x32, 0x63, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbf, 0x01, 0x0a, 0x17, 0x52, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x44, 0x61, 0x6d, 0x61, 0x67, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x44, 0x74, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3e, 0x0a, 0x08, 0x63, 0x68, 0x69, 0x6c,
	0x64, 0x72, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x72, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x44, 0x61, 0x6d, 0x61, 0x67, 0x65,
	0x53, 0x75, 0x62, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x44, 0x74, 0x6f, 0x52, 0x08,
	0x63, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x44, 0x61, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x44, 0x61, 0x6d, 0x61, 0x67, 0x65, 0x22, 0xb8, 0x01, 0x0a, 0x1a, 0x52,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x44, 0x61, 0x6d, 0x61, 0x67, 0x65, 0x53, 0x75, 0x62, 0x43, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x44, 0x74, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x34, 0x0a,
	0x06, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x44, 0x61, 0x6d,
	0x61, 0x67, 0x65, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x44, 0x74, 0x6f, 0x52, 0x06, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x6f, 0x69, 0x6e,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x6f,
	0x69, 0x6e, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x44, 0x61, 0x6d, 0x61,
	0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x44,
	0x61, 0x6d, 0x61, 0x67, 0x65, 0x22, 0x9c, 0x02, 0x0a, 0x14, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x44, 0x61, 0x6d, 0x61, 0x67, 0x65, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x44, 0x74, 0x6f, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x73,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x73, 0x12, 0x1a,
	0x0a, 0x08, 0x69, 0x73, 0x50, 0x61, 0x73, 0x73, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x08, 0x69, 0x73, 0x50, 0x61, 0x73, 0x73, 0x65, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x73,
	0x4e, 0x41, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x69, 0x73, 0x4e, 0x41, 0x12, 0x1a,
	0x0a, 0x08, 0x74, 0x61, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x74, 0x61, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65,
	0x6d, 0x61, 0x72, 0x6b, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x6d, 0x61,
	0x72, 0x6b, 0x12, 0x22, 0x0a, 0x0c, 0x64, 0x61, 0x6d, 0x61, 0x67, 0x65, 0x50, 0x68, 0x6f, 0x74,
	0x6f, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x61, 0x6d, 0x61, 0x67, 0x65,
	0x50, 0x68, 0x6f, 0x74, 0x6f, 0x73, 0x12, 0x2c, 0x0a, 0x11, 0x72, 0x65, 0x63, 0x6f, 0x6e, 0x64,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x11, 0x72, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x68,
	0x6f, 0x74, 0x6f, 0x73, 0x22, 0x80, 0x01, 0x0a, 0x13, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x65,
	0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x74, 0x6f, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65,
	0x6d, 0x61, 0x72, 0x6b, 0x12, 0x41, 0x0a, 0x0f, 0x72, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e,
	0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0f, 0x72, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x22, 0x34, 0x0a, 0x0e, 0x44, 0x61, 0x6d, 0x61, 0x67,
	0x65, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x44, 0x74, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x72, 0x6c,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x75, 0x72, 0x6c, 0x73, 0x2a, 0x32, 0x0a,
	0x0f, 0x52, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x06, 0x0a, 0x02, 0x4e, 0x4f, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x46, 0x49, 0x58, 0x45,
	0x44, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x52, 0x45, 0x50, 0x4c, 0x41, 0x43, 0x45, 0x44, 0x10,
	0x02, 0x32, 0xfd, 0x01, 0x0a, 0x06, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x30, 0x0a, 0x0d,
	0x47, 0x65, 0x74, 0x44, 0x61, 0x6d, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0d, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x49, 0x64, 0x44, 0x74, 0x6f, 0x1a, 0x10, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43,
	0x0a, 0x12, 0x44, 0x6f, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x50, 0x6f,
	0x69, 0x6e, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x74,
	0x6f, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x10, 0x42, 0x69, 0x6e, 0x64, 0x44, 0x61, 0x6d, 0x61, 0x67,
	0x65, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x73, 0x12, 0x16, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x2e, 0x44, 0x61, 0x6d, 0x61, 0x67, 0x65, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x44, 0x74, 0x6f, 0x1a,
	0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x3e, 0x0a, 0x12, 0x55, 0x6e, 0x62, 0x69, 0x6e, 0x64, 0x44, 0x61, 0x6d, 0x61, 0x67,
	0x65, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x73, 0x12, 0x16, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x2e, 0x44, 0x61, 0x6d, 0x61, 0x67, 0x65, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x44, 0x74, 0x6f, 0x1a,
	0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x63, 0x61, 0x72, 0x67, 0x6f, 0x64, 0x2d, 0x62, 0x6a, 0x2f, 0x62, 0x32, 0x63, 0x2d, 0x63, 0x61,
	0x72, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_report_report_proto_rawDescOnce sync.Once
	file_report_report_proto_rawDescData = file_report_report_proto_rawDesc
)

func file_report_report_proto_rawDescGZIP() []byte {
	file_report_report_proto_rawDescOnce.Do(func() {
		file_report_report_proto_rawDescData = protoimpl.X.CompressGZIP(file_report_report_proto_rawDescData)
	})
	return file_report_report_proto_rawDescData
}

var file_report_report_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_report_report_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_report_report_proto_goTypes = []interface{}{
	(ReconditionType)(0),               // 0: report.ReconditionType
	(*ReportDamageCategoryDto)(nil),    // 1: report.ReportDamageCategoryDto
	(*ReportDamageSubCategoryDto)(nil), // 2: report.ReportDamageSubCategoryDto
	(*ReportDamagePointDto)(nil),       // 3: report.ReportDamagePointDto
	(*PointReconditionDto)(nil),        // 4: report.PointReconditionDto
	(*DamagePhotoDto)(nil),             // 5: report.DamagePhotoDto
	(*common.IdDto)(nil),               // 6: common.IdDto
	(*common.Response)(nil),            // 7: common.Response
}
var file_report_report_proto_depIdxs = []int32{
	2, // 0: report.ReportDamageCategoryDto.children:type_name -> report.ReportDamageSubCategoryDto
	3, // 1: report.ReportDamageSubCategoryDto.points:type_name -> report.ReportDamagePointDto
	0, // 2: report.PointReconditionDto.reconditionType:type_name -> report.ReconditionType
	6, // 3: report.Report.GetDamageInfo:input_type -> common.IdDto
	4, // 4: report.Report.DoPointRecondition:input_type -> report.PointReconditionDto
	5, // 5: report.Report.BindDamagePhotos:input_type -> report.DamagePhotoDto
	5, // 6: report.Report.UnbindDamagePhotos:input_type -> report.DamagePhotoDto
	7, // 7: report.Report.GetDamageInfo:output_type -> common.Response
	7, // 8: report.Report.DoPointRecondition:output_type -> common.Response
	7, // 9: report.Report.BindDamagePhotos:output_type -> common.Response
	7, // 10: report.Report.UnbindDamagePhotos:output_type -> common.Response
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_report_report_proto_init() }
func file_report_report_proto_init() {
	if File_report_report_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_report_report_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_report_report_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_report_report_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_report_report_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PointReconditionDto); i {
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
		file_report_report_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DamagePhotoDto); i {
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
			RawDescriptor: file_report_report_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_report_report_proto_goTypes,
		DependencyIndexes: file_report_report_proto_depIdxs,
		EnumInfos:         file_report_report_proto_enumTypes,
		MessageInfos:      file_report_report_proto_msgTypes,
	}.Build()
	File_report_report_proto = out.File
	file_report_report_proto_rawDesc = nil
	file_report_report_proto_goTypes = nil
	file_report_report_proto_depIdxs = nil
}
