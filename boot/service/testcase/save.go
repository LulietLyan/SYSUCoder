package testcase

import (
	"SYSUCODER/boot/model"
	"SYSUCODER/utils"
	"strconv"
	"time"
)

// 保存到json文件
func SaveJson(t model.Testcase) (string, error) {
	// 获取当前时间戳
	timestamp := time.Now().Unix()
	path := "output/testcase/" + strconv.FormatInt(timestamp, 10) + ".json"

	err := utils.WriteJson(t, path)
	if err != nil {
		return path, err
	}

	return path, nil
}
