package system_menu_service

import (
	"crmeb_go/internal/data/admin"
	service_data "crmeb_go/internal/data/service"
	"crmeb_go/internal/server"
)

type SystemMenuServiceImpl interface {
	GetPermissionByUserId(params service_data.BaseServiceParams) (resp []admin.Permission, err error)
}

type SystemMenuService struct {
	svc *server.SvcContext
}

func NewSystemMenuService(svc *server.SvcContext) *SystemMenuService {
	return &SystemMenuService{svc: svc}
}

func (s *SystemMenuService) GetPermissionByUserId(params service_data.BaseServiceParams) (resp []admin.Permission, err error) {
	return NewPermissionByUserIdService(s.svc).GetPermissionByUserId(params)
}
