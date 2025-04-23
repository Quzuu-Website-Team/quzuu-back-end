package models

import "github.com/google/uuid"

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email" binding:"required,email"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type CreateEmailVerificationRequest struct {
	AccountId uuid.UUID `json:"account_id" binding:"required"`
}

type ChangePasswordRequest struct {
	OldPassword string `json:"old_password" binding:"required" `
	NewPassword string `json:"new_password" binding:"required" `
}

type EventDetailRequest struct {
	IdUser  uuid.UUID `json:"id_user"`
	EventId uuid.UUID `json:"id_event"`
}

type JoinEventRequest struct {
	EventId   uuid.UUID `json:"id_event"`
	EventCode string    `json:"event_code"`
}

type CreateVerifyEmailRequest struct {
	Token uint `json:"token" binding:"required"`
}

type OptionsRequest struct {
	OptionName  string   `json:"option_name" binding:"required"`
	OptionValue []string `json:"option_values" binding:"required"`
}

type ExternalAuthRequest struct {
	OauthID       string `json:"oauth_id" binding:"required"`
	OauthProvider string `json:"oauth_provider" binding:"required"`
}

type ForgotPasswordRequest struct {
	Email string `json:"email" binding:"required,email"`
}
type ValidateForgotPasswordRequest struct {
	Token       uint   `json:"token" binding:"required"`
	NewPassword string `json:"new_password"`
}
