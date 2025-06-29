package repositories

import (
	"log"

	"github.com/google/uuid"
	"godp.abdanhafidz.com/models"
)

func GetAllEventsPaginate(pagination PaginationConstructor) Repository[models.Events, []models.Events] {
	repo := Construct[models.Events, []models.Events](
		models.Events{},
	)
	repo.Pagination = pagination

	// Add debug log to verify pagination values
	log.Printf("Pagination - Limit: %d, Offset: %d", pagination.Limit, pagination.Offset)

	// Transactions that execute the query
	repo.Transactions(
		FindAllPaginate[models.Events, []models.Events],
	)

	// Check if there's an error or no records
	if repo.RowsCount == 0 {
		log.Println("No events found with the provided pagination")
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
