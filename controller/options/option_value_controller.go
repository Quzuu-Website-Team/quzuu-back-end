package options

import (
	"github.com/gin-gonic/gin"
	"godp.abdanhafidz.com/controller"
	"godp.abdanhafidz.com/models"
	"godp.abdanhafidz.com/services"
)

func List(c *gin.Context) {
	options := services.OptionValueService{}
	optionValueController := controller.Controller[any, models.OptionCategory, models.Options]{
		Service: &options.Service,
	}
	slug := c.Param("slug")
	options.Constructor.OptionSlug = slug
	options.Retrieve()
	optionValueController.Response(c)
}
