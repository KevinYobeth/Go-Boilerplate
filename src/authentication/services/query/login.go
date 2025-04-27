package query

import (
	"context"
	"go-boilerplate/config"
	"go-boilerplate/shared/decorator"
	"go-boilerplate/shared/errors"
	"go-boilerplate/src/authentication/domain/token"
	"go-boilerplate/src/authentication/infrastructure/repository"
	"time"

	"github.com/golang-jwt/jwt/v5"
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

	appConfig := config.LoadAppConfig()
	jwtConfig := config.LoadJWTConfig()

	issuedAt := time.Now().UTC()
	expiredAt := issuedAt.Add(jwtConfig.JWTShortLife)
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
		"sub":   user.ID,
		"iss":   appConfig.AppName,
		"exp":   jwt.NewNumericDate(expiredAt),
		"iat":   jwt.NewNumericDate(issuedAt),
	})
	tokenString, err := jwtToken.SignedString([]byte(jwtConfig.JWTSecret))
	if err != nil {
		return nil, errors.NewGenericError(err, "failed to sign JWT token")
	}

	return &token.Token{
		Token:     tokenString,
		ExpiredAt: expiredAt,
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
