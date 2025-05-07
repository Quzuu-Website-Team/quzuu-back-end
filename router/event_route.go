package router

import (
	"github.com/gin-gonic/gin"
	"godp.abdanhafidz.com/controller/event"
	"godp.abdanhafidz.com/middleware"
)

func EventRoute(router *gin.Engine) {
	routerGroup := router.Group("api/v1/events")
	{
		routerGroup.GET("/", event.EventList)
		routerGroup.GET("/details/:id_event", middleware.AuthUser, event.EventDetail)
		routerGroup.POST("/register-event", middleware.AuthUser, event.Register)
	}
}
