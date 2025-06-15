package intraprocess

import (
	"context"

	"github.com/google/uuid"
	"github.com/kevinyobeth/go-boilerplate/internal/authentication/domain/user"
)

type UserIntraprocess interface {
	GetUser(ctx context.Context, id uuid.UUID) (*user.User, error)
	GetUserByEmail(ctx context.Context, email string) (*user.User, error)
}
