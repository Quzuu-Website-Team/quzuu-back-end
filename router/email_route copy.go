package router

import (
	"github.com/gin-gonic/gin"
	EmailController "godp.abdanhafidz.com/controller/email"
)

func EmailRoute(router *gin.Engine) {
	routerGroup := router.Group("/api/v1/email")
	{
		routerGroup.POST("/verify", EmailController.Verify)
		routerGroup.POST("/create-verification", EmailController.CreateVerification)
	}
}
