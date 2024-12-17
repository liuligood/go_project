package admin_service

import (
	service_data "crmeb_go/internal/data/request"
	"crmeb_go/internal/data/response"
	"crmeb_go/internal/server"
)

type AdminServiceImpl interface {
	GetValidateCode() (resp response.ValidateCodeResp, err error)
	GetLoginPic(params service_data.GetLoginPicParams) (resp response.GetLoginPicResp, err error)
	UploadFile(params service_data.UploadParams) (resp response.UploadResp, err error)
	Login(params service_data.LoginParams) (resp response.LoginResp, err error)
	LoginUserInfo(params service_data.LoginUserInfoParams) (resp response.LoginUserInfoResp, err error)
}

type AdminService struct {
	svc *server.SvcContext
}

func NewAdminService(svc *server.SvcContext) *AdminService {
	return &AdminService{svc: svc}
}

func (a AdminService) GetValidateCode() (data response.ValidateCodeResp, err error) {
	return NewGetValidateCodeService(a.svc).GetValidateCode()
}

func (a AdminService) GetLoginPic(params service_data.GetLoginPicParams) (data response.GetLoginPicResp, err error) {
	return NewGetLoginPicService(a.svc).GetLoginPic(params)
}

func (a AdminService) UploadFile(params service_data.UploadParams) (data response.UploadResp, err error) {
	return NewUploadService(a.svc).UploadFile(params)
}

func (a AdminService) Login(params service_data.LoginParams) (resp response.LoginResp, err error) {
	return NewLoginService(a.svc).Login(params)
}

func (a AdminService) LoginUserInfo(params service_data.LoginUserInfoParams) (resp response.LoginUserInfoResp, err error) {
	return NewLoginUserInfoService(a.svc).LoginUserInfo(params)
}
