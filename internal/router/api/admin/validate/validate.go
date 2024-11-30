package validate

import (
	"crmeb_go/internal"
	"crmeb_go/internal/controller/validate_controller"
	"github.com/gin-gonic/gin"
)

func Register(app *gin.RouterGroup, appCxt *internal.AppContent) {
	// 获取验证码
	app.GET("code", validate_controller.GetValidateCode(appCxt.Service))
}
