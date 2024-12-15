package router

import (
	"crmeb_go/internal"
	"crmeb_go/internal/middleware"
	"crmeb_go/internal/router/api"
	"crmeb_go/utils/izap"
	"github.com/gin-gonic/gin"
	"time"
)

func Register(app *gin.Engine, appCxt *internal.AppContent) {
	app.Use(
		middleware.SessionMiddleware(appCxt.Svc),
		middleware.GinZap(izap.Log, time.RFC3339, true),
		middleware.RecoveryWithZap(izap.Log, true),
	)

	api.ApiRegister(app.Group("/api"), appCxt)
}
