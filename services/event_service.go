package services

import (
	"godp.abdanhafidz.com/models"
	"godp.abdanhafidz.com/repositories"
)

type EventService struct {
	Service[models.Events, []models.Events]
}

func (s *EventService) GetAllEventPaginate(pagination repositories.PaginationConstructor) {
	eventsRepo := repositories.GetAllEventsPaginate(pagination)

	events := eventsRepo.Result

	totalRecords := eventsRepo.RowsCount
	totalPages := (totalRecords / pagination.Limit) + 1

	metadata := repositories.PaginationMetadata{
		TotalRecords: totalRecords,
		TotalPages:   totalPages,
		CurrentPage:  (pagination.Offset / pagination.Limit) + 1,
		PageSize:     pagination.Limit,
	}

	s.Result = events
	s.MetaData = metadata
}
