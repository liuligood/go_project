package system_config_controller

import (
	"crmeb_go/internal/container/service"
	"crmeb_go/internal/data/request"
	"crmeb_go/internal/validation"
	"crmeb_go/utils/ihttp"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUniq(svc *service.Container) gin.HandlerFunc {
	return func(c *gin.Context) {
		var param validation.GetUniq
		if err := c.ShouldBindQuery(&param); err != nil {
			c.JSON(http.StatusOK, ihttp.Error(err))
		}

		params := request.GetSystemConfigParams{
			Name: param.Key,
		}
		params.BaseServiceParams.SetSessionContext(c)

		res, err := svc.SystemConfigService.GetSystemConfigInfo(&params)
		if err != nil {
			c.JSON(http.StatusBadRequest, ihttp.Error(err))
			return
		}

		c.JSON(http.StatusOK, ihttp.Data(res))
	}
}
