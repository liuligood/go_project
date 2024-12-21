package request

type SystemStaffListParams struct {
	BaseServiceParams BaseServiceParams
	StoreId           int64 `json:"storeId,default=0"`
	Page              int64 `json:"page"`
	PageSize          int64 `json:"pageSize"`
}
