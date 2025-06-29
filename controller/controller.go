package controller

import (
	"github.com/gin-gonic/gin"
	"godp.abdanhafidz.com/models"
	"godp.abdanhafidz.com/repositories"
	"godp.abdanhafidz.com/services"
	"godp.abdanhafidz.com/utils"
)

type (
	Controllers interface {
		RequestJSON(c *gin.Context)
		Response(c *gin.Context)
	}
	Controller[TRequest any, TConstructor any, TResult any] struct {
		AccountData models.AccountData
		Request     TRequest
		Service     *services.Service[TConstructor, TResult]
	}
)

func (controller *Controller[T1, T2, T3]) HeaderParse(c *gin.Context, act func()) {
	cParam, _ := c.Get("accountData")
	if cParam != nil {
		controller.AccountData = cParam.(models.AccountData)
	}
	act()
}
func (controller *Controller[T1, T2, T3]) RequestJSON(c *gin.Context, act func()) {
	cParam, _ := c.Get("accountData")
	if cParam != nil {
		controller.AccountData = cParam.(models.AccountData)
	}
	errBinding := c.ShouldBindJSON(&controller.Request)
	if errBinding != nil {
		utils.ResponseFAIL(c, 400, models.Exception{
			BadRequest: true,
			Message:    "Invalid Request!, recheck your request, there's must be some problem about required parameter or type parameter",
		})
		return
	} else {
		act()
		controller.Response(c)
	}
}
func (controller *Controller[T1, T2, T3]) Response(c *gin.Context) {
	switch {
	case controller.Service.Error != nil:
		utils.LogError(controller.Service.Error)
		utils.ResponseFAIL(c, 500, models.Exception{
			InternalServerError: true,
			Message:             "Internal Server Error",
		})

	case controller.Service.Exception.DataDuplicate:
		utils.ResponseFAIL(c, 400, controller.Service.Exception)
	case controller.Service.Exception.Unauthorized:
		utils.ResponseFAIL(c, 401, controller.Service.Exception)
	case controller.Service.Exception.DataNotFound:
		utils.ResponseFAIL(c, 404, controller.Service.Exception)
	case controller.Service.Exception.Message != "":
		utils.ResponseFAIL(c, 400, controller.Service.Exception)
	default:
		if controller.Service.MetaData != (repositories.PaginationMetadata{}) {
			utils.ResponseOK(c, controller.Service.Result, controller.Service.MetaData)
		} else {
			utils.ResponseOK(c, controller.Service.Result, struct{}{})
		}
	}
}
