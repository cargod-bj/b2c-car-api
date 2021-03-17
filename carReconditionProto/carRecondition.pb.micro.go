// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: carReconditionProto/carRecondition.proto

package carReconditionProto

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

// Api Endpoints for CarRecondition service

func NewCarReconditionEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for CarRecondition service

type CarReconditionService interface {
	// 添加整备类型，返回 data：nil
	AddReconditionType(ctx context.Context, in *ReconditionTypePDto, opts ...client.CallOption) (*common.Response, error)
	// 批量添加整备类型，返回 data：nil
	AddReconditionTypeBatch(ctx context.Context, in *AddReconditionBatchPDto, opts ...client.CallOption) (*common.Response, error)
	// 更新指定整备类型，返回 data：nil
	UpdateReconditionType(ctx context.Context, in *ReconditionTypePDto, opts ...client.CallOption) (*common.Response, error)
	// 删除指定整备类型，返回 data：nil
	DeleteReconditionType(ctx context.Context, in *ReconditionTypeIdPDto, opts ...client.CallOption) (*common.Response, error)
	// 查询整备类型列表：返回 Data = common.PageList{
	//              List = List<ReconditionTypePDto>
	//          }
	ListReconditionType(ctx context.Context, in *ReconditionTypeListReq, opts ...client.CallOption) (*common.Response, error)
	// 以Category-Detail的父子结构返回所有的recondition类型：
	//     返回 Data = common.PageList{
	//              List = List<ReconditionCategoryPDto>
	//          }
	ListAllReconditionTypeWithCategory(ctx context.Context, in *ReconditionTypeIdPDto, opts ...client.CallOption) (*common.Response, error)
	// 查询整备类型列表：返回 Data = common.PageList{
	//              List = List<ReconditionTypePDto>
	//          }
	ListReconditionTypeByWorkshop(ctx context.Context, in *ReconditionTypeIdPDto, opts ...client.CallOption) (*common.Response, error)
	// 添加整备门店，返回 data：nil
	AddReconditionWorkshop(ctx context.Context, in *ReconditionWorkshopPDto, opts ...client.CallOption) (*common.Response, error)
	// 更新指定整备门店，返回 data：nil
	UpdateReconditionWorkshop(ctx context.Context, in *ReconditionWorkshopPDto, opts ...client.CallOption) (*common.Response, error)
	// 删除指定整备门店，返回 data：nil
	DeleteReconditionWorkshop(ctx context.Context, in *ReconditionTypeIdPDto, opts ...client.CallOption) (*common.Response, error)
	// 查询整备门店列表：返回 Data = common.PageList{
	//              List = List<ReconditionWorkshopPDto>
	//          }
	ListReconditionWorkshop(ctx context.Context, in *ReconditionWorkshopListReq, opts ...client.CallOption) (*common.Response, error)
}

type carReconditionService struct {
	c    client.Client
	name string
}

func NewCarReconditionService(name string, c client.Client) CarReconditionService {
	return &carReconditionService{
		c:    c,
		name: name,
	}
}

func (c *carReconditionService) AddReconditionType(ctx context.Context, in *ReconditionTypePDto, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "CarRecondition.AddReconditionType", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carReconditionService) AddReconditionTypeBatch(ctx context.Context, in *AddReconditionBatchPDto, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "CarRecondition.AddReconditionTypeBatch", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carReconditionService) UpdateReconditionType(ctx context.Context, in *ReconditionTypePDto, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "CarRecondition.UpdateReconditionType", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carReconditionService) DeleteReconditionType(ctx context.Context, in *ReconditionTypeIdPDto, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "CarRecondition.DeleteReconditionType", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carReconditionService) ListReconditionType(ctx context.Context, in *ReconditionTypeListReq, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "CarRecondition.ListReconditionType", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carReconditionService) ListAllReconditionTypeWithCategory(ctx context.Context, in *ReconditionTypeIdPDto, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "CarRecondition.ListAllReconditionTypeWithCategory", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carReconditionService) ListReconditionTypeByWorkshop(ctx context.Context, in *ReconditionTypeIdPDto, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "CarRecondition.ListReconditionTypeByWorkshop", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carReconditionService) AddReconditionWorkshop(ctx context.Context, in *ReconditionWorkshopPDto, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "CarRecondition.AddReconditionWorkshop", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carReconditionService) UpdateReconditionWorkshop(ctx context.Context, in *ReconditionWorkshopPDto, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "CarRecondition.UpdateReconditionWorkshop", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carReconditionService) DeleteReconditionWorkshop(ctx context.Context, in *ReconditionTypeIdPDto, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "CarRecondition.DeleteReconditionWorkshop", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carReconditionService) ListReconditionWorkshop(ctx context.Context, in *ReconditionWorkshopListReq, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "CarRecondition.ListReconditionWorkshop", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CarRecondition service

type CarReconditionHandler interface {
	// 添加整备类型，返回 data：nil
	AddReconditionType(context.Context, *ReconditionTypePDto, *common.Response) error
	// 批量添加整备类型，返回 data：nil
	AddReconditionTypeBatch(context.Context, *AddReconditionBatchPDto, *common.Response) error
	// 更新指定整备类型，返回 data：nil
	UpdateReconditionType(context.Context, *ReconditionTypePDto, *common.Response) error
	// 删除指定整备类型，返回 data：nil
	DeleteReconditionType(context.Context, *ReconditionTypeIdPDto, *common.Response) error
	// 查询整备类型列表：返回 Data = common.PageList{
	//              List = List<ReconditionTypePDto>
	//          }
	ListReconditionType(context.Context, *ReconditionTypeListReq, *common.Response) error
	// 以Category-Detail的父子结构返回所有的recondition类型：
	//     返回 Data = common.PageList{
	//              List = List<ReconditionCategoryPDto>
	//          }
	ListAllReconditionTypeWithCategory(context.Context, *ReconditionTypeIdPDto, *common.Response) error
	// 查询整备类型列表：返回 Data = common.PageList{
	//              List = List<ReconditionTypePDto>
	//          }
	ListReconditionTypeByWorkshop(context.Context, *ReconditionTypeIdPDto, *common.Response) error
	// 添加整备门店，返回 data：nil
	AddReconditionWorkshop(context.Context, *ReconditionWorkshopPDto, *common.Response) error
	// 更新指定整备门店，返回 data：nil
	UpdateReconditionWorkshop(context.Context, *ReconditionWorkshopPDto, *common.Response) error
	// 删除指定整备门店，返回 data：nil
	DeleteReconditionWorkshop(context.Context, *ReconditionTypeIdPDto, *common.Response) error
	// 查询整备门店列表：返回 Data = common.PageList{
	//              List = List<ReconditionWorkshopPDto>
	//          }
	ListReconditionWorkshop(context.Context, *ReconditionWorkshopListReq, *common.Response) error
}

func RegisterCarReconditionHandler(s server.Server, hdlr CarReconditionHandler, opts ...server.HandlerOption) error {
	type carRecondition interface {
		AddReconditionType(ctx context.Context, in *ReconditionTypePDto, out *common.Response) error
		AddReconditionTypeBatch(ctx context.Context, in *AddReconditionBatchPDto, out *common.Response) error
		UpdateReconditionType(ctx context.Context, in *ReconditionTypePDto, out *common.Response) error
		DeleteReconditionType(ctx context.Context, in *ReconditionTypeIdPDto, out *common.Response) error
		ListReconditionType(ctx context.Context, in *ReconditionTypeListReq, out *common.Response) error
		ListAllReconditionTypeWithCategory(ctx context.Context, in *ReconditionTypeIdPDto, out *common.Response) error
		ListReconditionTypeByWorkshop(ctx context.Context, in *ReconditionTypeIdPDto, out *common.Response) error
		AddReconditionWorkshop(ctx context.Context, in *ReconditionWorkshopPDto, out *common.Response) error
		UpdateReconditionWorkshop(ctx context.Context, in *ReconditionWorkshopPDto, out *common.Response) error
		DeleteReconditionWorkshop(ctx context.Context, in *ReconditionTypeIdPDto, out *common.Response) error
		ListReconditionWorkshop(ctx context.Context, in *ReconditionWorkshopListReq, out *common.Response) error
	}
	type CarRecondition struct {
		carRecondition
	}
	h := &carReconditionHandler{hdlr}
	return s.Handle(s.NewHandler(&CarRecondition{h}, opts...))
}

type carReconditionHandler struct {
	CarReconditionHandler
}

func (h *carReconditionHandler) AddReconditionType(ctx context.Context, in *ReconditionTypePDto, out *common.Response) error {
	return h.CarReconditionHandler.AddReconditionType(ctx, in, out)
}

func (h *carReconditionHandler) AddReconditionTypeBatch(ctx context.Context, in *AddReconditionBatchPDto, out *common.Response) error {
	return h.CarReconditionHandler.AddReconditionTypeBatch(ctx, in, out)
}

func (h *carReconditionHandler) UpdateReconditionType(ctx context.Context, in *ReconditionTypePDto, out *common.Response) error {
	return h.CarReconditionHandler.UpdateReconditionType(ctx, in, out)
}

func (h *carReconditionHandler) DeleteReconditionType(ctx context.Context, in *ReconditionTypeIdPDto, out *common.Response) error {
	return h.CarReconditionHandler.DeleteReconditionType(ctx, in, out)
}

func (h *carReconditionHandler) ListReconditionType(ctx context.Context, in *ReconditionTypeListReq, out *common.Response) error {
	return h.CarReconditionHandler.ListReconditionType(ctx, in, out)
}

func (h *carReconditionHandler) ListAllReconditionTypeWithCategory(ctx context.Context, in *ReconditionTypeIdPDto, out *common.Response) error {
	return h.CarReconditionHandler.ListAllReconditionTypeWithCategory(ctx, in, out)
}

func (h *carReconditionHandler) ListReconditionTypeByWorkshop(ctx context.Context, in *ReconditionTypeIdPDto, out *common.Response) error {
	return h.CarReconditionHandler.ListReconditionTypeByWorkshop(ctx, in, out)
}

func (h *carReconditionHandler) AddReconditionWorkshop(ctx context.Context, in *ReconditionWorkshopPDto, out *common.Response) error {
	return h.CarReconditionHandler.AddReconditionWorkshop(ctx, in, out)
}

func (h *carReconditionHandler) UpdateReconditionWorkshop(ctx context.Context, in *ReconditionWorkshopPDto, out *common.Response) error {
	return h.CarReconditionHandler.UpdateReconditionWorkshop(ctx, in, out)
}

func (h *carReconditionHandler) DeleteReconditionWorkshop(ctx context.Context, in *ReconditionTypeIdPDto, out *common.Response) error {
	return h.CarReconditionHandler.DeleteReconditionWorkshop(ctx, in, out)
}

func (h *carReconditionHandler) ListReconditionWorkshop(ctx context.Context, in *ReconditionWorkshopListReq, out *common.Response) error {
	return h.CarReconditionHandler.ListReconditionWorkshop(ctx, in, out)
}
