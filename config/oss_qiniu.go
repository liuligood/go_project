package config

type Qiniu struct {
	Zone          string `json:"zone"`            // 存储区域
	Bucket        string `json:"bucket"`          // 空间名称
	ImgPath       string `json:"img-path"`        // CDN加速域名
	AccessKey     string `json:"access-key"`      // 秘钥AK
	SecretKey     string `json:"secret-key"`      // 秘钥SK
	UseHTTPS      bool   `json:"use-https"`       // 是否使用https
	UseCdnDomains bool   `json:"use-cdn-domains"` // 上传是否使用CDN上传加速
}
