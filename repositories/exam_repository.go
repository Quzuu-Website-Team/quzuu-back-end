package repositories

import (
	"github.com/google/uuid"
	"godp.abdanhafidz.com/models"
)

func GetExamProgress(user_id uuid.UUID) Repository[models.ExamProgress, models.ExamProgress] {
	repo := Construct[models.ExamProgress, models.ExamProgress](
		models.ExamProgress{
			AccountId: user_id,
		})
	repo.Transactions(
		WhereGivenConstructor[models.ExamProgress, models.ExamProgress],
		Find[models.ExamProgress, models.ExamProgress],
	)

	return *repo
}

func GetExamResult(user_id uuid.UUID) Repository[models.Result, models.Result] {
	repo := Construct[models.Result, models.Result](
		models.Result{
			AccountId: user_id,
		})

	repo.Transactions(
		WhereGivenConstructor[models.Result, models.Result],
		Find[models.Result, models.Result],
	)

	return *repo
}

func GetAllQuestionWithProblemSetID(problemset_id uuid.UUID) Repository[models.Questions, []models.Questions] {
	repo := Construct[models.Questions, []models.Questions](
		models.Questions{
			ProblemSetId: problemset_id,
		})

	repo.Transactions(
		WhereGivenConstructor[models.Questions, []models.Questions],
		Find[models.Questions, []models.Questions],
	)
	return *repo
}
