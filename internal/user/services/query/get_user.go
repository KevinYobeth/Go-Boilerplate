package query

import (
	"context"

	"github.com/kevinyobeth/go-boilerplate/internal/user/domain/user"
	"github.com/kevinyobeth/go-boilerplate/internal/user/infrastructure/repository"
	"github.com/kevinyobeth/go-boilerplate/pkg/common/decorator"
	"github.com/kevinyobeth/go-boilerplate/pkg/common/errors"
	"github.com/kevinyobeth/go-boilerplate/pkg/common/metrics"
	"github.com/ztrue/tracerr"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type GetUserRequest struct {
	ID             uuid.UUID
	Email          *string
	SilentNotFound bool
}

type getUserHandler struct {
	repository repository.Repository
}

type GetUserHandler decorator.QueryHandler[*GetUserRequest, *user.User]

func (h getUserHandler) Handle(c context.Context, params *GetUserRequest) (*user.User, error) {
	if err := validateRequest(params); err != nil {
		return nil, tracerr.Wrap(err)
	}

	var user *user.User
	var err error

	if params.ID != uuid.Nil {
		user, err = h.repository.GetUser(c, params.ID)
	}

	if params.Email != nil {
		user, err = h.repository.GetUserByEmail(c, *params.Email)
	}

	if err != nil {
		return nil, errors.NewGenericError(err, "failed to get user")
	}
	if !params.SilentNotFound && user == nil {
		return nil, errors.NewGenericError(nil, "user not found")
	}

	return user, nil
}

func NewGetUserHandler(repository repository.Repository, logger *zap.SugaredLogger, metricsClient metrics.Client) GetUserHandler {
	return decorator.ApplyQueryDecorators(
		getUserHandler{
			repository: repository,
		}, logger, metricsClient,
	)
}

func validateRequest(request *GetUserRequest) error {
	if request.ID == uuid.Nil && request.Email == nil {
		return errors.NewGenericError(nil, "either ID or Email must be provided")
	}

	if request.ID != uuid.Nil && request.Email != nil {
		return errors.NewGenericError(nil, "only one of ID or Email should be provided")
	}

	return nil
}
