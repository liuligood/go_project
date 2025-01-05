package spu_repository

import (
	"context"
	"crmeb_go/internal/repository/base_repository"
	"crmeb_go/internal/repository/gen"
	"crmeb_go/internal/validation"
	"crmeb_go/utils/izap"
	"gorm.io/gorm"
)

type SpuRepository struct {
	*base_repository.DBRepository
}

func NewSpuRepository(db *gorm.DB, gen *gen.Query) *SpuRepository {
	return &SpuRepository{base_repository.NewRepository(db, gen)}
}

func (s *SpuRepository) CreateProduct(ctx context.Context, params *validation.CreateProductParam) error {
	izap.Log.WithContext(ctx).Error("这是错误的")
	return nil
}
