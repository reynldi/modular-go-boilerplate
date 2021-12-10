package graphql

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/reynldi/modular-go-boilerplate/pkg/application/config"
	"github.com/reynldi/modular-go-boilerplate/pkg/delivery/graphql/generated"
	graphql "github.com/reynldi/modular-go-boilerplate/pkg/delivery/graphql/resolver"
)


func GqlServer() gin.HandlerFunc {
	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graphql.Resolver{}}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func GqlPlayground() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/graph/query")

	return func(c *gin.Context) {
		if config.GetConfig().Application.Environment == "development" {
			h.ServeHTTP(c.Writer, c.Request)
		}

		c.Done()
	}
}
