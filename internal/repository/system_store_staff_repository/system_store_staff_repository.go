package system_store_staff_repository

import (
	"context"
	"crmeb_go/internal/model"
	"crmeb_go/internal/repository/base_repository"
	"crmeb_go/internal/repository/gen"
	"gorm.io/gorm"
)

type SystemStoreStaffRepository struct {
	*base_repository.DBRepository
}

func NewSystemStoreStaffRepository(db *gorm.DB, gen *gen.Query) *SystemStoreStaffRepository {
	return &SystemStoreStaffRepository{base_repository.NewRepository(db, gen)}
}

func (s *SystemStoreStaffRepository) FindListByStoreId(ctx context.Context, storeId int64) (data []*model.SystemStoreStaff, err error) {
	if storeId > 0 {
		return s.Gen.WithContext(ctx).SystemStoreStaff.Where(s.Gen.SystemStoreStaff.StoreID.Eq(storeId)).Find()
	} else {
		return s.Gen.WithContext(ctx).SystemStoreStaff.Debug().Find()
	}
}

func (s *SystemStoreStaffRepository) FindListByStoreIdCount(ctx context.Context, storeId int64) (data int64, err error) {
	if storeId > 0 {
		return s.Gen.WithContext(ctx).SystemStoreStaff.Where(s.Gen.SystemStoreStaff.StoreID.Eq(storeId)).Count()
	} else {
		return s.Gen.WithContext(ctx).SystemStoreStaff.Count()
	}
}
