// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.14.0
// source: carLogProto/car_log.proto

package carLogProto

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

type CarLogReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	PageNo   uint32 `protobuf:"varint,2,opt,name=pageNo,proto3" json:"pageNo"`
	PageSize uint32 `protobuf:"varint,3,opt,name=pageSize,proto3" json:"pageSize"`
}

func (x *CarLogReq) Reset() {
	*x = CarLogReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_carLogProto_car_log_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CarLogReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CarLogReq) ProtoMessage() {}

func (x *CarLogReq) ProtoReflect() protoreflect.Message {
	mi := &file_carLogProto_car_log_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CarLogReq.ProtoReflect.Descriptor instead.
func (*CarLogReq) Descriptor() ([]byte, []int) {
	return file_carLogProto_car_log_proto_rawDescGZIP(), []int{0}
}

func (x *CarLogReq) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CarLogReq) GetPageNo() uint32 {
	if x != nil {
		return x.PageNo
	}
	return 0
}

func (x *CarLogReq) GetPageSize() uint32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

// 车辆日志结果
type CarLogResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 日志id
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	// 车辆id
	CarId uint64 `protobuf:"varint,2,opt,name=carId,proto3" json:"carId"`
	// 日志的uuid
	Uuid string `protobuf:"bytes,3,opt,name=uuid,proto3" json:"uuid"`
	// 日志remark
	Remark string `protobuf:"bytes,4,opt,name=remark,proto3" json:"remark"`
	// 当前操作者名称
	Operator string `protobuf:"bytes,5,opt,name=operator,proto3" json:"operator"`
	// 操作者id
	OperatorId uint64 `protobuf:"varint,6,opt,name=operatorId,proto3" json:"operatorId"`
	// 操作后的车辆状态
	CarState uint32 `protobuf:"varint,7,opt,name=carState,proto3" json:"carState"`
	// 操作后的车辆状态名称
	CarStateName string `protobuf:"bytes,8,opt,name=carStateName,proto3" json:"carStateName"`
	// 操作后的上架状态
	SaleStatus uint32 `protobuf:"varint,9,opt,name=saleStatus,proto3" json:"saleStatus"`
	// 操作后的上架状态名称
	SaleStatusName string `protobuf:"bytes,10,opt,name=saleStatusName,proto3" json:"saleStatusName"`
	// 操作后的库存状态
	InventoryStatus uint32 `protobuf:"varint,11,opt,name=inventoryStatus,proto3" json:"inventoryStatus"`
	// 操作后的库存状态名称
	InventoryStatusName string `protobuf:"bytes,12,opt,name=inventoryStatusName,proto3" json:"inventoryStatusName"`
	// 操作时间
	CreateTime uint64 `protobuf:"varint,14,opt,name=createTime,proto3" json:"createTime"`
	// 价格
	ExpSellingPrice string `protobuf:"bytes,15,opt,name=expSellingPrice,proto3" json:"expSellingPrice"`
}

func (x *CarLogResp) Reset() {
	*x = CarLogResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_carLogProto_car_log_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CarLogResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CarLogResp) ProtoMessage() {}

func (x *CarLogResp) ProtoReflect() protoreflect.Message {
	mi := &file_carLogProto_car_log_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CarLogResp.ProtoReflect.Descriptor instead.
func (*CarLogResp) Descriptor() ([]byte, []int) {
	return file_carLogProto_car_log_proto_rawDescGZIP(), []int{1}
}

func (x *CarLogResp) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CarLogResp) GetCarId() uint64 {
	if x != nil {
		return x.CarId
	}
	return 0
}

func (x *CarLogResp) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *CarLogResp) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

func (x *CarLogResp) GetOperator() string {
	if x != nil {
		return x.Operator
	}
	return ""
}

func (x *CarLogResp) GetOperatorId() uint64 {
	if x != nil {
		return x.OperatorId
	}
	return 0
}

func (x *CarLogResp) GetCarState() uint32 {
	if x != nil {
		return x.CarState
	}
	return 0
}

func (x *CarLogResp) GetCarStateName() string {
	if x != nil {
		return x.CarStateName
	}
	return ""
}

func (x *CarLogResp) GetSaleStatus() uint32 {
	if x != nil {
		return x.SaleStatus
	}
	return 0
}

func (x *CarLogResp) GetSaleStatusName() string {
	if x != nil {
		return x.SaleStatusName
	}
	return ""
}

func (x *CarLogResp) GetInventoryStatus() uint32 {
	if x != nil {
		return x.InventoryStatus
	}
	return 0
}

func (x *CarLogResp) GetInventoryStatusName() string {
	if x != nil {
		return x.InventoryStatusName
	}
	return ""
}

func (x *CarLogResp) GetCreateTime() uint64 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

func (x *CarLogResp) GetExpSellingPrice() string {
	if x != nil {
		return x.ExpSellingPrice
	}
	return ""
}

var File_carLogProto_car_log_proto protoreflect.FileDescriptor

var file_carLogProto_car_log_proto_rawDesc = []byte{
	0x0a, 0x19, 0x63, 0x61, 0x72, 0x4c, 0x6f, 0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x61,
	0x72, 0x5f, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x63, 0x61, 0x72,
	0x4c, 0x6f, 0x67, 0x1a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x63, 0x61, 0x72, 0x67, 0x6f, 0x64, 0x2d, 0x62, 0x6a, 0x2f, 0x62, 0x32, 0x63, 0x2d, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2d, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4f,
	0x0a, 0x09, 0x43, 0x61, 0x72, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x70,
	0x61, 0x67, 0x65, 0x4e, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x70, 0x61, 0x67,
	0x65, 0x4e, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x22,
	0xc8, 0x03, 0x0a, 0x0a, 0x43, 0x61, 0x72, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x63, 0x61, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x63,
	0x61, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x6d, 0x61,
	0x72, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b,
	0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x1e, 0x0a, 0x0a,
	0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x0a, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08,
	0x63, 0x61, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08,
	0x63, 0x61, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x61, 0x72, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x63, 0x61, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a,
	0x73, 0x61, 0x6c, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0a, 0x73, 0x61, 0x6c, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x26, 0x0a, 0x0e,
	0x73, 0x61, 0x6c, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x73, 0x61, 0x6c, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72,
	0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x69,
	0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x30,
	0x0a, 0x13, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x69, 0x6e, 0x76,
	0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x0e,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x28, 0x0a, 0x0f, 0x65, 0x78, 0x70, 0x53, 0x65, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x65, 0x78, 0x70, 0x53, 0x65,
	0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x69, 0x63, 0x65, 0x32, 0x37, 0x0a, 0x06, 0x43, 0x61,
	0x72, 0x4c, 0x6f, 0x67, 0x12, 0x2d, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x11, 0x2e, 0x63,
	0x61, 0x72, 0x4c, 0x6f, 0x67, 0x2e, 0x43, 0x61, 0x72, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x1a,
	0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x63, 0x61, 0x72, 0x67, 0x6f, 0x64, 0x2d, 0x62, 0x6a, 0x2f, 0x62, 0x32, 0x63, 0x2d,
	0x63, 0x61, 0x72, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x61, 0x72, 0x4c, 0x6f, 0x67, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_carLogProto_car_log_proto_rawDescOnce sync.Once
	file_carLogProto_car_log_proto_rawDescData = file_carLogProto_car_log_proto_rawDesc
)

func file_carLogProto_car_log_proto_rawDescGZIP() []byte {
	file_carLogProto_car_log_proto_rawDescOnce.Do(func() {
		file_carLogProto_car_log_proto_rawDescData = protoimpl.X.CompressGZIP(file_carLogProto_car_log_proto_rawDescData)
	})
	return file_carLogProto_car_log_proto_rawDescData
}

var file_carLogProto_car_log_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_carLogProto_car_log_proto_goTypes = []interface{}{
	(*CarLogReq)(nil),       // 0: carLog.CarLogReq
	(*CarLogResp)(nil),      // 1: carLog.CarLogResp
	(*common.Response)(nil), // 2: common.Response
}
var file_carLogProto_car_log_proto_depIdxs = []int32{
	0, // 0: carLog.CarLog.List:input_type -> carLog.CarLogReq
	2, // 1: carLog.CarLog.List:output_type -> common.Response
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_carLogProto_car_log_proto_init() }
func file_carLogProto_car_log_proto_init() {
	if File_carLogProto_car_log_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_carLogProto_car_log_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CarLogReq); i {
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
		file_carLogProto_car_log_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CarLogResp); i {
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
			RawDescriptor: file_carLogProto_car_log_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_carLogProto_car_log_proto_goTypes,
		DependencyIndexes: file_carLogProto_car_log_proto_depIdxs,
		MessageInfos:      file_carLogProto_car_log_proto_msgTypes,
	}.Build()
	File_carLogProto_car_log_proto = out.File
	file_carLogProto_car_log_proto_rawDesc = nil
	file_carLogProto_car_log_proto_goTypes = nil
	file_carLogProto_car_log_proto_depIdxs = nil
}
