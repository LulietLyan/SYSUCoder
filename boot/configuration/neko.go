package configuration

// NekoConf 使用 Neko 进行 AI 生成内容的 API 相关配置
type NekoConf struct {
	Host  string `yaml:"host" json:"host"`
	Port  string `yaml:"port" json:"port"`
	Token string `yaml:"token" json:"token"`
}

func (n *NekoConf) Default() {
	n.Host = "127.0.0.1"
	n.Port = "14515"
	n.Token = "131Fa67D3C52f8d27BbDb777eC23dE3Bb68A24109Fb4c46Cd78471fAa3f9d529"
}
