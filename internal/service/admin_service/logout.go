package admin_service

import (
	"crmeb_go/internal/data/request"
	"crmeb_go/internal/server"
)

type LogoutServiceImpl interface {
	Logout(*request.BaseServiceParams) (err error)
}

type LogoutService struct {
	svc *server.SvcContext
}

func NewLogoutService(svc *server.SvcContext) *LogoutService {
	return &LogoutService{svc: svc}
}

func (l *LogoutService) Logout(*request.BaseServiceParams) (err error) {

	return
}
