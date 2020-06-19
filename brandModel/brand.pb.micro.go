// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: brandModel/brand.proto

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

// Api Endpoints for Brand service

func NewBrandEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Brand service

type BrandService interface {
	// 新增车辆品牌
	// response.Data = nil
	Add(ctx context.Context, in *BrandDto, opts ...client.CallOption) (*common.Response, error)
	// 删除指定id的车辆品牌
	// response.Data = nil
	Delete(ctx context.Context, in *common.IdDto, opts ...client.CallOption) (*common.Response, error)
	// 更新指定车辆品牌，其中有部分不能更新字段，请参照 BrandDto 字段说明
	// response.Data = nil
	Update(ctx context.Context, in *BrandDto, opts ...client.CallOption) (*common.Response, error)
	// 获取指定id的车辆品牌
	// response.Data = BrandDto
	Get(ctx context.Context, in *common.IdDto, opts ...client.CallOption) (*common.Response, error)
	List(ctx context.Context, in *common.Page, opts ...client.CallOption) (*common.Response, error)
}

type brandService struct {
	c    client.Client
	name string
}

func NewBrandService(name string, c client.Client) BrandService {
	return &brandService{
		c:    c,
		name: name,
	}
}

func (c *brandService) Add(ctx context.Context, in *BrandDto, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "Brand.Add", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *brandService) Delete(ctx context.Context, in *common.IdDto, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "Brand.Delete", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *brandService) Update(ctx context.Context, in *BrandDto, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "Brand.Update", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *brandService) Get(ctx context.Context, in *common.IdDto, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "Brand.Get", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *brandService) List(ctx context.Context, in *common.Page, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "Brand.List", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Brand service

type BrandHandler interface {
	// 新增车辆品牌
	// response.Data = nil
	Add(context.Context, *BrandDto, *common.Response) error
	// 删除指定id的车辆品牌
	// response.Data = nil
	Delete(context.Context, *common.IdDto, *common.Response) error
	// 更新指定车辆品牌，其中有部分不能更新字段，请参照 BrandDto 字段说明
	// response.Data = nil
	Update(context.Context, *BrandDto, *common.Response) error
	// 获取指定id的车辆品牌
	// response.Data = BrandDto
	Get(context.Context, *common.IdDto, *common.Response) error
	List(context.Context, *common.Page, *common.Response) error
}

func RegisterBrandHandler(s server.Server, hdlr BrandHandler, opts ...server.HandlerOption) error {
	type brand interface {
		Add(ctx context.Context, in *BrandDto, out *common.Response) error
		Delete(ctx context.Context, in *common.IdDto, out *common.Response) error
		Update(ctx context.Context, in *BrandDto, out *common.Response) error
		Get(ctx context.Context, in *common.IdDto, out *common.Response) error
		List(ctx context.Context, in *common.Page, out *common.Response) error
	}
	type Brand struct {
		brand
	}
	h := &brandHandler{hdlr}
	return s.Handle(s.NewHandler(&Brand{h}, opts...))
}

type brandHandler struct {
	BrandHandler
}

func (h *brandHandler) Add(ctx context.Context, in *BrandDto, out *common.Response) error {
	return h.BrandHandler.Add(ctx, in, out)
}

func (h *brandHandler) Delete(ctx context.Context, in *common.IdDto, out *common.Response) error {
	return h.BrandHandler.Delete(ctx, in, out)
}

func (h *brandHandler) Update(ctx context.Context, in *BrandDto, out *common.Response) error {
	return h.BrandHandler.Update(ctx, in, out)
}

func (h *brandHandler) Get(ctx context.Context, in *common.IdDto, out *common.Response) error {
	return h.BrandHandler.Get(ctx, in, out)
}

func (h *brandHandler) List(ctx context.Context, in *common.Page, out *common.Response) error {
	return h.BrandHandler.List(ctx, in, out)
}
