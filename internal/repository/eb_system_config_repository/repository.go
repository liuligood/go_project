package eb_system_config_repository

import (
	"context"
	"crmeb_go/internal/model"
	"crmeb_go/internal/repository"
	"crmeb_go/internal/repository/gen"
	"gorm.io/gorm"
)

// EbSystemConfigRepository 配置表 模型仓库
type EbSystemConfigRepository struct {
	*repository.DBRepository
}

// NewEbUserRepository 配置表 模型仓库实例
func NewEbSystemConfigRepository(db *gorm.DB, gen *gen.Query) *EbSystemConfigRepository {
	return &EbSystemConfigRepository{repository.NewRepository(db, gen)}
}

func (r *EbSystemConfigRepository) All(ctx context.Context) (data []*model.EbSystemConfig, err error) {
	ebSystemConfigList, err := r.Gen.EbSystemConfig.WithContext(ctx).Find()
	if err != nil {
		return data, err
	}

	data = ebSystemConfigList

	return data, err
}

func (r *EbSystemConfigRepository) QueryByName(ctx context.Context, name string) (data *model.EbSystemConfig, err error) {
	ebSystemConfig, err := r.Gen.EbSystemConfig.WithContext(ctx).Where(r.Gen.EbSystemConfig.Name.Eq("name")).First()
	if err != nil {
		return data, err
	}

	data = ebSystemConfig

	return data, err
}
