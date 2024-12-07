package repository

import (
	"crmeb_go/internal/repository/eb_system_config_repository"
	"crmeb_go/internal/repository/eb_user_repository"
	"crmeb_go/internal/repository/gen"
	"gorm.io/gorm"
)

type Container struct {
	EbUserRepository         *eb_user_repository.EbUserRepository                  // 用户表 模型
	EbSystemConfigRepository *eb_system_config_repository.EbSystemConfigRepository // 配置表
}

func Register(db *gorm.DB, gen *gen.Query) *Container {
	return &Container{
		EbUserRepository:         eb_user_repository.NewEbUserRepository(db, gen),                  // 用户表 模型
		EbSystemConfigRepository: eb_system_config_repository.NewEbSystemConfigRepository(db, gen), // 配置表
	}
}
