package configuration

// YukiConf 使用 Yuki 图床服务 API 所需的配置
type YukiConf struct {
	Host  string `yaml:"host" json:"host"`
	Port  string `yaml:"port" json:"port"`
	Token string `yaml:"token" json:"token"`
}

func (y *YukiConf) Default() {
	y.Host = "STUOJ-yuki"
	y.Port = "7415"
	y.Token = ""
}
