// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: history/history.proto

package historyProto

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

// Api Endpoints for History service

func NewHistoryEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for History service

type HistoryService interface {
	// 添加用户历史记录，返回 data：History 类型
	SaveHistory(ctx context.Context, in *History, opts ...client.CallOption) (*common.Response, error)
	// 根据Userid获取历史记录，返回 data：nil
	GetHistory(ctx context.Context, in *History, opts ...client.CallOption) (*common.Response, error)
}

type historyService struct {
	c    client.Client
	name string
}

func NewHistoryService(name string, c client.Client) HistoryService {
	return &historyService{
		c:    c,
		name: name,
	}
}

func (c *historyService) SaveHistory(ctx context.Context, in *History, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "History.SaveHistory", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *historyService) GetHistory(ctx context.Context, in *History, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "History.GetHistory", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for History service

type HistoryHandler interface {
	// 添加用户历史记录，返回 data：History 类型
	SaveHistory(context.Context, *History, *common.Response) error
	// 根据Userid获取历史记录，返回 data：nil
	GetHistory(context.Context, *History, *common.Response) error
}

func RegisterHistoryHandler(s server.Server, hdlr HistoryHandler, opts ...server.HandlerOption) error {
	type history interface {
		SaveHistory(ctx context.Context, in *History, out *common.Response) error
		GetHistory(ctx context.Context, in *History, out *common.Response) error
	}
	type History struct {
		history
	}
	h := &historyHandler{hdlr}
	return s.Handle(s.NewHandler(&History{h}, opts...))
}

type historyHandler struct {
	HistoryHandler
}

func (h *historyHandler) SaveHistory(ctx context.Context, in *History, out *common.Response) error {
	return h.HistoryHandler.SaveHistory(ctx, in, out)
}

func (h *historyHandler) GetHistory(ctx context.Context, in *History, out *common.Response) error {
	return h.HistoryHandler.GetHistory(ctx, in, out)
}
