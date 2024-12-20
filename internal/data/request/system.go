package request

type GetSystemConfigParams struct {
	BaseServiceParams
	Name string `json:"name"`
}
