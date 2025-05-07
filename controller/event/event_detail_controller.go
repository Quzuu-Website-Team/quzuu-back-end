package event

import (
	"github.com/gin-gonic/gin"
	"godp.abdanhafidz.com/controller"
	"godp.abdanhafidz.com/models"
	"godp.abdanhafidz.com/services"
)

func EventDetail(c *gin.Context) {
	eventDetail := services.EventDetailService{}
	eventDetailController := controller.Controller[any, models.Events, models.EventDetailResponse]{
		Service: &eventDetail.Service,
	}

	eventDetailController.HeaderParse(c, func() {
		eventDetailController.Service.Constructor.Slug = c.Param("event_slug")
		eventDetail.Retrieve(eventDetailController.AccountData.UserID)
		eventDetailController.Response(c)
	})
}
