package sku_attribute_repository

import (
	"crmeb_go/internal/repository/base_repository"
	"crmeb_go/internal/repository/gen"
	"gorm.io/gorm"
)

type SkuAttributeRepository struct {
	*base_repository.DBRepository
}

func NewSkuAttributeRepository(db *gorm.DB, gen *gen.Query) *SkuAttributeRepository {
	return &SkuAttributeRepository{base_repository.NewRepository(db, gen)}
}
