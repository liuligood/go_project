package user_service

import (
	"crmeb_go/internal/server"
)

type EbUserServiceImpl interface {
}

// EbUserService 用户表 模型服务
type EbUserService struct {
	svc *server.SvcContext
}

// NewEbUserService 新用户表 模型服务实例
func NewEbUserService(svc *server.SvcContext) *EbUserService {
	return &EbUserService{svc: svc}
}
