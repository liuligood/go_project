package system

import (
	"crmeb_go/internal"
	"crmeb_go/internal/router/api/admin/system/config"
	"crmeb_go/internal/router/api/admin/system/store"
	"github.com/gin-gonic/gin"
)

func Register(app *gin.RouterGroup, appCxt *internal.AppContent) {
	config.Register(app.Group("/config"), appCxt)
	store.Register(app.Group("/store"), appCxt)
}
