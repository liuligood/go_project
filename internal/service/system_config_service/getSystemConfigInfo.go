package system_config_service

import (
	"crmeb_go/define"
	service_data "crmeb_go/internal/data/service"
	"crmeb_go/internal/model"
	"crmeb_go/internal/server"
	"crmeb_go/utils/izap"
	"errors"
	"github.com/redis/go-redis/v9"
	"github.com/samber/lo"
	"go.uber.org/zap"
)

type GetSystemConfigInfoImpl interface {
	GetSystemConfigInfo(params service_data.GetSystemConfigParams) (data service_data.GetSystemConfigResult, err error)
}

type GetSystemConfigInfoService struct {
	svc *server.SvcContext
}

func NewGetSystemConfigInfoService(svc *server.SvcContext) *GetSystemConfigInfoService {
	return &GetSystemConfigInfoService{svc: svc}
}

func (g *GetSystemConfigInfoService) GetSystemConfigInfo(params service_data.GetSystemConfigParams) (data service_data.GetSystemConfigResult, err error) {
	// 如果同步配置没有开启
	if !g.svc.Conf.System.AsyncConfig {
		systemConfig, err := g.svc.Repo.SystemConfigRepository.QueryByName(params.Ctx, params.Name)
		if err != nil {
			izap.Log.Error("EbSystemConfigRepository.QueryOne [http_err]:%v", zap.Error(err))

			return data, err
		}

		return service_data.GetSystemConfigResult{Value: systemConfig.Value, Name: systemConfig.Name}, nil
	}

	// 检测redis是否为空
	exists, err := g.svc.RedisClient.Exists(params.Ctx, define.RedisConfigListKey).Result()
	if err != nil && !errors.Is(err, redis.Nil) {
		izap.Log.Error("g.svc.RedisClient.HMGet [http_err]:%v", zap.Error(err))

		return data, err
	}

	if exists != 1 {
		systemConfigs, err := g.svc.Repo.SystemConfigRepository.All(params.Ctx)
		if err != nil {
			izap.Log.Error("EbSystemConfigRepository.QueryAll [http_err]:%v", zap.Error(err))

			return data, err
		}

		lo.ForEach(systemConfigs, func(item *model.SystemConfig, index int) {
			err := g.svc.RedisClient.HMSet(params.Ctx, define.RedisConfigListKey, item.Name, item.Value).Err()
			if err != nil {
				izap.Log.Error("RedisClient.HMSet [http_err]:%v", zap.Error(err))
			}
		})
	}

	value, err := g.svc.RedisClient.HMGet(params.Ctx, define.RedisConfigListKey, params.Name).Result()
	if err != nil {
		izap.Log.Error("EbSystemConfigRepository.QueryAll [http_err]:%v", zap.Error(err))

		return data, err
	}

	data.Value = value[0].(string)
	data.Name = params.Name
	return data, nil
}
