package request

type SystemStaffListParams struct {
	BaseServiceParams BaseServiceParams
	StoreId           int64 `json:"storeId,default=0"`
	Page              int   `json:"page"`
	Limit             int   `json:"limit"`
}
