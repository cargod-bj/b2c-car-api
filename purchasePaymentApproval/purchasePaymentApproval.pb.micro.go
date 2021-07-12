// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: purchasePaymentApproval/purchasePaymentApproval.proto

package purchasePaymentApproval

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

// Api Endpoints for PurchasePaymentApproval service

func NewPurchasePaymentApprovalEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for PurchasePaymentApproval service

type PurchasePaymentApprovalService interface {
	// 扫描生成审批记录数据
	Scan(ctx context.Context, in *ApprovalScanReq, opts ...client.CallOption) (*common.Response, error)
	List(ctx context.Context, in *ApprovalApprovalCondition, opts ...client.CallOption) (*common.Response, error)
	Apply(ctx context.Context, in *ApplyReq, opts ...client.CallOption) (*common.Response, error)
	Cancel(ctx context.Context, in *IdWithUserReq, opts ...client.CallOption) (*common.Response, error)
	Callback(ctx context.Context, in *DingTalkCallbackReq, opts ...client.CallOption) (*common.Response, error)
}

type purchasePaymentApprovalService struct {
	c    client.Client
	name string
}

func NewPurchasePaymentApprovalService(name string, c client.Client) PurchasePaymentApprovalService {
	return &purchasePaymentApprovalService{
		c:    c,
		name: name,
	}
}

func (c *purchasePaymentApprovalService) Scan(ctx context.Context, in *ApprovalScanReq, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "PurchasePaymentApproval.Scan", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *purchasePaymentApprovalService) List(ctx context.Context, in *ApprovalApprovalCondition, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "PurchasePaymentApproval.List", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *purchasePaymentApprovalService) Apply(ctx context.Context, in *ApplyReq, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "PurchasePaymentApproval.Apply", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *purchasePaymentApprovalService) Cancel(ctx context.Context, in *IdWithUserReq, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "PurchasePaymentApproval.Cancel", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *purchasePaymentApprovalService) Callback(ctx context.Context, in *DingTalkCallbackReq, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "PurchasePaymentApproval.Callback", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for PurchasePaymentApproval service

type PurchasePaymentApprovalHandler interface {
	// 扫描生成审批记录数据
	Scan(context.Context, *ApprovalScanReq, *common.Response) error
	List(context.Context, *ApprovalApprovalCondition, *common.Response) error
	Apply(context.Context, *ApplyReq, *common.Response) error
	Cancel(context.Context, *IdWithUserReq, *common.Response) error
	Callback(context.Context, *DingTalkCallbackReq, *common.Response) error
}

func RegisterPurchasePaymentApprovalHandler(s server.Server, hdlr PurchasePaymentApprovalHandler, opts ...server.HandlerOption) error {
	type purchasePaymentApproval interface {
		Scan(ctx context.Context, in *ApprovalScanReq, out *common.Response) error
		List(ctx context.Context, in *ApprovalApprovalCondition, out *common.Response) error
		Apply(ctx context.Context, in *ApplyReq, out *common.Response) error
		Cancel(ctx context.Context, in *IdWithUserReq, out *common.Response) error
		Callback(ctx context.Context, in *DingTalkCallbackReq, out *common.Response) error
	}
	type PurchasePaymentApproval struct {
		purchasePaymentApproval
	}
	h := &purchasePaymentApprovalHandler{hdlr}
	return s.Handle(s.NewHandler(&PurchasePaymentApproval{h}, opts...))
}

type purchasePaymentApprovalHandler struct {
	PurchasePaymentApprovalHandler
}

func (h *purchasePaymentApprovalHandler) Scan(ctx context.Context, in *ApprovalScanReq, out *common.Response) error {
	return h.PurchasePaymentApprovalHandler.Scan(ctx, in, out)
}

func (h *purchasePaymentApprovalHandler) List(ctx context.Context, in *ApprovalApprovalCondition, out *common.Response) error {
	return h.PurchasePaymentApprovalHandler.List(ctx, in, out)
}

func (h *purchasePaymentApprovalHandler) Apply(ctx context.Context, in *ApplyReq, out *common.Response) error {
	return h.PurchasePaymentApprovalHandler.Apply(ctx, in, out)
}

func (h *purchasePaymentApprovalHandler) Cancel(ctx context.Context, in *IdWithUserReq, out *common.Response) error {
	return h.PurchasePaymentApprovalHandler.Cancel(ctx, in, out)
}

func (h *purchasePaymentApprovalHandler) Callback(ctx context.Context, in *DingTalkCallbackReq, out *common.Response) error {
	return h.PurchasePaymentApprovalHandler.Callback(ctx, in, out)
}
