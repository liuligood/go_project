package home

import (
	"crmeb_go/internal"
	"crmeb_go/internal/controller/home_controller"
	"github.com/gin-gonic/gin"
)

func Register(app *gin.RouterGroup, appCxt *internal.AppContent) {
	// 获取微服务token
	app.GET("index", home_controller.IndexDate(appCxt.Service))
	app.GET("chart/order", home_controller.ChartOrder(appCxt.Service))
	app.GET("chart/order/week", home_controller.ChartOrderInWeek(appCxt.Service))
	app.GET("chart/order/month", home_controller.ChartOrderInMonth(appCxt.Service))
	app.GET("chart/order/year", home_controller.ChartOrderInYear(appCxt.Service))
}
