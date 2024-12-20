package data

import (
	"fmt"
	"testing"
	"time"

	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/mengbin92/explorer/internal/biz"
	"github.com/mengbin92/explorer/internal/conf"
	"github.com/stretchr/testify/assert"
)

func loadConfig() *conf.Bootstrap {
	c := config.New(
		config.WithSource(file.NewSource("../../configs/config.yaml")),
	)
	if err := c.Load(); err != nil {
		panic(err)
	}

	bc := &conf.Bootstrap{}
	if err := c.Scan(bc); err != nil {
		panic(err)
	}
	return bc
}

func TestLoadConfig(t *testing.T) {
	bc := loadConfig()
	assert.NotNil(t, bc)

	t.Logf("database driver: %s, source: %s", bc.Data.Database.Driver, bc.Data.Database.Source)
}

func TestNewData(t *testing.T) {
	bc := loadConfig()
	assert.NotNil(t, bc)

	data, _, err := NewData(bc.Data, nil)
	assert.Nil(t, err)
	assert.NotNil(t, data)
}

func TestCreateUser(t *testing.T) {
	bc := loadConfig()
	assert.NotNil(t, bc)

	data, _, err := NewData(bc.Data, nil)
	assert.Nil(t, err)
	assert.NotNil(t, data)

	// 创建一个新用户
	now := time.Now()
	newUser := biz.User{
		Username:         "john_doe0878",
		Email:            "john77@example.com",
		PasswordHash:     "hashedpassword123", // 在实际使用中，密码应该经过哈希处理
		LastLoginAt:      &now,
		ApiKey:           "new-api-key", // 在这里，用户本身的APIKey可以暂时存储，但是我们会另外生成API密钥插入api_keys表
		TwoFactorEnabled: false,
		Status:           "active",
		Role:             "user",
	}
	// 插入新用户
	result := data.db.Create(&newUser)
	if result.Error != nil {
		fmt.Println("Error creating user:", result.Error)
	} else {
		fmt.Println("User created successfully!")
	}

	// 创建 API 密钥记录
	apiKey := biz.ApiKey{
		UserID:    newUser.ID,
		ApiKey:    "new-api-key",
		CreatedAt: time.Now(),
		ExpiresAt: time.Now().Add(time.Hour * 24 * 30), // 30天后过期
		Permissions: "{\"read\": true}", // API 密钥的权限为 "read" 和 "write"
		Status:    "active", // API 密钥的状态为 "active"
	}

	// 将 API 密钥插入到 api_keys 表
	apiKeyResult := data.db.Create(&apiKey)
	if apiKeyResult.Error != nil {
		fmt.Println("Error creating API key:", apiKeyResult.Error)
	} else {
		fmt.Println("API key created successfully!")
	}

	// 创建用户活动记录
	userActivity := biz.UserActivity{
		UserID:       newUser.ID,
		ActivityType: "create_project", // 活动类型为用户创建
		Details:      "User created through API",
		CreatedAt:    time.Now(),
	}

	// 插入用户活动记录
	activityResult := data.db.Create(&userActivity)
	if activityResult.Error != nil {
		fmt.Println("Error creating user activity:", activityResult.Error)
	} else {
		fmt.Println("User activity created successfully!")
	}
}

func TestLoadUser(t *testing.T) {
	bc := loadConfig()
	assert.NotNil(t, bc)

	data, _, err := NewData(bc.Data, nil)
	assert.Nil(t, err)
	assert.NotNil(t, data)

	var users []biz.User
	if err := data.db.Model(&biz.User{}).Find(&users).Error; err != nil {
		t.Error(err)
	}
	for _, user := range users {
		t.Logf("user proto : %s\n", user.Toproto().String())
	}
}

func TestLoadApikey(t *testing.T) {
	bc := loadConfig()
	assert.NotNil(t, bc)

	data, _, err := NewData(bc.Data, nil)
	assert.Nil(t, err)
	assert.NotNil(t, data)

	var apikeys []biz.ApiKey
	if err := data.db.Model(&biz.ApiKey{}).Find(&apikeys).Error; err != nil {
		t.Error(err)
	}
	for _, apikey := range apikeys {
		t.Logf("apikey proto : %s\n", apikey.Toproto().String())
	}
}

func TestLoadUserActivity(t *testing.T) {
	bc := loadConfig()
	assert.NotNil(t, bc)

	data, _, err := NewData(bc.Data, nil)
	assert.Nil(t, err)
	assert.NotNil(t, data)

	var userActivities []biz.UserActivity
	if err := data.db.Model(&biz.UserActivity{}).Find(&userActivities).Error; err != nil {
		t.Error(err)
	}
	for _, userActivity := range userActivities {
		t.Logf("userActivity proto : %s\n", userActivity.Toproto().String())
	}
}
