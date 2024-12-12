package admin

type ValidateCodeData struct {
	Key  string `json:"key"`
	Code string `json:"code"`
}

type GetLoginPicResp struct {
	Map map[string]any `json:"map"`
}

type UploadResp struct {
	Url string `json:"url"`
}
