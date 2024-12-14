package config

type Conf struct {
	App         App       `json:"app"`         // 应用配置
	Server      Server    `json:"server"`      // 服务配置
	Log         Zap       `json:"log"`         // 日志配置
	System      System    `json:"system"`      // 系统配置
	Mysql       Mysql     `json:"mysql"`       // mysql
	Pgsql       Pgsql     `json:"pgsql"`       // pgsql
	Cors        CORS      `json:"cors"`        // 跨域配置
	Redis       Redis     `json:"redis"`       // redis
	RedisList   []Redis   `json:"redis_list"`  // redisList
	QiniuyunOSS Qiniu     `json:"qiniu-oss"`   // 七牛云
	AliyunOSS   AliyunOSS `json:"aliyun-oss"`  // 啊里云
	JWT         JWT       `json:"jwt"`         // token
	PictureUrl  string    `json:"picture_url"` // 图片地址
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

type System struct {
	DbType        string `mapstructure:"db-type" json:"db-type" yaml:"db-type"`       // 数据库类型:mysql(默认)|sqlite|sqlserver|postgresql
	OssType       string `mapstructure:"oss-type" json:"oss-type" yaml:"oss-type"`    // Oss类型
	FileSize      int64  `mapstructure:"file-size" json:"file-size" yaml:"file-size"` // 是否同步config表数据到redis
	RouterPrefix  string `mapstructure:"router-prefix" json:"router-prefix" yaml:"router-prefix"`
	Addr          int64  `mapstructure:"addr" json:"addr" yaml:"addr"` // 端口值
	LimitCountIP  int64  `mapstructure:"iplimit-count" json:"iplimit-count" yaml:"iplimit-count"`
	LimitTimeIP   int64  `mapstructure:"iplimit-time" json:"iplimit-time" yaml:"iplimit-time"`
	UseMultipoint bool   `mapstructure:"use-multipoint" json:"use-multipoint" yaml:"use-multipoint"`    // 多点登录拦截
	UseRedis      bool   `mapstructure:"use-redis" json:"use-redis" yaml:"use-redis"`                   // 使用redis
	UseMongo      bool   `mapstructure:"use-mongo" json:"use-mongo" yaml:"use-mongo"`                   // 使用mongo
	UseStrictAuth bool   `mapstructure:"use-strict-auth" json:"use-strict-auth" yaml:"use-strict-auth"` // 使用树形角色分配模式
	AsyncConfig   bool   `mapstructure:"async-config" json:"async-config" yaml:"async-config"`          // 是否同步config表数据到redis
}
