package admin_service

import (
	"crmeb_go/internal/data/admin_data"
	"crmeb_go/internal/server"
)

type AdminServiceImpl interface {
	GetValidateCode() (data admin_data.ValidateCodeData, err error)
}

type AdminService struct {
	svc *server.SvcContext
}

func NewAdminService(svc *server.SvcContext) *AdminService {
	return &AdminService{svc: svc}
}

func (a AdminService) GetValidateCode() (data admin_data.ValidateCodeData, err error) {
	return NewGetValidateCodeService(a.svc).GetValidateCode()
}
