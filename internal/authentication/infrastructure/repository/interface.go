package repository

import (
	"context"

	"github.com/kevinyobeth/go-boilerplate/internal/authentication/domain/user"
)

type Repository interface {
	Register(c context.Context, request *user.UserDto) error
}
