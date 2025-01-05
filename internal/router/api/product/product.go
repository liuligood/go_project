package product

import (
	"crmeb_go/internal"
	"crmeb_go/internal/controller/product_controller"
	"github.com/gin-gonic/gin"
)

func Register(app *gin.RouterGroup, appCxt *internal.AppContent) {
	// 获取微服务token
	app.POST("create", product_controller.CreateProduct(appCxt.Service))
}
