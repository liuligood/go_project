package config

type AliyunOSS struct {
	Endpoint        string `json:"endpoint"`
	AccessKeyId     string `json:"access-key-id"`
	AccessKeySecret string `json:"access-key-secret"`
	BucketName      string `json:"bucket-name"`
	BucketUrl       string `json:"bucket-url"`
	BasePath        string `json:"base-path"`
	Region          string `json:"region"`
	RoleArn         string `json:"role-arn"`
}
