package repositories

import (
	"log"
	"time"

	"github.com/google/uuid"
	"godp.abdanhafidz.com/models"
)

func EventListWithFilter(userId uuid.UUID, pagination PaginationConstructor) Repository[models.Events, []models.Events] {
	repo := Construct[models.Events, []models.Events](
		models.Events{},
	)
	repo.Pagination = pagination

	log.Printf("Pagination - Limit: %d, Offset: %d", pagination.Limit, pagination.Offset)

	repo.Transactions(
		FindAllPaginate[models.Events, []models.Events],
	)

	currentTime := time.Now()

	filteredEvents := []models.Events{}
	for _, event := range repo.Result {
		if event.IsPublic || isAssignedToEvent(userId, event.Id) {
			if event.EndEvent.After(currentTime) {
				filteredEvents = append(filteredEvents, event)
			}
		}
	}

	repo.Result = filteredEvents
	repo.RowsCount = len(filteredEvents)

	if repo.RowsCount == 0 {
		log.Println("No accessible events found for the user")
	}

	return *repo
}

func GetEventDetailByEventId(EventId uuid.UUID) Repository[models.Events, models.Events] {
	repo := Construct[models.Events, models.Events](
		models.Events{
			Id: EventId,
		},
	)
	repo.Transactions(
		WhereGivenConstructor[models.Events, models.Events],
		Find[models.Events, models.Events],
	)
	return *repo
}

func GetEventDetailBySlug(slug string) Repository[models.Events, models.Events] {
	repo := Construct[models.Events, models.Events](
		models.Events{
			Slug: slug,
		},
	)
	repo.Transactions(
		WhereGivenConstructor[models.Events, models.Events],
		Find[models.Events, models.Events],
	)
	return *repo
}

func GetEventAssigned(EventId uuid.UUID, AccountId uuid.UUID) Repository[models.EventAssign, models.EventAssign] {
	repo := Construct[models.EventAssign, models.EventAssign](
		models.EventAssign{
			EventId:   EventId,
			AccountId: AccountId,
		},
	)
	repo.Transactions(
		WhereGivenConstructor[models.EventAssign, models.EventAssign],
		Find[models.EventAssign, models.EventAssign],
	)
	return *repo
}

func GetEventByCode(code string) Repository[models.Events, models.Events] {
	repo := Construct[models.Events, models.Events](
		models.Events{EventCode: code},
	)

	repo.Transactions(
		WhereGivenConstructor[models.Events, models.Events],
		Find[models.Events, models.Events],
	)
	if repo.RowsCount == 0 {
		log.Println("No events found with the provided code")
	}
	return *repo
}

func AssignEvent(eventAssign models.EventAssign) Repository[models.EventAssign, models.EventAssign] {
	repo := Construct[models.EventAssign, models.EventAssign](
		eventAssign,
	)
	Create(repo)
	return *repo
}

func isAssignedToEvent(userId uuid.UUID, eventId uuid.UUID) bool {
	repo := GetEventAssigned(eventId, userId)

	return repo.RowsCount > 0
}
