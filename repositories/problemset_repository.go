package repositories

import (
	"github.com/google/uuid"
	"godp.abdanhafidz.com/models"
)

func GetProblemSet(idEvent uuid.UUID) Repository[models.ProblemSetAssign, []models.ProblemSet] {
	repo := Construct[models.ProblemSetAssign, []models.ProblemSet](
		models.ProblemSetAssign{
			IDEvent: idEvent,
		},
	)

	repo.Transactions(
		PreloadQuery[models.ProblemSetAssign, []models.ProblemSet]("Event", "ProblemSet"),
		FindAllPaginate[models.ProblemSetAssign, []models.ProblemSet],
	)
	return *repo
}
