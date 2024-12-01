package eb_system_config_service

import (
	service_data "crmeb_go/internal/data/sevice_data"
	"crmeb_go/internal/server"
)

type EbSystemConfigServiceImpl interface {
	GetSystemConfigInfo(params service_data.GetSystemConfigParams) (data service_data.GetSystemConfigResult, err error)
}

// EbSystemConfigService 配置表 模型服务
type EbSystemConfigService struct {
	svc *server.SvcContext
}

// NewEbSystemConfigService 新配置表 模型服务实例
func NewEbSystemConfigService(svc *server.SvcContext) *EbSystemConfigService {
	return &EbSystemConfigService{svc: svc}
}

func (s EbSystemConfigService) GetSystemConfigInfo(params service_data.GetSystemConfigParams) (data service_data.GetSystemConfigResult, err error) {
	return NewGetSystemConfigInfoService(s.svc).GetSystemConfigInfo(params)
}
