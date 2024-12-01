package eb_system_config_repository

import (
	"crmeb_go/internal/model/eb_system_config_model"
	"crmeb_go/internal/repository"
	"go.uber.org/zap"
)

// EbSystemConfigRepository 配置表 模型仓库
type EbSystemConfigRepository struct {
	*repository.DBRepository
}

// NewEbSystemConfigRepository 新配置表 模型仓库实例
func NewEbSystemConfigRepository(model *eb_system_config_model.EbSystemConfigModel, log *zap.Logger) *EbSystemConfigRepository {
	return &EbSystemConfigRepository{repository.NewDBRepository(model, log)}
}
