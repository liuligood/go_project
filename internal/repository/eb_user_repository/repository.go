package eb_user_repository

import (
	"context"
	"crmeb_go/internal/model"
	"crmeb_go/internal/repository"
	"crmeb_go/internal/repository/gen"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"
)

// EbUserRepository 用户表 模型仓库
type EbUserRepository struct {
	*repository.DBRepository
}

// NewEbUserRepository 新用户表 模型仓库实例
func NewEbUserRepository(db *gorm.DB, log *zap.Logger) *EbUserRepository {
	return &EbUserRepository{repository.NewRepository(db, log)}
}

func (r *EbUserRepository) GetRealName(ctx context.Context, userId uint64) (data model.EbUser, err error) {
	recordDo := r.Gen.EbSmsRecord.WithContext(ctx)
	take, err := recordDo.Take()
	if err != nil {
		r.Log.Error("EbSmsRecord.Take [err]:%v", zap.Error(err))

		return
	}

	ssx := model.EbSystemRole{
		ID:       int32(23),
		RoleName: "123456",
	}

	ssx22 := model.EbSystemRole{
		ID:       int32(24),
		RoleName: "123456",
	}

	fmt.Println("1", time.Now())
	r.Transaction(ctx, func(query *gen.Query) error {
		if err := query.EbSystemRole.WithContext(ctx).Create(&ssx22); err != nil {
			r.Log.Error("Create.Take [err]:%v", zap.Error(err))

			return err
		}

		if err := query.EbSystemRole.WithContext(ctx).Create(&ssx); err != nil {
			r.Log.Error("Create.Take [err]:%v", zap.Error(err))

			return err
		}

		return nil
	})
	fmt.Println("2", time.Now())
	fmt.Println(take)

	var userInfoModel model.EbUser
	options := map[string]any{
		"select": "uid,real_name,account",
	}

	if err := r.QueryOne(ctx, "uid = ?", []any{userId}, &userInfoModel, options); err != nil {
		r.Log.Error("EbUserRepository.QueryOne [err]:%v", zap.Error(err))

		return data, err
	}

	data = userInfoModel

	return data, err
}
