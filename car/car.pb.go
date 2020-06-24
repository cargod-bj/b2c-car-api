// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0-devel
// 	protoc        v3.7.1
// source: car/car.proto

package carProto

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

type SourceParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page   *common.Page `protobuf:"bytes,1,opt,name=page,proto3" json:"page,omitempty"`
	LeadId string       `protobuf:"bytes,2,opt,name=lead_id,json=leadId,proto3" json:"lead_id,omitempty"`
}

func (x *SourceParams) Reset() {
	*x = SourceParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_car_car_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SourceParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SourceParams) ProtoMessage() {}

func (x *SourceParams) ProtoReflect() protoreflect.Message {
	mi := &file_car_car_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SourceParams.ProtoReflect.Descriptor instead.
func (*SourceParams) Descriptor() ([]byte, []int) {
	return file_car_car_proto_rawDescGZIP(), []int{0}
}

func (x *SourceParams) GetPage() *common.Page {
	if x != nil {
		return x.Page
	}
	return nil
}

func (x *SourceParams) GetLeadId() string {
	if x != nil {
		return x.LeadId
	}
	return ""
}

type CarListParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page  *common.Page `protobuf:"bytes,1,opt,name=page,proto3" json:"page,omitempty"`
	Local string       `protobuf:"bytes,2,opt,name=local,proto3" json:"local,omitempty"`
}

func (x *CarListParams) Reset() {
	*x = CarListParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_car_car_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CarListParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CarListParams) ProtoMessage() {}

func (x *CarListParams) ProtoReflect() protoreflect.Message {
	mi := &file_car_car_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CarListParams.ProtoReflect.Descriptor instead.
func (*CarListParams) Descriptor() ([]byte, []int) {
	return file_car_car_proto_rawDescGZIP(), []int{1}
}

func (x *CarListParams) GetPage() *common.Page {
	if x != nil {
		return x.Page
	}
	return nil
}

func (x *CarListParams) GetLocal() string {
	if x != nil {
		return x.Local
	}
	return ""
}

type CarSourceDto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Location        string `protobuf:"bytes,1,opt,name=location,proto3" json:"location,omitempty"`
	InspectionId    uint64 `protobuf:"varint,2,opt,name=inspection_id,json=inspectionId,proto3" json:"inspection_id,omitempty"`
	LeadId          uint64 `protobuf:"varint,3,opt,name=lead_id,json=leadId,proto3" json:"lead_id,omitempty"`
	InspectionDate  uint64 `protobuf:"varint,4,opt,name=inspection_date,json=inspectionDate,proto3" json:"inspection_date,omitempty"`
	Inspector       string `protobuf:"bytes,5,opt,name=inspector,proto3" json:"inspector,omitempty"`
	CarBrand        string `protobuf:"bytes,6,opt,name=car_brand,json=carBrand,proto3" json:"car_brand,omitempty"`
	CarModel        string `protobuf:"bytes,7,opt,name=car_model,json=carModel,proto3" json:"car_model,omitempty"`
	CarVariant      string `protobuf:"bytes,8,opt,name=car_variant,json=carVariant,proto3" json:"car_variant,omitempty"`
	CarEngine       string `protobuf:"bytes,9,opt,name=car_engine,json=carEngine,proto3" json:"car_engine,omitempty"`
	CarYear         string `protobuf:"bytes,10,opt,name=car_year,json=carYear,proto3" json:"car_year,omitempty"`
	CarMileage      string `protobuf:"bytes,11,opt,name=car_mileage,json=carMileage,proto3" json:"car_mileage,omitempty"`
	CarTransmission string `protobuf:"bytes,12,opt,name=car_transmission,json=carTransmission,proto3" json:"car_transmission,omitempty"`
	LicensePlate    string `protobuf:"bytes,13,opt,name=license_plate,json=licensePlate,proto3" json:"license_plate,omitempty"`
	Status          int32  `protobuf:"varint,14,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *CarSourceDto) Reset() {
	*x = CarSourceDto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_car_car_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CarSourceDto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CarSourceDto) ProtoMessage() {}

func (x *CarSourceDto) ProtoReflect() protoreflect.Message {
	mi := &file_car_car_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CarSourceDto.ProtoReflect.Descriptor instead.
func (*CarSourceDto) Descriptor() ([]byte, []int) {
	return file_car_car_proto_rawDescGZIP(), []int{2}
}

func (x *CarSourceDto) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *CarSourceDto) GetInspectionId() uint64 {
	if x != nil {
		return x.InspectionId
	}
	return 0
}

func (x *CarSourceDto) GetLeadId() uint64 {
	if x != nil {
		return x.LeadId
	}
	return 0
}

func (x *CarSourceDto) GetInspectionDate() uint64 {
	if x != nil {
		return x.InspectionDate
	}
	return 0
}

func (x *CarSourceDto) GetInspector() string {
	if x != nil {
		return x.Inspector
	}
	return ""
}

func (x *CarSourceDto) GetCarBrand() string {
	if x != nil {
		return x.CarBrand
	}
	return ""
}

func (x *CarSourceDto) GetCarModel() string {
	if x != nil {
		return x.CarModel
	}
	return ""
}

func (x *CarSourceDto) GetCarVariant() string {
	if x != nil {
		return x.CarVariant
	}
	return ""
}

func (x *CarSourceDto) GetCarEngine() string {
	if x != nil {
		return x.CarEngine
	}
	return ""
}

func (x *CarSourceDto) GetCarYear() string {
	if x != nil {
		return x.CarYear
	}
	return ""
}

func (x *CarSourceDto) GetCarMileage() string {
	if x != nil {
		return x.CarMileage
	}
	return ""
}

func (x *CarSourceDto) GetCarTransmission() string {
	if x != nil {
		return x.CarTransmission
	}
	return ""
}

func (x *CarSourceDto) GetLicensePlate() string {
	if x != nil {
		return x.LicensePlate
	}
	return ""
}

func (x *CarSourceDto) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

// 新增车辆dto
type CarDto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 车辆id，如果为新增，则此字段为0，否则为车辆id
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// 车辆的创建时间，如果为新增，或更新，则此字段无效
	CreateTime uint64 `protobuf:"varint,2,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// 车辆的更新时间，如果为新增，或更新，则此字段无效
	UpdateTime uint64 `protobuf:"varint,3,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	// 检查报告id
	InspectionId uint64 `protobuf:"varint,101,opt,name=inspection_id,json=inspectionId,proto3" json:"inspection_id,omitempty"`
	// 车辆状态
	Status uint32 `protobuf:"varint,102,opt,name=status,proto3" json:"status,omitempty"`
	// 门店id
	StoreId uint64 `protobuf:"varint,103,opt,name=store_id,json=storeId,proto3" json:"store_id,omitempty"`
	// 停车场id
	ParkingId uint64 `protobuf:"varint,104,opt,name=parking_id,json=parkingId,proto3" json:"parking_id,omitempty"`
	// 线索id
	LeadId uint64 `protobuf:"varint,105,opt,name=lead_id,json=leadId,proto3" json:"lead_id,omitempty"`
	// 品牌id
	BrandId uint32 `protobuf:"varint,201,opt,name=brand_id,json=brandId,proto3" json:"brand_id,omitempty"`
	// 车型id
	ModelId uint32 `protobuf:"varint,202,opt,name=model_id,json=modelId,proto3" json:"model_id,omitempty"`
	// 车variant
	CarVariant string `protobuf:"bytes,203,opt,name=car_variant,json=carVariant,proto3" json:"car_variant,omitempty"`
	// 品牌id
	BrandName string `protobuf:"bytes,204,opt,name=brand_name,json=brandName,proto3" json:"brand_name,omitempty"`
	// 车型id
	ModelName string `protobuf:"bytes,205,opt,name=model_name,json=modelName,proto3" json:"model_name,omitempty"`
	// 车variant
	VariantId uint32 `protobuf:"varint,206,opt,name=variant_id,json=variantId,proto3" json:"variant_id,omitempty"`
	// 车照片url
	Image string `protobuf:"bytes,301,opt,name=image,proto3" json:"image,omitempty"`
	// 排量
	CarEngine string `protobuf:"bytes,401,opt,name=car_engine,json=carEngine,proto3" json:"car_engine,omitempty"`
	// 年份
	CarYear uint32 `protobuf:"varint,402,opt,name=car_year,json=carYear,proto3" json:"car_year,omitempty"`
	// 里程
	CarMileage uint32 `protobuf:"varint,403,opt,name=car_mileage,json=carMileage,proto3" json:"car_mileage,omitempty"`
	// 驾驶类型 手动,自动
	CarTransmission      uint32 `protobuf:"varint,404,opt,name=car_transmission,json=carTransmission,proto3" json:"car_transmission,omitempty"`
	CarTransmissionName  string `protobuf:"bytes,405,opt,name=car_transmission_name,json=carTransmissionName,proto3" json:"car_transmission_name,omitempty"`
	RegistrationType     uint32 `protobuf:"varint,406,opt,name=registration_type,json=registrationType,proto3" json:"registration_type,omitempty"`
	RegistrationTypeName string `protobuf:"bytes,407,opt,name=registration_type_name,json=registrationTypeName,proto3" json:"registration_type_name,omitempty"`
	RegistrationDate     uint64 `protobuf:"varint,408,opt,name=registration_date,json=registrationDate,proto3" json:"registration_date,omitempty"`
	Seat                 uint32 `protobuf:"varint,409,opt,name=seat,proto3" json:"seat,omitempty"`
	Color                uint32 `protobuf:"varint,410,opt,name=color,proto3" json:"color,omitempty"`
	FuelType             uint32 `protobuf:"varint,411,opt,name=fuel_type,json=fuelType,proto3" json:"fuel_type,omitempty"`
	FuelTypeName         string `protobuf:"bytes,412,opt,name=fuel_type_name,json=fuelTypeName,proto3" json:"fuel_type_name,omitempty"`
	ColorName            string `protobuf:"bytes,413,opt,name=colorName,proto3" json:"colorName,omitempty"`
	// 卖价
	ExpSellingPrice float64 `protobuf:"fixed64,501,opt,name=exp_selling_price,json=expSellingPrice,proto3" json:"exp_selling_price,omitempty"`
	// 来源
	Source string `protobuf:"bytes,601,opt,name=source,proto3" json:"source,omitempty"`
	// 城市地区
	Location string `protobuf:"bytes,701,opt,name=location,proto3" json:"location,omitempty"`
	// 城市地区id
	LocationId uint64 `protobuf:"varint,702,opt,name=location_id,json=locationId,proto3" json:"location_id,omitempty"`
}

func (x *CarDto) Reset() {
	*x = CarDto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_car_car_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CarDto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CarDto) ProtoMessage() {}

func (x *CarDto) ProtoReflect() protoreflect.Message {
	mi := &file_car_car_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CarDto.ProtoReflect.Descriptor instead.
func (*CarDto) Descriptor() ([]byte, []int) {
	return file_car_car_proto_rawDescGZIP(), []int{3}
}

func (x *CarDto) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CarDto) GetCreateTime() uint64 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

func (x *CarDto) GetUpdateTime() uint64 {
	if x != nil {
		return x.UpdateTime
	}
	return 0
}

func (x *CarDto) GetInspectionId() uint64 {
	if x != nil {
		return x.InspectionId
	}
	return 0
}

func (x *CarDto) GetStatus() uint32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *CarDto) GetStoreId() uint64 {
	if x != nil {
		return x.StoreId
	}
	return 0
}

func (x *CarDto) GetParkingId() uint64 {
	if x != nil {
		return x.ParkingId
	}
	return 0
}

func (x *CarDto) GetLeadId() uint64 {
	if x != nil {
		return x.LeadId
	}
	return 0
}

func (x *CarDto) GetBrandId() uint32 {
	if x != nil {
		return x.BrandId
	}
	return 0
}

func (x *CarDto) GetModelId() uint32 {
	if x != nil {
		return x.ModelId
	}
	return 0
}

func (x *CarDto) GetCarVariant() string {
	if x != nil {
		return x.CarVariant
	}
	return ""
}

func (x *CarDto) GetBrandName() string {
	if x != nil {
		return x.BrandName
	}
	return ""
}

func (x *CarDto) GetModelName() string {
	if x != nil {
		return x.ModelName
	}
	return ""
}

func (x *CarDto) GetVariantId() uint32 {
	if x != nil {
		return x.VariantId
	}
	return 0
}

func (x *CarDto) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *CarDto) GetCarEngine() string {
	if x != nil {
		return x.CarEngine
	}
	return ""
}

func (x *CarDto) GetCarYear() uint32 {
	if x != nil {
		return x.CarYear
	}
	return 0
}

func (x *CarDto) GetCarMileage() uint32 {
	if x != nil {
		return x.CarMileage
	}
	return 0
}

func (x *CarDto) GetCarTransmission() uint32 {
	if x != nil {
		return x.CarTransmission
	}
	return 0
}

func (x *CarDto) GetCarTransmissionName() string {
	if x != nil {
		return x.CarTransmissionName
	}
	return ""
}

func (x *CarDto) GetRegistrationType() uint32 {
	if x != nil {
		return x.RegistrationType
	}
	return 0
}

func (x *CarDto) GetRegistrationTypeName() string {
	if x != nil {
		return x.RegistrationTypeName
	}
	return ""
}

func (x *CarDto) GetRegistrationDate() uint64 {
	if x != nil {
		return x.RegistrationDate
	}
	return 0
}

func (x *CarDto) GetSeat() uint32 {
	if x != nil {
		return x.Seat
	}
	return 0
}

func (x *CarDto) GetColor() uint32 {
	if x != nil {
		return x.Color
	}
	return 0
}

func (x *CarDto) GetFuelType() uint32 {
	if x != nil {
		return x.FuelType
	}
	return 0
}

func (x *CarDto) GetFuelTypeName() string {
	if x != nil {
		return x.FuelTypeName
	}
	return ""
}

func (x *CarDto) GetColorName() string {
	if x != nil {
		return x.ColorName
	}
	return ""
}

func (x *CarDto) GetExpSellingPrice() float64 {
	if x != nil {
		return x.ExpSellingPrice
	}
	return 0
}

func (x *CarDto) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *CarDto) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *CarDto) GetLocationId() uint64 {
	if x != nil {
		return x.LocationId
	}
	return 0
}

// 车辆的id
type CarIdsDto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 车辆id
	Ids   []uint64 `protobuf:"varint,1,rep,packed,name=ids,proto3" json:"ids,omitempty"`
	Local string   `protobuf:"bytes,2,opt,name=local,proto3" json:"local,omitempty"`
}

func (x *CarIdsDto) Reset() {
	*x = CarIdsDto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_car_car_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CarIdsDto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CarIdsDto) ProtoMessage() {}

func (x *CarIdsDto) ProtoReflect() protoreflect.Message {
	mi := &file_car_car_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CarIdsDto.ProtoReflect.Descriptor instead.
func (*CarIdsDto) Descriptor() ([]byte, []int) {
	return file_car_car_proto_rawDescGZIP(), []int{4}
}

func (x *CarIdsDto) GetIds() []uint64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *CarIdsDto) GetLocal() string {
	if x != nil {
		return x.Local
	}
	return ""
}

type CarDtoList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CarDtoList []*CarDto `protobuf:"bytes,1,rep,name=CarDtoList,proto3" json:"CarDtoList,omitempty"`
}

func (x *CarDtoList) Reset() {
	*x = CarDtoList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_car_car_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CarDtoList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CarDtoList) ProtoMessage() {}

func (x *CarDtoList) ProtoReflect() protoreflect.Message {
	mi := &file_car_car_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CarDtoList.ProtoReflect.Descriptor instead.
func (*CarDtoList) Descriptor() ([]byte, []int) {
	return file_car_car_proto_rawDescGZIP(), []int{5}
}

func (x *CarDtoList) GetCarDtoList() []*CarDto {
	if x != nil {
		return x.CarDtoList
	}
	return nil
}

var File_car_car_proto protoreflect.FileDescriptor

var file_car_car_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x63, 0x61, 0x72, 0x2f, 0x63, 0x61, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x03, 0x63, 0x61, 0x72, 0x1a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x63, 0x61, 0x72, 0x67, 0x6f, 0x64, 0x2d, 0x62, 0x6a, 0x2f, 0x62, 0x32, 0x63, 0x2d, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x49, 0x0a, 0x0c, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12,
	0x20, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x04, 0x70, 0x61, 0x67,
	0x65, 0x12, 0x17, 0x0a, 0x07, 0x6c, 0x65, 0x61, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x6c, 0x65, 0x61, 0x64, 0x49, 0x64, 0x22, 0x47, 0x0a, 0x0d, 0x43, 0x61,
	0x72, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x20, 0x0a, 0x04, 0x70,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x6f,
	0x63, 0x61, 0x6c, 0x22, 0xcd, 0x03, 0x0a, 0x0c, 0x43, 0x61, 0x72, 0x53, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x44, 0x74, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x23, 0x0a, 0x0d, 0x69, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x69, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x6c, 0x65, 0x61, 0x64, 0x5f, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x6c, 0x65, 0x61, 0x64, 0x49, 0x64, 0x12, 0x27,
	0x0a, 0x0f, 0x69, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0e, 0x69, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x6e, 0x73, 0x70, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x6e, 0x73, 0x70,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x61, 0x72, 0x5f, 0x62, 0x72, 0x61,
	0x6e, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x72, 0x42, 0x72, 0x61,
	0x6e, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x61, 0x72, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x72, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12,
	0x1f, 0x0a, 0x0b, 0x63, 0x61, 0x72, 0x5f, 0x76, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x61, 0x72, 0x56, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74,
	0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x61, 0x72, 0x5f, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x61, 0x72, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x12,
	0x19, 0x0a, 0x08, 0x63, 0x61, 0x72, 0x5f, 0x79, 0x65, 0x61, 0x72, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x63, 0x61, 0x72, 0x59, 0x65, 0x61, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x61,
	0x72, 0x5f, 0x6d, 0x69, 0x6c, 0x65, 0x61, 0x67, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x63, 0x61, 0x72, 0x4d, 0x69, 0x6c, 0x65, 0x61, 0x67, 0x65, 0x12, 0x29, 0x0a, 0x10, 0x63,
	0x61, 0x72, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x63, 0x61, 0x72, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73,
	0x65, 0x5f, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6c,
	0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x50, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x22, 0xa2, 0x08, 0x0a, 0x06, 0x43, 0x61, 0x72, 0x44, 0x74, 0x6f, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1f,
	0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x1f, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x23, 0x0a, 0x0d, 0x69, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69,
	0x64, 0x18, 0x65, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x69, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x66, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x19, 0x0a,
	0x08, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x67, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x07, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x72, 0x6b,
	0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x68, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x70, 0x61,
	0x72, 0x6b, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x6c, 0x65, 0x61, 0x64, 0x5f,
	0x69, 0x64, 0x18, 0x69, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x6c, 0x65, 0x61, 0x64, 0x49, 0x64,
	0x12, 0x1a, 0x0a, 0x08, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x5f, 0x69, 0x64, 0x18, 0xc9, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x07, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0xca, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x07, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x61, 0x72, 0x5f,
	0x76, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x18, 0xcb, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x63, 0x61, 0x72, 0x56, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x62, 0x72,
	0x61, 0x6e, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0xcc, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0xcd, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x76, 0x61,
	0x72, 0x69, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0xce, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x09, 0x76, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x05, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x18, 0xad, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x61, 0x72, 0x5f, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x18,
	0x91, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x61, 0x72, 0x45, 0x6e, 0x67, 0x69, 0x6e,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x72, 0x5f, 0x79, 0x65, 0x61, 0x72, 0x18, 0x92, 0x03,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x63, 0x61, 0x72, 0x59, 0x65, 0x61, 0x72, 0x12, 0x20, 0x0a,
	0x0b, 0x63, 0x61, 0x72, 0x5f, 0x6d, 0x69, 0x6c, 0x65, 0x61, 0x67, 0x65, 0x18, 0x93, 0x03, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0a, 0x63, 0x61, 0x72, 0x4d, 0x69, 0x6c, 0x65, 0x61, 0x67, 0x65, 0x12,
	0x2a, 0x0a, 0x10, 0x63, 0x61, 0x72, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x94, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x63, 0x61, 0x72, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x33, 0x0a, 0x15, 0x63,
	0x61, 0x72, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x95, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x63, 0x61, 0x72,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x2c, 0x0a, 0x11, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x96, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x10, 0x72, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x35,
	0x0a, 0x16, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x97, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x14, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2c, 0x0a, 0x11, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x98, 0x03, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x10, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44,
	0x61, 0x74, 0x65, 0x12, 0x13, 0x0a, 0x04, 0x73, 0x65, 0x61, 0x74, 0x18, 0x99, 0x03, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x04, 0x73, 0x65, 0x61, 0x74, 0x12, 0x15, 0x0a, 0x05, 0x63, 0x6f, 0x6c, 0x6f,
	0x72, 0x18, 0x9a, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x12,
	0x1c, 0x0a, 0x09, 0x66, 0x75, 0x65, 0x6c, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x9b, 0x03, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x08, 0x66, 0x75, 0x65, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x12, 0x25, 0x0a,
	0x0e, 0x66, 0x75, 0x65, 0x6c, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x9c, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x66, 0x75, 0x65, 0x6c, 0x54, 0x79, 0x70, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x09, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x9d, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x2b, 0x0a, 0x11, 0x65, 0x78, 0x70, 0x5f, 0x73, 0x65, 0x6c, 0x6c, 0x69,
	0x6e, 0x67, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0xf5, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x0f, 0x65, 0x78, 0x70, 0x53, 0x65, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x69, 0x63, 0x65,
	0x12, 0x17, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0xd9, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x1b, 0x0a, 0x08, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0xbd, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0xbe, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x33, 0x0a, 0x09, 0x43, 0x61, 0x72, 0x49,
	0x64, 0x73, 0x44, 0x74, 0x6f, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x04, 0x52, 0x03, 0x69, 0x64, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x6f, 0x63, 0x61, 0x6c,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x22, 0x39, 0x0a,
	0x0a, 0x43, 0x61, 0x72, 0x44, 0x74, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2b, 0x0a, 0x0a, 0x43,
	0x61, 0x72, 0x44, 0x74, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0b, 0x2e, 0x63, 0x61, 0x72, 0x2e, 0x43, 0x61, 0x72, 0x44, 0x74, 0x6f, 0x52, 0x0a, 0x43, 0x61,
	0x72, 0x44, 0x74, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x32, 0xca, 0x02, 0x0a, 0x03, 0x43, 0x61, 0x72,
	0x12, 0x26, 0x0a, 0x03, 0x41, 0x64, 0x64, 0x12, 0x0b, 0x2e, 0x63, 0x61, 0x72, 0x2e, 0x43, 0x61,
	0x72, 0x44, 0x74, 0x6f, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2c, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x12, 0x0e, 0x2e, 0x63, 0x61, 0x72, 0x2e, 0x43, 0x61, 0x72, 0x49, 0x64, 0x73, 0x44,
	0x74, 0x6f, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x29, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x12, 0x0b, 0x2e, 0x63, 0x61, 0x72, 0x2e, 0x43, 0x61, 0x72, 0x44, 0x74, 0x6f, 0x1a, 0x10, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x29, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x0e, 0x2e, 0x63, 0x61, 0x72, 0x2e, 0x43,
	0x61, 0x72, 0x49, 0x64, 0x73, 0x44, 0x74, 0x6f, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2e, 0x0a, 0x04,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x12, 0x2e, 0x63, 0x61, 0x72, 0x2e, 0x43, 0x61, 0x72, 0x4c, 0x69,
	0x73, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x33, 0x0a, 0x0a,
	0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x11, 0x2e, 0x63, 0x61, 0x72,
	0x2e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x10, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x32, 0x0a, 0x0d, 0x41, 0x64, 0x64, 0x46, 0x72, 0x6f, 0x6d, 0x53, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x12, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x49, 0x64, 0x44, 0x74,
	0x6f, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x61, 0x72, 0x67, 0x6f, 0x64, 0x2d, 0x62, 0x6a, 0x2f, 0x62, 0x32,
	0x63, 0x2d, 0x63, 0x61, 0x72, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x61, 0x72, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_car_car_proto_rawDescOnce sync.Once
	file_car_car_proto_rawDescData = file_car_car_proto_rawDesc
)

func file_car_car_proto_rawDescGZIP() []byte {
	file_car_car_proto_rawDescOnce.Do(func() {
		file_car_car_proto_rawDescData = protoimpl.X.CompressGZIP(file_car_car_proto_rawDescData)
	})
	return file_car_car_proto_rawDescData
}

var file_car_car_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_car_car_proto_goTypes = []interface{}{
	(*SourceParams)(nil),    // 0: car.SourceParams
	(*CarListParams)(nil),   // 1: car.CarListParams
	(*CarSourceDto)(nil),    // 2: car.CarSourceDto
	(*CarDto)(nil),          // 3: car.CarDto
	(*CarIdsDto)(nil),       // 4: car.CarIdsDto
	(*CarDtoList)(nil),      // 5: car.CarDtoList
	(*common.Page)(nil),     // 6: common.Page
	(*common.IdDto)(nil),    // 7: common.IdDto
	(*common.Response)(nil), // 8: common.Response
}
var file_car_car_proto_depIdxs = []int32{
	6,  // 0: car.SourceParams.page:type_name -> common.Page
	6,  // 1: car.CarListParams.page:type_name -> common.Page
	3,  // 2: car.CarDtoList.CarDtoList:type_name -> car.CarDto
	3,  // 3: car.Car.Add:input_type -> car.CarDto
	4,  // 4: car.Car.Delete:input_type -> car.CarIdsDto
	3,  // 5: car.Car.Update:input_type -> car.CarDto
	4,  // 6: car.Car.Get:input_type -> car.CarIdsDto
	1,  // 7: car.Car.List:input_type -> car.CarListParams
	0,  // 8: car.Car.SourceList:input_type -> car.SourceParams
	7,  // 9: car.Car.AddFromSource:input_type -> common.IdDto
	8,  // 10: car.Car.Add:output_type -> common.Response
	8,  // 11: car.Car.Delete:output_type -> common.Response
	8,  // 12: car.Car.Update:output_type -> common.Response
	8,  // 13: car.Car.Get:output_type -> common.Response
	8,  // 14: car.Car.List:output_type -> common.Response
	8,  // 15: car.Car.SourceList:output_type -> common.Response
	8,  // 16: car.Car.AddFromSource:output_type -> common.Response
	10, // [10:17] is the sub-list for method output_type
	3,  // [3:10] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_car_car_proto_init() }
func file_car_car_proto_init() {
	if File_car_car_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_car_car_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SourceParams); i {
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
		file_car_car_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CarListParams); i {
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
		file_car_car_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CarSourceDto); i {
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
		file_car_car_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CarDto); i {
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
		file_car_car_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CarIdsDto); i {
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
		file_car_car_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CarDtoList); i {
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
			RawDescriptor: file_car_car_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_car_car_proto_goTypes,
		DependencyIndexes: file_car_car_proto_depIdxs,
		MessageInfos:      file_car_car_proto_msgTypes,
	}.Build()
	File_car_car_proto = out.File
	file_car_car_proto_rawDesc = nil
	file_car_car_proto_goTypes = nil
	file_car_car_proto_depIdxs = nil
}
