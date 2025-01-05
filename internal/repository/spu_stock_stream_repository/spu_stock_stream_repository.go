package spu_stock_stream_repository

import (
	"crmeb_go/internal/repository/base_repository"
	"crmeb_go/internal/repository/gen"
	"gorm.io/gorm"
)

type SpuStockStreamRepository struct {
	*base_repository.DBRepository
}

func NewSpuStockStreamRepository(db *gorm.DB, gen *gen.Query) *SpuStockStreamRepository {
	return &SpuStockStreamRepository{base_repository.NewRepository(db, gen)}
}
