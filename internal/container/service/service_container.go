package service

import (
	"crmeb_go/internal/server"
	"crmeb_go/internal/service/admin_service"
	"crmeb_go/internal/service/eb_article_service"
	"crmeb_go/internal/service/eb_user_service"
)

type Container struct {
	UserService    eb_user_service.EbUserServiceImpl
	ArticleService eb_article_service.EbArticleServiceImpl
	AdminService   admin_service.AdminServiceImpl
}

func Register(svc *server.SvcContext) *Container {
	return &Container{
		UserService:    eb_user_service.NewEbUserService(svc),
		ArticleService: eb_article_service.NewEbArticleService(svc),
		AdminService:   admin_service.NewAdminService(svc),
	}
}
