package intraprocess

import (
	"context"

	"github.com/google/uuid"
	"github.com/kevinyobeth/go-boilerplate/internal/shared/interfaces"
	"github.com/kevinyobeth/go-boilerplate/internal/user/services"
	"github.com/kevinyobeth/go-boilerplate/internal/user/services/query"
	"github.com/ztrue/tracerr"
)

type UserIntraprocess struct {
	service services.Application
}

func NewUserIntraprocessService(service services.Application) interfaces.UserIntraprocess {
	return &UserIntraprocess{service: service}
}

func (i *UserIntraprocess) GetUser(ctx context.Context, id uuid.UUID) (*interfaces.User, error) {
	user, err := i.service.Queries.GetUser.Handle(ctx, &query.GetUserRequest{
		ID: id,
	})
	if err != nil {
		return nil, tracerr.Wrap(err)
	}

	return TransformToIntraprocessUser(user), nil
}

func (i *UserIntraprocess) GetUserByEmail(ctx context.Context, email string) (*interfaces.User, error) {
	user, err := i.service.Queries.GetUser.Handle(ctx, &query.GetUserRequest{
		Email: &email,
	})
	if err != nil {
		return nil, tracerr.Wrap(err)
	}

	return TransformToIntraprocessUser(user), nil
}
