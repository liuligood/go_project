package system_store_staff_service

import (
	"crmeb_go/internal/data/request"
	"crmeb_go/internal/data/response"
	"crmeb_go/internal/server"
)

type SystemStoreStaffServiceImpl interface {
	SystemStoreStaffList(params *request.SystemStaffListParams) (data *response.GetRealNameResp, err error)
}

type SystemStoreStaffService struct {
	svc *server.SvcContext
}

func NewSystemStoreStaffService(svc *server.SvcContext) *SystemStoreStaffService {
	return &SystemStoreStaffService{svc: svc}
}

func (s UserService) GetRealName(params *request.GetRealNameParams) (data *response.GetRealNameResp, err error) {
}
