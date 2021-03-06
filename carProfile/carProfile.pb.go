// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.14.0
// source: carProfile/carProfile.proto

package carProfile

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

type CarProfileCond struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CarModel string `protobuf:"bytes,1,opt,name=carModel,proto3" json:"carModel"`
}

func (x *CarProfileCond) Reset() {
	*x = CarProfileCond{}
	if protoimpl.UnsafeEnabled {
		mi := &file_carProfile_carProfile_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CarProfileCond) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CarProfileCond) ProtoMessage() {}

func (x *CarProfileCond) ProtoReflect() protoreflect.Message {
	mi := &file_carProfile_carProfile_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CarProfileCond.ProtoReflect.Descriptor instead.
func (*CarProfileCond) Descriptor() ([]byte, []int) {
	return file_carProfile_carProfile_proto_rawDescGZIP(), []int{0}
}

func (x *CarProfileCond) GetCarModel() string {
	if x != nil {
		return x.CarModel
	}
	return ""
}

type CarProfileResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	Brand        string `protobuf:"bytes,2,opt,name=brand,proto3" json:"brand"`
	Model        string `protobuf:"bytes,3,opt,name=model,proto3" json:"model"`
	Variant      string `protobuf:"bytes,4,opt,name=variant,proto3" json:"variant"`
	Year         string `protobuf:"bytes,5,opt,name=year,proto3" json:"year"`
	Engine       string `protobuf:"bytes,6,opt,name=engine,proto3" json:"engine"`
	Transmission string `protobuf:"bytes,7,opt,name=transmission,proto3" json:"transmission"`
	BodyType     string `protobuf:"bytes,8,opt,name=bodyType,proto3" json:"bodyType"`
	UniqueCode   string `protobuf:"bytes,9,opt,name=uniqueCode,proto3" json:"uniqueCode"`
	Profile      string `protobuf:"bytes,10,opt,name=profile,proto3" json:"profile"`
	TypicalId    int64  `protobuf:"varint,11,opt,name=typicalId,proto3" json:"typicalId"`
	CreateTime   uint64 `protobuf:"varint,12,opt,name=createTime,proto3" json:"createTime"`
	UpdateTime   uint64 `protobuf:"varint,13,opt,name=updateTime,proto3" json:"updateTime"`
}

func (x *CarProfileResp) Reset() {
	*x = CarProfileResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_carProfile_carProfile_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CarProfileResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CarProfileResp) ProtoMessage() {}

func (x *CarProfileResp) ProtoReflect() protoreflect.Message {
	mi := &file_carProfile_carProfile_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CarProfileResp.ProtoReflect.Descriptor instead.
func (*CarProfileResp) Descriptor() ([]byte, []int) {
	return file_carProfile_carProfile_proto_rawDescGZIP(), []int{1}
}

func (x *CarProfileResp) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CarProfileResp) GetBrand() string {
	if x != nil {
		return x.Brand
	}
	return ""
}

func (x *CarProfileResp) GetModel() string {
	if x != nil {
		return x.Model
	}
	return ""
}

func (x *CarProfileResp) GetVariant() string {
	if x != nil {
		return x.Variant
	}
	return ""
}

func (x *CarProfileResp) GetYear() string {
	if x != nil {
		return x.Year
	}
	return ""
}

func (x *CarProfileResp) GetEngine() string {
	if x != nil {
		return x.Engine
	}
	return ""
}

func (x *CarProfileResp) GetTransmission() string {
	if x != nil {
		return x.Transmission
	}
	return ""
}

func (x *CarProfileResp) GetBodyType() string {
	if x != nil {
		return x.BodyType
	}
	return ""
}

func (x *CarProfileResp) GetUniqueCode() string {
	if x != nil {
		return x.UniqueCode
	}
	return ""
}

func (x *CarProfileResp) GetProfile() string {
	if x != nil {
		return x.Profile
	}
	return ""
}

func (x *CarProfileResp) GetTypicalId() int64 {
	if x != nil {
		return x.TypicalId
	}
	return 0
}

func (x *CarProfileResp) GetCreateTime() uint64 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

func (x *CarProfileResp) GetUpdateTime() uint64 {
	if x != nil {
		return x.UpdateTime
	}
	return 0
}

var File_carProfile_carProfile_proto protoreflect.FileDescriptor

var file_carProfile_carProfile_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x63, 0x61, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2f, 0x63, 0x61, 0x72,
	0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x63,
	0x61, 0x72, 0x1a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63,
	0x61, 0x72, 0x67, 0x6f, 0x64, 0x2d, 0x62, 0x6a, 0x2f, 0x62, 0x32, 0x63, 0x2d, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2d, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2c, 0x0a,
	0x0e, 0x43, 0x61, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x64, 0x12,
	0x1a, 0x0a, 0x08, 0x63, 0x61, 0x72, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x63, 0x61, 0x72, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x22, 0xea, 0x02, 0x0a, 0x0e,
	0x43, 0x61, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x62,
	0x72, 0x61, 0x6e, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x61,
	0x72, 0x69, 0x61, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x61, 0x72,
	0x69, 0x61, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x79, 0x65, 0x61, 0x72, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x79, 0x65, 0x61, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x6e, 0x67, 0x69,
	0x6e, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65,
	0x12, 0x22, 0x0a, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x62, 0x6f, 0x64, 0x79, 0x54, 0x79, 0x70, 0x65,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x6f, 0x64, 0x79, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x1e, 0x0a, 0x0a, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x79,
	0x70, 0x69, 0x63, 0x61, 0x6c, 0x49, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74,
	0x79, 0x70, 0x69, 0x63, 0x61, 0x6c, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x32, 0x71, 0x0a, 0x0a, 0x43, 0x61, 0x72, 0x50,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x30, 0x0a, 0x08, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x73,
	0x69, 0x73, 0x12, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x44, 0x74, 0x6f, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x06, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x12, 0x13, 0x2e, 0x63, 0x61, 0x72, 0x2e, 0x43, 0x61, 0x72, 0x50, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x64, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x2d, 0x5a, 0x2b, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x61, 0x72, 0x67, 0x6f, 0x64,
	0x2d, 0x62, 0x6a, 0x2f, 0x62, 0x32, 0x63, 0x2d, 0x63, 0x61, 0x72, 0x2d, 0x61, 0x70, 0x69, 0x2f,
	0x63, 0x61, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_carProfile_carProfile_proto_rawDescOnce sync.Once
	file_carProfile_carProfile_proto_rawDescData = file_carProfile_carProfile_proto_rawDesc
)

func file_carProfile_carProfile_proto_rawDescGZIP() []byte {
	file_carProfile_carProfile_proto_rawDescOnce.Do(func() {
		file_carProfile_carProfile_proto_rawDescData = protoimpl.X.CompressGZIP(file_carProfile_carProfile_proto_rawDescData)
	})
	return file_carProfile_carProfile_proto_rawDescData
}

var file_carProfile_carProfile_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_carProfile_carProfile_proto_goTypes = []interface{}{
	(*CarProfileCond)(nil),  // 0: car.CarProfileCond
	(*CarProfileResp)(nil),  // 1: car.CarProfileResp
	(*common.EmptyDto)(nil), // 2: common.EmptyDto
	(*common.Response)(nil), // 3: common.Response
}
var file_carProfile_carProfile_proto_depIdxs = []int32{
	2, // 0: car.CarProfile.Analysis:input_type -> common.EmptyDto
	0, // 1: car.CarProfile.Detail:input_type -> car.CarProfileCond
	3, // 2: car.CarProfile.Analysis:output_type -> common.Response
	3, // 3: car.CarProfile.Detail:output_type -> common.Response
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_carProfile_carProfile_proto_init() }
func file_carProfile_carProfile_proto_init() {
	if File_carProfile_carProfile_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_carProfile_carProfile_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CarProfileCond); i {
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
		file_carProfile_carProfile_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CarProfileResp); i {
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
			RawDescriptor: file_carProfile_carProfile_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_carProfile_carProfile_proto_goTypes,
		DependencyIndexes: file_carProfile_carProfile_proto_depIdxs,
		MessageInfos:      file_carProfile_carProfile_proto_msgTypes,
	}.Build()
	File_carProfile_carProfile_proto = out.File
	file_carProfile_carProfile_proto_rawDesc = nil
	file_carProfile_carProfile_proto_goTypes = nil
	file_carProfile_carProfile_proto_depIdxs = nil
}
