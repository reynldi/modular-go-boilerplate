package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/reynldi/modular-go-boilerplate/pkg/datastore/postgres/repository"
	"github.com/reynldi/modular-go-boilerplate/pkg/module/auth/delivery/graphql"
	"github.com/reynldi/modular-go-boilerplate/pkg/module/auth/delivery/rest"
	"github.com/reynldi/modular-go-boilerplate/pkg/module/auth/router"
	"github.com/reynldi/modular-go-boilerplate/pkg/module/auth/service"
)

var moduleName = "auth"

type Module struct {
	Router *router.Router
}

func Init(r *gin.Engine, repo *repository.Repository) {
	service := service.NewService(repo)
	handler := rest.NewHandler(service)
	router := router.NewRouter(handler)
	graphql := graphql.GraphServer(*service)

	r.POST(moduleName+"/graph/query", graphql)
	router.Register(r.Group(moduleName))
}
