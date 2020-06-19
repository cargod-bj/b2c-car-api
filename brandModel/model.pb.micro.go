// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: brandModel/model.proto

package brandModelProto

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

// Api Endpoints for Model service

func NewModelEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Model service

type ModelService interface {
	// 新增车辆品牌
	// response.Data = nil
	Add(ctx context.Context, in *ModelDto, opts ...client.CallOption) (*common.Response, error)
	// 删除指定id的车辆品牌
	// response.Data = nil
	Delete(ctx context.Context, in *common.IdDto, opts ...client.CallOption) (*common.Response, error)
	// 更新指定车辆品牌，其中有部分不能更新字段，请参照 BrandDto 字段说明
	// response.Data = nil
	Update(ctx context.Context, in *ModelDto, opts ...client.CallOption) (*common.Response, error)
	// 获取指定id的车辆品牌
	// response.Data = BrandDto
	Get(ctx context.Context, in *common.IdDto, opts ...client.CallOption) (*common.Response, error)
	ListByBrandId(ctx context.Context, in *common.IdDto, opts ...client.CallOption) (*common.Response, error)
}

type modelService struct {
	c    client.Client
	name string
}

func NewModelService(name string, c client.Client) ModelService {
	return &modelService{
		c:    c,
		name: name,
	}
}

func (c *modelService) Add(ctx context.Context, in *ModelDto, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "Model.Add", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelService) Delete(ctx context.Context, in *common.IdDto, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "Model.Delete", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelService) Update(ctx context.Context, in *ModelDto, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "Model.Update", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelService) Get(ctx context.Context, in *common.IdDto, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "Model.Get", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelService) ListByBrandId(ctx context.Context, in *common.IdDto, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "Model.ListByBrandId", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Model service

type ModelHandler interface {
	// 新增车辆品牌
	// response.Data = nil
	Add(context.Context, *ModelDto, *common.Response) error
	// 删除指定id的车辆品牌
	// response.Data = nil
	Delete(context.Context, *common.IdDto, *common.Response) error
	// 更新指定车辆品牌，其中有部分不能更新字段，请参照 BrandDto 字段说明
	// response.Data = nil
	Update(context.Context, *ModelDto, *common.Response) error
	// 获取指定id的车辆品牌
	// response.Data = BrandDto
	Get(context.Context, *common.IdDto, *common.Response) error
	ListByBrandId(context.Context, *common.IdDto, *common.Response) error
}

func RegisterModelHandler(s server.Server, hdlr ModelHandler, opts ...server.HandlerOption) error {
	type model interface {
		Add(ctx context.Context, in *ModelDto, out *common.Response) error
		Delete(ctx context.Context, in *common.IdDto, out *common.Response) error
		Update(ctx context.Context, in *ModelDto, out *common.Response) error
		Get(ctx context.Context, in *common.IdDto, out *common.Response) error
		ListByBrandId(ctx context.Context, in *common.IdDto, out *common.Response) error
	}
	type Model struct {
		model
	}
	h := &modelHandler{hdlr}
	return s.Handle(s.NewHandler(&Model{h}, opts...))
}

type modelHandler struct {
	ModelHandler
}

func (h *modelHandler) Add(ctx context.Context, in *ModelDto, out *common.Response) error {
	return h.ModelHandler.Add(ctx, in, out)
}

func (h *modelHandler) Delete(ctx context.Context, in *common.IdDto, out *common.Response) error {
	return h.ModelHandler.Delete(ctx, in, out)
}

func (h *modelHandler) Update(ctx context.Context, in *ModelDto, out *common.Response) error {
	return h.ModelHandler.Update(ctx, in, out)
}

func (h *modelHandler) Get(ctx context.Context, in *common.IdDto, out *common.Response) error {
	return h.ModelHandler.Get(ctx, in, out)
}

func (h *modelHandler) ListByBrandId(ctx context.Context, in *common.IdDto, out *common.Response) error {
	return h.ModelHandler.ListByBrandId(ctx, in, out)
}
