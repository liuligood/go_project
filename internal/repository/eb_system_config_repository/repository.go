package eb_system_config_repository

import (
	"context"
	"crmeb_go/internal/common/define"
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

func (e EbSystemConfigRepository) getSystemConfigInfoValue(ctx context.Context, name string) (data string, err error) {
	var model eb_system_config_model.EbSystemConfigModel

	options := map[string]any{
		"select": "name,value",
	}

	err = e.QueryOne(ctx, "name = ? and status = ?", []any{name, define.NoStatus}, &model, options)
	if err != nil {
		e.Log.Error("EbSystemConfigRepository.QueryOne [err]:%v", zap.Error(err))

		return
	}

	data = model.Value

	return
}
