package biz

import (
	"context"
)

type UserManagementRepo interface {
	CreateUser(ctx context.Context, user *User) error
	GetUser(ctx context.Context, id int64) (*User, error)
	UpdateUser(ctx context.Context, user *User) error
	DeleteUser(ctx context.Context, id int64) error
	CreateApiKey(ctx context.Context, id int64, permissions string) (*ApiKey, error)
	GetApiKey(ctx context.Context, id int64) ([]*ApiKey, error)
	RevokeApiKey(ctx context.Context, id int64) error
}
