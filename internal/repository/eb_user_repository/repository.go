package eb_user_repository

import (
	"context"
	"crmeb_go/internal/model"
	"crmeb_go/internal/repository"
	"crmeb_go/internal/repository/gen"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// EbUserRepository 用户表 模型仓库
type EbUserRepository struct {
	*repository.DBRepository
}

// NewEbUserRepository 新用户表 模型仓库实例
func NewEbUserRepository(db *gorm.DB, log *zap.Logger, gen *gen.Query) *EbUserRepository {
	return &EbUserRepository{repository.NewRepository(db, log, gen)}
}

func (r *EbUserRepository) GetRealName(ctx context.Context, userId uint64) (data model.EbUser, err error) {
	ebUser, err := r.Gen.EbUser.WithContext(ctx).Where(r.Gen.EbUser.UID.Eq(int32(userId))).First()
	if err != nil {
		return data, err
	}

	data = *ebUser

	result, count, err := r.Gen.EbSystemRole.WithContext(ctx).Where(r.Gen.EbSystemRole.Level.Eq(int32(userId))).FindByPage(1, 4)
	if err != nil {
		return data, err
	}

	fmt.Println(result, count)
	data = *ebUser

	return data, err
}
