package query

import (
	"context"
	"go-boilerplate/shared/decorator"
	"go-boilerplate/shared/errors"
	"go-boilerplate/src/authentication/domain/token"
	"go-boilerplate/src/authentication/infrastructure/repository"
	"go-boilerplate/src/authentication/services/helper"

	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type loginHandler struct {
	repository repository.Repository
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

	return &token.Token{
		Token:     jwtToken.Token,
		ExpiredAt: jwtToken.ExpiredAt,
		RefreshToken: token.RefreshToken{
			Token:     jwtToken.RefreshToken.Token,
			ExpiredAt: jwtToken.RefreshToken.ExpiredAt,
		},
	}, nil
}

func NewLoginHandler(repository repository.Repository, logger *zap.SugaredLogger) LoginHandler {
	if repository == nil {
		panic("repository is required")
	}

	return decorator.ApplyQueryDecorators(
		loginHandler{
			repository: repository,
		}, logger,
	)
}

func comparePassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
