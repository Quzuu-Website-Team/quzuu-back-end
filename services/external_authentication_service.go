package services

import (
	"context"
	"errors"

	"godp.abdanhafidz.com/models"
	"godp.abdanhafidz.com/repositories"
	"google.golang.org/api/idtoken"
)

type GoogleAuthService struct {
	Service[models.ExternalAuth, models.AuthenticatedUser]
}

func (s *GoogleAuthService) Authenticate(isAgree bool) {
	GoogleAuth := repositories.GetExternalAccountByOauthId(s.Constructor.OauthID)
	payload, errGoogleAuth := idtoken.Validate(context.Background(), s.Constructor.OauthID, "")
	s.Error = errGoogleAuth
	if errGoogleAuth != nil {
		s.Exception.Unauthorized = true
		s.Exception.Message = "Oauth Provider Failed Login (Google Authentication)"
		return
	}
	email := payload.Claims["email"]
	checkRegisteredEmail := repositories.GetAccountbyEmail(email.(string))
	if !checkRegisteredEmail.NoRecord {
		token, _ := GenerateToken(&checkRegisteredEmail.Result)
		checkRegisteredEmail.Result.Password = "SECRET"
		s.Result = models.AuthenticatedUser{
			Account: checkRegisteredEmail.Result,
			Token:   token,
		}
		return
	}
	if GoogleAuth.NoRecord {
		if !isAgree {
			s.Exception.BadRequest = true
			s.Exception.Message = "Please agree to the terms and conditions to create an account"
			return
		}
		s.Constructor.OauthProvider = "Google"

		createAccount := repositories.CreateAccount(models.Account{
			Email:           email.(string),
			IsEmailVerified: true,
		})

		s.Constructor.AccountId = createAccount.Result.Id
		createGoogleAuth := repositories.CreateExternalAuth(s.Constructor)

		GoogleAuth.Result.AccountId = createGoogleAuth.Result.AccountId
		userProfile := UserProfileService{}
		userProfile.Constructor.AccountId = GoogleAuth.Result.AccountId
		userProfile.Create()
		if userProfile.Error != nil {
			s.Error = userProfile.Error
			return
		}
		s.Error = createGoogleAuth.RowsError
		s.Error = errors.Join(s.Error, createAccount.RowsError)
	}

	accountData := repositories.GetAccountById(GoogleAuth.Result.AccountId)
	token, err_tok := GenerateToken(&accountData.Result)

	if err_tok != nil {
		s.Error = errors.Join(s.Error, err_tok)
	}

	accountData.Result.Password = "SECRET"
	s.Result = models.AuthenticatedUser{
		Account: accountData.Result,
		Token:   token,
	}
	s.Error = accountData.RowsError

}
