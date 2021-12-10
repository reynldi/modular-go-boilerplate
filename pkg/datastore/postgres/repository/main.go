package repository

import "gorm.io/gorm"

type Repository struct {
	AuthRepository *authRepository
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		AuthRepository: NewAuthRepository(db),
	}
}
