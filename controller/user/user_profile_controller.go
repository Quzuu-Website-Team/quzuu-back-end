package user

import (
	"github.com/gin-gonic/gin"
	"godp.abdanhafidz.com/controller"
	"godp.abdanhafidz.com/models"
	"godp.abdanhafidz.com/services"
)

func Profile(c *gin.Context) {
	userProfile := services.UserProfileService{}
	userProfileController := controller.Controller[any, models.AccountDetails, models.UserProfileResponse]{
		Service: &userProfile.Service,
	}
	userProfileController.HeaderParse(c, func() {
		userProfileController.Service.Constructor.AccountId = userProfileController.AccountData.UserID
		userProfile.Retrieve()
		userProfileController.Response(c)
	},
	)
}
