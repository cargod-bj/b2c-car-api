// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: brandModel/year_engine_trans.proto

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

// Api Endpoints for YearEngineTrans service

func NewYearEngineTransEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for YearEngineTrans service

type YearEngineTransService interface {
	// 根据model查询year列表
	// response.Data = common.Page{
	//        List = List<IdNameDto>
	//      }
	FindYearByModel(ctx context.Context, in *common.IdDto, opts ...client.CallOption) (*common.Response, error)
	// 根据year查询engine列表
	// response.Data = common.Page{
	//         List = List<IdNameDto>
	//       }
	FindVariantByYear(ctx context.Context, in *ModelYearReq, opts ...client.CallOption) (*common.Response, error)
	// 根据year查询engine列表
	// response.Data = common.Page{
	//         List = List<IdNameDto>
	//       }
	FindEngineByVariant(ctx context.Context, in *ModelYearVariantReq, opts ...client.CallOption) (*common.Response, error)
	// 根据engine查询transmission列表
	// response.Data = common.Page{
	//          List = List<IdNameDto>
	//       }
	FindTransmissionByEngine(ctx context.Context, in *ModelYearVariantEngineReq, opts ...client.CallOption) (*common.Response, error)
	// 根据model查询year列表, 如果modelId=0, 将查询所有有车的model列表
	// response.Data = common.Page{
	//          List = List<IdNameDto>
	//       }
	FindYearByModelHasCar(ctx context.Context, in *common.IdLocalDTO, opts ...client.CallOption) (*common.Response, error)
	// 根据year查询variant列表, year不能为null, modelId 不能为null
	// response.Data = common.Page{
	//          List = List<IdNameDto>
	//       }
	FindVariantByYearHasCar(ctx context.Context, in *ModelYearReq, opts ...client.CallOption) (*common.Response, error)
	// 根据variant查询engine列表, 入参为variantId
	// response.Data = common.Page{
	//          List = List<IdNameDto>
	//       }
	FindEngineByVariantHasCar(ctx context.Context, in *FindEngineByVariantHasCarReq, opts ...client.CallOption) (*common.Response, error)
	// 根据brand、model、year、variant、engine的入参查询有车辆的列表
	// response.Data = common.Page{
	//          List = List<IdNameDto>
	//       }
	FindTransmissionHasCar(ctx context.Context, in *ModelYearVariantEngineReq, opts ...client.CallOption) (*common.Response, error)
}

type yearEngineTransService struct {
	c    client.Client
	name string
}

func NewYearEngineTransService(name string, c client.Client) YearEngineTransService {
	return &yearEngineTransService{
		c:    c,
		name: name,
	}
}

func (c *yearEngineTransService) FindYearByModel(ctx context.Context, in *common.IdDto, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "YearEngineTrans.FindYearByModel", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *yearEngineTransService) FindVariantByYear(ctx context.Context, in *ModelYearReq, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "YearEngineTrans.FindVariantByYear", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *yearEngineTransService) FindEngineByVariant(ctx context.Context, in *ModelYearVariantReq, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "YearEngineTrans.FindEngineByVariant", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *yearEngineTransService) FindTransmissionByEngine(ctx context.Context, in *ModelYearVariantEngineReq, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "YearEngineTrans.FindTransmissionByEngine", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *yearEngineTransService) FindYearByModelHasCar(ctx context.Context, in *common.IdLocalDTO, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "YearEngineTrans.FindYearByModelHasCar", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *yearEngineTransService) FindVariantByYearHasCar(ctx context.Context, in *ModelYearReq, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "YearEngineTrans.FindVariantByYearHasCar", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *yearEngineTransService) FindEngineByVariantHasCar(ctx context.Context, in *FindEngineByVariantHasCarReq, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "YearEngineTrans.FindEngineByVariantHasCar", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *yearEngineTransService) FindTransmissionHasCar(ctx context.Context, in *ModelYearVariantEngineReq, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "YearEngineTrans.FindTransmissionHasCar", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for YearEngineTrans service

type YearEngineTransHandler interface {
	// 根据model查询year列表
	// response.Data = common.Page{
	//        List = List<IdNameDto>
	//      }
	FindYearByModel(context.Context, *common.IdDto, *common.Response) error
	// 根据year查询engine列表
	// response.Data = common.Page{
	//         List = List<IdNameDto>
	//       }
	FindVariantByYear(context.Context, *ModelYearReq, *common.Response) error
	// 根据year查询engine列表
	// response.Data = common.Page{
	//         List = List<IdNameDto>
	//       }
	FindEngineByVariant(context.Context, *ModelYearVariantReq, *common.Response) error
	// 根据engine查询transmission列表
	// response.Data = common.Page{
	//          List = List<IdNameDto>
	//       }
	FindTransmissionByEngine(context.Context, *ModelYearVariantEngineReq, *common.Response) error
	// 根据model查询year列表, 如果modelId=0, 将查询所有有车的model列表
	// response.Data = common.Page{
	//          List = List<IdNameDto>
	//       }
	FindYearByModelHasCar(context.Context, *common.IdLocalDTO, *common.Response) error
	// 根据year查询variant列表, year不能为null, modelId 不能为null
	// response.Data = common.Page{
	//          List = List<IdNameDto>
	//       }
	FindVariantByYearHasCar(context.Context, *ModelYearReq, *common.Response) error
	// 根据variant查询engine列表, 入参为variantId
	// response.Data = common.Page{
	//          List = List<IdNameDto>
	//       }
	FindEngineByVariantHasCar(context.Context, *FindEngineByVariantHasCarReq, *common.Response) error
	// 根据brand、model、year、variant、engine的入参查询有车辆的列表
	// response.Data = common.Page{
	//          List = List<IdNameDto>
	//       }
	FindTransmissionHasCar(context.Context, *ModelYearVariantEngineReq, *common.Response) error
}

func RegisterYearEngineTransHandler(s server.Server, hdlr YearEngineTransHandler, opts ...server.HandlerOption) error {
	type yearEngineTrans interface {
		FindYearByModel(ctx context.Context, in *common.IdDto, out *common.Response) error
		FindVariantByYear(ctx context.Context, in *ModelYearReq, out *common.Response) error
		FindEngineByVariant(ctx context.Context, in *ModelYearVariantReq, out *common.Response) error
		FindTransmissionByEngine(ctx context.Context, in *ModelYearVariantEngineReq, out *common.Response) error
		FindYearByModelHasCar(ctx context.Context, in *common.IdLocalDTO, out *common.Response) error
		FindVariantByYearHasCar(ctx context.Context, in *ModelYearReq, out *common.Response) error
		FindEngineByVariantHasCar(ctx context.Context, in *FindEngineByVariantHasCarReq, out *common.Response) error
		FindTransmissionHasCar(ctx context.Context, in *ModelYearVariantEngineReq, out *common.Response) error
	}
	type YearEngineTrans struct {
		yearEngineTrans
	}
	h := &yearEngineTransHandler{hdlr}
	return s.Handle(s.NewHandler(&YearEngineTrans{h}, opts...))
}

type yearEngineTransHandler struct {
	YearEngineTransHandler
}

func (h *yearEngineTransHandler) FindYearByModel(ctx context.Context, in *common.IdDto, out *common.Response) error {
	return h.YearEngineTransHandler.FindYearByModel(ctx, in, out)
}

func (h *yearEngineTransHandler) FindVariantByYear(ctx context.Context, in *ModelYearReq, out *common.Response) error {
	return h.YearEngineTransHandler.FindVariantByYear(ctx, in, out)
}

func (h *yearEngineTransHandler) FindEngineByVariant(ctx context.Context, in *ModelYearVariantReq, out *common.Response) error {
	return h.YearEngineTransHandler.FindEngineByVariant(ctx, in, out)
}

func (h *yearEngineTransHandler) FindTransmissionByEngine(ctx context.Context, in *ModelYearVariantEngineReq, out *common.Response) error {
	return h.YearEngineTransHandler.FindTransmissionByEngine(ctx, in, out)
}

func (h *yearEngineTransHandler) FindYearByModelHasCar(ctx context.Context, in *common.IdLocalDTO, out *common.Response) error {
	return h.YearEngineTransHandler.FindYearByModelHasCar(ctx, in, out)
}

func (h *yearEngineTransHandler) FindVariantByYearHasCar(ctx context.Context, in *ModelYearReq, out *common.Response) error {
	return h.YearEngineTransHandler.FindVariantByYearHasCar(ctx, in, out)
}

func (h *yearEngineTransHandler) FindEngineByVariantHasCar(ctx context.Context, in *FindEngineByVariantHasCarReq, out *common.Response) error {
	return h.YearEngineTransHandler.FindEngineByVariantHasCar(ctx, in, out)
}

func (h *yearEngineTransHandler) FindTransmissionHasCar(ctx context.Context, in *ModelYearVariantEngineReq, out *common.Response) error {
	return h.YearEngineTransHandler.FindTransmissionHasCar(ctx, in, out)
}
