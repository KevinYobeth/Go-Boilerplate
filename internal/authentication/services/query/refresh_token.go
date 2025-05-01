package query

import (
	"context"

	"github.com/kevinyobeth/go-boilerplate/config"
	"github.com/kevinyobeth/go-boilerplate/internal/authentication/domain/token"
	"github.com/kevinyobeth/go-boilerplate/internal/authentication/infrastructure/repository"
	"github.com/kevinyobeth/go-boilerplate/internal/authentication/services/helper"
	"github.com/kevinyobeth/go-boilerplate/shared/decorator"
	"github.com/kevinyobeth/go-boilerplate/shared/errors"

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
	if err != nil {
		return nil, errors.NewUnauthenticatedError(err)
	}

	user, err := h.repository.GetUser(c, uuid.MustParse(sub))
	if err != nil {
		return nil, errors.NewGenericError(err, "failed to get user by id")
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

func NewRefreshTokenHandler(repository repository.Repository, logger *zap.SugaredLogger, metricsClient metrics.Client) RefreshTokenHandler {
	if repository == nil {
		panic("repository is required")
	}

	return decorator.ApplyQueryDecorators(
		refreshTokenHandler{
			repository: repository,
		}, logger, metricsClient,
	)
}
