package repositories

import "godp.abdanhafidz.com/models"

func CreateForgotPassword(forgotPassword models.ForgotPassword) Repository[models.ForgotPassword, models.ForgotPassword] {
	repo := Construct[models.ForgotPassword, models.ForgotPassword](
		forgotPassword,
	)
	Create(repo)
	return *repo
}

func GetForgotPasswordByToken(token uint) Repository[models.ForgotPassword, models.ForgotPassword] {
	repo := Construct[models.ForgotPassword, models.ForgotPassword](
		models.ForgotPassword{Token: token},
	)
	repo.Transactions(
		WhereGivenConstructor[models.ForgotPassword, models.ForgotPassword],
		Find[models.ForgotPassword, models.ForgotPassword],
	)
	return *repo
}
