package product_service

import (
	"crmeb_go/internal/data/request"
	"crmeb_go/internal/data/response"
	"crmeb_go/internal/server"
	"crmeb_go/utils/izap"
	"go.uber.org/zap"
)

type ProductServiceImpl interface {
	CreateProduct(params *request.CreateProductParams) (resp *response.CreateProductResp, err error)
}

type ProductService struct {
	svc *server.SvcContext
}

func NewProductService(svc *server.SvcContext) *ProductService {
	return &ProductService{svc: svc}
}

func (p *ProductService) CreateProduct(params *request.CreateProductParams) (resp *response.CreateProductResp, err error) {
	izap.Log.WithContext(params.Ctx).Info("ss", zap.String("错误", "2"))
	p.svc.Repo.SpuRepository.CreateProduct(params.Ctx, params.CreateProduct)
	return
}
