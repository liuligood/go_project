package service

import (
	"crmeb_go/internal/server"
	"crmeb_go/internal/service/admin_service"
	"crmeb_go/internal/service/article_service"
	"crmeb_go/internal/service/home_service"
	"crmeb_go/internal/service/product_service"
	"crmeb_go/internal/service/store_order_service"
	"crmeb_go/internal/service/system_config_service"
	"crmeb_go/internal/service/system_group_data_service"
	"crmeb_go/internal/service/system_menu_service"
	"crmeb_go/internal/service/system_store_staff_service"
	"crmeb_go/internal/service/user_service"
	"crmeb_go/internal/service/user_visit_record_service"
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
	UserVisitRecordService  user_visit_record_service.UserVisitRecordServiceImpl
	StoreOrderService       store_order_service.StoreOrderServiceImpl
	ProductService          product_service.ProductServiceImpl
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
		UserVisitRecordService:  user_visit_record_service.NewUserVisitRecordService(svc),
		StoreOrderService:       store_order_service.NewStoreOrderService(svc),
		ProductService:          product_service.NewProductService(svc),
	}
}
