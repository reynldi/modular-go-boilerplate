package model

type UserAuth struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type UserAuthInput struct {
	Email    string `json:"email" validate:"required, email"`
	Password string `json:"password" validate:"required, alphanum"`
}
