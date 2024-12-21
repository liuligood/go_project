package response

import "github.com/shopspring/decimal"

type HomeRateResp struct {
	Sales               decimal.Decimal `json:"sales"`               // 今日销售额
	YesterdaySales      decimal.Decimal `json:"yesterdaySales"`      // 昨日销售额
	PageViews           int64           `json:"pageviews"`           // 今日访问量
	YesterdayPageViews  int64           `json:"yesterdayPageviews"`  // 昨日访问量
	OrderNum            int64           `json:"orderNum"`            // 今日订单量
	YesterdayOrderNum   int64           `json:"yesterdayOrderNum"`   // 昨日订单量
	NewUserNum          int64           `json:"newUserNum"`          // 今日新增用户
	YesterdayNewUserNum int64           `json:"yesterdayNewUserNum"` // 昨日新增用户
}
