package query

import (
	"context"

	"github.com/kevinyobeth/go-boilerplate/internal/authentication/domain/token"
	"github.com/kevinyobeth/go-boilerplate/internal/authentication/infrastructure/repository"
	"github.com/kevinyobeth/go-boilerplate/internal/authentication/services/helper"
	"github.com/kevinyobeth/go-boilerplate/shared/decorator"
	"github.com/kevinyobeth/go-boilerplate/shared/errors"

	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type loginHandler struct {
	repository repository.Repository
	logger     *zap.SugaredLogger
}

type LoginHandler decorator.QueryHandler[LoginRequest, *token.Token]

func (h loginHandler) Handle(c context.Context, params LoginRequest) (*token.Token, error) {

	user, err := h.repository.GetUserByEmail(c, params.Email)
	if err != nil {
		return nil, errors.NewGenericError(err, "failed to get user by email")
	}

	if user == nil {
		return nil, errors.NewIncorrectInputError(nil, "wrong email or password")
	}

	if err := comparePassword(user.Password, params.Password); err != nil {
		return nil, errors.NewIncorrectInputError(nil, "wrong email or password")
	}

	jwtToken, err := helper.GenerateToken(c, helper.GenerateTokenOpts{
		Params: helper.GenerateTokenRequest{
			User: *user,
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

func NewLoginHandler(repository repository.Repository, logger *zap.SugaredLogger, metricsClient metrics.Client) LoginHandler {
	if repository == nil {
		panic("repository is required")
	}

	return decorator.ApplyQueryDecorators(
		loginHandler{
			repository: repository,
			logger:     logger,
		}, logger, metricsClient,
	)
}

func comparePassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
