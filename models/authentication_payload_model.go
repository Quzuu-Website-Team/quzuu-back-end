package models

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type AccountData struct {
	UserID       uuid.UUID
	VerifyStatus string
	Role         string
	ErrVerif     error
}
type CustomClaims struct {
	jwt.RegisteredClaims
	UserID uuid.UUID `json:"id"`
	Role   string    `json:"role"`
}
