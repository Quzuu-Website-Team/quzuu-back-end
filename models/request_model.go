package models

import "github.com/google/uuid"

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email" binding:"required,email"`
	Phone    string `json:"phone"`
	Password string `json:"password" binding:"required"`
}

type CreateEmailVerificationRequest struct {
	AccountID int `json:"account_id" binding:"required"`
}

type ChangePasswordRequest struct {
	OldPassword string `json:"old_password" binding:"required" `
	NewPassword string `json:"new_password" binding:"required" `
}

type GetProblemSetByEventID struct {
	IdEvent uuid.UUID `json:"id_event" binding:"required"`
}
type EventDetailRequest struct {
	IdUser  uuid.UUID `json:"id_user"`
	IdEvent uuid.UUID `json:"id_event"`
}

type JoinEventRequest struct {
	IdEvent   uuid.UUID `json:"id_event"`
	EventCode string    `json:"event_code"`
}
