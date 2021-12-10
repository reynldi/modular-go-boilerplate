package graphql

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/gin-gonic/gin"
	"github.com/reynldi/modular-go-boilerplate/pkg/module/auth/delivery/graphql/directive"
	"github.com/reynldi/modular-go-boilerplate/pkg/module/auth/delivery/graphql/generated"
	graphql "github.com/reynldi/modular-go-boilerplate/pkg/module/auth/delivery/graphql/resolver"
	"github.com/reynldi/modular-go-boilerplate/pkg/module/auth/service"
)

func GraphServer(service service.Service) gin.HandlerFunc {
	resolver := graphql.NewRootResolver(service)

	c := generated.Config{Resolvers: resolver}
	c.Directives.Binding = directive.Binding

	h := handler.NewDefaultServer(generated.NewExecutableSchema(c))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
