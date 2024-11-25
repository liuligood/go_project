package config

type Conf struct {
	App    App    `json:"app"`    // 应用配置
	Server Server `json:"server"` // 服务配置
	Log    Zap    `json:"log"`    // 日志配置
}

type App struct {
	Name         string `json:"name"` // 应用名称
	Env          string `json:"env"`  // 环境
	Key          string `json:"key"`
	ServerNumber int    `json:"server_number"` // 服务器编号
}

type Server struct {
	Http Network `json:"http"` // http配置
	Rpc  Network `json:"rpc"`  // rpc配置
}

type Network struct {
	Addr string `json:"addr"` // 地址
	Mode string `json:"mode"` // 模式（etcd、直连）
}
