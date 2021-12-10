package graphql

import (
	"context"

	"github.com/reynldi/modular-go-boilerplate/pkg/helper"
	"github.com/reynldi/modular-go-boilerplate/pkg/module/auth/delivery/graphql/generated"
	"github.com/reynldi/modular-go-boilerplate/pkg/module/auth/model"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func (r *queryResolver) InfoQuery(ctx context.Context, input *model.InfoQueryInput) (*model.InfoQuery, error) {
	user, err := r.service.AuthUserService.FindUserByEmail(input.Email)

	if err != nil {
		return nil, gqlerror.Errorf("User not found " + input.Email)
	}

	token := helper.GenSHA256(input.Email)

	return &model.InfoQuery{
		Email:           user.Email,
		IsEmailVerified: user.IsEmailVerified,
		ActionToken:     token,
	}, nil
}

func (r *RootResolver) Query() generated.QueryResolver { return &queryResolver{r} }
