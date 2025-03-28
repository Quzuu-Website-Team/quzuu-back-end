package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"godp.abdanhafidz.com/controller"
	"godp.abdanhafidz.com/models"
	"godp.abdanhafidz.com/services"
)

func Verify(c *gin.Context) {
	emailVerification := services.EmailVerificationService{}
	emailVerificationController := controller.Controller[any, models.EmailVerification, models.EmailVerification]{
		Service: &emailVerification.Service,
	}
	query, _ := c.GetQuery("account_id")
	accountId, _ := uuid.Parse(query)
	emailVerificationController.Service.Constructor.AccountID = accountId
	emailVerification.Validate()
	emailVerificationController.Response(c)
}
