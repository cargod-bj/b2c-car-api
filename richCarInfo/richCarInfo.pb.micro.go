// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: richCarInfo/richCarInfo.proto

package richCarInfo

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

// Api Endpoints for RichCarInfo service

func NewRichCarInfoEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for RichCarInfo service

type RichCarInfoService interface {
	Create(ctx context.Context, in *CreateReq, opts ...client.CallOption) (*common.Response, error)
	List(ctx context.Context, in *ListReq, opts ...client.CallOption) (*common.Response, error)
	Upadte(ctx context.Context, in *UpdateReq, opts ...client.CallOption) (*common.Response, error)
	UpadteIcon(ctx context.Context, in *UpadteIconReq, opts ...client.CallOption) (*common.Response, error)
	Delete(ctx context.Context, in *DeleteReq, opts ...client.CallOption) (*common.Response, error)
	Detial(ctx context.Context, in *common.EmptyDto, opts ...client.CallOption) (*common.Response, error)
	UpadteCar(ctx context.Context, in *UpadteCarReq, opts ...client.CallOption) (*common.Response, error)
}

type richCarInfoService struct {
	c    client.Client
	name string
}

func NewRichCarInfoService(name string, c client.Client) RichCarInfoService {
	return &richCarInfoService{
		c:    c,
		name: name,
	}
}

func (c *richCarInfoService) Create(ctx context.Context, in *CreateReq, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "RichCarInfo.Create", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *richCarInfoService) List(ctx context.Context, in *ListReq, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "RichCarInfo.List", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *richCarInfoService) Upadte(ctx context.Context, in *UpdateReq, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "RichCarInfo.Upadte", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *richCarInfoService) UpadteIcon(ctx context.Context, in *UpadteIconReq, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "RichCarInfo.UpadteIcon", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *richCarInfoService) Delete(ctx context.Context, in *DeleteReq, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "RichCarInfo.Delete", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *richCarInfoService) Detial(ctx context.Context, in *common.EmptyDto, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "RichCarInfo.Detial", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *richCarInfoService) UpadteCar(ctx context.Context, in *UpadteCarReq, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "RichCarInfo.UpadteCar", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RichCarInfo service

type RichCarInfoHandler interface {
	Create(context.Context, *CreateReq, *common.Response) error
	List(context.Context, *ListReq, *common.Response) error
	Upadte(context.Context, *UpdateReq, *common.Response) error
	UpadteIcon(context.Context, *UpadteIconReq, *common.Response) error
	Delete(context.Context, *DeleteReq, *common.Response) error
	Detial(context.Context, *common.EmptyDto, *common.Response) error
	UpadteCar(context.Context, *UpadteCarReq, *common.Response) error
}

func RegisterRichCarInfoHandler(s server.Server, hdlr RichCarInfoHandler, opts ...server.HandlerOption) error {
	type richCarInfo interface {
		Create(ctx context.Context, in *CreateReq, out *common.Response) error
		List(ctx context.Context, in *ListReq, out *common.Response) error
		Upadte(ctx context.Context, in *UpdateReq, out *common.Response) error
		UpadteIcon(ctx context.Context, in *UpadteIconReq, out *common.Response) error
		Delete(ctx context.Context, in *DeleteReq, out *common.Response) error
		Detial(ctx context.Context, in *common.EmptyDto, out *common.Response) error
		UpadteCar(ctx context.Context, in *UpadteCarReq, out *common.Response) error
	}
	type RichCarInfo struct {
		richCarInfo
	}
	h := &richCarInfoHandler{hdlr}
	return s.Handle(s.NewHandler(&RichCarInfo{h}, opts...))
}

type richCarInfoHandler struct {
	RichCarInfoHandler
}

func (h *richCarInfoHandler) Create(ctx context.Context, in *CreateReq, out *common.Response) error {
	return h.RichCarInfoHandler.Create(ctx, in, out)
}

func (h *richCarInfoHandler) List(ctx context.Context, in *ListReq, out *common.Response) error {
	return h.RichCarInfoHandler.List(ctx, in, out)
}

func (h *richCarInfoHandler) Upadte(ctx context.Context, in *UpdateReq, out *common.Response) error {
	return h.RichCarInfoHandler.Upadte(ctx, in, out)
}

func (h *richCarInfoHandler) UpadteIcon(ctx context.Context, in *UpadteIconReq, out *common.Response) error {
	return h.RichCarInfoHandler.UpadteIcon(ctx, in, out)
}

func (h *richCarInfoHandler) Delete(ctx context.Context, in *DeleteReq, out *common.Response) error {
	return h.RichCarInfoHandler.Delete(ctx, in, out)
}

func (h *richCarInfoHandler) Detial(ctx context.Context, in *common.EmptyDto, out *common.Response) error {
	return h.RichCarInfoHandler.Detial(ctx, in, out)
}

func (h *richCarInfoHandler) UpadteCar(ctx context.Context, in *UpadteCarReq, out *common.Response) error {
	return h.RichCarInfoHandler.UpadteCar(ctx, in, out)
}
