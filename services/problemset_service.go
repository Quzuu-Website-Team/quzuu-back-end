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
	problemsetAssignRepo := repositories.GetProblemSet(idEvent)
	if problemsetAssignRepo.NoRecord == true {
		s.Error = problemsetAssignRepo.RowsError
		s.Exception.DataNotFound = true
		s.Exception.Message = "no problemset found in the event"
		return
	}

	var problemSets []models.ProblemSet
	for _, assign := range problemsetAssignRepo.Result {
		if assign.ProblemSet != nil {
			problemSets = append(problemSets, *assign.ProblemSet)
		}
	}

	s.Result = problemSets
}
