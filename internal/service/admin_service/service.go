package admin_service

import (
	"crmeb_go/internal/data/admin_data"
	service_data "crmeb_go/internal/data/sevice_data"
	"crmeb_go/internal/server"
)

type AdminServiceImpl interface {
	GetValidateCode() (data admin_data.ValidateCodeData, err error)
	GetLoginPic(params service_data.GetLoginPicParams) (data admin_data.GetLoginPicResp, err error)
	Upload(params service_data.UploadParams) (data admin_data.UploadResp, err error)
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

func (a AdminService) GetLoginPic(params service_data.GetLoginPicParams) (data admin_data.GetLoginPicResp, err error) {
	return NewGetLoginPicService(a.svc).GetLoginPic(params)
}

func (a AdminService) Upload(params service_data.UploadParams) (data admin_data.UploadResp, err error) {
	return NewUploadService(a.svc).Upload(params)
}
