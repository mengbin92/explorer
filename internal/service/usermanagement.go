package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	pb "github.com/mengbin92/explorer/api/user/v1"
	"github.com/mengbin92/explorer/internal/biz"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type UserManagementService struct {
	user biz.UserManagementRepo
	pb.UnimplementedUserManagementServer
}

func NewUserManagementService(repo *biz.UserManagementCase, logger log.Logger) *UserManagementService {
	return &UserManagementService{
		user: repo,
	}
}

func (s *UserManagementService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	user, err := s.user.GetUser(ctx, req.UserId)
	if err != nil {
		return nil, errors.Wrap(err, "GetUser failed")
	}
	return &pb.GetUserResponse{
		User: user.Toproto(),
	}, nil
}
func (s *UserManagementService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	user := &biz.User{
		Username:     req.Username,
		PasswordHash: req.Password,
		Email:        req.Email,
		Role:         req.Role,
	}
	err := s.user.CreateUser(ctx, user)
	if err != nil {
		return nil, errors.Wrap(err, "CreateUser failed")
	}
	return &pb.CreateUserResponse{
		UserId: int64(user.ID),
		Status: user.Status,
	}, nil
}
func (s *UserManagementService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	userUpdate := &biz.User{
		ID:           int(req.UserId),
		Username:     req.Username,
		PasswordHash: req.Password,
		Email:        req.Email,
	}
	err := s.user.UpdateUser(ctx, userUpdate)
	if err != nil {
		return nil, errors.Wrap(err, "UpdateUser failed")
	}
	return &pb.UpdateUserResponse{
		Success: true,
		Message: "Update user success",
	}, nil
}
func (s *UserManagementService) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	err := s.user.DeleteUser(ctx, req.UserId)
	if err != nil {
		return nil, errors.Wrap(err, "DeleteUser failed")
	}
	return &pb.DeleteUserResponse{
		Success: true,
		Message: "Delete user success",
	}, nil
}
func (s *UserManagementService) GetApiKey(ctx context.Context, req *pb.GetApiKeyRequest) (*pb.GetApiKeyResponse, error) {
	keys, err := s.user.GetApiKey(ctx, req.UserId)
	if err != nil {
		return nil, errors.Wrap(err, "GetApiKey failed")
	}

	keysProto := make([]*pb.ApiKey, len(keys))
	for i, key := range keys {
		keysProto[i] = key.Toproto()
	}
	return &pb.GetApiKeyResponse{
		ApiKeys: keysProto,
	}, nil
}
func (s *UserManagementService) CreateApiKey(ctx context.Context, req *pb.CreateApiKeyRequest) (*pb.CreateApiKeyResponse, error) {
	key, err := s.user.CreateApiKey(ctx, req.UserId, req.Permissions)
	if err != nil {
		return nil, errors.Wrap(err, "CreateApiKey failed")
	}
	return &pb.CreateApiKeyResponse{
		ApiKey:    key.ApiKey,
		CreatedAt: timestamppb.New(key.CreatedAt),
		ExpiresAt: timestamppb.New(key.ExpiresAt),
		Status:    key.Status,
	}, nil
}
func (s *UserManagementService) RevokeApiKey(ctx context.Context, req *pb.RevokeApiKeyRequest) (*pb.RevokeApiKeyResponse, error) {
	err := s.user.RevokeApiKey(ctx, req.ApiKeyId)
	if err != nil {
		return nil, errors.Wrap(err, "RevokeApiKey failed")
	}
	return &pb.RevokeApiKeyResponse{
		Success: true,
		Message: "Revoke api key success",
	}, nil
}
