package home_service

import (
	"crmeb_go/internal/data/request"
	"crmeb_go/internal/data/response"
	"crmeb_go/internal/server"
	"crmeb_go/internal/service/store_order_service"
	"crmeb_go/internal/service/user_service"
	"crmeb_go/internal/service/user_visit_record_service"
	"time"
)

type HomeServiceImpl interface {
	IndexDate(params *request.BaseServiceParams) (data *response.HomeRateResp, err error)
	ChartOrder(params *request.BaseServiceParams) (data *map[string]interface{}, err error)
}

type HomeService struct {
	svc *server.SvcContext
}

func NewHomeService(svc *server.SvcContext) *HomeService {
	return &HomeService{svc: svc}
}

// IndexDate 首页数据
func (h *HomeService) IndexDate(params *request.BaseServiceParams) (data *response.HomeRateResp, err error) {
	var resp response.HomeRateResp
	now := time.Now()
	// 今天0点时间
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	todayUnix := today.Unix()
	// 计算昨天0点的时间
	yesterdayZero := today.Add(-24 * time.Hour).Unix()
	// 计算昨天24点的时间
	yesterdayEnd := today.Add(-1 * time.Second).Unix()
	// 现在的时间
	nowUnix := time.Now().Unix() - 1
	dateParams := &request.DateParams{
		BaseServiceParams: *params,
		Start:             todayUnix,
		End:               nowUnix,
		YesterdayStart:    yesterdayZero,
		YesterdayEnd:      yesterdayEnd,
	}

	// 取今天，咋日访问量
	visitRecordService := user_visit_record_service.NewUserVisitRecordService(h.svc)
	pageViews, err := visitRecordService.FindPageViewsByDate(dateParams)
	if err != nil {
		return &resp, err
	}

	// 今日咋日注册用户数量
	userService := user_service.NewUserService(h.svc)
	registerNum, err := userService.GetRegisterNumByDate(dateParams)
	if err != nil {
		return &resp, err
	}

	// 今日咋日注册用户数量
	storeOrderService := store_order_service.NewStoreOrderService(h.svc)
	orderNum, err := storeOrderService.GetOrderNumByDate(dateParams)
	if err != nil {
		return &resp, err
	}

	// 今日咋日销售额
	payOrderAmount, err := storeOrderService.GetPayOrderAmountByDate(dateParams)
	if err != nil {
		return &resp, err
	}

	resp.PageViews = pageViews.NowData
	resp.YesterdayPageViews = pageViews.YesterdayData
	resp.NewUserNum = registerNum.NowData
	resp.YesterdayNewUserNum = registerNum.YesterdayData
	resp.OrderNum = orderNum.NowData
	resp.YesterdayOrderNum = orderNum.YesterdayData
	resp.Sale = payOrderAmount.NowSale
	resp.YesterdaySale = payOrderAmount.YesterdaySale
	return &resp, nil
}

// ChartOrder 30天订单趋势
func (h *HomeService) ChartOrder(params *request.BaseServiceParams) (data *map[string]interface{}, err error) {
	storeOrderService := store_order_service.NewStoreOrderService(h.svc)
	storeOrderService.GetOrderGroupByDate()
}
