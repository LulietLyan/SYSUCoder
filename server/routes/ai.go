package routes

import (
	"SYSUCODER/server/middlewares"
	"SYSUCODER/tools/neko"

	"github.com/gin-gonic/gin"
)

func InitAiRouter(ginServer *gin.Engine) {
	aiUserRouter := ginServer.Group("/ai")
	{
		// 使用中间件
		aiUserRouter.Use(middlewares.TokenAuthUser())

		aiUserRouter.POST("/chat/assistant", neko.ChatAssistant)
		aiUserRouter.GET("/misc/joke", neko.GenerateJoke)
		aiUserRouter.POST("/judge/submit", neko.JudgeSubmit)
	}

	aiEditorRouter := ginServer.Group("/ai")
	{
		// 使用中间件
		aiEditorRouter.Use(middlewares.TokenAuthEditor())

		aiEditorRouter.POST("/problem/parse", neko.ParseProblem)
		aiEditorRouter.POST("/problem/translate", neko.TranslateProblem)
		aiEditorRouter.POST("/problem/generate", neko.GenerateProblem)
		aiEditorRouter.POST("/testcase/generate", neko.GenerateTestcase)
		aiEditorRouter.POST("/solution/generate", neko.GenerateSolution)
	}
}
