package system_config_service

import (
	service_data "crmeb_go/internal/data/sevice_data"
	"crmeb_go/internal/server"
)

type SystemConfigServiceImpl interface {
	GetSystemConfigInfo(params service_data.GetSystemConfigParams) (data service_data.GetSystemConfigResult, err error)
}

// SystemConfigService 配置表 模型服务
type SystemConfigService struct {
	svc *server.SvcContext
}

// NewSystemConfigService 新配置表 模型服务实例
func NewSystemConfigService(svc *server.SvcContext) *SystemConfigService {
	return &SystemConfigService{svc: svc}
}

func (s SystemConfigService) GetSystemConfigInfo(params service_data.GetSystemConfigParams) (data service_data.GetSystemConfigResult, err error) {
	return NewGetSystemConfigInfoService(s.svc).GetSystemConfigInfo(params)
}
