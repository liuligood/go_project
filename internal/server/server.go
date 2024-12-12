package server

import (
	"context"
	"crmeb_go/config"
	"crmeb_go/internal/container/repository"
	"crmeb_go/internal/repository/gen"
	"crmeb_go/utils/captcha"
	"crmeb_go/utils/icache"
	"crmeb_go/utils/igorm"
	oss "crmeb_go/utils/ioss"
	iredis "crmeb_go/utils/iredis"
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
	CaptchaClient   *captcha.CaptchaClient
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

	if c.System.UseRedis {
		// 初始化redis服务
		svc.RedisClient = iredis.Redis(c)
		svc.RedisClientList = iredis.RedisList(c)
		// 初始化分布式锁
		svc.Redsync = redsync.New(goredis.NewPool(svc.RedisClient))
		// 初始化本地缓存
		icache.InitLocalCache()
		// 使用stsToken 上传
		oss.InitOSS(svc.RedisClient, c)
		// 初始化oss
		oss.GetOssClient(c)

		svc.CaptchaClient = captcha.NewCaptchaClient(svc.RedisClient, svc.Ctx)
	}

	return svc
}
