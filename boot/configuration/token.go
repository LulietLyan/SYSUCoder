package configuration

// TokenConf OJ 服务器所需的 Token
type TokenConf struct {
	Expire  uint64 `yaml:"expire" json:"expire"`
	Refresh uint64 `yaml:"refresh" json:"refresh"`
	Secret  string `yaml:"secret" json:"secret"`
}

func (t *TokenConf) Default() {
	t.Expire = 864000
	t.Refresh = 432000
	t.Secret = "SYSUCODER"
}
