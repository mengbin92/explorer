package service

import (
	"context"

	pb "github.com/mengbin92/explorer/api/user/v1"
)

type UserManagementService struct {
	pb.UnimplementedUserManagementServer
}

func NewUserManagementService() *UserManagementService {
	return &UserManagementService{}
}

func (s *UserManagementService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	return &pb.GetUserResponse{}, nil
}
func (s *UserManagementService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	return &pb.CreateUserResponse{}, nil
}
func (s *UserManagementService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	return &pb.UpdateUserResponse{}, nil
}
func (s *UserManagementService) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	return &pb.DeleteUserResponse{}, nil
}
func (s *UserManagementService) GetApiKey(ctx context.Context, req *pb.GetApiKeyRequest) (*pb.GetApiKeyResponse, error) {
	return &pb.GetApiKeyResponse{}, nil
}
func (s *UserManagementService) CreateApiKey(ctx context.Context, req *pb.CreateApiKeyRequest) (*pb.CreateApiKeyResponse, error) {
	return &pb.CreateApiKeyResponse{}, nil
}
func (s *UserManagementService) RevokeApiKey(ctx context.Context, req *pb.RevokeApiKeyRequest) (*pb.RevokeApiKeyResponse, error) {
	return &pb.RevokeApiKeyResponse{}, nil
}
