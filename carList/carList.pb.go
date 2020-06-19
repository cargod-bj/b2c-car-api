// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.11.4
// source: carList/carList.proto

package carList

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

// 新增上架车辆dto
type CarListDto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 车辆id，如果为新增，则此字段为0，否则为车辆id
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	//车辆在car表里的id
	CarId uint64 `protobuf:"varint,2,opt,name=car_id,json=carId,proto3" json:"car_id,omitempty"`
	// 车辆的创建时间，如果为新增，或更新，则此字段无效
	CreateTime uint64 `protobuf:"varint,3,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// 车辆的更新时间，如果为新增，或更新，则此字段无效
	UpdateTime uint64 `protobuf:"varint,4,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	// 检查报告id
	InspectionId uint64 `protobuf:"varint,101,opt,name=inspection_id,json=inspectionId,proto3" json:"inspection_id,omitempty"`
	// 车辆状态
	Status uint32 `protobuf:"varint,102,opt,name=status,proto3" json:"status,omitempty"`
	// 门店id
	StoreId uint64 `protobuf:"varint,103,opt,name=store_id,json=storeId,proto3" json:"store_id,omitempty"`
	// 停车场id
	ParkingId uint64 `protobuf:"varint,104,opt,name=parking_id,json=parkingId,proto3" json:"parking_id,omitempty"`
	// 品牌id
	BrandId uint32 `protobuf:"varint,201,opt,name=brand_id,json=brandId,proto3" json:"brand_id,omitempty"`
	// 车型id
	ModelId uint32 `protobuf:"varint,202,opt,name=model_id,json=modelId,proto3" json:"model_id,omitempty"`
	// 车variant
	CarVariant string `protobuf:"bytes,203,opt,name=car_variant,json=carVariant,proto3" json:"car_variant,omitempty"`
	// 车照片url
	Image string `protobuf:"bytes,301,opt,name=image,proto3" json:"image,omitempty"`
	// 排量
	CarEngine string `protobuf:"bytes,401,opt,name=car_engine,json=carEngine,proto3" json:"car_engine,omitempty"`
	// 年份
	CarYear uint32 `protobuf:"varint,402,opt,name=car_year,json=carYear,proto3" json:"car_year,omitempty"`
	// 里程
	CarMilage uint32 `protobuf:"varint,403,opt,name=car_milage,json=carMilage,proto3" json:"car_milage,omitempty"`
	// 驾驶类型 手动,自动
	CarTransmission string `protobuf:"bytes,404,opt,name=car_transmission,json=carTransmission,proto3" json:"car_transmission,omitempty"`
	// 卖价
	ExpSellingPrice float64 `protobuf:"fixed64,501,opt,name=exp_selling_price,json=expSellingPrice,proto3" json:"exp_selling_price,omitempty"`
	// 来源
	Source string `protobuf:"bytes,601,opt,name=source,proto3" json:"source,omitempty"`
	// 城市地区
	Location string `protobuf:"bytes,701,opt,name=location,proto3" json:"location,omitempty"`
	// 城市地区
	LocationAddress string `protobuf:"bytes,702,opt,name=location_address,json=locationAddress,proto3" json:"location_address,omitempty"`
	// 是否是Dealer发布
	IsDealer uint32 `protobuf:"varint,703,opt,name=is_dealer,json=isDealer,proto3" json:"is_dealer,omitempty"`
	//排序
	Seq uint32 `protobuf:"varint,704,opt,name=seq,proto3" json:"seq,omitempty"`
	//汽车类型例如suv
	CarTypeId uint32 `protobuf:"varint,707,opt,name=car_type_id,json=carTypeId,proto3" json:"car_type_id,omitempty"`
	//经销商id
	DealerId uint64 `protobuf:"varint,708,opt,name=dealer_id,json=dealerId,proto3" json:"dealer_id,omitempty"`
	//燃油类型
	FuelType uint32 `protobuf:"varint,709,opt,name=fuel_type,json=fuelType,proto3" json:"fuel_type,omitempty"`
	//座椅个数
	SeatSize uint32 `protobuf:"varint,710,opt,name=seat_size,json=seatSize,proto3" json:"seat_size,omitempty"`
	//车漆颜色
	Color uint32 `protobuf:"varint,711,opt,name=color,proto3" json:"color,omitempty"`
	//车辆特征
	CarTypeName string `protobuf:"bytes,712,opt,name=car_type_name,json=carTypeName,proto3" json:"car_type_name,omitempty"`
	ModelName   string `protobuf:"bytes,705,opt,name=model_name,json=modelName,proto3" json:"model_name,omitempty"`
	BrandName   string `protobuf:"bytes,706,opt,name=brand_name,json=brandName,proto3" json:"brand_name,omitempty"`
}

func (x *CarListDto) Reset() {
	*x = CarListDto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_carList_carList_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CarListDto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CarListDto) ProtoMessage() {}

func (x *CarListDto) ProtoReflect() protoreflect.Message {
	mi := &file_carList_carList_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CarListDto.ProtoReflect.Descriptor instead.
func (*CarListDto) Descriptor() ([]byte, []int) {
	return file_carList_carList_proto_rawDescGZIP(), []int{0}
}

func (x *CarListDto) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CarListDto) GetCarId() uint64 {
	if x != nil {
		return x.CarId
	}
	return 0
}

func (x *CarListDto) GetCreateTime() uint64 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

func (x *CarListDto) GetUpdateTime() uint64 {
	if x != nil {
		return x.UpdateTime
	}
	return 0
}

func (x *CarListDto) GetInspectionId() uint64 {
	if x != nil {
		return x.InspectionId
	}
	return 0
}

func (x *CarListDto) GetStatus() uint32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *CarListDto) GetStoreId() uint64 {
	if x != nil {
		return x.StoreId
	}
	return 0
}

func (x *CarListDto) GetParkingId() uint64 {
	if x != nil {
		return x.ParkingId
	}
	return 0
}

func (x *CarListDto) GetBrandId() uint32 {
	if x != nil {
		return x.BrandId
	}
	return 0
}

func (x *CarListDto) GetModelId() uint32 {
	if x != nil {
		return x.ModelId
	}
	return 0
}

func (x *CarListDto) GetCarVariant() string {
	if x != nil {
		return x.CarVariant
	}
	return ""
}

func (x *CarListDto) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *CarListDto) GetCarEngine() string {
	if x != nil {
		return x.CarEngine
	}
	return ""
}

func (x *CarListDto) GetCarYear() uint32 {
	if x != nil {
		return x.CarYear
	}
	return 0
}

func (x *CarListDto) GetCarMilage() uint32 {
	if x != nil {
		return x.CarMilage
	}
	return 0
}

func (x *CarListDto) GetCarTransmission() string {
	if x != nil {
		return x.CarTransmission
	}
	return ""
}

func (x *CarListDto) GetExpSellingPrice() float64 {
	if x != nil {
		return x.ExpSellingPrice
	}
	return 0
}

func (x *CarListDto) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *CarListDto) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *CarListDto) GetLocationAddress() string {
	if x != nil {
		return x.LocationAddress
	}
	return ""
}

func (x *CarListDto) GetIsDealer() uint32 {
	if x != nil {
		return x.IsDealer
	}
	return 0
}

func (x *CarListDto) GetSeq() uint32 {
	if x != nil {
		return x.Seq
	}
	return 0
}

func (x *CarListDto) GetCarTypeId() uint32 {
	if x != nil {
		return x.CarTypeId
	}
	return 0
}

func (x *CarListDto) GetDealerId() uint64 {
	if x != nil {
		return x.DealerId
	}
	return 0
}

func (x *CarListDto) GetFuelType() uint32 {
	if x != nil {
		return x.FuelType
	}
	return 0
}

func (x *CarListDto) GetSeatSize() uint32 {
	if x != nil {
		return x.SeatSize
	}
	return 0
}

func (x *CarListDto) GetColor() uint32 {
	if x != nil {
		return x.Color
	}
	return 0
}

func (x *CarListDto) GetCarTypeName() string {
	if x != nil {
		return x.CarTypeName
	}
	return ""
}

func (x *CarListDto) GetModelName() string {
	if x != nil {
		return x.ModelName
	}
	return ""
}

func (x *CarListDto) GetBrandName() string {
	if x != nil {
		return x.BrandName
	}
	return ""
}

//carlist查询条件
type CarListCondition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ExpSellingPriceFrom uint64   `protobuf:"varint,1,opt,name=exp_selling_price_from,json=expSellingPriceFrom,proto3" json:"exp_selling_price_from,omitempty"` //卖价
	ExpSellingPriceTo   uint64   `protobuf:"varint,2,opt,name=exp_selling_price_to,json=expSellingPriceTo,proto3" json:"exp_selling_price_to,omitempty"`
	ModelId             []uint64 `protobuf:"varint,3,rep,packed,name=model_id,json=modelId,proto3" json:"model_id,omitempty"` //车系
	ModelName           []string `protobuf:"bytes,4,rep,name=model_name,json=modelName,proto3" json:"model_name,omitempty"`
	BrandId             []uint64 `protobuf:"varint,5,rep,packed,name=brand_id,json=brandId,proto3" json:"brand_id,omitempty"` //品牌
	BrandName           []string `protobuf:"bytes,6,rep,name=brand_name,json=brandName,proto3" json:"brand_name,omitempty"`
	Variant             []string `protobuf:"bytes,7,rep,name=variant,proto3" json:"variant,omitempty"` //型号
	Status              uint32   `protobuf:"varint,8,opt,name=status,proto3" json:"status,omitempty"`
	YearFrom            uint32   `protobuf:"varint,9,opt,name=year_from,json=yearFrom,proto3" json:"year_from,omitempty"` //年份
	YearTo              uint32   `protobuf:"varint,10,opt,name=year_to,json=yearTo,proto3" json:"year_to,omitempty"`
	MilageFrom          uint64   `protobuf:"varint,11,opt,name=milage_from,json=milageFrom,proto3" json:"milage_from,omitempty"` //里程
	MilageTo            uint64   `protobuf:"varint,12,opt,name=milage_to,json=milageTo,proto3" json:"milage_to,omitempty"`
	Transmission        []string `protobuf:"bytes,13,rep,name=transmission,proto3" json:"transmission,omitempty"`                       //驾驶类型 手动，自动
	Engine              []string `protobuf:"bytes,14,rep,name=engine,proto3" json:"engine,omitempty"`                                   //排量
	Color               []uint32 `protobuf:"varint,15,rep,packed,name=color,proto3" json:"color,omitempty"`                             //颜色
	FuelType            []uint32 `protobuf:"varint,16,rep,packed,name=fuel_type,json=fuelType,proto3" json:"fuel_type,omitempty"`       //燃油
	SeatType            []uint32 `protobuf:"varint,17,rep,packed,name=seat_type,json=seatType,proto3" json:"seat_type,omitempty"`       //座位数
	ParkingId           []uint64 `protobuf:"varint,18,rep,packed,name=parking_id,json=parkingId,proto3" json:"parking_id,omitempty"`    //停车场
	StoreId             []uint64 `protobuf:"varint,20,rep,packed,name=store_id,json=storeId,proto3" json:"store_id,omitempty"`          //店铺
	Source              []string `protobuf:"bytes,22,rep,name=source,proto3" json:"source,omitempty"`                                   //来源
	Location            string   `protobuf:"bytes,23,opt,name=location,proto3" json:"location,omitempty"`                               //城市
	LocationId          []uint64 `protobuf:"varint,27,rep,packed,name=location_id,json=locationId,proto3" json:"location_id,omitempty"` //地址id
	DealerId            []uint64 `protobuf:"varint,24,rep,packed,name=dealer_id,json=dealerId,proto3" json:"dealer_id,omitempty"`       //经销商id
	CarTypeId           []uint32 `protobuf:"varint,30,rep,packed,name=car_type_id,json=carTypeId,proto3" json:"car_type_id,omitempty"`  //汽车类型
	//carlist排序条件
	SortKey  string `protobuf:"bytes,25,opt,name=sort_key,json=sortKey,proto3" json:"sort_key,omitempty"`
	SortType uint32 `protobuf:"varint,26,opt,name=sort_type,json=sortType,proto3" json:"sort_type,omitempty"`
	Page     uint32 `protobuf:"varint,28,opt,name=page,proto3" json:"page,omitempty"`
	Limit    uint32 `protobuf:"varint,29,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *CarListCondition) Reset() {
	*x = CarListCondition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_carList_carList_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CarListCondition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CarListCondition) ProtoMessage() {}

func (x *CarListCondition) ProtoReflect() protoreflect.Message {
	mi := &file_carList_carList_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CarListCondition.ProtoReflect.Descriptor instead.
func (*CarListCondition) Descriptor() ([]byte, []int) {
	return file_carList_carList_proto_rawDescGZIP(), []int{1}
}

func (x *CarListCondition) GetExpSellingPriceFrom() uint64 {
	if x != nil {
		return x.ExpSellingPriceFrom
	}
	return 0
}

func (x *CarListCondition) GetExpSellingPriceTo() uint64 {
	if x != nil {
		return x.ExpSellingPriceTo
	}
	return 0
}

func (x *CarListCondition) GetModelId() []uint64 {
	if x != nil {
		return x.ModelId
	}
	return nil
}

func (x *CarListCondition) GetModelName() []string {
	if x != nil {
		return x.ModelName
	}
	return nil
}

func (x *CarListCondition) GetBrandId() []uint64 {
	if x != nil {
		return x.BrandId
	}
	return nil
}

func (x *CarListCondition) GetBrandName() []string {
	if x != nil {
		return x.BrandName
	}
	return nil
}

func (x *CarListCondition) GetVariant() []string {
	if x != nil {
		return x.Variant
	}
	return nil
}

func (x *CarListCondition) GetStatus() uint32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *CarListCondition) GetYearFrom() uint32 {
	if x != nil {
		return x.YearFrom
	}
	return 0
}

func (x *CarListCondition) GetYearTo() uint32 {
	if x != nil {
		return x.YearTo
	}
	return 0
}

func (x *CarListCondition) GetMilageFrom() uint64 {
	if x != nil {
		return x.MilageFrom
	}
	return 0
}

func (x *CarListCondition) GetMilageTo() uint64 {
	if x != nil {
		return x.MilageTo
	}
	return 0
}

func (x *CarListCondition) GetTransmission() []string {
	if x != nil {
		return x.Transmission
	}
	return nil
}

func (x *CarListCondition) GetEngine() []string {
	if x != nil {
		return x.Engine
	}
	return nil
}

func (x *CarListCondition) GetColor() []uint32 {
	if x != nil {
		return x.Color
	}
	return nil
}

func (x *CarListCondition) GetFuelType() []uint32 {
	if x != nil {
		return x.FuelType
	}
	return nil
}

func (x *CarListCondition) GetSeatType() []uint32 {
	if x != nil {
		return x.SeatType
	}
	return nil
}

func (x *CarListCondition) GetParkingId() []uint64 {
	if x != nil {
		return x.ParkingId
	}
	return nil
}

func (x *CarListCondition) GetStoreId() []uint64 {
	if x != nil {
		return x.StoreId
	}
	return nil
}

func (x *CarListCondition) GetSource() []string {
	if x != nil {
		return x.Source
	}
	return nil
}

func (x *CarListCondition) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *CarListCondition) GetLocationId() []uint64 {
	if x != nil {
		return x.LocationId
	}
	return nil
}

func (x *CarListCondition) GetDealerId() []uint64 {
	if x != nil {
		return x.DealerId
	}
	return nil
}

func (x *CarListCondition) GetCarTypeId() []uint32 {
	if x != nil {
		return x.CarTypeId
	}
	return nil
}

func (x *CarListCondition) GetSortKey() string {
	if x != nil {
		return x.SortKey
	}
	return ""
}

func (x *CarListCondition) GetSortType() uint32 {
	if x != nil {
		return x.SortType
	}
	return 0
}

func (x *CarListCondition) GetPage() uint32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *CarListCondition) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

// 上架车辆的id
type CarListIdDto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 车辆id
	Id []uint64 `protobuf:"varint,1,rep,packed,name=id,proto3" json:"id,omitempty"`
}

func (x *CarListIdDto) Reset() {
	*x = CarListIdDto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_carList_carList_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CarListIdDto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CarListIdDto) ProtoMessage() {}

func (x *CarListIdDto) ProtoReflect() protoreflect.Message {
	mi := &file_carList_carList_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CarListIdDto.ProtoReflect.Descriptor instead.
func (*CarListIdDto) Descriptor() ([]byte, []int) {
	return file_carList_carList_proto_rawDescGZIP(), []int{2}
}

func (x *CarListIdDto) GetId() []uint64 {
	if x != nil {
		return x.Id
	}
	return nil
}

var File_carList_carList_proto protoreflect.FileDescriptor

var file_carList_carList_proto_rawDesc = []byte{
	0x0a, 0x15, 0x63, 0x61, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x2f, 0x63, 0x61, 0x72, 0x4c, 0x69, 0x73,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x63, 0x61, 0x72, 0x1a, 0x39, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x61, 0x72, 0x67, 0x6f, 0x64, 0x2d,
	0x62, 0x6a, 0x2f, 0x62, 0x32, 0x63, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9c, 0x07, 0x0a, 0x0a, 0x43, 0x61, 0x72, 0x4c,
	0x69, 0x73, 0x74, 0x44, 0x74, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x63, 0x61, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x63, 0x61, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a,
	0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1f,
	0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x23, 0x0a, 0x0d, 0x69, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64,
	0x18, 0x65, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x69, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x66,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x19, 0x0a, 0x08,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x67, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x72, 0x6b, 0x69,
	0x6e, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x68, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x70, 0x61, 0x72,
	0x6b, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x5f,
	0x69, 0x64, 0x18, 0xc9, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x62, 0x72, 0x61, 0x6e, 0x64,
	0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0xca,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x49, 0x64, 0x12, 0x20,
	0x0a, 0x0b, 0x63, 0x61, 0x72, 0x5f, 0x76, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x18, 0xcb, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x61, 0x72, 0x56, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74,
	0x12, 0x15, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0xad, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x61, 0x72, 0x5f, 0x65,
	0x6e, 0x67, 0x69, 0x6e, 0x65, 0x18, 0x91, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x61,
	0x72, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x72, 0x5f, 0x79,
	0x65, 0x61, 0x72, 0x18, 0x92, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x63, 0x61, 0x72, 0x59,
	0x65, 0x61, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x61, 0x72, 0x5f, 0x6d, 0x69, 0x6c, 0x61, 0x67,
	0x65, 0x18, 0x93, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x63, 0x61, 0x72, 0x4d, 0x69, 0x6c,
	0x61, 0x67, 0x65, 0x12, 0x2a, 0x0a, 0x10, 0x63, 0x61, 0x72, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x94, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f,
	0x63, 0x61, 0x72, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x2b, 0x0a, 0x11, 0x65, 0x78, 0x70, 0x5f, 0x73, 0x65, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x18, 0xf5, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0f, 0x65, 0x78, 0x70,
	0x53, 0x65, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x17, 0x0a, 0x06,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0xd9, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x1b, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0xbd, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x2a, 0x0a, 0x10, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0xbe, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1c,
	0x0a, 0x09, 0x69, 0x73, 0x5f, 0x64, 0x65, 0x61, 0x6c, 0x65, 0x72, 0x18, 0xbf, 0x05, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x08, 0x69, 0x73, 0x44, 0x65, 0x61, 0x6c, 0x65, 0x72, 0x12, 0x11, 0x0a, 0x03,
	0x73, 0x65, 0x71, 0x18, 0xc0, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x73, 0x65, 0x71, 0x12,
	0x1f, 0x0a, 0x0b, 0x63, 0x61, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x69, 0x64, 0x18, 0xc3,
	0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x63, 0x61, 0x72, 0x54, 0x79, 0x70, 0x65, 0x49, 0x64,
	0x12, 0x1c, 0x0a, 0x09, 0x64, 0x65, 0x61, 0x6c, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0xc4, 0x05,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x64, 0x65, 0x61, 0x6c, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1c,
	0x0a, 0x09, 0x66, 0x75, 0x65, 0x6c, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0xc5, 0x05, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x08, 0x66, 0x75, 0x65, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x73, 0x65, 0x61, 0x74, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0xc6, 0x05, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x08, 0x73, 0x65, 0x61, 0x74, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x15, 0x0a, 0x05, 0x63, 0x6f,
	0x6c, 0x6f, 0x72, 0x18, 0xc7, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x63, 0x6f, 0x6c, 0x6f,
	0x72, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x61, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0xc8, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x61, 0x72, 0x54, 0x79,
	0x70, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0xc1, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0xc2, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x62, 0x72, 0x61,
	0x6e, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0xcc, 0x06, 0x0a, 0x10, 0x43, 0x61, 0x72, 0x4c, 0x69,
	0x73, 0x74, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x33, 0x0a, 0x16, 0x65,
	0x78, 0x70, 0x5f, 0x73, 0x65, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x13, 0x65, 0x78, 0x70,
	0x53, 0x65, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x69, 0x63, 0x65, 0x46, 0x72, 0x6f, 0x6d,
	0x12, 0x2f, 0x0a, 0x14, 0x65, 0x78, 0x70, 0x5f, 0x73, 0x65, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x5f,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x5f, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x11,
	0x65, 0x78, 0x70, 0x53, 0x65, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x69, 0x63, 0x65, 0x54,
	0x6f, 0x12, 0x19, 0x0a, 0x08, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x04, 0x52, 0x07, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x09, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x62,
	0x72, 0x61, 0x6e, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x03, 0x28, 0x04, 0x52, 0x07, 0x62,
	0x72, 0x61, 0x6e, 0x64, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x62, 0x72, 0x61, 0x6e,
	0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74,
	0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x76, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x79, 0x65, 0x61, 0x72, 0x5f,
	0x66, 0x72, 0x6f, 0x6d, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x79, 0x65, 0x61, 0x72,
	0x46, 0x72, 0x6f, 0x6d, 0x12, 0x17, 0x0a, 0x07, 0x79, 0x65, 0x61, 0x72, 0x5f, 0x74, 0x6f, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x79, 0x65, 0x61, 0x72, 0x54, 0x6f, 0x12, 0x1f, 0x0a,
	0x0b, 0x6d, 0x69, 0x6c, 0x61, 0x67, 0x65, 0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x0a, 0x6d, 0x69, 0x6c, 0x61, 0x67, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x12, 0x1b,
	0x0a, 0x09, 0x6d, 0x69, 0x6c, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x08, 0x6d, 0x69, 0x6c, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x12, 0x22, 0x0a, 0x0c, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x0d, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x16, 0x0a, 0x06, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x06, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72,
	0x18, 0x0f, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x12, 0x1b, 0x0a,
	0x09, 0x66, 0x75, 0x65, 0x6c, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x10, 0x20, 0x03, 0x28, 0x0d,
	0x52, 0x08, 0x66, 0x75, 0x65, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65,
	0x61, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x11, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x08, 0x73,
	0x65, 0x61, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x72, 0x6b, 0x69,
	0x6e, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x12, 0x20, 0x03, 0x28, 0x04, 0x52, 0x09, 0x70, 0x61, 0x72,
	0x6b, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x14, 0x20, 0x03, 0x28, 0x04, 0x52, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x49,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x16, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x17, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x69, 0x64, 0x18, 0x1b, 0x20, 0x03, 0x28, 0x04, 0x52, 0x0a, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x65, 0x61, 0x6c, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x18, 0x20, 0x03, 0x28, 0x04, 0x52, 0x08, 0x64, 0x65, 0x61, 0x6c, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0b, 0x63, 0x61, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x1e, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x09, 0x63, 0x61, 0x72, 0x54, 0x79, 0x70,
	0x65, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x6f, 0x72, 0x74, 0x5f, 0x6b, 0x65, 0x79, 0x18,
	0x19, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x6f, 0x72, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x1b,
	0x0a, 0x09, 0x73, 0x6f, 0x72, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x1a, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x08, 0x73, 0x6f, 0x72, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70,
	0x61, 0x67, 0x65, 0x18, 0x1c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x1d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x1e, 0x0a, 0x0c, 0x43, 0x61, 0x72, 0x4c, 0x69, 0x73, 0x74,
	0x49, 0x64, 0x44, 0x74, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x04, 0x52, 0x02, 0x69, 0x64, 0x32, 0xa9, 0x02, 0x0a, 0x07, 0x43, 0x61, 0x72, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x2a, 0x0a, 0x03, 0x41, 0x64, 0x64, 0x12, 0x0f, 0x2e, 0x63, 0x61, 0x72, 0x2e, 0x43,
	0x61, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x74, 0x6f, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2f, 0x0a,
	0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x11, 0x2e, 0x63, 0x61, 0x72, 0x2e, 0x43, 0x61,
	0x72, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x64, 0x44, 0x74, 0x6f, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2d,
	0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x0f, 0x2e, 0x63, 0x61, 0x72, 0x2e, 0x43,
	0x61, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x74, 0x6f, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2c, 0x0a,
	0x03, 0x47, 0x65, 0x74, 0x12, 0x11, 0x2e, 0x63, 0x61, 0x72, 0x2e, 0x43, 0x61, 0x72, 0x4c, 0x69,
	0x73, 0x74, 0x49, 0x64, 0x44, 0x74, 0x6f, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x28, 0x0a, 0x04, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67,
	0x65, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6e,
	0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x15, 0x2e, 0x63, 0x61, 0x72, 0x2e, 0x43, 0x61, 0x72,
	0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x10, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x63, 0x61, 0x72, 0x67, 0x6f, 0x64, 0x2d, 0x62, 0x6a, 0x2f, 0x62, 0x32, 0x63, 0x2d, 0x63, 0x61,
	0x72, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x61, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_carList_carList_proto_rawDescOnce sync.Once
	file_carList_carList_proto_rawDescData = file_carList_carList_proto_rawDesc
)

func file_carList_carList_proto_rawDescGZIP() []byte {
	file_carList_carList_proto_rawDescOnce.Do(func() {
		file_carList_carList_proto_rawDescData = protoimpl.X.CompressGZIP(file_carList_carList_proto_rawDescData)
	})
	return file_carList_carList_proto_rawDescData
}

var file_carList_carList_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_carList_carList_proto_goTypes = []interface{}{
	(*CarListDto)(nil),       // 0: car.CarListDto
	(*CarListCondition)(nil), // 1: car.CarListCondition
	(*CarListIdDto)(nil),     // 2: car.CarListIdDto
	(*common.Page)(nil),      // 3: common.Page
	(*common.Response)(nil),  // 4: common.Response
}
var file_carList_carList_proto_depIdxs = []int32{
	0, // 0: car.CarList.Add:input_type -> car.CarListDto
	2, // 1: car.CarList.Delete:input_type -> car.CarListIdDto
	0, // 2: car.CarList.Update:input_type -> car.CarListDto
	2, // 3: car.CarList.Get:input_type -> car.CarListIdDto
	3, // 4: car.CarList.List:input_type -> common.Page
	1, // 5: car.CarList.ListCondition:input_type -> car.CarListCondition
	4, // 6: car.CarList.Add:output_type -> common.Response
	4, // 7: car.CarList.Delete:output_type -> common.Response
	4, // 8: car.CarList.Update:output_type -> common.Response
	4, // 9: car.CarList.Get:output_type -> common.Response
	4, // 10: car.CarList.List:output_type -> common.Response
	4, // 11: car.CarList.ListCondition:output_type -> common.Response
	6, // [6:12] is the sub-list for method output_type
	0, // [0:6] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_carList_carList_proto_init() }
func file_carList_carList_proto_init() {
	if File_carList_carList_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_carList_carList_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CarListDto); i {
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
		file_carList_carList_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CarListCondition); i {
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
		file_carList_carList_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CarListIdDto); i {
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
			RawDescriptor: file_carList_carList_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_carList_carList_proto_goTypes,
		DependencyIndexes: file_carList_carList_proto_depIdxs,
		MessageInfos:      file_carList_carList_proto_msgTypes,
	}.Build()
	File_carList_carList_proto = out.File
	file_carList_carList_proto_rawDesc = nil
	file_carList_carList_proto_goTypes = nil
	file_carList_carList_proto_depIdxs = nil
}
