package auth

import (
	"github.com/gin-gonic/gin"
	"godp.abdanhafidz.com/controller"
	"godp.abdanhafidz.com/models"
	"godp.abdanhafidz.com/services"
)

func CreateForgotPassword(c *gin.Context) {
	ForgotPassword := services.ForgotPasswordService{}
	ForgotPasswordController := controller.Controller[models.ForgotPasswordRequest, models.ForgotPassword, models.ForgotPassword]{
		Service: &ForgotPassword.Service,
	}
	ForgotPasswordController.RequestJSON(c, func() {
		ForgotPassword.Create(ForgotPasswordController.Request.Email)
	})

}
func ValidateForgotPassword(c *gin.Context) {
	ForgotPassword := services.ForgotPasswordService{}
	ForgotPasswordController := controller.Controller[models.ValidateForgotPasswordRequest, models.ForgotPassword, models.ForgotPassword]{
		Service: &ForgotPassword.Service,
	}
	ForgotPasswordController.RequestJSON(c, func() {
		ForgotPasswordController.Service.Constructor.Token = ForgotPasswordController.Request.Token
		ForgotPassword.Validate(&ForgotPasswordController.Request.NewPassword)
	})
}
