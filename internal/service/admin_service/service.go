package admin_service

import (
	"crmeb_go/internal/data/admin"
	service_data "crmeb_go/internal/data/service"
	"crmeb_go/internal/server"
)

type AdminServiceImpl interface {
	GetValidateCode() (data admin.ValidateCodeData, err error)
	GetLoginPic(params service_data.GetLoginPicParams) (data admin.GetLoginPicResp, err error)
	UploadFile(params service_data.UploadParams) (data admin.UploadResp, err error)
}

type AdminService struct {
	svc *server.SvcContext
}

func NewAdminService(svc *server.SvcContext) *AdminService {
	return &AdminService{svc: svc}
}

func (a AdminService) GetValidateCode() (data admin.ValidateCodeData, err error) {
	return NewGetValidateCodeService(a.svc).GetValidateCode()
}

func (a AdminService) GetLoginPic(params service_data.GetLoginPicParams) (data admin.GetLoginPicResp, err error) {
	return NewGetLoginPicService(a.svc).GetLoginPic(params)
}

func (a AdminService) UploadFile(params service_data.UploadParams) (data admin.UploadResp, err error) {
	return NewUploadService(a.svc).UploadFile(params)
}
