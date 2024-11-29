package service

import (
	"crmeb_go/internal/server"
	"crmeb_go/internal/service/eb_article_service"
	"crmeb_go/internal/service/eb_user_service"
)

type Container struct {
	UserService    eb_user_service.EbUserServiceImpl
	ArticleService eb_article_service.EbArticleServiceImpl
}

func Register(svc *server.SvcContext) *Container {
	return &Container{
		UserService:    eb_user_service.NewEbUserService(svc),
		ArticleService: eb_article_service.NewEbArticleService(svc),
	}
}
