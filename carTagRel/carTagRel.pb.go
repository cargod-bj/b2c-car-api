// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0
// 	protoc        v3.9.0
// source: carTagRel/carTagRel.proto

package carTagRel

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

// 新增Tagreldto
type CarTagRelDto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 标签id，如果为新增，则此字段为0
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	//tag的id
	TagId uint64 `protobuf:"varint,2,opt,name=tagId,proto3" json:"tagId,omitempty"`
	// taglist类型id
	CarListId uint64 `protobuf:"varint,3,opt,name=carListId,proto3" json:"carListId,omitempty"`
	//tagrel新增时间
	CreateTime uint64 `protobuf:"varint,5,opt,name=createTime,proto3" json:"createTime,omitempty"`
	// tagrel的更新时间，如果为新增，或更新，则此字段无效
	UpdateTime uint64 `protobuf:"varint,6,opt,name=updateTime,proto3" json:"updateTime,omitempty"`
}

func (x *CarTagRelDto) Reset() {
	*x = CarTagRelDto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_carTagRel_carTagRel_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CarTagRelDto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CarTagRelDto) ProtoMessage() {}

func (x *CarTagRelDto) ProtoReflect() protoreflect.Message {
	mi := &file_carTagRel_carTagRel_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CarTagRelDto.ProtoReflect.Descriptor instead.
func (*CarTagRelDto) Descriptor() ([]byte, []int) {
	return file_carTagRel_carTagRel_proto_rawDescGZIP(), []int{0}
}

func (x *CarTagRelDto) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CarTagRelDto) GetTagId() uint64 {
	if x != nil {
		return x.TagId
	}
	return 0
}

func (x *CarTagRelDto) GetCarListId() uint64 {
	if x != nil {
		return x.CarListId
	}
	return 0
}

func (x *CarTagRelDto) GetCreateTime() uint64 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

func (x *CarTagRelDto) GetUpdateTime() uint64 {
	if x != nil {
		return x.UpdateTime
	}
	return 0
}

var File_carTagRel_carTagRel_proto protoreflect.FileDescriptor

var file_carTagRel_carTagRel_proto_rawDesc = []byte{
	0x0a, 0x19, 0x63, 0x61, 0x72, 0x54, 0x61, 0x67, 0x52, 0x65, 0x6c, 0x2f, 0x63, 0x61, 0x72, 0x54,
	0x61, 0x67, 0x52, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x63, 0x61, 0x72,
	0x1a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x61, 0x72,
	0x67, 0x6f, 0x64, 0x2d, 0x62, 0x6a, 0x2f, 0x62, 0x32, 0x63, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2d, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x92, 0x01, 0x0a, 0x0c,
	0x43, 0x61, 0x72, 0x54, 0x61, 0x67, 0x52, 0x65, 0x6c, 0x44, 0x74, 0x6f, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x61, 0x67, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x74, 0x61, 0x67,
	0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x61, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x63, 0x61, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x64,
	0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x1e, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x32, 0xc5, 0x01, 0x0a, 0x09, 0x43, 0x61, 0x72, 0x54, 0x61, 0x67, 0x52, 0x65, 0x6c, 0x12, 0x2c,
	0x0a, 0x03, 0x41, 0x64, 0x64, 0x12, 0x11, 0x2e, 0x63, 0x61, 0x72, 0x2e, 0x43, 0x61, 0x72, 0x54,
	0x61, 0x67, 0x52, 0x65, 0x6c, 0x44, 0x74, 0x6f, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2f, 0x0a, 0x06,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x11, 0x2e, 0x63, 0x61, 0x72, 0x2e, 0x43, 0x61, 0x72,
	0x54, 0x61, 0x67, 0x52, 0x65, 0x6c, 0x44, 0x74, 0x6f, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2f, 0x0a,
	0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x11, 0x2e, 0x63, 0x61, 0x72, 0x2e, 0x43, 0x61,
	0x72, 0x54, 0x61, 0x67, 0x52, 0x65, 0x6c, 0x44, 0x74, 0x6f, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x28,
	0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x50, 0x61, 0x67, 0x65, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x2c, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x61, 0x72, 0x67, 0x6f, 0x64, 0x2d, 0x62, 0x6a,
	0x2f, 0x62, 0x32, 0x63, 0x2d, 0x63, 0x61, 0x72, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x61, 0x72,
	0x54, 0x61, 0x67, 0x52, 0x65, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_carTagRel_carTagRel_proto_rawDescOnce sync.Once
	file_carTagRel_carTagRel_proto_rawDescData = file_carTagRel_carTagRel_proto_rawDesc
)

func file_carTagRel_carTagRel_proto_rawDescGZIP() []byte {
	file_carTagRel_carTagRel_proto_rawDescOnce.Do(func() {
		file_carTagRel_carTagRel_proto_rawDescData = protoimpl.X.CompressGZIP(file_carTagRel_carTagRel_proto_rawDescData)
	})
	return file_carTagRel_carTagRel_proto_rawDescData
}

var file_carTagRel_carTagRel_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_carTagRel_carTagRel_proto_goTypes = []interface{}{
	(*CarTagRelDto)(nil),    // 0: car.CarTagRelDto
	(*common.Page)(nil),     // 1: common.Page
	(*common.Response)(nil), // 2: common.Response
}
var file_carTagRel_carTagRel_proto_depIdxs = []int32{
	0, // 0: car.CarTagRel.Add:input_type -> car.CarTagRelDto
	0, // 1: car.CarTagRel.Delete:input_type -> car.CarTagRelDto
	0, // 2: car.CarTagRel.Update:input_type -> car.CarTagRelDto
	1, // 3: car.CarTagRel.List:input_type -> common.Page
	2, // 4: car.CarTagRel.Add:output_type -> common.Response
	2, // 5: car.CarTagRel.Delete:output_type -> common.Response
	2, // 6: car.CarTagRel.Update:output_type -> common.Response
	2, // 7: car.CarTagRel.List:output_type -> common.Response
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_carTagRel_carTagRel_proto_init() }
func file_carTagRel_carTagRel_proto_init() {
	if File_carTagRel_carTagRel_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_carTagRel_carTagRel_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CarTagRelDto); i {
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
			RawDescriptor: file_carTagRel_carTagRel_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_carTagRel_carTagRel_proto_goTypes,
		DependencyIndexes: file_carTagRel_carTagRel_proto_depIdxs,
		MessageInfos:      file_carTagRel_carTagRel_proto_msgTypes,
	}.Build()
	File_carTagRel_carTagRel_proto = out.File
	file_carTagRel_carTagRel_proto_rawDesc = nil
	file_carTagRel_carTagRel_proto_goTypes = nil
	file_carTagRel_carTagRel_proto_depIdxs = nil
}
