package system_config_service

import (
	"crmeb_go/define"
	"crmeb_go/internal/data/request"
	"crmeb_go/internal/data/response"
	"crmeb_go/internal/model"
	"crmeb_go/internal/server"
	"crmeb_go/utils/izap"
	"errors"
	"github.com/redis/go-redis/v9"
	"github.com/samber/lo"
	"go.uber.org/zap"
)

type GetSystemConfigInfoImpl interface {
	GetSystemConfigInfo(params *request.GetSystemConfigParams) (data *response.GetSystemConfigResult, err error)
}

type GetSystemConfigInfoService struct {
	svc *server.SvcContext
}

func NewGetSystemConfigInfoService(svc *server.SvcContext) *GetSystemConfigInfoService {
	return &GetSystemConfigInfoService{svc: svc}
}

func (g *GetSystemConfigInfoService) GetSystemConfigInfo(params *request.GetSystemConfigParams) (data *response.GetSystemConfigResult, err error) {
	// 如果同步配置没有开启
	if !g.svc.Conf.System.AsyncConfig {
		systemConfig, err := g.svc.Repo.SystemConfigRepository.QueryByName(params.Ctx, params.Name)
		if err != nil {
			izap.Log.Error("EbSystemConfigRepository.QueryOne [errorm]:%v", zap.Error(err))

			return data, err
		}

		return &response.GetSystemConfigResult{Value: systemConfig.Value, Name: systemConfig.Name}, nil
	}

	// 检测redis是否为空
	exists, err := g.svc.RedisClient.Exists(params.Ctx, define.RedisConfigListKey).Result()
	if err != nil && !errors.Is(err, redis.Nil) {
		izap.Log.Error("g.svc.RedisClient.HMGet [errorm]:%v", zap.Error(err))

		return data, err
	}

	if exists != 1 {
		systemConfigs, err := g.svc.Repo.SystemConfigRepository.All(params.Ctx)
		if err != nil {
			izap.Log.Error("EbSystemConfigRepository.QueryAll [errorm]:%v", zap.Error(err))

			return data, err
		}

		lo.ForEach(systemConfigs, func(item *model.SystemConfig, index int) {
			err := g.svc.RedisClient.HMSet(params.Ctx, define.RedisConfigListKey, item.Name, item.Value).Err()
			if err != nil {
				izap.Log.Error("RedisClient.HMSet [errorm]:%v", zap.Error(err))
			}
		})
	}

	value, err := g.svc.RedisClient.HMGet(params.Ctx, define.RedisConfigListKey, params.Name).Result()
	if err != nil {
		izap.Log.Error("EbSystemConfigRepository.QueryAll [errorm]:%v", zap.Error(err))

		return data, err
	}

	// 是图片的value
	var pictureKey = []string{define.AdminSitLogoLeftTop, define.AdminLoginBgPic, define.AdminSiteLogoLogin, define.AdminSiteLogoSquare}
	if lo.Contains(pictureKey, params.Name) {
		value = []interface{}{g.svc.Conf.PictureUrl + value[0].(string)}
	}

	return &response.GetSystemConfigResult{Value: value[0].(string), Name: params.Name}, nil
}
