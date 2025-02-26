package configuration

// JudgeConf 自动判题系统 API 配置信息
type JudgeConf struct {
	Host  string `yaml:"host" json:"host"`
	Port  string `yaml:"port" json:"port"`
	Token string `yaml:"token" json:"token"`
}

func (j *JudgeConf) Default() {
	j.Host = "judge0"
	j.Port = "2358"
	j.Token = ""
}
