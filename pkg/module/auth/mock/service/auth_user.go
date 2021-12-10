package mockService

import (
	"github.com/reynldi/modular-go-boilerplate/pkg/datastore/postgres/entity"
	"github.com/stretchr/testify/mock"
)

type AuthUserService struct {
	mock.Mock
}

func (s *AuthUserService) FindUserByEmail(email string) (*entity.AuthUser, error) {
	args := s.Called(email)
	return args.Get(0).(*entity.AuthUser), args.Error(1)
}
