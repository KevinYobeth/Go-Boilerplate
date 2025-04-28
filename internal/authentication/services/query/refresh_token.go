package query

import (
	"context"
	"go-boilerplate/config"
	"go-boilerplate/internal/authentication/domain/token"
	"go-boilerplate/internal/authentication/infrastructure/repository"
	"go-boilerplate/internal/authentication/services/helper"
	"go-boilerplate/shared/decorator"
	"go-boilerplate/shared/errors"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token"`
}

type refreshTokenHandler struct {
	repository repository.Repository
}

type RefreshTokenHandler decorator.QueryHandler[RefreshTokenRequest, *token.Token]

func (h refreshTokenHandler) Handle(c context.Context, params RefreshTokenRequest) (*token.Token, error) {
	jwtConfig := config.LoadJWTConfig()
	jwtRefreshToken, err := jwt.Parse(params.RefreshToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.NewIncorrectInputError(nil, "invalid signing method")
		}
		return []byte(jwtConfig.JWTRefreshSecret), nil
	})
	if err != nil {
		return nil, errors.NewUnauthenticatedError(err)
	}
	if !jwtRefreshToken.Valid {
		return nil, errors.NewUnauthenticatedError(nil)
	}

	sub, err := jwtRefreshToken.Claims.GetSubject()
	user, err := h.repository.GetUser(c, uuid.MustParse(sub))

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

func NewRefreshTokenHandler(repository repository.Repository, logger *zap.SugaredLogger) RefreshTokenHandler {
	if repository == nil {
		panic("repository is required")
	}

	return decorator.ApplyQueryDecorators(
		refreshTokenHandler{
			repository: repository,
		}, logger,
	)
}
