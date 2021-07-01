// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: carPurchase/car_purchase.proto

package carPurchaseProto

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

// Api Endpoints for CarPurchase service

func NewCarPurchaseEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for CarPurchase service

type CarPurchaseService interface {
	GetCarByInspection(ctx context.Context, in *InspectionReq, opts ...client.CallOption) (*common.Response, error)
}

type carPurchaseService struct {
	c    client.Client
	name string
}

func NewCarPurchaseService(name string, c client.Client) CarPurchaseService {
	return &carPurchaseService{
		c:    c,
		name: name,
	}
}

func (c *carPurchaseService) GetCarByInspection(ctx context.Context, in *InspectionReq, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "CarPurchase.GetCarByInspection", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CarPurchase service

type CarPurchaseHandler interface {
	GetCarByInspection(context.Context, *InspectionReq, *common.Response) error
}

func RegisterCarPurchaseHandler(s server.Server, hdlr CarPurchaseHandler, opts ...server.HandlerOption) error {
	type carPurchase interface {
		GetCarByInspection(ctx context.Context, in *InspectionReq, out *common.Response) error
	}
	type CarPurchase struct {
		carPurchase
	}
	h := &carPurchaseHandler{hdlr}
	return s.Handle(s.NewHandler(&CarPurchase{h}, opts...))
}

type carPurchaseHandler struct {
	CarPurchaseHandler
}

func (h *carPurchaseHandler) GetCarByInspection(ctx context.Context, in *InspectionReq, out *common.Response) error {
	return h.CarPurchaseHandler.GetCarByInspection(ctx, in, out)
}
