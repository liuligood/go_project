package request

import "mime/multipart"

type GetLoginPicParams struct {
	BaseServiceParams
	Temp int64 `json:"temp"`
}

type UploadParams struct {
	BaseServiceParams
	File          multipart.File
	FileHeader    *multipart.FileHeader
	FileKey       string
	PathName      string
	FileType      string
	ContentLength int64
}

type LoginParams struct {
	BaseServiceParams
	Account string `json:"account"` // 账号
	Pwd     string `json:"pwd"`     // 密码
	Key     string `json:"key"`     // 验证码key
	Code    string `json:"code"`    // 验证码
}

type LoginUserInfoParams struct {
	BaseServiceParams BaseServiceParams
	Token             string `json:"token"`
	Temp              int64  `json:"temp"`
}

type GetMenusParams struct {
	BaseServiceParams BaseServiceParams
	Temp              int64 `json:"temp"`
}
