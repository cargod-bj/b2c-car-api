// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: carCampaignProto/carCampaign.proto

package carCampaignProto

import (
	fmt "fmt"
	_ "github.com/cargod-bj/b2c-car-api/carProto"
	common "github.com/cargod-bj/b2c-proto-common/common"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
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

// Api Endpoints for CarCampaign service

func NewCarCampaignEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for CarCampaign service

type CarCampaignService interface {
	// 添加活动，返回 data：nil
	Add(ctx context.Context, in *CarCampaignReq, opts ...client.CallOption) (*common.Response, error)
	// 停止指定活动，返回 data：nil
	Stop(ctx context.Context, in *DeleteCarCampaignReq, opts ...client.CallOption) (*common.Response, error)
	// 获取指定id的活动详情：返回 data: CarCampaignDetailWithCars
	Get(ctx context.Context, in *common.IdDto, opts ...client.CallOption) (*common.Response, error)
	// 查询活动列表：返回 Data = common.PageList{
	//              List = List<CarCampaignList>
	//          }
	List(ctx context.Context, in *ListCarCampaignReq, opts ...client.CallOption) (*common.Response, error)
	// 更新活动，返回 data：nil
	Update(ctx context.Context, in *CarCampaignReq, opts ...client.CallOption) (*common.Response, error)
	// 获取当前生效的限时活动，返回 data：CarCampaignDetail
	GetCurrentActiveCampaign(ctx context.Context, in *common.IdDto, opts ...client.CallOption) (*common.Response, error)
	// 获取指定ID的车辆活动信息
	GetCarCampaignInfo(ctx context.Context, in *CarCampaignCond, opts ...client.CallOption) (*common.Response, error)
	// 查询活动列表：返回 Data = common.PageList{
	//              List = List<CarCampaignLogList>
	//          }
	RemarkList(ctx context.Context, in *ListCarCampaignLogReq, opts ...client.CallOption) (*common.Response, error)
	// 检查导入的车辆信息：返回 Data = CheckImportCarsResp
	CheckImportCars(ctx context.Context, in *CheckImportCarsReq, opts ...client.CallOption) (*common.Response, error)
}

type carCampaignService struct {
	c    client.Client
	name string
}

func NewCarCampaignService(name string, c client.Client) CarCampaignService {
	return &carCampaignService{
		c:    c,
		name: name,
	}
}

func (c *carCampaignService) Add(ctx context.Context, in *CarCampaignReq, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "CarCampaign.Add", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carCampaignService) Stop(ctx context.Context, in *DeleteCarCampaignReq, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "CarCampaign.Stop", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carCampaignService) Get(ctx context.Context, in *common.IdDto, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "CarCampaign.Get", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carCampaignService) List(ctx context.Context, in *ListCarCampaignReq, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "CarCampaign.List", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carCampaignService) Update(ctx context.Context, in *CarCampaignReq, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "CarCampaign.Update", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carCampaignService) GetCurrentActiveCampaign(ctx context.Context, in *common.IdDto, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "CarCampaign.GetCurrentActiveCampaign", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carCampaignService) GetCarCampaignInfo(ctx context.Context, in *CarCampaignCond, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "CarCampaign.GetCarCampaignInfo", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carCampaignService) RemarkList(ctx context.Context, in *ListCarCampaignLogReq, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "CarCampaign.RemarkList", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carCampaignService) CheckImportCars(ctx context.Context, in *CheckImportCarsReq, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "CarCampaign.CheckImportCars", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CarCampaign service

type CarCampaignHandler interface {
	// 添加活动，返回 data：nil
	Add(context.Context, *CarCampaignReq, *common.Response) error
	// 停止指定活动，返回 data：nil
	Stop(context.Context, *DeleteCarCampaignReq, *common.Response) error
	// 获取指定id的活动详情：返回 data: CarCampaignDetailWithCars
	Get(context.Context, *common.IdDto, *common.Response) error
	// 查询活动列表：返回 Data = common.PageList{
	//              List = List<CarCampaignList>
	//          }
	List(context.Context, *ListCarCampaignReq, *common.Response) error
	// 更新活动，返回 data：nil
	Update(context.Context, *CarCampaignReq, *common.Response) error
	// 获取当前生效的限时活动，返回 data：CarCampaignDetail
	GetCurrentActiveCampaign(context.Context, *common.IdDto, *common.Response) error
	// 获取指定ID的车辆活动信息
	GetCarCampaignInfo(context.Context, *CarCampaignCond, *common.Response) error
	// 查询活动列表：返回 Data = common.PageList{
	//              List = List<CarCampaignLogList>
	//          }
	RemarkList(context.Context, *ListCarCampaignLogReq, *common.Response) error
	// 检查导入的车辆信息：返回 Data = CheckImportCarsResp
	CheckImportCars(context.Context, *CheckImportCarsReq, *common.Response) error
}

func RegisterCarCampaignHandler(s server.Server, hdlr CarCampaignHandler, opts ...server.HandlerOption) error {
	type carCampaign interface {
		Add(ctx context.Context, in *CarCampaignReq, out *common.Response) error
		Stop(ctx context.Context, in *DeleteCarCampaignReq, out *common.Response) error
		Get(ctx context.Context, in *common.IdDto, out *common.Response) error
		List(ctx context.Context, in *ListCarCampaignReq, out *common.Response) error
		Update(ctx context.Context, in *CarCampaignReq, out *common.Response) error
		GetCurrentActiveCampaign(ctx context.Context, in *common.IdDto, out *common.Response) error
		GetCarCampaignInfo(ctx context.Context, in *CarCampaignCond, out *common.Response) error
		RemarkList(ctx context.Context, in *ListCarCampaignLogReq, out *common.Response) error
		CheckImportCars(ctx context.Context, in *CheckImportCarsReq, out *common.Response) error
	}
	type CarCampaign struct {
		carCampaign
	}
	h := &carCampaignHandler{hdlr}
	return s.Handle(s.NewHandler(&CarCampaign{h}, opts...))
}

type carCampaignHandler struct {
	CarCampaignHandler
}

func (h *carCampaignHandler) Add(ctx context.Context, in *CarCampaignReq, out *common.Response) error {
	return h.CarCampaignHandler.Add(ctx, in, out)
}

func (h *carCampaignHandler) Stop(ctx context.Context, in *DeleteCarCampaignReq, out *common.Response) error {
	return h.CarCampaignHandler.Stop(ctx, in, out)
}

func (h *carCampaignHandler) Get(ctx context.Context, in *common.IdDto, out *common.Response) error {
	return h.CarCampaignHandler.Get(ctx, in, out)
}

func (h *carCampaignHandler) List(ctx context.Context, in *ListCarCampaignReq, out *common.Response) error {
	return h.CarCampaignHandler.List(ctx, in, out)
}

func (h *carCampaignHandler) Update(ctx context.Context, in *CarCampaignReq, out *common.Response) error {
	return h.CarCampaignHandler.Update(ctx, in, out)
}

func (h *carCampaignHandler) GetCurrentActiveCampaign(ctx context.Context, in *common.IdDto, out *common.Response) error {
	return h.CarCampaignHandler.GetCurrentActiveCampaign(ctx, in, out)
}

func (h *carCampaignHandler) GetCarCampaignInfo(ctx context.Context, in *CarCampaignCond, out *common.Response) error {
	return h.CarCampaignHandler.GetCarCampaignInfo(ctx, in, out)
}

func (h *carCampaignHandler) RemarkList(ctx context.Context, in *ListCarCampaignLogReq, out *common.Response) error {
	return h.CarCampaignHandler.RemarkList(ctx, in, out)
}

func (h *carCampaignHandler) CheckImportCars(ctx context.Context, in *CheckImportCarsReq, out *common.Response) error {
	return h.CarCampaignHandler.CheckImportCars(ctx, in, out)
}
