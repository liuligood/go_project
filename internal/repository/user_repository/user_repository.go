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

func (u *UserRepository) FindRealName(ctx context.Context, userId uint64) (data *model.User, err error) {
	ebUser, err := u.Gen.User.WithContext(ctx).Where(u.Gen.User.UID.Eq(int64(userId))).First()
	if err != nil {
		return data, err
	}

	data = ebUser
	return data, err
}

func (u *UserRepository) FindUserByAccount(ctx context.Context, account string) (data *model.User, err error) {
	return u.Gen.User.WithContext(ctx).Where(u.Gen.User.Account.Eq(account)).First()
}

func (u *UserRepository) FindListById(ctx context.Context, userIdList []int64) (data []*model.User, err error) {
	return u.Gen.User.WithContext(ctx).Where(u.Gen.User.UID.In(userIdList...)).Find()
}

func (u *UserRepository) GetRegisterNumByDate(ctx context.Context, start, end int64) (data int64, err error) {
	return u.Gen.User.WithContext(ctx).Where(u.Gen.User.CreatedAt.Between(start, end)).Count()
}
