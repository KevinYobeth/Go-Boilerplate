package command

import (
	"context"
	"go-boilerplate/shared/decorator"
	"go-boilerplate/shared/errors"
	"go-boilerplate/src/authentication/domain/user"
	"go-boilerplate/src/authentication/infrastructure/repository"

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

	dto := user.NewRegisterDto(
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

func NewRegisterHandler(repositor repository.Repository, logger *zap.SugaredLogger) RegisterHandler {
	if repositor == nil {
		panic("repository is required")
	}

	return decorator.ApplyCommandDecorators(
		registerHandler{
			repository: repositor,
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
