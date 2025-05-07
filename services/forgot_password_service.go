package services

import (
	"fmt"
	"log"
	"math/rand/v2"
	"net/smtp"
	"time"

	"godp.abdanhafidz.com/config"
	"godp.abdanhafidz.com/models"
	"godp.abdanhafidz.com/repositories"
)

type ForgotPasswordService struct {
	Service[models.ForgotPassword, models.ForgotPassword]
}

func (s *ForgotPasswordService) Create(email string) {
	if email == "" {
		s.Exception.BadRequest = true
		s.Exception.Message = "Email is required!"
		return
	}
	accountRepo := repositories.GetAccountbyEmail(email)
	if accountRepo.NoRecord {
		s.Error = accountRepo.RowsError
		s.Exception.DataNotFound = true
		s.Exception.Message = "There is no account data with given credentials!"
		return
	}

	remainingTime := time.Duration(config.EMAIL_VERIFICATION_DURATION) * time.Hour
	dueTime := CalculateDueTime(remainingTime)

	token := uint(rand.IntN(999999-100000) + 100000)
	s.Constructor.ExpiredAt = dueTime
	s.Constructor.AccountId = accountRepo.Result.Id
	s.Constructor.Token = token
	repo := repositories.CreateForgotPassword(s.Constructor)

	s.Error = repo.RowsError
	s.Result = repo.Result
	// ⬇ Kirim token ke email user menggunakan SMTP
	go func(toEmail string, token uint) {
		from := config.SMTP_SENDER_EMAIL
		password := config.SMTP_SENDER_PASSWORD
		smtpHost := config.SMTP_HOST
		smtpPort := config.SMTP_PORT

		auth := smtp.PlainAuth("", from, password, smtpHost)

		subject := "Forgot Password Token"
		body := fmt.Sprintf("Your Forgot Password token is: %06d\nPlease use it before it expires.", token)

		msg := []byte("To: " + toEmail + "\r\n" +
			"Subject: " + subject + "\r\n" +
			"\r\n" +
			body + "\r\n")

		err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{toEmail}, msg)
		if err != nil {
			s.Error = err
			log.Printf("Error sending verification email: %v", err)
			return
		}
	}(accountRepo.Result.Email, token)
	s.Result.Token = 0
}

func (s *ForgotPasswordService) Validate(newPassword *string) {

	fgPasswordRepo := repositories.GetForgotPasswordByToken(s.Constructor.Token)
	s.Error = fgPasswordRepo.RowsError
	if fgPasswordRepo.NoRecord {
		s.Exception.DataNotFound = true
		s.Exception.Message = "There is no forgot password data with given credentials!"
		return
	}
	if fgPasswordRepo.Result.ExpiredAt.Before(time.Now()) {
		s.Exception.Unauthorized = true
		s.Exception.Message = "Token has expired!"
		return
	}

	accountRepo := repositories.GetAccountById(fgPasswordRepo.Result.AccountId)
	if accountRepo.NoRecord {
		s.Error = accountRepo.RowsError
		s.Exception.DataNotFound = true
		s.Exception.Message = "There is no account data with given credentials!"
		return
	}
	s.Result = fgPasswordRepo.Result
	if newPassword == nil {
		return
	}
	// fmt.Println("Previous Account", accountRepo.Result)
	// fmt.Println("New password", *newPassword)
	hashed_password, _ := HashPassword(*newPassword)
	accountRepo.Result.Password = hashed_password
	changePassword := repositories.UpdateAccount(accountRepo.Result)
	// fmt.Println("New Account", changePassword.Result)
	if changePassword.RowsError != nil {
		s.Error = changePassword.RowsError
		s.Exception.QueryError = true
		s.Exception.Message = "Failed to update password!"
		return
	}
	fgPasswordRepo.Result.Token = 0

}
