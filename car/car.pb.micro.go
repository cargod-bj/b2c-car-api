// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: car/car.proto

package carProto

import (
	fmt "fmt"
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

// Api Endpoints for Car service

func NewCarEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Car service

type CarService interface {
	// 添加指定车辆，返回 data：CarDto 类型
	Add(ctx context.Context, in *CarDto, opts ...client.CallOption) (*common.Response, error)
	// 根据车辆id删除车辆，返回 data：nil
	Delete(ctx context.Context, in *CarIdDto, opts ...client.CallOption) (*common.Response, error)
	// 更新指定车辆，返回 data：nil
	Update(ctx context.Context, in *CarDto, opts ...client.CallOption) (*common.Response, error)
	// 获取指定id的车辆：返回 data: CarDto
	Get(ctx context.Context, in *CarIdDto, opts ...client.CallOption) (*common.Response, error)
	//获取车辆列表: 返回data：common.PagedList：CarDto
	List(ctx context.Context, in *common.Page, opts ...client.CallOption) (*common.Response, error)
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

func (c *carService) Delete(ctx context.Context, in *CarIdDto, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "Car.Delete", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carService) Update(ctx context.Context, in *CarDto, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "Car.Update", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carService) Get(ctx context.Context, in *CarIdDto, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "Car.Get", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carService) List(ctx context.Context, in *common.Page, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "Car.List", in)
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
	Delete(context.Context, *CarIdDto, *common.Response) error
	// 更新指定车辆，返回 data：nil
	Update(context.Context, *CarDto, *common.Response) error
	// 获取指定id的车辆：返回 data: CarDto
	Get(context.Context, *CarIdDto, *common.Response) error
	//获取车辆列表: 返回data：common.PagedList：CarDto
	List(context.Context, *common.Page, *common.Response) error
}

func RegisterCarHandler(s server.Server, hdlr CarHandler, opts ...server.HandlerOption) error {
	type car interface {
		Add(ctx context.Context, in *CarDto, out *common.Response) error
		Delete(ctx context.Context, in *CarIdDto, out *common.Response) error
		Update(ctx context.Context, in *CarDto, out *common.Response) error
		Get(ctx context.Context, in *CarIdDto, out *common.Response) error
		List(ctx context.Context, in *common.Page, out *common.Response) error
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

func (h *carHandler) Delete(ctx context.Context, in *CarIdDto, out *common.Response) error {
	return h.CarHandler.Delete(ctx, in, out)
}

func (h *carHandler) Update(ctx context.Context, in *CarDto, out *common.Response) error {
	return h.CarHandler.Update(ctx, in, out)
}

func (h *carHandler) Get(ctx context.Context, in *CarIdDto, out *common.Response) error {
	return h.CarHandler.Get(ctx, in, out)
}

func (h *carHandler) List(ctx context.Context, in *common.Page, out *common.Response) error {
	return h.CarHandler.List(ctx, in, out)
}
