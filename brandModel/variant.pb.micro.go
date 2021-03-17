// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: brandModel/variant.proto

package brandModelProto

import (
	fmt "fmt"
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

// Api Endpoints for Variant service

func NewVariantEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Variant service

type VariantService interface {
	// 新增车辆品牌
	// response.Data = nil
	Add(ctx context.Context, in *VariantDto, opts ...client.CallOption) (*common.Response, error)
	// 删除指定id的车辆品牌
	// response.Data = nil
	Delete(ctx context.Context, in *common.IdDto, opts ...client.CallOption) (*common.Response, error)
	// 更新指定车辆品牌，其中有部分不能更新字段，请参照 BrandDto 字段说明
	// response.Data = nil
	Update(ctx context.Context, in *VariantDto, opts ...client.CallOption) (*common.Response, error)
	// 获取指定id的车辆品牌
	// response.Data = BrandDto
	Get(ctx context.Context, in *common.IdLocalDTO, opts ...client.CallOption) (*common.Response, error)
	ListByModelId(ctx context.Context, in *common.IdLocalDTO, opts ...client.CallOption) (*common.Response, error)
}

type variantService struct {
	c    client.Client
	name string
}

func NewVariantService(name string, c client.Client) VariantService {
	return &variantService{
		c:    c,
		name: name,
	}
}

func (c *variantService) Add(ctx context.Context, in *VariantDto, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "Variant.Add", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *variantService) Delete(ctx context.Context, in *common.IdDto, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "Variant.Delete", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *variantService) Update(ctx context.Context, in *VariantDto, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "Variant.Update", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *variantService) Get(ctx context.Context, in *common.IdLocalDTO, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "Variant.Get", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *variantService) ListByModelId(ctx context.Context, in *common.IdLocalDTO, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "Variant.ListByModelId", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Variant service

type VariantHandler interface {
	// 新增车辆品牌
	// response.Data = nil
	Add(context.Context, *VariantDto, *common.Response) error
	// 删除指定id的车辆品牌
	// response.Data = nil
	Delete(context.Context, *common.IdDto, *common.Response) error
	// 更新指定车辆品牌，其中有部分不能更新字段，请参照 BrandDto 字段说明
	// response.Data = nil
	Update(context.Context, *VariantDto, *common.Response) error
	// 获取指定id的车辆品牌
	// response.Data = BrandDto
	Get(context.Context, *common.IdLocalDTO, *common.Response) error
	ListByModelId(context.Context, *common.IdLocalDTO, *common.Response) error
}

func RegisterVariantHandler(s server.Server, hdlr VariantHandler, opts ...server.HandlerOption) error {
	type variant interface {
		Add(ctx context.Context, in *VariantDto, out *common.Response) error
		Delete(ctx context.Context, in *common.IdDto, out *common.Response) error
		Update(ctx context.Context, in *VariantDto, out *common.Response) error
		Get(ctx context.Context, in *common.IdLocalDTO, out *common.Response) error
		ListByModelId(ctx context.Context, in *common.IdLocalDTO, out *common.Response) error
	}
	type Variant struct {
		variant
	}
	h := &variantHandler{hdlr}
	return s.Handle(s.NewHandler(&Variant{h}, opts...))
}

type variantHandler struct {
	VariantHandler
}

func (h *variantHandler) Add(ctx context.Context, in *VariantDto, out *common.Response) error {
	return h.VariantHandler.Add(ctx, in, out)
}

func (h *variantHandler) Delete(ctx context.Context, in *common.IdDto, out *common.Response) error {
	return h.VariantHandler.Delete(ctx, in, out)
}

func (h *variantHandler) Update(ctx context.Context, in *VariantDto, out *common.Response) error {
	return h.VariantHandler.Update(ctx, in, out)
}

func (h *variantHandler) Get(ctx context.Context, in *common.IdLocalDTO, out *common.Response) error {
	return h.VariantHandler.Get(ctx, in, out)
}

func (h *variantHandler) ListByModelId(ctx context.Context, in *common.IdLocalDTO, out *common.Response) error {
	return h.VariantHandler.ListByModelId(ctx, in, out)
}
