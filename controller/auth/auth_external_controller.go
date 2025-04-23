package auth

import (
	"github.com/gin-gonic/gin"
	"godp.abdanhafidz.com/controller"
	"godp.abdanhafidz.com/models"
	"godp.abdanhafidz.com/services"
)

func ExternalAuth(c *gin.Context) {
	ExternalAuthController := controller.Controller[models.ExternalAuthRequest, models.ExternalAuth, models.AuthenticatedUser]{}
	ExternalAuthController.RequestJSON(c, func() {
		if ExternalAuthController.Request.OauthProvider == "google" {
			GoogleLogin := services.GoogleAuthService{}
			ExternalAuthController.Service = &GoogleLogin.Service
			ExternalAuthController.Service.Constructor.OauthID = ExternalAuthController.Request.OauthID
			GoogleLogin.Authenticate(true)
		}
	})
}
