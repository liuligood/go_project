package user_service

import (
	"crmeb_go/internal/data/request"
	"crmeb_go/internal/data/response"
	"crmeb_go/internal/server"
)

type UserServiceImpl interface {
	GetRealName(params *request.GetRealNameParams) (data *response.GetRealNameResp, err error)
}

// UserService 用户表 模型服务
type UserService struct {
	svc *server.SvcContext
}

// NewUserService 新用户表 模型服务实例
func NewUserService(svc *server.SvcContext) *UserService {
	return &UserService{svc: svc}
}

func (s UserService) GetRealName(params *request.GetRealNameParams) (data *response.GetRealNameResp, err error) {
	service := NewGetRealNameService(s.svc)
	data, err = service.GetRealName(params)

	return
}
