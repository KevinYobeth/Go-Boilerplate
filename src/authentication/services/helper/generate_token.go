package helper

import (
	"context"
	"go-boilerplate/config"
	"go-boilerplate/shared/errors"
	"go-boilerplate/src/authentication/domain/token"
	"go-boilerplate/src/authentication/domain/user"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type GenerateTokenRequest struct {
	User user.User
}

type GenerateTokenOpts struct {
	Params GenerateTokenRequest
}

func GenerateToken(c context.Context, opts GenerateTokenOpts) (*token.Token, error) {
	appConfig := config.LoadAppConfig()
	jwtConfig := config.LoadJWTConfig()

	issuedAt := time.Now().UTC()
	expiredAt := issuedAt.Add(jwtConfig.JWTShortLife)
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": opts.Params.User.Email,
		"sub":   opts.Params.User.ID,
		"iss":   appConfig.AppName,
		"exp":   jwt.NewNumericDate(expiredAt),
		"iat":   jwt.NewNumericDate(issuedAt),
	})
	tokenString, err := jwtToken.SignedString([]byte(jwtConfig.JWTSecret))
	if err != nil {
		return nil, errors.NewGenericError(err, "failed to sign JWT token")
	}

	refreshExpiredAt := issuedAt.Add(jwtConfig.JWTLongLife)
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": opts.Params.User.ID,
		"exp": jwt.NewNumericDate(refreshExpiredAt),
		"iat": jwt.NewNumericDate(issuedAt),
	})
	refreshTokenString, err := refreshToken.SignedString([]byte(jwtConfig.JWTRefreshSecret))
	if err != nil {
		return nil, errors.NewGenericError(err, "failed to sign JWT refresh token")
	}

	return &token.Token{
		Token:     tokenString,
		ExpiredAt: expiredAt,
		RefreshToken: token.RefreshToken{
			Token:     refreshTokenString,
			ExpiredAt: refreshExpiredAt,
		},
	}, nil
}
