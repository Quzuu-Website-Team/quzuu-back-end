package repositories

import (
	"github.com/google/uuid"
	"godp.abdanhafidz.com/models"
)

func CreateExternalAuth(oauth models.ExternalAuth) Repository[models.ExternalAuth, models.ExternalAuth] {
	repo := Construct[models.ExternalAuth, models.ExternalAuth](
		oauth,
	)
	Create(repo)
	return *repo
}

func GetExternalAuthByAccountId(AccountId uuid.UUID) Repository[models.ExternalAuth, []models.ExternalAuth] {
	repo := Construct[models.ExternalAuth, []models.ExternalAuth](
		models.ExternalAuth{
			AccountId: AccountId,
		},
	)
	repo.Transactions(
		WhereGivenConstructor[models.ExternalAuth, []models.ExternalAuth],
		Find[models.ExternalAuth, []models.ExternalAuth],
	)
	return *repo
}

func GetExternalAccountByOauthId(oauthId string) Repository[models.ExternalAuth, models.ExternalAuth] {
	repo := Construct[models.ExternalAuth, models.ExternalAuth](
		models.ExternalAuth{
			OauthID: oauthId,
		},
	)
	repo.Transactions(
		WhereGivenConstructor[models.ExternalAuth, models.ExternalAuth],
		Find[models.ExternalAuth, models.ExternalAuth],
	)
	return *repo
}
