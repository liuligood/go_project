package eb_article_service

import (
	"crmeb_go/internal/server"
)

type EbArticleServiceImpl interface {
}

// EbArticleService 服务
type EbArticleService struct {
	svc *server.SvcContext
}

// NewEbArticleService 新服务实例
func NewEbArticleService(svc *server.SvcContext) *EbArticleService {
	return &EbArticleService{svc: svc}
}
