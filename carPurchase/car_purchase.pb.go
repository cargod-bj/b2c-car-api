// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: carPurchase/car_purchase.proto

package carPurchaseProto

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

type InspectionReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InspectionId uint64 `protobuf:"varint,1,opt,name=inspectionId,proto3" json:"inspectionId,omitempty"`
}

func (x *InspectionReq) Reset() {
	*x = InspectionReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_carPurchase_car_purchase_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InspectionReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InspectionReq) ProtoMessage() {}

func (x *InspectionReq) ProtoReflect() protoreflect.Message {
	mi := &file_carPurchase_car_purchase_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InspectionReq.ProtoReflect.Descriptor instead.
func (*InspectionReq) Descriptor() ([]byte, []int) {
	return file_carPurchase_car_purchase_proto_rawDescGZIP(), []int{0}
}

func (x *InspectionReq) GetInspectionId() uint64 {
	if x != nil {
		return x.InspectionId
	}
	return 0
}

var File_carPurchase_car_purchase_proto protoreflect.FileDescriptor

var file_carPurchase_car_purchase_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x63, 0x61, 0x72, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x2f, 0x63, 0x61,
	0x72, 0x5f, 0x70, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x10, 0x63, 0x61, 0x72, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63,
	0x61, 0x72, 0x67, 0x6f, 0x64, 0x2d, 0x62, 0x6a, 0x2f, 0x62, 0x32, 0x63, 0x2d, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2d, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x33, 0x0a,
	0x0d, 0x49, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x22,
	0x0a, 0x0c, 0x69, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x69, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x64, 0x32, 0x58, 0x0a, 0x0b, 0x43, 0x61, 0x72, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73,
	0x65, 0x12, 0x49, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x42, 0x79, 0x49, 0x6e, 0x73,
	0x70, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x2e, 0x63, 0x61, 0x72, 0x50, 0x75, 0x72,
	0x63, 0x68, 0x61, 0x73, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x6e, 0x73, 0x70, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x33, 0x5a, 0x31,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x61, 0x72, 0x67, 0x6f,
	0x64, 0x2d, 0x62, 0x6a, 0x2f, 0x62, 0x32, 0x63, 0x2d, 0x63, 0x61, 0x72, 0x2d, 0x61, 0x70, 0x69,
	0x2f, 0x63, 0x61, 0x72, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_carPurchase_car_purchase_proto_rawDescOnce sync.Once
	file_carPurchase_car_purchase_proto_rawDescData = file_carPurchase_car_purchase_proto_rawDesc
)

func file_carPurchase_car_purchase_proto_rawDescGZIP() []byte {
	file_carPurchase_car_purchase_proto_rawDescOnce.Do(func() {
		file_carPurchase_car_purchase_proto_rawDescData = protoimpl.X.CompressGZIP(file_carPurchase_car_purchase_proto_rawDescData)
	})
	return file_carPurchase_car_purchase_proto_rawDescData
}

var file_carPurchase_car_purchase_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_carPurchase_car_purchase_proto_goTypes = []interface{}{
	(*InspectionReq)(nil),   // 0: carPurchaseProto.InspectionReq
	(*common.Response)(nil), // 1: common.Response
}
var file_carPurchase_car_purchase_proto_depIdxs = []int32{
	0, // 0: carPurchaseProto.CarPurchase.GetCarByInspection:input_type -> carPurchaseProto.InspectionReq
	1, // 1: carPurchaseProto.CarPurchase.GetCarByInspection:output_type -> common.Response
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_carPurchase_car_purchase_proto_init() }
func file_carPurchase_car_purchase_proto_init() {
	if File_carPurchase_car_purchase_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_carPurchase_car_purchase_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InspectionReq); i {
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
			RawDescriptor: file_carPurchase_car_purchase_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_carPurchase_car_purchase_proto_goTypes,
		DependencyIndexes: file_carPurchase_car_purchase_proto_depIdxs,
		MessageInfos:      file_carPurchase_car_purchase_proto_msgTypes,
	}.Build()
	File_carPurchase_car_purchase_proto = out.File
	file_carPurchase_car_purchase_proto_rawDesc = nil
	file_carPurchase_car_purchase_proto_goTypes = nil
	file_carPurchase_car_purchase_proto_depIdxs = nil
}