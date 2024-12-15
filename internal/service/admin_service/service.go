package admin_service

import (
	"crmeb_go/internal/data/admin"
	service_data "crmeb_go/internal/data/service"
	"crmeb_go/internal/server"
)

type AdminServiceImpl interface {
	GetValidateCode() (resp admin.ValidateCodeResp, err error)
	GetLoginPic(params service_data.GetLoginPicParams) (resp admin.GetLoginPicResp, err error)
	UploadFile(params service_data.UploadParams) (resp admin.UploadResp, err error)
	Login(params service_data.LoginParams) (resp admin.LoginResp, err error)
	LoginUserInfo(params service_data.LoginUserInfoParams) (resp admin.LoginUserInfoResp, err error)
}

type AdminService struct {
	svc *server.SvcContext
}

func NewAdminService(svc *server.SvcContext) *AdminService {
	return &AdminService{svc: svc}
}

func (a AdminService) GetValidateCode() (data admin.ValidateCodeResp, err error) {
	return NewGetValidateCodeService(a.svc).GetValidateCode()
}

func (a AdminService) GetLoginPic(params service_data.GetLoginPicParams) (data admin.GetLoginPicResp, err error) {
	return NewGetLoginPicService(a.svc).GetLoginPic(params)
}

func (a AdminService) UploadFile(params service_data.UploadParams) (data admin.UploadResp, err error) {
	return NewUploadService(a.svc).UploadFile(params)
}

func (a AdminService) Login(params service_data.LoginParams) (resp admin.LoginResp, err error) {
	return NewLoginService(a.svc).Login(params)
}

func (a AdminService) LoginUserInfo(params service_data.LoginUserInfoParams) (resp admin.LoginUserInfoResp, err error) {
	return NewLoginUserInfoService(a.svc).LoginUserInfo(params)
}
