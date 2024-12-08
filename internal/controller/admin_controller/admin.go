package admin_controller

import (
	"crmeb_go/internal/container/service"
	service_data "crmeb_go/internal/data/sevice_data"
	"crmeb_go/internal/validation"
	"crmeb_go/utils/ihttp"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetLoginPic(svc *service.Container) gin.HandlerFunc {
	return func(c *gin.Context) {
		var param validation.GetLoginPicParams

		if err := c.ShouldBindQuery(&param); err != nil {
			c.JSON(http.StatusOK, ihttp.Error(err))
			return
		}

		params := service_data.GetLoginPicParams{
			Temp: param.Temp,
		}

		params.SetSessionContext(c)

		res, err := svc.AdminService.GetLoginPic(params)
		if err != nil {
			c.JSON(http.StatusOK, ihttp.Error(err))
			return
		}

		c.JSON(http.StatusOK, ihttp.Data(res))
	}
}
