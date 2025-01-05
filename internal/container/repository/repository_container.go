package repository

import (
	"crmeb_go/internal/repository/gen"
	"crmeb_go/internal/repository/sku_attribute_repository"
	"crmeb_go/internal/repository/sku_repository"
	"crmeb_go/internal/repository/spu_log_repository"
	"crmeb_go/internal/repository/spu_logistics_repository"
	"crmeb_go/internal/repository/spu_picture_repository"
	"crmeb_go/internal/repository/spu_repository"
	"crmeb_go/internal/repository/spu_service_repository"
	"crmeb_go/internal/repository/spu_stock_stream_repository"
	"crmeb_go/internal/repository/store_order_repository"
	"crmeb_go/internal/repository/system_admin_repository"
	"crmeb_go/internal/repository/system_config_repository"
	"crmeb_go/internal/repository/system_group_data_repository"
	"crmeb_go/internal/repository/system_menu_repository"
	"crmeb_go/internal/repository/system_store_repository"
	"crmeb_go/internal/repository/system_store_staff_repository"
	"crmeb_go/internal/repository/user_repository"
	"crmeb_go/internal/repository/user_visit_repository"
	"gorm.io/gorm"
)

type Container struct {
	UserRepository             *user_repository.UserRepository                  // 用户表 模型
	SystemConfigRepository     *system_config_repository.SystemConfigRepository // 配置表
	SystemGroupDataRepository  *system_group_data_repository.SystemGroupDataRepository
	SystemAdminRepository      *system_admin_repository.SystemAdminRepository
	SystemMenuRepository       *system_menu_repository.SystemMenuRepository
	SystemStoreRepository      *system_store_repository.SystemStoreRepository
	SystemStoreStaffRepository *system_store_staff_repository.SystemStoreStaffRepository
	UserVisitRecordRepository  *user_visit_repository.UserVisitRecordRepository
	StoreOrderRepository       *store_order_repository.StoreOrderRepository
	SpuRepository              *spu_repository.SpuRepository
	SkuRepository              *sku_repository.SkuRepository
	SkuAttributeRepository     *sku_attribute_repository.SkuAttributeRepository
	SpuLogRepository           *spu_log_repository.SpuLogRepository
	SpuLogisticsRepository     *spu_logistics_repository.SpuLogisticsRepository
	SpuPictureRepository       *spu_picture_repository.SpuPictureRepository
	SpuServiceRepository       *spu_service_repository.SpuServiceRepository
	SpuStockStreamRepository   *spu_stock_stream_repository.SpuStockStreamRepository
}

func Register(db *gorm.DB, gen *gen.Query) *Container {
	return &Container{
		UserRepository:             user_repository.NewUserRepository(db, gen),                           // 用户表
		SystemConfigRepository:     system_config_repository.NewSystemConfigRepository(db, gen),          // 配置表
		SystemGroupDataRepository:  system_group_data_repository.NewSystemGroupDataRepository(db, gen),   // 配置表
		SystemAdminRepository:      system_admin_repository.NewSystemAdminRepository(db, gen),            // 后台管理表
		SystemMenuRepository:       system_menu_repository.NewSystemMenuRepository(db, gen),              // 菜单表
		SystemStoreRepository:      system_store_repository.NewSystemStoreRepository(db, gen),            // 门店自提表
		SystemStoreStaffRepository: system_store_staff_repository.NewSystemStoreStaffRepository(db, gen), // 门店店员表
		UserVisitRecordRepository:  user_visit_repository.NewUserVisitRecordRepository(db, gen),          // 用户访问记录表
		StoreOrderRepository:       store_order_repository.NewStoreOrderRepository(db, gen),              // 订单表
		SpuRepository:              spu_repository.NewSpuRepository(db, gen),                             // 商品表
		SkuRepository:              sku_repository.NewSkuRepository(db, gen),                             // 套餐表
		SkuAttributeRepository:     sku_attribute_repository.NewSkuAttributeRepository(db, gen),          // 套餐属性表
		SpuLogRepository:           spu_log_repository.NewSpuLogRepository(db, gen),                      // 商品日志表
		SpuLogisticsRepository:     spu_logistics_repository.NewSpuLogisticsRepository(db, gen),          // 商品物流表
		SpuPictureRepository:       spu_picture_repository.NewSpuPictureRepository(db, gen),              // 商品图片表
		SpuServiceRepository:       spu_service_repository.NewSpuServiceRepository(db, gen),              // 商品服务表
		SpuStockStreamRepository:   spu_stock_stream_repository.NewSpuStockStreamRepository(db, gen),     // 商品库流水表
	}
}
