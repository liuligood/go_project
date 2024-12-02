package eb_user_repository

import (
	"context"
	"crmeb_go/internal/model/eb_user_model"
	"crmeb_go/internal/repository"
	"go.uber.org/zap"
)

// EbUserRepository 用户表 模型仓库
type EbUserRepository struct {
	*repository.DBRepository
}

// NewEbUserRepository 新用户表 模型仓库实例
func NewEbUserRepository(model *eb_user_model.EbUserModel, log *zap.Logger) *EbUserRepository {
	return &EbUserRepository{repository.NewDBRepository(model, log)}
}

func (r *EbUserRepository) GetRealName(ctx context.Context, userId uint64) (data eb_user_model.EbUserModel, err error) {
	var userInfoModel eb_user_model.EbUserModel
	options := map[string]any{
		"select": "uid,real_name,account",
	}

	if err := r.QueryOne(ctx, "uid = ?", []any{userId}, &userInfoModel, options); err != nil {
		r.Log.Error("EbUserRepository.QueryOne [err]:%v", zap.Error(err))

		return
	}

	return
}
