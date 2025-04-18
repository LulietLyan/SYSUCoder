package routes

import (
	"SYSUCODER/server/handler"
	"SYSUCODER/server/middlewares"

	"github.com/gin-gonic/gin"
)

func InitUserRoute(ginServer *gin.Engine) {
	// Finished
	userPublicRoute := ginServer.Group("/user")
	{
		userPublicRoute.GET("/:id", handler.UserInfo)
		userPublicRoute.POST("/login", handler.UserLogin)
		userPublicRoute.POST("/register", handler.UserRegister)
		userPublicRoute.PUT("/password", handler.UserChangePassword)
	}

	// Finished
	userUserRoute := ginServer.Group("/user")
	{
		// 使用中间件
		userUserRoute.Use(middlewares.TokenAuthUser())

		userUserRoute.GET("/current", handler.UserCurrentId)
		userUserRoute.PUT("/modify/:id", handler.UserModify)
		userUserRoute.POST("/avatar/:id", handler.ModifyUserAvatar)
	}

	// Finished
	userAdminRoute := ginServer.Group("/user")
	{
		// 使用中间件
		userAdminRoute.Use(middlewares.TokenAuthAdmin())

		// Think about this
		userAdminRoute.GET("/", handler.UserList)
		userAdminRoute.POST("/", handler.UserAdd)
		userAdminRoute.PUT("/role", handler.UserModifyRole)

	}

	// Finished
	userRootRoute := ginServer.Group("/user")
	{
		// 使用中间件
		userRootRoute.Use(middlewares.TokenAuthRoot())

		userRootRoute.DELETE("/:id", handler.UserRemove)
	}
}
