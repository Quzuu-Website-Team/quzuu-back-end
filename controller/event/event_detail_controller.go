package event

import (
	"github.com/gin-gonic/gin"
	"godp.abdanhafidz.com/controller"
	"godp.abdanhafidz.com/models"
	"godp.abdanhafidz.com/services"
)

func EventDetail(c *gin.Context) {
	eventDetail := services.EventDetailService{}
	eventDetailController := controller.Controller[models.EventDetailRequest, models.EventDetailRequest, models.EventDetailResponse]{
		Service: &eventDetail.Service,
	}

	eventDetailController.RequestJSON(c, func() {
		eventDetailController.Service.Constructor.IdEvent = eventDetailController.Request.IdEvent
		eventDetailController.Service.Constructor.IdUser = eventDetailController.Request.IdUser
		eventDetail.Retrieve()
	})
}
