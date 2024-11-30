package admin

import (
	"crmeb_go/internal"
	"crmeb_go/internal/controller/user_controller"
	"crmeb_go/internal/router/api/admin/validate"
	"github.com/gin-gonic/gin"
)

func Register(app *gin.RouterGroup, appCxt *internal.AppContent) {
	// 获取微服务token
	app.GET("login", user_controller.GetRealName(appCxt.Service))

	// 授权路由
	validate.Register(app.Group("/validate"), appCxt)
}
