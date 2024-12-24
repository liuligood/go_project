package home_service

import (
	"crmeb_go/define"
	"crmeb_go/internal/data/request"
	"crmeb_go/internal/data/response"
	"crmeb_go/internal/model/model_data"
	"crmeb_go/internal/server"
	"crmeb_go/internal/service/store_order_service"
	"crmeb_go/internal/service/user_service"
	"crmeb_go/internal/service/user_visit_record_service"
	"crmeb_go/utils/itime"
	"errors"
	"time"
)

type HomeServiceImpl interface {
	IndexDate(params *request.BaseServiceParams) (data *response.HomeRateResp, err error)
	ChartOrder(params *request.BaseServiceParams) (data *map[string]interface{}, err error)
	ChartOrderInWeek(params *request.BaseServiceParams) (data *map[string]interface{}, err error)
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
	everyDateResp, err := storeOrderService.GetOrderGroupByDate(&request.SearchDateParams{BaseServiceParams: *params, Date: define.AdminSearchDateLatelyThirty})
	if err != nil {
		return nil, err
	}

	listDate := itime.GetListDate(define.AdminSearchDateLatelyThirty)
	resp := make(map[string]interface{}, len(everyDateResp))
	priceMap, idMap, err := h.getPriceAndIdMap(everyDateResp, define.AdminSearchDateLatelyThirty)
	if err != nil {
		return nil, err
	}

	h.setValue(listDate, priceMap)
	h.setValue(listDate, idMap)
	resp["price"] = priceMap
	resp["quality"] = idMap
	return &resp, nil
}

// ChartOrderInWeek 周订单量趋势
func (h *HomeService) ChartOrderInWeek(params *request.BaseServiceParams) (data *map[string]interface{}, err error) {
	storeOrderService := store_order_service.NewStoreOrderService(h.svc)
	listDate := itime.GetListDate(define.AdminSearchDateWeek)

	// 本周
	weekData, err := storeOrderService.GetOrderGroupByDate(&request.SearchDateParams{BaseServiceParams: *params, Date: define.AdminSearchDateWeek})
	if err != nil {
		return nil, err
	}

	resp := make(map[string]interface{}, len(weekData))
	priceMap, idMap, err := h.getPriceAndIdMap(weekData, define.AdminSearchDateWeek)
	if err != nil {
		return nil, err
	}

	h.setValue(listDate, priceMap)
	h.setValue(listDate, idMap)

	// 上周
	preWeekData, err := storeOrderService.GetOrderGroupByDate(&request.SearchDateParams{BaseServiceParams: *params, Date: define.AdminSearchDatePreWeek})
	if err != nil {
		return nil, err
	}

	prePriceMap, preIdMap, err := h.getPriceAndIdMap(preWeekData, define.AdminSearchDateWeek)
	if err != nil {
		return nil, err
	}

	h.setValue(listDate, prePriceMap)
	h.setValue(listDate, preIdMap)
	resp["prePrice"] = prePriceMap
	resp["preQuality"] = preIdMap
	resp["price"] = priceMap
	resp["quality"] = idMap
	return &resp, nil
}

func (h *HomeService) setValue(listDate []string, priceMap map[string]interface{}) {
	for _, v := range listDate {
		if _, ok := priceMap[v]; ok {
			continue
		}
		priceMap[v] = 0
	}
}

func (h *HomeService) getPriceAndIdMap(weekData []*model_data.EveryDateResp, data string) (map[string]interface{}, map[string]interface{}, error) {
	priceMap := make(map[string]interface{}, len(weekData))
	idMap := make(map[string]interface{}, len(weekData))
	for _, v := range weekData {
		switch data {
		case define.AdminSearchDateLatelyThirty:
			parse, err := time.Parse(time.RFC3339, v.EveryDate)
			if err != nil {
				return nil, nil, errors.New("解析日期时间失败")
			}
			formatDate := parse.Format(define.SystemTimeMonthDayFormat)
			priceMap[formatDate] = v.PayPrice
			idMap[formatDate] = v.ID

		case define.AdminSearchDateWeek:
			parse, err := time.Parse(time.RFC3339, v.EveryDate)
			if err != nil {
				return nil, nil, errors.New("解析日期时间失败")
			}
			// 获取星期几
			weekday := parse.Weekday()
			weekdays := []string{"星期一", "星期二", "星期三", "星期四", "星期五", "星期六", "星期日"}
			weekdayStr := weekdays[weekday-1]
			priceMap[weekdayStr] = v.PayPrice
			idMap[weekdayStr] = v.ID
		}
	}

	return priceMap, idMap, nil
}
