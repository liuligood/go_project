package config

type Local struct {
	Path      string `json:"path"`       // 本地文件访问路径
	StorePath string `json:"store-path"` // 本地文件存储路径
}
