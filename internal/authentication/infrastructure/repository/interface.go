package repository

import (
	"context"

	"github.com/kevinyobeth/go-boilerplate/internal/authentication/domain/user"
	interfaces "github.com/kevinyobeth/go-boilerplate/internal/shared/interfaces/event"

	"github.com/google/uuid"
)

type Repository interface {
	Register(c context.Context, request *user.UserDto) error
	GetUser(c context.Context, id uuid.UUID) (*user.User, error)
	GetUserByEmail(c context.Context, email string) (*user.User, error)
}

type Publisher interface {
	UserRegistered(c context.Context, payload interfaces.UserRegistered) error
}
