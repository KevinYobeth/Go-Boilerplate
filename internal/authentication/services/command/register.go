package command

import (
	"context"

	"github.com/kevinyobeth/go-boilerplate/internal/authentication/domain/user"
	"github.com/kevinyobeth/go-boilerplate/internal/authentication/infrastructure/publisher"
	"github.com/kevinyobeth/go-boilerplate/internal/authentication/infrastructure/repository"
	eventcontract "github.com/kevinyobeth/go-boilerplate/internal/shared/event_contract"
	intraprocesscontract "github.com/kevinyobeth/go-boilerplate/internal/shared/intraprocess_contract"
	"github.com/kevinyobeth/go-boilerplate/pkg/common/decorator"
	"github.com/kevinyobeth/go-boilerplate/pkg/common/errors"
	"github.com/kevinyobeth/go-boilerplate/pkg/common/metrics"
	"github.com/kevinyobeth/go-boilerplate/pkg/common/validator"

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
	repository  repository.Repository
	userService intraprocesscontract.UserInterface
	publisher   publisher.Publisher
}

type RegisterHandler decorator.CommandHandler[*RegisterRequest]

func (h registerHandler) Handle(c context.Context, params *RegisterRequest) error {
	if err := validator.ValidateStruct(params); err != nil {
		return tracerr.Wrap(err)
	}

	userObj, err := h.userService.GetUserByEmail(c, params.Email)
	if err != nil {
		return tracerr.Wrap(err)
	}
	if userObj != nil {
		return errors.NewIncorrectInputError(nil, "user with this email already exists")
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

	err = h.publisher.UserRegistered(c, eventcontract.UserRegistered{
		UserID: dto.ID,
		Email:  dto.Email,
		Name:   dto.FirstName + " " + dto.LastName,
	})
	if err != nil {
		return errors.NewGenericError(err, "failed to publish user registered event")
	}

	return nil
}

func NewRegisterHandler(repository repository.Repository, userService intraprocesscontract.UserInterface, publisher publisher.Publisher, logger *zap.SugaredLogger, metricsClient metrics.Client) RegisterHandler {
	if repository == nil {
		panic("repository is required")
	}

	if userService == nil {
		panic("userService is required")
	}

	if publisher == nil {
		panic("publisher is required")
	}

	return decorator.ApplyCommandDecorators(
		registerHandler{
			repository:  repository,
			userService: userService,
			publisher:   publisher,
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
