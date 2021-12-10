package service

import (
	"github.com/reynldi/modular-go-boilerplate/pkg/datastore/postgres/entity"
	"github.com/reynldi/modular-go-boilerplate/pkg/datastore/postgres/repository"
)

type IAuthUserService interface {
	FindUserByEmail(email string) (*entity.AuthUser, error)
}

type authService struct {
	repository *repository.Repository
}

func NewAuthService(repository *repository.Repository) *authService {
	return &authService{repository: repository}
}

func (s *authService) FindUserByEmail(email string) (*entity.AuthUser, error) {
	var user entity.AuthUser
	user, err := s.repository.AuthRepository.FindUserByEmail(email)

	if err != nil {
		return nil, err
	}

	return &user, nil

}
