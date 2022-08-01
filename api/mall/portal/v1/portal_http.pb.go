// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.1.2

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

type PortalHTTPServer interface {
	GetUserList(context.Context, *GetUserListRequest) (*GetUserListReply, error)
	GetUserProfile(context.Context, *GetUserProfileRequest) (*UserProfileReply, error)
	Login(context.Context, *LoginRequest) (*LoginReply, error)
	Logout(context.Context, *LogoutRequest) (*LogoutReply, error)
	Register(context.Context, *RegisterRequest) (*RegisterReply, error)
	Update(context.Context, *UpdateRequest) (*UpdateReply, error)
}

func RegisterPortalHTTPServer(s *http.Server, srv PortalHTTPServer) {
	r := s.Route("/")
	r.POST("/v1/register", _Portal_Register0_HTTP_Handler(srv))
	r.POST("/v1/login", _Portal_Login0_HTTP_Handler(srv))
	r.POST("/v1/logout", _Portal_Logout0_HTTP_Handler(srv))
	r.PUT("/v1/user/update", _Portal_Update0_HTTP_Handler(srv))
	r.GET("/v1/user/profile/{uid}", _Portal_GetUserProfile0_HTTP_Handler(srv))
	r.POST("/v1/user/list", _Portal_GetUserList0_HTTP_Handler(srv))
}

func _Portal_Register0_HTTP_Handler(srv PortalHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in RegisterRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/mall.portal.v1.Portal/Register")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Register(ctx, req.(*RegisterRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*RegisterReply)
		return ctx.Result(200, reply)
	}
}

func _Portal_Login0_HTTP_Handler(srv PortalHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in LoginRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/mall.portal.v1.Portal/Login")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Login(ctx, req.(*LoginRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*LoginReply)
		return ctx.Result(200, reply)
	}
}

func _Portal_Logout0_HTTP_Handler(srv PortalHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in LogoutRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/mall.portal.v1.Portal/Logout")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Logout(ctx, req.(*LogoutRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*LogoutReply)
		return ctx.Result(200, reply)
	}
}

func _Portal_Update0_HTTP_Handler(srv PortalHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/mall.portal.v1.Portal/Update")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Update(ctx, req.(*UpdateRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateReply)
		return ctx.Result(200, reply)
	}
}

func _Portal_GetUserProfile0_HTTP_Handler(srv PortalHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetUserProfileRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/mall.portal.v1.Portal/GetUserProfile")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetUserProfile(ctx, req.(*GetUserProfileRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UserProfileReply)
		return ctx.Result(200, reply)
	}
}

func _Portal_GetUserList0_HTTP_Handler(srv PortalHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetUserListRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/mall.portal.v1.Portal/GetUserList")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetUserList(ctx, req.(*GetUserListRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetUserListReply)
		return ctx.Result(200, reply)
	}
}

type PortalHTTPClient interface {
	GetUserList(ctx context.Context, req *GetUserListRequest, opts ...http.CallOption) (rsp *GetUserListReply, err error)
	GetUserProfile(ctx context.Context, req *GetUserProfileRequest, opts ...http.CallOption) (rsp *UserProfileReply, err error)
	Login(ctx context.Context, req *LoginRequest, opts ...http.CallOption) (rsp *LoginReply, err error)
	Logout(ctx context.Context, req *LogoutRequest, opts ...http.CallOption) (rsp *LogoutReply, err error)
	Register(ctx context.Context, req *RegisterRequest, opts ...http.CallOption) (rsp *RegisterReply, err error)
	Update(ctx context.Context, req *UpdateRequest, opts ...http.CallOption) (rsp *UpdateReply, err error)
}

type PortalHTTPClientImpl struct {
	cc *http.Client
}

func NewPortalHTTPClient(client *http.Client) PortalHTTPClient {
	return &PortalHTTPClientImpl{client}
}

func (c *PortalHTTPClientImpl) GetUserList(ctx context.Context, in *GetUserListRequest, opts ...http.CallOption) (*GetUserListReply, error) {
	var out GetUserListReply
	pattern := "/v1/user/list"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/mall.portal.v1.Portal/GetUserList"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *PortalHTTPClientImpl) GetUserProfile(ctx context.Context, in *GetUserProfileRequest, opts ...http.CallOption) (*UserProfileReply, error) {
	var out UserProfileReply
	pattern := "/v1/user/profile/{uid}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/mall.portal.v1.Portal/GetUserProfile"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *PortalHTTPClientImpl) Login(ctx context.Context, in *LoginRequest, opts ...http.CallOption) (*LoginReply, error) {
	var out LoginReply
	pattern := "/v1/login"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/mall.portal.v1.Portal/Login"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *PortalHTTPClientImpl) Logout(ctx context.Context, in *LogoutRequest, opts ...http.CallOption) (*LogoutReply, error) {
	var out LogoutReply
	pattern := "/v1/logout"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/mall.portal.v1.Portal/Logout"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *PortalHTTPClientImpl) Register(ctx context.Context, in *RegisterRequest, opts ...http.CallOption) (*RegisterReply, error) {
	var out RegisterReply
	pattern := "/v1/register"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/mall.portal.v1.Portal/Register"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *PortalHTTPClientImpl) Update(ctx context.Context, in *UpdateRequest, opts ...http.CallOption) (*UpdateReply, error) {
	var out UpdateReply
	pattern := "/v1/user/update"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/mall.portal.v1.Portal/Update"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
