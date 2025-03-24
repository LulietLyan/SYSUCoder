package judge

import (
	"SYSUCODER/boot/DAO"
	"SYSUCODER/boot/model"
	"SYSUCODER/tools/judge0"
	"errors"
	"log"
)

func GetStatistics() (model.Judge0Statistics, error) {
	var err error
	var stats model.Judge0Statistics

	// 统计语言数量
	stats.LanguageCount, err = DAO.CountLanguages()
	if err != nil {
		log.Println(err)
		return model.Judge0Statistics{}, errors.New("统计语言数量失败")
	}

	// 获取评测机统计信息
	stats.JudgeStatistics, err = judge0.GetStatistics()
	if err != nil {
		log.Println(err)
		return model.Judge0Statistics{}, errors.New("获取评测机统计信息失败")
	}

	// 获取评测机系统信息
	stats.JudgeSystemInfo, err = judge0.GetSystemInfo()
	if err != nil {
		log.Println(err)
		return model.Judge0Statistics{}, errors.New("获取评测机系统信息失败")
	}

	return stats, nil
}
