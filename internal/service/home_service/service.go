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
	"github.com/iancoleman/orderedmap"
	"time"
)

type HomeServiceImpl interface {
	IndexDate(params *request.BaseServiceParams) (data *response.HomeRateResp, err error)
	ChartOrder(params *request.BaseServiceParams) (data *response.ChartOrder, err error)
	ChartOrderInWeek(params *request.BaseServiceParams) (data *response.ChartOrder, err error)
	ChartOrderInMonth(params *request.BaseServiceParams) (data *response.ChartOrder, err error)
	ChartOrderInYear(params *request.BaseServiceParams) (data *response.ChartOrder, err error)
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

	// 今日咋日订单数量
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
func (h *HomeService) ChartOrder(params *request.BaseServiceParams) (data *response.ChartOrder, err error) {
	storeOrderService := store_order_service.NewStoreOrderService(h.svc)
	var resp response.ChartOrder
	listDate := itime.GetListDate(define.AdminSearchDateLatelyThirty)

	everyDateResp, err := storeOrderService.GetOrderGroupByDate(&request.SearchDateParams{BaseServiceParams: *params, Date: define.AdminSearchDateLatelyThirty})
	if err != nil {
		return nil, err
	}

	priceMap, idMap, err := h.getPriceAndIdMap(everyDateResp, define.AdminSearchDateLatelyThirty, listDate)
	if err != nil {
		return nil, err
	}

	resp.Price = priceMap
	resp.Quality = idMap
	return &resp, nil
}

// ChartOrderInWeek 周订单量趋势
func (h *HomeService) ChartOrderInWeek(params *request.BaseServiceParams) (data *response.ChartOrder, err error) {
	storeOrderService := store_order_service.NewStoreOrderService(h.svc)
	listDate := itime.GetListDate(define.AdminSearchDateWeek)
	var resp response.ChartOrder

	// 本周
	weekData, err := storeOrderService.GetOrderGroupByDate(&request.SearchDateParams{BaseServiceParams: *params, Date: define.AdminSearchDateWeek})
	if err != nil {
		return nil, err
	}

	priceMap, idMap, err := h.getPriceAndIdMap(weekData, define.AdminSearchDateWeek, listDate)
	if err != nil {
		return nil, err
	}

	// 上周
	preWeekData, err := storeOrderService.GetOrderGroupByDate(&request.SearchDateParams{BaseServiceParams: *params, Date: define.AdminSearchDatePreWeek})
	if err != nil {
		return nil, err
	}

	prePriceMap, preIdMap, err := h.getPriceAndIdMap(preWeekData, define.AdminSearchDateWeek, listDate)
	if err != nil {
		return nil, err
	}

	resp.PrePrice = prePriceMap
	resp.PreQuality = preIdMap
	resp.Price = priceMap
	resp.Quality = idMap
	return &resp, nil
}

func (h *HomeService) setValue(listDate []string, priceMap *orderedmap.OrderedMap) {
	for _, v := range listDate {
		if _, ok := priceMap.Get(v); ok {
			continue
		}
		priceMap.Set(v, 0)
	}
}

func (h *HomeService) getPriceAndIdMap(weekData []*model_data.EveryDateResp, data string, listDate []string) (*orderedmap.OrderedMap, *orderedmap.OrderedMap, error) {
	priceMap := orderedmap.New()
	idMap := orderedmap.New()
	h.setValue(listDate, priceMap)
	h.setValue(listDate, idMap)
	for _, v := range weekData {
		switch data {
		case define.AdminSearchDateLatelyThirty:
			parse, err := time.Parse(time.RFC3339, v.EveryDate)
			if err != nil {
				return nil, nil, errors.New("解析日期时间失败")
			}

			formatDate := parse.Format(define.SystemTimeMonthDayFormat)
			priceMap.Set(formatDate, v.PayPrice)
			idMap.Set(formatDate, v.ID)
		case define.AdminSearchDateWeek:
			parse, err := time.Parse(time.RFC3339, v.EveryDate)
			if err != nil {
				return nil, nil, errors.New("解析日期时间失败")
			}

			// 获取星期几
			weekday := parse.Weekday()
			weekdayList := []string{"星期日", "星期一", "星期二", "星期三", "星期四", "星期五", "星期六"}
			weekdayStr := weekdayList[weekday]
			priceMap.Set(weekdayStr, v.PayPrice)
			idMap.Set(weekdayStr, v.ID)
		case define.AdminSearchDateMonth:
			parse, err := time.Parse(time.RFC3339, v.EveryDate)
			if err != nil {
				return nil, nil, errors.New("解析日期时间失败")
			}

			formatDate := parse.Format(define.SystemTimeDayFormat)
			priceMap.Set(formatDate, v.PayPrice)
			idMap.Set(formatDate, v.ID)
		case define.AdminSearchDatePreMonth:
			parse, err := time.Parse(time.RFC3339, v.EveryDate)
			if err != nil {
				return nil, nil, errors.New("解析日期时间失败")
			}

			formatDate := parse.Format(define.SystemTimeDayFormat)
			priceMap.Set(formatDate, v.PayPrice)
			idMap.Set(formatDate, v.ID)
		case define.AdminSearchDateYear:
			parse, err := time.Parse(time.RFC3339, v.EveryDate)
			if err != nil {
				return nil, nil, errors.New("解析日期时间失败")
			}
			// 获取是当月几号
			month := parse.Month()
			monthList := []string{"一月", "二月", "三月", "四月", "五月", "六月", "七月", "八月", "九月", "十月", "十一月", "十二月"}
			monthStr := monthList[month-1]
			priceMap.Set(monthStr, v.PayPrice)
			idMap.Set(monthStr, v.ID)
		}
	}

	return priceMap, idMap, nil
}

// ChartOrderInMonth 本月上月订单量趋势
func (h *HomeService) ChartOrderInMonth(params *request.BaseServiceParams) (data *response.ChartOrder, err error) {
	storeOrderService := store_order_service.NewStoreOrderService(h.svc)
	listDate := itime.GetListDate(define.AdminSearchDateMonth)
	var resp response.ChartOrder

	// 本月
	monthData, err := storeOrderService.GetOrderGroupByDate(&request.SearchDateParams{BaseServiceParams: *params, Date: define.AdminSearchDateMonth})
	if err != nil {
		return nil, err
	}

	priceMap, idMap, err := h.getPriceAndIdMap(monthData, define.AdminSearchDateMonth, listDate)

	if err != nil {
		return nil, err
	}

	// 上个月
	preMonthData, err := storeOrderService.GetOrderGroupByDate(&request.SearchDateParams{BaseServiceParams: *params, Date: define.AdminSearchDatePreMonth})
	if err != nil {
		return nil, err
	}

	preListDate := itime.GetListDate(define.AdminSearchDatePreMonth)
	prePriceMap, preIdMap, err := h.getPriceAndIdMap(preMonthData, define.AdminSearchDatePreMonth, preListDate)
	if err != nil {
		return nil, err
	}

	resp.PrePrice = prePriceMap
	resp.PreQuality = preIdMap
	resp.Price = priceMap
	resp.Quality = idMap
	return &resp, nil
}

// ChartOrderInYear 今年去年订单量趋势
func (h *HomeService) ChartOrderInYear(params *request.BaseServiceParams) (data *response.ChartOrder, err error) {
	storeOrderService := store_order_service.NewStoreOrderService(h.svc)
	listDate := itime.GetListDate(define.AdminSearchDateYear)
	var resp response.ChartOrder

	// 今年
	monthData, err := storeOrderService.GetOrderGroupByDate(&request.SearchDateParams{BaseServiceParams: *params, Date: define.AdminSearchDateYear})
	if err != nil {
		return nil, err
	}

	priceMap, idMap, err := h.getPriceAndIdMap(monthData, define.AdminSearchDateYear, listDate)

	if err != nil {
		return nil, err
	}

	// 去年
	preMonthData, err := storeOrderService.GetOrderGroupByDate(&request.SearchDateParams{BaseServiceParams: *params, Date: define.AdminSearchDatePreYear})
	if err != nil {
		return nil, err
	}

	prePriceMap, preIdMap, err := h.getPriceAndIdMap(preMonthData, define.AdminSearchDateYear, listDate)
	if err != nil {
		return nil, err
	}

	resp.PrePrice = prePriceMap
	resp.PreQuality = preIdMap
	resp.Price = priceMap
	resp.Quality = idMap
	return &resp, nil
}
