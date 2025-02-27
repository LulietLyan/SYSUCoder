package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"strings"
	"time"
)

// 全局变量分别是 JWT 过期时间、刷新令牌的过期时间。用于签名和验证 JWT 的密钥
var (
	Expire  uint64
	Refresh uint64
	Secret  string
)

// GenerateToken 生成 token
func GenerateToken(id int64) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["uid"] = id
	claims["exp"] = time.Now().Add(time.Second * time.Duration(Expire)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(Secret))
}

// GetToken 提取 token
func GetToken(c *gin.Context) string {
	bearerToken := c.GetHeader("Authorization")

	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}
