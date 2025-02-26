package configuration

// DatabaseConf 连接后台数据库所需的配置信息
type DatabaseConf struct {
	Host    string `yaml:"host" json:"host"`
	Port    string `yaml:"port" json:"port"`
	Name    string `yaml:"name" json:"name"`
	User    string `yaml:"user" json:"user"`
	Pwd     string `yaml:"password" json:"password"`
	MaxConn int    `yaml:"max_open_conns" json:"max_open_conns"`
	MaxIdle int    `yaml:"max_idle_conns" json:"max_idle_conns"`
}

func (d *DatabaseConf) Default() {
	d.Host = "STUOJ-db"
	d.Port = "3306"
	d.Name = "STUOJ-db"
	d.User = "STUOJ"
	d.Pwd = "STUOJ"
	d.MaxConn = 10
	d.MaxIdle = 5
}
