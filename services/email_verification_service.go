package services

import (
	"math/rand/v2"
	"time"

	"godp.abdanhafidz.com/config"
	"godp.abdanhafidz.com/models"
	"godp.abdanhafidz.com/repositories"
)

type EmailVerificationService struct {
	Service[models.EmailVerification, models.EmailVerification]
}

func (s *EmailVerificationService) Create(email string) {
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

	repo := repositories.CreateEmailVerification(s.Constructor.Id, accountRepo.Result.Id, dueTime, token)

	s.Error = repo.RowsError
	s.Result = repo.Result

	// // ⬇ Kirim token ke email user menggunakan SMTP
	// go func(toEmail string, token uint) {
	// 	from := config.SMTP_SENDER_EMAIL
	// 	password := config.SMTP_SENDER_PASSWORD
	// 	smtpHost := config.SMTP_HOST
	// 	smtpPort := config.SMTP_PORT

	// 	auth := smtp.PlainAuth("", from, password, smtpHost)

	// 	subject := "Email Verification Token"
	// 	body := fmt.Sprintf("Your verification token is: %06d\nPlease use it before it expires.", token)

	// 	msg := []byte("To: " + toEmail + "\r\n" +
	// 		"Subject: " + subject + "\r\n" +
	// 		"\r\n" +
	// 		body + "\r\n")

	// 	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{toEmail}, msg)
	// 	if err != nil {
	// 		s.Error = err
	// 		log.Printf("Error sending verification email: %v", err)
	// 		return
	// 	}
	// }(accountRepo.Result.Email, token)
	// s.Result.Token = repo.Result.Token
}

func (s *EmailVerificationService) Validate(email string) {
	accountRepo := repositories.GetAccountbyEmail(email)
	if accountRepo.NoRecord {
		s.Error = accountRepo.RowsError
		s.Exception.DataNotFound = true
		s.Exception.Message = "There is no account data with given credentials!"
		return
	}

	repo := repositories.GetEmailVerification(accountRepo.Result.Id, s.Constructor.Token)
	s.Error = repo.RowsError

	if repo.NoRecord {
		s.Exception.DataNotFound = true
		s.Exception.Message = "Invalid token!"
		return
	}

	if repo.Result.ExpiredAt.Before(time.Now()) {
		s.Exception.Unauthorized = true
		s.Exception.Message = "Token has expired!"
		repositories.UpdateExpiredEmailVerification(s.Constructor.Id)
		s.Delete()
		return
	}
	account := repositories.GetAccountById(repo.Result.AccountId)
	account.Result.IsEmailVerified = true

	repositories.UpdateAccount(account.Result)
	s.Result = repo.Result
}

func (s *EmailVerificationService) Delete() {
	repo := repositories.DeleteEmailVerification(s.Constructor.Token)
	s.Error = repo.RowsError
	if repo.NoRecord {
		s.Exception.DataNotFound = true
		s.Exception.Message = "Invalid token!"
		return
	}
	s.Result = repo.Result
}
