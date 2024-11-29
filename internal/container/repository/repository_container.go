package repository

import (
	"crmeb_go/internal/model/eb_article_model"
	"crmeb_go/internal/model/eb_user_model"
	"crmeb_go/internal/repository/eb_article_repository"
	"crmeb_go/internal/repository/eb_user_repository"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Container struct {
	EbArticleRepository *eb_article_repository.EbArticleRepository //
	EbUserRepository    *eb_user_repository.EbUserRepository       // 用户表 模型
}

func Register(db *gorm.DB, log *zap.Logger) *Container {
	return &Container{
		EbArticleRepository: eb_article_repository.NewEbArticleRepository(eb_article_model.NewEbArticleModel(db), log), //
		EbUserRepository:    eb_user_repository.NewEbUserRepository(eb_user_model.NewEbUserModel(db), log),             // 用户表 模型
	}
}
