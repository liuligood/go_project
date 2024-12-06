package config

type CloudflareR2 struct {
	Bucket          string `json:"bucket"`
	BaseURL         string `json:"base-url"`
	Path            string `json:"path"`
	AccountID       string `json:"account-id"`
	AccessKeyID     string `json:"access-key-id"`
	SecretAccessKey string `json:"secret-access-key"`
}
