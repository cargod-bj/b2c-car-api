// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0-devel
// 	protoc        v3.15.8
// source: brandModel/model.proto

package brandModelProto

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

// 品牌dto
type ModelDto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id，如果为新增，则此字段为0，否则为id
	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// 创建时间，如果为新增，或更新，则此字段无效
	CreateTime uint64 `protobuf:"varint,2,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// 更新时间，如果为新增，或更新，则此字段无效
	UpdateTime uint64 `protobuf:"varint,3,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	// 当前key的描述信息, 如果为更新，则此字段无效
	Name    string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	BrandId uint32 `protobuf:"varint,5,opt,name=brand_id,json=brandId,proto3" json:"brand_id,omitempty"`
	Status  uint32 `protobuf:"varint,6,opt,name=status,proto3" json:"status,omitempty"`
	Order   uint32 `protobuf:"varint,7,opt,name=order,proto3" json:"order,omitempty"`
}

func (x *ModelDto) Reset() {
	*x = ModelDto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_brandModel_model_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModelDto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModelDto) ProtoMessage() {}

func (x *ModelDto) ProtoReflect() protoreflect.Message {
	mi := &file_brandModel_model_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModelDto.ProtoReflect.Descriptor instead.
func (*ModelDto) Descriptor() ([]byte, []int) {
	return file_brandModel_model_proto_rawDescGZIP(), []int{0}
}

func (x *ModelDto) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ModelDto) GetCreateTime() uint64 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

func (x *ModelDto) GetUpdateTime() uint64 {
	if x != nil {
		return x.UpdateTime
	}
	return 0
}

func (x *ModelDto) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ModelDto) GetBrandId() uint32 {
	if x != nil {
		return x.BrandId
	}
	return 0
}

func (x *ModelDto) GetStatus() uint32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *ModelDto) GetOrder() uint32 {
	if x != nil {
		return x.Order
	}
	return 0
}

var File_brandModel_model_proto protoreflect.FileDescriptor

var file_brandModel_model_proto_rawDesc = []byte{
	0x0a, 0x16, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x4d,
	0x6f, 0x64, 0x65, 0x6c, 0x1a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x63, 0x61, 0x72, 0x67, 0x6f, 0x64, 0x2d, 0x62, 0x6a, 0x2f, 0x62, 0x32, 0x63, 0x2d, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xb9, 0x01, 0x0a, 0x08, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x44, 0x74, 0x6f, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1f, 0x0a,
	0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x49, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x32, 0xc1, 0x02, 0x0a, 0x05,
	0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x2f, 0x0a, 0x03, 0x41, 0x64, 0x64, 0x12, 0x14, 0x2e, 0x62,
	0x72, 0x61, 0x6e, 0x64, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x44,
	0x74, 0x6f, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x30, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x12, 0x12, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x49, 0x64, 0x4c, 0x6f, 0x63, 0x61,
	0x6c, 0x44, 0x54, 0x4f, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x32, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x12, 0x14, 0x2e, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x2e,
	0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x44, 0x74, 0x6f, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2d, 0x0a, 0x03,
	0x47, 0x65, 0x74, 0x12, 0x12, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x49, 0x64, 0x4c,
	0x6f, 0x63, 0x61, 0x6c, 0x44, 0x54, 0x4f, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x0d, 0x4c,
	0x69, 0x73, 0x74, 0x42, 0x79, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x49, 0x64, 0x12, 0x12, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x49, 0x64, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x44, 0x54, 0x4f,
	0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x6f, 0x64, 0x65,
	0x6c, 0x48, 0x61, 0x73, 0x43, 0x61, 0x72, 0x12, 0x12, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x49, 0x64, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x44, 0x54, 0x4f, 0x1a, 0x10, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x32, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x61,
	0x72, 0x67, 0x6f, 0x64, 0x2d, 0x62, 0x6a, 0x2f, 0x62, 0x32, 0x63, 0x2d, 0x63, 0x61, 0x72, 0x2d,
	0x61, 0x70, 0x69, 0x2f, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_brandModel_model_proto_rawDescOnce sync.Once
	file_brandModel_model_proto_rawDescData = file_brandModel_model_proto_rawDesc
)

func file_brandModel_model_proto_rawDescGZIP() []byte {
	file_brandModel_model_proto_rawDescOnce.Do(func() {
		file_brandModel_model_proto_rawDescData = protoimpl.X.CompressGZIP(file_brandModel_model_proto_rawDescData)
	})
	return file_brandModel_model_proto_rawDescData
}

var file_brandModel_model_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_brandModel_model_proto_goTypes = []interface{}{
	(*ModelDto)(nil),          // 0: brandModel.ModelDto
	(*common.IdLocalDTO)(nil), // 1: common.IdLocalDTO
	(*common.Response)(nil),   // 2: common.Response
}
var file_brandModel_model_proto_depIdxs = []int32{
	0, // 0: brandModel.Model.Add:input_type -> brandModel.ModelDto
	1, // 1: brandModel.Model.Delete:input_type -> common.IdLocalDTO
	0, // 2: brandModel.Model.Update:input_type -> brandModel.ModelDto
	1, // 3: brandModel.Model.Get:input_type -> common.IdLocalDTO
	1, // 4: brandModel.Model.ListByBrandId:input_type -> common.IdLocalDTO
	1, // 5: brandModel.Model.ListModelHasCar:input_type -> common.IdLocalDTO
	2, // 6: brandModel.Model.Add:output_type -> common.Response
	2, // 7: brandModel.Model.Delete:output_type -> common.Response
	2, // 8: brandModel.Model.Update:output_type -> common.Response
	2, // 9: brandModel.Model.Get:output_type -> common.Response
	2, // 10: brandModel.Model.ListByBrandId:output_type -> common.Response
	2, // 11: brandModel.Model.ListModelHasCar:output_type -> common.Response
	6, // [6:12] is the sub-list for method output_type
	0, // [0:6] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_brandModel_model_proto_init() }
func file_brandModel_model_proto_init() {
	if File_brandModel_model_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_brandModel_model_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModelDto); i {
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
			RawDescriptor: file_brandModel_model_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_brandModel_model_proto_goTypes,
		DependencyIndexes: file_brandModel_model_proto_depIdxs,
		MessageInfos:      file_brandModel_model_proto_msgTypes,
	}.Build()
	File_brandModel_model_proto = out.File
	file_brandModel_model_proto_rawDesc = nil
	file_brandModel_model_proto_goTypes = nil
	file_brandModel_model_proto_depIdxs = nil
}
