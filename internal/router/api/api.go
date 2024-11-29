package api

import (
	"crmeb_go/internal"
	"crmeb_go/internal/middleware"
	"crmeb_go/internal/router/api/user"
	"github.com/gin-gonic/gin"
)

// 静默授权api
func ApiRegister(app *gin.RouterGroup, appCxt *internal.AppContent) {

	// 中间件：鉴权
	app.Use(middleware.ApiAuthVisitorMiddleWare(appCxt.Svc))

	// 用户路由
	user.Register(app.Group("/user"), appCxt)

}
