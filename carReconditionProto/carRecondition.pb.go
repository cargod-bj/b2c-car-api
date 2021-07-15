// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.14.0
// source: carReconditionProto/carRecondition.proto

package carReconditionProto

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

// 添加整备类型
type ReconditionTypePDto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	// 整备类别id
	CategoryId uint32 `protobuf:"varint,2,opt,name=categoryId,proto3" json:"categoryId"`
	// 整备类别Name
	CategoryName string `protobuf:"bytes,3,opt,name=categoryName,proto3" json:"categoryName"`
	// 详细描述
	ReconditionDetail string `protobuf:"bytes,4,opt,name=reconditionDetail,proto3" json:"reconditionDetail"`
	// 更新时间
	UpdateTime uint64 `protobuf:"varint,5,opt,name=updateTime,proto3" json:"updateTime"`
	// 创建时间
	CreateTime uint64 `protobuf:"varint,6,opt,name=createTime,proto3" json:"createTime"`
	// 最后操作人员id
	LastOptPic uint64 `protobuf:"varint,7,opt,name=lastOptPic,proto3" json:"lastOptPic"`
}

func (x *ReconditionTypePDto) Reset() {
	*x = ReconditionTypePDto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_carReconditionProto_carRecondition_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReconditionTypePDto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReconditionTypePDto) ProtoMessage() {}

func (x *ReconditionTypePDto) ProtoReflect() protoreflect.Message {
	mi := &file_carReconditionProto_carRecondition_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReconditionTypePDto.ProtoReflect.Descriptor instead.
func (*ReconditionTypePDto) Descriptor() ([]byte, []int) {
	return file_carReconditionProto_carRecondition_proto_rawDescGZIP(), []int{0}
}

func (x *ReconditionTypePDto) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ReconditionTypePDto) GetCategoryId() uint32 {
	if x != nil {
		return x.CategoryId
	}
	return 0
}

func (x *ReconditionTypePDto) GetCategoryName() string {
	if x != nil {
		return x.CategoryName
	}
	return ""
}

func (x *ReconditionTypePDto) GetReconditionDetail() string {
	if x != nil {
		return x.ReconditionDetail
	}
	return ""
}

func (x *ReconditionTypePDto) GetUpdateTime() uint64 {
	if x != nil {
		return x.UpdateTime
	}
	return 0
}

func (x *ReconditionTypePDto) GetCreateTime() uint64 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

func (x *ReconditionTypePDto) GetLastOptPic() uint64 {
	if x != nil {
		return x.LastOptPic
	}
	return 0
}

type ReconditionCategoryPDto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 整备类别id
	CategoryId uint32 `protobuf:"varint,1,opt,name=categoryId,proto3" json:"categoryId"`
	// 整备类别Name
	CategoryName string `protobuf:"bytes,2,opt,name=categoryName,proto3" json:"categoryName"`
	// 当前类型下的detail列表
	Details []*ReconditionTypePDto `protobuf:"bytes,3,rep,name=details,proto3" json:"details"`
}

func (x *ReconditionCategoryPDto) Reset() {
	*x = ReconditionCategoryPDto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_carReconditionProto_carRecondition_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReconditionCategoryPDto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReconditionCategoryPDto) ProtoMessage() {}

func (x *ReconditionCategoryPDto) ProtoReflect() protoreflect.Message {
	mi := &file_carReconditionProto_carRecondition_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReconditionCategoryPDto.ProtoReflect.Descriptor instead.
func (*ReconditionCategoryPDto) Descriptor() ([]byte, []int) {
	return file_carReconditionProto_carRecondition_proto_rawDescGZIP(), []int{1}
}

func (x *ReconditionCategoryPDto) GetCategoryId() uint32 {
	if x != nil {
		return x.CategoryId
	}
	return 0
}

func (x *ReconditionCategoryPDto) GetCategoryName() string {
	if x != nil {
		return x.CategoryName
	}
	return ""
}

func (x *ReconditionCategoryPDto) GetDetails() []*ReconditionTypePDto {
	if x != nil {
		return x.Details
	}
	return nil
}

// 添加整备类型
type AddReconditionBatchPDto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 整备类别id
	CategoryId uint32 `protobuf:"varint,2,opt,name=categoryId,proto3" json:"categoryId"`
	// 详细描述
	Details []string `protobuf:"bytes,3,rep,name=details,proto3" json:"details"`
	// 最后操作人员id
	LastOptPic uint64 `protobuf:"varint,6,opt,name=lastOptPic,proto3" json:"lastOptPic"`
}

func (x *AddReconditionBatchPDto) Reset() {
	*x = AddReconditionBatchPDto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_carReconditionProto_carRecondition_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddReconditionBatchPDto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddReconditionBatchPDto) ProtoMessage() {}

func (x *AddReconditionBatchPDto) ProtoReflect() protoreflect.Message {
	mi := &file_carReconditionProto_carRecondition_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddReconditionBatchPDto.ProtoReflect.Descriptor instead.
func (*AddReconditionBatchPDto) Descriptor() ([]byte, []int) {
	return file_carReconditionProto_carRecondition_proto_rawDescGZIP(), []int{2}
}

func (x *AddReconditionBatchPDto) GetCategoryId() uint32 {
	if x != nil {
		return x.CategoryId
	}
	return 0
}

func (x *AddReconditionBatchPDto) GetDetails() []string {
	if x != nil {
		return x.Details
	}
	return nil
}

func (x *AddReconditionBatchPDto) GetLastOptPic() uint64 {
	if x != nil {
		return x.LastOptPic
	}
	return 0
}

// id dto
type ReconditionTypeIdPDto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	// 执行当前操作的用户id
	OperatorId uint64 `protobuf:"varint,2,opt,name=operatorId,proto3" json:"operatorId"`
}

func (x *ReconditionTypeIdPDto) Reset() {
	*x = ReconditionTypeIdPDto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_carReconditionProto_carRecondition_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReconditionTypeIdPDto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReconditionTypeIdPDto) ProtoMessage() {}

func (x *ReconditionTypeIdPDto) ProtoReflect() protoreflect.Message {
	mi := &file_carReconditionProto_carRecondition_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReconditionTypeIdPDto.ProtoReflect.Descriptor instead.
func (*ReconditionTypeIdPDto) Descriptor() ([]byte, []int) {
	return file_carReconditionProto_carRecondition_proto_rawDescGZIP(), []int{3}
}

func (x *ReconditionTypeIdPDto) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ReconditionTypeIdPDto) GetOperatorId() uint64 {
	if x != nil {
		return x.OperatorId
	}
	return 0
}

// 获取车辆整备类型列表
type ReconditionTypeListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 车辆整备类别id
	CategoryIds []uint64 `protobuf:"varint,1,rep,packed,name=categoryIds,proto3" json:"categoryIds"`
	// 车辆整备详情
	Detail string `protobuf:"bytes,2,opt,name=detail,proto3" json:"detail"`
	// 分页信息
	Page *common.Page `protobuf:"bytes,101,opt,name=page,proto3" json:"page"`
}

func (x *ReconditionTypeListReq) Reset() {
	*x = ReconditionTypeListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_carReconditionProto_carRecondition_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReconditionTypeListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReconditionTypeListReq) ProtoMessage() {}

func (x *ReconditionTypeListReq) ProtoReflect() protoreflect.Message {
	mi := &file_carReconditionProto_carRecondition_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReconditionTypeListReq.ProtoReflect.Descriptor instead.
func (*ReconditionTypeListReq) Descriptor() ([]byte, []int) {
	return file_carReconditionProto_carRecondition_proto_rawDescGZIP(), []int{4}
}

func (x *ReconditionTypeListReq) GetCategoryIds() []uint64 {
	if x != nil {
		return x.CategoryIds
	}
	return nil
}

func (x *ReconditionTypeListReq) GetDetail() string {
	if x != nil {
		return x.Detail
	}
	return ""
}

func (x *ReconditionTypeListReq) GetPage() *common.Page {
	if x != nil {
		return x.Page
	}
	return nil
}

// 添加整备类型
type ReconditionWorkshopPDto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	// 门店名称
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name"`
	// 门店所属州
	StateCode string `protobuf:"bytes,3,opt,name=stateCode,proto3" json:"stateCode"`
	// 门店所属城市
	AreaCode string `protobuf:"bytes,4,opt,name=areaCode,proto3" json:"areaCode"`
	// 门店详细地址
	Address string `protobuf:"bytes,5,opt,name=address,proto3" json:"address"`
	// 更新时间
	UpdateTime uint64 `protobuf:"varint,6,opt,name=updateTime,proto3" json:"updateTime"`
	// 创建时间
	CreateTime uint64 `protobuf:"varint,7,opt,name=createTime,proto3" json:"createTime"`
	// 最后操作人员id
	LastOptPic uint64 `protobuf:"varint,8,opt,name=lastOptPic,proto3" json:"lastOptPic"`
	// 门店提供的具体整备服务数量
	ReconditionCount uint64 `protobuf:"varint,9,opt,name=reconditionCount,proto3" json:"reconditionCount"`
	// 门店提供的具体整备id
	ReconditionDetailIds []uint64 `protobuf:"varint,10,rep,packed,name=reconditionDetailIds,proto3" json:"reconditionDetailIds"`
}

func (x *ReconditionWorkshopPDto) Reset() {
	*x = ReconditionWorkshopPDto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_carReconditionProto_carRecondition_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReconditionWorkshopPDto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReconditionWorkshopPDto) ProtoMessage() {}

func (x *ReconditionWorkshopPDto) ProtoReflect() protoreflect.Message {
	mi := &file_carReconditionProto_carRecondition_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReconditionWorkshopPDto.ProtoReflect.Descriptor instead.
func (*ReconditionWorkshopPDto) Descriptor() ([]byte, []int) {
	return file_carReconditionProto_carRecondition_proto_rawDescGZIP(), []int{5}
}

func (x *ReconditionWorkshopPDto) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ReconditionWorkshopPDto) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ReconditionWorkshopPDto) GetStateCode() string {
	if x != nil {
		return x.StateCode
	}
	return ""
}

func (x *ReconditionWorkshopPDto) GetAreaCode() string {
	if x != nil {
		return x.AreaCode
	}
	return ""
}

func (x *ReconditionWorkshopPDto) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *ReconditionWorkshopPDto) GetUpdateTime() uint64 {
	if x != nil {
		return x.UpdateTime
	}
	return 0
}

func (x *ReconditionWorkshopPDto) GetCreateTime() uint64 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

func (x *ReconditionWorkshopPDto) GetLastOptPic() uint64 {
	if x != nil {
		return x.LastOptPic
	}
	return 0
}

func (x *ReconditionWorkshopPDto) GetReconditionCount() uint64 {
	if x != nil {
		return x.ReconditionCount
	}
	return 0
}

func (x *ReconditionWorkshopPDto) GetReconditionDetailIds() []uint64 {
	if x != nil {
		return x.ReconditionDetailIds
	}
	return nil
}

// 获取整备门店列表
type ReconditionWorkshopListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 整备门店名称
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name"`
	// 分页信息
	Page *common.Page `protobuf:"bytes,101,opt,name=page,proto3" json:"page"`
}

func (x *ReconditionWorkshopListReq) Reset() {
	*x = ReconditionWorkshopListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_carReconditionProto_carRecondition_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReconditionWorkshopListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReconditionWorkshopListReq) ProtoMessage() {}

func (x *ReconditionWorkshopListReq) ProtoReflect() protoreflect.Message {
	mi := &file_carReconditionProto_carRecondition_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReconditionWorkshopListReq.ProtoReflect.Descriptor instead.
func (*ReconditionWorkshopListReq) Descriptor() ([]byte, []int) {
	return file_carReconditionProto_carRecondition_proto_rawDescGZIP(), []int{6}
}

func (x *ReconditionWorkshopListReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ReconditionWorkshopListReq) GetPage() *common.Page {
	if x != nil {
		return x.Page
	}
	return nil
}

var File_carReconditionProto_carRecondition_proto protoreflect.FileDescriptor

var file_carReconditionProto_carRecondition_proto_rawDesc = []byte{
	0x0a, 0x28, 0x63, 0x61, 0x72, 0x52, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x61, 0x72, 0x52, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x63, 0x61, 0x72, 0x43,
	0x6f, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x61, 0x72, 0x67, 0x6f, 0x64, 0x2d, 0x62, 0x6a, 0x2f, 0x62,
	0x32, 0x63, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xf7, 0x01, 0x0a, 0x13, 0x52, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x50, 0x44, 0x74, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x63,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x63,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x2c, 0x0a, 0x11, 0x72, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x72, 0x65, 0x63, 0x6f,
	0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x1e, 0x0a,
	0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1e, 0x0a,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1e, 0x0a,
	0x0a, 0x6c, 0x61, 0x73, 0x74, 0x4f, 0x70, 0x74, 0x50, 0x69, 0x63, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x0a, 0x6c, 0x61, 0x73, 0x74, 0x4f, 0x70, 0x74, 0x50, 0x69, 0x63, 0x22, 0x9a, 0x01,
	0x0a, 0x17, 0x52, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x50, 0x44, 0x74, 0x6f, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x63,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x3b, 0x0a,
	0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21,
	0x2e, 0x63, 0x61, 0x72, 0x43, 0x6f, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65,
	0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x50, 0x44, 0x74,
	0x6f, 0x52, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x22, 0x73, 0x0a, 0x17, 0x41, 0x64,
	0x64, 0x52, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x61, 0x74, 0x63,
	0x68, 0x50, 0x44, 0x74, 0x6f, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12,
	0x1e, 0x0a, 0x0a, 0x6c, 0x61, 0x73, 0x74, 0x4f, 0x70, 0x74, 0x50, 0x69, 0x63, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x0a, 0x6c, 0x61, 0x73, 0x74, 0x4f, 0x70, 0x74, 0x50, 0x69, 0x63, 0x22,
	0x47, 0x0a, 0x15, 0x52, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79,
	0x70, 0x65, 0x49, 0x64, 0x50, 0x44, 0x74, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x6f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x22, 0x74, 0x0a, 0x16, 0x52, 0x65, 0x63, 0x6f,
	0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x04, 0x52, 0x0b, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x49, 0x64, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x20, 0x0a, 0x04,
	0x70, 0x61, 0x67, 0x65, 0x18, 0x65, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x22, 0xd1,
	0x02, 0x0a, 0x17, 0x52, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x57, 0x6f,
	0x72, 0x6b, 0x73, 0x68, 0x6f, 0x70, 0x50, 0x44, 0x74, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c,
	0x0a, 0x09, 0x73, 0x74, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x61, 0x72, 0x65, 0x61, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x61, 0x72, 0x65, 0x61, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x6c, 0x61, 0x73, 0x74, 0x4f, 0x70, 0x74, 0x50, 0x69, 0x63,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x6c, 0x61, 0x73, 0x74, 0x4f, 0x70, 0x74, 0x50,
	0x69, 0x63, 0x12, 0x2a, 0x0a, 0x10, 0x72, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x04, 0x52, 0x10, 0x72, 0x65,
	0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x32,
	0x0a, 0x14, 0x72, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x49, 0x64, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x04, 0x52, 0x14, 0x72, 0x65,
	0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x49,
	0x64, 0x73, 0x22, 0x52, 0x0a, 0x1a, 0x52, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x68, 0x6f, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x65, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65,
	0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x32, 0xbb, 0x07, 0x0a, 0x0e, 0x43, 0x61, 0x72, 0x52, 0x65,
	0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x4b, 0x0a, 0x12, 0x41, 0x64, 0x64,
	0x52, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x21, 0x2e, 0x63, 0x61, 0x72, 0x43, 0x6f, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52,
	0x65, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x50, 0x44,
	0x74, 0x6f, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x54, 0x0a, 0x17, 0x41, 0x64, 0x64, 0x52, 0x65, 0x63,
	0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x42, 0x61, 0x74, 0x63,
	0x68, 0x12, 0x25, 0x2e, 0x63, 0x61, 0x72, 0x43, 0x6f, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x41, 0x64, 0x64, 0x52, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x42,
	0x61, 0x74, 0x63, 0x68, 0x50, 0x44, 0x74, 0x6f, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4e, 0x0a, 0x15,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x21, 0x2e, 0x63, 0x61, 0x72, 0x43, 0x6f, 0x73, 0x74, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x54, 0x79, 0x70, 0x65, 0x50, 0x44, 0x74, 0x6f, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x50, 0x0a, 0x15,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x23, 0x2e, 0x63, 0x61, 0x72, 0x43, 0x6f, 0x73, 0x74, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x54, 0x79, 0x70, 0x65, 0x49, 0x64, 0x50, 0x44, 0x74, 0x6f, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4f,
	0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x24, 0x2e, 0x63, 0x61, 0x72, 0x43, 0x6f, 0x73, 0x74, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x54, 0x79, 0x70, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x10, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x5d, 0x0a, 0x22, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x63, 0x6f, 0x6e, 0x64,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x57, 0x69, 0x74, 0x68, 0x43, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x23, 0x2e, 0x63, 0x61, 0x72, 0x43, 0x6f, 0x73, 0x74, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x54, 0x79, 0x70, 0x65, 0x49, 0x64, 0x50, 0x44, 0x74, 0x6f, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x58,
	0x0a, 0x1d, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x54, 0x79, 0x70, 0x65, 0x42, 0x79, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x68, 0x6f, 0x70, 0x12,
	0x23, 0x2e, 0x63, 0x61, 0x72, 0x43, 0x6f, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52,
	0x65, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x49, 0x64,
	0x50, 0x44, 0x74, 0x6f, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x53, 0x0a, 0x16, 0x41, 0x64, 0x64, 0x52,
	0x65, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x68,
	0x6f, 0x70, 0x12, 0x25, 0x2e, 0x63, 0x61, 0x72, 0x43, 0x6f, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x57, 0x6f, 0x72,
	0x6b, 0x73, 0x68, 0x6f, 0x70, 0x50, 0x44, 0x74, 0x6f, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x56, 0x0a,
	0x19, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x68, 0x6f, 0x70, 0x12, 0x25, 0x2e, 0x63, 0x61, 0x72,
	0x43, 0x6f, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x6e, 0x64,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x68, 0x6f, 0x70, 0x50, 0x44, 0x74,
	0x6f, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x54, 0x0a, 0x19, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52,
	0x65, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x68,
	0x6f, 0x70, 0x12, 0x23, 0x2e, 0x63, 0x61, 0x72, 0x43, 0x6f, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70,
	0x65, 0x49, 0x64, 0x50, 0x44, 0x74, 0x6f, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x57, 0x0a, 0x17, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x57, 0x6f,
	0x72, 0x6b, 0x73, 0x68, 0x6f, 0x70, 0x12, 0x28, 0x2e, 0x63, 0x61, 0x72, 0x43, 0x6f, 0x73, 0x74,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x68, 0x6f, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x36, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x63, 0x61, 0x72, 0x67, 0x6f, 0x64, 0x2d, 0x62, 0x6a, 0x2f, 0x62, 0x32, 0x63,
	0x2d, 0x63, 0x61, 0x72, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x61, 0x72, 0x52, 0x65, 0x63, 0x6f,
	0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_carReconditionProto_carRecondition_proto_rawDescOnce sync.Once
	file_carReconditionProto_carRecondition_proto_rawDescData = file_carReconditionProto_carRecondition_proto_rawDesc
)

func file_carReconditionProto_carRecondition_proto_rawDescGZIP() []byte {
	file_carReconditionProto_carRecondition_proto_rawDescOnce.Do(func() {
		file_carReconditionProto_carRecondition_proto_rawDescData = protoimpl.X.CompressGZIP(file_carReconditionProto_carRecondition_proto_rawDescData)
	})
	return file_carReconditionProto_carRecondition_proto_rawDescData
}

var file_carReconditionProto_carRecondition_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_carReconditionProto_carRecondition_proto_goTypes = []interface{}{
	(*ReconditionTypePDto)(nil),        // 0: carCostProto.ReconditionTypePDto
	(*ReconditionCategoryPDto)(nil),    // 1: carCostProto.ReconditionCategoryPDto
	(*AddReconditionBatchPDto)(nil),    // 2: carCostProto.AddReconditionBatchPDto
	(*ReconditionTypeIdPDto)(nil),      // 3: carCostProto.ReconditionTypeIdPDto
	(*ReconditionTypeListReq)(nil),     // 4: carCostProto.ReconditionTypeListReq
	(*ReconditionWorkshopPDto)(nil),    // 5: carCostProto.ReconditionWorkshopPDto
	(*ReconditionWorkshopListReq)(nil), // 6: carCostProto.ReconditionWorkshopListReq
	(*common.Page)(nil),                // 7: common.Page
	(*common.Response)(nil),            // 8: common.Response
}
var file_carReconditionProto_carRecondition_proto_depIdxs = []int32{
	0,  // 0: carCostProto.ReconditionCategoryPDto.details:type_name -> carCostProto.ReconditionTypePDto
	7,  // 1: carCostProto.ReconditionTypeListReq.page:type_name -> common.Page
	7,  // 2: carCostProto.ReconditionWorkshopListReq.page:type_name -> common.Page
	0,  // 3: carCostProto.CarRecondition.AddReconditionType:input_type -> carCostProto.ReconditionTypePDto
	2,  // 4: carCostProto.CarRecondition.AddReconditionTypeBatch:input_type -> carCostProto.AddReconditionBatchPDto
	0,  // 5: carCostProto.CarRecondition.UpdateReconditionType:input_type -> carCostProto.ReconditionTypePDto
	3,  // 6: carCostProto.CarRecondition.DeleteReconditionType:input_type -> carCostProto.ReconditionTypeIdPDto
	4,  // 7: carCostProto.CarRecondition.ListReconditionType:input_type -> carCostProto.ReconditionTypeListReq
	3,  // 8: carCostProto.CarRecondition.ListAllReconditionTypeWithCategory:input_type -> carCostProto.ReconditionTypeIdPDto
	3,  // 9: carCostProto.CarRecondition.ListReconditionTypeByWorkshop:input_type -> carCostProto.ReconditionTypeIdPDto
	5,  // 10: carCostProto.CarRecondition.AddReconditionWorkshop:input_type -> carCostProto.ReconditionWorkshopPDto
	5,  // 11: carCostProto.CarRecondition.UpdateReconditionWorkshop:input_type -> carCostProto.ReconditionWorkshopPDto
	3,  // 12: carCostProto.CarRecondition.DeleteReconditionWorkshop:input_type -> carCostProto.ReconditionTypeIdPDto
	6,  // 13: carCostProto.CarRecondition.ListReconditionWorkshop:input_type -> carCostProto.ReconditionWorkshopListReq
	8,  // 14: carCostProto.CarRecondition.AddReconditionType:output_type -> common.Response
	8,  // 15: carCostProto.CarRecondition.AddReconditionTypeBatch:output_type -> common.Response
	8,  // 16: carCostProto.CarRecondition.UpdateReconditionType:output_type -> common.Response
	8,  // 17: carCostProto.CarRecondition.DeleteReconditionType:output_type -> common.Response
	8,  // 18: carCostProto.CarRecondition.ListReconditionType:output_type -> common.Response
	8,  // 19: carCostProto.CarRecondition.ListAllReconditionTypeWithCategory:output_type -> common.Response
	8,  // 20: carCostProto.CarRecondition.ListReconditionTypeByWorkshop:output_type -> common.Response
	8,  // 21: carCostProto.CarRecondition.AddReconditionWorkshop:output_type -> common.Response
	8,  // 22: carCostProto.CarRecondition.UpdateReconditionWorkshop:output_type -> common.Response
	8,  // 23: carCostProto.CarRecondition.DeleteReconditionWorkshop:output_type -> common.Response
	8,  // 24: carCostProto.CarRecondition.ListReconditionWorkshop:output_type -> common.Response
	14, // [14:25] is the sub-list for method output_type
	3,  // [3:14] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_carReconditionProto_carRecondition_proto_init() }
func file_carReconditionProto_carRecondition_proto_init() {
	if File_carReconditionProto_carRecondition_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_carReconditionProto_carRecondition_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReconditionTypePDto); i {
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
		file_carReconditionProto_carRecondition_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReconditionCategoryPDto); i {
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
		file_carReconditionProto_carRecondition_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddReconditionBatchPDto); i {
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
		file_carReconditionProto_carRecondition_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReconditionTypeIdPDto); i {
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
		file_carReconditionProto_carRecondition_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReconditionTypeListReq); i {
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
		file_carReconditionProto_carRecondition_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReconditionWorkshopPDto); i {
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
		file_carReconditionProto_carRecondition_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReconditionWorkshopListReq); i {
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
			RawDescriptor: file_carReconditionProto_carRecondition_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_carReconditionProto_carRecondition_proto_goTypes,
		DependencyIndexes: file_carReconditionProto_carRecondition_proto_depIdxs,
		MessageInfos:      file_carReconditionProto_carRecondition_proto_msgTypes,
	}.Build()
	File_carReconditionProto_carRecondition_proto = out.File
	file_carReconditionProto_carRecondition_proto_rawDesc = nil
	file_carReconditionProto_carRecondition_proto_goTypes = nil
	file_carReconditionProto_carRecondition_proto_depIdxs = nil
}
