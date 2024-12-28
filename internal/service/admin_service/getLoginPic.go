package admin_service

import (
	"crmeb_go/define"
	"crmeb_go/internal/data/request"
	"crmeb_go/internal/data/response"
	"crmeb_go/internal/server"
	"crmeb_go/internal/service/system_config_service"
	"crmeb_go/internal/service/system_group_data_service"
	"crmeb_go/utils/izap"
	"github.com/samber/lo"
	"go.uber.org/zap"
)

type GetLoginPicImpl interface {
	GetLoginPic(params *request.GetLoginPicParams) (data *response.GetLoginPicResp, err error)
}

type GetLoginPicService struct {
	svc *server.SvcContext
}

func NewGetLoginPicService(svc *server.SvcContext) *GetLoginPicService {
	return &GetLoginPicService{svc: svc}
}

func (a *GetLoginPicService) GetLoginPic(params *request.GetLoginPicParams) (*response.GetLoginPicResp, error) {
	var systemConfigParam request.GetSystemConfigParams
	var data response.GetLoginPicResp
	systemConfigNameList := []string{
		define.AdminLoginBgPic,
		define.AdminSitLogoLeftTop,
		define.AdminSiteLogoLogin,
	}

	configInfoService := system_config_service.NewGetSystemConfigInfoService(a.svc)
	systemConfigParam.BaseServiceParams = params.BaseServiceParams

	for _, v := range systemConfigNameList {
		systemConfigParam.Name = v
		systemConfigInfo, err := configInfoService.GetSystemConfigInfo(&systemConfigParam)
		if err != nil {
			izap.Log.Error("查询系统配置错误:", zap.String("name", v), zap.Error(err))

			return &data, err
		}

		if systemConfigInfo.Name == define.AdminLoginBgPic {
			data.BackgroundImage = systemConfigInfo.Value
		}

		if systemConfigInfo.Name == define.AdminSitLogoLeftTop {
			data.Logo = systemConfigInfo.Value
		}

		if systemConfigInfo.Name == define.AdminSiteLogoLogin {
			data.LoginLogo = systemConfigInfo.Value
		}
	}

	// 轮播图
	list, err := system_group_data_service.NewGetValueListService(a.svc).
		GetValueList(&request.GetGetValueListParams{BaseServiceParams: params.BaseServiceParams, Gid: define.AdminLoginBannerImageList})
	if err != nil {
		izap.Log.Error("查询轮播图失败:", zap.Error(err))

		return nil, err
	}

	if valueList, ok := list.Data[define.AdminLoginBannerImageList]; ok {
		lo.ForEach(valueList, func(item response.ValueData, index int) {
			bannerList := make([]response.Banner, 0, len(item.Fields))
			lo.ForEach(item.Fields, func(item response.Fields, index int) {
				var banner response.Banner
				banner.Pic = a.svc.Conf.PictureUrl + item.Value
				bannerList = append(bannerList, banner)
			})
			data.BannerList = bannerList
		})
	}

	return &data, nil
}
