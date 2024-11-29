package eb_user_service

import (
	"crmeb_go/internal/data/eb_user_data"
	service_data "crmeb_go/internal/data/sevice_data"
	"crmeb_go/internal/server"
)

type EbUserServiceImpl interface {
	GetRealNameImpl
}

// EbUserService 用户表 模型服务
type EbUserService struct {
	svc *server.SvcContext
}

// NewEbUserService 新用户表 模型服务实例
func NewEbUserService(svc *server.SvcContext) *EbUserService {
	return &EbUserService{svc: svc}
}

func (s EbUserService) GetRealName(params service_data.GetRealNameParams) (data eb_user_data.GetRealNameData, err error) {
	service := NewGetRealNameService(s.svc)
	data, err = service.GetRealName(params)

	return
}
