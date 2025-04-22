package problemset

import (
	"github.com/gin-gonic/gin"
	"godp.abdanhafidz.com/controller"
	"godp.abdanhafidz.com/models"
	"godp.abdanhafidz.com/services"
)

func GetProblemsetsByEvent(c *gin.Context) {
	problemsets := services.GetProblemSetService{}
	problemsetsController := controller.Controller[models.GetProblemSetByEventID, models.ProblemSetAssign, []models.ProblemSet]{
		Service: &problemsets.Service,
	}
	problemsetsController.RequestJSON(c, func() {
		problemsets.GetProblemSetByEventID(problemsetsController.Request.IdEvent)
	})
}
