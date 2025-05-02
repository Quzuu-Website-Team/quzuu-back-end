package exam

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"godp.abdanhafidz.com/controller"
	"godp.abdanhafidz.com/models"
	"godp.abdanhafidz.com/services"
)

func ExamQuestion(c *gin.Context) {
	examQuestion := services.ProblemSetService{}
	examQuestionController := controller.Controller[models.ExamRequest, models.ProblemSetAssign, []models.Questions]{
		Service: &examQuestion.Service,
	}

	examQuestionController.HeaderParse(c, func() {
		userid := examQuestionController.AccountData.UserID
		examQuestionController.RequestJSON(c, func() {
			examQuestionController.Service.Constructor.ProblemSetId, _ = uuid.Parse(examQuestionController.Request.IdProblemsets)
			examQuestionController.Service.Constructor.EventId, _ = uuid.Parse(examQuestionController.Request.IdEvent)
			examQuestion.Retrieve(userid)
		})
	})
}
