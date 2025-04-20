package repository

import (
	"context"
	"go-boilerplate/src/authentication/domain/user"
)

type Repository interface {
	Register(c context.Context, request *user.RegisterDto) error
}
