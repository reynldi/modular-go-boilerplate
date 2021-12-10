package service

import (
	"github.com/reynldi/modular-go-boilerplate/pkg/datastore/postgres/repository"
)

type Service struct {
	AuthUserService IAuthUserService
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		AuthUserService: NewAuthService(repository),
	}
}
