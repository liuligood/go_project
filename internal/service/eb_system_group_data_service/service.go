package eb_system_group_data_service

import (
	"crmeb_go/internal/server"
)

type EbSystemGroupDataServiceImpl interface {
}

// EbSystemGroupDataService 组合数据详情表 模型服务
type EbSystemGroupDataService struct {
	svc *server.SvcContext
}

// NewEbSystemGroupDataService 新组合数据详情表 模型服务实例
func NewEbSystemGroupDataService(svc *server.SvcContext) *EbSystemGroupDataService {
	return &EbSystemGroupDataService{svc: svc}
}
