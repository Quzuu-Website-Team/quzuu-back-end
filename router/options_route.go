package router

import (
	"github.com/gin-gonic/gin"
	OptionsController "godp.abdanhafidz.com/controller/options"
	CityController "godp.abdanhafidz.com/controller/region/city"
	ProvinceController "godp.abdanhafidz.com/controller/region/province"
)

func OptionsRoute(router *gin.Engine) {
	routerGroup := router.Group("/api/v1/options")
	{
		routerGroup.POST("/create", OptionsController.AddOptions)
		routerGroup.GET("/list/:slug", OptionsController.List)
		routerGroup.GET("/region/provinces", ProvinceController.List)
		routerGroup.GET("/region/cities", CityController.List)
		routerGroup.POST("/region/seed-provinces", ProvinceController.Seeds)
		routerGroup.POST("/region/seed-cities", CityController.Seeds)
	}
}
