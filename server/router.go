package server

import (
	"SYSUCODER/boot/model"
	"SYSUCODER/server/middlewares"
	"SYSUCODER/server/routes"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRoute() error {
	// 访问根目录时返回 200
	ginServer.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, model.RespOk("SYSUCODER back end start running successfully!", nil))
	})

	// 当请求报文的路由不存在时返回 404
	ginServer.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, model.RespError("404 Not Found", nil))
	})

	// 注册一个全局中间件 TokenGetInfo，所有后续定义的路由都会经过此中间件
	// 注意：前面注册的路由不会经过此中间件
	ginServer.Use(middlewares.TokenGetInfo())

	// 模块化路由
	// Finished
	routes.InitUserRoute(ginServer)
	// TODO
	routes.InitProblemRoute(ginServer)
	// TODO
	routes.InitTagRoute(ginServer)
	// TODO
	routes.InitSolutionRoute(ginServer)
	// TODO
	routes.InitTestcaseRoute(ginServer)
	// TODO
	routes.InitJudgeRoute(ginServer)
	// TODO
	routes.InitRecordRoute(ginServer)
	// Finished
	routes.InitBlogRoute(ginServer)
	// Finished
	routes.InitCommentRoute(ginServer)
	// Finished but need to read twice
	routes.InitAiRouter(ginServer)
	// Finished
	routes.InitStatisticsRoute(ginServer)
	// Finished
	routes.InitSystemRoute(ginServer)
	// Finished
	routes.InitMiscRoute(ginServer)

	return nil
}
