package handler

import (
	"SYSUCODER/boot/configuration"
	"SYSUCODER/boot/model"
	"SYSUCODER/tools/judge0"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 获取设置列表
func ConfigList(c *gin.Context) {
	var err error
	config := model.Configuration{}

	config.System = *configuration.Conf
	config.Judge, err = judge0.GetConfigInfo()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, model.RespError("获取配置信息失败", nil))
		return
	}

	c.JSON(http.StatusOK, model.RespOk("OK", config))
}
