package configuration

// ServerConf OJ 服务器相关的配置信息
type ServerConf struct {
	Port          string `yaml:"port" json:"port"`
	DatamakeLimit uint64 `yaml:"datamake_limit" json:"datamake_limit"`
}

func (s *ServerConf) Default() {
	s.Port = "14514"
	s.DatamakeLimit = 10000
}
