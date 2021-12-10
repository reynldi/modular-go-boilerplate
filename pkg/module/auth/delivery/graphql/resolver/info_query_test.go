package graphql

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/reynldi/modular-go-boilerplate/pkg/datastore/postgres/entity"
	"github.com/reynldi/modular-go-boilerplate/pkg/helper"
	globalmodel "github.com/reynldi/modular-go-boilerplate/pkg/model"
	"github.com/reynldi/modular-go-boilerplate/pkg/module/auth/delivery/graphql/directive"
	"github.com/reynldi/modular-go-boilerplate/pkg/module/auth/delivery/graphql/generated"
	mockService "github.com/reynldi/modular-go-boilerplate/pkg/module/auth/mock/service"
	"github.com/reynldi/modular-go-boilerplate/pkg/module/auth/model"
	"github.com/reynldi/modular-go-boilerplate/pkg/module/auth/service"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

var (
	userId          = "123456"
	email           = "email@mail.com"
	IsEmailVerified = true
	passwordHash    = "xxx"
	status          = 1
)

func TestQueryResolver_InfoQuery(t *testing.T) {
	t.Run("Error invalid email format", func(t *testing.T) {
		userAuthService := new(mockService.AuthUserService)
		resolvers := NewRootResolver(service.Service{
			AuthUserService: userAuthService,
		})

		conf := generated.Config{Resolvers: resolvers}
		conf.Directives.Binding = directive.Binding

		c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(conf)))
		q := `
			query {
			  infoQuery(input: { email: "email.com" }) {
				isEmailVerified
				email
				actionToken
			  }
			}
		`

		res, _ := c.RawPost(q)
		var graphErr []globalmodel.GraphErr

		errJson := json.Unmarshal(res.Errors, &graphErr)
		if errJson != nil {
			panic(errJson)
		}

<<<<<<< HEAD
		require.Equal(t, "email must be a valid email address", graphErr[0].Message)
		require.Equal(t, "email", graphErr[0].Path[2])
=======
		require.Equal(t, "email must be a valid email address", graphErr[0].Message )
		require.Equal(t, "email", graphErr[0].Path[2] )
>>>>>>> 5e80ffaf63d025da43d9475f8d9dd7544e70f26f
	})

	t.Run("Error user not found", func(t *testing.T) {
		userAuthService := new(mockService.AuthUserService)
		resolvers := NewRootResolver(service.Service{
			AuthUserService: userAuthService,
		})

		conf := generated.Config{Resolvers: resolvers}
		conf.Directives.Binding = directive.Binding
		c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(conf)))

		mockData := entity.AuthUser{}
		userAuthService.On("FindUserByEmail", mock.Anything).Return(&mockData, errors.New("failed"))
		q := `
			query {
			  infoQuery(input: { email: "different@email.com" }) {
				isEmailVerified
				email
				actionToken
			  }
			}
		`

		res, _ := c.RawPost(q)
		var graphErr []globalmodel.GraphErr

		errJson := json.Unmarshal(res.Errors, &graphErr)
		if errJson != nil {
			panic(errJson)
		}

		require.Equal(t, "User not found different@email.com", graphErr[0].Message)
	})

	t.Run("Should return info query data", func(t *testing.T) {
		userAuthService := new(mockService.AuthUserService)
		resolvers := NewRootResolver(service.Service{
			AuthUserService: userAuthService,
		})

		conf := generated.Config{Resolvers: resolvers}
		conf.Directives.Binding = directive.Binding
		c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(conf)))

		mockData := entity.AuthUser{
			UserId:          userId,
			Email:           email,
			IsEmailVerified: IsEmailVerified,
			PasswordHash:    passwordHash,
			Status:          status,
		}

		userAuthService.On("FindUserByEmail", mock.AnythingOfType("string")).Return(&mockData, nil)
		var res struct {
			InfoQuery model.InfoQuery
		}

		q := `
			query {
			  infoQuery(input: { email: "email@mail.com" }) {
				isEmailVerified
				email
				actionToken
			  }
			}
		`

		token := helper.GenSHA256(email)

		c.MustPost(q, &res)
		userAuthService.AssertExpectations(t)
		require.Equal(t, email, res.InfoQuery.Email)
		require.Equal(t, IsEmailVerified, res.InfoQuery.IsEmailVerified)
		require.Equal(t, token, res.InfoQuery.ActionToken)

	})
}
