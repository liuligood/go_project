package server

import (
	"context"
	"crmeb_go/config"
	"crmeb_go/internal/container/repository"
	"crmeb_go/utils/igorm"
	iredis "crmeb_go/utils/iredis"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type SvcContext struct {
	Ctx             context.Context
	Conf            config.Conf
	Logger          *zap.Logger
	Gorm            *gorm.DB
	Repo            *repository.Container
	RedisClient     redis.UniversalClient
	RedisClientList map[string]redis.UniversalClient
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

	if c.System.UseRedis {
		// 初始化redis服务
		svc.RedisClient = iredis.Redis(c, logger)
		svc.RedisClientList = iredis.RedisList(c, logger)
	}

	return svc
}
