package service_data

type GetLoginPicParams struct {
	BaseServiceParams
	Temp int64 `json:"temp"`
}

const (
	CONFIG_KEY_ADMIN_LOGIN_BACKGROUND_IMAGE = "admin_login_bg_pic"
	CONFIG_KEY_ADMIN_LOGIN_LOGO_LEFT_TOP    = "site_logo_lefttop" // 登录页LOGO
	CONFIG_KEY_ADMIN_LOGIN_LOGO_LOGIN       = "site_logo_login"
)
