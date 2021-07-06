// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: uploadPurchaseFile.proto

package uploadPurchaseFile

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

// Api Endpoints for UploadPurchaseFile service

func NewUploadPurchaseFileEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for UploadPurchaseFile service

type UploadPurchaseFileService interface {
	// save fileadress to mysql service
	SaveAdress(ctx context.Context, in *SaveAdressReq, opts ...client.CallOption) (*common.Response, error)
}

type uploadPurchaseFileService struct {
	c    client.Client
	name string
}

func NewUploadPurchaseFileService(name string, c client.Client) UploadPurchaseFileService {
	return &uploadPurchaseFileService{
		c:    c,
		name: name,
	}
}

func (c *uploadPurchaseFileService) SaveAdress(ctx context.Context, in *SaveAdressReq, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "UploadPurchaseFile.SaveAdress", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UploadPurchaseFile service

type UploadPurchaseFileHandler interface {
	// save fileadress to mysql service
	SaveAdress(context.Context, *SaveAdressReq, *common.Response) error
}

func RegisterUploadPurchaseFileHandler(s server.Server, hdlr UploadPurchaseFileHandler, opts ...server.HandlerOption) error {
	type uploadPurchaseFile interface {
		SaveAdress(ctx context.Context, in *SaveAdressReq, out *common.Response) error
	}
	type UploadPurchaseFile struct {
		uploadPurchaseFile
	}
	h := &uploadPurchaseFileHandler{hdlr}
	return s.Handle(s.NewHandler(&UploadPurchaseFile{h}, opts...))
}

type uploadPurchaseFileHandler struct {
	UploadPurchaseFileHandler
}

func (h *uploadPurchaseFileHandler) SaveAdress(ctx context.Context, in *SaveAdressReq, out *common.Response) error {
	return h.UploadPurchaseFileHandler.SaveAdress(ctx, in, out)
}
