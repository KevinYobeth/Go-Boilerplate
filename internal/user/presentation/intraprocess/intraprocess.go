package intraprocess

import (
	"context"

	"github.com/google/uuid"
	"github.com/kevinyobeth/go-boilerplate/internal/shared/contract"
	"github.com/kevinyobeth/go-boilerplate/internal/user/services"
	"github.com/kevinyobeth/go-boilerplate/internal/user/services/query"
	"github.com/ztrue/tracerr"
)

type UserIntraprocess struct {
	service services.Application
}

func NewUserIntraprocessService(service services.Application) contract.UserIntraprocess {
	return &UserIntraprocess{service: service}
}

func (i *UserIntraprocess) GetUser(ctx context.Context, id uuid.UUID) (*contract.User, error) {
	user, err := i.service.Queries.GetUser.Handle(ctx, &query.GetUserRequest{
		ID: id,
	})
	if err != nil {
		return nil, tracerr.Wrap(err)
	}

	return TransformToIntraprocessUser(user), nil
}

func (i *UserIntraprocess) GetUserByEmail(ctx context.Context, email string) (*contract.User, error) {
	user, err := i.service.Queries.GetUser.Handle(ctx, &query.GetUserRequest{
		Email:          &email,
		SilentNotFound: true,
	})
	if err != nil {
		return nil, tracerr.Wrap(err)
	}

	return TransformToIntraprocessUser(user), nil
}
