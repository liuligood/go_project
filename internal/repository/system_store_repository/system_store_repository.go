package system_store_repository

import (
	"context"
	"crmeb_go/internal/model"
	"crmeb_go/internal/repository/base_repository"
	"crmeb_go/internal/repository/gen"
	"gorm.io/gorm"
)

type SystemStoreRepository struct {
	*base_repository.DBRepository
}

func NewSystemStoreRepository(db *gorm.DB, gen *gen.Query) *SystemStoreRepository {
	return &SystemStoreRepository{base_repository.NewRepository(db, gen)}
}

func (s *SystemStoreRepository) FindListById(ctx context.Context, storeId []int64) (data []*model.SystemStore, err error) {
	return s.Gen.WithContext(ctx).SystemStore.Debug().Where(s.Gen.SystemStore.ID.In(storeId...)).Find()
}
