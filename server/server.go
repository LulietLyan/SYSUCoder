package server

import (
	"SYSUCODER/boot/configuration"

	"github.com/gin-gonic/gin"
)

var (
	ginServer *gin.Engine
)

func InitServer() error {
	config := configuration.Conf.Server

	// 创建 gin 实例
	ginServer = gin.Default()

	// 初始化路由
	err := InitRoute()
	if err != nil {
		return err
	}

	// 启动服务
	err = ginServer.Run(":", config.Port)
	if err != nil {
		return err
	}

	return nil
}
