package user_repository

import (
	"context"
	"crmeb_go/internal/model"
	"crmeb_go/internal/repository/base_repository"
	"crmeb_go/internal/repository/gen"
	"gorm.io/gorm"
)

// UserRepository 用户表 模型仓库
type UserRepository struct {
	*base_repository.DBRepository
}

// NewUserRepository 新用户表 模型仓库实例
func NewUserRepository(db *gorm.DB, gen *gen.Query) *UserRepository {
	return &UserRepository{base_repository.NewRepository(db, gen)}
}

func (r *UserRepository) QueryRealName(ctx context.Context, userId uint64) (data model.User, err error) {
	ebUser, err := r.Gen.User.WithContext(ctx).Where(r.Gen.User.UID.Eq(int64(userId))).First()
	if err != nil {
		return data, err
	}

	data = *ebUser
	return data, err
}
