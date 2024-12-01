package repository

import (
	"crmeb_go/internal/model/eb_article_model"
	"crmeb_go/internal/model/eb_system_config_model"
	"crmeb_go/internal/model/eb_system_group_data_model"
	"crmeb_go/internal/model/eb_user_model"
	"crmeb_go/internal/repository/eb_article_repository"
	"crmeb_go/internal/repository/eb_system_config_repository"
	"crmeb_go/internal/repository/eb_system_group_data_repository"
	"crmeb_go/internal/repository/eb_user_repository"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Container struct {
	EbArticleRepository         *eb_article_repository.EbArticleRepository                   //
	EbUserRepository            *eb_user_repository.EbUserRepository                         // 用户表 模型
	EbSystemGroupDataRepository *eb_system_group_data_repository.EbSystemGroupDataRepository // 组合数据详情表 模型
	EbSystemConfigRepository    *eb_system_config_repository.EbSystemConfigRepository        // 配置表 模型
}

func Register(db *gorm.DB, log *zap.Logger) *Container {
	return &Container{
		EbArticleRepository:         eb_article_repository.NewEbArticleRepository(eb_article_model.NewEbArticleModel(db), log),                                      //
		EbUserRepository:            eb_user_repository.NewEbUserRepository(eb_user_model.NewEbUserModel(db), log),                                                  // 用户表 模型
		EbSystemGroupDataRepository: eb_system_group_data_repository.NewEbSystemGroupDataRepository(eb_system_group_data_model.NewEbSystemGroupDatumModel(db), log), // 组合数据详情表 模型
		EbSystemConfigRepository:    eb_system_config_repository.NewEbSystemConfigRepository(eb_system_config_model.NewEbSystemConfigModel(db), log),                // 配置表 模型
	}
}
