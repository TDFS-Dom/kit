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

type AuthHTTPServer interface {
	ForgetPassword(context.Context, *ForgetPasswordRequest) (*ForgetPasswordReply, error)
	GetCsrfToken(context.Context, *GetCsrfTokenRequest) (*GetCsrfTokenResponse, error)
	GetLoginForm(context.Context, *GetLoginFormRequest) (*GetLoginFormResponse, error)
	Login(context.Context, *LoginAuthRequest) (*LoginAuthReply, error)
	LoginPasswordless(context.Context, *LoginPasswordlessRequest) (*LoginPasswordlessReply, error)
	Refresh(context.Context, *RefreshTokenAuthRequest) (*RefreshTokenAuthReply, error)
	Register(context.Context, *RegisterAuthRequest) (*RegisterAuthReply, error)
	SendForgetPasswordToken(context.Context, *ForgetPasswordTokenRequest) (*ForgetPasswordTokenReply, error)
	SendPasswordlessToken(context.Context, *PasswordlessTokenAuthRequest) (*PasswordlessTokenAuthReply, error)
	Token(context.Context, *TokenRequest) (*TokenReply, error)
	ValidatePassword(context.Context, *ValidatePasswordRequest) (*ValidatePasswordReply, error)
}

func RegisterAuthHTTPServer(s *http.Server, srv AuthHTTPServer) {
	r := s.Route("/")
	r.POST("/v1/auth/register", _Auth_Register0_HTTP_Handler(srv))
	r.POST("/v1/auth/login", _Auth_Login0_HTTP_Handler(srv))
	r.GET("/v1/auth/login", _Auth_GetLoginForm0_HTTP_Handler(srv))
	r.POST("/v1/auth/token", _Auth_Token0_HTTP_Handler(srv))
	r.POST("/v1/auth/refresh", _Auth_Refresh0_HTTP_Handler(srv))
	r.POST("/v1/auth/action/passwordless", _Auth_SendPasswordlessToken0_HTTP_Handler(srv))
	r.POST("/v1/auth/login/passwordless", _Auth_LoginPasswordless0_HTTP_Handler(srv))
	r.POST("/v1/auth/action/forget", _Auth_SendForgetPasswordToken0_HTTP_Handler(srv))
	r.POST("/v1/auth/password/forget", _Auth_ForgetPassword0_HTTP_Handler(srv))
	r.POST("/v1/auth/password/validate", _Auth_ValidatePassword0_HTTP_Handler(srv))
	r.GET("/v1/auth/csrf", _Auth_GetCsrfToken0_HTTP_Handler(srv))
}

func _Auth_Register0_HTTP_Handler(srv AuthHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in RegisterAuthRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/user.api.auth.v1.Auth/Register")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Register(ctx, req.(*RegisterAuthRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*RegisterAuthReply)
		return ctx.Result(200, reply)
	}
}

func _Auth_Login0_HTTP_Handler(srv AuthHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in LoginAuthRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/user.api.auth.v1.Auth/Login")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Login(ctx, req.(*LoginAuthRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*LoginAuthReply)
		return ctx.Result(200, reply)
	}
}

func _Auth_GetLoginForm0_HTTP_Handler(srv AuthHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetLoginFormRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/user.api.auth.v1.Auth/GetLoginForm")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetLoginForm(ctx, req.(*GetLoginFormRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetLoginFormResponse)
		return ctx.Result(200, reply)
	}
}

func _Auth_Token0_HTTP_Handler(srv AuthHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in TokenRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/user.api.auth.v1.Auth/Token")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Token(ctx, req.(*TokenRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*TokenReply)
		return ctx.Result(200, reply)
	}
}

func _Auth_Refresh0_HTTP_Handler(srv AuthHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in RefreshTokenAuthRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/user.api.auth.v1.Auth/Refresh")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Refresh(ctx, req.(*RefreshTokenAuthRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*RefreshTokenAuthReply)
		return ctx.Result(200, reply)
	}
}

func _Auth_SendPasswordlessToken0_HTTP_Handler(srv AuthHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in PasswordlessTokenAuthRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/user.api.auth.v1.Auth/SendPasswordlessToken")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SendPasswordlessToken(ctx, req.(*PasswordlessTokenAuthRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*PasswordlessTokenAuthReply)
		return ctx.Result(200, reply)
	}
}

func _Auth_LoginPasswordless0_HTTP_Handler(srv AuthHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in LoginPasswordlessRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/user.api.auth.v1.Auth/LoginPasswordless")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.LoginPasswordless(ctx, req.(*LoginPasswordlessRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*LoginPasswordlessReply)
		return ctx.Result(200, reply)
	}
}

func _Auth_SendForgetPasswordToken0_HTTP_Handler(srv AuthHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ForgetPasswordTokenRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/user.api.auth.v1.Auth/SendForgetPasswordToken")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SendForgetPasswordToken(ctx, req.(*ForgetPasswordTokenRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ForgetPasswordTokenReply)
		return ctx.Result(200, reply)
	}
}

func _Auth_ForgetPassword0_HTTP_Handler(srv AuthHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ForgetPasswordRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/user.api.auth.v1.Auth/ForgetPassword")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ForgetPassword(ctx, req.(*ForgetPasswordRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ForgetPasswordReply)
		return ctx.Result(200, reply)
	}
}

func _Auth_ValidatePassword0_HTTP_Handler(srv AuthHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ValidatePasswordRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/user.api.auth.v1.Auth/ValidatePassword")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ValidatePassword(ctx, req.(*ValidatePasswordRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ValidatePasswordReply)
		return ctx.Result(200, reply)
	}
}

func _Auth_GetCsrfToken0_HTTP_Handler(srv AuthHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetCsrfTokenRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/user.api.auth.v1.Auth/GetCsrfToken")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetCsrfToken(ctx, req.(*GetCsrfTokenRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetCsrfTokenResponse)
		return ctx.Result(200, reply)
	}
}

type AuthHTTPClient interface {
	ForgetPassword(ctx context.Context, req *ForgetPasswordRequest, opts ...http.CallOption) (rsp *ForgetPasswordReply, err error)
	GetCsrfToken(ctx context.Context, req *GetCsrfTokenRequest, opts ...http.CallOption) (rsp *GetCsrfTokenResponse, err error)
	GetLoginForm(ctx context.Context, req *GetLoginFormRequest, opts ...http.CallOption) (rsp *GetLoginFormResponse, err error)
	Login(ctx context.Context, req *LoginAuthRequest, opts ...http.CallOption) (rsp *LoginAuthReply, err error)
	LoginPasswordless(ctx context.Context, req *LoginPasswordlessRequest, opts ...http.CallOption) (rsp *LoginPasswordlessReply, err error)
	Refresh(ctx context.Context, req *RefreshTokenAuthRequest, opts ...http.CallOption) (rsp *RefreshTokenAuthReply, err error)
	Register(ctx context.Context, req *RegisterAuthRequest, opts ...http.CallOption) (rsp *RegisterAuthReply, err error)
	SendForgetPasswordToken(ctx context.Context, req *ForgetPasswordTokenRequest, opts ...http.CallOption) (rsp *ForgetPasswordTokenReply, err error)
	SendPasswordlessToken(ctx context.Context, req *PasswordlessTokenAuthRequest, opts ...http.CallOption) (rsp *PasswordlessTokenAuthReply, err error)
	Token(ctx context.Context, req *TokenRequest, opts ...http.CallOption) (rsp *TokenReply, err error)
	ValidatePassword(ctx context.Context, req *ValidatePasswordRequest, opts ...http.CallOption) (rsp *ValidatePasswordReply, err error)
}

type AuthHTTPClientImpl struct {
	cc *http.Client
}

func NewAuthHTTPClient(client *http.Client) AuthHTTPClient {
	return &AuthHTTPClientImpl{client}
}

func (c *AuthHTTPClientImpl) ForgetPassword(ctx context.Context, in *ForgetPasswordRequest, opts ...http.CallOption) (*ForgetPasswordReply, error) {
	var out ForgetPasswordReply
	pattern := "/v1/auth/password/forget"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/user.api.auth.v1.Auth/ForgetPassword"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AuthHTTPClientImpl) GetCsrfToken(ctx context.Context, in *GetCsrfTokenRequest, opts ...http.CallOption) (*GetCsrfTokenResponse, error) {
	var out GetCsrfTokenResponse
	pattern := "/v1/auth/csrf"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/user.api.auth.v1.Auth/GetCsrfToken"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AuthHTTPClientImpl) GetLoginForm(ctx context.Context, in *GetLoginFormRequest, opts ...http.CallOption) (*GetLoginFormResponse, error) {
	var out GetLoginFormResponse
	pattern := "/v1/auth/login"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/user.api.auth.v1.Auth/GetLoginForm"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AuthHTTPClientImpl) Login(ctx context.Context, in *LoginAuthRequest, opts ...http.CallOption) (*LoginAuthReply, error) {
	var out LoginAuthReply
	pattern := "/v1/auth/login"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/user.api.auth.v1.Auth/Login"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AuthHTTPClientImpl) LoginPasswordless(ctx context.Context, in *LoginPasswordlessRequest, opts ...http.CallOption) (*LoginPasswordlessReply, error) {
	var out LoginPasswordlessReply
	pattern := "/v1/auth/login/passwordless"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/user.api.auth.v1.Auth/LoginPasswordless"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AuthHTTPClientImpl) Refresh(ctx context.Context, in *RefreshTokenAuthRequest, opts ...http.CallOption) (*RefreshTokenAuthReply, error) {
	var out RefreshTokenAuthReply
	pattern := "/v1/auth/refresh"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/user.api.auth.v1.Auth/Refresh"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AuthHTTPClientImpl) Register(ctx context.Context, in *RegisterAuthRequest, opts ...http.CallOption) (*RegisterAuthReply, error) {
	var out RegisterAuthReply
	pattern := "/v1/auth/register"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/user.api.auth.v1.Auth/Register"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AuthHTTPClientImpl) SendForgetPasswordToken(ctx context.Context, in *ForgetPasswordTokenRequest, opts ...http.CallOption) (*ForgetPasswordTokenReply, error) {
	var out ForgetPasswordTokenReply
	pattern := "/v1/auth/action/forget"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/user.api.auth.v1.Auth/SendForgetPasswordToken"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AuthHTTPClientImpl) SendPasswordlessToken(ctx context.Context, in *PasswordlessTokenAuthRequest, opts ...http.CallOption) (*PasswordlessTokenAuthReply, error) {
	var out PasswordlessTokenAuthReply
	pattern := "/v1/auth/action/passwordless"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/user.api.auth.v1.Auth/SendPasswordlessToken"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AuthHTTPClientImpl) Token(ctx context.Context, in *TokenRequest, opts ...http.CallOption) (*TokenReply, error) {
	var out TokenReply
	pattern := "/v1/auth/token"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/user.api.auth.v1.Auth/Token"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AuthHTTPClientImpl) ValidatePassword(ctx context.Context, in *ValidatePasswordRequest, opts ...http.CallOption) (*ValidatePasswordReply, error) {
	var out ValidatePasswordReply
	pattern := "/v1/auth/password/validate"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/user.api.auth.v1.Auth/ValidatePassword"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
