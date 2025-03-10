package judge0

import (
	"SYSUCODER/boot/entity"
	"SYSUCODER/boot/model"
	"encoding/json"
)

// GetLanguage 请求 judge0 返回编程语言信息
func GetLanguage() ([]entity.Language, error) {
	bodystr, err := httpInteraction("/languages", "GET", nil)
	if err != nil {
		return nil, err
	}
	var languages []entity.Language
	err = json.Unmarshal([]byte(bodystr), &languages)
	if err != nil {
		return nil, err
	}
	return languages, nil
}

// GetConfigInfo 返回 judge0 相关的配置信息
func GetConfigInfo() (model.JudgeConfigInfo, error) {
	bodystr, err := httpInteraction("/config_info", "GET", nil)
	if err != nil {
		return model.JudgeConfigInfo{}, err
	}
	var config model.JudgeConfigInfo
	err = json.Unmarshal([]byte(bodystr), &config)
	if err != nil {
		return model.JudgeConfigInfo{}, err
	}
	return config, nil
}

// GetSystemInfo 返回 Judge0 相关的系统信息
func GetSystemInfo() (model.JudgeSystemInfo, error) {
	bodystr, err := httpInteraction("/system_info", "GET", nil)
	if err != nil {
		return model.JudgeSystemInfo{}, err
	}
	var system model.JudgeSystemInfo
	err = json.Unmarshal([]byte(bodystr), &system)
	if err != nil {
		return model.JudgeSystemInfo{}, err
	}
	return system, nil
}

// GetStatistics 返回题目提交情况等等信息
func GetStatistics() (model.JudgeStatistics, error) {
	bodystr, err := httpInteraction("/statistics", "GET", nil)
	if err != nil {
		return model.JudgeStatistics{}, err
	}
	var statistics model.JudgeStatistics
	err = json.Unmarshal([]byte(bodystr), &statistics)
	if err != nil {
		return model.JudgeStatistics{}, err
	}
	return statistics, nil
}

// GetAbout 返回版本等信息
func GetAbout() (model.JudgeAbout, error) {
	bodystr, err := httpInteraction("/about", "GET", nil)
	if err != nil {
		return model.JudgeAbout{}, err
	}
	var about model.JudgeAbout
	err = json.Unmarshal([]byte(bodystr), &about)
	if err != nil {
		return model.JudgeAbout{}, err
	}
	return about, nil
}

// GetWorkers 返回 judge0 工作状态，包括工作队列、队列大小等等
func GetWorkers() ([]model.JudgeWorker, error) {
	bodystr, err := httpInteraction("/workers", "GET", nil)
	if err != nil {
		return nil, err
	}
	var workers []model.JudgeWorker
	err = json.Unmarshal([]byte(bodystr), &workers)
	if err != nil {
		return nil, err
	}
	return workers, nil
}

// GetLicense 许可证信息
func GetLicense() (string, error) {
	bodystr, err := httpInteraction("/license", "GET", nil)
	if err != nil {
		return "", err
	}
	return bodystr, nil
}

// GetIsolate 这是啥？
func GetIsolate() (string, error) {
	bodystr, err := httpInteraction("/isolate", "GET", nil)
	if err != nil {
		return "", err
	}
	return bodystr, nil
}

// GetVersion 版本信息
func GetVersion() (string, error) {
	bodystr, err := httpInteraction("/version", "GET", nil)
	if err != nil {
		return "", err
	}
	return bodystr, nil
}
