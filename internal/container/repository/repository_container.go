package repository

import (
	"crmeb_go/internal/repository/eb_user_repository"
	"crmeb_go/internal/repository/gen"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Container struct {
	EbUserRepository *eb_user_repository.EbUserRepository // 用户表 模型
}

func Register(db *gorm.DB, log *zap.Logger, gen *gen.Query) *Container {
	return &Container{
		EbUserRepository: eb_user_repository.NewEbUserRepository(db, log, gen), // 用户表 模型
	}
}
