package admin_service

import (
	"crmeb_go/define"
	"crmeb_go/internal/data/admin_data"
	service_data "crmeb_go/internal/data/sevice_data"
	"crmeb_go/internal/server"
	"crmeb_go/internal/service/system_config_service"
	"crmeb_go/utils/izap"
	"go.uber.org/zap"
)

type GetLoginPicImpl interface {
	GetLoginPic(params service_data.GetLoginPicParams) (data admin_data.GetLoginPicResp, err error)
}

type GetLoginPicService struct {
	svc *server.SvcContext
}

func NewGetLoginPicService(svc *server.SvcContext) *GetLoginPicService {
	return &GetLoginPicService{svc: svc}
}

func (a GetLoginPicService) GetLoginPic(params service_data.GetLoginPicParams) (data admin_data.GetLoginPicResp, err error) {
	result := make(map[string]any)
	var systemConfigParam service_data.GetSystemConfigParams

	//背景图
	systemConfigParam.BaseServiceParams = params.BaseServiceParams
	systemConfigParam.Name = define.CONFIG_KEY_ADMIN_LOGIN_BACKGROUND_IMAGE
	backgroundImage, err := system_config_service.NewGetSystemConfigInfoService(a.svc).GetSystemConfigInfo(systemConfigParam)

	if err != nil {
		izap.Log.Error("查询系统配置错误:", zap.String("name", define.CONFIG_KEY_ADMIN_LOGIN_BACKGROUND_IMAGE), zap.Error(err))

		return data, err
	}

	result["backgroundImage"] = backgroundImage.Value

	// logo
	systemConfigParam.Name = define.CONFIG_KEY_ADMIN_LOGIN_LOGO_LEFT_TOP
	logo, err := system_config_service.NewGetSystemConfigInfoService(a.svc).GetSystemConfigInfo(systemConfigParam)

	if err != nil {
		izap.Log.Error("查询系统配置错误:", zap.String("name", define.CONFIG_KEY_ADMIN_LOGIN_LOGO_LEFT_TOP), zap.Error(err))

		return data, err
	}

	result["logo"] = logo.Value

	// loginLogo
	systemConfigParam.Name = define.CONFIG_KEY_ADMIN_LOGIN_LOGO_LOGIN
	loginLogo, err := system_config_service.NewGetSystemConfigInfoService(a.svc).GetSystemConfigInfo(systemConfigParam)

	if err != nil {
		izap.Log.Error("查询系统配置错误:", zap.String("name", define.CONFIG_KEY_ADMIN_LOGIN_LOGO_LOGIN), zap.Error(err))

		return data, err
	}

	result["loginLogo"] = loginLogo.Value

	// todo 轮播图
	result["banner"] = []string{"1", "2", "3"}
	data.Map = result

	return data, nil
}
