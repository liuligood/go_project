package validation

type SystemStaffListParam struct {
	StoreId int64 `form:"storeId,default=0"`
	Page    int   `form:"page"`
	Limit   int   `form:"limit"`
}
