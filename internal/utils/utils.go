package utils

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
)

// GenerateAPIKey 用于生成一个新的API密钥
func GenerateAPIKey() string {
	// 使用随机数生成API密钥，或者可以选择其他生成策略
	buf := make([]byte, 32)
	_, err := rand.Read(buf)
	if err != nil {
		fmt.Println("Error generating random API key:", err)
		return ""
	}
	hash := sha256.Sum256(buf)
	return fmt.Sprintf("%x", hash)
}
