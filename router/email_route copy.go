package router

import (
	"github.com/gin-gonic/gin"
	EmailController "godp.abdanhafidz.com/controller/email"
	"godp.abdanhafidz.com/middleware"
)

func EmailRoute(router *gin.Engine) {
	routerGroup := router.Group("/api/v1/email")
	{
		routerGroup.POST("/verify", middleware.AuthUser, EmailController.Verify)
		routerGroup.POST("/create-verification", middleware.AuthUser, EmailController.CreateVerification)
	}
}
