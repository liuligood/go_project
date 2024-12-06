package config

type AwsS3 struct {
	Bucket           string `json:"bucket"`
	Region           string `json:"region"`
	Endpoint         string `json:"endpoint" `
	SecretID         string `json:"secret-id"`
	SecretKey        string `json:"secret-key"`
	BaseURL          string `json:"base-url"`
	PathPrefix       string `json:"path-prefix"`
	S3ForcePathStyle bool   `json:"s3-force-path-style"`
	DisableSSL       bool   `json:"disable-ssl"`
}
