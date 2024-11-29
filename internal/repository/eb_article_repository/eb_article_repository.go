package eb_article_repository

import (
	"crmeb_go/internal/model/eb_article_model"
	"crmeb_go/internal/repository"
	"go.uber.org/zap"
)

// EbArticleRepository 仓库
type EbArticleRepository struct {
	*repository.DBRepository
}

// NewEbArticleRepository 新仓库实例
func NewEbArticleRepository(model *eb_article_model.EbArticleModel, log *zap.Logger) *EbArticleRepository {
	return &EbArticleRepository{repository.NewDBRepository(model, log)}
}
