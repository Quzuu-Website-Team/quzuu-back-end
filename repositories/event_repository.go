package repositories

import (
	"github.com/google/uuid"
	"godp.abdanhafidz.com/models"
	"log"
)

func GetAllEventsPaginate(pagination PaginationConstructor) Repository[models.Events, []models.Events] {
	repo := Construct[models.Events, []models.Events](
		models.Events{},
	)
	repo.Pagination = pagination

	repo.Transactions(
		FindAllPaginate[models.Events, []models.Events],
	)

	if repo.RowsCount == 0 {
		log.Println("No events found with the provided pagination")
	}

	return *repo
}

func GetDetailEvent(idEvent uuid.UUID) Repository[models.Events, models.Events] {
	repo := Construct[models.Events, models.Events](
		models.Events{
			ID: idEvent,
		},
	)
	repo.Transactions(
		WhereGivenConstructor[models.Events, models.Events],
		Find[models.Events, models.Events],
	)
	return *repo
}

func GetEventAssigned(idEvent uuid.UUID, idAccount uuid.UUID) Repository[models.EventAssign, models.EventAssign] {
	repo := Construct[models.EventAssign, models.EventAssign](
		models.EventAssign{
			IDEvent:   idEvent,
			IDAccount: idAccount,
		},
	)
	repo.Transactions(
		WhereGivenConstructor[models.EventAssign, models.EventAssign],
		Find[models.EventAssign, models.EventAssign],
	)
	return *repo
}
