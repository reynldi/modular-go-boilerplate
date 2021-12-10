package graphql

import (
	"context"

	"github.com/reynldi/modular-go-boilerplate/pkg/module/auth/delivery/graphql/generated"
	"github.com/reynldi/modular-go-boilerplate/pkg/module/auth/model"
)

func (r *queryResolver) UserAuth(ctx context.Context, input *model.UserAuthInput) (*model.UserAuth, error) {

	return &model.UserAuth{
		AccessToken:  "foo",
		RefreshToken: "bar",
	}, nil
}

func (r *RootResolver) UserAuthQuery() generated.QueryResolver { return &queryResolver{r} }
