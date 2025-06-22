package repository

import (
	"context"

	"github.com/kevinyobeth/go-boilerplate/internal/authentication/domain/user"
	eventcontract "github.com/kevinyobeth/go-boilerplate/internal/shared/event_contract"
)

type Repository interface {
	Register(c context.Context, request *user.UserDto) error
}

type Publisher interface {
	UserRegistered(c context.Context, payload eventcontract.UserRegistered) error
}
