package session_context_data

import (
	"crmeb_go/config"
	"crmeb_go/internal/container/repository"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type SvcContext struct {
	//RedisClient *xredis.RedisClient
	Logger *zap.Logger
	Repo   *repository.Container
	Conf   config.Conf
	Gorm   *gorm.DB
	//MQClient *rocketmq.Client
}
