package admin_service

import (
	"crmeb_go/define"
	"crmeb_go/internal/data/admin"
	service_data "crmeb_go/internal/data/service"
	"crmeb_go/internal/server"
	"crmeb_go/internal/service/system_config_service"
	"crmeb_go/utils/izap"
	"go.uber.org/zap"
)

type GetLoginPicImpl interface {
	GetLoginPic(params service_data.GetLoginPicParams) (data admin.GetLoginPicResp, err error)
}

type GetLoginPicService struct {
	svc *server.SvcContext
}

func NewGetLoginPicService(svc *server.SvcContext) *GetLoginPicService {
	return &GetLoginPicService{svc: svc}
}

func (a GetLoginPicService) GetLoginPic(params service_data.GetLoginPicParams) (data admin.GetLoginPicResp, err error) {
	var systemConfigParam service_data.GetSystemConfigParams
	result := make(map[string]any)
	systemConfigNameList := []string{
		define.AdminLoginBgPic,
		define.AdminSitLogoLeftTop,
		define.AdminSiteLogoLogin,
	}

	configInfoService := system_config_service.NewGetSystemConfigInfoService(a.svc)
	systemConfigParam.BaseServiceParams = params.BaseServiceParams

	for _, v := range systemConfigNameList {
		systemConfigParam.Name = v
		systemConfigInfo, err := configInfoService.GetSystemConfigInfo(systemConfigParam)
		if err != nil {
			izap.Log.Error("查询系统配置错误:", zap.String("name", v), zap.Error(err))

			return data, err
		}

		if systemConfigInfo.Name == define.AdminLoginBgPic {
			result["backgroundImage"] = systemConfigInfo.Value
		}

		if systemConfigInfo.Name == define.AdminSitLogoLeftTop {
			result["logo"] = systemConfigInfo.Value
		}

		if systemConfigInfo.Name == define.AdminSiteLogoLogin {
			result["loginLogo"] = systemConfigInfo.Value
		}
	}

	// todo 轮播图
	result["banner"] = []string{"1", "2", "3"}
	return admin.GetLoginPicResp{Map: result}, nil
}
