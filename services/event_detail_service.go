package services

import (
	"godp.abdanhafidz.com/models"
	"godp.abdanhafidz.com/repositories"
)

type EventDetailService struct {
	Service[models.EventDetailRequest, models.EventDetailResponse]
}

func (s *EventDetailService) Retrieve() {
	// ngecek event nya dulu
	detail := repositories.GetDetailEvent(s.Constructor.IdEvent)
	if detail.NoRecord {
		s.Error = detail.RowsError
		s.Exception.DataNotFound = true
		s.Exception.Message = "Event detail not found"
		return
	}
	// ngecek apakah si event dan si user udah ke assign/register
	assigned := repositories.GetEventAssigned(s.Constructor.IdEvent, s.Constructor.IdUser)

	// kalo eventnya private dan di assigned itu ga ditemuin data antara event dan account ke register
	// bakal ke tolak karena unauthorized
	if detail.Result.Public == "N" && assigned.NoRecord == true {
		s.Error = assigned.RowsError
		s.Exception.Unauthorized = true
		s.Exception.Message = "your account doesnt have access to this event"
		return
	}

	// ini ngecek kalo ke assign/register ke event public atau ngga
	var registerStatus int
	if assigned.NoRecord == true {
		registerStatus = 0
	} else {
		registerStatus = 1
	}

	s.Error = detail.RowsError
	s.Result.Data = &detail.Result
	s.Result.RegisterStatus = registerStatus
}
