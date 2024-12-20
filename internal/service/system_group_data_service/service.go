package system_group_data_service

import (
	"crmeb_go/internal/data/request"
	"crmeb_go/internal/data/response"
	"crmeb_go/internal/server"
)

type SystemGroupDataServiceImpl interface {
	GetValueList(params *request.GetGetValueListParams) (data *response.GetGetValueListResult, err error)
}

// SystemGroupDataService 组合数据详情表 模型服务
type SystemGroupDataService struct {
	svc *server.SvcContext
}

// NewSystemGroupDataService 新组合数据详情表 模型服务实例
func NewSystemGroupDataService(svc *server.SvcContext) *SystemGroupDataService {
	return &SystemGroupDataService{svc: svc}
}

func (s *SystemGroupDataService) GetValueList(params *request.GetGetValueListParams) (data *response.GetGetValueListResult, err error) {
	return NewGetValueListService(s.svc).GetValueList(params)
}
