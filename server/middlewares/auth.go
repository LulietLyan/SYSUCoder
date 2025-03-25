package middlewares

import (
	"SYSUCODER/boot/configuration"
	"SYSUCODER/boot/entity"
	"SYSUCODER/boot/model"
	"SYSUCODER/boot/service/user"
	"SYSUCODER/utils"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func TokenGetInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 检查 token 是否存在，不存在设置为访客
		if utils.GetToken(c) == "" {
			c.Set("id", uint64(0))
			c.Set("role", entity.RoleVisitor)
			c.Next()
			return
		}

		// 从 token 解析用户 ID
		uid, err := utils.GetTokenUid(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, model.RespError("获取用户id失败", nil))
			c.Abort()
			return
		}

		// 从 token 解析用户角色
		role, err := getUserRole(uid)
		if err != nil {
			c.JSON(http.StatusUnauthorized, model.RespError("无法查询到用户组", nil))
			c.Abort()
			return
		}

		// token 自动刷新
		if role != entity.RoleVisitor {
			err = tokenAutoRefresh(c)
			if err != nil {
				log.Println(err)
			}
		}

		// 5. 将用户信息存入上下文
		c.Set("id", uid)
		c.Set("role", role)
		c.Next()
	}
}

func TokenAuthUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, _ := utils.GetUserInfo(c)
		if role < entity.RoleUser {
			c.JSON(http.StatusForbidden, model.RespError("请登录", nil))
			c.Abort()
			return
		}
		// 放行
		c.Next()
	}
}

func TokenAuthEditor() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, _ := utils.GetUserInfo(c)
		if role < entity.RoleEditor {
			c.JSON(http.StatusForbidden, model.RespError("权限不足", nil))
			c.Abort()
			return
		}
		// 放行
		c.Next()
	}
}

func TokenAuthAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, _ := utils.GetUserInfo(c)
		if role < entity.RoleAdmin {
			c.JSON(http.StatusForbidden, model.RespError("权限不足", nil))
			c.Abort()
			return
		}

		// 放行
		c.Next()
	}
}

func TokenAuthRoot() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, _ := utils.GetUserInfo(c)
		if role < entity.RoleRoot {
			c.JSON(http.StatusForbidden, model.RespError("权限不足", nil))
			c.Abort()
			return
		}

		// 放行
		c.Next()
	}
}

func tokenAutoRefresh(c *gin.Context) error {
	// 1. 获取 Token 过期时间
	config := configuration.Conf.Token
	exp, err := utils.GetTokenExpire(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, model.RespError("token无效，获取用户信息失败", nil))
		c.Abort()
		return err
	}

	// 计算 token 剩余时间
	timeLeft := exp - uint64(time.Now().Unix())
	if timeLeft > config.Refresh {
		return nil
	}

	_, id_ := utils.GetUserInfo(c)
	uid := uint64(id_)

	// 生成新 token
	token, err := utils.GenerateToken(int64(uid))
	if err != nil {
		c.JSON(http.StatusUnauthorized, model.RespError("token 刷新失败", nil))
		c.Abort()
		return err
	}

	// 4. 返回新 Token 并要求客户端重试
	c.JSON(http.StatusUnauthorized, model.RespRetry("token 已刷新，请重新发送请求", token))
	c.Abort()
	return nil
}

func getUserRole(uid uint64) (entity.Role, error) {
	// 获取用户信息
	u, err := user.SelectById(uid)
	if err != nil {
		return 0, err
	}

	return u.Role, nil
}

func tokenVerify(c *gin.Context) error {
	err := utils.VerifyToken(c)
	if err != nil {
		return errors.New("token 无效")
	}

	return nil
}
