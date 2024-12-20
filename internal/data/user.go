package data

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/mengbin92/explorer/internal/biz"
	"github.com/mengbin92/explorer/internal/utils"
	"github.com/pkg/errors"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserManagementRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (ur *userRepo) CreateUser(ctx context.Context, user *biz.User) error {
	// create new user
	if err := ur.data.db.Create(user).Error; err != nil {
		ur.log.Error(err)
		return errors.Wrap(err, "create user failed")
	}
	if err := ur.insertActivityRecord(ctx, int(user.ID), biz.ActivityType_CREATE, "create new user"); err != nil {
		return errors.Wrap(err, "create activity record failed")
	}

	// create new api key record
	apiKey := &biz.ApiKey{
		UserID:      user.ID,
		ApiKey:      user.ApiKey,
		Permissions: "{\"read\": true}",
		CreatedAt:   user.CreatedAt,
		ExpiresAt:   time.Now().Add(time.Hour * 24 * 30), // 30天后过期,
		Status:      "active",
	}
	if err := ur.data.db.Create(apiKey).Error; err != nil {
		ur.log.Error(err)
		return errors.Wrap(err, "create api key failed")
	}
	if err := ur.insertActivityRecord(ctx, int(user.ID), biz.ActivityType_CREATE, "create new user api key"); err != nil {
		return errors.Wrap(err, "create activity record failed")
	}
	return nil
}

func (ur *userRepo) GetUser(ctx context.Context, id int64) (user *biz.User, err error) {
	if err = ur.data.db.First(&user, id).Error; err != nil {
		ur.log.Error(err)
		return nil, errors.Wrap(err, "get user failed")
	}
	return
}

func (ur *userRepo) UpdateUser(ctx context.Context, user *biz.User) error {
	if err := ur.data.db.Save(user).Error; err != nil {
		ur.log.Error(err)
		return errors.Wrap(err, "update user failed")
	}
	return nil
}
func (ur *userRepo) DeleteUser(ctx context.Context, id int64) error {
	if err := ur.data.db.Delete(&biz.User{}, id).Error; err != nil {
		ur.log.Error(err)
		return errors.Wrap(err, "delete user failed")
	}
	return nil
}
func (ur *userRepo) CreateApiKey(ctx context.Context, id int64, permissions string) (*biz.ApiKey, error) {
	apiKey := &biz.ApiKey{
		UserID:      int(id),
		Permissions: permissions,
		ExpiresAt:   time.Now().Add(time.Hour * 24 * 30), // 30天后过期,
		Status:      "active",
		ApiKey:      utils.GenerateAPIKey(),
	}
	if err := ur.data.db.Create(apiKey).Error; err != nil {
		ur.log.Error(err)
		return nil, errors.Wrap(err, "create api key failed")
	}
	return nil, nil
}
func (ur *userRepo) GetApiKey(ctx context.Context, id int64) ([]*biz.ApiKey, error) {
	var apiKeys []*biz.ApiKey
	if err := ur.data.db.Find(&apiKeys, "user_id = ?", id).Error; err != nil {
		ur.log.Error(err)
		return nil, errors.Wrap(err, "get api key failed")
	}
	return apiKeys, nil
}
func (ur *userRepo) RevokeApiKey(ctx context.Context, index int64) error {
	return nil
}

func (ur *userRepo) insertActivityRecord(ctx context.Context, user_id int, activity_type biz.ActivityType, msg string) error {
	userActivity := &biz.UserActivity{
		UserID:       user_id,
		Details:      msg,
		ActivityType: biz.SwitchActivityType(activity_type),
	}
	if err := ur.data.db.Create(userActivity).Error; err != nil {
		return errors.Wrap(err, "create activity record failed")
	}
	return nil
}
