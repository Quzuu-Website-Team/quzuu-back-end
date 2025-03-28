package user

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"godp.abdanhafidz.com/controller"
	"godp.abdanhafidz.com/models"
	"godp.abdanhafidz.com/services"
)

func ChangePassword(c *gin.Context) {
	authentication := services.AuthenticationService{}
	changePasswordController := controller.Controller[models.ChangePasswordRequest, models.Account, models.AuthenticatedUser]{
		Service: &authentication.Service,
	}
	changePasswordController.HeaderParse(c, func() {
		changePasswordController.Service.Constructor.ID, _ = uuid.Parse(changePasswordController.AccountData.UserID)
	})
	changePasswordController.RequestJSON(c, func() {
		authentication.Update(changePasswordController.Request.OldPassword, changePasswordController.Request.NewPassword)
	})
}
