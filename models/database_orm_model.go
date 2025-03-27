package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm/dialects/postgres"
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
	AccountID uuid.UUID `json:"account_id"`
	IsExpired bool      `json:"is_expired"`
	CreatedAt time.Time `json:"created_at"`
	ExpiredAt time.Time `json:"expired_at"`

	Account *Account `gorm:"foreignKey:AccountID"`
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

type Announcement struct {
	IDAnnouncement uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id_announcement"`
	Title          string    `json:"title"`
	CreatedAt      time.Time `json:"created_at"`
	Message        string    `json:"message"`
	Publisher      string    `json:"publisher"`
	IDEvent        uint      `json:"id_event"`
}

type ProblemSet struct {
	IDProblemSet uuid.UUID     `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id_problem_set"`
	Title        string        `json:"title"`
	Duration     time.Duration `json:"duration"`
	Randomize    uint          `json:"randomize"`
	MC_Count     uint          `json:"mc_count"`
	SA_Count     uint          `json:"sa_count"`
	Essay_Count  uint          `json:"essay_count"`
}

type Questions struct {
	IDQuestion   uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id_question"`
	Type         string    `json:"type"` //MultChoices, ShortAns, Essay, IntPuzzle, IntType
	Question     string    `json:"question"`
	Options      []string  `gorm:"type:text[]" json:"options"`
	AnsKey       []string  `gorm:"type:text[]" json:"ans_key"`
	CorrMark     float64   `json:"corr_mark"`
	IncorrMark   float64   `json:"incorr_mark"`
	NullMark     float64   `json:"null_mark"`
	IDProblemSet uuid.UUID `json:"id_problem_set"`

	ProblemSet *ProblemSet `gorm:"foreignKey:IDProblemSet"`
}

type EventAssign struct {
	IDAssign   uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id_assign"`
	IDAccount  uuid.UUID `json:"id_account"`
	IDEvent    uuid.UUID `json:"id_event"`
	AssignedAt time.Time `json:"assigned_at"`

	Account *Account `gorm:"foreignKey:IDAccount"`
	Event   *Events  `gorm:"foreignKey:IDEvent"`
}

type ProblemSetAssign struct {
	IDProblemSetAssign uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id_problem_set_assign"`
	IDEvent            uuid.UUID `json:"id_event"`
	IDProblemSet       uuid.UUID `json:"id_problem_set"`

	Event      *Events     `gorm:"foreignKey:IDEvent"`
	ProblemSet *ProblemSet `gorm:"foreignKey:IDProblemSet"`
}

type Result struct {
	IDResult      uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id_result"`
	IDAccount     uuid.UUID `json:"id_account"`
	IDEvent       uuid.UUID `json:"id_event"`
	IDProblemSet  uuid.UUID `json:"id_problem_set"`
	IDProgress    uuid.UUID `json:"id_progress"`
	FinishTime    time.Time `json:"finish_time"`
	Correct       uint      `json:"correct"`
	Incorrect     uint      `json:"incorrect"`
	Empty         uint      `json:"empty"`
	OnCorrection  uint      `json:"on_correction"`
	ManualScoring float64   `json:"manual_scoring"`
	MCScore       float64   `json:"mc_score"`
	ManualScore   float64   `json:"manual_score"`
	FinalScore    float64   `json:"final_score"`

	Account      *Account      `gorm:"foreignKey:IDAccount"`
	Event        *Events       `gorm:"foreignKey:IDEvent"`
	ProblemSet   *ProblemSet   `gorm:"foreignKey:IDProblemSet"`
	ExamProgress *ExamProgress `gorm:"foreignKey:IDProgress"`
}

type ExamProgress struct {
	IDProgress     uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id_progress"`
	IDAccount      uuid.UUID `json:"id_account"`
	IDEvent        uuid.UUID `json:"id_event"`
	IDProblemSet   uuid.UUID `json:"id_problem_set"`
	CreatedAt      time.Time `json:"created_at"`
	DueAt          time.Time `json:"due_at"`
	QuestionsOrder []string  `gorm:"type:text[]" json:"questions_order"`
	Answers        any       `gorm:"type:jsonb" json:"answers"`

	Account    *Account    `gorm:"foreignKey:IDAccount"`
	Event      *Events     `gorm:"foreignKey:IDEvent"`
	ProblemSet *ProblemSet `gorm:"foreignKey:IDProblemSet"`
}
type ExamProgress_Result struct {
	IDProgress     uuid.UUID      `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id_progress"`
	IDAccount      uuid.UUID      `json:"id_account"`
	IDEvent        uuid.UUID      `json:"id_event"`
	IDProblemSet   uuid.UUID      `json:"id_problem_set"`
	CreatedAt      time.Time      `json:"created_at"`
	DueAt          time.Time      `json:"due_at"`
	QuestionsOrder []string       `gorm:"type:text[]" json:"questions_order"`
	Answers        postgres.Jsonb `gorm:"type:jsonb" json:"answers"`

	Account    *Account    `gorm:"foreignKey:IDAccount"`
	Event      *Events     `gorm:"foreignKey:IDEvent"`
	ProblemSet *ProblemSet `gorm:"foreignKey:IDProblemSet"`
}

// Gorm table name settings
func (Account) TableName() string           { return "account" }
func (AccountDetails) TableName() string    { return "account_details" }
func (EmailVerification) TableName() string { return "email_verifications" }
func (ExternalAuth) TableName() string      { return "extern_auth" }
func (ForgotPassword) TableName() string    { return "forgot_password" }
func (Events) TableName() string            { return "events" }
