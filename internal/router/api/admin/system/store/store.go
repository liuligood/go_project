package store

import (
	"crmeb_go/internal"
	"crmeb_go/internal/controller/system_store_staff_controller"
	"github.com/gin-gonic/gin"
)

func Register(app *gin.RouterGroup, appCxt *internal.AppContent) {
	app.GET("staff/list", system_store_staff_controller.SystemStaffList(appCxt.Service)) // 表单配置根据key获取
}
