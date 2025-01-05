package product_controller

import (
	"crmeb_go/internal/container/service"
	"crmeb_go/internal/data/request"
	"crmeb_go/internal/validation"
	"crmeb_go/utils/ihttp"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateProduct(svc *service.Container) gin.HandlerFunc {
	return func(c *gin.Context) {
		var param validation.CreateProductParam
		if err := c.ShouldBindJSON(&param); err != nil {
			c.JSON(http.StatusOK, ihttp.Error(err))
		}

		params := request.CreateProductParams{
			CreateProduct: &param,
		}
		params.SetSessionContext(c)

		res, err := svc.ProductService.CreateProduct(&params)
		if err != nil {
			c.JSON(http.StatusOK, ihttp.Error(err))

			return
		}

		c.JSON(http.StatusOK, ihttp.Data(res))
	}
}
