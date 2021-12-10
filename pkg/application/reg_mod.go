package application

import (
	"github.com/gin-gonic/gin"
	"github.com/reynldi/modular-go-boilerplate/pkg/datastore/postgres/repository"
	"github.com/reynldi/modular-go-boilerplate/pkg/module/auth"
)

type regModuleOptions struct {
	GinEngine  *gin.Engine
	Repository *repository.Repository
}

func RegisterModule(options regModuleOptions) {
	auth.Init(options.GinEngine, options.Repository)
}