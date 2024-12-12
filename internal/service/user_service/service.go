package user_service

import (
	service_data "crmeb_go/internal/data/service"
	"crmeb_go/internal/data/user"
	"crmeb_go/internal/server"
)

type UserServiceImpl interface {
	GetRealName(params service_data.GetRealNameParams) (data user.GetRealNameData, err error)
}

// UserService 用户表 模型服务
type UserService struct {
	svc *server.SvcContext
}

// NewUserService 新用户表 模型服务实例
func NewUserService(svc *server.SvcContext) *UserService {
	return &UserService{svc: svc}
}

func (s UserService) GetRealName(params service_data.GetRealNameParams) (data user.GetRealNameData, err error) {
	service := NewGetRealNameService(s.svc)
	data, err = service.GetRealName(params)

	return
}
