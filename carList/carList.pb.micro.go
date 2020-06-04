// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: carList/carList.proto

package carList

import (
	fmt "fmt"
	common "github.com/cargod-bj/b2c-proto-common/common"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/timestamp"
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

// Api Endpoints for CarList service

func NewCarListEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for CarList service

type CarListService interface {
	// 添加指定车辆，返回 data：nil
	Add(ctx context.Context, in *CarListDto, opts ...client.CallOption) (*common.Response, error)
	// 根据车辆id删除车辆，返回 data：nil
	Delete(ctx context.Context, in *CarListIdDto, opts ...client.CallOption) (*common.Response, error)
	// 更新指定车辆，返回 data：nil
	Update(ctx context.Context, in *CarListDto, opts ...client.CallOption) (*common.Response, error)
	// 获取指定id的车辆：返回 data: CarListDto
	Get(ctx context.Context, in *CarListIdDto, opts ...client.CallOption) (*common.Response, error)
	//获取车辆列表信息: 返回data：common.PagedList
	List(ctx context.Context, in *common.Page, opts ...client.CallOption) (*common.Response, error)
}

type carListService struct {
	c    client.Client
	name string
}

func NewCarListService(name string, c client.Client) CarListService {
	return &carListService{
		c:    c,
		name: name,
	}
}

func (c *carListService) Add(ctx context.Context, in *CarListDto, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "CarList.Add", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carListService) Delete(ctx context.Context, in *CarListIdDto, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "CarList.Delete", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carListService) Update(ctx context.Context, in *CarListDto, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "CarList.Update", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carListService) Get(ctx context.Context, in *CarListIdDto, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "CarList.Get", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carListService) List(ctx context.Context, in *common.Page, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "CarList.List", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CarList service

type CarListHandler interface {
	// 添加指定车辆，返回 data：nil
	Add(context.Context, *CarListDto, *common.Response) error
	// 根据车辆id删除车辆，返回 data：nil
	Delete(context.Context, *CarListIdDto, *common.Response) error
	// 更新指定车辆，返回 data：nil
	Update(context.Context, *CarListDto, *common.Response) error
	// 获取指定id的车辆：返回 data: CarListDto
	Get(context.Context, *CarListIdDto, *common.Response) error
	//获取车辆列表信息: 返回data：common.PagedList
	List(context.Context, *common.Page, *common.Response) error
}

func RegisterCarListHandler(s server.Server, hdlr CarListHandler, opts ...server.HandlerOption) error {
	type carList interface {
		Add(ctx context.Context, in *CarListDto, out *common.Response) error
		Delete(ctx context.Context, in *CarListIdDto, out *common.Response) error
		Update(ctx context.Context, in *CarListDto, out *common.Response) error
		Get(ctx context.Context, in *CarListIdDto, out *common.Response) error
		List(ctx context.Context, in *common.Page, out *common.Response) error
	}
	type CarList struct {
		carList
	}
	h := &carListHandler{hdlr}
	return s.Handle(s.NewHandler(&CarList{h}, opts...))
}

type carListHandler struct {
	CarListHandler
}

func (h *carListHandler) Add(ctx context.Context, in *CarListDto, out *common.Response) error {
	return h.CarListHandler.Add(ctx, in, out)
}

func (h *carListHandler) Delete(ctx context.Context, in *CarListIdDto, out *common.Response) error {
	return h.CarListHandler.Delete(ctx, in, out)
}

func (h *carListHandler) Update(ctx context.Context, in *CarListDto, out *common.Response) error {
	return h.CarListHandler.Update(ctx, in, out)
}

func (h *carListHandler) Get(ctx context.Context, in *CarListIdDto, out *common.Response) error {
	return h.CarListHandler.Get(ctx, in, out)
}

func (h *carListHandler) List(ctx context.Context, in *common.Page, out *common.Response) error {
	return h.CarListHandler.List(ctx, in, out)
}
