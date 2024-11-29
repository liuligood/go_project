package server

import (
	"context"
	"crmeb_go/config"
	"crmeb_go/internal/container/repository"
	"crmeb_go/utils/igorm"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type SvcContext struct {
	Ctx    context.Context
	Conf   config.Conf
	Logger *zap.Logger
	Gorm   *gorm.DB
	Repo   *repository.Container
}

func NewSvcContext(c config.Conf, logger *zap.Logger) *SvcContext {
	svc := &SvcContext{
		Ctx:    context.Background(),
		Conf:   c,
		Logger: logger,
	}

	db := igorm.NewDB(c)
	svc.Gorm = igorm.Gorm(db)
	svc.Repo = repository.Register(svc.Gorm, svc.Logger)

	return svc
}
