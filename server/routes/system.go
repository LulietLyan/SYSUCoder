package routes

import (
	"SYSUCODER/server/handler"
	"SYSUCODER/server/middlewares"

	"github.com/gin-gonic/gin"
)

func InitSystemRoute(ginServer *gin.Engine) {
	rootPrivateRoute := ginServer.Group("/system")
	{
		// 使用中间件
		rootPrivateRoute.Use(middlewares.TokenAuthRoot())

		rootPrivateRoute.GET("/config", handler.ConfigList)
	}
}
