package router

import (
	"github.com/gin-gonic/gin"
	"godp.abdanhafidz.com/controller/event"
)

func EventRoute(router *gin.Engine) {
	routerGroup := router.Group("api/v1/event")
	{
		routerGroup.GET("/", event.GetAllEvent)
	}
}
