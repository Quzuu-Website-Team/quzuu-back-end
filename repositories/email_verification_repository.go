package repositories

import (
	"time"

	"github.com/google/uuid"
	"godp.abdanhafidz.com/models"
)

func CreateEmailVerification(uuid uuid.UUID, AccountId uuid.UUID, dueTime time.Time, token uint) Repository[models.EmailVerification, models.EmailVerification] {
	repo := Construct[models.EmailVerification, models.EmailVerification](
		models.EmailVerification{
			AccountId: AccountId,
			IsExpired: false,
			ExpiredAt: dueTime,
			Token:     token,
			Id:        uuid,
		},
	)
	Create(repo)
	return *repo
}

func GetEmailVerification(AccountId uuid.UUID, token uint) Repository[models.EmailVerification, models.EmailVerification] {
	repo := Construct[models.EmailVerification, models.EmailVerification](
		models.EmailVerification{
			AccountId: AccountId,
			IsExpired: false,
			Token:     token,
		},
	)
	repo.Transactions(
		WhereGivenConstructor[models.EmailVerification, models.EmailVerification],
		Find[models.EmailVerification, models.EmailVerification],
	)
	return *repo
}

func UpdateExpiredEmailVerification(uuid uuid.UUID) Repository[models.EmailVerification, models.EmailVerification] {
	repo := Construct[models.EmailVerification, models.EmailVerification](
		models.EmailVerification{Id: uuid},
	)

	repo.Transaction.Where("UUID = ?", uuid).First(&repo.Constructor)
	repo.Constructor.IsExpired = true
	repo.Transaction.Updates(repo.Constructor)
	repo.Result = repo.Constructor
	return *repo
}

func DeleteEmailVerification(token uint) Repository[models.EmailVerification, models.EmailVerification] {
	repo := Construct[models.EmailVerification, models.EmailVerification](
		models.EmailVerification{
			Token: token,
		},
	)

	repo.Transactions(
		WhereGivenConstructor[models.EmailVerification, models.EmailVerification],
		Delete[models.EmailVerification],
	)
	return *repo
}
