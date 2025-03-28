package services

import (
	"github.com/google/uuid"
	"godp.abdanhafidz.com/models"
	"godp.abdanhafidz.com/repositories"
)

type GetProblemSetService struct {
	Service[models.ProblemSetAssign, []models.ProblemSet]
}

func (s *GetProblemSetService) GetProblemSetByEventID(idEvent uuid.UUID) {
	problemsetsRepo := repositories.GetProblemSet(idEvent)
	if problemsetsRepo.NoRecord == true {
		s.Error = problemsetsRepo.RowsError
		s.Exception.DataNotFound = true
		s.Exception.Message = "no problemset found in the event"
		return
	}
	s.Result = problemsetsRepo.Result
}
