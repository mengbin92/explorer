// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.8.2
// - protoc             v5.29.0
// source: user/v1/user_management.proto

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

const OperationUserManagementCreateApiKey = "/usermanagement.v1.UserManagement/CreateApiKey"
const OperationUserManagementCreateUser = "/usermanagement.v1.UserManagement/CreateUser"
const OperationUserManagementDeleteUser = "/usermanagement.v1.UserManagement/DeleteUser"
const OperationUserManagementGetApiKey = "/usermanagement.v1.UserManagement/GetApiKey"
const OperationUserManagementGetUser = "/usermanagement.v1.UserManagement/GetUser"
const OperationUserManagementRevokeApiKey = "/usermanagement.v1.UserManagement/RevokeApiKey"
const OperationUserManagementUpdateUser = "/usermanagement.v1.UserManagement/UpdateUser"

type UserManagementHTTPServer interface {
	CreateApiKey(context.Context, *CreateApiKeyRequest) (*CreateApiKeyResponse, error)
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	DeleteUser(context.Context, *DeleteUserRequest) (*DeleteUserResponse, error)
	GetApiKey(context.Context, *GetApiKeyRequest) (*GetApiKeyResponse, error)
	GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error)
	RevokeApiKey(context.Context, *RevokeApiKeyRequest) (*RevokeApiKeyResponse, error)
	UpdateUser(context.Context, *UpdateUserRequest) (*UpdateUserResponse, error)
}

func RegisterUserManagementHTTPServer(s *http.Server, srv UserManagementHTTPServer) {
	r := s.Route("/")
	r.GET("/user/v1/{user_id}", _UserManagement_GetUser0_HTTP_Handler(srv))
	r.POST("user/v1/create", _UserManagement_CreateUser0_HTTP_Handler(srv))
	r.PUT("user/v1/update", _UserManagement_UpdateUser0_HTTP_Handler(srv))
	r.DELETE("user/v1/delete/{user_id}", _UserManagement_DeleteUser0_HTTP_Handler(srv))
	r.GET("/user/v1/key/{user_id}", _UserManagement_GetApiKey0_HTTP_Handler(srv))
	r.POST("user/v1/key/create", _UserManagement_CreateApiKey0_HTTP_Handler(srv))
	r.PUT("user/v1/key/revoke/{api_key_id}", _UserManagement_RevokeApiKey0_HTTP_Handler(srv))
}

func _UserManagement_GetUser0_HTTP_Handler(srv UserManagementHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetUserRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserManagementGetUser)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetUser(ctx, req.(*GetUserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetUserResponse)
		return ctx.Result(200, reply)
	}
}

func _UserManagement_CreateUser0_HTTP_Handler(srv UserManagementHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateUserRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserManagementCreateUser)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateUser(ctx, req.(*CreateUserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateUserResponse)
		return ctx.Result(200, reply)
	}
}

func _UserManagement_UpdateUser0_HTTP_Handler(srv UserManagementHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateUserRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserManagementUpdateUser)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateUser(ctx, req.(*UpdateUserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateUserResponse)
		return ctx.Result(200, reply)
	}
}

func _UserManagement_DeleteUser0_HTTP_Handler(srv UserManagementHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteUserRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserManagementDeleteUser)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteUser(ctx, req.(*DeleteUserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteUserResponse)
		return ctx.Result(200, reply)
	}
}

func _UserManagement_GetApiKey0_HTTP_Handler(srv UserManagementHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetApiKeyRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserManagementGetApiKey)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetApiKey(ctx, req.(*GetApiKeyRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetApiKeyResponse)
		return ctx.Result(200, reply)
	}
}

func _UserManagement_CreateApiKey0_HTTP_Handler(srv UserManagementHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateApiKeyRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserManagementCreateApiKey)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateApiKey(ctx, req.(*CreateApiKeyRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateApiKeyResponse)
		return ctx.Result(200, reply)
	}
}

func _UserManagement_RevokeApiKey0_HTTP_Handler(srv UserManagementHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in RevokeApiKeyRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserManagementRevokeApiKey)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.RevokeApiKey(ctx, req.(*RevokeApiKeyRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*RevokeApiKeyResponse)
		return ctx.Result(200, reply)
	}
}

type UserManagementHTTPClient interface {
	CreateApiKey(ctx context.Context, req *CreateApiKeyRequest, opts ...http.CallOption) (rsp *CreateApiKeyResponse, err error)
	CreateUser(ctx context.Context, req *CreateUserRequest, opts ...http.CallOption) (rsp *CreateUserResponse, err error)
	DeleteUser(ctx context.Context, req *DeleteUserRequest, opts ...http.CallOption) (rsp *DeleteUserResponse, err error)
	GetApiKey(ctx context.Context, req *GetApiKeyRequest, opts ...http.CallOption) (rsp *GetApiKeyResponse, err error)
	GetUser(ctx context.Context, req *GetUserRequest, opts ...http.CallOption) (rsp *GetUserResponse, err error)
	RevokeApiKey(ctx context.Context, req *RevokeApiKeyRequest, opts ...http.CallOption) (rsp *RevokeApiKeyResponse, err error)
	UpdateUser(ctx context.Context, req *UpdateUserRequest, opts ...http.CallOption) (rsp *UpdateUserResponse, err error)
}

type UserManagementHTTPClientImpl struct {
	cc *http.Client
}

func NewUserManagementHTTPClient(client *http.Client) UserManagementHTTPClient {
	return &UserManagementHTTPClientImpl{client}
}

func (c *UserManagementHTTPClientImpl) CreateApiKey(ctx context.Context, in *CreateApiKeyRequest, opts ...http.CallOption) (*CreateApiKeyResponse, error) {
	var out CreateApiKeyResponse
	pattern := "user/v1/key/create"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserManagementCreateApiKey))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *UserManagementHTTPClientImpl) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...http.CallOption) (*CreateUserResponse, error) {
	var out CreateUserResponse
	pattern := "user/v1/create"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserManagementCreateUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *UserManagementHTTPClientImpl) DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...http.CallOption) (*DeleteUserResponse, error) {
	var out DeleteUserResponse
	pattern := "user/v1/delete/{user_id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationUserManagementDeleteUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *UserManagementHTTPClientImpl) GetApiKey(ctx context.Context, in *GetApiKeyRequest, opts ...http.CallOption) (*GetApiKeyResponse, error) {
	var out GetApiKeyResponse
	pattern := "/user/v1/key/{user_id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationUserManagementGetApiKey))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *UserManagementHTTPClientImpl) GetUser(ctx context.Context, in *GetUserRequest, opts ...http.CallOption) (*GetUserResponse, error) {
	var out GetUserResponse
	pattern := "/user/v1/{user_id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationUserManagementGetUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *UserManagementHTTPClientImpl) RevokeApiKey(ctx context.Context, in *RevokeApiKeyRequest, opts ...http.CallOption) (*RevokeApiKeyResponse, error) {
	var out RevokeApiKeyResponse
	pattern := "user/v1/key/revoke/{api_key_id}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserManagementRevokeApiKey))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *UserManagementHTTPClientImpl) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...http.CallOption) (*UpdateUserResponse, error) {
	var out UpdateUserResponse
	pattern := "user/v1/update"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserManagementUpdateUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
