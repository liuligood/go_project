package repository

import (
	"crmeb_go/internal/repository/eb_user_repository"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Container struct {
	EbUserRepository *eb_user_repository.EbUserRepository // 用户表 模型
}

func Register(db *gorm.DB, log *zap.Logger) *Container {
	return &Container{
		EbUserRepository: eb_user_repository.NewEbUserRepository(db, log), // 用户表 模型
	}
}
