package input

type GetUserByEmailInput struct {
	Email string `json:"email" binding:"required"`
}

type UserAuth struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
