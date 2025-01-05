package spu_picture_repository

import (
	"crmeb_go/internal/repository/base_repository"
	"crmeb_go/internal/repository/gen"
	"gorm.io/gorm"
)

type SpuPictureRepository struct {
	*base_repository.DBRepository
}

func NewSpuPictureRepository(db *gorm.DB, gen *gen.Query) *SpuPictureRepository {
	return &SpuPictureRepository{base_repository.NewRepository(db, gen)}
}
