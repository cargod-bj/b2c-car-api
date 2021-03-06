// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: carTag/carTag.proto

package carTag

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

// Api Endpoints for CarTag service

func NewCarTagEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for CarTag service

type CarTagService interface {
	// 添加Tag，返回 data：nil
	Add(ctx context.Context, in *CarTagDto, opts ...client.CallOption) (*common.Response, error)
	// 根据Tagid删除车辆，返回 data：nil
	Delete(ctx context.Context, in *CarTagConditionDto, opts ...client.CallOption) (*common.Response, error)
	// 更新Tag，返回 data：nil
	Update(ctx context.Context, in *CarTagDto, opts ...client.CallOption) (*common.Response, error)
	//获取Tag信息: 返回data：common.PagedList
	List(ctx context.Context, in *CarTagConditionDto, opts ...client.CallOption) (*common.Response, error)
	//获取Tag类型
	TagType(ctx context.Context, in *CarTagType, opts ...client.CallOption) (*common.Response, error)
}

type carTagService struct {
	c    client.Client
	name string
}

func NewCarTagService(name string, c client.Client) CarTagService {
	return &carTagService{
		c:    c,
		name: name,
	}
}

func (c *carTagService) Add(ctx context.Context, in *CarTagDto, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "CarTag.Add", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carTagService) Delete(ctx context.Context, in *CarTagConditionDto, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "CarTag.Delete", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carTagService) Update(ctx context.Context, in *CarTagDto, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "CarTag.Update", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carTagService) List(ctx context.Context, in *CarTagConditionDto, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "CarTag.List", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carTagService) TagType(ctx context.Context, in *CarTagType, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "CarTag.TagType", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CarTag service

type CarTagHandler interface {
	// 添加Tag，返回 data：nil
	Add(context.Context, *CarTagDto, *common.Response) error
	// 根据Tagid删除车辆，返回 data：nil
	Delete(context.Context, *CarTagConditionDto, *common.Response) error
	// 更新Tag，返回 data：nil
	Update(context.Context, *CarTagDto, *common.Response) error
	//获取Tag信息: 返回data：common.PagedList
	List(context.Context, *CarTagConditionDto, *common.Response) error
	//获取Tag类型
	TagType(context.Context, *CarTagType, *common.Response) error
}

func RegisterCarTagHandler(s server.Server, hdlr CarTagHandler, opts ...server.HandlerOption) error {
	type carTag interface {
		Add(ctx context.Context, in *CarTagDto, out *common.Response) error
		Delete(ctx context.Context, in *CarTagConditionDto, out *common.Response) error
		Update(ctx context.Context, in *CarTagDto, out *common.Response) error
		List(ctx context.Context, in *CarTagConditionDto, out *common.Response) error
		TagType(ctx context.Context, in *CarTagType, out *common.Response) error
	}
	type CarTag struct {
		carTag
	}
	h := &carTagHandler{hdlr}
	return s.Handle(s.NewHandler(&CarTag{h}, opts...))
}

type carTagHandler struct {
	CarTagHandler
}

func (h *carTagHandler) Add(ctx context.Context, in *CarTagDto, out *common.Response) error {
	return h.CarTagHandler.Add(ctx, in, out)
}

func (h *carTagHandler) Delete(ctx context.Context, in *CarTagConditionDto, out *common.Response) error {
	return h.CarTagHandler.Delete(ctx, in, out)
}

func (h *carTagHandler) Update(ctx context.Context, in *CarTagDto, out *common.Response) error {
	return h.CarTagHandler.Update(ctx, in, out)
}

func (h *carTagHandler) List(ctx context.Context, in *CarTagConditionDto, out *common.Response) error {
	return h.CarTagHandler.List(ctx, in, out)
}

func (h *carTagHandler) TagType(ctx context.Context, in *CarTagType, out *common.Response) error {
	return h.CarTagHandler.TagType(ctx, in, out)
}
