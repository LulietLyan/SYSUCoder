package neko

import (
	"SYSUCODER/boot/model"
	"SYSUCODER/boot/service/chat"
	"SYSUCODER/boot/service/judge"
	"SYSUCODER/boot/service/misc"
	"SYSUCODER/boot/service/problem"
	"SYSUCODER/boot/service/solution"
	"SYSUCODER/boot/service/testcase"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// 对应路由组会跳转到该函数转发到 Neko 服务对应的套接字
func ForwardHandler(c *gin.Context) {
	var err error
	url := preUrl + strings.Replace(c.Request.URL.Path, "/ai", "", 1)

	log.Println("NekoACM 请求转发到: " + url)
	req, err := http.NewRequest(c.Request.Method, url, c.Request.Body)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, model.RespError("请求失败！", nil))
		return
	}

	req.Header.Set("Authorization", "Bearer "+config.Token)
	req.Header.Set("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, model.RespError("请求失败！", nil))
		return
	}
	defer res.Body.Close()

	_, err = io.Copy(c.Writer, res.Body)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, model.RespError("请求失败！", nil))
		return
	}
}

// 对话
func ChatAssistant(c *gin.Context) {
	var req model.ChatMsg

	// 参数绑定
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.RespError("参数错误", nil))
		return
	}

	p, err := chat.Assistant(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, model.RespOk("OK", p))
}

// 提交评测
func JudgeSubmit(c *gin.Context) {
	var req model.Submission

	// 参数绑定
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.RespError("参数错误", nil))
		return
	}

	j, err := judge.Submit(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, model.RespOk("OK", j))
}

// 生成笑话
func GenerateJoke(c *gin.Context) {
	p, err := misc.TellJoke()
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, model.RespOk("OK", p))
}

// 生成题目
func GenerateProblem(c *gin.Context) {
	var req model.ProblemInstruction

	// 参数绑定
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.RespError("参数错误", nil))
		return
	}

	// 生成题目
	p, err := problem.Generate(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, model.RespOk("OK", p))
}

// 翻译题目
func TranslateProblem(c *gin.Context) {
	var req model.TranslateInstruction

	// 参数绑定
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.RespError("参数错误", nil))
		return
	}

	// 翻译题目
	p, err := problem.Translate(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, model.RespOk("OK", p))
}

// 解析题目
func ParseProblem(c *gin.Context) {
	var req model.ProblemData

	// 参数绑定
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.RespError("参数错误", nil))
		return
	}

	// 解析题目
	p, err := problem.Parse(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, model.RespOk("OK", p))
}

// 生成题解
func GenerateSolution(c *gin.Context) {
	var req model.SolutionInstruction

	// 参数绑定
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.RespError("参数错误", nil))
		return
	}

	// 生成题解
	p, err := solution.Generate(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, model.RespOk("OK", p))
}

// 生成测试用例
func GenerateTestcase(c *gin.Context) {
	var req model.TestcaseInstruction

	// 参数绑定
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.RespError("参数错误", nil))
		return
	}

	// 生成测试用例
	p, err := testcase.Generate(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.RespError(err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, model.RespOk("OK", p))
}
