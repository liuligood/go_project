package store_order_service

import (
	"crmeb_go/internal/data/request"
	"crmeb_go/internal/data/response"
	"crmeb_go/internal/model/model_data"
	"crmeb_go/internal/server"
	"crmeb_go/utils/itime"
	"crmeb_go/utils/izap"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type StoreOrderServiceImpl interface {
	GetOrderNumByDate(params *request.DateParams) (*response.DateResp, error)
	GetPayOrderAmountByDate(params *request.DateParams) (*response.DateResp, error)
	GetOrderGroupByDate(params *request.SearchDateParams) ([]*model_data.EveryDateResp, error)
}

type StoreOrderService struct {
	svc *server.SvcContext
}

func NewStoreOrderService(svc *server.SvcContext) *StoreOrderService {
	return &StoreOrderService{svc: svc}
}

func (s *StoreOrderService) GetOrderNumByDate(params *request.DateParams) (*response.DateResp, error) {
	var data response.DateResp
	// 今日订单数量
	nowPageViews, err := s.svc.Repo.StoreOrderRepository.FindOrderNumByDate(params.BaseServiceParams.Ctx, params.Start, params.End)
	if err != nil {
		izap.Log.Error("查询今日访问量失败", zap.String("开始时间", itime.Format(params.Start)), zap.String("结束时间", itime.Format(params.End)), zap.Error(err))

		return &data, err
	}

	// 昨天订单数量
	yesterdayPageViews, err := s.svc.Repo.StoreOrderRepository.FindOrderNumByDate(params.BaseServiceParams.Ctx, params.YesterdayStart, params.YesterdayEnd)
	if err != nil {
		izap.Log.Error("查询今日访问量失败", zap.String("开始时间", itime.Format(params.YesterdayStart)), zap.String("结束时间", itime.Format(params.YesterdayEnd)), zap.Error(err))

		return &data, err
	}

	data.NowData = nowPageViews
	data.YesterdayData = yesterdayPageViews
	return &data, nil
}

func (s *StoreOrderService) GetPayOrderAmountByDate(params *request.DateParams) (*response.DateResp, error) {
	var data response.DateResp
	// 今日销售额
	nowSale, err := s.svc.Repo.StoreOrderRepository.FindPayOrderAmountByDate(params.BaseServiceParams.Ctx, params.Start, params.End)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		izap.Log.Error("查询今日访问量失败", zap.String("开始时间", itime.Format(params.Start)), zap.String("结束时间", itime.Format(params.End)), zap.Error(err))

		return &data, err
	}

	// 昨天销售额
	yesterdaySale, err := s.svc.Repo.StoreOrderRepository.FindPayOrderAmountByDate(params.BaseServiceParams.Ctx, params.YesterdayStart, params.YesterdayEnd)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		izap.Log.Error("查询今日访问量失败", zap.String("开始时间", itime.Format(params.YesterdayStart)), zap.String("结束时间", itime.Format(params.YesterdayEnd)), zap.Error(err))

		return &data, err
	}

	data.NowSale = nowSale
	data.YesterdaySale = yesterdaySale
	return &data, nil
}

func (s *StoreOrderService) GetOrderGroupByDate(params *request.SearchDateParams) (data []*model_data.EveryDateResp, err error) {
	start, end := itime.CalculateDateRange(params.Date)
	// 计算时间范围
	data, err = s.svc.Repo.StoreOrderRepository.FindOrderGroupByDate(params.BaseServiceParams.Ctx, start, end)
	if err != nil {
		izap.Log.Error("根据时间查询订单趋势错误", zap.String("时间范围:", fmt.Sprintf("start:%v, end:%v", start, end)), zap.Error(err))

		return nil, err
	}
	return data, nil
}
