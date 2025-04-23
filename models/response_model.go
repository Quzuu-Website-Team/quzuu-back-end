package models

type SuccessResponse struct {
	Status   string `json:"status"`
	Message  string `json:"message"`
	Data     any    `json:"data"`
	MetaData any    `json:"meta_data"`
}

type ErrorResponse struct {
	Status   string    `json:"status"`
	Message  string    `json:"message"`
	Errors   Exception `json:"errors"`
	MetaData any       `json:"meta_data"`
}
type AuthenticatedUser struct {
	Account Account `json:"account"`
	Token   string  `json:"token"`
}

type EventDetailResponse struct {
	Data           *Events
	RegisterStatus int `json:"register_status"`
}

type Options struct {
	OptionCategory OptionCategory `json:"option_category"`
	OptionValues   []OptionValues `json:"option_values"`
}
type OptionsResponse struct {
	Options []Options `json:"options"`
}

type UserProfileResponse struct {
	Account Account        `json:"account"`
	Details AccountDetails `json:"details"`
}

type AcademyMaterialResponse struct {
	Materials AcademyMaterial
	Contents  []AcademyContent
}
type AcademyResponse struct {
	Academy   Academy                   `json:"academy"`
	Materials []AcademyMaterialResponse `json:"academy_materials"`
}

type AllAcademyResponse struct {
	Academies []AcademyResponse `json:"academy_dasar"`
}
