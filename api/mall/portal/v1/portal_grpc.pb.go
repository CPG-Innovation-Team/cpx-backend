// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// PortalClient is the client API for Portal service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PortalClient interface {
	// Register user
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterReply, error)
	// The user logs in
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginReply, error)
	// The user logs out
	Logout(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*LogoutReply, error)
	// Update user information
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateReply, error)
	// Get Single User Profile
	GetUserProfile(ctx context.Context, in *GetUserProfileRequest, opts ...grpc.CallOption) (*UserProfileReply, error)
	// List of users
	GetUserList(ctx context.Context, in *GetUserListRequest, opts ...grpc.CallOption) (*GetUserListReply, error)
	// Add Product
	AddProduct(ctx context.Context, in *AddProductRequest, opts ...grpc.CallOption) (*AddProductReply, error)
	// Update Product
	UpdateProduct(ctx context.Context, in *UpdateProductRequest, opts ...grpc.CallOption) (*UpdateProductReply, error)
	// Get Product Detail
	GetProductDetail(ctx context.Context, in *GetProductDetailRequest, opts ...grpc.CallOption) (*GetProductDetailReply, error)
	// Get Product List
	GetProductList(ctx context.Context, in *GetProductListRequest, opts ...grpc.CallOption) (*GetProductListReply, error)
}

type portalClient struct {
	cc grpc.ClientConnInterface
}

func NewPortalClient(cc grpc.ClientConnInterface) PortalClient {
	return &portalClient{cc}
}

func (c *portalClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterReply, error) {
	out := new(RegisterReply)
	err := c.cc.Invoke(ctx, "/mall.portal.v1.Portal/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *portalClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginReply, error) {
	out := new(LoginReply)
	err := c.cc.Invoke(ctx, "/mall.portal.v1.Portal/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *portalClient) Logout(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*LogoutReply, error) {
	out := new(LogoutReply)
	err := c.cc.Invoke(ctx, "/mall.portal.v1.Portal/Logout", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *portalClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateReply, error) {
	out := new(UpdateReply)
	err := c.cc.Invoke(ctx, "/mall.portal.v1.Portal/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *portalClient) GetUserProfile(ctx context.Context, in *GetUserProfileRequest, opts ...grpc.CallOption) (*UserProfileReply, error) {
	out := new(UserProfileReply)
	err := c.cc.Invoke(ctx, "/mall.portal.v1.Portal/GetUserProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *portalClient) GetUserList(ctx context.Context, in *GetUserListRequest, opts ...grpc.CallOption) (*GetUserListReply, error) {
	out := new(GetUserListReply)
	err := c.cc.Invoke(ctx, "/mall.portal.v1.Portal/GetUserList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *portalClient) AddProduct(ctx context.Context, in *AddProductRequest, opts ...grpc.CallOption) (*AddProductReply, error) {
	out := new(AddProductReply)
	err := c.cc.Invoke(ctx, "/mall.portal.v1.Portal/AddProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *portalClient) UpdateProduct(ctx context.Context, in *UpdateProductRequest, opts ...grpc.CallOption) (*UpdateProductReply, error) {
	out := new(UpdateProductReply)
	err := c.cc.Invoke(ctx, "/mall.portal.v1.Portal/UpdateProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *portalClient) GetProductDetail(ctx context.Context, in *GetProductDetailRequest, opts ...grpc.CallOption) (*GetProductDetailReply, error) {
	out := new(GetProductDetailReply)
	err := c.cc.Invoke(ctx, "/mall.portal.v1.Portal/GetProductDetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *portalClient) GetProductList(ctx context.Context, in *GetProductListRequest, opts ...grpc.CallOption) (*GetProductListReply, error) {
	out := new(GetProductListReply)
	err := c.cc.Invoke(ctx, "/mall.portal.v1.Portal/GetProductList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PortalServer is the server API for Portal service.
// All implementations must embed UnimplementedPortalServer
// for forward compatibility
type PortalServer interface {
	// Register user
	Register(context.Context, *RegisterRequest) (*RegisterReply, error)
	// The user logs in
	Login(context.Context, *LoginRequest) (*LoginReply, error)
	// The user logs out
	Logout(context.Context, *LogoutRequest) (*LogoutReply, error)
	// Update user information
	Update(context.Context, *UpdateRequest) (*UpdateReply, error)
	// Get Single User Profile
	GetUserProfile(context.Context, *GetUserProfileRequest) (*UserProfileReply, error)
	// List of users
	GetUserList(context.Context, *GetUserListRequest) (*GetUserListReply, error)
	// Add Product
	AddProduct(context.Context, *AddProductRequest) (*AddProductReply, error)
	// Update Product
	UpdateProduct(context.Context, *UpdateProductRequest) (*UpdateProductReply, error)
	// Get Product Detail
	GetProductDetail(context.Context, *GetProductDetailRequest) (*GetProductDetailReply, error)
	// Get Product List
	GetProductList(context.Context, *GetProductListRequest) (*GetProductListReply, error)
	mustEmbedUnimplementedPortalServer()
}

// UnimplementedPortalServer must be embedded to have forward compatible implementations.
type UnimplementedPortalServer struct {
}

func (UnimplementedPortalServer) Register(context.Context, *RegisterRequest) (*RegisterReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedPortalServer) Login(context.Context, *LoginRequest) (*LoginReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedPortalServer) Logout(context.Context, *LogoutRequest) (*LogoutReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Logout not implemented")
}
func (UnimplementedPortalServer) Update(context.Context, *UpdateRequest) (*UpdateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedPortalServer) GetUserProfile(context.Context, *GetUserProfileRequest) (*UserProfileReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserProfile not implemented")
}
func (UnimplementedPortalServer) GetUserList(context.Context, *GetUserListRequest) (*GetUserListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserList not implemented")
}
func (UnimplementedPortalServer) AddProduct(context.Context, *AddProductRequest) (*AddProductReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddProduct not implemented")
}
func (UnimplementedPortalServer) UpdateProduct(context.Context, *UpdateProductRequest) (*UpdateProductReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProduct not implemented")
}
func (UnimplementedPortalServer) GetProductDetail(context.Context, *GetProductDetailRequest) (*GetProductDetailReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProductDetail not implemented")
}
func (UnimplementedPortalServer) GetProductList(context.Context, *GetProductListRequest) (*GetProductListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProductList not implemented")
}
func (UnimplementedPortalServer) mustEmbedUnimplementedPortalServer() {}

// UnsafePortalServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PortalServer will
// result in compilation errors.
type UnsafePortalServer interface {
	mustEmbedUnimplementedPortalServer()
}

func RegisterPortalServer(s grpc.ServiceRegistrar, srv PortalServer) {
	s.RegisterService(&Portal_ServiceDesc, srv)
}

func _Portal_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortalServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mall.portal.v1.Portal/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortalServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Portal_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortalServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mall.portal.v1.Portal/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortalServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Portal_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogoutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortalServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mall.portal.v1.Portal/Logout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortalServer).Logout(ctx, req.(*LogoutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Portal_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortalServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mall.portal.v1.Portal/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortalServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Portal_GetUserProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortalServer).GetUserProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mall.portal.v1.Portal/GetUserProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortalServer).GetUserProfile(ctx, req.(*GetUserProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Portal_GetUserList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortalServer).GetUserList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mall.portal.v1.Portal/GetUserList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortalServer).GetUserList(ctx, req.(*GetUserListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Portal_AddProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortalServer).AddProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mall.portal.v1.Portal/AddProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortalServer).AddProduct(ctx, req.(*AddProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Portal_UpdateProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortalServer).UpdateProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mall.portal.v1.Portal/UpdateProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortalServer).UpdateProduct(ctx, req.(*UpdateProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Portal_GetProductDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProductDetailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortalServer).GetProductDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mall.portal.v1.Portal/GetProductDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortalServer).GetProductDetail(ctx, req.(*GetProductDetailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Portal_GetProductList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProductListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortalServer).GetProductList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mall.portal.v1.Portal/GetProductList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortalServer).GetProductList(ctx, req.(*GetProductListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Portal_ServiceDesc is the grpc.ServiceDesc for Portal service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Portal_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mall.portal.v1.Portal",
	HandlerType: (*PortalServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _Portal_Register_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _Portal_Login_Handler,
		},
		{
			MethodName: "Logout",
			Handler:    _Portal_Logout_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Portal_Update_Handler,
		},
		{
			MethodName: "GetUserProfile",
			Handler:    _Portal_GetUserProfile_Handler,
		},
		{
			MethodName: "GetUserList",
			Handler:    _Portal_GetUserList_Handler,
		},
		{
			MethodName: "AddProduct",
			Handler:    _Portal_AddProduct_Handler,
		},
		{
			MethodName: "UpdateProduct",
			Handler:    _Portal_UpdateProduct_Handler,
		},
		{
			MethodName: "GetProductDetail",
			Handler:    _Portal_GetProductDetail_Handler,
		},
		{
			MethodName: "GetProductList",
			Handler:    _Portal_GetProductList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/mall/portal/v1/portal.proto",
}
