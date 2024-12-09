package server

import (
	"context"
	"crmeb_go/config"
	"crmeb_go/internal/container/repository"
	"crmeb_go/internal/repository/gen"
	"crmeb_go/utils/icache"
	"crmeb_go/utils/igorm"
	iredis "crmeb_go/utils/iredis"
	"crmeb_go/utils/iupload"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v9"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type SvcContext struct {
	Ctx             context.Context
	Conf            config.Conf
	Gorm            *gorm.DB
	Repo            *repository.Container
	RedisClient     redis.UniversalClient
	RedisClientList map[string]redis.UniversalClient
	Redsync         *redsync.Redsync
	Gen             *gen.Query
	Oss             iupload.OSS
}

func NewSvcContext(c config.Conf) *SvcContext {
	svc := &SvcContext{
		Ctx:  context.Background(),
		Conf: c,
	}

	db := igorm.NewDB(c)
	svc.Gorm = igorm.Gorm(db)
	svc.Gen = gen.Use(svc.Gorm)
	svc.Repo = repository.Register(svc.Gorm, svc.Gen)
	svc.Oss = iupload.NewOss(c)

	if c.System.UseRedis {
		// 初始化redis服务
		svc.RedisClient = iredis.Redis(c)
		svc.RedisClientList = iredis.RedisList(c)
		svc.Redsync = redsync.New(goredis.NewPool(svc.RedisClient))
		icache.InitLocalCache()
	}

	return svc
}
