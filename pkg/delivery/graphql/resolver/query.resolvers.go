package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/reynldi/modular-go-boilerplate/pkg/application/config"
	"github.com/reynldi/modular-go-boilerplate/pkg/delivery/graphql/generated"
	"github.com/reynldi/modular-go-boilerplate/pkg/delivery/graphql/model"
	"github.com/reynldi/modular-go-boilerplate/pkg/helper"
)

func (r *queryResolver) HealthCheck(ctx context.Context) (*model.HealthCheck, error) {
	return &model.HealthCheck{
		ServiceName: config.GetConfig().Application.AppName,
		Uptime: helper.GetUptime(),
	}, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
