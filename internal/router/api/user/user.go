package user

import (
	"crmeb_go/internal"
	"crmeb_go/internal/controller/user_controller"
	"github.com/gin-gonic/gin"
)

func Register(app *gin.RouterGroup, appCxt *internal.AppContent) {
	// 获取微服务token
	app.GET("test", user_controller.GetRealName(appCxt.Service))
}
