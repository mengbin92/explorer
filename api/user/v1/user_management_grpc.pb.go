// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.0
// source: user/v1/user_management.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	UserManagement_GetUser_FullMethodName      = "/usermanagement.v1.UserManagement/GetUser"
	UserManagement_CreateUser_FullMethodName   = "/usermanagement.v1.UserManagement/CreateUser"
	UserManagement_UpdateUser_FullMethodName   = "/usermanagement.v1.UserManagement/UpdateUser"
	UserManagement_DeleteUser_FullMethodName   = "/usermanagement.v1.UserManagement/DeleteUser"
	UserManagement_GetApiKey_FullMethodName    = "/usermanagement.v1.UserManagement/GetApiKey"
	UserManagement_CreateApiKey_FullMethodName = "/usermanagement.v1.UserManagement/CreateApiKey"
	UserManagement_RevokeApiKey_FullMethodName = "/usermanagement.v1.UserManagement/RevokeApiKey"
)

// UserManagementClient is the client API for UserManagement service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// gRPC 服务接口
type UserManagementClient interface {
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error)
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
	UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UpdateUserResponse, error)
	DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*DeleteUserResponse, error)
	GetApiKey(ctx context.Context, in *GetApiKeyRequest, opts ...grpc.CallOption) (*GetApiKeyResponse, error)
	CreateApiKey(ctx context.Context, in *CreateApiKeyRequest, opts ...grpc.CallOption) (*CreateApiKeyResponse, error)
	RevokeApiKey(ctx context.Context, in *RevokeApiKeyRequest, opts ...grpc.CallOption) (*RevokeApiKeyResponse, error)
}

type userManagementClient struct {
	cc grpc.ClientConnInterface
}

func NewUserManagementClient(cc grpc.ClientConnInterface) UserManagementClient {
	return &userManagementClient{cc}
}

func (c *userManagementClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetUserResponse)
	err := c.cc.Invoke(ctx, UserManagement_GetUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userManagementClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, UserManagement_CreateUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userManagementClient) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UpdateUserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateUserResponse)
	err := c.cc.Invoke(ctx, UserManagement_UpdateUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userManagementClient) DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*DeleteUserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteUserResponse)
	err := c.cc.Invoke(ctx, UserManagement_DeleteUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userManagementClient) GetApiKey(ctx context.Context, in *GetApiKeyRequest, opts ...grpc.CallOption) (*GetApiKeyResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetApiKeyResponse)
	err := c.cc.Invoke(ctx, UserManagement_GetApiKey_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userManagementClient) CreateApiKey(ctx context.Context, in *CreateApiKeyRequest, opts ...grpc.CallOption) (*CreateApiKeyResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateApiKeyResponse)
	err := c.cc.Invoke(ctx, UserManagement_CreateApiKey_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userManagementClient) RevokeApiKey(ctx context.Context, in *RevokeApiKeyRequest, opts ...grpc.CallOption) (*RevokeApiKeyResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RevokeApiKeyResponse)
	err := c.cc.Invoke(ctx, UserManagement_RevokeApiKey_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserManagementServer is the server API for UserManagement service.
// All implementations must embed UnimplementedUserManagementServer
// for forward compatibility.
//
// gRPC 服务接口
type UserManagementServer interface {
	GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error)
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	UpdateUser(context.Context, *UpdateUserRequest) (*UpdateUserResponse, error)
	DeleteUser(context.Context, *DeleteUserRequest) (*DeleteUserResponse, error)
	GetApiKey(context.Context, *GetApiKeyRequest) (*GetApiKeyResponse, error)
	CreateApiKey(context.Context, *CreateApiKeyRequest) (*CreateApiKeyResponse, error)
	RevokeApiKey(context.Context, *RevokeApiKeyRequest) (*RevokeApiKeyResponse, error)
	mustEmbedUnimplementedUserManagementServer()
}

// UnimplementedUserManagementServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedUserManagementServer struct{}

func (UnimplementedUserManagementServer) GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedUserManagementServer) CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedUserManagementServer) UpdateUser(context.Context, *UpdateUserRequest) (*UpdateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (UnimplementedUserManagementServer) DeleteUser(context.Context, *DeleteUserRequest) (*DeleteUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (UnimplementedUserManagementServer) GetApiKey(context.Context, *GetApiKeyRequest) (*GetApiKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetApiKey not implemented")
}
func (UnimplementedUserManagementServer) CreateApiKey(context.Context, *CreateApiKeyRequest) (*CreateApiKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateApiKey not implemented")
}
func (UnimplementedUserManagementServer) RevokeApiKey(context.Context, *RevokeApiKeyRequest) (*RevokeApiKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RevokeApiKey not implemented")
}
func (UnimplementedUserManagementServer) mustEmbedUnimplementedUserManagementServer() {}
func (UnimplementedUserManagementServer) testEmbeddedByValue()                        {}

// UnsafeUserManagementServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserManagementServer will
// result in compilation errors.
type UnsafeUserManagementServer interface {
	mustEmbedUnimplementedUserManagementServer()
}

func RegisterUserManagementServer(s grpc.ServiceRegistrar, srv UserManagementServer) {
	// If the following call pancis, it indicates UnimplementedUserManagementServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&UserManagement_ServiceDesc, srv)
}

func _UserManagement_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserManagementServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserManagement_GetUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserManagementServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserManagement_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserManagementServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserManagement_CreateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserManagementServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserManagement_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserManagementServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserManagement_UpdateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserManagementServer).UpdateUser(ctx, req.(*UpdateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserManagement_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserManagementServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserManagement_DeleteUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserManagementServer).DeleteUser(ctx, req.(*DeleteUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserManagement_GetApiKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetApiKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserManagementServer).GetApiKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserManagement_GetApiKey_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserManagementServer).GetApiKey(ctx, req.(*GetApiKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserManagement_CreateApiKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateApiKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserManagementServer).CreateApiKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserManagement_CreateApiKey_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserManagementServer).CreateApiKey(ctx, req.(*CreateApiKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserManagement_RevokeApiKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RevokeApiKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserManagementServer).RevokeApiKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserManagement_RevokeApiKey_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserManagementServer).RevokeApiKey(ctx, req.(*RevokeApiKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserManagement_ServiceDesc is the grpc.ServiceDesc for UserManagement service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserManagement_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "usermanagement.v1.UserManagement",
	HandlerType: (*UserManagementServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUser",
			Handler:    _UserManagement_GetUser_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _UserManagement_CreateUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _UserManagement_UpdateUser_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _UserManagement_DeleteUser_Handler,
		},
		{
			MethodName: "GetApiKey",
			Handler:    _UserManagement_GetApiKey_Handler,
		},
		{
			MethodName: "CreateApiKey",
			Handler:    _UserManagement_CreateApiKey_Handler,
		},
		{
			MethodName: "RevokeApiKey",
			Handler:    _UserManagement_RevokeApiKey_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user/v1/user_management.proto",
}
