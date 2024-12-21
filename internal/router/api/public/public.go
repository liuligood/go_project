package public

import (
	"crmeb_go/internal"
	"crmeb_go/internal/controller/public_controller"
	"github.com/gin-gonic/gin"
)

func Register(app *gin.RouterGroup, appCxt *internal.AppContent) {
	// 获取微服务token
	app.GET("jsconfig/getcrmebchatconfig", public_controller.GetCrmebChatConfig(appCxt.Service))
}
