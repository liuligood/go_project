package eb_user_repository

import (
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
