package services

import (
	"github.com/google/uuid"
	"godp.abdanhafidz.com/models"
	"godp.abdanhafidz.com/repositories"
)

type EventList struct {
	Service[models.Events, []models.Events]
}

func (s *EventList) EventListWithFilter(userid uuid.UUID, pagination repositories.PaginationConstructor) {
	eventsRepo := repositories.EventListWithFilter(userid, pagination)

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
