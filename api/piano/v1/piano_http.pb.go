// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.0.0-rc7

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

type PianoHTTPServer interface {
	CreatePiano(context.Context, *CreatePianoRequest) (*CreatePianoReply, error)
	LoginIn(context.Context, *CreateLogInRequest) (*CommonResponse, error)
	RegisterUser(context.Context, *CreateRegisterRequest) (*CommonResponse, error)
}

func RegisterPianoHTTPServer(s *http.Server, srv PianoHTTPServer) {
	r := s.Route("/")
	r.POST("/v1.service.create", _Piano_CreatePiano0_HTTP_Handler(srv))
	r.POST("/v1.service.register", _Piano_RegisterUser0_HTTP_Handler(srv))
	r.POST("/v1.service.login", _Piano_LoginIn0_HTTP_Handler(srv))
}

func _Piano_CreatePiano0_HTTP_Handler(srv PianoHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreatePianoRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.piano.v1.Piano/CreatePiano")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreatePiano(ctx, req.(*CreatePianoRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreatePianoReply)
		return ctx.Result(200, reply)
	}
}

func _Piano_RegisterUser0_HTTP_Handler(srv PianoHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateRegisterRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.piano.v1.Piano/RegisterUser")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.RegisterUser(ctx, req.(*CreateRegisterRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CommonResponse)
		return ctx.Result(200, reply)
	}
}

func _Piano_LoginIn0_HTTP_Handler(srv PianoHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateLogInRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.piano.v1.Piano/LoginIn")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.LoginIn(ctx, req.(*CreateLogInRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CommonResponse)
		return ctx.Result(200, reply)
	}
}

type PianoHTTPClient interface {
	CreatePiano(ctx context.Context, req *CreatePianoRequest, opts ...http.CallOption) (rsp *CreatePianoReply, err error)
	LoginIn(ctx context.Context, req *CreateLogInRequest, opts ...http.CallOption) (rsp *CommonResponse, err error)
	RegisterUser(ctx context.Context, req *CreateRegisterRequest, opts ...http.CallOption) (rsp *CommonResponse, err error)
}

type PianoHTTPClientImpl struct {
	cc *http.Client
}

func NewPianoHTTPClient(client *http.Client) PianoHTTPClient {
	return &PianoHTTPClientImpl{client}
}

func (c *PianoHTTPClientImpl) CreatePiano(ctx context.Context, in *CreatePianoRequest, opts ...http.CallOption) (*CreatePianoReply, error) {
	var out CreatePianoReply
	pattern := "/v1.service.create"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/api.piano.v1.Piano/CreatePiano"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *PianoHTTPClientImpl) LoginIn(ctx context.Context, in *CreateLogInRequest, opts ...http.CallOption) (*CommonResponse, error) {
	var out CommonResponse
	pattern := "/v1.service.login"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/api.piano.v1.Piano/LoginIn"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *PianoHTTPClientImpl) RegisterUser(ctx context.Context, in *CreateRegisterRequest, opts ...http.CallOption) (*CommonResponse, error) {
	var out CommonResponse
	pattern := "/v1.service.register"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/api.piano.v1.Piano/RegisterUser"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
