package config

type Minio struct {
	Endpoint        string `json:"endpoint"`
	AccessKeyId     string `json:"access-key-id"`
	AccessKeySecret string `json:"access-key-secret"`
	BucketName      string `json:"bucket-name"`
	UseSSL          bool   `json:"use-ssl"`
	BasePath        string `json:"base-path"`
	BucketUrl       string `json:"bucket-url"`
}
