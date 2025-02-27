package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"strconv"
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

// VetifyToken 验证 token
func VerifyToken(c *gin.Context) error {
	tokenString := GetToken(c)
	_, err := jwt.Parse(tokenString, func(t *jwt.Token) (any, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(Secret), nil
	})
	if err != nil {
		return err
	}
	return nil
}

// GetTokenUid 提取 token 中的 uid
func GetTokenUid(c *gin.Context) (uint64, error) {
	tokenString := GetToken(c)
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (any, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(Secret), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		uid, err := strconv.ParseUint(fmt.Sprintf("%.0f", claims["uid"]), 10, 64)
		if err != nil {
			return 0, err
		}
		return uid, nil
	}
	return 0, nil
}

// GetTokenExpire 提取 token 过期时间
func GetTokenExpire(c *gin.Context) (uint64, error) {
	tokenString := GetToken(c)
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (any, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(Secret), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		exp, err := strconv.ParseUint(fmt.Sprintf("%.0f", claims["exp"]), 10, 64)
		if err != nil {
			return 0, err
		}
		return exp, nil
	}
	return 0, nil
}
