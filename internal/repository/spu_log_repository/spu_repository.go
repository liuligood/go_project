package spu_log_repository

import (
	"crmeb_go/internal/repository/base_repository"
	"crmeb_go/internal/repository/gen"
	"gorm.io/gorm"
)

type SpuLogRepository struct {
	*base_repository.DBRepository
}

func NewSpuLogRepository(db *gorm.DB, gen *gen.Query) *SpuLogRepository {
	return &SpuLogRepository{base_repository.NewRepository(db, gen)}
}
