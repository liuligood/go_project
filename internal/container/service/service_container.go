package service

import (
	"crmeb_go/internal/server"
	"crmeb_go/internal/service/admin_service"
	"crmeb_go/internal/service/article_service"
	"crmeb_go/internal/service/home_service"
	"crmeb_go/internal/service/system_config_service"
	"crmeb_go/internal/service/system_group_data_service"
	"crmeb_go/internal/service/system_menu_service"
	"crmeb_go/internal/service/system_store_staff_service"
	"crmeb_go/internal/service/user_service"
)

type Container struct {
	UserService             user_service.UserServiceImpl
	ArticleService          article_service.ArticleServiceImpl
	AdminService            admin_service.AdminServiceImpl
	SystemGroupDataService  system_group_data_service.SystemGroupDataServiceImpl // 组合数据详情表 模型
	SystemConfigService     system_config_service.SystemConfigServiceImpl        // 配置表 模型
	SystemMenuService       system_menu_service.SystemMenuServiceImpl
	SystemStoreStaffService system_store_staff_service.SystemStoreStaffServiceImpl
	HomeService             home_service.HomeServiceImpl
}

func Register(svc *server.SvcContext) *Container {
	return &Container{
		UserService:             user_service.NewUserService(svc),
		ArticleService:          article_service.NewArticleService(svc),
		AdminService:            admin_service.NewAdminService(svc),
		SystemGroupDataService:  system_group_data_service.NewSystemGroupDataService(svc), // 组合数据详情表 模型
		SystemConfigService:     system_config_service.NewSystemConfigService(svc),        // 配置表 模型
		SystemMenuService:       system_menu_service.NewSystemMenuService(svc),
		SystemStoreStaffService: system_store_staff_service.NewSystemStoreStaffService(svc),
		HomeService:             home_service.NewHomeService(svc),
	}
}
