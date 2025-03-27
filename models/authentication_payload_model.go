package models

import "github.com/google/uuid"

type AccountData struct {
	UserID       uuid.UUID
	VerifyStatus string
	ErrVerif     error
}
