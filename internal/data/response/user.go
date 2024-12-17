package response

type GetRealNameResp struct {
	UserId   uint64 `json:"user_id"`
	RealName string `json:"real_name"`
}
