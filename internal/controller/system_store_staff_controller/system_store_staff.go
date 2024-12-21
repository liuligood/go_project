package system_store_staff_controller

import (
	"crmeb_go/internal/container/service"
	"crmeb_go/internal/data/request"
	"crmeb_go/internal/validation"
	"crmeb_go/utils/ihttp"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SystemStaffList(svc *service.Container) gin.HandlerFunc {
	return func(c *gin.Context) {
		var param validation.SystemStaffListParam
		if err := c.ShouldBind(&param); err != nil {
			c.JSON(http.StatusOK, ihttp.Error(err))
		}

		params := request.SystemStaffListParams{
			StoreId:  param.StoreId,
			Page:     param.Page,
			PageSize: param.PageSize,
		}
		params.BaseServiceParams.SetSessionContext(c)

		res, err := svc.SystemStoreStaffService.SystemStoreStaffList(&params)
		if err != nil {
			c.JSON(http.StatusBadRequest, ihttp.Error(err))
			return
		}

		c.JSON(http.StatusOK, ihttp.Paginate(res.List, res.Count, params.Page, params.PageSize))
	}
}
