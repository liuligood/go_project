package store_order_repository

import (
	"context"
	"crmeb_go/define"
	"crmeb_go/internal/model"
	"crmeb_go/internal/model/model_data"
	"crmeb_go/internal/repository/base_repository"
	"crmeb_go/internal/repository/gen"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type StoreOrderRepository struct {
	*base_repository.DBRepository
}

func NewStoreOrderRepository(db *gorm.DB, gen *gen.Query) *StoreOrderRepository {
	return &StoreOrderRepository{base_repository.NewRepository(db, gen)}
}

func (s *StoreOrderRepository) FindOrderNumByDate(ctx context.Context, start int64, end int64) (data int64, err error) {
	return s.Gen.WithContext(ctx).StoreOrder.Select(s.Gen.StoreOrder.ID).
		Where(s.Gen.StoreOrder.CreatedAt.Between(start, end), s.Gen.StoreOrder.Paid.Eq(define.StoreOrderPaidValid)).Count()
}

func (s *StoreOrderRepository) FindPayOrderAmountByDate(ctx context.Context, start int64, end int64) (data decimal.Decimal, err error) {
	storeOrder, err := s.Gen.WithContext(ctx).StoreOrder.Select(s.Gen.StoreOrder.PayPrice).
		Where(s.Gen.StoreOrder.CreatedAt.Between(start, end), s.Gen.StoreOrder.Paid.Eq(define.StoreOrderPaidValid)).
		First()

	if storeOrder != nil {
		data = storeOrder.PayPrice
	}

	return data, err
}

func (s *StoreOrderRepository) FindOrderGroupByDate(ctx context.Context, start int64, end int64) (data []*model.StoreOrder, err error) {
	return s.Gen.StoreOrder.WithContext(ctx).Debug().QueryOrderGroupByDate(&model_data.DateCondition{Start: start, End: end})
}
