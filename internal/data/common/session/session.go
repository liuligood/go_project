package session

import (
	"crmeb_go/config"
	"crmeb_go/internal/container/repository"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type SvcContext struct {
	RedisClient     redis.UniversalClient
	Logger          *zap.Logger
	Repo            *repository.Container
	Conf            config.Conf
	Gorm            *gorm.DB
	RedisClientList map[string]redis.UniversalClient
}

type LoginUserInfo struct {
	UserId          int64    `json:"user_id,omitempty"`         // 用户id
	Roles           string   `json:"roles,omitempty"`           // 用户权限
	PermissionsList []string `json:"permissionsList,omitempty"` // 权限标识
	ClientIp        string   `json:"client_ip,omitempty"`       // ip地址
}
