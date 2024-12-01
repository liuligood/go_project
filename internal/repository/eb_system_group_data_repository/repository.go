package eb_system_group_data_repository

import (
	"crmeb_go/internal/model/eb_system_group_data_model"
	"crmeb_go/internal/repository"
	"go.uber.org/zap"
)

// EbSystemGroupDataRepository 组合数据详情表 模型仓库
type EbSystemGroupDataRepository struct {
	*repository.DBRepository
}

// NewEbSystemGroupDataRepository 新组合数据详情表 模型仓库实例
func NewEbSystemGroupDataRepository(model *eb_system_group_data_model.EbSystemGroupDatumModel, zap *zap.Logger) *EbSystemGroupDataRepository {
	return &EbSystemGroupDataRepository{repository.NewDBRepository(model, zap)}
}
