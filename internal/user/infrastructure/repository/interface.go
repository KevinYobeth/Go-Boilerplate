package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/kevinyobeth/go-boilerplate/internal/user/domain/user"
)

type Repository interface {
	GetUser(c context.Context, id uuid.UUID) (*user.User, error)
	GetUserByEmail(c context.Context, email string) (*user.User, error)
}
