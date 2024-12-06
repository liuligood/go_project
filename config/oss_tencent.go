package config

type TencentCOS struct {
	Bucket     string `json:"bucket"`
	Region     string `json:"region"`
	SecretID   string `json:"secret-id"`
	SecretKey  string `json:"secret-key"`
	BaseURL    string `json:"base-url"`
	PathPrefix string `json:"path-prefix"`
}
