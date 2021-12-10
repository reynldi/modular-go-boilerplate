package rest

import (
	"github.com/reynldi/modular-go-boilerplate/pkg/module/auth/service"
)

type Handler struct {
	AuthUserHandler AuthUserHandler
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		AuthUserHandler: NewAuthUserHandler(service),
	}
}
