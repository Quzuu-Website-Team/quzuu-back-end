package router

import (
	"github.com/gin-gonic/gin"
	EmailController "godp.abdanhafidz.com/controller/email"
)

func EmailRoute(router *gin.Engine) {
	routerGroup := router.Group("/api/v1/emails")
	{
		routerGroup.POST("/verify", EmailController.CreateVerification)
		routerGroup.POST("/create-verification", EmailController.CreateVerification)
		routerGroup.DELETE("/delete-verification", EmailController.DeleteVerification)
	}
}
