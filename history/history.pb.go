// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0-devel
// 	protoc        v3.12.1
// source: history/history.proto

package historyProto

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

type History struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CarId  uint64 `protobuf:"varint,1,opt,name=CarId,proto3" json:"CarId,omitempty"`
	UserId uint64 `protobuf:"varint,2,opt,name=UserId,proto3" json:"UserId,omitempty"`
	Local  string `protobuf:"bytes,3,opt,name=Local,proto3" json:"Local,omitempty"`
}

func (x *History) Reset() {
	*x = History{}
	if protoimpl.UnsafeEnabled {
		mi := &file_history_history_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *History) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*History) ProtoMessage() {}

func (x *History) ProtoReflect() protoreflect.Message {
	mi := &file_history_history_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use History.ProtoReflect.Descriptor instead.
func (*History) Descriptor() ([]byte, []int) {
	return file_history_history_proto_rawDescGZIP(), []int{0}
}

func (x *History) GetCarId() uint64 {
	if x != nil {
		return x.CarId
	}
	return 0
}

func (x *History) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *History) GetLocal() string {
	if x != nil {
		return x.Local
	}
	return ""
}

var File_history_history_proto protoreflect.FileDescriptor

var file_history_history_proto_rawDesc = []byte{
	0x0a, 0x15, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79,
	0x1a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x61, 0x72,
	0x67, 0x6f, 0x64, 0x2d, 0x62, 0x6a, 0x2f, 0x62, 0x32, 0x63, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2d, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4d, 0x0a, 0x07, 0x48,
	0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x43, 0x61, 0x72, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x43, 0x61, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x32, 0x72, 0x0a, 0x07, 0x68, 0x69,
	0x73, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x33, 0x0a, 0x0b, 0x53, 0x61, 0x76, 0x65, 0x48, 0x69, 0x73,
	0x74, 0x6f, 0x72, 0x79, 0x12, 0x10, 0x2e, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x48,
	0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x32, 0x0a, 0x0a, 0x47, 0x65,
	0x74, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x10, 0x2e, 0x68, 0x69, 0x73, 0x74, 0x6f,
	0x72, 0x79, 0x2e, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x2f,
	0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x61, 0x72,
	0x67, 0x6f, 0x64, 0x2d, 0x62, 0x6a, 0x2f, 0x62, 0x32, 0x63, 0x2d, 0x63, 0x61, 0x72, 0x2d, 0x61,
	0x70, 0x69, 0x2f, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_history_history_proto_rawDescOnce sync.Once
	file_history_history_proto_rawDescData = file_history_history_proto_rawDesc
)

func file_history_history_proto_rawDescGZIP() []byte {
	file_history_history_proto_rawDescOnce.Do(func() {
		file_history_history_proto_rawDescData = protoimpl.X.CompressGZIP(file_history_history_proto_rawDescData)
	})
	return file_history_history_proto_rawDescData
}

var file_history_history_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_history_history_proto_goTypes = []interface{}{
	(*History)(nil),         // 0: history.History
	(*common.Response)(nil), // 1: common.Response
}
var file_history_history_proto_depIdxs = []int32{
	0, // 0: history.history.SaveHistory:input_type -> history.History
	0, // 1: history.history.GetHistory:input_type -> history.History
	1, // 2: history.history.SaveHistory:output_type -> common.Response
	1, // 3: history.history.GetHistory:output_type -> common.Response
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_history_history_proto_init() }
func file_history_history_proto_init() {
	if File_history_history_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_history_history_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*History); i {
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
			RawDescriptor: file_history_history_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_history_history_proto_goTypes,
		DependencyIndexes: file_history_history_proto_depIdxs,
		MessageInfos:      file_history_history_proto_msgTypes,
	}.Build()
	File_history_history_proto = out.File
	file_history_history_proto_rawDesc = nil
	file_history_history_proto_goTypes = nil
	file_history_history_proto_depIdxs = nil
}