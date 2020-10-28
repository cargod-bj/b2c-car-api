// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0-devel
// 	protoc        v3.12.1
// source: carCampaign/carCampaign.proto

package carCampaign

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

// 限时活动
type CarCampaignDto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 活动id
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// 活动名称
	Name                      uint64 `protobuf:"varint,2,opt,name=name,proto3" json:"name,omitempty"`
	ResWebFilterTagSelected   string `protobuf:"bytes,4,opt,name=ResWebFilterTagSelected,proto3" json:"ResWebFilterTagSelected,omitempty"`
	ResWebFilterTagUnselected string `protobuf:"bytes,5,opt,name=ResWebFilterTagUnselected,proto3" json:"ResWebFilterTagUnselected,omitempty"`
	ResH5FilterTagSelected    string `protobuf:"bytes,6,opt,name=ResH5FilterTagSelected,proto3" json:"ResH5FilterTagSelected,omitempty"`
	ResH5FilterTagUnselected  string `protobuf:"bytes,7,opt,name=ResH5FilterTagUnselected,proto3" json:"ResH5FilterTagUnselected,omitempty"`
}

func (x *CarCampaignDto) Reset() {
	*x = CarCampaignDto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_carCampaign_carCampaign_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CarCampaignDto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CarCampaignDto) ProtoMessage() {}

func (x *CarCampaignDto) ProtoReflect() protoreflect.Message {
	mi := &file_carCampaign_carCampaign_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CarCampaignDto.ProtoReflect.Descriptor instead.
func (*CarCampaignDto) Descriptor() ([]byte, []int) {
	return file_carCampaign_carCampaign_proto_rawDescGZIP(), []int{0}
}

func (x *CarCampaignDto) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CarCampaignDto) GetName() uint64 {
	if x != nil {
		return x.Name
	}
	return 0
}

func (x *CarCampaignDto) GetResWebFilterTagSelected() string {
	if x != nil {
		return x.ResWebFilterTagSelected
	}
	return ""
}

func (x *CarCampaignDto) GetResWebFilterTagUnselected() string {
	if x != nil {
		return x.ResWebFilterTagUnselected
	}
	return ""
}

func (x *CarCampaignDto) GetResH5FilterTagSelected() string {
	if x != nil {
		return x.ResH5FilterTagSelected
	}
	return ""
}

func (x *CarCampaignDto) GetResH5FilterTagUnselected() string {
	if x != nil {
		return x.ResH5FilterTagUnselected
	}
	return ""
}

var File_carCampaign_carCampaign_proto protoreflect.FileDescriptor

var file_carCampaign_carCampaign_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x63, 0x61, 0x72, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x2f, 0x63, 0x61,
	0x72, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0b, 0x63, 0x61, 0x72, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x1a, 0x39, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x61, 0x72, 0x67, 0x6f, 0x64, 0x2d,
	0x62, 0x6a, 0x2f, 0x62, 0x32, 0x63, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa0, 0x02, 0x0a, 0x0e, 0x43, 0x61, 0x72, 0x43,
	0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x44, 0x74, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x38,
	0x0a, 0x17, 0x52, 0x65, 0x73, 0x57, 0x65, 0x62, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x54, 0x61,
	0x67, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x17, 0x52, 0x65, 0x73, 0x57, 0x65, 0x62, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x54, 0x61, 0x67,
	0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x12, 0x3c, 0x0a, 0x19, 0x52, 0x65, 0x73, 0x57,
	0x65, 0x62, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x54, 0x61, 0x67, 0x55, 0x6e, 0x73, 0x65, 0x6c,
	0x65, 0x63, 0x74, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x19, 0x52, 0x65, 0x73,
	0x57, 0x65, 0x62, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x54, 0x61, 0x67, 0x55, 0x6e, 0x73, 0x65,
	0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x12, 0x36, 0x0a, 0x16, 0x52, 0x65, 0x73, 0x48, 0x35, 0x46,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x54, 0x61, 0x67, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x16, 0x52, 0x65, 0x73, 0x48, 0x35, 0x46, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x54, 0x61, 0x67, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x12, 0x3a,
	0x0a, 0x18, 0x52, 0x65, 0x73, 0x48, 0x35, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x54, 0x61, 0x67,
	0x55, 0x6e, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x18, 0x52, 0x65, 0x73, 0x48, 0x35, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x54, 0x61, 0x67,
	0x55, 0x6e, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x32, 0x4c, 0x0a, 0x0b, 0x43, 0x61,
	0x72, 0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x12, 0x3d, 0x0a, 0x18, 0x47, 0x65, 0x74,
	0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x43, 0x61, 0x6d,
	0x70, 0x61, 0x69, 0x67, 0x6e, 0x12, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x49,
	0x64, 0x44, 0x74, 0x6f, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x61, 0x72, 0x67, 0x6f, 0x64, 0x2d, 0x62, 0x6a,
	0x2f, 0x62, 0x32, 0x63, 0x2d, 0x63, 0x61, 0x72, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x61, 0x72,
	0x43, 0x61, 0x6d, 0x70, 0x61, 0x69, 0x67, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_carCampaign_carCampaign_proto_rawDescOnce sync.Once
	file_carCampaign_carCampaign_proto_rawDescData = file_carCampaign_carCampaign_proto_rawDesc
)

func file_carCampaign_carCampaign_proto_rawDescGZIP() []byte {
	file_carCampaign_carCampaign_proto_rawDescOnce.Do(func() {
		file_carCampaign_carCampaign_proto_rawDescData = protoimpl.X.CompressGZIP(file_carCampaign_carCampaign_proto_rawDescData)
	})
	return file_carCampaign_carCampaign_proto_rawDescData
}

var file_carCampaign_carCampaign_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_carCampaign_carCampaign_proto_goTypes = []interface{}{
	(*CarCampaignDto)(nil),  // 0: carCampaign.CarCampaignDto
	(*common.IdDto)(nil),    // 1: common.IdDto
	(*common.Response)(nil), // 2: common.Response
}
var file_carCampaign_carCampaign_proto_depIdxs = []int32{
	1, // 0: carCampaign.CarCampaign.GetCurrentActiveCampaign:input_type -> common.IdDto
	2, // 1: carCampaign.CarCampaign.GetCurrentActiveCampaign:output_type -> common.Response
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_carCampaign_carCampaign_proto_init() }
func file_carCampaign_carCampaign_proto_init() {
	if File_carCampaign_carCampaign_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_carCampaign_carCampaign_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CarCampaignDto); i {
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
			RawDescriptor: file_carCampaign_carCampaign_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_carCampaign_carCampaign_proto_goTypes,
		DependencyIndexes: file_carCampaign_carCampaign_proto_depIdxs,
		MessageInfos:      file_carCampaign_carCampaign_proto_msgTypes,
	}.Build()
	File_carCampaign_carCampaign_proto = out.File
	file_carCampaign_carCampaign_proto_rawDesc = nil
	file_carCampaign_carCampaign_proto_goTypes = nil
	file_carCampaign_carCampaign_proto_depIdxs = nil
}
