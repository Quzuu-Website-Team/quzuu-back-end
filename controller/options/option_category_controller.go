package options

import (
	"github.com/gin-gonic/gin"
	"godp.abdanhafidz.com/controller"
	"godp.abdanhafidz.com/models"
	"godp.abdanhafidz.com/services"
)

func AddOptions(c *gin.Context) {
	options := services.OptionService{}
	addOptionController := controller.Controller[[]models.OptionsRequest, []models.OptionsRequest, models.OptionsResponse]{
		Service: &options.Service,
	}
	addOptionController.RequestJSON(c, func() {
		options.Constructor = addOptionController.Request
		options.Create()
	})
}
