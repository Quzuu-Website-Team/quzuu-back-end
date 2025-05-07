package city

import (
	"github.com/gin-gonic/gin"
	"godp.abdanhafidz.com/controller"
	"godp.abdanhafidz.com/models"
	"godp.abdanhafidz.com/services"
)

func Seeds(c *gin.Context) {
	city := services.CityService{}
	CityController := controller.Controller[any, models.RegionCity, []models.RegionCity]{
		Service: &city.Service,
	}
	city.Create()
	CityController.Response(c)
}
