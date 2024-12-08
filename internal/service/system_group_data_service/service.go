package system_group_data_service

import (
	"crmeb_go/internal/server"
)

type SystemGroupDataServiceImpl interface {
}

// SystemGroupDataService 组合数据详情表 模型服务
type SystemGroupDataService struct {
	svc *server.SvcContext
}

// NewSystemGroupDataService 新组合数据详情表 模型服务实例
func NewSystemGroupDataService(svc *server.SvcContext) *SystemGroupDataService {
	return &SystemGroupDataService{svc: svc}
}
