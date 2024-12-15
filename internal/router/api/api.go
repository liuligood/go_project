package api

import (
	"crmeb_go/internal"
	"crmeb_go/internal/middleware"
	"crmeb_go/internal/router/api/admin"
	"crmeb_go/internal/router/api/user"
	"github.com/gin-gonic/gin"
)

func ApiRegister(app *gin.RouterGroup, appCxt *internal.AppContent) {

	// 中间件：鉴权
	app.Use(middleware.ApiAuthVisitorMiddleWare(appCxt.Svc))

	// 用户路由
	user.Register(app.Group("/user"), appCxt)

	// 授权路由
	admin.Register(app.Group("/admin"), appCxt)
}
