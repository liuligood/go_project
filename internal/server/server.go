package server

import (
	"context"
	"crmeb_go/config"
	"crmeb_go/internal/container/repository"
	"crmeb_go/internal/repository/gen"
	"crmeb_go/utils/upload"
	"github.com/go-redsync/redsync/v4/redis/goredis/v9"

	"crmeb_go/utils/igorm"
	iredis "crmeb_go/utils/iredis"
	"github.com/go-redsync/redsync/v4"
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
	Redsync         *redsync.Redsync
	Gen             *gen.Query
	Oss             upload.OSS
}

func NewSvcContext(c config.Conf, logger *zap.Logger) *SvcContext {
	svc := &SvcContext{
		Ctx:    context.Background(),
		Conf:   c,
		Logger: logger,
	}

	db := igorm.NewDB(c)
	svc.Gorm = igorm.Gorm(db)
	svc.Gen = gen.Use(svc.Gorm)
	svc.Repo = repository.Register(svc.Gorm, svc.Logger, svc.Gen)
	svc.Oss = upload.NewOss(c, logger)

	if c.System.UseRedis {
		// 初始化redis服务
		svc.RedisClient = iredis.Redis(c, logger)
		svc.RedisClientList = iredis.RedisList(c, logger)
		svc.Redsync = redsync.New(goredis.NewPool(svc.RedisClient))
	}

	return svc
}
