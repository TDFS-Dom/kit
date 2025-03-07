// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: saas/api/tenant/v1/tenant.proto

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

// TenantServiceClient is the client API for TenantService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TenantServiceClient interface {
	//CreateTenant
	//authz: saas.tenant,*,create
	CreateTenant(ctx context.Context, in *CreateTenantRequest, opts ...grpc.CallOption) (*Tenant, error)
	//UpdateTenant
	//authz: saas.tenant,{id},update
	UpdateTenant(ctx context.Context, in *UpdateTenantRequest, opts ...grpc.CallOption) (*Tenant, error)
	//DeleteTenant
	//authz: saas.tenant,{id},delete
	DeleteTenant(ctx context.Context, in *DeleteTenantRequest, opts ...grpc.CallOption) (*DeleteTenantReply, error)
	// GetTenant
	// authz: saas.tenant,{id},get
	GetTenant(ctx context.Context, in *GetTenantRequest, opts ...grpc.CallOption) (*Tenant, error)
	// GetTenant
	// authz: saas.tenant,{id},get
	GetTenantPublic(ctx context.Context, in *GetTenantPublicRequest, opts ...grpc.CallOption) (*TenantInfo, error)
	//ListTenant
	//authz: saas.tenant,*,list
	ListTenant(ctx context.Context, in *ListTenantRequest, opts ...grpc.CallOption) (*ListTenantReply, error)
	//GetCurrentTenant
	GetCurrentTenant(ctx context.Context, in *GetCurrentTenantRequest, opts ...grpc.CallOption) (*GetCurrentTenantReply, error)
	ChangeTenant(ctx context.Context, in *ChangeTenantRequest, opts ...grpc.CallOption) (*ChangeTenantReply, error)
}

type tenantServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTenantServiceClient(cc grpc.ClientConnInterface) TenantServiceClient {
	return &tenantServiceClient{cc}
}

func (c *tenantServiceClient) CreateTenant(ctx context.Context, in *CreateTenantRequest, opts ...grpc.CallOption) (*Tenant, error) {
	out := new(Tenant)
	err := c.cc.Invoke(ctx, "/saas.api.tenant.v1.TenantService/CreateTenant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tenantServiceClient) UpdateTenant(ctx context.Context, in *UpdateTenantRequest, opts ...grpc.CallOption) (*Tenant, error) {
	out := new(Tenant)
	err := c.cc.Invoke(ctx, "/saas.api.tenant.v1.TenantService/UpdateTenant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tenantServiceClient) DeleteTenant(ctx context.Context, in *DeleteTenantRequest, opts ...grpc.CallOption) (*DeleteTenantReply, error) {
	out := new(DeleteTenantReply)
	err := c.cc.Invoke(ctx, "/saas.api.tenant.v1.TenantService/DeleteTenant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tenantServiceClient) GetTenant(ctx context.Context, in *GetTenantRequest, opts ...grpc.CallOption) (*Tenant, error) {
	out := new(Tenant)
	err := c.cc.Invoke(ctx, "/saas.api.tenant.v1.TenantService/GetTenant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tenantServiceClient) GetTenantPublic(ctx context.Context, in *GetTenantPublicRequest, opts ...grpc.CallOption) (*TenantInfo, error) {
	out := new(TenantInfo)
	err := c.cc.Invoke(ctx, "/saas.api.tenant.v1.TenantService/GetTenantPublic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tenantServiceClient) ListTenant(ctx context.Context, in *ListTenantRequest, opts ...grpc.CallOption) (*ListTenantReply, error) {
	out := new(ListTenantReply)
	err := c.cc.Invoke(ctx, "/saas.api.tenant.v1.TenantService/ListTenant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tenantServiceClient) GetCurrentTenant(ctx context.Context, in *GetCurrentTenantRequest, opts ...grpc.CallOption) (*GetCurrentTenantReply, error) {
	out := new(GetCurrentTenantReply)
	err := c.cc.Invoke(ctx, "/saas.api.tenant.v1.TenantService/GetCurrentTenant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tenantServiceClient) ChangeTenant(ctx context.Context, in *ChangeTenantRequest, opts ...grpc.CallOption) (*ChangeTenantReply, error) {
	out := new(ChangeTenantReply)
	err := c.cc.Invoke(ctx, "/saas.api.tenant.v1.TenantService/ChangeTenant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TenantServiceServer is the server API for TenantService service.
// All implementations should embed UnimplementedTenantServiceServer
// for forward compatibility
type TenantServiceServer interface {
	//CreateTenant
	//authz: saas.tenant,*,create
	CreateTenant(context.Context, *CreateTenantRequest) (*Tenant, error)
	//UpdateTenant
	//authz: saas.tenant,{id},update
	UpdateTenant(context.Context, *UpdateTenantRequest) (*Tenant, error)
	//DeleteTenant
	//authz: saas.tenant,{id},delete
	DeleteTenant(context.Context, *DeleteTenantRequest) (*DeleteTenantReply, error)
	// GetTenant
	// authz: saas.tenant,{id},get
	GetTenant(context.Context, *GetTenantRequest) (*Tenant, error)
	// GetTenant
	// authz: saas.tenant,{id},get
	GetTenantPublic(context.Context, *GetTenantPublicRequest) (*TenantInfo, error)
	//ListTenant
	//authz: saas.tenant,*,list
	ListTenant(context.Context, *ListTenantRequest) (*ListTenantReply, error)
	//GetCurrentTenant
	GetCurrentTenant(context.Context, *GetCurrentTenantRequest) (*GetCurrentTenantReply, error)
	ChangeTenant(context.Context, *ChangeTenantRequest) (*ChangeTenantReply, error)
}

// UnimplementedTenantServiceServer should be embedded to have forward compatible implementations.
type UnimplementedTenantServiceServer struct {
}

func (UnimplementedTenantServiceServer) CreateTenant(context.Context, *CreateTenantRequest) (*Tenant, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTenant not implemented")
}
func (UnimplementedTenantServiceServer) UpdateTenant(context.Context, *UpdateTenantRequest) (*Tenant, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTenant not implemented")
}
func (UnimplementedTenantServiceServer) DeleteTenant(context.Context, *DeleteTenantRequest) (*DeleteTenantReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTenant not implemented")
}
func (UnimplementedTenantServiceServer) GetTenant(context.Context, *GetTenantRequest) (*Tenant, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTenant not implemented")
}
func (UnimplementedTenantServiceServer) GetTenantPublic(context.Context, *GetTenantPublicRequest) (*TenantInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTenantPublic not implemented")
}
func (UnimplementedTenantServiceServer) ListTenant(context.Context, *ListTenantRequest) (*ListTenantReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTenant not implemented")
}
func (UnimplementedTenantServiceServer) GetCurrentTenant(context.Context, *GetCurrentTenantRequest) (*GetCurrentTenantReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCurrentTenant not implemented")
}
func (UnimplementedTenantServiceServer) ChangeTenant(context.Context, *ChangeTenantRequest) (*ChangeTenantReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeTenant not implemented")
}

// UnsafeTenantServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TenantServiceServer will
// result in compilation errors.
type UnsafeTenantServiceServer interface {
	mustEmbedUnimplementedTenantServiceServer()
}

func RegisterTenantServiceServer(s grpc.ServiceRegistrar, srv TenantServiceServer) {
	s.RegisterService(&TenantService_ServiceDesc, srv)
}

func _TenantService_CreateTenant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTenantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenantServiceServer).CreateTenant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/saas.api.tenant.v1.TenantService/CreateTenant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenantServiceServer).CreateTenant(ctx, req.(*CreateTenantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TenantService_UpdateTenant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTenantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenantServiceServer).UpdateTenant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/saas.api.tenant.v1.TenantService/UpdateTenant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenantServiceServer).UpdateTenant(ctx, req.(*UpdateTenantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TenantService_DeleteTenant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTenantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenantServiceServer).DeleteTenant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/saas.api.tenant.v1.TenantService/DeleteTenant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenantServiceServer).DeleteTenant(ctx, req.(*DeleteTenantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TenantService_GetTenant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTenantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenantServiceServer).GetTenant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/saas.api.tenant.v1.TenantService/GetTenant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenantServiceServer).GetTenant(ctx, req.(*GetTenantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TenantService_GetTenantPublic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTenantPublicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenantServiceServer).GetTenantPublic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/saas.api.tenant.v1.TenantService/GetTenantPublic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenantServiceServer).GetTenantPublic(ctx, req.(*GetTenantPublicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TenantService_ListTenant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTenantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenantServiceServer).ListTenant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/saas.api.tenant.v1.TenantService/ListTenant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenantServiceServer).ListTenant(ctx, req.(*ListTenantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TenantService_GetCurrentTenant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCurrentTenantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenantServiceServer).GetCurrentTenant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/saas.api.tenant.v1.TenantService/GetCurrentTenant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenantServiceServer).GetCurrentTenant(ctx, req.(*GetCurrentTenantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TenantService_ChangeTenant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeTenantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenantServiceServer).ChangeTenant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/saas.api.tenant.v1.TenantService/ChangeTenant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenantServiceServer).ChangeTenant(ctx, req.(*ChangeTenantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TenantService_ServiceDesc is the grpc.ServiceDesc for TenantService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TenantService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "saas.api.tenant.v1.TenantService",
	HandlerType: (*TenantServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTenant",
			Handler:    _TenantService_CreateTenant_Handler,
		},
		{
			MethodName: "UpdateTenant",
			Handler:    _TenantService_UpdateTenant_Handler,
		},
		{
			MethodName: "DeleteTenant",
			Handler:    _TenantService_DeleteTenant_Handler,
		},
		{
			MethodName: "GetTenant",
			Handler:    _TenantService_GetTenant_Handler,
		},
		{
			MethodName: "GetTenantPublic",
			Handler:    _TenantService_GetTenantPublic_Handler,
		},
		{
			MethodName: "ListTenant",
			Handler:    _TenantService_ListTenant_Handler,
		},
		{
			MethodName: "GetCurrentTenant",
			Handler:    _TenantService_GetCurrentTenant_Handler,
		},
		{
			MethodName: "ChangeTenant",
			Handler:    _TenantService_ChangeTenant_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "saas/api/tenant/v1/tenant.proto",
}

// TenantInternalServiceClient is the client API for TenantInternalService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TenantInternalServiceClient interface {
	//GetTenant internal api for remote tenant store
	GetTenant(ctx context.Context, in *GetTenantRequest, opts ...grpc.CallOption) (*Tenant, error)
}

type tenantInternalServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTenantInternalServiceClient(cc grpc.ClientConnInterface) TenantInternalServiceClient {
	return &tenantInternalServiceClient{cc}
}

func (c *tenantInternalServiceClient) GetTenant(ctx context.Context, in *GetTenantRequest, opts ...grpc.CallOption) (*Tenant, error) {
	out := new(Tenant)
	err := c.cc.Invoke(ctx, "/saas.api.tenant.v1.TenantInternalService/GetTenant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TenantInternalServiceServer is the server API for TenantInternalService service.
// All implementations should embed UnimplementedTenantInternalServiceServer
// for forward compatibility
type TenantInternalServiceServer interface {
	//GetTenant internal api for remote tenant store
	GetTenant(context.Context, *GetTenantRequest) (*Tenant, error)
}

// UnimplementedTenantInternalServiceServer should be embedded to have forward compatible implementations.
type UnimplementedTenantInternalServiceServer struct {
}

func (UnimplementedTenantInternalServiceServer) GetTenant(context.Context, *GetTenantRequest) (*Tenant, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTenant not implemented")
}

// UnsafeTenantInternalServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TenantInternalServiceServer will
// result in compilation errors.
type UnsafeTenantInternalServiceServer interface {
	mustEmbedUnimplementedTenantInternalServiceServer()
}

func RegisterTenantInternalServiceServer(s grpc.ServiceRegistrar, srv TenantInternalServiceServer) {
	s.RegisterService(&TenantInternalService_ServiceDesc, srv)
}

func _TenantInternalService_GetTenant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTenantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenantInternalServiceServer).GetTenant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/saas.api.tenant.v1.TenantInternalService/GetTenant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenantInternalServiceServer).GetTenant(ctx, req.(*GetTenantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TenantInternalService_ServiceDesc is the grpc.ServiceDesc for TenantInternalService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TenantInternalService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "saas.api.tenant.v1.TenantInternalService",
	HandlerType: (*TenantInternalServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTenant",
			Handler:    _TenantInternalService_GetTenant_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "saas/api/tenant/v1/tenant.proto",
}
