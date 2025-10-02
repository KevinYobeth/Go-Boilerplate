package query

import (
	"context"

	"github.com/kevinyobeth/go-boilerplate/internal/authentication/domain/token"
	"github.com/kevinyobeth/go-boilerplate/internal/authentication/domain/user"
	"github.com/kevinyobeth/go-boilerplate/internal/authentication/infrastructure/repository"
	"github.com/kevinyobeth/go-boilerplate/internal/authentication/services/helper"
	intraprocesscontract "github.com/kevinyobeth/go-boilerplate/internal/shared/intraprocess_contract"
	"github.com/kevinyobeth/go-boilerplate/shared/decorator"
	"github.com/kevinyobeth/go-boilerplate/shared/errors"
	"github.com/kevinyobeth/go-boilerplate/shared/metrics"
	"github.com/kevinyobeth/go-boilerplate/shared/validator"

	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type LoginRequest struct {
	Email    string `conform:"trim" validate:"required,email,min=3,max=255"`
	Password string `validate:"required,min=8,max=255"`
}

type loginHandler struct {
	repository  repository.Repository
	userService intraprocesscontract.UserInterface
	logger      *zap.SugaredLogger
}

type LoginHandler decorator.QueryHandler[*LoginRequest, *token.Token]

func (h loginHandler) Handle(c context.Context, params *LoginRequest) (*token.Token, error) {
	if err := validator.ValidateStruct(params); err != nil {
		return nil, errors.NewIncorrectInputError(err, err.Error())
	}

	userObj, err := h.userService.GetUserByEmail(c, params.Email)
	if err != nil {
		return nil, errors.NewGenericError(err, "failed to get user by email")
	}
	if userObj == nil {
		return nil, errors.NewIncorrectInputError(nil, "wrong email or password")
	}

	if err := comparePassword(userObj.Password, params.Password); err != nil {
		return nil, errors.NewIncorrectInputError(nil, "wrong email or password")
	}

	jwtToken, err := helper.GenerateToken(c, helper.GenerateTokenOpts{
		Params: helper.GenerateTokenRequest{
			User: user.User{
				ID:        userObj.ID,
				FirstName: userObj.FirstName,
				LastName:  userObj.LastName,
				Email:     userObj.Email,
				Password:  userObj.Password,
			},
		},
	})
	if err != nil {
		return nil, errors.NewGenericError(err, "failed to generate token")
	}

	return &token.Token{
		Token:     jwtToken.Token,
		ExpiredAt: jwtToken.ExpiredAt,
		RefreshToken: token.RefreshToken{
			Token:     jwtToken.RefreshToken.Token,
			ExpiredAt: jwtToken.RefreshToken.ExpiredAt,
		},
	}, nil
}

func NewLoginHandler(repository repository.Repository, userService intraprocesscontract.UserInterface, logger *zap.SugaredLogger, metricsClient metrics.Client) LoginHandler {
	if repository == nil {
		panic("repository is required")
	}

	return decorator.ApplyQueryDecorators(
		loginHandler{
			repository:  repository,
			userService: userService,
			logger:      logger,
		}, logger, metricsClient,
	)
}

func comparePassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
