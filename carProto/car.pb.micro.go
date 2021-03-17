// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: carProto/car.proto

package carProto

import (
	fmt "fmt"
	_ "github.com/cargod-bj/b2c-car-api/reportProto"
	common "github.com/cargod-bj/b2c-proto-common/common"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/micro/v3/service/api"
	client "github.com/micro/micro/v3/service/client"
	server "github.com/micro/micro/v3/service/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Car service

func NewCarEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Car service

type CarService interface {
	// 添加指定车辆，返回 data：CarDto 类型
	Add(ctx context.Context, in *CarDto, opts ...client.CallOption) (*common.Response, error)
	// 根据车辆id删除车辆，返回 data：nil
	Delete(ctx context.Context, in *CarIdsDto, opts ...client.CallOption) (*common.Response, error)
	// 更新指定车辆，返回 data：nil
	Update(ctx context.Context, in *UpdateCarDto, opts ...client.CallOption) (*common.Response, error)
	// 获取指定id的车辆：返回 data: CarDto List
	Get(ctx context.Context, in *CarIdsDto, opts ...client.CallOption) (*common.Response, error)
	// 获取指定id的车辆,并且可以指定获取内容：返回 data: CarDto List
	GetSimpleInfo(ctx context.Context, in *CarSimpleInfoParams, opts ...client.CallOption) (*common.Response, error)
	// 给website专用的接口，获取指定id的车辆：返回 data: CarDto List
	GetForWebsite(ctx context.Context, in *CarIdsDto, opts ...client.CallOption) (*common.Response, error)
	//获取车辆列表: 返回data：common.PagedList：CarDto
	List(ctx context.Context, in *CarListParams, opts ...client.CallOption) (*common.Response, error)
	//获取车辆来源李彪
	SourceList(ctx context.Context, in *SourceParams, opts ...client.CallOption) (*common.Response, error)
	//添加车辆到平台
	AddFromSource(ctx context.Context, in *AddFromSourceParams, opts ...client.CallOption) (*common.Response, error)
	AddFromSourceList(ctx context.Context, in *AddFromSourceListParams, opts ...client.CallOption) (*common.Response, error)
	// 添加第三方车辆，返回 data：nil
	AddPartnerCar(ctx context.Context, in *AddPartnerCarReq, opts ...client.CallOption) (*common.Response, error)
	// 上架车辆
	LaunchCar(ctx context.Context, in *common.IdDto, opts ...client.CallOption) (*common.Response, error)
	// 查询车辆可能变更的状态列表：返回 Data = common.PageList{
	//              List = List<carEnumProto.KeyValueDto>
	//          }
	GetValidState(ctx context.Context, in *GetValidStateReq, opts ...client.CallOption) (*common.Response, error)
	// 更改车辆状态：返回 Data = nil
	ChangeCarState(ctx context.Context, in *ChangeCarStateReq, opts ...client.CallOption) (*common.Response, error)
	// 更改库存状态
	ChangeInventoryStatus(ctx context.Context, in *ChangeCarStateReq, opts ...client.CallOption) (*common.Response, error)
	// 获取指定carNo的车辆：返回 data: CarDto
	GetCarByCarNo(ctx context.Context, in *CarNoDto, opts ...client.CallOption) (*common.Response, error)
	// 获取车辆No模糊搜索车辆信息：返回 data: CarDtoList
	GetCarByNoFuzzy(ctx context.Context, in *CarNoDto, opts ...client.CallOption) (*common.Response, error)
	// 获取车辆详情页的访问token：返回 Data = CarDetailAccessDataResp
	GetCarDetailAccessData(ctx context.Context, in *CarNoReq, opts ...client.CallOption) (*common.Response, error)
	// 获取指定carNo或licensePlate的车辆：返回 data: CarDto
	GetCarByCarNoOrLicensePlate(ctx context.Context, in *CarNoOrLicensePlateDto, opts ...client.CallOption) (*common.Response, error)
	//  车辆变更门店
	TransferStore(ctx context.Context, in *CarTransferStoreDto, opts ...client.CallOption) (*common.Response, error)
	// 车辆变更门店记录
	TransferStoreList(ctx context.Context, in *CarTransferStoreList, opts ...client.CallOption) (*common.Response, error)
}

type carService struct {
	c    client.Client
	name string
}

func NewCarService(name string, c client.Client) CarService {
	return &carService{
		c:    c,
		name: name,
	}
}

func (c *carService) Add(ctx context.Context, in *CarDto, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "Car.Add", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carService) Delete(ctx context.Context, in *CarIdsDto, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "Car.Delete", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carService) Update(ctx context.Context, in *UpdateCarDto, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "Car.Update", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carService) Get(ctx context.Context, in *CarIdsDto, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "Car.Get", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carService) GetSimpleInfo(ctx context.Context, in *CarSimpleInfoParams, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "Car.GetSimpleInfo", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carService) GetForWebsite(ctx context.Context, in *CarIdsDto, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "Car.GetForWebsite", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carService) List(ctx context.Context, in *CarListParams, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "Car.List", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carService) SourceList(ctx context.Context, in *SourceParams, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "Car.SourceList", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carService) AddFromSource(ctx context.Context, in *AddFromSourceParams, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "Car.AddFromSource", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carService) AddFromSourceList(ctx context.Context, in *AddFromSourceListParams, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "Car.AddFromSourceList", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carService) AddPartnerCar(ctx context.Context, in *AddPartnerCarReq, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "Car.AddPartnerCar", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carService) LaunchCar(ctx context.Context, in *common.IdDto, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "Car.LaunchCar", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carService) GetValidState(ctx context.Context, in *GetValidStateReq, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "Car.GetValidState", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carService) ChangeCarState(ctx context.Context, in *ChangeCarStateReq, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "Car.ChangeCarState", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carService) ChangeInventoryStatus(ctx context.Context, in *ChangeCarStateReq, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "Car.ChangeInventoryStatus", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carService) GetCarByCarNo(ctx context.Context, in *CarNoDto, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "Car.GetCarByCarNo", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carService) GetCarByNoFuzzy(ctx context.Context, in *CarNoDto, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "Car.GetCarByNoFuzzy", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carService) GetCarDetailAccessData(ctx context.Context, in *CarNoReq, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "Car.GetCarDetailAccessData", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carService) GetCarByCarNoOrLicensePlate(ctx context.Context, in *CarNoOrLicensePlateDto, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "Car.GetCarByCarNoOrLicensePlate", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carService) TransferStore(ctx context.Context, in *CarTransferStoreDto, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "Car.TransferStore", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carService) TransferStoreList(ctx context.Context, in *CarTransferStoreList, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "Car.TransferStoreList", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Car service

type CarHandler interface {
	// 添加指定车辆，返回 data：CarDto 类型
	Add(context.Context, *CarDto, *common.Response) error
	// 根据车辆id删除车辆，返回 data：nil
	Delete(context.Context, *CarIdsDto, *common.Response) error
	// 更新指定车辆，返回 data：nil
	Update(context.Context, *UpdateCarDto, *common.Response) error
	// 获取指定id的车辆：返回 data: CarDto List
	Get(context.Context, *CarIdsDto, *common.Response) error
	// 获取指定id的车辆,并且可以指定获取内容：返回 data: CarDto List
	GetSimpleInfo(context.Context, *CarSimpleInfoParams, *common.Response) error
	// 给website专用的接口，获取指定id的车辆：返回 data: CarDto List
	GetForWebsite(context.Context, *CarIdsDto, *common.Response) error
	//获取车辆列表: 返回data：common.PagedList：CarDto
	List(context.Context, *CarListParams, *common.Response) error
	//获取车辆来源李彪
	SourceList(context.Context, *SourceParams, *common.Response) error
	//添加车辆到平台
	AddFromSource(context.Context, *AddFromSourceParams, *common.Response) error
	AddFromSourceList(context.Context, *AddFromSourceListParams, *common.Response) error
	// 添加第三方车辆，返回 data：nil
	AddPartnerCar(context.Context, *AddPartnerCarReq, *common.Response) error
	// 上架车辆
	LaunchCar(context.Context, *common.IdDto, *common.Response) error
	// 查询车辆可能变更的状态列表：返回 Data = common.PageList{
	//              List = List<carEnumProto.KeyValueDto>
	//          }
	GetValidState(context.Context, *GetValidStateReq, *common.Response) error
	// 更改车辆状态：返回 Data = nil
	ChangeCarState(context.Context, *ChangeCarStateReq, *common.Response) error
	// 更改库存状态
	ChangeInventoryStatus(context.Context, *ChangeCarStateReq, *common.Response) error
	// 获取指定carNo的车辆：返回 data: CarDto
	GetCarByCarNo(context.Context, *CarNoDto, *common.Response) error
	// 获取车辆No模糊搜索车辆信息：返回 data: CarDtoList
	GetCarByNoFuzzy(context.Context, *CarNoDto, *common.Response) error
	// 获取车辆详情页的访问token：返回 Data = CarDetailAccessDataResp
	GetCarDetailAccessData(context.Context, *CarNoReq, *common.Response) error
	// 获取指定carNo或licensePlate的车辆：返回 data: CarDto
	GetCarByCarNoOrLicensePlate(context.Context, *CarNoOrLicensePlateDto, *common.Response) error
	//  车辆变更门店
	TransferStore(context.Context, *CarTransferStoreDto, *common.Response) error
	// 车辆变更门店记录
	TransferStoreList(context.Context, *CarTransferStoreList, *common.Response) error
}

func RegisterCarHandler(s server.Server, hdlr CarHandler, opts ...server.HandlerOption) error {
	type car interface {
		Add(ctx context.Context, in *CarDto, out *common.Response) error
		Delete(ctx context.Context, in *CarIdsDto, out *common.Response) error
		Update(ctx context.Context, in *UpdateCarDto, out *common.Response) error
		Get(ctx context.Context, in *CarIdsDto, out *common.Response) error
		GetSimpleInfo(ctx context.Context, in *CarSimpleInfoParams, out *common.Response) error
		GetForWebsite(ctx context.Context, in *CarIdsDto, out *common.Response) error
		List(ctx context.Context, in *CarListParams, out *common.Response) error
		SourceList(ctx context.Context, in *SourceParams, out *common.Response) error
		AddFromSource(ctx context.Context, in *AddFromSourceParams, out *common.Response) error
		AddFromSourceList(ctx context.Context, in *AddFromSourceListParams, out *common.Response) error
		AddPartnerCar(ctx context.Context, in *AddPartnerCarReq, out *common.Response) error
		LaunchCar(ctx context.Context, in *common.IdDto, out *common.Response) error
		GetValidState(ctx context.Context, in *GetValidStateReq, out *common.Response) error
		ChangeCarState(ctx context.Context, in *ChangeCarStateReq, out *common.Response) error
		ChangeInventoryStatus(ctx context.Context, in *ChangeCarStateReq, out *common.Response) error
		GetCarByCarNo(ctx context.Context, in *CarNoDto, out *common.Response) error
		GetCarByNoFuzzy(ctx context.Context, in *CarNoDto, out *common.Response) error
		GetCarDetailAccessData(ctx context.Context, in *CarNoReq, out *common.Response) error
		GetCarByCarNoOrLicensePlate(ctx context.Context, in *CarNoOrLicensePlateDto, out *common.Response) error
		TransferStore(ctx context.Context, in *CarTransferStoreDto, out *common.Response) error
		TransferStoreList(ctx context.Context, in *CarTransferStoreList, out *common.Response) error
	}
	type Car struct {
		car
	}
	h := &carHandler{hdlr}
	return s.Handle(s.NewHandler(&Car{h}, opts...))
}

type carHandler struct {
	CarHandler
}

func (h *carHandler) Add(ctx context.Context, in *CarDto, out *common.Response) error {
	return h.CarHandler.Add(ctx, in, out)
}

func (h *carHandler) Delete(ctx context.Context, in *CarIdsDto, out *common.Response) error {
	return h.CarHandler.Delete(ctx, in, out)
}

func (h *carHandler) Update(ctx context.Context, in *UpdateCarDto, out *common.Response) error {
	return h.CarHandler.Update(ctx, in, out)
}

func (h *carHandler) Get(ctx context.Context, in *CarIdsDto, out *common.Response) error {
	return h.CarHandler.Get(ctx, in, out)
}

func (h *carHandler) GetSimpleInfo(ctx context.Context, in *CarSimpleInfoParams, out *common.Response) error {
	return h.CarHandler.GetSimpleInfo(ctx, in, out)
}

func (h *carHandler) GetForWebsite(ctx context.Context, in *CarIdsDto, out *common.Response) error {
	return h.CarHandler.GetForWebsite(ctx, in, out)
}

func (h *carHandler) List(ctx context.Context, in *CarListParams, out *common.Response) error {
	return h.CarHandler.List(ctx, in, out)
}

func (h *carHandler) SourceList(ctx context.Context, in *SourceParams, out *common.Response) error {
	return h.CarHandler.SourceList(ctx, in, out)
}

func (h *carHandler) AddFromSource(ctx context.Context, in *AddFromSourceParams, out *common.Response) error {
	return h.CarHandler.AddFromSource(ctx, in, out)
}

func (h *carHandler) AddFromSourceList(ctx context.Context, in *AddFromSourceListParams, out *common.Response) error {
	return h.CarHandler.AddFromSourceList(ctx, in, out)
}

func (h *carHandler) AddPartnerCar(ctx context.Context, in *AddPartnerCarReq, out *common.Response) error {
	return h.CarHandler.AddPartnerCar(ctx, in, out)
}

func (h *carHandler) LaunchCar(ctx context.Context, in *common.IdDto, out *common.Response) error {
	return h.CarHandler.LaunchCar(ctx, in, out)
}

func (h *carHandler) GetValidState(ctx context.Context, in *GetValidStateReq, out *common.Response) error {
	return h.CarHandler.GetValidState(ctx, in, out)
}

func (h *carHandler) ChangeCarState(ctx context.Context, in *ChangeCarStateReq, out *common.Response) error {
	return h.CarHandler.ChangeCarState(ctx, in, out)
}

func (h *carHandler) ChangeInventoryStatus(ctx context.Context, in *ChangeCarStateReq, out *common.Response) error {
	return h.CarHandler.ChangeInventoryStatus(ctx, in, out)
}

func (h *carHandler) GetCarByCarNo(ctx context.Context, in *CarNoDto, out *common.Response) error {
	return h.CarHandler.GetCarByCarNo(ctx, in, out)
}

func (h *carHandler) GetCarByNoFuzzy(ctx context.Context, in *CarNoDto, out *common.Response) error {
	return h.CarHandler.GetCarByNoFuzzy(ctx, in, out)
}

func (h *carHandler) GetCarDetailAccessData(ctx context.Context, in *CarNoReq, out *common.Response) error {
	return h.CarHandler.GetCarDetailAccessData(ctx, in, out)
}

func (h *carHandler) GetCarByCarNoOrLicensePlate(ctx context.Context, in *CarNoOrLicensePlateDto, out *common.Response) error {
	return h.CarHandler.GetCarByCarNoOrLicensePlate(ctx, in, out)
}

func (h *carHandler) TransferStore(ctx context.Context, in *CarTransferStoreDto, out *common.Response) error {
	return h.CarHandler.TransferStore(ctx, in, out)
}

func (h *carHandler) TransferStoreList(ctx context.Context, in *CarTransferStoreList, out *common.Response) error {
	return h.CarHandler.TransferStoreList(ctx, in, out)
}
