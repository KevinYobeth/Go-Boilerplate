package query

import (
	"context"
	"go-boilerplate/shared/decorator"
	"go-boilerplate/shared/errors"
	"go-boilerplate/src/authentication/domain/user"
	"go-boilerplate/src/authentication/infrastructure/repository"

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

func NewGetUserHandler(repository repository.Repository, logger *zap.SugaredLogger) GetUserHandler {
	return decorator.ApplyQueryDecorators(
		getUserHandler{
			repository: repository,
		}, logger,
	)
}
