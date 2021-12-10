package repository

import (
	"github.com/reynldi/modular-go-boilerplate/pkg/datastore/postgres/entity"
	"gorm.io/gorm"
)

type IAuthRepository interface {
	FindUserByEmail(email string) (entity.AuthUser, error)
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *authRepository {
	return &authRepository{db}
}

func (r *authRepository) FindUserByEmail(email string) (entity.AuthUser, error) {
	var user entity.AuthUser
	err := r.db.Where("email = ?", email).First(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}
