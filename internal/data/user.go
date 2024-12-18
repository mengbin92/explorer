package data

import (
	"time"

	"github.com/golang/protobuf/ptypes/timestamp"
	pb "github.com/mengbin92/explorer/api/user/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// 生成的枚举类型 ActivityType
type ActivityType int32

const (
	ActivityType_ACTIVITY_TYPE_UNSPECIFIED ActivityType = 0
	ActivityType_LOGIN                     ActivityType = 1
	ActivityType_API_CALL                  ActivityType = 2
	ActivityType_UPDATE            ActivityType = 3
	ActivityType_CREATE           ActivityType = 4
)

// User 表示用户表
type User struct {
	ID               int        `gorm:"primaryKey;autoIncrement" json:"id"`
	Username         string     `gorm:"size:255;unique" json:"username"`
	Email            string     `gorm:"size:255;unique" json:"email"`
	PasswordHash     string     `gorm:"size:255" json:"password_hash"`
	CreatedAt        time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt        time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
	Status           string     `gorm:"size:20" json:"status"`
	LastLoginAt      *time.Time `json:"last_login_at"`
	ApiKey           string     `gorm:"size:255" json:"api_key"`
	TwoFactorEnabled bool       `json:"two_factor_enabled"`
	Role             string     `gorm:"size:20" json:"role"`
}

func (u *User) Toproto() *pb.User {
	return &pb.User{
		Id:               int32(u.ID),
		Username:         u.Username,
		Email:            u.Email,
		PasswordHash:     u.PasswordHash,
		CreatedAt:        timestamppb.New(u.CreatedAt),
		UpdatedAt:        timestamppb.New(u.UpdatedAt),
		LastLoginAt:      timestampFromTimePtr(u.LastLoginAt),
		ApiKey:           u.ApiKey,
		TwoFactorEnabled: u.TwoFactorEnabled,
		Status:           u.Status,
		Role:             u.Role,
	}
}

// ApiKey 表示 API 密钥表
type ApiKey struct {
	ID          int       `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID      int       `gorm:"index" json:"user_id"`
	ApiKey      string    `gorm:"size:255;unique" json:"api_key"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	ExpiresAt   time.Time `json:"expires_at"`
	Permissions string    `json:"permissions"` // JSON 格式的权限数据
	Status      string    `gorm:"size:20" json:"status"`
}

func (k *ApiKey) Toproto() *pb.ApiKey {
	return &pb.ApiKey{
		Id:          int32(k.ID),
		UserId:      int32(k.UserID),
		ApiKey:      k.ApiKey,
		CreatedAt:   timestamppb.New(k.CreatedAt),
		ExpiresAt:   timestamppb.New(k.ExpiresAt),
		Permissions: k.Permissions,
		Status:      k.Status,
	}
}

// UserActivity 表示用户活动日志
type UserActivity struct {
	ID           int          `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID       int          `gorm:"index" json:"user_id"`
	ActivityType string `gorm:"size:50" json:"activity_type"`
	Details      string       `json:"details"`
	CreatedAt    time.Time    `gorm:"autoCreateTime" json:"created_at"`
}

func (a *UserActivity) TableName() string {
	return "user_activity"
}

func (a *UserActivity) Toproto() *pb.UserActivity {
	ua := &pb.UserActivity{
		Id:        int32(a.ID),
		UserId:    int32(a.UserID),
		Details:   a.Details,
		CreatedAt: timestamppb.New(a.CreatedAt),
	}
	switch a.ActivityType {
	case "login":
		ua.ActivityType = pb.ActivityType_LOGIN
	case "api_call":
		ua.ActivityType = pb.ActivityType_API_CALL
	case "update_profile":
		ua.ActivityType = pb.ActivityType_UPDATE
	case "create_project":
		ua.ActivityType = pb.ActivityType_CREATE
	default:
		ua.ActivityType = pb.ActivityType_ACTIVITY_TYPE_UNSPECIFIED
	}
	return ua
}

// 将 protobuf 的 Timestamp 转换为 Go 的 time.Time
func timeFromTimestamp(t *timestamp.Timestamp) time.Time {
	return time.Unix(t.GetSeconds(), int64(t.GetNanos()))
}

// 将 protobuf 的 Timestamp 转换为 Go 的 *time.Time
func timePtrFromTimestamp(t *timestamp.Timestamp) *time.Time {
	if t == nil {
		return nil
	}
	tm := timeFromTimestamp(t)
	return &tm
}

// 将 Go 的 time.Time 转换为 protobuf 的 Timestamp
func timestampFromTime(t time.Time) *timestamp.Timestamp {
	return &timestamp.Timestamp{
		Seconds: t.Unix(),
		Nanos:   int32(t.Nanosecond()),
	}
}

// 将 Go 的 *time.Time 转换为 protobuf 的 Timestamp
func timestampFromTimePtr(t *time.Time) *timestamp.Timestamp {
	if t == nil {
		return nil
	}
	return timestampFromTime(*t)
}
