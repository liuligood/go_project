package system_store_staff_service

import (
	"crmeb_go/internal/data/request"
	"crmeb_go/internal/data/response"
	"crmeb_go/internal/server"
)

type SystemStoreStaffServiceImpl interface {
	SystemStoreStaffList(params *request.SystemStaffListParams) (data *response.SystemStaffListResp, err error)
}

type SystemStoreStaffService struct {
	svc *server.SvcContext
}

func NewSystemStoreStaffService(svc *server.SvcContext) *SystemStoreStaffService {
	return &SystemStoreStaffService{svc: svc}
}

func (s *SystemStoreStaffService) SystemStoreStaffList(params *request.SystemStaffListParams) (data *response.SystemStaffListResp, err error) {
	return NewSystemStoreStaffListService(s.svc).SystemStoreStaffList(params)
}
