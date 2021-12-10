package module

import (
	"github.com/reynldi/modular-go-boilerplate/pkg/datastore/postgres/repository"
)

type ModuleConfig struct {
	Repository repository.Repository
}
