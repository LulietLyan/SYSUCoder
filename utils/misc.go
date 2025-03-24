package utils

import (
	"SYSUCODER/boot/entity"
	"encoding/json"
	"fmt"
	"math"
	"os"
	"reflect"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/exp/rand"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

// PrettyStruct 将任意结构体格式化为美观的JSON字符串，便于调试或日志输出
func PrettyStruct(data interface{}) (string, error) {
	val, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return "", err
	}
	return string(val), nil
}

// GetRandKey 生成一个6位的随机字符串，包含大小写字母和数字
func GetRandKey() string {
	rand.Seed(uint64(time.Now().UnixNano()))
	key := make([]rune, 6)
	for i := range key {
		key[i] = letters[rand.Intn(len(letters))]
	}
	return string(key)
}

// IsFileExists 检查指定路径的文件是否存在
func IsFileExists(filePath string) (bool, error) {
	_, err := os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil // 文件不存在，返回false和不为nil的error
		}
		return false, err // 其他错误，返回false和错误
	}
	return true, nil // 文件存在，返回true和nil的error
}

// GetUserInfo 从Gin上下文（如JWT中间件）中获取用户角色和ID
func GetUserInfo(c *gin.Context) (entity.Role, uint64) {
	role, exist := c.Get("role")
	if !exist {
		role = entity.RoleVisitor
	}
	id, exist := c.Get("id")
	if !exist {
		id = uint64(0)
	}

	return role.(entity.Role), id.(uint64)
}

// ConvertStringToType 将字符串转换为泛型 T 指定的类型，用于HTTP请求参数的类型转换（如将查询参数解析为整数）
func ConvertStringToType[T any](str string, result *interface{}) error {
	var tmp T
	switch any(tmp).(type) {
	case int:
		parsed, err := strconv.Atoi(str)
		if err != nil {
			return fmt.Errorf("failed to parse value for field: %w", err)
		}
		*result = parsed
	case int8:
		parsed, err := strconv.ParseInt(str, 10, 8)
		if err != nil || parsed < math.MinInt8 || parsed > math.MaxInt8 {
			return fmt.Errorf("failed to parse value for field: %w", err)
		}
		*result = parsed
	case entity.BlogStatus:
		parsed, err := strconv.ParseUint(str, 10, 8)
		if err != nil || parsed > math.MaxUint8 {
			return fmt.Errorf("failed to parse value for field: %w", err)
		}
		*result = parsed
	case int16:
		parsed, err := strconv.ParseInt(str, 10, 16)
		if err != nil || parsed < math.MinInt16 || parsed > math.MaxInt16 {
			return fmt.Errorf("failed to parse value for field: %w", err)
		}
		*result = parsed
	case int32:
		parsed, err := strconv.ParseInt(str, 10, 32)
		if err != nil || parsed < math.MinInt32 || parsed > math.MaxInt32 {
			return fmt.Errorf("failed to parse value for field: %w", err)
		}
		*result = parsed
	case int64:
		parsed, err := strconv.ParseInt(str, 10, 64)
		if err != nil {
			return fmt.Errorf("failed to parse value for field: %w", err)
		}
		*result = parsed
	case uint:
		parsed, err := strconv.ParseUint(str, 10, 0)
		if err != nil || parsed > math.MaxUint {
			return fmt.Errorf("failed to parse value for field: %w", err)
		}
		*result = parsed
	case uint8:
		parsed, err := strconv.ParseUint(str, 10, 8)
		if err != nil || parsed > math.MaxUint8 {
			return fmt.Errorf("failed to parse value for field: %w", err)
		}
		*result = parsed
	case uint16:
		parsed, err := strconv.ParseUint(str, 10, 16)
		if err != nil || parsed > math.MaxUint16 {
			return fmt.Errorf("failed to parse value for field: %w", err)
		}
		*result = parsed
	case uint32:
		parsed, err := strconv.ParseUint(str, 10, 32)
		if err != nil || parsed > math.MaxUint32 {
			return fmt.Errorf("failed to parse value for field: %w", err)
		}
		*result = parsed
	case uint64:
		parsed, err := strconv.ParseUint(str, 10, 64)
		if err != nil {
			return fmt.Errorf("failed to parse value for field: %w", err)
		}
		*result = parsed
	case float32:
		parsed, err := strconv.ParseFloat(str, 32)
		if err != nil {
			return fmt.Errorf("failed to parse value for field: %w", err)
		}
		*result = parsed
	case float64:
		parsed, err := strconv.ParseFloat(str, 64)
		if err != nil {
			return fmt.Errorf("failed to parse value for field: %w", err)
		}
		*result = parsed
	case bool:
		parsed, err := strconv.ParseBool(str)
		if err != nil {
			return fmt.Errorf("failed to parse value for field: %w", err)
		}
		*result = parsed
	case string:
		*result = str // 直接赋值字符串
	default:
		return fmt.Errorf("unsupported type: %v", reflect.TypeOf(tmp))
	}
	return nil
}
