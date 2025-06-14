package repository

import (
	"context"

	"github.com/kevinyobeth/go-boilerplate/internal/authentication/domain/user"
	interfaces "github.com/kevinyobeth/go-boilerplate/internal/shared/interfaces/event"
)

type Repository interface {
	Register(c context.Context, request *user.UserDto) error
}

type Publisher interface {
	UserRegistered(c context.Context, payload interfaces.UserRegistered) error
}
