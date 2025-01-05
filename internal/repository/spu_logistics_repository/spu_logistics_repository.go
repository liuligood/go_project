package spu_logistics_repository

import (
	"crmeb_go/internal/repository/base_repository"
	"crmeb_go/internal/repository/gen"
	"gorm.io/gorm"
)

type SpuLogisticsRepository struct {
	*base_repository.DBRepository
}

func NewSpuLogisticsRepository(db *gorm.DB, gen *gen.Query) *SpuLogisticsRepository {
	return &SpuLogisticsRepository{base_repository.NewRepository(db, gen)}
}
