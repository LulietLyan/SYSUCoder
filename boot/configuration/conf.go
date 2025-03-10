package configuration

import (
	"SYSUCODER/boot/model"
	"SYSUCODER/utils"
)

type Config struct {
	Datebase  DatabaseConf `yaml:"database" json:"database"`
	Judge     JudgeConf    `yaml:"judge" json:"judge"`
	YukiImage YukiConf     `yaml:"yuki-image" json:"yuki_image"`
	NekoAcm   NekoConf     `yaml:"neko-acm" json:"neko_acm"`
	Server    ServerConf   `yaml:"server" json:"server"`
	Email     EmailConf    `yaml:"email" json:"email"`
	Token     TokenConf    `yaml:"token" json:"token"`
}

// Config 初始化
func InitConfig() error {
	// 检查全局的配置文件是否存在
	v, err := utils.IsFileExists("config.yaml")
	if err != nil {
		return err
	}

	// 不存在配置文件时需要写入默认配置
	if !v {
		Conf.Default()
		err = utils.WriteYaml(&Conf, "config.yaml")
		if err != nil {
			return err
		}
	}

	//	尝试读取配置文件
	err = utils.ReadYaml(&Conf, "config.yaml")
	if err != nil {
		return err
	}

	// 保存配置相关信息到各种变量中，需要时使用
	utils.Expire = Conf.Token.Expire
	utils.Secret = Conf.Token.Secret
	utils.Refresh = Conf.Token.Refresh
	model.DatamakeLimit = Conf.Server.DatamakeLimit
	utils.EmailHost = Conf.Email.Host
	utils.EmailPort = Conf.Email.Port
	utils.FromEmail = Conf.Email.Email
	utils.FromEmailSmtpPwd = Conf.Email.SmtpPwd
	return nil
}

// 为 Config 结构体定义 Default 方法，可以初始化配置文件及配置相关的变量
func (c *Config) Default() {
	c.Datebase.Default()
	c.Judge.Default()
	c.YukiImage.Default()
	c.NekoAcm.Default()
	c.Server.Default()
	c.Token.Default()
}
