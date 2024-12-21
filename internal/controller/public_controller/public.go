package public_controller

import (
	"crmeb_go/define"
	"crmeb_go/internal/container/service"
	"crmeb_go/internal/data/request"
	"crmeb_go/utils/ihttp"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetCrmebChatConfig(svc *service.Container) gin.HandlerFunc {
	return func(c *gin.Context) {
		params := request.GetSystemConfigParams{
			Name: define.AdminChatStatistics,
		}
		params.BaseServiceParams.SetSessionContext(c)

		res, err := svc.SystemConfigService.GetSystemConfigInfo(&params)
		if err != nil {
			c.JSON(http.StatusBadRequest, ihttp.Error(err))
			return
		}

		c.JSON(http.StatusOK, ihttp.Data(res.Name))
	}
}
