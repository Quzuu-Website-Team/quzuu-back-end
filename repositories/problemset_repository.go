package repositories

import (
	"github.com/google/uuid"
	"godp.abdanhafidz.com/models"
	"log"
)

func GetProblemSet(idEvent uuid.UUID) Repository[models.ProblemSetAssign, []models.ProblemSetAssign] {
	repo := Construct[models.ProblemSetAssign, []models.ProblemSetAssign](
		models.ProblemSetAssign{
			IDEvent: idEvent,
		},
	)
	repo.PreloadQuery = PreloadQueryConstructor{
		Tables: []string{"ProblemSet", "Event"},
	}

	repo.Transactions(
		WhereGivenConstructor[models.ProblemSetAssign, []models.ProblemSetAssign],
		Find[models.ProblemSetAssign, []models.ProblemSetAssign],
	)
	if repo.RowsCount == 0 {
		log.Println("No events found with the provided pagination")
	}
	return *repo
}
