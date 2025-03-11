package server

import (
	"SYSUCODER/boot/model"
	"SYSUCODER/server/middlewares"
	"SYSUCODER/server/routes"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRoute() error {
	// index
	ginServer.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, model.RespOk("SYSUCODER back end start running successfully!", nil))
	})

	// 404
	ginServer.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, model.RespError("404 Not Found", nil))
	})

	ginServer.Use(middlewares.TokenGetInfo())

	// 初始化路由
	routes.InitUserRoute(ginServer)
	routes.InitProblemRoute(ginServer)
	routes.InitTagRoute(ginServer)
	routes.InitSolutionRoute(ginServer)
	routes.InitTestcaseRoute(ginServer)
	routes.InitJudgeRoute(ginServer)
	routes.InitRecordRoute(ginServer)
	routes.InitBlogRoute(ginServer)
	routes.InitCommentRoute(ginServer)
	routes.InitAiRouter(ginServer)
	routes.InitStatisticsRoute(ginServer)
	routes.InitSystemRoute(ginServer)
	routes.InitMiscRoute(ginServer)

	return nil
}
