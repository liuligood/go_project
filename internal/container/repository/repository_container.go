package repository

import (
	"crmeb_go/internal/repository/gen"
	"crmeb_go/internal/repository/system_config_repository"
	"crmeb_go/internal/repository/system_group_data_repository"
	"crmeb_go/internal/repository/user_repository"
	"gorm.io/gorm"
)

type Container struct {
	UserRepository            *user_repository.UserRepository                  // 用户表 模型
	SystemConfigRepository    *system_config_repository.SystemConfigRepository // 配置表
	SystemGroupDataRepository *system_group_data_repository.SystemGroupDataRepository
}

func Register(db *gorm.DB, gen *gen.Query) *Container {
	return &Container{
		UserRepository:            user_repository.NewUserRepository(db, gen),                         // 用户表 模型
		SystemConfigRepository:    system_config_repository.NewSystemConfigRepository(db, gen),        // 配置表
		SystemGroupDataRepository: system_group_data_repository.NewSystemGroupDataRepository(db, gen), // 配置表
	}
}
