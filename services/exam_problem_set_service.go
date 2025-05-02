package services

import (
	"fmt"
	"github.com/google/uuid"
	"godp.abdanhafidz.com/models"
	"godp.abdanhafidz.com/repositories"
)

type ProblemSetService struct {
	Service[models.ProblemSetAssign, []models.Questions]
}

func (s *ProblemSetService) Retrieve(userid uuid.UUID) {
	examProgress := repositories.GetExamProgress(userid)

	var questions []models.Questions
	if examProgress.NoRecord {
		// create question order
		problemSetID := s.Constructor.ProblemSetId
		// get all the questions
		q := repositories.GetAllQuestionWithProblemSetID(problemSetID)

		for _, question := range q.Result {
			fmt.Println(question)
		}
		questions = q.Result
	}
	s.Result = questions
}
