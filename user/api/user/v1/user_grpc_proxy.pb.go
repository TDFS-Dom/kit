// Code generated by protoc-gen-go-grpc-proxy. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc-proxy v1.2.0
// - protoc             (unknown)
// source: user/api/user/v1/user.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

var _ UserServiceServer = (*userServiceClientProxy)(nil)

// userServiceClientProxy is the proxy to turn UserService client to server interface.
type userServiceClientProxy struct {
	cc UserServiceClient
}

func NewUserServiceClientProxy(cc UserServiceClient) UserServiceServer {
	return &userServiceClientProxy{cc}
}

func (c *userServiceClientProxy) ListUsers(ctx context.Context, in *ListUsersRequest) (*ListUsersResponse, error) {
	return c.cc.ListUsers(ctx, in)
}
func (c *userServiceClientProxy) GetUser(ctx context.Context, in *GetUserRequest) (*User, error) {
	return c.cc.GetUser(ctx, in)
}
func (c *userServiceClientProxy) CreateUser(ctx context.Context, in *CreateUserRequest) (*User, error) {
	return c.cc.CreateUser(ctx, in)
}
func (c *userServiceClientProxy) UpdateUser(ctx context.Context, in *UpdateUserRequest) (*User, error) {
	return c.cc.UpdateUser(ctx, in)
}
func (c *userServiceClientProxy) DeleteUser(ctx context.Context, in *DeleteUserRequest) (*DeleteUserResponse, error) {
	return c.cc.DeleteUser(ctx, in)
}
func (c *userServiceClientProxy) GetUserRoles(ctx context.Context, in *GetUserRoleRequest) (*GetUserRoleReply, error) {
	return c.cc.GetUserRoles(ctx, in)
}
func (c *userServiceClientProxy) InviteUser(ctx context.Context, in *InviteUserRequest) (*InviteUserReply, error) {
	return c.cc.InviteUser(ctx, in)
}
func (c *userServiceClientProxy) SearchUser(ctx context.Context, in *SearchUserRequest) (*SearchUserResponse, error) {
	return c.cc.SearchUser(ctx, in)
}
func (c *userServiceClientProxy) CheckUserTenant(ctx context.Context, in *CheckUserTenantRequest) (*CheckUserTenantReply, error) {
	return c.cc.CheckUserTenant(ctx, in)
}
