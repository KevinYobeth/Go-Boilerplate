package repository

import (
	"context"
	"go-boilerplate/src/authentication/domain/user"

	"github.com/google/uuid"
)

type Repository interface {
	Register(c context.Context, request *user.UserDto) error
	GetUser(c context.Context, id uuid.UUID) (*user.User, error)
	GetUserByEmail(c context.Context, email string) (*user.User, error)
}
