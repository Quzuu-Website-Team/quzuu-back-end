package event

import (
	"github.com/gin-gonic/gin"
	"godp.abdanhafidz.com/controller"
	"godp.abdanhafidz.com/models"
	"godp.abdanhafidz.com/repositories"
	"godp.abdanhafidz.com/services"
	"godp.abdanhafidz.com/utils"
	"strconv"
)

func GetAllEvent(c *gin.Context) {
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil {
		service := services.Service[any, any]{
			Exception: models.Exception{
				Message: "Invalid limit parameter",
			},
		}
		utils.SendResponse(c, service)
		return
	}

	offset, err := strconv.Atoi(c.DefaultQuery("offset", "0"))
	if err != nil {
		service := services.Service[any, any]{
			Exception: models.Exception{
				Message: "Invalid offset parameter",
			},
		}
		utils.SendResponse(c, service)
		return
	}
	filter := c.DefaultQuery("filter", "")
	filterBy := c.DefaultQuery("filter_by", "")

	pagination := repositories.PaginationConstructor{
		Limit:    limit,
		Offset:   offset,
		Filter:   filter,
		FilterBy: filterBy,
	}

	eventsService := services.GetAllEventService{}
	getAllEventController := controller.Controller[any, models.Events, []models.Events]{
		Service: &eventsService.Service,
	}

	eventsService.GetAllEventPaginate(pagination)

	getAllEventController.Response(c)
}
