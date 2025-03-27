package user

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"godp.abdanhafidz.com/controller"
	"godp.abdanhafidz.com/models"
	"godp.abdanhafidz.com/services"
)

func UpdateProfile(c *gin.Context) {
	userProfile := services.UserProfileService{}
	userUpdateProfileController := controller.Controller[models.AccountDetails, models.AccountDetails, models.AccountDetails]{
		Service: &userProfile.Service,
	}

	userUpdateProfileController.RequestJSON(c, func() {
		userUpdateProfileController.Service.Constructor = userUpdateProfileController.Request
		userUpdateProfileController.HeaderParse(c, func() {
			userUpdateProfileController.Service.Constructor.AccountID, _ = uuid.Parse(userUpdateProfileController.AccountData.UserID)
		})
		userProfile.Update()
	},
	)
}
