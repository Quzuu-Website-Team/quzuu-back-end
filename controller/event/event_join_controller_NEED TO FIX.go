package event

import (
	"github.com/gin-gonic/gin"
	"godp.abdanhafidz.com/controller"
	"godp.abdanhafidz.com/models"
	"godp.abdanhafidz.com/services"
)

func JoinEvent(c *gin.Context) {
	eventAssign := services.JoinEventService{}
	eventAssignController := controller.Controller[models.JoinEventRequest, models.JoinEventRequest, models.EventDetailResponse]{
		Service: &eventAssign.Service,
	}

	eventAssignController.RequestJSON(c, func() {
		eventAssignController.Service.Constructor.EventId = eventAssignController.Request.EventId
		eventAssignController.Service.Constructor.EventCode = eventAssignController.Request.EventCode
		idUser := eventAssignController.AccountData.UserID
		eventAssign.Create(idUser)
	})
}
