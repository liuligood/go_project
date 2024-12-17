package response

type Permission struct {
	Id   int64  `json:"id"`
	Pid  int64  `json:"pid"`
	Name string `json:"name"`
	Path string `json:"path"`
	Sort int64  `json:"sort"`
}
