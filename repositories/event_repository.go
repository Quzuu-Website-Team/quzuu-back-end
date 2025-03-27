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

func GetDetailEvent(idEvent uuid.UUID) Repository[models.Events, models.Events] {
	repo := Construct[models.Events, models.Events](
		models.Events{
			IDEvent: idEvent,
		},
	)
	repo.Transactions(
		WhereGivenConstructor[models.Events, models.Events],
		Find[models.Events, models.Events],
	)
	return *repo
}
