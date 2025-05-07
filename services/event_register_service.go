package services

import (
	"github.com/google/uuid"
	"godp.abdanhafidz.com/models"
	"godp.abdanhafidz.com/repositories"
)

type JoinEventService struct {
	Service[models.JoinEventRequest, models.EventDetailResponse]
}

func (s *JoinEventService) Create(AccountId uuid.UUID) {
	event := repositories.GetEventByCode(s.Constructor.EventCode)
	// log.Printf("event: %v", event)
	if event.NoRecord {
		s.Error = event.RowsError
		s.Exception.DataNotFound = true
		s.Exception.Message = "event not found"
		return
	}
	// ngecek apakah si event dan si user udah ke assign/register
	assigned := repositories.GetEventAssigned(s.Constructor.EventId, AccountId)
	if assigned.NoRecord == true {
		accountAssigned := &models.EventAssign{
			Id:        uuid.New(),
			EventId:   s.Constructor.EventId,
			AccountId: AccountId,
		}
		repositories.AssignEvent(*accountAssigned)
	} else {
		s.Exception.DataDuplicate = true
		s.Exception.Message = "account already assigned to this event"
		return
	}
	s.Result.Data = &event.Result
	s.Result.RegisterStatus = 1
}
