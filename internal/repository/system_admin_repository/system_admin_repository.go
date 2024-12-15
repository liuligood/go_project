package system_admin_repository

import (
	"context"
	"crmeb_go/internal/model"
	"crmeb_go/internal/repository/base_repository"
	"crmeb_go/internal/repository/gen"
	"gorm.io/gorm"
)

type SystemAdminRepository struct {
	*base_repository.DBRepository
}

func NewSystemAdminRepository(db *gorm.DB, gen *gen.Query) *SystemAdminRepository {
	return &SystemAdminRepository{base_repository.NewRepository(db, gen)}
}

func (r *SystemAdminRepository) QueryAdminByAccount(ctx context.Context, account string) (data *model.SystemAdmin, err error) {
	return r.Gen.SystemAdmin.WithContext(ctx).Where(r.Gen.SystemAdmin.Account.Eq(account)).First()
}
func (r *SystemAdminRepository) QueryAdminById(ctx context.Context, userId int64) (data *model.SystemAdmin, err error) {
	return r.Gen.SystemAdmin.WithContext(ctx).Where(r.Gen.SystemAdmin.ID.Eq(userId)).First()
}
