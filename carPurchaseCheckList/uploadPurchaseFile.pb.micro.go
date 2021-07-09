// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: carPurchaseCheckList/uploadPurchaseFile.proto

package CarPurCheckList

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

// Api Endpoints for CheckListService service

func NewCheckListServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for CheckListService service

type CheckListService interface {
	// save fileadress to database table carsource
	SaveAdress(ctx context.Context, in *PurchaseFileDto, opts ...client.CallOption) (*common.Response, error)
	//create a checklist based on the car-purchase ID
	Create(ctx context.Context, in *CreateReq, opts ...client.CallOption) (*common.Response, error)
	Creates(ctx context.Context, in *BatchCreate, opts ...client.CallOption) (*common.Response, error)
	//get a checklist according to the car-purchase ID
	GetList(ctx context.Context, in *CarPurchaseIdReq, opts ...client.CallOption) (*common.Response, error)
	//update checklist according to the ID
	Update(ctx context.Context, in *UpdateReq, opts ...client.CallOption) (*common.Response, error)
	//drop checklist according to the ID
	Delete(ctx context.Context, in *IdReq, opts ...client.CallOption) (*common.Response, error)
}

type checkListService struct {
	c    client.Client
	name string
}

func NewCheckListService(name string, c client.Client) CheckListService {
	return &checkListService{
		c:    c,
		name: name,
	}
}

func (c *checkListService) SaveAdress(ctx context.Context, in *PurchaseFileDto, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "CheckListService.SaveAdress", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *checkListService) Create(ctx context.Context, in *CreateReq, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "CheckListService.Create", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *checkListService) Creates(ctx context.Context, in *BatchCreate, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "CheckListService.Creates", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *checkListService) GetList(ctx context.Context, in *CarPurchaseIdReq, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "CheckListService.GetList", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *checkListService) Update(ctx context.Context, in *UpdateReq, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "CheckListService.Update", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *checkListService) Delete(ctx context.Context, in *IdReq, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "CheckListService.Delete", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CheckListService service

type CheckListServiceHandler interface {
	// save fileadress to database table carsource
	SaveAdress(context.Context, *PurchaseFileDto, *common.Response) error
	//create a checklist based on the car-purchase ID
	Create(context.Context, *CreateReq, *common.Response) error
	Creates(context.Context, *BatchCreate, *common.Response) error
	//get a checklist according to the car-purchase ID
	GetList(context.Context, *CarPurchaseIdReq, *common.Response) error
	//update checklist according to the ID
	Update(context.Context, *UpdateReq, *common.Response) error
	//drop checklist according to the ID
	Delete(context.Context, *IdReq, *common.Response) error
}

func RegisterCheckListServiceHandler(s server.Server, hdlr CheckListServiceHandler, opts ...server.HandlerOption) error {
	type checkListService interface {
		SaveAdress(ctx context.Context, in *PurchaseFileDto, out *common.Response) error
		Create(ctx context.Context, in *CreateReq, out *common.Response) error
		Creates(ctx context.Context, in *BatchCreate, out *common.Response) error
		GetList(ctx context.Context, in *CarPurchaseIdReq, out *common.Response) error
		Update(ctx context.Context, in *UpdateReq, out *common.Response) error
		Delete(ctx context.Context, in *IdReq, out *common.Response) error
	}
	type CheckListService struct {
		checkListService
	}
	h := &checkListServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&CheckListService{h}, opts...))
}

type checkListServiceHandler struct {
	CheckListServiceHandler
}

func (h *checkListServiceHandler) SaveAdress(ctx context.Context, in *PurchaseFileDto, out *common.Response) error {
	return h.CheckListServiceHandler.SaveAdress(ctx, in, out)
}

func (h *checkListServiceHandler) Create(ctx context.Context, in *CreateReq, out *common.Response) error {
	return h.CheckListServiceHandler.Create(ctx, in, out)
}

func (h *checkListServiceHandler) Creates(ctx context.Context, in *BatchCreate, out *common.Response) error {
	return h.CheckListServiceHandler.Creates(ctx, in, out)
}

func (h *checkListServiceHandler) GetList(ctx context.Context, in *CarPurchaseIdReq, out *common.Response) error {
	return h.CheckListServiceHandler.GetList(ctx, in, out)
}

func (h *checkListServiceHandler) Update(ctx context.Context, in *UpdateReq, out *common.Response) error {
	return h.CheckListServiceHandler.Update(ctx, in, out)
}

func (h *checkListServiceHandler) Delete(ctx context.Context, in *IdReq, out *common.Response) error {
	return h.CheckListServiceHandler.Delete(ctx, in, out)
}
