package province

import (
	"github.com/gin-gonic/gin"
	"godp.abdanhafidz.com/controller"
	"godp.abdanhafidz.com/models"
	"godp.abdanhafidz.com/services"
)

func Seeds(c *gin.Context) {
	province := services.ProvinceService{}
	ProvinceController := controller.Controller[any, models.RegionProvince, []models.RegionProvince]{
		Service: &province.Service,
	}
	province.Create()
	ProvinceController.Response(c)
}
