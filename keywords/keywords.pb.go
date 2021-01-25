// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.1
// source: keywords/keywords.proto

package keywords

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

// 新增Tagdto
type KeywordsDto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// keywords id，如果为新增，则此字段为0
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	//keywords名称
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// keywords类型
	Type string `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *KeywordsDto) Reset() {
	*x = KeywordsDto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_keywords_keywords_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeywordsDto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeywordsDto) ProtoMessage() {}

func (x *KeywordsDto) ProtoReflect() protoreflect.Message {
	mi := &file_keywords_keywords_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeywordsDto.ProtoReflect.Descriptor instead.
func (*KeywordsDto) Descriptor() ([]byte, []int) {
	return file_keywords_keywords_proto_rawDescGZIP(), []int{0}
}

func (x *KeywordsDto) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *KeywordsDto) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *KeywordsDto) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

var File_keywords_keywords_proto protoreflect.FileDescriptor

var file_keywords_keywords_proto_rawDesc = []byte{
	0x0a, 0x17, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x2f, 0x6b, 0x65, 0x79, 0x77, 0x6f,
	0x72, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x63, 0x61, 0x72, 0x1a, 0x39,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x61, 0x72, 0x67, 0x6f,
	0x64, 0x2d, 0x62, 0x6a, 0x2f, 0x62, 0x32, 0x63, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x45, 0x0a, 0x0b, 0x4b, 0x65, 0x79,
	0x77, 0x6f, 0x72, 0x64, 0x73, 0x44, 0x74, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x32, 0xe2, 0x01, 0x0a, 0x08, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x28, 0x0a,
	0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50,
	0x61, 0x67, 0x65, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x33, 0x0a, 0x0f, 0x50, 0x75, 0x62, 0x6c, 0x69,
	0x73, 0x74, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x0c, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x14,
	0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x79, 0x4b, 0x65, 0x79, 0x77,
	0x6f, 0x72, 0x64, 0x73, 0x12, 0x10, 0x2e, 0x63, 0x61, 0x72, 0x2e, 0x4b, 0x65, 0x79, 0x77, 0x6f,
	0x72, 0x64, 0x73, 0x44, 0x74, 0x6f, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x11, 0x47, 0x65,
	0x74, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x10, 0x2e, 0x63, 0x61, 0x72, 0x2e, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x44, 0x74,
	0x6f, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x61, 0x72, 0x67, 0x6f, 0x64, 0x2d, 0x62, 0x6a, 0x2f, 0x62, 0x32,
	0x63, 0x2d, 0x63, 0x61, 0x72, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72,
	0x64, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_keywords_keywords_proto_rawDescOnce sync.Once
	file_keywords_keywords_proto_rawDescData = file_keywords_keywords_proto_rawDesc
)

func file_keywords_keywords_proto_rawDescGZIP() []byte {
	file_keywords_keywords_proto_rawDescOnce.Do(func() {
		file_keywords_keywords_proto_rawDescData = protoimpl.X.CompressGZIP(file_keywords_keywords_proto_rawDescData)
	})
	return file_keywords_keywords_proto_rawDescData
}

var file_keywords_keywords_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_keywords_keywords_proto_goTypes = []interface{}{
	(*KeywordsDto)(nil),     // 0: car.KeywordsDto
	(*common.Page)(nil),     // 1: common.Page
	(*common.Response)(nil), // 2: common.Response
}
var file_keywords_keywords_proto_depIdxs = []int32{
	1, // 0: car.Keywords.List:input_type -> common.Page
	1, // 1: car.Keywords.PublistKeywords:input_type -> common.Page
	0, // 2: car.Keywords.GetCarListByKeywords:input_type -> car.KeywordsDto
	0, // 3: car.Keywords.GetKeywordsByName:input_type -> car.KeywordsDto
	2, // 4: car.Keywords.List:output_type -> common.Response
	2, // 5: car.Keywords.PublistKeywords:output_type -> common.Response
	2, // 6: car.Keywords.GetCarListByKeywords:output_type -> common.Response
	2, // 7: car.Keywords.GetKeywordsByName:output_type -> common.Response
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_keywords_keywords_proto_init() }
func file_keywords_keywords_proto_init() {
	if File_keywords_keywords_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_keywords_keywords_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeywordsDto); i {
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
			RawDescriptor: file_keywords_keywords_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_keywords_keywords_proto_goTypes,
		DependencyIndexes: file_keywords_keywords_proto_depIdxs,
		MessageInfos:      file_keywords_keywords_proto_msgTypes,
	}.Build()
	File_keywords_keywords_proto = out.File
	file_keywords_keywords_proto_rawDesc = nil
	file_keywords_keywords_proto_goTypes = nil
	file_keywords_keywords_proto_depIdxs = nil
}
