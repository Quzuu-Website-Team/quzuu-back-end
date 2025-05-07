package services

import (
	"github.com/google/uuid"
	"godp.abdanhafidz.com/models"
	"godp.abdanhafidz.com/repositories"
)

type EventDetailService struct {
	Service[models.Events, models.EventDetailResponse]
}

func (s *EventDetailService) Retrieve(userId uuid.UUID) {
	// ngecek event nya dulu
	detail := repositories.GetEventDetailBySlug(s.Constructor.Slug)
	if detail.NoRecord {
		s.Error = detail.RowsError
		s.Exception.DataNotFound = true
		s.Exception.Message = "Event detail not found"
		return
	}
	// ngecek apakah si event dan si user udah ke assign/register
	assigned := repositories.GetEventAssigned(detail.Result.Id, userId)

	// kalo eventnya private dan di assigned itu ga ditemuin data antara event dan account ke register
	// bakal ke tolak karena unauthorized
	// el event tidak public mamka tidak boleh masuk
	if !detail.Result.IsPublic && assigned.NoRecord {
		s.Error = assigned.RowsError
		s.Exception.Unauthorized = true
		s.Exception.Message = "your account doesnt have access to this event"
		return
	}

	// ini ngecek kalo ke assign/register ke event public atau ngga
	// sudah boolean, jadi ga usah dibandingin
	var registerStatus int
	if assigned.NoRecord {
		registerStatus = 0
	} else {
		registerStatus = 1
	}

	s.Error = detail.RowsError
	s.Result.Data = &detail.Result
	s.Result.RegisterStatus = registerStatus
}
