package validation

type GetLoginPicParams struct {
	Temp int64 `json:"temp"`
}

type LoginParam struct {
	Account string `json:"account"` // 账号
	Pwd     string `json:"pwd"`     // 密码
	Key     string `json:"key"`     // 验证码key
	Code    string `json:"code"`    // 验证码
}

type LoginUserInfoParam struct {
	Token string `json:"token"`
	Temp  int64  `json:"temp"`
}

type GetMenusParam struct {
	Temp int64 `json:"temp"`
}
