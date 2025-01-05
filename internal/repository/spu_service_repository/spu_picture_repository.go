package spu_service_repository

import (
	"crmeb_go/internal/repository/base_repository"
	"crmeb_go/internal/repository/gen"
	"gorm.io/gorm"
)

type SpuServiceRepository struct {
	*base_repository.DBRepository
}

func NewSpuServiceRepository(db *gorm.DB, gen *gen.Query) *SpuServiceRepository {
	return &SpuServiceRepository{base_repository.NewRepository(db, gen)}
}
