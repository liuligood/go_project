package router

import (
	"crmeb_go/internal"
	"crmeb_go/internal/middleware"
	"crmeb_go/internal/router/api"
	"github.com/gin-gonic/gin"
)

func Register(app *gin.Engine, appCxt *internal.AppContent) {
	app.Use(
		middleware.SessionMiddleware(appCxt.Svc),
	)

	// api接口：鉴权中间件
	api.ApiRegister(app.Group("/api"), appCxt)
}
