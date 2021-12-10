package graphql

import (
	"github.com/reynldi/modular-go-boilerplate/pkg/module/auth/service"
)

type RootResolver struct {
	service service.Service
}

func NewRootResolver(service service.Service) *RootResolver {
	return &RootResolver{
		service: service,
	}
}

type queryResolver struct{ *RootResolver }
