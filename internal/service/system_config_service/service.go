package system_config_service

import (
	"crmeb_go/internal/data/request"
	"crmeb_go/internal/data/response"
	"crmeb_go/internal/server"
)

type SystemConfigServiceImpl interface {
	GetSystemConfigInfo(params *request.GetSystemConfigParams) (data *response.GetSystemConfigResult, err error)
}

// SystemConfigService 配置表 模型服务
type SystemConfigService struct {
	svc *server.SvcContext
}

// NewSystemConfigService 新配置表 模型服务实例
func NewSystemConfigService(svc *server.SvcContext) *SystemConfigService {
	return &SystemConfigService{svc: svc}
}

func (s SystemConfigService) GetSystemConfigInfo(params *request.GetSystemConfigParams) (data *response.GetSystemConfigResult, err error) {
	return NewGetSystemConfigInfoService(s.svc).GetSystemConfigInfo(params)
}
