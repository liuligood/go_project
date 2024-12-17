package service_data

type GetSystemConfigParams struct {
	BaseServiceParams
	Name string `json:"name"`
}
