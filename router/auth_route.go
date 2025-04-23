package router

import (
	"github.com/gin-gonic/gin"
	AuthController "godp.abdanhafidz.com/controller/auth"
	"godp.abdanhafidz.com/middleware"
)

func AuthRoute(router *gin.Engine) {
	routerGroup := router.Group("/api/v1/auth")
	{
		routerGroup.POST("/external-login", AuthController.ExternalAuth)
		routerGroup.POST("/login", AuthController.Login)
		routerGroup.POST("/register", AuthController.Register)
		routerGroup.PUT("/change-password", middleware.AuthUser, AuthController.ChangePassword)
		routerGroup.POST("/forgot-password", AuthController.CreateForgotPassword)
		routerGroup.PUT("/forgot-password", AuthController.ValidateForgotPassword)
	}
}
