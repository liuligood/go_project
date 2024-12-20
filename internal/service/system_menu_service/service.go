package system_menu_service

import (
	"crmeb_go/internal/data/request"
	"crmeb_go/internal/data/response"
	"crmeb_go/internal/server"
)

type SystemMenuServiceImpl interface {
	GetPermissionByUserId(params request.BaseServiceParams) (resp []response.Permission, err error)
}

type SystemMenuService struct {
	svc *server.SvcContext
}

func NewSystemMenuService(svc *server.SvcContext) *SystemMenuService {
	return &SystemMenuService{svc: svc}
}

func (s *SystemMenuService) GetPermissionByUserId(params request.BaseServiceParams) (resp []response.Permission, err error) {
	return NewPermissionByUserIdService(s.svc).GetPermissionByUserId(params)
}
