package config

type Qiniu struct {
	Zone      string `json:"zone"`       // 存储区域
	Bucket    string `json:"bucket"`     // 空间名称
	AccessKey string `json:"access-key"` // 秘钥AK
	SecretKey string `json:"secret-key"` // 秘钥SK
}
