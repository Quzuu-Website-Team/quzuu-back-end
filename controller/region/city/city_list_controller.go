package city

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"godp.abdanhafidz.com/controller"
	"godp.abdanhafidz.com/models"
	"godp.abdanhafidz.com/services"
)

func List(c *gin.Context) {
	city := services.CityService{}
	CityController := controller.Controller[any, models.RegionCity, []models.RegionCity]{
		Service: &city.Service,
	}
	ProvinceID, _ := strconv.Atoi(c.Query("province_id"))
	city.Constructor.ProvinceId = uint(ProvinceID)
	city.Retrieve()
	CityController.Response(c)
}
