package command

import (
	"context"

	"github.com/kevinyobeth/go-boilerplate/internal/authentication/domain/user"
	"github.com/kevinyobeth/go-boilerplate/internal/authentication/infrastructure/repository"
	"github.com/kevinyobeth/go-boilerplate/shared/decorator"
	"github.com/kevinyobeth/go-boilerplate/shared/errors"

	"github.com/ztrue/tracerr"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type RegisterRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type registerHandler struct {
	repository repository.Repository
}

type RegisterHandler decorator.CommandHandler[RegisterRequest]

func (h registerHandler) Handle(c context.Context, params RegisterRequest) error {
	hashedPassword, err := hashPassword(params.Password)
	if err != nil {
		return tracerr.Wrap(err)
	}

	dto := user.NewUserDto(
		params.FirstName,
		params.LastName,
		params.Email,
		hashedPassword,
	)
	err = h.repository.Register(c, dto)
	if err != nil {
		return errors.NewGenericError(err, "failed to register user")
	}

	return nil
}

func NewRegisterHandler(repository repository.Repository, logger *zap.SugaredLogger) RegisterHandler {
	if repository == nil {
		panic("repository is required")
	}

	return decorator.ApplyCommandDecorators(
		registerHandler{
			repository: repository,
		}, logger,
	)
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return "", errors.NewGenericError(err, "failed to hash password")
	}
	return string(bytes), nil
}
