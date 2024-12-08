package article_service

import (
	"crmeb_go/internal/server"
)

type ArticleServiceImpl interface {
}

// ArticleService 服务
type ArticleService struct {
	svc *server.SvcContext
}

// NewArticleService 新服务实例
func NewArticleService(svc *server.SvcContext) *ArticleService {
	return &ArticleService{svc: svc}
}
