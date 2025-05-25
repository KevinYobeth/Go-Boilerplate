package command

import (
	"context"

	"github.com/kevinyobeth/go-boilerplate/internal/authentication/domain/user"
	"github.com/kevinyobeth/go-boilerplate/internal/authentication/infrastructure/repository"
	"github.com/kevinyobeth/go-boilerplate/shared/decorator"
	"github.com/kevinyobeth/go-boilerplate/shared/errors"
	"github.com/kevinyobeth/go-boilerplate/shared/metrics"
	"github.com/kevinyobeth/go-boilerplate/shared/notification"
	"github.com/kevinyobeth/go-boilerplate/shared/validator"

	"github.com/ztrue/tracerr"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type RegisterRequest struct {
	FirstName string `conform:"trim" validate:"required,min=3,max=255"`
	LastName  string `conform:"trim" validate:"required,min=3,max=255"`
	Email     string `conform:"trim" validate:"required,email,min=3,max=255"`
	Password  string `validate:"required,min=8,max=255"`
}

type registerHandler struct {
	repository repository.Repository
}

type RegisterHandler decorator.CommandHandler[*RegisterRequest]

func (h registerHandler) Handle(c context.Context, params *RegisterRequest) error {
	if err := validator.ValidateStruct(params); err != nil {
		return tracerr.Wrap(err)
	}

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

	emailStrategy, err := notification.NewEmailNotificationStrategy()
	if err != nil {
		return tracerr.Wrap(err)
	}
	notification := notification.NewNotification(emailStrategy)

	err = notification.Send("me@kevinyobeth.com", []string{params.Email}, "Welcome to Go Boilerplate",
		"Hello "+params.FirstName+",\n\nThank you for registering with Go Boilerplate.\n\nBest regards,\nGo Boilerplate Team")
	if err != nil {
		return tracerr.Wrap(err)
	}

	return nil
}

func NewRegisterHandler(repository repository.Repository, logger *zap.SugaredLogger, metricsClient metrics.Client) RegisterHandler {
	if repository == nil {
		panic("repository is required")
	}

	return decorator.ApplyCommandDecorators(
		registerHandler{
			repository: repository,
		}, logger, metricsClient,
	)
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return "", errors.NewGenericError(err, "failed to hash password")
	}
	return string(bytes), nil
}
