package store_order_service

import (
	"crmeb_go/internal/data/request"
	"crmeb_go/internal/data/response"
	"crmeb_go/internal/server"
	"crmeb_go/utils/itime"
	"crmeb_go/utils/izap"
	"errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type StoreOrderServiceImpl interface {
	GetOrderNumByDate(params *request.DateParams) (*response.DateResp, error)
	GetPayOrderAmountByDate(params *request.DateParams) (*response.DateResp, error)
}

type StoreOrderService struct {
	svc *server.SvcContext
}

func NewStoreOrderService(svc *server.SvcContext) *StoreOrderService {
	return &StoreOrderService{svc: svc}
}

func (u *StoreOrderService) GetOrderNumByDate(params *request.DateParams) (*response.DateResp, error) {
	var data response.DateResp
	// 今日订单数量
	nowPageViews, err := u.svc.Repo.StoreOrderRepository.FindOrderNumByDate(params.BaseServiceParams.Ctx, params.Start, params.End)
	if err != nil {
		izap.Log.Error("查询今日访问量失败", zap.String("开始时间", itime.Format(params.Start)), zap.String("结束时间", itime.Format(params.End)), zap.Error(err))

		return &data, err
	}

	// 昨天订单数量
	yesterdayPageViews, err := u.svc.Repo.StoreOrderRepository.FindOrderNumByDate(params.BaseServiceParams.Ctx, params.YesterdayStart, params.YesterdayEnd)
	if err != nil {
		izap.Log.Error("查询今日访问量失败", zap.String("开始时间", itime.Format(params.YesterdayStart)), zap.String("结束时间", itime.Format(params.YesterdayEnd)), zap.Error(err))

		return &data, err
	}

	data.NowData = nowPageViews
	data.YesterdayData = yesterdayPageViews
	return &data, nil
}

func (u *StoreOrderService) GetPayOrderAmountByDate(params *request.DateParams) (*response.DateResp, error) {
	var data response.DateResp
	// 今日销售额
	nowSale, err := u.svc.Repo.StoreOrderRepository.FindPayOrderAmountByDate(params.BaseServiceParams.Ctx, params.Start, params.End)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		izap.Log.Error("查询今日访问量失败", zap.String("开始时间", itime.Format(params.Start)), zap.String("结束时间", itime.Format(params.End)), zap.Error(err))

		return &data, err
	}

	// 昨天销售额
	yesterdaySale, err := u.svc.Repo.StoreOrderRepository.FindPayOrderAmountByDate(params.BaseServiceParams.Ctx, params.YesterdayStart, params.YesterdayEnd)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		izap.Log.Error("查询今日访问量失败", zap.String("开始时间", itime.Format(params.YesterdayStart)), zap.String("结束时间", itime.Format(params.YesterdayEnd)), zap.Error(err))

		return &data, err
	}

	data.NowSale = nowSale
	data.YesterdaySale = yesterdaySale
	return &data, nil
}
