package internal

import (
	"crmeb_go/internal/container/service"
	"crmeb_go/internal/server"
)

type AppContent struct {
	Svc     *server.SvcContext
	Service *service.Container
}

func Register(svc *server.SvcContext) *AppContent {
	return &AppContent{
		Svc:     svc,
		Service: service.Register(svc),
	}
}
