package query

import (
	"context"

	"github.com/kevinyobeth/go-boilerplate/internal/authentication/domain/user"
	"github.com/kevinyobeth/go-boilerplate/internal/authentication/infrastructure/repository"
	"github.com/kevinyobeth/go-boilerplate/shared/decorator"
	"github.com/kevinyobeth/go-boilerplate/shared/errors"
	"github.com/kevinyobeth/go-boilerplate/shared/metrics"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type GetUserRequest struct {
	ID uuid.UUID
}

type getUserHandler struct {
	repository repository.Repository
}

type GetUserHandler decorator.QueryHandler[GetUserRequest, *user.User]

func (h getUserHandler) Handle(c context.Context, params GetUserRequest) (*user.User, error) {
	user, err := h.repository.GetUser(c, params.ID)
	if err != nil {
		return nil, errors.NewGenericError(err, "failed to get user")
	}
	if user == nil {
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
