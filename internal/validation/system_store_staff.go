package validation

type SystemStaffListParam struct {
	StoreId  int64 `form:"storeId,default=0"`
	Page     int64 `form:"page"`
	PageSize int64 `form:"pageSize"`
}
