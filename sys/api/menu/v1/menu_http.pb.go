// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.1.4

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

type MenuServiceHTTPServer interface {
	CreateMenu(context.Context, *CreateMenuRequest) (*Menu, error)
	DeleteMenu(context.Context, *DeleteMenuRequest) (*DeleteMenuReply, error)
	GetMenu(context.Context, *GetMenuRequest) (*Menu, error)
	ListMenu(context.Context, *ListMenuRequest) (*ListMenuReply, error)
	UpdateMenu(context.Context, *UpdateMenuRequest) (*Menu, error)
}

func RegisterMenuServiceHTTPServer(s *http.Server, srv MenuServiceHTTPServer) {
	r := s.Route("/")
	r.POST("/v1/sys/menu/list", _MenuService_ListMenu0_HTTP_Handler(srv))
	r.GET("/v1/sys/menus", _MenuService_ListMenu1_HTTP_Handler(srv))
	r.GET("/v1/sys/menu/{id}", _MenuService_GetMenu0_HTTP_Handler(srv))
	r.POST("/v1/sys/menu", _MenuService_CreateMenu0_HTTP_Handler(srv))
	r.PATCH("/v1/sys/menu/{menu.id}", _MenuService_UpdateMenu0_HTTP_Handler(srv))
	r.PUT("/v1/sys/menu/{menu.id}", _MenuService_UpdateMenu1_HTTP_Handler(srv))
	r.DELETE("/v1/sys/menu/{id}", _MenuService_DeleteMenu0_HTTP_Handler(srv))
}

func _MenuService_ListMenu0_HTTP_Handler(srv MenuServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListMenuRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.menu.v1.MenuService/ListMenu")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListMenu(ctx, req.(*ListMenuRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListMenuReply)
		return ctx.Result(200, reply)
	}
}

func _MenuService_ListMenu1_HTTP_Handler(srv MenuServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListMenuRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.menu.v1.MenuService/ListMenu")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListMenu(ctx, req.(*ListMenuRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListMenuReply)
		return ctx.Result(200, reply)
	}
}

func _MenuService_GetMenu0_HTTP_Handler(srv MenuServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetMenuRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.menu.v1.MenuService/GetMenu")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetMenu(ctx, req.(*GetMenuRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Menu)
		return ctx.Result(200, reply)
	}
}

func _MenuService_CreateMenu0_HTTP_Handler(srv MenuServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateMenuRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.menu.v1.MenuService/CreateMenu")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateMenu(ctx, req.(*CreateMenuRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Menu)
		return ctx.Result(200, reply)
	}
}

func _MenuService_UpdateMenu0_HTTP_Handler(srv MenuServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateMenuRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.menu.v1.MenuService/UpdateMenu")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateMenu(ctx, req.(*UpdateMenuRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Menu)
		return ctx.Result(200, reply)
	}
}

func _MenuService_UpdateMenu1_HTTP_Handler(srv MenuServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateMenuRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.menu.v1.MenuService/UpdateMenu")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateMenu(ctx, req.(*UpdateMenuRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Menu)
		return ctx.Result(200, reply)
	}
}

func _MenuService_DeleteMenu0_HTTP_Handler(srv MenuServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteMenuRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.menu.v1.MenuService/DeleteMenu")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteMenu(ctx, req.(*DeleteMenuRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteMenuReply)
		return ctx.Result(200, reply)
	}
}

type MenuServiceHTTPClient interface {
	CreateMenu(ctx context.Context, req *CreateMenuRequest, opts ...http.CallOption) (rsp *Menu, err error)
	DeleteMenu(ctx context.Context, req *DeleteMenuRequest, opts ...http.CallOption) (rsp *DeleteMenuReply, err error)
	GetMenu(ctx context.Context, req *GetMenuRequest, opts ...http.CallOption) (rsp *Menu, err error)
	ListMenu(ctx context.Context, req *ListMenuRequest, opts ...http.CallOption) (rsp *ListMenuReply, err error)
	UpdateMenu(ctx context.Context, req *UpdateMenuRequest, opts ...http.CallOption) (rsp *Menu, err error)
}

type MenuServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewMenuServiceHTTPClient(client *http.Client) MenuServiceHTTPClient {
	return &MenuServiceHTTPClientImpl{client}
}

func (c *MenuServiceHTTPClientImpl) CreateMenu(ctx context.Context, in *CreateMenuRequest, opts ...http.CallOption) (*Menu, error) {
	var out Menu
	pattern := "/v1/sys/menu"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/api.menu.v1.MenuService/CreateMenu"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *MenuServiceHTTPClientImpl) DeleteMenu(ctx context.Context, in *DeleteMenuRequest, opts ...http.CallOption) (*DeleteMenuReply, error) {
	var out DeleteMenuReply
	pattern := "/v1/sys/menu/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/api.menu.v1.MenuService/DeleteMenu"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *MenuServiceHTTPClientImpl) GetMenu(ctx context.Context, in *GetMenuRequest, opts ...http.CallOption) (*Menu, error) {
	var out Menu
	pattern := "/v1/sys/menu/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/api.menu.v1.MenuService/GetMenu"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *MenuServiceHTTPClientImpl) ListMenu(ctx context.Context, in *ListMenuRequest, opts ...http.CallOption) (*ListMenuReply, error) {
	var out ListMenuReply
	pattern := "/v1/sys/menus"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/api.menu.v1.MenuService/ListMenu"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *MenuServiceHTTPClientImpl) UpdateMenu(ctx context.Context, in *UpdateMenuRequest, opts ...http.CallOption) (*Menu, error) {
	var out Menu
	pattern := "/v1/sys/menu/{menu.id}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/api.menu.v1.MenuService/UpdateMenu"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
