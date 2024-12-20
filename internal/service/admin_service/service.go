package admin_service

import (
	"crmeb_go/internal/data/request"
	"crmeb_go/internal/data/response"
	"crmeb_go/internal/server"
)

type AdminServiceImpl interface {
	GetValidateCode() (resp response.ValidateCodeResp, err error)
	GetLoginPic(params request.GetLoginPicParams) (resp response.GetLoginPicResp, err error)
	UploadFile(params request.UploadParams) (resp response.UploadResp, err error)
	Login(params request.LoginParams) (resp response.LoginResp, err error)
	LoginUserInfo(params request.LoginUserInfoParams) (resp response.LoginUserInfoResp, err error)
	GetMenus(params request.GetMenusParams) (resp []*response.GetMenusResp, err error)
}

type AdminService struct {
	svc *server.SvcContext
}

func NewAdminService(svc *server.SvcContext) *AdminService {
	return &AdminService{svc: svc}
}

func (a *AdminService) GetValidateCode() (data response.ValidateCodeResp, err error) {
	return NewGetValidateCodeService(a.svc).GetValidateCode()
}

func (a *AdminService) GetLoginPic(params request.GetLoginPicParams) (data response.GetLoginPicResp, err error) {
	return NewGetLoginPicService(a.svc).GetLoginPic(params)
}

func (a *AdminService) UploadFile(params request.UploadParams) (data response.UploadResp, err error) {
	return NewUploadService(a.svc).UploadFile(params)
}

func (a *AdminService) Login(params request.LoginParams) (resp response.LoginResp, err error) {
	return NewLoginService(a.svc).Login(params)
}

func (a *AdminService) LoginUserInfo(params request.LoginUserInfoParams) (resp response.LoginUserInfoResp, err error) {
	return NewLoginUserInfoService(a.svc).LoginUserInfo(params)
}

func (a *AdminService) GetMenus(params request.GetMenusParams) (resp []*response.GetMenusResp, err error) {
	return NewGetMenusService(a.svc).GetMenus(params)
}
