// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: purchasePaymentApproval/purchasePaymentApproval.proto

package purchasePaymentApproval

import (
	common "github.com/cargod-bj/b2c-proto-common/common"
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

type ApprovalScanReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ApprovalScanReq) Reset() {
	*x = ApprovalScanReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_purchasePaymentApproval_purchasePaymentApproval_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApprovalScanReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApprovalScanReq) ProtoMessage() {}

func (x *ApprovalScanReq) ProtoReflect() protoreflect.Message {
	mi := &file_purchasePaymentApproval_purchasePaymentApproval_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApprovalScanReq.ProtoReflect.Descriptor instead.
func (*ApprovalScanReq) Descriptor() ([]byte, []int) {
	return file_purchasePaymentApproval_purchasePaymentApproval_proto_rawDescGZIP(), []int{0}
}

type ApprovalApprovalCondition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CarNo             string `protobuf:"bytes,1,opt,name=carNo,proto3" json:"carNo,omitempty"`
	PlateNumber       string `protobuf:"bytes,2,opt,name=plateNumber,proto3" json:"plateNumber,omitempty"`
	InspectionId      uint64 `protobuf:"varint,3,opt,name=inspectionId,proto3" json:"inspectionId,omitempty"`
	PurchaseNo        string `protobuf:"bytes,4,opt,name=purchaseNo,proto3" json:"purchaseNo,omitempty"`
	DtStatus          uint64 `protobuf:"varint,5,opt,name=dtStatus,proto3" json:"dtStatus,omitempty"`
	ApprovalStartDate uint64 `protobuf:"varint,6,opt,name=approvalStartDate,proto3" json:"approvalStartDate,omitempty"`
	ApprovalEndDate   uint64 `protobuf:"varint,7,opt,name=approvalEndDate,proto3" json:"approvalEndDate,omitempty"`
	PageNo            uint32 `protobuf:"varint,8,opt,name=pageNo,proto3" json:"pageNo,omitempty"`
	PageSize          uint32 `protobuf:"varint,9,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
}

func (x *ApprovalApprovalCondition) Reset() {
	*x = ApprovalApprovalCondition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_purchasePaymentApproval_purchasePaymentApproval_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApprovalApprovalCondition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApprovalApprovalCondition) ProtoMessage() {}

func (x *ApprovalApprovalCondition) ProtoReflect() protoreflect.Message {
	mi := &file_purchasePaymentApproval_purchasePaymentApproval_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApprovalApprovalCondition.ProtoReflect.Descriptor instead.
func (*ApprovalApprovalCondition) Descriptor() ([]byte, []int) {
	return file_purchasePaymentApproval_purchasePaymentApproval_proto_rawDescGZIP(), []int{1}
}

func (x *ApprovalApprovalCondition) GetCarNo() string {
	if x != nil {
		return x.CarNo
	}
	return ""
}

func (x *ApprovalApprovalCondition) GetPlateNumber() string {
	if x != nil {
		return x.PlateNumber
	}
	return ""
}

func (x *ApprovalApprovalCondition) GetInspectionId() uint64 {
	if x != nil {
		return x.InspectionId
	}
	return 0
}

func (x *ApprovalApprovalCondition) GetPurchaseNo() string {
	if x != nil {
		return x.PurchaseNo
	}
	return ""
}

func (x *ApprovalApprovalCondition) GetDtStatus() uint64 {
	if x != nil {
		return x.DtStatus
	}
	return 0
}

func (x *ApprovalApprovalCondition) GetApprovalStartDate() uint64 {
	if x != nil {
		return x.ApprovalStartDate
	}
	return 0
}

func (x *ApprovalApprovalCondition) GetApprovalEndDate() uint64 {
	if x != nil {
		return x.ApprovalEndDate
	}
	return 0
}

func (x *ApprovalApprovalCondition) GetPageNo() uint32 {
	if x != nil {
		return x.PageNo
	}
	return 0
}

func (x *ApprovalApprovalCondition) GetPageSize() uint32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type ApprovalDto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"` // 审批编号
	PurchaseNo        string `protobuf:"bytes,2,opt,name=purchaseNo,proto3" json:"purchaseNo,omitempty"`
	CarNo             string `protobuf:"bytes,3,opt,name=carNo,proto3" json:"carNo,omitempty"`
	InspectionId      uint64 `protobuf:"varint,4,opt,name=inspectionId,proto3" json:"inspectionId,omitempty"`
	CarModel          string `protobuf:"bytes,5,opt,name=carModel,proto3" json:"carModel,omitempty"`
	PlateNumber       string `protobuf:"bytes,6,opt,name=plateNumber,proto3" json:"plateNumber,omitempty"`
	DtStatus          uint64 `protobuf:"varint,7,opt,name=dtStatus,proto3" json:"dtStatus,omitempty"`
	ApprovalStartDate uint64 `protobuf:"varint,8,opt,name=approvalStartDate,proto3" json:"approvalStartDate,omitempty"`
	ApprovalEndDate   uint64 `protobuf:"varint,9,opt,name=approvalEndDate,proto3" json:"approvalEndDate,omitempty"`
	PoStatus          uint64 `protobuf:"varint,10,opt,name=poStatus,proto3" json:"poStatus,omitempty"`
	UploadResult      uint64 `protobuf:"varint,11,opt,name=uploadResult,proto3" json:"uploadResult,omitempty"`
	PurchaseId        uint64 `protobuf:"varint,12,opt,name=purchaseId,proto3" json:"purchaseId,omitempty"`
	Year              uint64 `protobuf:"varint,13,opt,name=year,proto3" json:"year,omitempty"`
	BrandId           uint64 `protobuf:"varint,14,opt,name=brandId,proto3" json:"brandId,omitempty"`
	ModelId           uint64 `protobuf:"varint,15,opt,name=modelId,proto3" json:"modelId,omitempty"`
	VariantId         uint64 `protobuf:"varint,16,opt,name=variantId,proto3" json:"variantId,omitempty"`
	CarEngine         string `protobuf:"bytes,17,opt,name=carEngine,proto3" json:"carEngine,omitempty"`
	CarTransmission   uint64 `protobuf:"varint,18,opt,name=carTransmission,proto3" json:"carTransmission,omitempty"`
}

func (x *ApprovalDto) Reset() {
	*x = ApprovalDto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_purchasePaymentApproval_purchasePaymentApproval_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApprovalDto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApprovalDto) ProtoMessage() {}

func (x *ApprovalDto) ProtoReflect() protoreflect.Message {
	mi := &file_purchasePaymentApproval_purchasePaymentApproval_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApprovalDto.ProtoReflect.Descriptor instead.
func (*ApprovalDto) Descriptor() ([]byte, []int) {
	return file_purchasePaymentApproval_purchasePaymentApproval_proto_rawDescGZIP(), []int{2}
}

func (x *ApprovalDto) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ApprovalDto) GetPurchaseNo() string {
	if x != nil {
		return x.PurchaseNo
	}
	return ""
}

func (x *ApprovalDto) GetCarNo() string {
	if x != nil {
		return x.CarNo
	}
	return ""
}

func (x *ApprovalDto) GetInspectionId() uint64 {
	if x != nil {
		return x.InspectionId
	}
	return 0
}

func (x *ApprovalDto) GetCarModel() string {
	if x != nil {
		return x.CarModel
	}
	return ""
}

func (x *ApprovalDto) GetPlateNumber() string {
	if x != nil {
		return x.PlateNumber
	}
	return ""
}

func (x *ApprovalDto) GetDtStatus() uint64 {
	if x != nil {
		return x.DtStatus
	}
	return 0
}

func (x *ApprovalDto) GetApprovalStartDate() uint64 {
	if x != nil {
		return x.ApprovalStartDate
	}
	return 0
}

func (x *ApprovalDto) GetApprovalEndDate() uint64 {
	if x != nil {
		return x.ApprovalEndDate
	}
	return 0
}

func (x *ApprovalDto) GetPoStatus() uint64 {
	if x != nil {
		return x.PoStatus
	}
	return 0
}

func (x *ApprovalDto) GetUploadResult() uint64 {
	if x != nil {
		return x.UploadResult
	}
	return 0
}

func (x *ApprovalDto) GetPurchaseId() uint64 {
	if x != nil {
		return x.PurchaseId
	}
	return 0
}

func (x *ApprovalDto) GetYear() uint64 {
	if x != nil {
		return x.Year
	}
	return 0
}

func (x *ApprovalDto) GetBrandId() uint64 {
	if x != nil {
		return x.BrandId
	}
	return 0
}

func (x *ApprovalDto) GetModelId() uint64 {
	if x != nil {
		return x.ModelId
	}
	return 0
}

func (x *ApprovalDto) GetVariantId() uint64 {
	if x != nil {
		return x.VariantId
	}
	return 0
}

func (x *ApprovalDto) GetCarEngine() string {
	if x != nil {
		return x.CarEngine
	}
	return ""
}

func (x *ApprovalDto) GetCarTransmission() uint64 {
	if x != nil {
		return x.CarTransmission
	}
	return 0
}

type ApplyReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"` // 审批编号
	UserId string `protobuf:"bytes,2,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *ApplyReq) Reset() {
	*x = ApplyReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_purchasePaymentApproval_purchasePaymentApproval_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApplyReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApplyReq) ProtoMessage() {}

func (x *ApplyReq) ProtoReflect() protoreflect.Message {
	mi := &file_purchasePaymentApproval_purchasePaymentApproval_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApplyReq.ProtoReflect.Descriptor instead.
func (*ApplyReq) Descriptor() ([]byte, []int) {
	return file_purchasePaymentApproval_purchasePaymentApproval_proto_rawDescGZIP(), []int{3}
}

func (x *ApplyReq) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ApplyReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type IdWithUserReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"` // 审批编号
	UserId    string `protobuf:"bytes,2,opt,name=userId,proto3" json:"userId,omitempty"`
	SysUserId uint64 `protobuf:"varint,3,opt,name=sysUserId,proto3" json:"sysUserId,omitempty"`
}

func (x *IdWithUserReq) Reset() {
	*x = IdWithUserReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_purchasePaymentApproval_purchasePaymentApproval_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IdWithUserReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdWithUserReq) ProtoMessage() {}

func (x *IdWithUserReq) ProtoReflect() protoreflect.Message {
	mi := &file_purchasePaymentApproval_purchasePaymentApproval_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdWithUserReq.ProtoReflect.Descriptor instead.
func (*IdWithUserReq) Descriptor() ([]byte, []int) {
	return file_purchasePaymentApproval_purchasePaymentApproval_proto_rawDescGZIP(), []int{4}
}

func (x *IdWithUserReq) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *IdWithUserReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *IdWithUserReq) GetSysUserId() uint64 {
	if x != nil {
		return x.SysUserId
	}
	return 0
}

type DingTalkCallbackReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Signature string `protobuf:"bytes,1,opt,name=signature,proto3" json:"signature,omitempty"`
	Timestamp string `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Nonce     string `protobuf:"bytes,3,opt,name=nonce,proto3" json:"nonce,omitempty"`
	Encrypt   string `protobuf:"bytes,4,opt,name=encrypt,proto3" json:"encrypt,omitempty"`
}

func (x *DingTalkCallbackReq) Reset() {
	*x = DingTalkCallbackReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_purchasePaymentApproval_purchasePaymentApproval_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DingTalkCallbackReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DingTalkCallbackReq) ProtoMessage() {}

func (x *DingTalkCallbackReq) ProtoReflect() protoreflect.Message {
	mi := &file_purchasePaymentApproval_purchasePaymentApproval_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DingTalkCallbackReq.ProtoReflect.Descriptor instead.
func (*DingTalkCallbackReq) Descriptor() ([]byte, []int) {
	return file_purchasePaymentApproval_purchasePaymentApproval_proto_rawDescGZIP(), []int{5}
}

func (x *DingTalkCallbackReq) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

func (x *DingTalkCallbackReq) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

func (x *DingTalkCallbackReq) GetNonce() string {
	if x != nil {
		return x.Nonce
	}
	return ""
}

func (x *DingTalkCallbackReq) GetEncrypt() string {
	if x != nil {
		return x.Encrypt
	}
	return ""
}

type DingTalkCallbackRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MsgSignature string `protobuf:"bytes,1,opt,name=msg_signature,json=msgSignature,proto3" json:"msg_signature,omitempty"`
	TimeStamp    string `protobuf:"bytes,2,opt,name=timeStamp,proto3" json:"timeStamp,omitempty"`
	Nonce        string `protobuf:"bytes,3,opt,name=nonce,proto3" json:"nonce,omitempty"`
	Encrypt      string `protobuf:"bytes,4,opt,name=encrypt,proto3" json:"encrypt,omitempty"`
}

func (x *DingTalkCallbackRes) Reset() {
	*x = DingTalkCallbackRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_purchasePaymentApproval_purchasePaymentApproval_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DingTalkCallbackRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DingTalkCallbackRes) ProtoMessage() {}

func (x *DingTalkCallbackRes) ProtoReflect() protoreflect.Message {
	mi := &file_purchasePaymentApproval_purchasePaymentApproval_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DingTalkCallbackRes.ProtoReflect.Descriptor instead.
func (*DingTalkCallbackRes) Descriptor() ([]byte, []int) {
	return file_purchasePaymentApproval_purchasePaymentApproval_proto_rawDescGZIP(), []int{6}
}

func (x *DingTalkCallbackRes) GetMsgSignature() string {
	if x != nil {
		return x.MsgSignature
	}
	return ""
}

func (x *DingTalkCallbackRes) GetTimeStamp() string {
	if x != nil {
		return x.TimeStamp
	}
	return ""
}

func (x *DingTalkCallbackRes) GetNonce() string {
	if x != nil {
		return x.Nonce
	}
	return ""
}

func (x *DingTalkCallbackRes) GetEncrypt() string {
	if x != nil {
		return x.Encrypt
	}
	return ""
}

var File_purchasePaymentApproval_purchasePaymentApproval_proto protoreflect.FileDescriptor

var file_purchasePaymentApproval_purchasePaymentApproval_proto_rawDesc = []byte{
	0x0a, 0x35, 0x70, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x2f, 0x70, 0x75, 0x72, 0x63, 0x68, 0x61,
	0x73, 0x65, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61,
	0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x63, 0x61, 0x72, 0x1a, 0x39, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x61, 0x72, 0x67, 0x6f, 0x64, 0x2d,
	0x62, 0x6a, 0x2f, 0x62, 0x32, 0x63, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x11, 0x0a, 0x0f, 0x41, 0x70, 0x70, 0x72, 0x6f,
	0x76, 0x61, 0x6c, 0x53, 0x63, 0x61, 0x6e, 0x52, 0x65, 0x71, 0x22, 0xbf, 0x02, 0x0a, 0x19, 0x41,
	0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x43,
	0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x61, 0x72, 0x4e,
	0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x61, 0x72, 0x4e, 0x6f, 0x12, 0x20,
	0x0a, 0x0b, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x12, 0x22, 0x0a, 0x0c, 0x69, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x69, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65,
	0x4e, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x75, 0x72, 0x63, 0x68, 0x61,
	0x73, 0x65, 0x4e, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x64, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x2c, 0x0a, 0x11, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x53, 0x74, 0x61, 0x72,
	0x74, 0x44, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x11, 0x61, 0x70, 0x70,
	0x72, 0x6f, 0x76, 0x61, 0x6c, 0x53, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x28,
	0x0a, 0x0f, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x45, 0x6e, 0x64, 0x44, 0x61, 0x74,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0f, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61,
	0x6c, 0x45, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61, 0x67, 0x65,
	0x4e, 0x6f, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x6f,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x22, 0xb7, 0x04, 0x0a,
	0x0b, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x44, 0x74, 0x6f, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1e, 0x0a, 0x0a,
	0x70, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x4e, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x70, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x4e, 0x6f, 0x12, 0x14, 0x0a, 0x05,
	0x63, 0x61, 0x72, 0x4e, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x61, 0x72,
	0x4e, 0x6f, 0x12, 0x22, 0x0a, 0x0c, 0x69, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x69, 0x6e, 0x73, 0x70, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x72, 0x4d, 0x6f, 0x64,
	0x65, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x72, 0x4d, 0x6f, 0x64,
	0x65, 0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x64, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x2c, 0x0a, 0x11, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x53, 0x74, 0x61, 0x72,
	0x74, 0x44, 0x61, 0x74, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x04, 0x52, 0x11, 0x61, 0x70, 0x70,
	0x72, 0x6f, 0x76, 0x61, 0x6c, 0x53, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x28,
	0x0a, 0x0f, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x45, 0x6e, 0x64, 0x44, 0x61, 0x74,
	0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0f, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61,
	0x6c, 0x45, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6f, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x70, 0x6f, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x75, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x75, 0x72, 0x63,
	0x68, 0x61, 0x73, 0x65, 0x49, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x70, 0x75,
	0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x79, 0x65, 0x61, 0x72,
	0x18, 0x0d, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x79, 0x65, 0x61, 0x72, 0x12, 0x18, 0x0a, 0x07,
	0x62, 0x72, 0x61, 0x6e, 0x64, 0x49, 0x64, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x62,
	0x72, 0x61, 0x6e, 0x64, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x49,
	0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x49, 0x64,
	0x12, 0x1c, 0x0a, 0x09, 0x76, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x10, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x09, 0x76, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1c,
	0x0a, 0x09, 0x63, 0x61, 0x72, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x18, 0x11, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x63, 0x61, 0x72, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x12, 0x28, 0x0a, 0x0f,
	0x63, 0x61, 0x72, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x12, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0f, 0x63, 0x61, 0x72, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x32, 0x0a, 0x08, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x52,
	0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x55, 0x0a, 0x0d, 0x49, 0x64,
	0x57, 0x69, 0x74, 0x68, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x79, 0x73, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x73, 0x79, 0x73, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x22, 0x81, 0x01, 0x0a, 0x13, 0x44, 0x69, 0x6e, 0x67, 0x54, 0x61, 0x6c, 0x6b, 0x43, 0x61,
	0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67,
	0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x69,
	0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x65,
	0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6e,
	0x63, 0x72, 0x79, 0x70, 0x74, 0x22, 0x88, 0x01, 0x0a, 0x13, 0x44, 0x69, 0x6e, 0x67, 0x54, 0x61,
	0x6c, 0x6b, 0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x12, 0x23, 0x0a,
	0x0d, 0x6d, 0x73, 0x67, 0x5f, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6d, 0x73, 0x67, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x6d, 0x70, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x6d, 0x70,
	0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74,
	0x32, 0x9f, 0x02, 0x0a, 0x17, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x50, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x12, 0x30, 0x0a, 0x04,
	0x53, 0x63, 0x61, 0x6e, 0x12, 0x14, 0x2e, 0x63, 0x61, 0x72, 0x2e, 0x41, 0x70, 0x70, 0x72, 0x6f,
	0x76, 0x61, 0x6c, 0x53, 0x63, 0x61, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3a,
	0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1e, 0x2e, 0x63, 0x61, 0x72, 0x2e, 0x41, 0x70, 0x70,
	0x72, 0x6f, 0x76, 0x61, 0x6c, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x43, 0x6f, 0x6e,
	0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2a, 0x0a, 0x05, 0x41, 0x70,
	0x70, 0x6c, 0x79, 0x12, 0x0d, 0x2e, 0x63, 0x61, 0x72, 0x2e, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x52,
	0x65, 0x71, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x30, 0x0a, 0x06, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c,
	0x12, 0x12, 0x2e, 0x63, 0x61, 0x72, 0x2e, 0x49, 0x64, 0x57, 0x69, 0x74, 0x68, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x08, 0x43, 0x61, 0x6c, 0x6c,
	0x62, 0x61, 0x63, 0x6b, 0x12, 0x18, 0x2e, 0x63, 0x61, 0x72, 0x2e, 0x44, 0x69, 0x6e, 0x67, 0x54,
	0x61, 0x6c, 0x6b, 0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x1a, 0x10,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x3a, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x63, 0x61, 0x72, 0x67, 0x6f, 0x64, 0x2d, 0x62, 0x6a, 0x2f, 0x62, 0x32, 0x63, 0x2d, 0x63,
	0x61, 0x72, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x50,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_purchasePaymentApproval_purchasePaymentApproval_proto_rawDescOnce sync.Once
	file_purchasePaymentApproval_purchasePaymentApproval_proto_rawDescData = file_purchasePaymentApproval_purchasePaymentApproval_proto_rawDesc
)

func file_purchasePaymentApproval_purchasePaymentApproval_proto_rawDescGZIP() []byte {
	file_purchasePaymentApproval_purchasePaymentApproval_proto_rawDescOnce.Do(func() {
		file_purchasePaymentApproval_purchasePaymentApproval_proto_rawDescData = protoimpl.X.CompressGZIP(file_purchasePaymentApproval_purchasePaymentApproval_proto_rawDescData)
	})
	return file_purchasePaymentApproval_purchasePaymentApproval_proto_rawDescData
}

var file_purchasePaymentApproval_purchasePaymentApproval_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_purchasePaymentApproval_purchasePaymentApproval_proto_goTypes = []interface{}{
	(*ApprovalScanReq)(nil),           // 0: car.ApprovalScanReq
	(*ApprovalApprovalCondition)(nil), // 1: car.ApprovalApprovalCondition
	(*ApprovalDto)(nil),               // 2: car.ApprovalDto
	(*ApplyReq)(nil),                  // 3: car.ApplyReq
	(*IdWithUserReq)(nil),             // 4: car.IdWithUserReq
	(*DingTalkCallbackReq)(nil),       // 5: car.DingTalkCallbackReq
	(*DingTalkCallbackRes)(nil),       // 6: car.DingTalkCallbackRes
	(*common.Response)(nil),           // 7: common.Response
}
var file_purchasePaymentApproval_purchasePaymentApproval_proto_depIdxs = []int32{
	0, // 0: car.PurchasePaymentApproval.Scan:input_type -> car.ApprovalScanReq
	1, // 1: car.PurchasePaymentApproval.List:input_type -> car.ApprovalApprovalCondition
	3, // 2: car.PurchasePaymentApproval.Apply:input_type -> car.ApplyReq
	4, // 3: car.PurchasePaymentApproval.Cancel:input_type -> car.IdWithUserReq
	5, // 4: car.PurchasePaymentApproval.Callback:input_type -> car.DingTalkCallbackReq
	7, // 5: car.PurchasePaymentApproval.Scan:output_type -> common.Response
	7, // 6: car.PurchasePaymentApproval.List:output_type -> common.Response
	7, // 7: car.PurchasePaymentApproval.Apply:output_type -> common.Response
	7, // 8: car.PurchasePaymentApproval.Cancel:output_type -> common.Response
	7, // 9: car.PurchasePaymentApproval.Callback:output_type -> common.Response
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_purchasePaymentApproval_purchasePaymentApproval_proto_init() }
func file_purchasePaymentApproval_purchasePaymentApproval_proto_init() {
	if File_purchasePaymentApproval_purchasePaymentApproval_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_purchasePaymentApproval_purchasePaymentApproval_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApprovalScanReq); i {
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
		file_purchasePaymentApproval_purchasePaymentApproval_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApprovalApprovalCondition); i {
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
		file_purchasePaymentApproval_purchasePaymentApproval_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApprovalDto); i {
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
		file_purchasePaymentApproval_purchasePaymentApproval_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApplyReq); i {
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
		file_purchasePaymentApproval_purchasePaymentApproval_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IdWithUserReq); i {
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
		file_purchasePaymentApproval_purchasePaymentApproval_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DingTalkCallbackReq); i {
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
		file_purchasePaymentApproval_purchasePaymentApproval_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DingTalkCallbackRes); i {
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
			RawDescriptor: file_purchasePaymentApproval_purchasePaymentApproval_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_purchasePaymentApproval_purchasePaymentApproval_proto_goTypes,
		DependencyIndexes: file_purchasePaymentApproval_purchasePaymentApproval_proto_depIdxs,
		MessageInfos:      file_purchasePaymentApproval_purchasePaymentApproval_proto_msgTypes,
	}.Build()
	File_purchasePaymentApproval_purchasePaymentApproval_proto = out.File
	file_purchasePaymentApproval_purchasePaymentApproval_proto_rawDesc = nil
	file_purchasePaymentApproval_purchasePaymentApproval_proto_goTypes = nil
	file_purchasePaymentApproval_purchasePaymentApproval_proto_depIdxs = nil
}
