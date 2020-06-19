// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: keywords/keywords.proto

package keywords

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

// Api Endpoints for Keywords service

func NewKeywordsEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Keywords service

type KeywordsService interface {
	//获取keywords信息: 返回data：common.PagedList
	List(ctx context.Context, in *common.Page, opts ...client.CallOption) (*common.Response, error)
	//将所有keywords从redis取出生成json文件
	PublistKeywords(ctx context.Context, in *common.Page, opts ...client.CallOption) (*common.Response, error)
	//通过name查询出对应carlist列表返回
	GetCarListByKeywords(ctx context.Context, in *KeywordsDto, opts ...client.CallOption) (*common.Response, error)
}

type keywordsService struct {
	c    client.Client
	name string
}

func NewKeywordsService(name string, c client.Client) KeywordsService {
	return &keywordsService{
		c:    c,
		name: name,
	}
}

func (c *keywordsService) List(ctx context.Context, in *common.Page, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "Keywords.List", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keywordsService) PublistKeywords(ctx context.Context, in *common.Page, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "Keywords.PublistKeywords", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keywordsService) GetCarListByKeywords(ctx context.Context, in *KeywordsDto, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "Keywords.GetCarListByKeywords", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Keywords service

type KeywordsHandler interface {
	//获取keywords信息: 返回data：common.PagedList
	List(context.Context, *common.Page, *common.Response) error
	//将所有keywords从redis取出生成json文件
	PublistKeywords(context.Context, *common.Page, *common.Response) error
	//通过name查询出对应carlist列表返回
	GetCarListByKeywords(context.Context, *KeywordsDto, *common.Response) error
}

func RegisterKeywordsHandler(s server.Server, hdlr KeywordsHandler, opts ...server.HandlerOption) error {
	type keywords interface {
		List(ctx context.Context, in *common.Page, out *common.Response) error
		PublistKeywords(ctx context.Context, in *common.Page, out *common.Response) error
		GetCarListByKeywords(ctx context.Context, in *KeywordsDto, out *common.Response) error
	}
	type Keywords struct {
		keywords
	}
	h := &keywordsHandler{hdlr}
	return s.Handle(s.NewHandler(&Keywords{h}, opts...))
}

type keywordsHandler struct {
	KeywordsHandler
}

func (h *keywordsHandler) List(ctx context.Context, in *common.Page, out *common.Response) error {
	return h.KeywordsHandler.List(ctx, in, out)
}

func (h *keywordsHandler) PublistKeywords(ctx context.Context, in *common.Page, out *common.Response) error {
	return h.KeywordsHandler.PublistKeywords(ctx, in, out)
}

func (h *keywordsHandler) GetCarListByKeywords(ctx context.Context, in *KeywordsDto, out *common.Response) error {
	return h.KeywordsHandler.GetCarListByKeywords(ctx, in, out)
}