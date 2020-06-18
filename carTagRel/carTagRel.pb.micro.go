// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: carTagRel/carTagRel.proto

package carTagRel

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

// Api Endpoints for CarTagRel service

func NewCarTagRelEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for CarTagRel service

type CarTagRelService interface {
	// 添加Tagrel，返回 data：nil
	Add(ctx context.Context, in *CarTagRelDto, opts ...client.CallOption) (*common.Response, error)
	// 根据id删除车辆，返回 data：nil
	Delete(ctx context.Context, in *CarTagRelDto, opts ...client.CallOption) (*common.Response, error)
	// 更新Tagrel，返回 data：nil
	Update(ctx context.Context, in *CarTagRelDto, opts ...client.CallOption) (*common.Response, error)
	//获取Tagrel信息: 返回data：common.PagedList
	List(ctx context.Context, in *common.Page, opts ...client.CallOption) (*common.Response, error)
}

type carTagRelService struct {
	c    client.Client
	name string
}

func NewCarTagRelService(name string, c client.Client) CarTagRelService {
	return &carTagRelService{
		c:    c,
		name: name,
	}
}

func (c *carTagRelService) Add(ctx context.Context, in *CarTagRelDto, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "CarTagRel.Add", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carTagRelService) Delete(ctx context.Context, in *CarTagRelDto, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "CarTagRel.Delete", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carTagRelService) Update(ctx context.Context, in *CarTagRelDto, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "CarTagRel.Update", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carTagRelService) List(ctx context.Context, in *common.Page, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "CarTagRel.List", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CarTagRel service

type CarTagRelHandler interface {
	// 添加Tagrel，返回 data：nil
	Add(context.Context, *CarTagRelDto, *common.Response) error
	// 根据id删除车辆，返回 data：nil
	Delete(context.Context, *CarTagRelDto, *common.Response) error
	// 更新Tagrel，返回 data：nil
	Update(context.Context, *CarTagRelDto, *common.Response) error
	//获取Tagrel信息: 返回data：common.PagedList
	List(context.Context, *common.Page, *common.Response) error
}

func RegisterCarTagRelHandler(s server.Server, hdlr CarTagRelHandler, opts ...server.HandlerOption) error {
	type carTagRel interface {
		Add(ctx context.Context, in *CarTagRelDto, out *common.Response) error
		Delete(ctx context.Context, in *CarTagRelDto, out *common.Response) error
		Update(ctx context.Context, in *CarTagRelDto, out *common.Response) error
		List(ctx context.Context, in *common.Page, out *common.Response) error
	}
	type CarTagRel struct {
		carTagRel
	}
	h := &carTagRelHandler{hdlr}
	return s.Handle(s.NewHandler(&CarTagRel{h}, opts...))
}

type carTagRelHandler struct {
	CarTagRelHandler
}

func (h *carTagRelHandler) Add(ctx context.Context, in *CarTagRelDto, out *common.Response) error {
	return h.CarTagRelHandler.Add(ctx, in, out)
}

func (h *carTagRelHandler) Delete(ctx context.Context, in *CarTagRelDto, out *common.Response) error {
	return h.CarTagRelHandler.Delete(ctx, in, out)
}

func (h *carTagRelHandler) Update(ctx context.Context, in *CarTagRelDto, out *common.Response) error {
	return h.CarTagRelHandler.Update(ctx, in, out)
}

func (h *carTagRelHandler) List(ctx context.Context, in *common.Page, out *common.Response) error {
	return h.CarTagRelHandler.List(ctx, in, out)
}
