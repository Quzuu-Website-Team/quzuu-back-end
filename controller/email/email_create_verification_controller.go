package controller

import (
	"github.com/gin-gonic/gin"
	"godp.abdanhafidz.com/controller"
	"godp.abdanhafidz.com/models"
	"godp.abdanhafidz.com/services"
)

func CreateVerification(c *gin.Context) {
	emailVerification := services.EmailVerificationService{}
	emailVerificationController := controller.Controller[models.CreateEmailVerificationRequest, models.EmailVerification, models.EmailVerification]{
		Service: &emailVerification.Service,
	}
	emailVerificationController.RequestJSON(c, func() {
		emailVerification.Create(emailVerificationController.Request.Email)
		emailVerificationController.Response(c)
	})

}
