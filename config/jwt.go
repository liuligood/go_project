package config

type JWT struct {
	AccessSecret string `json:"access_secret"` // 密钥
	AccessExpire int64  `json:"access_expire"` // 过期时间
}
