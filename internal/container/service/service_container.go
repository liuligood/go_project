package service

import (
	"crmeb_go/internal/server"
	"crmeb_go/internal/service/admin_service"
	"crmeb_go/internal/service/eb_article_service"
	"crmeb_go/internal/service/eb_system_config_service"
	"crmeb_go/internal/service/eb_system_group_data_service"
	"crmeb_go/internal/service/eb_user_service"
)

type Container struct {
	UserService              eb_user_service.EbUserServiceImpl
	ArticleService           eb_article_service.EbArticleServiceImpl
	AdminService             admin_service.AdminServiceImpl
	EbSystemGroupDataService eb_system_group_data_service.EbSystemGroupDataServiceImpl // 组合数据详情表 模型
	EbSystemConfigService    eb_system_config_service.EbSystemConfigServiceImpl        // 配置表 模型
}

func Register(svc *server.SvcContext) *Container {
	return &Container{
		UserService:              eb_user_service.NewEbUserService(svc),
		ArticleService:           eb_article_service.NewEbArticleService(svc),
		AdminService:             admin_service.NewAdminService(svc),
		EbSystemGroupDataService: eb_system_group_data_service.NewEbSystemGroupDataService(svc), // 组合数据详情表 模型
		EbSystemConfigService:    eb_system_config_service.NewEbSystemConfigService(svc),        // 配置表 模型
	}
}
