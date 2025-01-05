package sku_repository

import (
	"crmeb_go/internal/repository/base_repository"
	"crmeb_go/internal/repository/gen"
	"gorm.io/gorm"
)

type SkuRepository struct {
	*base_repository.DBRepository
}

func NewSkuRepository(db *gorm.DB, gen *gen.Query) *SkuRepository {
	return &SkuRepository{base_repository.NewRepository(db, gen)}
}
