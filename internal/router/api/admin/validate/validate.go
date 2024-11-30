package validate

import (
	"crmeb_go/internal"
	"crmeb_go/internal/controller/validate_controller"
	"github.com/gin-gonic/gin"
)

func Register(app *gin.RouterGroup, appCxt *internal.AppContent) {
	// 获取微服务token
	app.GET("validate/code", validate_controller.GetValidateCode(appCxt.Service))
}
