// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: carPurchase/car_purchase.proto

package carPurchaseProto

import (
	fmt "fmt"
	_ "github.com/cargod-bj/b2c-car-api/CarPurchaseCheckList"
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

// Api Endpoints for CarPurchase service

func NewCarPurchaseEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for CarPurchase service

type CarPurchaseService interface {
	// 用inspectionId查询车辆
	GetCarByInspection(ctx context.Context, in *InspectionReq, opts ...client.CallOption) (*common.Response, error)
	// 创建C2B来的采购单
	CreateC2B(ctx context.Context, in *CreateC2BReq, opts ...client.CallOption) (*common.Response, error)
	// 创建采购单
	Create(ctx context.Context, in *CreateReq, opts ...client.CallOption) (*common.Response, error)
	// 编辑采购单
	Update(ctx context.Context, in *UpdateReq, opts ...client.CallOption) (*common.Response, error)
	// 获取采购单列表
	List(ctx context.Context, in *ListReq, opts ...client.CallOption) (*common.Response, error)
	// 采购单详情
	Detail(ctx context.Context, in *common.IdDto, opts ...client.CallOption) (*common.Response, error)
	// 删除采购单
	Delete(ctx context.Context, in *common.IdDto, opts ...client.CallOption) (*common.Response, error)
	// 获取采购单日志
	LogList(ctx context.Context, in *LogReq, opts ...client.CallOption) (*common.Response, error)
}

type carPurchaseService struct {
	c    client.Client
	name string
}

func NewCarPurchaseService(name string, c client.Client) CarPurchaseService {
	return &carPurchaseService{
		c:    c,
		name: name,
	}
}

func (c *carPurchaseService) GetCarByInspection(ctx context.Context, in *InspectionReq, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "CarPurchase.GetCarByInspection", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carPurchaseService) CreateC2B(ctx context.Context, in *CreateC2BReq, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "CarPurchase.CreateC2B", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carPurchaseService) Create(ctx context.Context, in *CreateReq, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "CarPurchase.Create", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carPurchaseService) Update(ctx context.Context, in *UpdateReq, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "CarPurchase.Update", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carPurchaseService) List(ctx context.Context, in *ListReq, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "CarPurchase.List", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carPurchaseService) Detail(ctx context.Context, in *common.IdDto, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "CarPurchase.Detail", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carPurchaseService) Delete(ctx context.Context, in *common.IdDto, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "CarPurchase.Delete", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carPurchaseService) LogList(ctx context.Context, in *LogReq, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "CarPurchase.LogList", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CarPurchase service

type CarPurchaseHandler interface {
	// 用inspectionId查询车辆
	GetCarByInspection(context.Context, *InspectionReq, *common.Response) error
	// 创建C2B来的采购单
	CreateC2B(context.Context, *CreateC2BReq, *common.Response) error
	// 创建采购单
	Create(context.Context, *CreateReq, *common.Response) error
	// 编辑采购单
	Update(context.Context, *UpdateReq, *common.Response) error
	// 获取采购单列表
	List(context.Context, *ListReq, *common.Response) error
	// 采购单详情
	Detail(context.Context, *common.IdDto, *common.Response) error
	// 删除采购单
	Delete(context.Context, *common.IdDto, *common.Response) error
	// 获取采购单日志
	LogList(context.Context, *LogReq, *common.Response) error
}

func RegisterCarPurchaseHandler(s server.Server, hdlr CarPurchaseHandler, opts ...server.HandlerOption) error {
	type carPurchase interface {
		GetCarByInspection(ctx context.Context, in *InspectionReq, out *common.Response) error
		CreateC2B(ctx context.Context, in *CreateC2BReq, out *common.Response) error
		Create(ctx context.Context, in *CreateReq, out *common.Response) error
		Update(ctx context.Context, in *UpdateReq, out *common.Response) error
		List(ctx context.Context, in *ListReq, out *common.Response) error
		Detail(ctx context.Context, in *common.IdDto, out *common.Response) error
		Delete(ctx context.Context, in *common.IdDto, out *common.Response) error
		LogList(ctx context.Context, in *LogReq, out *common.Response) error
	}
	type CarPurchase struct {
		carPurchase
	}
	h := &carPurchaseHandler{hdlr}
	return s.Handle(s.NewHandler(&CarPurchase{h}, opts...))
}

type carPurchaseHandler struct {
	CarPurchaseHandler
}

func (h *carPurchaseHandler) GetCarByInspection(ctx context.Context, in *InspectionReq, out *common.Response) error {
	return h.CarPurchaseHandler.GetCarByInspection(ctx, in, out)
}

func (h *carPurchaseHandler) CreateC2B(ctx context.Context, in *CreateC2BReq, out *common.Response) error {
	return h.CarPurchaseHandler.CreateC2B(ctx, in, out)
}

func (h *carPurchaseHandler) Create(ctx context.Context, in *CreateReq, out *common.Response) error {
	return h.CarPurchaseHandler.Create(ctx, in, out)
}

func (h *carPurchaseHandler) Update(ctx context.Context, in *UpdateReq, out *common.Response) error {
	return h.CarPurchaseHandler.Update(ctx, in, out)
}

func (h *carPurchaseHandler) List(ctx context.Context, in *ListReq, out *common.Response) error {
	return h.CarPurchaseHandler.List(ctx, in, out)
}

func (h *carPurchaseHandler) Detail(ctx context.Context, in *common.IdDto, out *common.Response) error {
	return h.CarPurchaseHandler.Detail(ctx, in, out)
}

func (h *carPurchaseHandler) Delete(ctx context.Context, in *common.IdDto, out *common.Response) error {
	return h.CarPurchaseHandler.Delete(ctx, in, out)
}

func (h *carPurchaseHandler) LogList(ctx context.Context, in *LogReq, out *common.Response) error {
	return h.CarPurchaseHandler.LogList(ctx, in, out)
}
