package define

const (
	SystemGroupDataOpenStatus  int64 = 1 // 开启
	SystemGroupDataCloseStatus int64 = 2 // 关闭
)

const (
	SystemAdminRoles       string = "1"
	SystemAdminAllPerm     string = "*:*:*"
	SystemAdminStatusValid int64  = 1 // 有效
	SystemAdminIsSmsValid  int64  = 1 // 发送短信
)

const (
	UserStatusValid int64 = 1 // 正常
	UserIsPromoter  int64 = 1 // 是推广员
)

const (
	StoreOrderPaidValid int64 = 1 // 已支付
)
