package system_config_repository

import (
	"context"
	"crmeb_go/internal/model"
	"crmeb_go/internal/repository/base_repository"
	"crmeb_go/internal/repository/gen"

	"gorm.io/gorm"
)

// SystemConfigRepository 配置表 模型仓库
type SystemConfigRepository struct {
	*base_repository.DBRepository
}

// NewUserRepository 配置表 模型仓库实例
func NewSystemConfigRepository(db *gorm.DB, gen *gen.Query) *SystemConfigRepository {
	return &SystemConfigRepository{base_repository.NewRepository(db, gen)}
}

func (r *SystemConfigRepository) All(ctx context.Context) (data []*model.SystemConfig, err error) {
	SystemConfigList, err := r.Gen.SystemConfig.WithContext(ctx).Find()
	if err != nil {
		return data, err
	}

	data = SystemConfigList

	return data, err
}

func (r *SystemConfigRepository) QueryByName(ctx context.Context, name string) (data *model.SystemConfig, err error) {
	SystemConfig, err := r.Gen.SystemConfig.WithContext(ctx).Where(r.Gen.SystemConfig.Name.Eq(name)).First()
	if err != nil {
		return data, err
	}

	data = SystemConfig

	return data, err
}
