package admin

type ValidateCodeData struct {
	Key  string `json:"key"`
	Code string `json:"code"`
}

type GetLoginPicResp struct {
	BackgroundImage string   `json:"backgroundImage"`
	LoginLogo       string   `json:"loginLogo"`
	Logo            string   `json:"logo"`
	BannerList      []Banner `json:"banner"`
}

type Banner struct {
	Pic string `json:"pic"`
}

type UploadResp struct {
	Url string `json:"url"`
}
