// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.14.0
// source: scheduleTask/scheduleTask.proto

package ScheduleTask

import (
	common "github.com/cargod-bj/b2c-proto-common/common"
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
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

var File_scheduleTask_scheduleTask_proto protoreflect.FileDescriptor

var file_scheduleTask_scheduleTask_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x2f, 0x73,
	0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0c, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x1a,
	0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x61, 0x72, 0x67,
	0x6f, 0x64, 0x2d, 0x62, 0x6a, 0x2f, 0x62, 0x32, 0x63, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x85, 0x04, 0x0a, 0x0c, 0x53,
	0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x3c, 0x0a, 0x14, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x54,
	0x61, 0x73, 0x6b, 0x12, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x44, 0x74, 0x6f, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x10, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x10, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x44, 0x74, 0x6f, 0x1a,
	0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x18, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72,
	0x4c, 0x69, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x61, 0x6c, 0x12,
	0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x44, 0x74,
	0x6f, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x14, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64,
	0x73, 0x43, 0x61, 0x63, 0x68, 0x65, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x12, 0x10, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x44, 0x74, 0x6f, 0x1a,
	0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x44, 0x0a, 0x1c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72,
	0x56, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x75, 0x69, 0x64, 0x65, 0x73, 0x53,
	0x79, 0x6e, 0x63, 0x12, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x44, 0x74, 0x6f, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x18, 0x50, 0x75, 0x62,
	0x6c, 0x69, 0x73, 0x68, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x43, 0x61, 0x72, 0x43, 0x61, 0x6d,
	0x70, 0x61, 0x69, 0x67, 0x6e, 0x12, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x44, 0x74, 0x6f, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x10, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x4c, 0x69, 0x62, 0x53, 0x79, 0x6e, 0x63, 0x12,
	0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x44, 0x74,
	0x6f, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x13, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x43, 0x61, 0x72, 0x57, 0x61, 0x72, 0x72, 0x61, 0x6e, 0x74, 0x79, 0x12, 0x10, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x44, 0x74, 0x6f, 0x1a, 0x10,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x63, 0x61, 0x72, 0x67, 0x6f, 0x64, 0x2d, 0x62, 0x6a, 0x2f, 0x62, 0x32, 0x63, 0x2d, 0x63,
	0x61, 0x72, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x54,
	0x61, 0x73, 0x6b, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_scheduleTask_scheduleTask_proto_goTypes = []interface{}{
	(*common.EmptyDto)(nil), // 0: common.EmptyDto
	(*common.Response)(nil), // 1: common.Response
}
var file_scheduleTask_scheduleTask_proto_depIdxs = []int32{
	0, // 0: ScheduleTask.ScheduleTask.UpdateBrandModelTask:input_type -> common.EmptyDto
	0, // 1: ScheduleTask.ScheduleTask.UpdateEnumConfig:input_type -> common.EmptyDto
	0, // 2: ScheduleTask.ScheduleTask.UpdateCarListStatistical:input_type -> common.EmptyDto
	0, // 3: ScheduleTask.ScheduleTask.KeywordsCacheRefresh:input_type -> common.EmptyDto
	0, // 4: ScheduleTask.ScheduleTask.UpdateCarValuationGuidesSync:input_type -> common.EmptyDto
	0, // 5: ScheduleTask.ScheduleTask.PublishActiveCarCampaign:input_type -> common.EmptyDto
	0, // 6: ScheduleTask.ScheduleTask.UpdateCarLibSync:input_type -> common.EmptyDto
	0, // 7: ScheduleTask.ScheduleTask.ValidateCarWarranty:input_type -> common.EmptyDto
	1, // 8: ScheduleTask.ScheduleTask.UpdateBrandModelTask:output_type -> common.Response
	1, // 9: ScheduleTask.ScheduleTask.UpdateEnumConfig:output_type -> common.Response
	1, // 10: ScheduleTask.ScheduleTask.UpdateCarListStatistical:output_type -> common.Response
	1, // 11: ScheduleTask.ScheduleTask.KeywordsCacheRefresh:output_type -> common.Response
	1, // 12: ScheduleTask.ScheduleTask.UpdateCarValuationGuidesSync:output_type -> common.Response
	1, // 13: ScheduleTask.ScheduleTask.PublishActiveCarCampaign:output_type -> common.Response
	1, // 14: ScheduleTask.ScheduleTask.UpdateCarLibSync:output_type -> common.Response
	1, // 15: ScheduleTask.ScheduleTask.ValidateCarWarranty:output_type -> common.Response
	8, // [8:16] is the sub-list for method output_type
	0, // [0:8] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_scheduleTask_scheduleTask_proto_init() }
func file_scheduleTask_scheduleTask_proto_init() {
	if File_scheduleTask_scheduleTask_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_scheduleTask_scheduleTask_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_scheduleTask_scheduleTask_proto_goTypes,
		DependencyIndexes: file_scheduleTask_scheduleTask_proto_depIdxs,
	}.Build()
	File_scheduleTask_scheduleTask_proto = out.File
	file_scheduleTask_scheduleTask_proto_rawDesc = nil
	file_scheduleTask_scheduleTask_proto_goTypes = nil
	file_scheduleTask_scheduleTask_proto_depIdxs = nil
}
