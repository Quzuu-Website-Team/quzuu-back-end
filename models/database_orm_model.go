package models

import (
	"time"

	"github.com/google/uuid"
)

type Account struct {
	Id                uuid.UUID  `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Email             string     `gorm:"uniqueIndex" json:"email"`
	Password          string     `json:"password"`
	IsEmailVerified   bool       `json:"is_email_verified"`
	IsDetailCompleted bool       `json:"is_detail_completed"`
	CreatedAt         time.Time  `json:"created_at"`
	DeletedAt         *time.Time `json:"deleted_at" gorm:"default:null"`
}

type AccountDetails struct {
	ID            uuid.UUID  `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	AccountID     uuid.UUID  `json:"account_id"`
	InitialName   string     `json:"initial_name"`
	FullName      *string    `json:"full_name"`
	DateOfBirth   *time.Time `json:"date_of_birth"`
	PlaceOfBirth  *string    `json:"place_of_birth"`
	Domicile      *string    `json:"domicile"`
	LastJob       *string    `json:"last_job"`
	Gender        *bool      `json:"gender"`
	LastEducation *string    `json:"last_education"`
	MaritalStatus *bool      `json:"marital_status"`
	Avatar        *string    `json:"avatar"`
	PhoneNumber   *string    `json:"phone_number"`

	Account *Account `gorm:"foreignKey:AccountID"`
}

type EmailVerification struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Token     uint      `json:"token"`
	AccountID uint      `json:"account_id"`
	IsExpired bool      `json:"is_expired"`
	CreatedAt time.Time `json:"created_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

type ExternalAuth struct {
	ID            uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	OauthID       string    `json:"oauth_id"`
	AccountID     uint      `json:"account_id"`
	OauthProvider string    `json:"oauth_provider"`
}

type FCM struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	AccountID uint      `json:"account_id"`
	FCMToken  string    `json:"fcm_token"`
}

type ForgotPassword struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Token     uint      `json:"token"`
	AccountID uint      `json:"account_id"`
	IsExpired bool      `json:"is_expired"`
	CreatedAt time.Time `json:"created_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

type Events struct {
	IDEvent    uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id_event"`
	Title      string    `json:"title"`
	StartEvent time.Time `json:"start_event"`
	EndEvent   time.Time `json:"end_event"`
	SID        string    `json:"sid"`
	Public     string    `json:"public"`
}

type ProblemSetAssign struct {
	IDProblemSetAssign uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id_problem_set_assign"`
	IDEvent            uuid.UUID `json:"id_event"`
	IDProblemSet       uuid.UUID `json:"id_problem_set"`

	Event *Events `gorm:"foreignKey:IDEvent"`
}

// Gorm table name settings
func (Account) TableName() string           { return "account" }
func (AccountDetails) TableName() string    { return "account_details" }
func (EmailVerification) TableName() string { return "email_verifications" }
func (ExternalAuth) TableName() string      { return "extern_auth" }
func (ForgotPassword) TableName() string    { return "forgot_password" }
func (Events) TableName() string            { return "events" }
