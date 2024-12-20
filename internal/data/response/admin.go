package response

type ValidateCodeResp struct {
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

type LoginResp struct {
	Id       int64  `json:"id"`
	Account  string `json:"account"`  // 账号
	RealName string `json:"realName"` // 管理员名称
	Token    string `json:"token"`    // token
	IsSms    bool   `json:"isSms"`    // 是否接收短信
}

type LoginUserInfoResp struct {
	ID              int64    `json:"id"`       // id
	Account         string   `json:"account"`  // 账号
	RealName        string   `json:"realName"` // 管理员名称
	Roles           string   `json:"roles"`    // 权限(menus_id)
	RoleNames       string   `json:"roleNames"`
	LastIP          string   `json:"lastIp"` // 最后登录一次的id
	LastTime        string   `json:"lastTime"`
	AddTime         string   `json:"addTime"`    // 创建时间
	LoginCount      int64    `json:"loginCount"` // 登录次数
	Level           int64    `json:"level"`      // 管理员级别
	Status          bool     `json:"status"`     // 是否有效
	Phone           string   `json:"phone"`
	IsSms           bool     `json:"isSms"`
	PermissionsList []string `json:"permissionsList"` // 权限标识
	Token           string   `json:"token"`
}

type GetMenusResp struct {
	ID        int         `json:"id"`        // ID
	Pid       int         `json:"pid"`       // 父级id
	Name      string      `json:"name"`      // 名称
	Icon      string      `json:"icon"`      // icon
	Perms     interface{} `json:"perms"`     // 权限标识
	Component string      `json:"component"` // 组件路径
	MenuType  string      `json:"menuType"`  // 类型
	Sort      int         `json:"sort"`      // 排序
	ChildList *RoleTrees  `json:"childList"` // 子对象列表
}

type RoleTrees []*GetMenusResp
