package model

type InfoQuery struct {
	Email           string `json:"email"`
	IsEmailVerified bool   `json:"isEmailVerified"`
	ActionToken     string `json:"actionToken"`
}

type InfoQueryInput struct {
	Email string `json:"email" validate:"required, email"`
}
