package config

import (
	"crmeb_go/internal"
	"crmeb_go/internal/controller/system_config_controller"
	"github.com/gin-gonic/gin"
)

func Register(app *gin.RouterGroup, appCxt *internal.AppContent) {
	app.GET("/getuniq", system_config_controller.GetUniq(appCxt.Service)) // 表单配置根据key获取
}
